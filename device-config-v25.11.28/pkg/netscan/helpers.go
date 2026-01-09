package netscan

import (
	"bufio"
	"net"
	"os"
	"os/exec"
	"regexp"
	"runtime"
	"strings"
)

func readARPCacheSafe() map[string]string {
	out := map[string]string{}
	if runtime.GOOS == "linux" {
		if f, err := os.Open("/proc/net/arp"); err == nil {
			defer f.Close()
			scanner := bufio.NewScanner(f)
			first := true
			for scanner.Scan() {
				line := scanner.Text()
				if first {
					first = false
					continue
				}
				fields := strings.Fields(line)
				if len(fields) >= 4 {
					ip := fields[0]
					mac := fields[3]
					if mac != "00:00:00:00:00:00" && mac != "" {
						out[ip] = mac
					}
				}
			}
			return out
		}
	}
	// fallback: arp -a
	cmd := exec.Command("arp", "-a")
	b, err := cmd.CombinedOutput()
	if err != nil {
		return out
	}
	text := string(b)
	ipRe := regexp.MustCompile(`\d+\.\d+\.\d+\.\d+`)
	macRe := regexp.MustCompile(`(?i)([0-9a-f]{2}(?::|-)){5}[0-9a-f]{2}`)
	scanner := bufio.NewScanner(strings.NewReader(text))
	for scanner.Scan() {
		line := scanner.Text()
		ip := ""
		mac := ""
		if m := ipRe.FindString(line); m != "" {
			ip = m
		}
		if mm := macRe.FindString(line); mm != "" {
			mac = strings.ReplaceAll(mm, "-", ":")
		}
		if ip != "" && mac != "" {
			out[ip] = strings.ToLower(mac)
		}
	}
	return out
}

func normalizeMAC(m string) string {
	m = strings.TrimSpace(strings.ToLower(m))
	m = strings.ReplaceAll(m, "-", ":")
	parts := strings.Split(m, ":")
	if len(parts) == 6 {
		for i := range parts {
			if len(parts[i]) == 1 {
				parts[i] = "0" + parts[i]
			}
		}
		return strings.Join(parts, ":")
	}
	return m
}

func ipLess(a, b string) bool {
	ipa := net.ParseIP(a).To4()
	ipb := net.ParseIP(b).To4()
	if ipa == nil || ipb == nil {
		return a < b
	}
	for i, va := range ipa {
		if vb := ipb[i]; va != vb {
			return va < vb
		}
	}
	return false
}
