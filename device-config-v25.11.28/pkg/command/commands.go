package command

var Reboot = Command{
	Code:        "0x0811",
	Type:        fFlag,
	Name:        "Reboot",
	WriteAction: "Rebooting",
}

var ResetFactory = Command{
	Code:        "0x09a0",
	Type:        fFlag,
	Name:        "Reset Factory",
	WriteAction: "Resetting factory",
}

var Reset = Command{
	Code:        "0x093d",
	Type:        fFlag,
	Name:        "Reset",
	WriteAction: "Resetting",
}

var NewServicePassword = Command{
	Code:        "0x028b",
	Type:        pString,
	Name:        "New Service Password",
	WriteAction: "Setting a new service password",
}

var ServicePassword = Command{
	Code:        "0x028b",
	Type:        pString,
	Name:        "Service Password",
	WriteAction: "Changing the service password",
}

var LivePassword = Command{
	Code:        "0x028b",
	Type:        pString,
	Name:        "Live Password",
	WriteAction: "Changing the live password",
}

var UserPassword = Command{
	Code:        "0x028b",
	Type:        pString,
	Name:        "User Password",
	WriteAction: "Changing the user password",
}

var UploadCertificate = Command{
	Name:        "Certificate",
	WriteAction: "Uploading certificate",
}

var CertificateUsage = Command{
	Code:        "0x0bf2",
	Type:        pOctet,
	Name:        "Certificate Usage",
	WriteAction: "Changing certificate usage",
}

var CloudCommission = Command{
	Code:        "0x0c72",
	Type:        pOctet,
	Name:        "Cloud Commission",
	WriteAction: "Commissioning in the cloud",
}

var Capabilities = Command{
	Code:        "0x0b60",
	Type:        pOctet,
	Name:        "Capabilities",
	ReadAction:  "Getting capabilities",
	WriteAction: "Setting capabilities",
}

var MacAddress = Command{
	Code:        "0x00BC",
	Type:        pOctet,
	Name:        "MAC Address",
	ReadAction:  "Getting mac address",
	WriteAction: "Setting mac address",
}

var Name = Command{
	Code:        "0x0024",
	Type:        pUnicode,
	Name:        "Name",
	ReadAction:  "Getting name",
	WriteAction: "Setting name",
}

var CloudDestination = Command{
	Code:        "0x0c75",
	Type:        pOctet,
	Name:        "Cloud Destination",
	ReadAction:  "Getting cloud destination",
	WriteAction: "Setting cloud destination",
}

var SocketKnockerDestination = Command{
	Code:        "0x0aee",
	Type:        pOctet,
	Name:        "Socket Knocker Destination",
	ReadAction:  "Getting socket knocker destination",
	WriteAction: "Setting socket knocker destination",
}

var SocketKnockerStatusReason = Command{
	Code:       "0x0b98",
	Type:       pOctet,
	Name:       "Socket Knocker Status/Reason",
	ReadAction: "Getting socket knocker status/reason",
}

var SocketKnockerMode = Command{
	Code:        "0x0b5c",
	Type:        tDword,
	Name:        "Socket Knocker Mode",
	ReadAction:  "Getting socket knocker mode",
	WriteAction: "Setting socket knocker mode",
}

var ProductName = Command{
	Code:       "0x0aea",
	Type:       pString,
	Name:       "Product Name",
	ReadAction: "Getting product name",
}

var CloudStatus = Command{
	Code:       "0x0c73",
	Type:       pOctet,
	Name:       "Cloud Status",
	ReadAction: "Getting cloud status",
}

var VCAProfile = Command{
	Code:        "0x0a65",
	Type:        tOctet,
	Name:        "VCA Profile",
	ReadAction:  "Getting vca profile",
	WriteAction: "Setting vca profile",
}

var FirmwareVersionFormatted = Command{
	Code:       "0x0cd4",
	Type:       pString,
	Name:       "Firmware Version",
	ReadAction: "Getting firmware version",
}

var FirmwareVersion = Command{
	Code:       "0x002f",
	Type:       pString,
	Name:       "Firmware Version",
	ReadAction: "Getting firmware version",
}

var DateTime = Command{
	Code:        "0x0ba8",
	Type:        pString,
	Name:        "Date/Time",
	ReadAction:  "Getting date/time",
	WriteAction: "Setting date/time",
}

var DateTimeOverwriteByDHCP = Command{
	Code:        "0x0c0e",
	Type:        tOctet,
	Name:        "Date/Time Overwrite by DHCP",
	ReadAction:  "Getting overwrite by dhcp",
	WriteAction: "Setting overwrite by dhcp",
}

var DateTimeFormat = Command{
	Code:        "0x01e9",
	Type:        tOctet,
	Name:        "Date/Time Format",
	ReadAction:  "Getting date/time format",
	WriteAction: "Setting date/time format",
}

var DateTimeYear = Command{
	Code:        "0x002a",
	Type:        tWord,
	Name:        "Date/Time Year",
	ReadAction:  "Getting date/time year",
	WriteAction: "Setting date/time year",
}

var DateTimeMonth = Command{
	Code:        "0x0029",
	Type:        tOctet,
	Name:        "Date/Time Month",
	ReadAction:  "Getting date/time month",
	WriteAction: "Setting date/time month",
}

var DateTimeDay = Command{
	Code:        "0x0028",
	Type:        tOctet,
	Name:        "Date/Time Day",
	ReadAction:  "Getting date/time day",
	WriteAction: "Setting date/time day",
}

var DateTimeHour = Command{
	Code:        "0x002d",
	Type:        tOctet,
	Name:        "Date/Time Hour",
	ReadAction:  "Getting date/time hour",
	WriteAction: "Setting date/time hour",
}

var DateTimeMinute = Command{
	Code:        "0x002c",
	Type:        tOctet,
	Name:        "Date/Time Minute",
	ReadAction:  "Getting date/time minute",
	WriteAction: "Setting date/time minute",
}

var DateTimeSecond = Command{
	Code:        "0x002b",
	Type:        tOctet,
	Name:        "Date/Time Second",
	ReadAction:  "Getting date/time second",
	WriteAction: "Setting date/time second",
}

var DateTimeTimezoneOffset = Command{
	Code:        "0x031f",
	Type:        tInt,
	Name:        "Date/Time Timezone Offset",
	ReadAction:  "Getting date/time timezone offset",
	WriteAction: "Setting date/time timezone offset",
}

var RelayOutputOn = Command{
	Code:        "0x01C1",
	Type:        fFlag,
	Name:        "Relay Output On",
	WriteAction: "Setting relay output on",
}

var RelayOutputOff = Command{
	Code:        "0x01C1",
	Type:        fFlag,
	Name:        "Relay Output Off",
	WriteAction: "Setting relay output off",
}

var RelayInputOn = Command{
	Code:        "0x008d",
	Type:        tOctet,
	Name:        "Relay Input On",
	WriteAction: "Setting relay input on",
}

var RelayInputOff = Command{
	Code:        "0x008d",
	Type:        tOctet,
	Name:        "Relay Input Off",
	WriteAction: "Setting relay input off",
}

var AudioOn = Command{
	Code:        "0x000c",
	Type:        fFlag,
	Name:        "Audio On",
	WriteAction: "Setting audio on",
}

var AudioOff = Command{
	Code:        "0x000c",
	Type:        fFlag,
	Name:        "Audio Off",
	WriteAction: "Setting audio off",
}

var CSRFProtectionOn = Command{
	Code:        "0x0d0d",
	Type:        fFlag,
	Name:        "CSRF Protection On",
	WriteAction: "Setting CSRF protection on",
}

var CSRFProtectionOff = Command{
	Code:        "0x0d0d",
	Type:        fFlag,
	Name:        "CSRF Protection Off",
	WriteAction: "Setting CSRF protection off",
}

var CSRFProtectionStatus = Command{
	Code:       "0x0d0d",
	Type:       fFlag,
	Name:       "CSRF Protection Status",
	ReadAction: "Getting CSRF protection status",
}

var DeviceType = Command{
	Name: "Device Type",
}

var PTZStatus = Command{
	Name: "PTZ",
}
