package request

import (
	"strconv"

	cmd "gitlab.cbsdev.net/quality-assurance/device-config/pkg/command"
	dir "gitlab.cbsdev.net/quality-assurance/device-config/pkg/direction"
	mt "gitlab.cbsdev.net/quality-assurance/device-config/pkg/method"
	rcp "gitlab.cbsdev.net/quality-assurance/device-config/pkg/rcp"
)

func (r *Request) Reboot() (resp string) {
	req := rcp.Request{
		IPAddress: r.Device.IPAddress,
		Password:  r.Device.Password,
		Method:    mt.Get,
		Command:   cmd.Reboot.Code,
		Type:      cmd.Reboot.Type,
		Direction: dir.Write,
	}
	resp = req.Send()
	return
}

func (r *Request) ResetFactory() (resp string) {
	req := rcp.Request{
		IPAddress: r.Device.IPAddress,
		Password:  r.Device.Password,
		Method:    mt.Get,
		Command:   cmd.ResetFactory.Code,
		Type:      cmd.ResetFactory.Type,
		Direction: dir.Write,
	}
	resp = req.Send()
	return
}

func (r *Request) Reset() (resp string) {
	req := rcp.Request{
		IPAddress: r.Device.IPAddress,
		Password:  r.Device.Password,
		Method:    mt.Get,
		Command:   cmd.Reset.Code,
		Type:      cmd.Reset.Type,
		Direction: dir.Write,
	}
	resp = req.Send()
	return
}

func (r *Request) SetANewServicePassword() (resp string) {
	req := rcp.Request{
		IPAddress: r.Device.IPAddress,
		Password:  "",
		Method:    mt.Get,
		Command:   cmd.NewServicePassword.Code,
		Type:      cmd.NewServicePassword.Type,
		Direction: dir.Write,
		Payload:   r.Device.ServicePassword,
		Num:       "2",
	}
	resp = req.Send()

	return
}

func (r *Request) ChangeServicePassword() (resp string) {
	req := rcp.Request{
		IPAddress: r.Device.IPAddress,
		Password:  r.Device.Password,
		Method:    mt.Get,
		Command:   cmd.ServicePassword.Code,
		Type:      cmd.ServicePassword.Type,
		Direction: dir.Write,
		Payload:   r.Device.ServicePassword,
		Num:       "2",
	}
	resp = req.Send()
	return
}

func (r *Request) ChangeLivePassword() (resp string) {
	req := rcp.Request{
		IPAddress: r.Device.IPAddress,
		Password:  r.Device.Password,
		Method:    mt.Get,
		Command:   cmd.LivePassword.Code,
		Type:      cmd.LivePassword.Type,
		Direction: dir.Write,
		Payload:   r.Device.LivePassword,
		Num:       "3",
	}
	resp = req.Send()
	return
}

func (r *Request) ChangeUserPassword() (resp string) {
	req := rcp.Request{
		IPAddress: r.Device.IPAddress,
		Password:  r.Device.Password,
		Method:    mt.Get,
		Command:   cmd.UserPassword.Code,
		Type:      cmd.UserPassword.Type,
		Direction: dir.Write,
		Payload:   r.Device.UserPassword,
		Num:       "1",
	}
	resp = req.Send()
	return
}

func (r *Request) SetCertificateUsage() (resp string) {
	req := rcp.Request{
		IPAddress: r.Device.IPAddress,
		Password:  r.Device.Password,
		Method:    mt.Get,
		Command:   cmd.CertificateUsage.Code,
		Type:      cmd.CertificateUsage.Type,
		Direction: dir.Write,
		Payload:   "0x0008000080000004001300016366673A4874747073536572766572001700017374617469633A636273547275737443657274",
		Num:       "1",
	}
	resp = req.Send()
	return
}

func (r *Request) SetSocketKnockerDestination() (resp string) {
	req := rcp.Request{
		IPAddress: r.Device.IPAddress,
		Password:  r.Device.Password,
		Method:    mt.Get,
		Command:   cmd.SocketKnockerDestination.Code,
		Type:      cmd.SocketKnockerDestination.Type,
		Direction: dir.Write,
		Payload:   "0x" + intToHex(r.Device.SocketKnockerPort) + "010100000000" + textToHex(r.Device.SocketKnockerUrl) + "00",
		Num:       "1",
	}
	resp = req.Send()
	return
}

func (r *Request) SetCloudDestination() (resp string) {
	req := rcp.Request{
		IPAddress: r.Device.IPAddress,
		Password:  r.Device.Password,
		Method:    mt.Get,
		Command:   cmd.CloudDestination.Code,
		Type:      cmd.CloudDestination.Type,
		Direction: dir.Write,
		Payload:   "0x" + intToHex(r.Device.CloudPort) + "010000000000" + textToHex(r.Device.CloudUrl) + "00",
		Num:       "1",
	}

	resp = req.Send()
	return
}

func (r *Request) SetCloudCommission() (resp string) {
	req := rcp.Request{
		IPAddress: r.Device.IPAddress,
		Password:  r.Device.Password,
		Method:    mt.Get,
		Command:   cmd.CloudCommission.Code,
		Type:      cmd.CloudCommission.Type,
		Direction: dir.Write,
		Payload:   "0x01" + textToHex(r.Device.CloudEmail) + "00" + textToHex(r.Device.CloudPassword) + "00",
		Num:       "1",
	}

	resp = req.Send()
	return
}

func (r *Request) SetName() (resp string) {
	req := rcp.Request{
		IPAddress: r.Device.IPAddress,
		Password:  r.Device.Password,
		Method:    mt.Get,
		Command:   cmd.Name.Code,
		Type:      cmd.Name.Type,
		Direction: dir.Write,
		Payload:   "0x" + nameToHex(r.Device.Name) + "0000",
		Num:       "1",
	}

	resp = req.Send()
	return
}

func (r *Request) SetSocketKnockerMode() (resp string) {
	req := rcp.Request{
		IPAddress: r.Device.IPAddress,
		Password:  r.Device.Password,
		Method:    mt.Get,
		Command:   cmd.SocketKnockerMode.Code,
		Type:      cmd.SocketKnockerMode.Type,
		Direction: dir.Write,
		Payload:   r.Device.SocketKnockerMode,
	}

	resp = req.Send()
	return
}

func (r *Request) SetDateTimeFormat() (resp string) {
	req := rcp.Request{
		IPAddress: r.Device.IPAddress,
		Password:  r.Device.Password,
		Method:    mt.Get,
		Command:   cmd.DateTimeFormat.Code,
		Type:      cmd.DateTimeFormat.Type,
		Direction: dir.Write,
		Num:       "1",
		IdString:  "selFormat",
		Payload:   "1",
	}
	resp = req.Send()
	return
}

func (r *Request) SetDateTimeOverwriteByDHCP() (resp string) {
	req := rcp.Request{
		IPAddress: r.Device.IPAddress,
		Password:  r.Device.Password,
		Method:    mt.Get,
		Command:   cmd.DateTimeOverwriteByDHCP.Code,
		Type:      cmd.DateTimeOverwriteByDHCP.Type,
		Direction: dir.Write,
		Num:       "1",
		Payload:   "0",
	}
	resp = req.Send()
	return
}

func (r *Request) SetDateTimeYear() (resp string) {
	req := rcp.Request{
		IPAddress: r.Device.IPAddress,
		Password:  r.Device.Password,
		Method:    mt.Get,
		Command:   cmd.DateTimeYear.Code,
		Type:      cmd.DateTimeYear.Type,
		Direction: dir.Write,
		Num:       "1",
		IdString:  "date3",
		Payload:   r.Device.DateTime.Year,
	}
	resp = req.Send()
	return
}

func (r *Request) SetDateTimeMonth() (resp string) {
	req := rcp.Request{
		IPAddress: r.Device.IPAddress,
		Password:  r.Device.Password,
		Method:    mt.Get,
		Command:   cmd.DateTimeMonth.Code,
		Type:      cmd.DateTimeMonth.Type,
		Direction: dir.Write,
		Num:       "1",
		IdString:  "date2",
		Payload:   r.Device.DateTime.Month,
	}
	resp = req.Send()
	return
}

func (r *Request) SetDateTimeDay() (resp string) {
	req := rcp.Request{
		IPAddress: r.Device.IPAddress,
		Password:  r.Device.Password,
		Method:    mt.Get,
		Command:   cmd.DateTimeDay.Code,
		Type:      cmd.DateTimeDay.Type,
		Direction: dir.Write,
		Num:       "1",
		IdString:  "date1",
		Payload:   r.Device.DateTime.Day,
	}
	resp = req.Send()
	return
}

func (r *Request) SetDateTimeHour() (resp string) {
	req := rcp.Request{
		IPAddress: r.Device.IPAddress,
		Password:  r.Device.Password,
		Method:    mt.Get,
		Command:   cmd.DateTimeHour.Code,
		Type:      cmd.DateTimeHour.Type,
		Direction: dir.Write,
		Num:       "1",
		IdString:  "hour",
		Payload:   r.Device.DateTime.Hour,
	}
	resp = req.Send()
	return
}

func (r *Request) SetDateTimeMinute() (resp string) {
	req := rcp.Request{
		IPAddress: r.Device.IPAddress,
		Password:  r.Device.Password,
		Method:    mt.Get,
		Command:   cmd.DateTimeMinute.Code,
		Type:      cmd.DateTimeMinute.Type,
		Direction: dir.Write,
		Num:       "1",
		IdString:  "minute",
		Payload:   r.Device.DateTime.Minute,
	}
	resp = req.Send()
	return
}

func (r *Request) SetDateTimeSecond() (resp string) {
	req := rcp.Request{
		IPAddress: r.Device.IPAddress,
		Password:  r.Device.Password,
		Method:    mt.Get,
		Command:   cmd.DateTimeSecond.Code,
		Type:      cmd.DateTimeSecond.Type,
		Direction: dir.Write,
		Num:       "1",
		IdString:  "second",
		Payload:   r.Device.DateTime.Second,
	}
	resp = req.Send()
	return
}

func (r *Request) SetDateTimeTimezoneOffset() (resp string) {
	req := rcp.Request{
		IPAddress: r.Device.IPAddress,
		Password:  r.Device.Password,
		Method:    mt.Get,
		Command:   cmd.DateTimeTimezoneOffset.Code,
		Type:      cmd.DateTimeTimezoneOffset.Type,
		Direction: dir.Write,
		Num:       "1",
		IdString:  "selTimeZone",
		Payload:   strconv.Itoa(r.Device.DateTime.TimezoneOffset),
	}
	resp = req.Send()
	return
}

func (r *Request) SetVCAProfile() (resp string) {
	req := rcp.Request{
		IPAddress: r.Device.IPAddress,
		Password:  r.Device.Password,
		Method:    mt.Get,
		Command:   cmd.VCAProfile.Code,
		Type:      cmd.VCAProfile.Type,
		Direction: dir.Write,
		Num:       "1",
		Payload:   r.Device.VCAProfile,
	}

	resp = req.Send()
	return
}

func (r *Request) SetRelayOutputOn() (resp string) {
	req := rcp.Request{
		IPAddress: r.Device.IPAddress,
		Password:  r.Device.Password,
		Method:    mt.Get,
		Command:   cmd.RelayOutputOn.Code,
		Type:      cmd.RelayOutputOn.Type,
		Direction: dir.Write,
		Num:       r.Device.RelayOutputOn,
		Payload:   "1",
	}

	resp = req.Send()
	return
}

func (r *Request) SetRelayOutputOff() (resp string) {
	req := rcp.Request{
		IPAddress: r.Device.IPAddress,
		Password:  r.Device.Password,
		Method:    mt.Get,
		Command:   cmd.RelayOutputOn.Code,
		Type:      cmd.RelayOutputOn.Type,
		Direction: dir.Write,
		Num:       r.Device.RelayOutputOff,
		Payload:   "0",
	}

	resp = req.Send()
	return
}

func (r *Request) SetRelayInputOn() (resp string) {
	req := rcp.Request{
		IPAddress: r.Device.IPAddress,
		Password:  r.Device.Password,
		Method:    mt.Get,
		Command:   cmd.RelayInputOn.Code,
		Type:      cmd.RelayInputOn.Type,
		Direction: dir.Write,
		Num:       r.Device.RelayInputOn,
		Payload:   "1",
	}

	resp = req.Send()
	return
}

func (r *Request) SetRelayInputOff() (resp string) {
	req := rcp.Request{
		IPAddress: r.Device.IPAddress,
		Password:  r.Device.Password,
		Method:    mt.Get,
		Command:   cmd.RelayInputOff.Code,
		Type:      cmd.RelayInputOff.Type,
		Direction: dir.Write,
		Num:       r.Device.RelayInputOff,
		Payload:   "2",
	}

	resp = req.Send()
	return
}

func (r *Request) SetAudioOn() (resp string) {
	req := rcp.Request{
		IPAddress: r.Device.IPAddress,
		Password:  r.Device.Password,
		Method:    mt.Get,
		Command:   cmd.AudioOn.Code,
		Type:      cmd.AudioOn.Type,
		Direction: dir.Write,
		Payload:   "1",
	}

	resp = req.Send()
	return
}

func (r *Request) SetAudioOff() (resp string) {
	req := rcp.Request{
		IPAddress: r.Device.IPAddress,
		Password:  r.Device.Password,
		Method:    mt.Get,
		Command:   cmd.AudioOff.Code,
		Type:      cmd.AudioOff.Type,
		Direction: dir.Write,
		Payload:   "0",
	}

	resp = req.Send()
	return
}

func (r *Request) SetCSRFProtectionOn() (resp string) {
	req := rcp.Request{
		IPAddress: r.Device.IPAddress,
		Password:  r.Device.Password,
		Method:    mt.Get,
		Command:   cmd.CSRFProtectionOn.Code,
		Type:      cmd.CSRFProtectionOn.Type,
		Direction: dir.Write,
		Payload:   "1",
	}

	resp = req.Send()
	return
}

func (r *Request) SetCSRFProtectionOff() (resp string) {
	req := rcp.Request{
		IPAddress: r.Device.IPAddress,
		Password:  r.Device.Password,
		Method:    mt.Get,
		Command:   cmd.CSRFProtectionOff.Code,
		Type:      cmd.CSRFProtectionOff.Type,
		Direction: dir.Write,
		Payload:   "0",
	}

	resp = req.Send()
	return
}
