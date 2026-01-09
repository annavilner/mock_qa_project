package request

import (
	cmd "gitlab.cbsdev.net/quality-assurance/device-config/pkg/command"
	dir "gitlab.cbsdev.net/quality-assurance/device-config/pkg/direction"
	mt "gitlab.cbsdev.net/quality-assurance/device-config/pkg/method"
	rcp "gitlab.cbsdev.net/quality-assurance/device-config/pkg/rcp"
)

func (r *Request) GetCapabilites() string {
	req := rcp.Request{
		IPAddress: r.Device.IPAddress,
		Password:  r.Device.Password,
		Method:    mt.Get,
		Command:   cmd.Capabilities.Code,
		Type:      cmd.Capabilities.Type,
		Direction: dir.Read,
	}
	return req.Send()
}

func (r *Request) GetMacAddress() string {
	req := rcp.Request{
		IPAddress: r.Device.IPAddress,
		Password:  r.Device.Password,
		Method:    mt.Get,
		Command:   cmd.MacAddress.Code,
		Type:      cmd.MacAddress.Type,
		Direction: dir.Read,
	}
	return req.Send()
}

func (r *Request) GetName() string {
	req := rcp.Request{
		IPAddress: r.Device.IPAddress,
		Password:  r.Device.Password,
		Method:    mt.Get,
		Command:   cmd.Name.Code,
		Type:      cmd.Name.Type,
		Direction: dir.Read,
	}
	return req.Send()
}

func (r *Request) GetCloudDestination() string {
	req := rcp.Request{
		IPAddress: r.Device.IPAddress,
		Password:  r.Device.Password,
		Method:    mt.Get,
		Command:   cmd.CloudDestination.Code,
		Type:      cmd.CloudDestination.Type,
		Direction: dir.Read,
	}
	return req.Send()
}

func (r *Request) GetSocketKnockerDestination() string {
	req := rcp.Request{
		IPAddress: r.Device.IPAddress,
		Password:  r.Device.Password,
		Method:    mt.Get,
		Command:   cmd.SocketKnockerDestination.Code,
		Type:      cmd.SocketKnockerDestination.Type,
		Direction: dir.Read,
	}
	return req.Send()
}

func (r *Request) GetFirmwareVersionFormatted() string{
	req := rcp.Request{
		IPAddress: r.Device.IPAddress,
		Password:  r.Device.Password,
		Method:    mt.Get,
		Command:   cmd.FirmwareVersionFormatted.Code,
		Type:      cmd.FirmwareVersionFormatted.Type,
		Direction: dir.Read,
	}
	return req.Send()
}

func (r *Request) GetFirmwareVersion() string{
	req := rcp.Request{
		IPAddress: r.Device.IPAddress,
		Password:  r.Device.Password,
		Method:    mt.Get,
		Command:   cmd.FirmwareVersion.Code,
		Type:      cmd.FirmwareVersion.Type,
		Direction: dir.Read,
	}
	return req.Send()
}

func (r *Request) GetProductName() string {
	req := rcp.Request{
		IPAddress: r.Device.IPAddress,
		Password:  r.Device.Password,
		Method:    mt.Get,
		Command:   cmd.ProductName.Code,
		Type:      cmd.ProductName.Type,
		Direction: dir.Read,
	}
	return req.Send()
}

func (r *Request) GetSocketKnockerStatusAndReason() string {
	req := rcp.Request{
		IPAddress: r.Device.IPAddress,
		Password:  r.Device.Password,
		Method:    mt.Get,
		Command:   cmd.SocketKnockerStatusReason.Code,
		Type:      cmd.SocketKnockerStatusReason.Type,
		Direction: dir.Read,
	}
	return req.Send()
}

func (r *Request) GetCloudStatus() string {
	req := rcp.Request{
		IPAddress: r.Device.IPAddress,
		Password:  r.Device.Password,
		Method:    mt.Get,
		Command:   cmd.CloudStatus.Code,
		Type:      cmd.CloudStatus.Type,
		Direction: dir.Read,
	}
	return req.Send()
}

func (r *Request) GetVCAProfile() string {
	req := rcp.Request{
		IPAddress: r.Device.IPAddress,
		Password:  r.Device.Password,
		Method:    mt.Get,
		Command:   cmd.VCAProfile.Code,
		Type:      cmd.VCAProfile.Type,
		Direction: dir.Read,
	}
	return req.Send()
}

func (r *Request) GetSocketKnockerMode() string {
	req := rcp.Request{
		IPAddress: r.Device.IPAddress,
		Password:  r.Device.Password,
		Method:    mt.Get,
		Command:   cmd.SocketKnockerMode.Code,
		Type:      cmd.SocketKnockerMode.Type,
		Direction: dir.Read,
	}
	return req.Send()
}

func (r *Request) GetDateTimeYear() string {
	req := rcp.Request{
		IPAddress: r.Device.IPAddress,
		Password:  r.Device.Password,
		Method:    mt.Get,
		Command:   cmd.DateTimeYear.Code,
		Type:      cmd.DateTimeYear.Type,
		Direction: dir.Read,
	}
	return req.Send()
}

func (r *Request) GetDateTimeMonth() string {
	req := rcp.Request{
		IPAddress: r.Device.IPAddress,
		Password:  r.Device.Password,
		Method:    mt.Get,
		Command:   cmd.DateTimeMonth.Code,
		Type:      cmd.DateTimeMonth.Type,
		Direction: dir.Read,
	}
	return req.Send()
}

func (r *Request) GetDateTimeDay() string {
	req := rcp.Request{
		IPAddress: r.Device.IPAddress,
		Password:  r.Device.Password,
		Method:    mt.Get,
		Command:   cmd.DateTimeDay.Code,
		Type:      cmd.DateTimeDay.Type,
		Direction: dir.Read,
	}
	return req.Send()
}

func (r *Request) GetDateTimeHour() string {
	req := rcp.Request{
		IPAddress: r.Device.IPAddress,
		Password:  r.Device.Password,
		Method:    mt.Get,
		Command:   cmd.DateTimeHour.Code,
		Type:      cmd.DateTimeHour.Type,
		Direction: dir.Read,
	}
	return req.Send()
}

func (r *Request) GetDateTimeMinute() string {
	req := rcp.Request{
		IPAddress: r.Device.IPAddress,
		Password:  r.Device.Password,
		Method:    mt.Get,
		Command:   cmd.DateTimeMinute.Code,
		Type:      cmd.DateTimeMinute.Type,
		Direction: dir.Read,
	}
	return req.Send()
}

func (r *Request) GetDateTimeSecond() string {
	req := rcp.Request{
		IPAddress: r.Device.IPAddress,
		Password:  r.Device.Password,
		Method:    mt.Get,
		Command:   cmd.DateTimeSecond.Code,
		Type:      cmd.DateTimeSecond.Type,
		Direction: dir.Read,
	}
	return req.Send()
}

func (r *Request) GetDateTimeTimezoneOffset() string {
	req := rcp.Request{
		IPAddress: r.Device.IPAddress,
		Password:  r.Device.Password,
		Method:    mt.Get,
		Command:   cmd.DateTimeTimezoneOffset.Code,
		Type:      cmd.DateTimeTimezoneOffset.Type,
		Direction: dir.Read,
	}
	return req.Send()
}

func (r *Request) GetCSRFProtectionStatus() string {
	req := rcp.Request{
		IPAddress: r.Device.IPAddress,
		Password:  r.Device.Password,
		Method:    mt.Get,
		Command:   cmd.CSRFProtectionStatus.Code,
		Type:      cmd.CSRFProtectionStatus.Type,
		Direction: dir.Read,
	}
	return req.Send()
}
