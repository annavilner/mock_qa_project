package request

import (
	"encoding/hex"
	"strconv"
	"strings"
)

func textToHex(text string) (hexOutput string) {
	hexOutput = hex.EncodeToString([]byte(text))
	hexOutput = strings.ToUpper(hexOutput)
	return
}

func intToHex(num int64) (hexOutput string) {
	hexOutput = strings.ToUpper(strconv.FormatInt(num, 16))
	for i := 0; i < (4 - len(hexOutput)); i++ {
		hexOutput = "0" + hexOutput
	}
	return
}

func nameToHex(name string) (hexOutput string) {
	for i := 0; i < len(name); i++ {
		letter := hex.EncodeToString([]byte(string(name[i])))
		letterSize := len(letter)
		for j := 0; j < (4 - letterSize); j++ {
			letter = "0" + letter;
		}
		hexOutput += letter;
	}
	hexOutput = strings.ToUpper(hexOutput)
	return
}
