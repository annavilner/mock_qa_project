package request

import (
	mt "gitlab.cbsdev.net/quality-assurance/device-config/pkg/method"
	rcp "gitlab.cbsdev.net/quality-assurance/device-config/pkg/rcp"
)

func (r *Request) UploadCertificate() (resp string) {
	req := rcp.Request{
		IPAddress: r.Device.IPAddress,
		Password:  r.Device.Password,
		Method:    mt.Post,
	}
	resp = req.Upload()
	return
}
