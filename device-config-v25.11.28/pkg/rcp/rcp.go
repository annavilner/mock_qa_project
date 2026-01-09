package rcp

import (
	"bytes"
	"crypto/tls"
	"embed"
	"io"
	"mime/multipart"
	"net/http"
	"time"
)

//go:embed cert/Keenfinity-DEV-Root-CA.crt
var crt embed.FS

var transport = &http.Transport{
	TLSClientConfig: &tls.Config{
		InsecureSkipVerify: true,
		MinVersion:         tls.VersionTLS10,
		MaxVersion:         tls.VersionTLS12,

		ServerName: "", // disable SNI
		CipherSuites: []uint16{
			tls.TLS_RSA_WITH_AES_128_CBC_SHA,
			tls.TLS_RSA_WITH_AES_256_CBC_SHA,
			tls.TLS_RSA_WITH_3DES_EDE_CBC_SHA,
		},
	},
	ForceAttemptHTTP2:  false,
	DisableKeepAlives:  false,
	MaxIdleConns:       10,
	IdleConnTimeout:    10 * time.Second,
	DisableCompression: true,
}
var client = &http.Client{
	Transport: transport,
	Timeout:   10 * time.Second,
}

func (r *Request) Send() string {
	req, err := http.NewRequest(r.Method, "https://"+r.IPAddress+"/rcp.xml", nil)
	if err != nil {
		return err.Error()
	}

	if r.Password != "" {
		req.Header.Add("Authorization", "Basic "+basicAuth(r.Password))
	}

	q := "command=" + r.Command
	q += "&type=" + r.Type
	q += "&direction=" + r.Direction
	if r.Num != "" {
		q += "&num=" + r.Num
	}
	if r.IdString != "" {
		q += "&idstring=" + r.IdString
	}
	if r.Payload != "" {
		q += "&payload=" + r.Payload
	}
	req.URL.RawQuery = q

	res, err := client.Do(req)
	if err != nil {
		return err.Error()
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err.Error()
	}

	return string(body)
}

func (r *Request) Upload() string {
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}

	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	file, err := crt.Open("cert/Keenfinity-DEV-Root-CA.crt")
	if err != nil {
		return err.Error()
	}

	defer file.Close()

	part1, err := writer.CreateFormFile("filename", "Keenfinity-DEV-Root-CA.crt")
	if err != nil {
		return err.Error()
	}

	_, err = io.Copy(part1, file)
	if err != nil {
		return err.Error()

	}
	err = writer.Close()
	if err != nil {
		return err.Error()
	}

	req, err := http.NewRequest(r.Method, "https://"+r.IPAddress+"/upload.htm", payload)
	if err != nil {
		return err.Error()
	}
	req.Close = true

	req.Header.Add("Authorization", "Basic "+basicAuth(r.Password))
	req.Header.Set("Content-Type", writer.FormDataContentType())

	res, err := client.Do(req)
	if err != nil {
		return err.Error()
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err.Error()
	}

	return string(body)
}
