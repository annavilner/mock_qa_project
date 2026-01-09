package device

type Device struct {
	IPAddress         string
	Password          string
	Name              string
	ServicePassword   string
	LivePassword      string
	UserPassword      string
	SocketKnockerUrl  string
	SocketKnockerPort int64
	SocketKnockerMode string
	CloudUrl          string
	CloudPort         int64
	CloudEmail        string
	CloudPassword     string
	Rebooted          bool
	Capabilities      string
	DateTimeUTC       string
	VCAProfile        string
	RelayOutputOn     string
	RelayOutputOff    string
	RelayInputOn      string
	RelayInputOff     string
	DateTime          DeviceDateTime
}

type DeviceDateTime struct {
	Year           string
	Month          string
	Day            string
	Hour           string
	Minute         string
	Second         string
	TimezoneOffset int
}
