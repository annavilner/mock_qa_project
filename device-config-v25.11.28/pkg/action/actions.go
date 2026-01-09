package action

import (
	"fmt"
	"sync"
	"time"

	cmd "gitlab.cbsdev.net/quality-assurance/device-config/pkg/command"
	o "gitlab.cbsdev.net/quality-assurance/device-config/pkg/output"
)

func (a *Action) ShowDeviceDetails(isFull bool) {
	a.Output.StartSpin()
	a.needWaitForReboot()

	if pwd := *a.Options["ChangeServicePasswordOption"].Value.(*string); pwd != "" {
		a.Request.Device.Password = pwd
	}

	fullText := ""
	if isFull {
		fullText = "full "
	}
	a.Output.SetSpinText(a.Request.Device.IPAddress + " - Getting device " + fullText + "details")

	outputLines := []o.DeviceDetailsLine{
		{Name: "IP Address", Value: a.Request.Device.IPAddress},
	}

	var mu sync.Mutex // protects outputLines

	appendLine := func(line o.DeviceDetailsLine) {
		if line.Name != "" {
			mu.Lock()
			outputLines = append(outputLines, line)
			mu.Unlock()
		}
	}

	// ----------------------------
	// Step 1: Safe concurrent requests (limited concurrency)
	// ----------------------------
	safeFuncs := []func() o.DeviceDetailsLine{
		func() o.DeviceDetailsLine {
			ret := a.Request.GetMacAddress()
			return o.DeviceDetailsLine{Name: cmd.MacAddress.Name, Value: a.Output.MacAddress(ret)}
		},
		func() o.DeviceDetailsLine {
			ret := a.Request.GetName()
			return o.DeviceDetailsLine{Name: cmd.Name.Name, Value: a.Output.Name(ret)}
		},	
	}

	if *a.Options["Password"].Value.(*string) != "" {
		safeFuncs = append(safeFuncs,
			func() o.DeviceDetailsLine {
				ret := a.Request.GetCloudDestination()
				return o.DeviceDetailsLine{Name: cmd.CloudDestination.Name, Value: a.Output.CloudDestination(ret)}
			},
			func() o.DeviceDetailsLine {
				ret := a.Request.GetSocketKnockerDestination()
				return o.DeviceDetailsLine{Name: cmd.SocketKnockerDestination.Name, Value: a.Output.SocketKnockerDestination(ret)}
			},
			func() o.DeviceDetailsLine {
				ret := a.Request.GetSocketKnockerMode()
				return o.DeviceDetailsLine{Name: cmd.SocketKnockerMode.Name, Value: a.Output.SocketKnockerMode(ret)}
			},
			func() o.DeviceDetailsLine {
				ret := a.Request.GetCloudStatus()
				return o.DeviceDetailsLine{Name: cmd.CloudStatus.Name, Value: a.Output.CloudStatus(ret)}
			},
			func() o.DeviceDetailsLine {
				ret := a.Request.GetCSRFProtectionStatus()
				return o.DeviceDetailsLine{Name: cmd.CSRFProtectionStatus.Name, Value: a.Output.CSRFProtectionStatus(ret)}
			},	
		)
	} 

	if *a.Options["Password"].Value.(*string) == "" && isFull{
		safeFuncs = append(safeFuncs,
			func() o.DeviceDetailsLine {
				ret := a.Request.GetFirmwareVersion()
				return o.DeviceDetailsLine{Name: cmd.FirmwareVersion.Name, Value: a.Output.FirmwareVersion(ret)}
			},
		)
	}

	// Limit concurrency to 5
	sem := make(chan struct{}, 5)
	var wg sync.WaitGroup

	for _, f := range safeFuncs {
		wg.Add(1)
		go func(fn func() o.DeviceDetailsLine) {
			defer wg.Done()
			sem <- struct{}{} // acquire slot
			line := fn()
			<-sem // release slot
			appendLine(line)
		}(f)
	}

	wg.Wait()

	// ----------------------------
	// Step 2: Sensitive / sequential requests
	// ----------------------------
	if *a.Options["Password"].Value.(*string) != "" && isFull {
		appendLine(o.DeviceDetailsLine{
			Name:  cmd.SocketKnockerStatusReason.Name,
			Value: a.Output.SocketKnockerStatusAndReason(a.Request.GetSocketKnockerStatusAndReason()),
		})
		appendLine(o.DeviceDetailsLine{
			Name:  cmd.ProductName.Name,
			Value: a.Output.ProductName(a.Request.GetProductName()),
		})
		appendLine(o.DeviceDetailsLine{
			Name:  cmd.VCAProfile.Name,
			Value: a.Output.VCAProfile(a.Request.GetVCAProfile()),
		})
		firmwareVersionLine := o.DeviceDetailsLine{
			Name:  cmd.FirmwareVersionFormatted.Name,
			Value: a.Output.FirmwareVersionFormatted(a.Request.GetFirmwareVersionFormatted()),
		}
		if firmwareVersionLine.Value != "" {
			appendLine(firmwareVersionLine)
		} else {
			appendLine(o.DeviceDetailsLine{
				Name:  cmd.FirmwareVersion.Name,
				Value: a.Output.FirmwareVersion(a.Request.GetFirmwareVersion()),
			})
		}

		// DateTime sequentially
		year := a.Request.GetDateTimeYear()
		month := a.Request.GetDateTimeMonth()
		day := a.Request.GetDateTimeDay()
		hour := a.Request.GetDateTimeHour()
		minute := a.Request.GetDateTimeMinute()
		second := a.Request.GetDateTimeSecond()
		offset := a.Request.GetDateTimeTimezoneOffset()
		appendLine(o.DeviceDetailsLine{
			Name: cmd.DateTime.Name,
			Value: a.Output.DateTimeDay(day) + "/" + a.Output.DateTimeMonth(month) + "/" + a.Output.DateTimeYear(year) +
				" " + a.Output.DateTimeHour(hour) + ":" + a.Output.DateTimeMinute(minute) + ":" + a.Output.DateTimeSecond(second) +
				"  UTC:" + a.Output.DateTimeTimezoneOffset(offset),
		})

		// Capabilities sequentially
		caps := a.Request.GetCapabilites()
		if a.Output.Capabilites(caps) {
			appendLine(o.DeviceDetailsLine{Name: cmd.DeviceType.Name, Value: a.Output.DeviceType()})
			appendLine(o.DeviceDetailsLine{Name: cmd.PTZStatus.Name, Value: a.Output.PTZStatus()})
		}
	}

	a.Output.StopSpin()

	// ----------------------------
	// Print output
	// ----------------------------
	maxNameSize, maxValueSize := a.Output.MaxSizeDeviceDetailsLine(outputLines)
	a.Output.PrintDeviceDetailsDivisionLine(maxNameSize + maxValueSize)
	a.Output.PrintDeviceDetailsLine(outputLines, maxNameSize, maxValueSize)
	a.Output.PrintDeviceDetailsDivisionLine(maxNameSize + maxValueSize)
}

func (a *Action) Reboot() {
	a.Output.StartSpin()
	if (*a.Options["ResetFactoryOption"].Value.(*bool) || *a.Options["ResetOption"].Value.(*bool)) && !a.Request.Device.Rebooted {
		a.waitForReboot()
	}
	a.Output.SetSpinText(cmd.Reboot.WriteAction)
	resp := a.Request.Reboot()
	a.Output.StopSpin()
	code, out := a.Output.ResponseResult(resp, cmd.Reboot.Name)
	a.Output.PrintOutput(out)
	a.Request.Device.Rebooted = true
	if code == "success" {
		a.Request.Device.Rebooted = false
	}
}

func (a *Action) ResetFactory() {
	a.Output.StartSpin()
	if (*a.Options["RebootOption"].Value.(*bool) || *a.Options["ResetOption"].Value.(*bool)) && !a.Request.Device.Rebooted {
		a.waitForReboot()
	}
	a.Output.SetSpinText(cmd.ResetFactory.WriteAction)
	resp := a.Request.ResetFactory()
	a.Output.StopSpin()
	code, out := a.Output.ResponseResult(resp, cmd.ResetFactory.Name)
	a.Output.PrintOutput(out)
	a.Request.Device.Rebooted = true
	if code == "success" {
		a.Request.Device.Rebooted = false
	}
}

func (a *Action) Reset() {
	a.Output.StartSpin()
	if (*a.Options["RebootOption"].Value.(*bool) || *a.Options["ResetFactoryOption"].Value.(*bool)) && !a.Request.Device.Rebooted {
		a.waitForReboot()
	}
	a.Output.SetSpinText(cmd.Reset.WriteAction)
	resp := a.Request.Reset()
	a.Output.StopSpin()
	code, out := a.Output.ResponseResult(resp, cmd.Reset.Name)
	a.Output.PrintOutput(out)
	a.Request.Device.Rebooted = true
	if code == "success" {
		a.Request.Device.Rebooted = false
	}
}

func (a *Action) SetANewServicePassword() {
	a.Request.Device.ServicePassword = *a.Options["Password"].Value.(*string)
	a.Output.StartSpin()
	a.needWaitForReboot()
	a.Output.SetSpinText(cmd.NewServicePassword.WriteAction)
	resp := a.Request.SetANewServicePassword()
	a.Output.StopSpin()
	_, out := a.Output.ResponseResult(resp, cmd.NewServicePassword.Name)
	a.Output.PrintOutput(out)
}

func (a *Action) ChangeServicePassword() {
	a.Request.Device.ServicePassword = *a.Options["ChangeServicePasswordOption"].Value.(*string)
	a.Output.StartSpin()
	a.needWaitForReboot()
	a.Output.SetSpinText(cmd.ServicePassword.WriteAction)
	resp := a.Request.ChangeServicePassword()
	a.Output.StopSpin()
	_, out := a.Output.ResponseResult(resp, cmd.ServicePassword.Name)
	a.Output.PrintOutput(out)
}

func (a *Action) ChangeLivePassword() {
	a.Request.Device.LivePassword = *a.Options["ChangeLivePasswordOption"].Value.(*string)
	a.Output.StartSpin()
	a.needWaitForReboot()
	a.Output.SetSpinText(cmd.LivePassword.WriteAction)
	resp := a.Request.ChangeLivePassword()
	a.Output.StopSpin()
	_, out := a.Output.ResponseResult(resp, cmd.LivePassword.Name)
	a.Output.PrintOutput(out)
}

func (a *Action) ChangeUserPassword() {
	a.Request.Device.UserPassword = *a.Options["ChangeUserPasswordOption"].Value.(*string)
	a.Output.StartSpin()
	a.needWaitForReboot()
	a.Output.SetSpinText(cmd.UserPassword.WriteAction)
	resp := a.Request.ChangeUserPassword()
	a.Output.StopSpin()
	_, out := a.Output.ResponseResult(resp, cmd.UserPassword.Name)
	a.Output.PrintOutput(out)
}

func (a *Action) AddCertificate() {
	a.Output.StartSpin()
	a.needWaitForReboot()
	a.Output.SetSpinText(cmd.UploadCertificate.WriteAction)
	resp := a.Request.UploadCertificate()
	a.Output.StopSpin()
	cd, out := a.Output.ResponseResult(resp, cmd.UploadCertificate.Name)
	a.Output.PrintOutput(out)

	if cd == "success" {
		a.Output.StartSpin()
		a.Output.SetSpinText(cmd.CertificateUsage.WriteAction)
		ta := time.Now().Add(90 * time.Second)
		code := ""
		out := ""
		for {
			tn := time.Now()
			resp = a.Request.SetCertificateUsage()
			code, out = a.Output.ResponseResult(resp, cmd.CertificateUsage.Name)
			if code == "success" || tn.After(ta) {
				break
			}
		}
		a.Output.StopSpin()
		a.Output.PrintOutput(out)
	}
}

func (a *Action) SetSocketKnockerDestination() {
	a.Request.Device.SocketKnockerUrl, a.Request.Device.SocketKnockerPort = checkSocketKnockerDestination(*a.Options["SetSocketKnockerDestinationOption"].Value.(*string))
	a.Output.StartSpin()
	a.needWaitForReboot()
	a.Output.SetSpinText(cmd.SocketKnockerDestination.WriteAction)
	resp := a.Request.SetSocketKnockerDestination()
	a.Output.StopSpin()
	_, out := a.Output.ResponseResult(resp, cmd.SocketKnockerDestination.Name)
	a.Output.PrintOutput(out)
}

func (a *Action) SetCloudDestination() {
	a.Request.Device.CloudUrl, a.Request.Device.CloudPort = checkCloudDestination(*a.Options["SetCloudDestinationOption"].Value.(*string))
	a.Output.StartSpin()
	a.needWaitForReboot()
	a.Output.SetSpinText(cmd.CloudDestination.WriteAction)
	resp := a.Request.SetCloudDestination()
	a.Output.StopSpin()
	_, out := a.Output.ResponseResult(resp, cmd.CloudDestination.Name)
	a.Output.PrintOutput(out)
}

func (a *Action) CloudCommision() {
	a.Request.Device.CloudEmail, a.Request.Device.CloudPassword = checkCloudCommission(*a.Options["CloudCommissionOption"].Value.(*string))
	a.Output.StartSpin()
	a.needWaitForReboot()
	a.Output.SetSpinText(cmd.CloudCommission.WriteAction)
	resp := a.Request.SetCloudCommission()
	cd, out := a.Output.ResponseResult(resp, cmd.CloudCommission.Name)
	a.Output.StopSpin()
	a.Output.PrintOutput(out)

	if cd == "success" {
		a.Output.StartSpin()
		ta := time.Now().Add(90 * time.Second)
		reg := false
		for {
			tn := time.Now()
			a.Output.SetSpinText("Waiting for the device to be commissioned in the cloud")
			out := a.Request.GetCloudStatus()
			if a.Output.CloudStatus(out) == "Registered" {
				reg = true
				break
			} else if tn.After(ta) {
				break
			}
		}
		a.Output.StopSpin()
		if reg {
			fmt.Println(a.Output.CommissionSuccessMsg())
		} else {
			fmt.Println(a.Output.CommissionErrorMsg())
		}
	}
}

func (a *Action) SetName() {
	a.Request.Device.Name = *a.Options["SetNameOption"].Value.(*string)
	a.Output.StartSpin()
	a.needWaitForReboot()
	a.Output.SetSpinText(cmd.Name.WriteAction)
	resp := a.Request.SetName()
	a.Output.StopSpin()
	_, out := a.Output.ResponseResult(resp, cmd.Name.Name)
	a.Output.PrintOutput(out)
}

func (a *Action) SetSocketKnockerMode() {
	a.Request.Device.SocketKnockerMode = checkSocketKnockerMode(*a.Options["SetSocketKnockerModeOption"].Value.(*string))
	a.Output.StartSpin()
	a.needWaitForReboot()
	a.Output.SetSpinText(cmd.SocketKnockerMode.WriteAction)
	resp := a.Request.SetSocketKnockerMode()
	a.Output.StopSpin()
	_, out := a.Output.ResponseResult(resp, cmd.SocketKnockerMode.Name)
	a.Output.PrintOutput(out)
}

func (a *Action) SyncDateTimeToPC() {
	a.Request.Device.DateTime = getPCDateTime()
	a.Output.StartSpin()
	a.needWaitForReboot()
	a.Output.SetSpinText(cmd.DateTime.WriteAction)
	resp := ""

	resp = a.Request.SetDateTimeFormat()
	_, _ = a.Output.ResponseResult(resp, cmd.DateTime.Name)

	resp = a.Request.SetDateTimeTimezoneOffset()
	_, _ = a.Output.ResponseResult(resp, cmd.DateTime.Name)

	fSecond := async(func() string {
		return a.Request.SetDateTimeSecond()
	})
	fMinute := async(func() string {
		return a.Request.SetDateTimeMinute()
	})
	fHour := async(func() string {
		return a.Request.SetDateTimeHour()
	})
	fDay := async(func() string {
		return a.Request.SetDateTimeDay()
	})
	fMonth := async(func() string {
		return a.Request.SetDateTimeMonth()
	})
	fYear := async(func() string {
		return a.Request.SetDateTimeYear()
	})

	resp = a.Request.SetDateTimeOverwriteByDHCP()
	_, out := a.Output.ResponseResult(resp, cmd.DateTime.Name)

	<-fSecond
	<-fMinute
	<-fHour
	<-fDay
	<-fMonth
	<-fYear

	a.Output.StopSpin()
	a.Output.PrintOutput(out)
}

func (a *Action) SetVCAProfile() {
	a.Request.Device.VCAProfile = checkVCAProfile(*a.Options["SetVCAProfileOption"].Value.(*string))
	a.Output.StartSpin()
	a.needWaitForReboot()
	a.Output.SetSpinText(cmd.VCAProfile.WriteAction)
	resp := a.Request.SetVCAProfile()
	a.Output.StopSpin()
	_, out := a.Output.ResponseResult(resp, cmd.VCAProfile.Name)
	a.Output.PrintOutput(out)
}

func (a *Action) SetRelayOutputOn() {
	a.Request.Device.RelayOutputOn = checkOutputNumber(*a.Options["SetRelayOutputOnOption"].Value.(*string))
	a.Output.StartSpin()
	a.needWaitForReboot()
	a.Output.SetSpinText(cmd.RelayOutputOn.WriteAction)
	resp := a.Request.SetRelayOutputOn()
	a.Output.StopSpin()
	_, out := a.Output.ResponseResult(resp, cmd.RelayOutputOn.Name)
	a.Output.PrintOutput(out)
}

func (a *Action) SetRelayOutputOff() {
	a.Request.Device.RelayOutputOff = checkOutputNumber(*a.Options["SetRelayOutputOffOption"].Value.(*string))
	a.Output.StartSpin()
	a.needWaitForReboot()
	a.Output.SetSpinText(cmd.RelayOutputOff.WriteAction)
	resp := a.Request.SetRelayOutputOff()
	a.Output.StopSpin()
	_, out := a.Output.ResponseResult(resp, cmd.RelayOutputOff.Name)
	a.Output.PrintOutput(out)
}

func (a *Action) SetRelayInputOn() {
	a.Request.Device.RelayInputOn = checkOutputNumber(*a.Options["SetRelayInputOnOption"].Value.(*string))
	a.Output.StartSpin()
	a.needWaitForReboot()
	a.Output.SetSpinText(cmd.RelayInputOn.WriteAction)
	resp := a.Request.SetRelayInputOn()
	a.Output.StopSpin()
	_, out := a.Output.ResponseResult(resp, cmd.RelayInputOn.Name)
	a.Output.PrintOutput(out)
}

func (a *Action) SetRelayInputOff() {
	a.Request.Device.RelayInputOff = checkOutputNumber(*a.Options["SetRelayInputOffOption"].Value.(*string))
	a.Output.StartSpin()
	a.needWaitForReboot()
	a.Output.SetSpinText(cmd.RelayInputOff.WriteAction)
	resp := a.Request.SetRelayInputOff()
	a.Output.StopSpin()
	_, out := a.Output.ResponseResult(resp, cmd.RelayInputOff.Name)
	a.Output.PrintOutput(out)
}

func (a *Action) SetAudioOn() {
	a.Output.StartSpin()
	a.needWaitForReboot()
	a.Output.SetSpinText(cmd.AudioOn.WriteAction)
	resp := a.Request.SetAudioOn()
	a.Output.StopSpin()
	_, out := a.Output.ResponseResult(resp, cmd.AudioOn.Name)
	a.Output.PrintOutput(out)
}

func (a *Action) SetAudioOff() {
	a.Output.StartSpin()
	a.needWaitForReboot()
	a.Output.SetSpinText(cmd.AudioOff.WriteAction)
	resp := a.Request.SetAudioOff()
	a.Output.StopSpin()
	_, out := a.Output.ResponseResult(resp, cmd.AudioOff.Name)
	a.Output.PrintOutput(out)
}

func (a *Action) SetCSRFProtectionOn() {
	a.Output.StartSpin()
	a.needWaitForReboot()
	a.Output.SetSpinText(cmd.CSRFProtectionOn.WriteAction)
	resp := a.Request.SetCSRFProtectionOn()
	a.Output.StopSpin()
	_, out := a.Output.ResponseResult(resp, cmd.CSRFProtectionOn.Name)
	a.Output.PrintOutput(out)
}

func (a *Action) SetCSRFProtectionOff() {
	a.Output.StartSpin()
	a.needWaitForReboot()
	a.Output.SetSpinText(cmd.CSRFProtectionOff.WriteAction)
	resp := a.Request.SetCSRFProtectionOff()
	a.Output.StopSpin()
	_, out := a.Output.ResponseResult(resp, cmd.CSRFProtectionOff.Name)
	a.Output.PrintOutput(out)
}
