package output

import (
	"encoding/xml"

	"github.com/briandowns/spinner"
)

type RCPResponse struct {
	XMLName xml.Name `xml:"rcp"`
	Text    string   `xml:",chardata"`
	Command struct {
		Text string `xml:",chardata"`
		Hex  string `xml:"hex"`
		Dec  string `xml:"dec"`
	} `xml:"command"`
	Type      string `xml:"type"`
	Direction string `xml:"direction"`
	Num       string `xml:"num"`
	Idstring  string `xml:"idstring"`
	Payload   string `xml:"payload"`
	Sessionid string `xml:"sessionid"`
	Auth      string `xml:"auth"`
	Protocol  string `xml:"protocol"`
	Result    struct {
		Text string `xml:",chardata"`
		Len  string `xml:"len"`
		Str  string `xml:"str"`
		Dec  string `xml:"dec"`
		Err  string `xml:"err"`
	} `xml:"result"`
}

type Output struct {
	Capabilities string
	DateTime     DateTime
	SimpleOutput bool
	NoOutput     bool
	Spin      *spinner.Spinner
}

type DateTime struct {
	Year   string
	Month  string
	Day    string
	Hour   string
	Minute string
	Second string
	UTC    string
}

type DeviceDetailsLine struct {
	Name  string
	Value string
}
