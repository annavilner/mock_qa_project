package output

import "strings"

func (o *Output) CPP() (cpp string) {
	if o.Capabilities == "" {
		return
	}
	cppId := strings.Split(o.Capabilities, "000d0001")[1]
	cppId = cppId[0:2]
	switch o.hexToDecimal(cppId) {
	case 1:
		cpp = "CPP3"
	case 2:
		cpp = "CPP4"
	case 3:
		cpp = "CPP-ENC"
	case 4:
		cpp = "CPP5"
	case 5:
		cpp = "CPP6"
	case 6:
		cpp = "CPP7"
	case 7:
		cpp = "CPP7.3"
	case 9:
		cpp = "CPP13"
	case 10:
		cpp = "CPP14"
	case 14:
		cpp = "CPP14.3"
	case 255:
		cpp = "OTHER"
	default:
		cpp = "OTHER"
	}
	cpp = "(" + cpp + ")"

	return
}

func (o *Output) PTZStatus() (ptzStatus string) {
	if o.Capabilities == "" {
		return
	}
	typeId := o.getIfExist(strings.Split(o.Capabilities, "00020001"),1)
	if len(typeId) >= 2 {
		typeId = typeId[0:2]
	}
	switch o.hexToDecimal(typeId) {
	case 0:
		ptzStatus = "No PTZ"
	case 1:
		ptzStatus = "Full PTZ"
	case 2:
		ptzStatus = "Zoom Only"
	}

	return
}

func (o *Output) DeviceType() (deviceType string) {
	if o.Capabilities == "" {
		return
	}
	typeId := strings.Split(o.Capabilities, "00010001")[1]
	typeId = typeId[0:2]
	switch o.hexToDecimal(typeId) {
	case 1:
		deviceType = "Encoder"
	case 2:
		deviceType = "Camera"
	case 3:
		deviceType = "Transcoder"
	case 4:
		deviceType = "Vrm"
	case 5:
		deviceType = "Decoder"
	case 6:
		deviceType = "Streaming gateway "
	case 8:
		deviceType = "Storage"
	case 0:
		deviceType = "Other"
	}

	return
}
