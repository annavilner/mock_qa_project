package output

import (
	"fmt"
	"strconv"
	"strings"
)

func (o *Output) Capabilites(response string) bool {
	if !o.isNotAuthorized(response) && !o.hasError(response) {
		resp, err := o.parseResponseXML(response)
		if err != nil {
			return false
		}

		str := strings.ReplaceAll(resp.Result.Str, "\n", "")
		str = strings.ReplaceAll(str, " ", "")
		o.Capabilities = str

		return true
	}
	return false
}

func (o *Output) MacAddress(response string) (macAddress string) {
	if !o.isNotAuthorized(response) && !o.hasError(response) {
		resp, err := o.parseResponseXML(response)
		if err == nil {
			macAddress = strings.TrimSuffix(resp.Result.Str, " ")
			macAddress = strings.ReplaceAll(macAddress, " ", "-")
			macAddress = strings.ToUpper(macAddress)
		}
	}
	return
}

func (o *Output) Name(response string) (name string) {
	if !o.isNotAuthorized(response) && !o.hasError(response) {
		resp, err := o.parseResponseXML(response)
		if err == nil {
			name = o.hexToText(resp.Result.Str)
		}
	}
	return
}

func (o *Output) CloudDestination(response string) (cloudDestination string) {
	if !o.isNotAuthorized(response) && !o.hasError(response) {
		resp, err := o.parseResponseXML(response)
		if err == nil {
			txt := strings.ReplaceAll(resp.Result.Str, " ", "")
			txt = strings.ReplaceAll(txt, "\n", "")
			url := txt[16:]
			size := len(url)
			url = url[:size-2]
			port := txt[:4]
			ssl := o.hexToDecimal(txt[4:6])
			sslText := "No"
			if ssl == 1 {
				sslText = "Yes"
			}
			cloudDestination = o.hexToText(url) + ":" + strconv.FormatInt(o.hexToDecimal(port), 10) + "  SSL:" + sslText
		}
	}
	return
}

func (o *Output) SocketKnockerDestination(response string) (socketKnockerDestination string) {
	if !o.isNotAuthorized(response) && !o.hasError(response) {
		resp, err := o.parseResponseXML(response)
		if err == nil {
			txt := strings.ReplaceAll(resp.Result.Str, " ", "")
			txt = strings.ReplaceAll(txt, "\n", "")
			url := txt[16:]
			size := len(url)
			url = url[:size-2]
			port := txt[:4]
			ssl := o.hexToDecimal(txt[4:6])
			sslText := "No"
			if ssl == 1 {
				sslText = "Yes"
			}
			socketKnockerDestination = o.hexToText(url) + ":" + strconv.FormatInt(o.hexToDecimal(port), 10) + "  SSL:" + sslText
		}
	}
	return
}

func (o *Output) FirmwareVersionFormatted(response string) (firmwareVersion string) {
	if !o.isNotAuthorized(response) && !o.hasError(response) {
		resp, err := o.parseResponseXML(response)
		if err == nil {
			firmwareVersion = resp.Result.Str + " " + o.CPP()
		}
	}
	return
}

func (o *Output) FirmwareVersion(response string) (firmwareVersion string) {
	
	resp, err := o.parseResponseXML(response)
	if err == nil {
		firmwareVersion = resp.Result.Str
	}
	
	if len(firmwareVersion) < 8 && !o.isNumeric(firmwareVersion) {
		return ""
	}
	
	build := firmwareVersion[0:2]
	bosch := firmwareVersion[3:4]
	major := firmwareVersion[4:6]
	minor := firmwareVersion[6:8]

	majorInt, _ := strconv.Atoi(major)
	majorClean := strconv.Itoa(majorInt)

	firmwareVersion = fmt.Sprintf("%s.%s.0%s%s", majorClean, minor, bosch, build)
	return
}

func (o *Output) ProductName(response string) (productName string) {
	if !o.isNotAuthorized(response) && !o.hasError(response) {
		resp, err := o.parseResponseXML(response)
		if err == nil {
			productName = resp.Result.Str
		}
	}
	return
}

func (o *Output) SocketKnockerStatusAndReason(response string) (socketKnockerStatusAndReason string) {
	if !o.isNotAuthorized(response) && !o.hasError(response) {
		resp, err := o.parseResponseXML(response)
		if err == nil {
			str := strings.Split(resp.Result.Str, " ")
			state := o.hexToDecimal(str[0])
			stateDesc := ""
			switch state {
			case 0:
				stateDesc = "Not running"
			case 1:
				stateDesc = "Trying to connect"
			case 2:
				stateDesc = "Connected (ready for cloud service)"
			}
			stateReason := o.hexToDecimal(str[1])
			stateReasonDesc := ""
			switch stateReason {
			case 0:
				stateReasonDesc = "As expected"
			case 1:
				stateReasonDesc = "Unknown"
			case 2:
				stateReasonDesc = "Auto mode: DHCP off or not successful"
			case 3:
				stateReasonDesc = "Auto mode: max knocking attempts reached"
			case 4:
				stateReasonDesc = "Auto mode: max knocking time reached"
			case 5:
				stateReasonDesc = "URL could not be resolved		"
			case 6:
				stateReasonDesc = "Destination not responding"
			}

			socketKnockerStatusAndReason = stateDesc + " / " + stateReasonDesc + " "
		}
	}
	return
}

func (o *Output) CloudStatus(response string) (cloudStatus string) {
	if !o.isNotAuthorized(response) && !o.hasError(response) {
		resp, err := o.parseResponseXML(response)
		if err == nil {
			str := strings.Split(resp.Result.Str, " ")
			status := o.hexToDecimal(str[0])
			statusDesc := ""
			switch status {
			case 0:
				statusDesc = "Registered"
			case 1:
				statusDesc = "OnGoing"
			case 2:
				statusDesc = "UnRegistered"
			case 255:
				statusDesc = "Unknown"
			}

			cloudStatus = statusDesc
		}
	}
	return
}

func (o *Output) VCAProfile(response string) (vcaProfile string) {
	if !o.isNotAuthorized(response) && !o.hasError(response) {
		resp, err := o.parseResponseXML(response)
		if err == nil {
			dec := strings.ReplaceAll(resp.Result.Dec, " ", "")
			num, err := strconv.Atoi(dec)
			if err != nil {
				return err.Error()
			}
			switch num {
			case 0:
				vcaProfile = "Silent"
			case 253:
				vcaProfile = "Off"
			case 254:
				vcaProfile = "Scheduler"
			case 255:
				vcaProfile = "Script mode"
			default:
				vcaProfile = "Profile " + strconv.Itoa(num)
			}
		}
	}
	return
}

func (o *Output) SocketKnockerMode(response string) (socketknockerMode string) {
	if !o.isNotAuthorized(response) && !o.hasError(response) {
		resp, err := o.parseResponseXML(response)
		if err == nil {
			dec := strings.ReplaceAll(resp.Result.Dec, " ", "")
			num, err := strconv.Atoi(dec)
			if err != nil {
				fmt.Println(err)
			}
			switch num {
			case 0:
				socketknockerMode = "Off"
			case 1:
				socketknockerMode = "On"
			case 2:
				socketknockerMode = "Auto"
			}
		}
	}
	return
}

func (o *Output) DateTimeYear(response string) (year string) {
	year = "0000"
	if !o.isNotAuthorized(response) && !o.hasError(response) {
		resp, err := o.parseResponseXML(response)
		if err == nil {
			year = strings.ReplaceAll(resp.Result.Dec, " ", "")
		}
	}
	return
}

func (o *Output) DateTimeMonth(response string) (month string) {
	month = "00"
	if !o.isNotAuthorized(response) && !o.hasError(response) {
		resp, err := o.parseResponseXML(response)
		if err == nil {
			month = strings.ReplaceAll(resp.Result.Dec, " ", "")
			if len(month) == 1 {
				month = "0"+month
			}
		}
	}
	return
}

func (o *Output) DateTimeDay(response string) (day string) {
	day = "00"
	if !o.isNotAuthorized(response) && !o.hasError(response) {
		resp, err := o.parseResponseXML(response)
		if err == nil {
			day = strings.ReplaceAll(resp.Result.Dec, " ", "")
			if len(day) == 1 {
				day = "0"+day
			}
		}
	}
	return
}

func (o *Output) DateTimeHour(response string) (hour string) {
	hour = "00"
	if !o.isNotAuthorized(response) && !o.hasError(response) {
		resp, err := o.parseResponseXML(response)
		if err == nil {
			hour = strings.ReplaceAll(resp.Result.Dec, " ", "")
			if len(hour) == 1 {
				hour = "0"+hour
			}
		}
	}
	return
}

func (o *Output) DateTimeMinute(response string) (minute string) {
	minute = "00"
	if !o.isNotAuthorized(response) && !o.hasError(response) {
		resp, err := o.parseResponseXML(response)
		if err == nil {
			minute = strings.ReplaceAll(resp.Result.Dec, " ", "")
			if len(minute) == 1 {
				minute = "0"+minute
			}
		}
	}
	return
}

func (o *Output) DateTimeSecond(response string) (second string) {
	second = "00"
	if !o.isNotAuthorized(response) && !o.hasError(response) {
		resp, err := o.parseResponseXML(response)
		if err == nil {
			second = strings.ReplaceAll(resp.Result.Dec, " ", "")
			if len(second) == 1 {
				second = "0"+second
			}
		}
	}
	return
}

func (o *Output) DateTimeTimezoneOffset(response string) (offset string) {
	offset = "00:00"
	negative := false
	if !o.isNotAuthorized(response) && !o.hasError(response) {
		resp, err := o.parseResponseXML(response)
		if err == nil {
			ret := strings.ReplaceAll(resp.Result.Dec, " ", "")
			if strings.Contains(ret, "-") {
				negative = true
			}
			offsetNum, err := strconv.Atoi(strings.ReplaceAll(ret, "-", ""))
			if err == nil {
				hour := strconv.Itoa(offsetNum/3600)
				minute := strconv.Itoa((offsetNum % 3600) / 60)

				if len(hour) == 1 {
					hour = "0"+hour
				}
				if len(minute) == 1 {
					minute = "0"+minute
				}
				if negative {
					offset = "-"
				} else {
					offset = "+"
				}

				offset += hour+":"+minute 
			}
		}
	}
	return
}

func (o *Output) CSRFProtectionStatus(response string) (csrfProtectionStatus string) {
	if !o.isNotAuthorized(response) && !o.hasError(response) {
		resp, err := o.parseResponseXML(response)
		if err == nil {
			ret := strings.ReplaceAll(resp.Result.Dec, " ", "")
			status := o.hexToDecimal(ret)
			statusDesc := ""
			switch status {
			case 0:
				statusDesc = "Off"
			case 1:
				statusDesc = "On"
			}

			csrfProtectionStatus = statusDesc
		}
	}
	return
}