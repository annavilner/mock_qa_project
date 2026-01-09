package action

import (
	"strconv"
	"time"

	"strings"

	d "gitlab.cbsdev.net/quality-assurance/device-config/pkg/device"
)

func checkSocketKnockerDestination(sk string) (url string, port int64) {
	url = ""
	port = 0
	socketKnocker := strings.Split(sk, ":")
	url = socketKnocker[0]
	if len(socketKnocker) == 2 {
		var err error
		port, err = strconv.ParseInt(socketKnocker[1], 10, 64)
		if err != nil {
			port = 42000
		}
	}
	if port == 0 {
		port = 42000
	}
	if url == "" {
		url = "dm201.cbs.boschsecurity.com"
	}
	return
}

func checkCloudDestination(cd string) (url string, port int64) {
	url = ""
	port = 0
	cloudDestination := strings.Split(cd, ":")
	url = cloudDestination[0]
	if len(cloudDestination) == 2 {
		var err error
		port, err = strconv.ParseInt(cloudDestination[1], 10, 64)
		if err != nil {
			port = 443
		}
	}
	if port == 0 {
		port = 443
	}
	if url == "" {
		url = "api.remoqa.cbs.boschsecurity.com"
	}
	return
}

func checkCloudCommission(cc string) (email string, password string) {
	email = ""
	password = ""
	cloudCommission := strings.Split(cc, ":")
	email = cloudCommission[0]
	if len(cloudCommission) == 2 {
		password = cloudCommission[1]
	}
	return
}

func checkSocketKnockerMode(sn string) (socketKnockerMode string) {
	socketKnockerMode = ""
	switch sn {
	case "on":
		socketKnockerMode = "1"
	case "off":
		socketKnockerMode = "0"
	case "auto":
		socketKnockerMode = "2"
	}
	return
}

func checkVCAProfile(profile string) (vcaProfile string) {
	vcaProfile = ""
	switch profile {
	case "silent":
		vcaProfile = "0"
	case "off":
		vcaProfile = "253"
	case "scheduler":
		vcaProfile = "254"
	case "script":
		vcaProfile = "255"
	default:
		num, err := strconv.Atoi(profile)
		if err == nil && num >= 1 && num <= 255 {
			vcaProfile = profile
		} else {
			vcaProfile = "0"
		}
	}
	return
}

func getPCDateTime() (ddt d.DeviceDateTime) {
	localTime := time.Now()
	_, offset := localTime.Zone()
	ddt.Year = strconv.Itoa(localTime.Year())
	ddt.Month = getMonth(localTime.Month())
	ddt.Day = strconv.Itoa(localTime.Day())
	ddt.Hour = strconv.Itoa(localTime.Hour())
	ddt.Minute = strconv.Itoa(localTime.Minute())
	ddt.Second = strconv.Itoa(localTime.Second())
	ddt.TimezoneOffset = offset

	if localTime.IsDST() {
		ddt.TimezoneOffset = offset - 3600
	}

	return
}

func checkOutputNumber(outputToCheck string) (output string) {
	num, err := strconv.Atoi(outputToCheck)
	if err == nil {
		output = strconv.Itoa(num)
	} else {
		output = "0"
	}
	return
}

func async[T any](f func() T) chan T {
	ch := make(chan T)
	go func() {
		ch <- f()
	}()
	return ch
}

func asyncNoReturn(f func()) {
	go f()
}

func getMonth(month time.Month) string {
	if int(month) < 10 {
		return "0" + strconv.Itoa(int(month))
	} else {
		return strconv.Itoa(int(month))
	}
}

func (a *Action) waitForReboot() {
	for i := 90; i > 0; i-- {
		a.Output.SetSpinText("Waiting for the device to reboot " + strconv.Itoa(i) + "s")
		time.Sleep(1 * time.Second)
	}
	a.Request.Device.Rebooted = true
}

func (a *Action) needWaitForReboot() {
	if (*a.Options["RebootOption"].Value.(*bool) || *a.Options["ResetFactoryOption"].Value.(*bool) || *a.Options["ResetOption"].Value.(*bool)) && !a.Request.Device.Rebooted {
		a.waitForReboot()
	}
}