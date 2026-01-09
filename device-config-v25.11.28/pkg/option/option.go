package option

func Major() []Option {
	return []Option{
		{
			Name:        "IPAddress",
			Flag:        "i",
			Type:        "string",
			Description: "Device IP address",
		},
		{
			Name:        "MACAddress",
			Flag:        "m",
			Type:        "string",
			Description: "Device MAC address",
		},
		{
			Name:        "Password",
			Flag:        "p",
			Type:        "string",
			Description: "Device password",
		},
	}
}

func ReadRequests() []Option {
	return []Option{
		{
			Name:        "DeviceDetailsOption",
			Flag:        "d",
			Type:        "bool",
			Description: "Show device details",
		},
		{
			Name:        "FullDeviceDetailsOption",
			Flag:        "fd",
			Type:        "bool",
			Description: "Show device full details",
		},
	}
}

func WriteRequests() []Option {
	return []Option{
		{
			Name:        "RebootOption",
			Flag:        "rb",
			Type:        "bool",
			Description: "Reboot device",
		},
		{
			Name:        "ResetFactoryOption",
			Flag:        "rf",
			Type:        "bool",
			Description: "Reset factory",
		},
		{
			Name:        "ResetOption",
			Flag:        "rs",
			Type:        "bool",
			Description: "Reset",
		},
		{
			Name:        "NewServicePasswordOption",
			Flag:        "nsp",
			Type:        "bool",
			Description: "Setting a new service password. The new password must be set using the option '-p'.",
		},
		{
			Name:        "ChangeServicePasswordOption",
			Flag:        "csp",
			Type:        "string",
			Description: "Change the service password",
		},
		{
			Name:        "ChangeLivePasswordOption",
			Flag:        "clp",
			Type:        "string",
			Description: "Change the live password",
		},
		{
			Name:        "ChangeUserPasswordOption",
			Flag:        "cup",
			Type:        "string",
			Description: "Change the user password",
		},
		{
			Name:        "AddCertificateOption",
			Flag:        "cert",
			Type:        "bool",
			Description: "Add certificate",
		},
		{
			Name:        "SetSocketKnockerDestinationOption",
			Flag:        "skd",
			Type:        "string",
			Description: "Setting socket knocker destination",
		},
		{
			Name:        "SetCloudDestinationOption",
			Flag:        "cd",
			Type:        "string",
			Description: "Setting cloud destination",
		},
		{
			Name:        "CloudCommissionOption",
			Flag:        "cc",
			Type:        "string",
			Description: "Cloud commission (e.g. \"user@email.com:password\")",
		},
		{
			Name:        "SetNameOption",
			Flag:        "n",
			Type:        "string",
			Description: "Setting name",
		},
		{
			Name:        "SetSocketKnockerModeOption",
			Flag:        "skm",
			Type:        "string",
			Description: "Setting socket knocker mode. Available modes: on, off, auto",
		},
		{
			Name:        "SyncDateTimeToPC",
			Flag:        "sdt",
			Type:        "bool",
			Description: "Sync Date/Time to PC",
		},
		{
			Name:        "SetVCAProfileOption",
			Flag:        "vca",
			Type:        "string",
			Description: "Setting VCA profile",
		},
		{
			Name:        "SetRelayOutputOnOption",
			Flag:        "ro-on",
			Type:        "string",
			Description: "Setting relay output to on",
		},
		{
			Name:        "SetRelayOutputOffOption",
			Flag:        "ro-off",
			Type:        "string",
			Description: "Setting relay output to off",
		},
		{
			Name:        "SetRelayInputOnOption",
			Flag:        "ri-on",
			Type:        "string",
			Description: "Setting relay input to on",
		},
		{
			Name:        "SetRelayInputOffOption",
			Flag:        "ri-off",
			Type:        "string",
			Description: "Setting relay input to off",
		},
		{
			Name:        "SetAudioOnOption",
			Flag:        "au-on",
			Type:        "bool",
			Description: "Setting audio to on",
		},
		{
			Name:        "SetAudioOffOption",
			Flag:        "au-off",
			Type:        "bool",
			Description: "Setting audio to off",
		},
		{
			Name:        "SetCSRFProtectionOnOption",
			Flag:        "csrf-on",
			Type:        "bool",
			Description: "Setting CSRF protection to on",
		},
		{
			Name:        "SetCSRFProtectionOffOption",
			Flag:        "csrf-off",
			Type:        "bool",
			Description: "Setting CSRF protection to off",
		},
	}
}

func SpecificRP() []Option {
	return []Option{
		{
			Name:        "PrepareToRPRemoqaOption",
			Flag:        "rp-remoqa",
			Type:        "bool",
			Description: "Prepare device to 'remoqa'",
		},
		{
			Name:        "PrepareToRPRemodevnewOption",
			Flag:        "rp-remodevnew",
			Type:        "bool",
			Description: "Prepare device to 'remodevnew'",
		},
	}
}

func SpecificAM() []Option {
	return []Option{
		{
			Name:        "PrepareToAMTestOption",
			Flag:        "am-test",
			Type:        "bool",
			Description: "Prepare device to 'test'",
		},
		{
			Name:        "PrepareToAMBtuhtestOption",
			Flag:        "am-btuhtest",
			Type:        "bool",
			Description: "Prepare device to 'btuhtest'",
		},
		{
			Name:        "PrepareToAMDevOption",
			Flag:        "am-dev",
			Type:        "bool",
			Description: "Prepare device to 'dev'",
		},
		{
			Name:        "PrepareToAMDemoOption",
			Flag:        "am-demo",
			Type:        "bool",
			Description: "Prepare device to 'demo'",
		},
	}
}

func Others() []Option {
	return []Option{
		{
			Name:        "SimpleOutputOption",
			Flag:        "so",
			Type:        "bool",
			Description: "Simple output. No color, spinner or formatting.",
		},
		{
			Name:        "NoOutputOption",
			Flag:        "no",
			Type:        "bool",
			Description: "No Output",
		},
		{
			Name:        "HelpOption",
			Flag:        "h",
			Type:        "bool",
			Description: "Show help",
		},
		{
			Name:        "VersionOption",
			Flag:        "v",
			Type:        "bool",
			Description: "Show version",
		},
	}
}

func AllOptions() []Option {
	return Merge(Major(), ReadRequests(), WriteRequests(), SpecificRP(), SpecificAM(), Others())
}

func Merge[T any](slices ...[]T) []T {
    var result []T
    for _, s := range slices {
        result = append(result, s...)
    }
    return result
}