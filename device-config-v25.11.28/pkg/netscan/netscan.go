package netscan

import (
	"bytes"
	"errors"
	"net"
	"sort"
	"strings"
	"sync"
	"time"
)

const (
	localPingTimeout    = 50 * time.Millisecond
	localPingWorkers    = 200
	workerQueueCapacity = 1024
)

// Bosch MAC prefixes
var boschMacPrefixes = []string{
	"00:1A:E8", "00:13:7A", "58:69:6C", "44:8A:5B", "C8:2B:96",
	"00:01:31", "00:04:63", "00:07:5F", "00:10:17", "00:1B:86",
	"00:1C:44", "A4:18:94", "50:FA:CB:9",
}

// Device holds IP and MAC
type Device struct {
	IP  string
	MAC string
}

// -------------------
// Scan network (IP + MAC)
// -------------------
func ScanNetwork() ([]Device, error) {
	cidrs, err := activeInterfaceCIDRs()
	if err != nil || len(cidrs) == 0 {
		cidrs, err = localIPv4CIDRs()
		if err != nil || len(cidrs) == 0 {
			return nil, errors.New("no suitable IPv4 local subnets found")
		}
	}

	var allHosts []string
	for _, c := range cidrs {
		hosts, err := hostsInCIDR(c)
		if err == nil && len(hosts) > 0 {
			allHosts = append(allHosts, hosts...)
		}
	}

	if len(allHosts) == 0 {
		return nil, errors.New("no hosts to scan")
	}

	// Ping hosts to populate ARP table
	pingHostsParallel(allHosts)

	// Read ARP cache
	arpMap := readARPCacheSafe()

	// Build device list with IP + MAC
	var devices []Device
	for ip, mac := range arpMap {
		if IsBoschMAC(mac) {
			devices = append(devices, Device{
				IP:  ip,
				MAC: normalizeMAC(mac),
			})
		}
	}

	// Deduplicate & sort by IP
	unique := map[string]Device{}
	for _, d := range devices {
		unique[d.IP] = d
	}
	out := make([]Device, 0, len(unique))
	for _, d := range unique {
		out = append(out, d)
	}
	sort.Slice(out, func(i, j int) bool { return ipLess(out[i].IP, out[j].IP) })

	return out, nil
}

// -------------------
// Find IP by MAC
// -------------------
func FindIPByMAC(targetMAC string) (string) {
	targetMAC = normalizeMAC(targetMAC)
	devices, err := ScanNetwork()
	if err != nil {
		return ""
	}

	for _, d := range devices {
		if d.MAC == targetMAC {
			return d.IP
		}
	}
	return ""
}

// -------------------
// Bosch MAC utilities
// -------------------
func IsBoschMAC(mac string) bool {
	norm := normalizeMAC(mac)
	for _, prefix := range boschMacPrefixes {
		if strings.HasPrefix(norm, normalizeMAC(prefix)) {
			return true
		}
	}
	return false
}

// -------------------
// Active interface CIDRs
// -------------------
func activeInterfaceCIDRs() ([]string, error) {
	conn, err := net.Dial("udp", "1.1.1.1:80")
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	udpAddr, ok := conn.LocalAddr().(*net.UDPAddr)
	if !ok || udpAddr.IP == nil {
		return nil, errors.New("failed to read local address")
	}
	localIP := udpAddr.IP.To4()
	if localIP == nil {
		return nil, errors.New("active IP is not IPv4")
	}

	ifaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}

	for _, inf := range ifaces {
		if inf.Flags&net.FlagUp == 0 || inf.Flags&net.FlagLoopback != 0 {
			continue
		}
		addrs, _ := inf.Addrs()
		for _, a := range addrs {
			ip, ipnet := parseIPNet(a)
			if ip != nil && ip.Equal(localIP) {
				return []string{ipnet.String()}, nil
			}
		}
	}

	return nil, errors.New("active interface CIDR not found")
}

func localIPv4CIDRs() ([]string, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}
	var cidrs []string
	for _, inf := range ifaces {
		if inf.Flags&net.FlagUp == 0 || inf.Flags&net.FlagLoopback != 0 {
			continue
		}
		addrs, _ := inf.Addrs()
		for _, a := range addrs {
			ip, ipnet := parseIPNet(a)
			if ip != nil && ip.To4() != nil {
				cidr := ipnet.String()
				if !contains(cidrs, cidr) {
					cidrs = append(cidrs, cidr)
				}
			}
		}
	}
	return cidrs, nil
}

func parseIPNet(a net.Addr) (net.IP, *net.IPNet) {
	switch v := a.(type) {
	case *net.IPNet:
		return v.IP, v
	case *net.IPAddr:
		return v.IP, &net.IPNet{IP: v.IP, Mask: net.CIDRMask(32, 32)}
	default:
		return nil, nil
	}
}

func contains(slice []string, val string) bool {
	for _, s := range slice {
		if s == val {
			return true
		}
	}
	return false
}

// -------------------
// Hosts in CIDR
// -------------------
func hostsInCIDR(cidr string) ([]string, error) {
	ip, ipnet, err := net.ParseCIDR(cidr)
	if err != nil {
		return nil, err
	}
	start := ip.Mask(ipnet.Mask).To4()
	if start == nil {
		return nil, errors.New("not IPv4 CIDR")
	}

	mask := net.IP(ipnet.Mask).To4()
	end := make(net.IP, len(start))
	for i, s := range start {
		end[i] = s | (^mask[i])
	}

	var hosts []string
	cur := make(net.IP, len(start))
	copy(cur, start)
	for {
		hosts = append(hosts, cur.String())
		if bytes.Equal(cur, end) {
			break
		}
		incIP(cur)
		if len(cur) == 0 {
			break
		}
	}

	if len(hosts) > 2 {
		hosts = hosts[1 : len(hosts)-1]
	}
	return hosts, nil
}

func incIP(ip net.IP) {
	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] != 0 {
			break
		}
	}
}

// -------------------
// Ping hosts
// -------------------
func pingHostsParallel(hosts []string) {
	jobs := make(chan string, workerQueueCapacity)
	var wg sync.WaitGroup

	for i := 0; i < localPingWorkers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for h := range jobs {
				pingHost(h)
			}
		}()
	}

	for _, h := range hosts {
		jobs <- h
	}
	close(jobs)
	wg.Wait()
}

func pingHost(ip string) {
	for _, port := range []string{"80", "554"} {
		addr := net.JoinHostPort(ip, port)
		if conn, err := net.DialTimeout("tcp", addr, localPingTimeout); err == nil {
			_ = conn.Close()
			return
		}
	}
}