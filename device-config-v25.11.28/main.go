package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/briandowns/spinner"
	ac "gitlab.cbsdev.net/quality-assurance/device-config/pkg/action"
	"gitlab.cbsdev.net/quality-assurance/device-config/pkg/netscan"
	op "gitlab.cbsdev.net/quality-assurance/device-config/pkg/option"
	o "gitlab.cbsdev.net/quality-assurance/device-config/pkg/output"
)

const (
	rpApiDomain = "cbs.keen.tech"
	dmDomain    = "cbs.boschsecurity.com"
)

func main() {

	errors := []string{}
	a := ac.Action{
		Output: o.Output{
			Spin: spinner.New([]string{
				"🌑",
				"🌒",
				"🌓",
				"🌔",
				"🌕",
				"🌔",
				"🌓",
				"🌒",
			}, 200*time.Millisecond, spinner.WithWriter(os.Stderr), spinner.WithColor("green")),
		},
		Options: make(map[string]ac.Option),
	}

	for _, opt := range op.AllOptions() {
		switch opt.Type {
		case "string":
			a.Options[opt.Name] = ac.Option{
				Value: flag.String(opt.Flag, "", opt.Description),
			}
		case "bool":
			a.Options[opt.Name] = ac.Option{
				Value: flag.Bool(opt.Flag, false, opt.Description),
			}
		}
	}

	flag.Usage = func() {
		a.Output.Usage()
	}

	flag.Parse()

	if *a.Options["NewServicePasswordOption"].Value.(*bool) && *a.Options["Password"].Value.(*string) == "" && (!*a.Options["HelpOption"].Value.(*bool) && !*a.Options["VersionOption"].Value.(*bool)) {
		errors = append(errors, "You must inform the service password using the option '-p'")
	}

	socketKnockerMode := strings.ToLower(*a.Options["SetSocketKnockerModeOption"].Value.(*string))
	a.Options["SetSocketKnockerModeOption"] = ac.Option{Value: &socketKnockerMode}
	if *a.Options["SetSocketKnockerModeOption"].Value.(*string) != "" && *a.Options["SetSocketKnockerModeOption"].Value.(*string) != "on" &&
		*a.Options["SetSocketKnockerModeOption"].Value.(*string) != "off" && *a.Options["SetSocketKnockerModeOption"].Value.(*string) != "auto" && (!*a.Options["HelpOption"].Value.(*bool) && !*a.Options["VersionOption"].Value.(*bool)) {
		errors = append(errors, "Invalid Socket Knocker Mode. Available modes: on, off, auto")
	}

	a.Output.SimpleOutput = *a.Options["SimpleOutputOption"].Value.(*bool)
	a.Output.NoOutput = *a.Options["NoOutputOption"].Value.(*bool)

	if len(errors) > 0 {

		for i := 0; i < len(errors); i++ {
			fmt.Print(a.Output.ErrorMsg(errors[i]))
		}
		a.Output.Usage()
		return

	} else {

		if *a.Options["HelpOption"].Value.(*bool) {
			a.Output.Usage()
			return
		}

		if *a.Options["VersionOption"].Value.(*bool) {
			a.Output.Version()
			return
		}

		var ipList []string

		if *a.Options["IPAddress"].Value.(*string) == "" && *a.Options["MACAddress"].Value.(*string) == "" {

			a.Output.StartSpin()
			a.Output.SetSpinText("Scanning for devices")
			devices, err := netscan.ScanNetwork()
			a.Output.StopSpin()
			if err != nil {
				fmt.Print(a.Output.ErrorMsg(err.Error()))
			}
			for _, d := range devices {
				if d.IP != "" {
					ipList = append(ipList, d.IP)
				}
			}
			devicePass := *a.Options["Password"].Value.(*string)
			noOutput := *a.Options["NoOutputOption"].Value.(*bool)
			ddOption := *a.Options["DeviceDetailsOption"].Value.(*bool)
			ffOption := *a.Options["FullDeviceDetailsOption"].Value.(*bool)

			if len(ipList) == 0 {
				fmt.Println(a.Output.ErrorMsg("No Devices Found"))
				return
			}

			for _, ip := range ipList {

				a.Request.Device.IPAddress = ip
				a.Request.Device.Password = devicePass

				// Show Device Details
				if ddOption && !noOutput {
					a.ShowDeviceDetails(false)
				}

				// Show Full Device Details
				if ffOption && !noOutput {
					a.ShowDeviceDetails(true)
				}
			}

		} else {
			ipAddress := ""
			if *a.Options["IPAddress"].Value.(*string) != "" {
				ipAddress = *a.Options["IPAddress"].Value.(*string)
			} else if *a.Options["IPAddress"].Value.(*string) == "" && *a.Options["MACAddress"].Value.(*string) != "" {
				a.Output.StartSpin()
				a.Output.SetSpinText("Scanning for device")
				ipAddress = netscan.FindIPByMAC(*a.Options["MACAddress"].Value.(*string))

				if ipAddress == "" {
					fmt.Println(a.Output.ErrorMsg("Device Not Found"))
					return
				}
			}

			a.Request.Device.IPAddress = ipAddress
			a.Request.Device.Password = *a.Options["Password"].Value.(*string)

			// Prepare to RP 'remoqa'
			if *a.Options["PrepareToRPRemoqaOption"].Value.(*bool) {
				*a.Options["AddCertificateOption"].Value.(*bool) = true
				*a.Options["SetSocketKnockerDestinationOption"].Value.(*string) = "dm201." + dmDomain + ":42000"
				*a.Options["SetCloudDestinationOption"].Value.(*string) = "api.remoqa." + rpApiDomain + ":443"
				*a.Options["SetSocketKnockerModeOption"].Value.(*string) = "on"
				*a.Options["SetCSRFProtectionOffOption"].Value.(*bool) = true
				*a.Options["SyncDateTimeToPC"].Value.(*bool) = true
			}

			// Prepare to RP 'remodevnew'
			if *a.Options["PrepareToRPRemodevnewOption"].Value.(*bool) {
				*a.Options["AddCertificateOption"].Value.(*bool) = true
				*a.Options["SetSocketKnockerDestinationOption"].Value.(*string) = "dm207." + dmDomain + ":42000"
				*a.Options["SetCloudDestinationOption"].Value.(*string) = "api.remodevnew." + rpApiDomain + ":443"
				*a.Options["SetSocketKnockerModeOption"].Value.(*string) = "on"
				*a.Options["SetCSRFProtectionOffOption"].Value.(*bool) = true
				*a.Options["SyncDateTimeToPC"].Value.(*bool) = true
			}

			// Prepare to AM 'test'
			if *a.Options["PrepareToAMTestOption"].Value.(*bool) {
				*a.Options["AddCertificateOption"].Value.(*bool) = true
				*a.Options["SetSocketKnockerDestinationOption"].Value.(*string) = "dm201." + dmDomain + ":42000"
				*a.Options["SetCloudDestinationOption"].Value.(*string) = "api.remoqa." + rpApiDomain + ":443"
				*a.Options["SetSocketKnockerModeOption"].Value.(*string) = "on"
				*a.Options["SetCSRFProtectionOffOption"].Value.(*bool) = true
				*a.Options["SyncDateTimeToPC"].Value.(*bool) = true
			}

			// Prepare to AM 'btuhtest'
			if *a.Options["PrepareToAMBtuhtestOption"].Value.(*bool) {
				*a.Options["AddCertificateOption"].Value.(*bool) = true
				*a.Options["SetSocketKnockerDestinationOption"].Value.(*string) = "dm207." + dmDomain + ":42000"
				*a.Options["SetCloudDestinationOption"].Value.(*string) = "api.remodevnew." + rpApiDomain + ":443"
				*a.Options["SetSocketKnockerModeOption"].Value.(*string) = "on"
				*a.Options["SetCSRFProtectionOffOption"].Value.(*bool) = true
				*a.Options["SyncDateTimeToPC"].Value.(*bool) = true
			}

			// Prepare to AM 'dev'
			if *a.Options["PrepareToAMDevOption"].Value.(*bool) {
				*a.Options["AddCertificateOption"].Value.(*bool) = true
				*a.Options["SetSocketKnockerDestinationOption"].Value.(*string) = "dm207." + dmDomain + ":42000"
				*a.Options["SetCloudDestinationOption"].Value.(*string) = "api.remodevnew." + rpApiDomain + ":443"
				*a.Options["SetSocketKnockerModeOption"].Value.(*string) = "on"
				*a.Options["SetCSRFProtectionOffOption"].Value.(*bool) = true
				*a.Options["SyncDateTimeToPC"].Value.(*bool) = true
			}

			// Prepare to AM 'demo'
			if *a.Options["PrepareToAMDemoOption"].Value.(*bool) {
				*a.Options["AddCertificateOption"].Value.(*bool) = true
				*a.Options["SetSocketKnockerDestinationOption"].Value.(*string) = "dm201.cbs.boschsecurity.com:42000"
				*a.Options["SetCloudDestinationOption"].Value.(*string) = "api.remote.boschsecurity.com:443"
				*a.Options["SetSocketKnockerModeOption"].Value.(*string) = "on"
				*a.Options["SetCSRFProtectionOffOption"].Value.(*bool) = true
				*a.Options["SyncDateTimeToPC"].Value.(*bool) = true
			}

			// Reboot
			if *a.Options["RebootOption"].Value.(*bool) {
				a.Reboot()
			}

			// Reset Factory
			if *a.Options["ResetFactoryOption"].Value.(*bool) {
				a.ResetFactory()
			}

			// Reset
			if *a.Options["ResetOption"].Value.(*bool) {
				a.Reset()
			}

			// Set a New Service Password
			if *a.Options["NewServicePasswordOption"].Value.(*bool) {
				a.SetANewServicePassword()
			}

			// Change Service Password
			if *a.Options["ChangeServicePasswordOption"].Value.(*string) != "" {
				a.ChangeServicePassword()
			}

			// Change Live Password
			if *a.Options["ChangeLivePasswordOption"].Value.(*string) != "" {
				a.ChangeLivePassword()
			}

			// Change User Password
			if *a.Options["ChangeUserPasswordOption"].Value.(*string) != "" {
				a.ChangeUserPassword()
			}

			// Add Certificate
			if *a.Options["AddCertificateOption"].Value.(*bool) {
				a.AddCertificate()
			}

			// Set Socket Knocker Destination
			if *a.Options["SetSocketKnockerDestinationOption"].Value.(*string) != "" {
				a.SetSocketKnockerDestination()
			}

			// Set Cloud Destination
			if *a.Options["SetCloudDestinationOption"].Value.(*string) != "" {
				a.SetCloudDestination()
			}

			// Set Name
			if *a.Options["SetNameOption"].Value.(*string) != "" {
				a.SetName()
			}

			// Set Socket Knocker Mode
			if *a.Options["SetSocketKnockerModeOption"].Value.(*string) != "" {
				a.SetSocketKnockerMode()
			}

			// Sync Date/Time to PC
			if *a.Options["SyncDateTimeToPC"].Value.(*bool) {
				a.SyncDateTimeToPC()
			}

			// Set CSRF Protection On
			if *a.Options["SetCSRFProtectionOnOption"].Value.(*bool) {
				a.SetCSRFProtectionOn()
			}

			// Set CSRF Protection Off
			if *a.Options["SetCSRFProtectionOffOption"].Value.(*bool) {
				a.SetCSRFProtectionOff()
			}

			// Cloud Commission
			if *a.Options["CloudCommissionOption"].Value.(*string) != "" {
				a.CloudCommision()
			}

			// Set VCA Profile
			if *a.Options["SetVCAProfileOption"].Value.(*string) != "" {
				a.SetVCAProfile()
			}

			// Set Relay Output On
			if *a.Options["SetRelayOutputOnOption"].Value.(*string) != "" {
				a.SetRelayOutputOn()
			}

			// Set Relay Output Off
			if *a.Options["SetRelayOutputOffOption"].Value.(*string) != "" {
				a.SetRelayOutputOff()
			}

			// Set Relay Input On
			if *a.Options["SetRelayInputOnOption"].Value.(*string) != "" {
				a.SetRelayInputOn()
			}

			// Set Relay Input Off
			if *a.Options["SetRelayInputOffOption"].Value.(*string) != "" {
				a.SetRelayInputOff()
			}

			// Set Audio On
			if *a.Options["SetAudioOnOption"].Value.(*bool) {
				a.SetAudioOn()
			}

			// Set Audio Off
			if *a.Options["SetAudioOffOption"].Value.(*bool) {
				a.SetAudioOff()
			}

			// Show Device Details
			if !*a.Options["FullDeviceDetailsOption"].Value.(*bool) && *a.Options["DeviceDetailsOption"].Value.(*bool) && !*a.Options["NoOutputOption"].Value.(*bool) {
				a.ShowDeviceDetails(false)
			}

			// Show Full Device Details
			if *a.Options["FullDeviceDetailsOption"].Value.(*bool) && !*a.Options["NoOutputOption"].Value.(*bool) {
				a.ShowDeviceDetails(true)
			}
		}

	}

}
