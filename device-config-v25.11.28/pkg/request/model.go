package request

import (
	d "gitlab.cbsdev.net/quality-assurance/device-config/pkg/device"
	o "gitlab.cbsdev.net/quality-assurance/device-config/pkg/output"
)

type Request struct {
	Method    string
	Command   string
	Type      string
	Direction string
	Payload   string
	Num       string
	Device    d.Device
	Output    o.Output
}
