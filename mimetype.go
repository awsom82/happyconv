package main

import (
	"net/http"
	"strings"
)

type MimeType uint8

const (
	Unsupported MimeType = iota
	JSON
	XML
)

// List of mime-type in server response
var ctypes = [...]string{XML: "application/xml", JSON: "application/json", Unsupported: "application/octet-stream"}

// FromString sets MimeType from string
// It receives a content-type, removes charset part if they presence in ct
// Ans sets proper value
func (m *MimeType) FromString(ct string) {

	// Remove encoding part and clear extra spaces
	cts := func() string {
		x := strings.Split(ct, ";")
		return strings.Replace(x[0], " ", "", -1)
	}()

	switch cts {

	case "application/json":
		*m = JSON

	case "text/xml", "application/xml":
		*m = XML

	default:
		*m = Unsupported
	}

}

// DetectContentType detects MIME-Type by reading it from header, if this sections is missing.
// Then we try to detect it by MIME Sniff algoritm (https://mimesniff.spec.whatwg.org/)
// this is embedded algoritm in net/http std package
func (m *MimeType) DetectContentType(ct string, body []byte) {

	if ct == "" || ct == "application/octet-stream" {
		ct = http.DetectContentType(body)
	}

	m.FromString(ct)

}

// String return mimetype string
func (m *MimeType) String() string {
	return ctypes[*m]
}
