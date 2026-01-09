package output

import (
	"embed"
	"encoding/hex"
	"encoding/xml"
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"unicode"

	f "gitlab.cbsdev.net/quality-assurance/device-config/pkg/font"
	op "gitlab.cbsdev.net/quality-assurance/device-config/pkg/option"
)

//go:embed version/version.txt
var v embed.FS

func (o *Output) parseResponseXML(stringXML string) (resp RCPResponse, err error) {
	err = xml.Unmarshal([]byte(stringXML), &resp)
	return
}

func (o *Output) hexToText(hexString string) string {
	bs, err := hex.DecodeString(strings.ReplaceAll(hexString, " ", ""))
	if err != nil {
		fmt.Println(err)
	}
	return string(bs)
}

func (o *Output) hexToDecimal(hexString string) (dec int64) {
	dec, err := strconv.ParseInt(hexString, 16, 64)
	if err != nil {
		fmt.Println(err)
	}
	return
}

func (o *Output) isNotAuthorized(response string) bool {
	return response == "<HTML><HEAD><TITLE>401 Authorization Required</TITLE></HEAD><BODY><H1>Authorization Required</H1>this server could not verify that you are authorized to access the document you requested.  Either you supplied the wrong credentials (e.g., bad password), or your browser doesnt understand how to supply the credentials required.</BODY></HTML>"
}

func (o *Output) hasError(response string) bool {
	return strings.Contains(o.removeInvisibleChar(response), "<result><err>") && strings.Contains(o.removeInvisibleChar(response), "</err></result>")
}

func (o *Output) successMsg(text string) string {
	if !o.NoOutput {
		if !o.SimpleOutput {
			return f.BoldGreen + text + " command successfully executed!" + f.Reset
		} else {
			return text + " command successfully executed!"
		}
	}
	return ""
}

func (o *Output) ErrorMsg(errMsg string) string {
	if !o.NoOutput {
		if !o.SimpleOutput {
			return f.BoldRed + errMsg + f.Reset
		} else {
			return errMsg
		}
	}
	return ""
}

func (o *Output) commandErrorMsgDescription(errMsg string, description string) string {
	if !o.NoOutput {
		if !o.SimpleOutput {
			return f.BoldRed + errMsg + " Error " + f.Reset + f.Italic + f.Red + description + f.Reset
		} else {
			return errMsg + " Error " + " - " + description
		}
	}
	return ""
}

func (o *Output) commandErrorMsg(text string, response RCPResponse) string {
	return o.commandErrorMsgDescription(text, o.commandResponseError(response))
}

func (o *Output) unauthorizedErrorMsg(text string) string {
	return o.commandErrorMsgDescription(text, "401 Authorization Required")
}

func (o *Output) certificateAlreadyExistErrorMsg(text string) string {
	return o.commandErrorMsgDescription(text, "Certificate already exists")
}

func (o *Output) badRequestErrorMsg(text string) string {
	return o.commandErrorMsgDescription(text, "HTTP/1.0 400 Bad Request")
}

func (o *Output) noPasswordErrorMsg(text string) string {
	return o.commandErrorMsgDescription(text, "No service password configure")
}

func (o *Output) CommissionErrorMsg() string {
	if !o.NoOutput {
		if !o.SimpleOutput {
			return f.BoldRed + "Error commissioning device to the cloud" + f.Reset
		} else {
			return "Error commissioning device to the cloud"
		}
	}
	return ""
}

func (o *Output) CommissionSuccessMsg() string {
	if !o.NoOutput {
		if !o.SimpleOutput {
			return f.BoldGreen + "Device successfully commissioned to the cloud!" + f.Reset
		} else {
			return "Device successfully commissioned to the cloud!"
		}
	}
	return ""
}

func (o *Output) ResponseResult(response string, commandDescription string) (code string, msg string) {
	if o.isNotAuthorized(response) {
		return "unauthorized", o.unauthorizedErrorMsg(commandDescription)
	} else if o.isUploadCertificateResponse(response) {
		return "success", o.successMsg(commandDescription)
	} else if o.certificateAlreadyExist(response) {
		return "error", o.certificateAlreadyExistErrorMsg(commandDescription)
	} else if o.isNoPasswordResponse(response) {
		return "nopass", o.noPasswordErrorMsg(commandDescription)
	} else if o.isBadRequestResponse(response) {
		return "error", o.badRequestErrorMsg(commandDescription)
	} else {
		resp, err := o.parseResponseXML(response)
		if err != nil {
			return "error", o.commandErrorMsgDescription(commandDescription, err.Error())
		}
		result := resp.Result
		if result.Err == "" {
			return "success", o.successMsg(commandDescription)
		} else {
			return "error", o.commandErrorMsg(commandDescription, resp)
		}
	}
}

func (o *Output) isUploadCertificateResponse(response string) bool {
	return response == `<table border="1" cellspacing="0" cellpadding="0" bordercolorlight="#FFFFFF" bordercolordark="#1884FF"><tr><td><table border="0" cellspacing="0" cellpadding="0" style="padding: 10px"><tr><td style="padding: 10px"><p style="margin-left: 6">Software update:</p></td><td style="padding: 10px"><form method="POST" action="upload.htm" enctype="multipart/form-data"><table border="0" cellspacing="0" cellpadding="0"><tr><td style="padding-bottom: 10px">Password (only for config):</td><td><input type="password" name="pwd" size="30"</td></tr><tr><td>File:</td><td><input type="file" name="net.bin" size="20" maxlength="8000000"></td></tr></table></td><td style="padding: 10px"><input type="submit" value="Upload" name="Set" border="0"></td></form></tr></table></td></tr></table>`
}

func (o *Output) certificateAlreadyExist(response string) bool {
	return response == "Upload failed. Error code 141"
}

func (o *Output) isNoPasswordResponse(response string) bool {
	return strings.Contains(response, "<title>nopwd_header</title>")
}

func (o *Output) isBadRequestResponse(response string) bool {
	return strings.Contains(response, "<body><h1>HTTP/1.0 400 Bad Request</h1></body>")
}

func (o *Output) commandResponseError(response RCPResponse) string {
	result := response.Result
	switch result.Err {
	case "0xFF":
		return "Unknown"
	case "0x10":
		return "Invalid version"
	case "0x20":
		return "Not Registered"
	case "0x21":
		return "Invalid Client ID"
	case "0x30":
		return "Invalid Method"
	case "0x40":
		return "Invalid CMD"
	case "0x50":
		return "Invalid Access Type"
	case "0x60":
		return "Invalid Data Type"
	case "0x70":
		return "Write Error"
	case "0x80":
		return "Packet Size"
	case "0x90":
		return "Read Not Supported"
	case "0xa0":
		return "Invalid Auth Level"
	case "0xb0":
		return "Invalid Session ID"
	case "0xc0":
		return "Try Later"
	case "0xd0":
		return "Timeout"
	case "0xe0":
		return "No License"
	case "0xf0":
		return "Command Specific"
	case "0xf1":
		return "Address Format"
	default:
		return ""
	}
}

func (o *Output) SetSpinText(text string) {
	if !o.SimpleOutput || !o.NoOutput {
		o.Spin.Suffix = " " + o.spinText(text)
	}
}

func (o *Output) StartSpin() {
	if !o.SimpleOutput && !o.NoOutput {
		o.Spin.Start()
	}
}

func (o *Output) StopSpin() {
	if !o.SimpleOutput && !o.NoOutput {
		o.Spin.Stop()
	}
}

func (o *Output) spinText(text string) string {
	return f.Yellow + f.Italic + text + f.Reset
}

func (o *Output) PrintOutput(text string) {
	if !o.NoOutput {
		fmt.Println(text)
	}
}

func (o *Output) MaxSizeDeviceDetailsLine(ddl []DeviceDetailsLine) (nameMaxSize int, valueMaxSize int) {
	nameMaxSize = 0
	valueMaxSize = 0

	for _, l := range ddl {
		if len(l.Name) > nameMaxSize {
			nameMaxSize = len(l.Name)
		}
		l.Value = strings.TrimSpace(o.removeInvisibleChar(l.Value))
		if len(l.Value) > valueMaxSize {
			valueMaxSize = len(l.Value)
		}
	}
	return
}

func (o *Output) PrintDeviceDetailsLine(ddl []DeviceDetailsLine, nameMaxSize int, valueMaxSize int) {

	sort.Slice(ddl, func(i, j int) bool {
		return ddl[i].Name < ddl[j].Name
	})

	if !o.NoOutput {
		if !o.SimpleOutput {
			for _, l := range ddl {
				l.Name = o.removeInvisibleChar(l.Name)
				l.Value = strings.TrimSpace(o.removeInvisibleChar(l.Value))

				if len(l.Name) < nameMaxSize {
					for i := len(l.Name); i < nameMaxSize; i++ {
						l.Name += " "
					}
				}
				if len(l.Value) < valueMaxSize {
					for j := len(l.Value); j < valueMaxSize; j++ {
						l.Value += " "
					}
				}
				leftPipe := "| "
				middlePipe := " | "
				rightPipe := " |"
				fmt.Println(f.BoldWhite + leftPipe + f.Reset +
					f.BoldCyan + l.Name + f.Reset +
					f.BoldWhite + middlePipe + f.Reset +
					f.Cyan + f.Italic + l.Value + f.Reset +
					f.BoldWhite + rightPipe + f.Reset)
			}
		} else {
			for _, l := range ddl {
				l.Name = o.removeInvisibleChar(l.Name)
				l.Value = strings.TrimSpace(o.removeInvisibleChar(l.Value))

				if l.Value != "" {
					fmt.Println(l.Name + ": " + l.Value)
				}
			}
		}
	}

}

func (o *Output) removeInvisibleChar(text string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsGraphic(r) {
			return r
		}
		return -1
	}, text)
}

func (o *Output) PrintDeviceDetailsDivisionLine(maxLineSize int) {
	if !o.NoOutput {
		maxLineSize = maxLineSize + 7
		line := ""
		for i := 0; i < maxLineSize; i++ {
			line += "-"
		}
		if !o.SimpleOutput {
			fmt.Println(f.BoldWhite + line + f.Reset)
		}
	}
}

func (o *Output) printHelpOption(tag string, tagType string, tagDescription string) {
	if !o.NoOutput {
		if !o.SimpleOutput {
			fmt.Print(f.BoldCyan + "-" + tag + f.Reset + f.Blue + "  " + tagType + f.Reset + "\n   " + f.Italic + f.Yellow + tagDescription + f.Reset + "\n")
		} else {
			fmt.Print("-" + tag + "  " + tagType + " - " + tagDescription + "\n")
		}
	}
}

func (o *Output) printHelpSectionTitle(title string) {
	if !o.NoOutput {
		if !o.SimpleOutput {
			fmt.Println(f.Magenta + title + f.Reset)
		} else {
			fmt.Println(title)
		}
	}
}

func (o *Output) printHelpTitle(title string) {
	if !o.NoOutput {
		if !o.SimpleOutput {
			fmt.Println("\n" + f.BoldCyan + title + f.Reset + "\n")
		} else {
			fmt.Println("\n" + title + "\n")
		}
	}
}

func (o *Output) Usage() {
	o.printHelpTitle("Available Options: ")

	o.printHelpSectionTitle("Major Parameters")
	for _, opt := range op.Major() {
		o.printHelpOption(opt.Flag, opt.Type, opt.Description)
	}

	o.printHelpSectionTitle("Read Requests")
	for _, opt := range op.ReadRequests() {
		o.printHelpOption(opt.Flag, opt.Type, opt.Description)
	}

	o.printHelpSectionTitle("Write Requests")
	for _, opt := range op.WriteRequests() {
		o.printHelpOption(opt.Flag, opt.Type, opt.Description)
	}

	o.printHelpSectionTitle("Preparing the device for a specific Remote Portal environment")
	for _, opt := range op.SpecificRP() {
		o.printHelpOption(opt.Flag, opt.Type, opt.Description)
	}

	o.printHelpSectionTitle("Preparing the device for a specific Alarm Management environment")
	for _, opt := range op.SpecificAM() {
		o.printHelpOption(opt.Flag, opt.Type, opt.Description)
	}

	o.printHelpSectionTitle("Others")
	for _, opt := range op.Others() {
		o.printHelpOption(opt.Flag, opt.Type, opt.Description)
	}

	fmt.Printf("\n")
}

func (o *Output) Version() {
	data, err := v.ReadFile("version/version.txt")
	if err != nil {
		fmt.Println(err.Error())
	}
	o.printVersion(strings.TrimSuffix(string(data), "\n"))
}

func (o *Output) printVersion(version string) {
	if !o.NoOutput {
		if !o.SimpleOutput {
			fmt.Println(f.BoldCyan + version + f.Reset)
		} else {
			fmt.Println(version)
		}
	}
}

func (o *Output) getIfExist(arr []string, index int) (val string) {
	if index >= 0 && index < len(arr) {
		return arr[index]
	}
	return ""
}

func (o *Output) isNumeric(s string) bool {
	return regexp.MustCompile(`^[0-9]+$`).MatchString(s)
}
