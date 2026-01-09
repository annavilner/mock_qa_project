package rcp

import "encoding/base64"

func basicAuth(password string) string {
	auth := "service" + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}
