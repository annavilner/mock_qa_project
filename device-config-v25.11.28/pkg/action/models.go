package action

import (
	o "gitlab.cbsdev.net/quality-assurance/device-config/pkg/output"
	r "gitlab.cbsdev.net/quality-assurance/device-config/pkg/request"
)

type Action struct {
	Options  map[string]Option
	Request r.Request
	Output  o.Output
}

type Option struct {
	Value any
}

/*
type Option struct {
	IPAddress                         *string
	Password                          *string
	DeviceDetailsOption               *bool
	FullDeviceDetailsOption           *bool
	RebootOption                      *bool
	ResetFactoryOption                *bool
	ResetOption                       *bool
	NewServicePasswordOption          *bool
	ChangeServicePasswordOption       *string
	ChangeLivePasswordOption          *string
	ChangeUserPasswordOption          *string
	AddCertificateOption              *bool
	SetSocketKnockerDestinationOption *string
	SetCloudDestinationOption         *string
	CloudCommissionOption             *string
	PrepareToRPRemoqaOption           *bool
	PrepareToRPRemodevnewOption       *bool
	PrepareToAMTestOption             *bool
	PrepareToAMBtuhtestOption         *bool
	SetNameOption                     *string
	SetSocketKnockerModeOption        *string
	SyncDateTimeToPC                  *bool
	SetVCAProfileOption               *string
	SetRelayOutputOnOption            *string
	SetRelayOutputOffOption           *string
	SetRelayInputOnOption             *string
	SetRelayInputOffOption            *string
	SetAudioOnOption                  *bool
	SetAudioOffOption                 *bool
	SimpleOutputOption                *bool
	NoOutputOption                    *bool
	HelpOption                        *bool
}
*/
