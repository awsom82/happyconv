package main

import (
	"testing"
)

func TestMimeFromString(t *testing.T) {

	var mime MimeType

	var tests = []struct {
		input string
		want  MimeType
	}{
		{"application/json", JSON},
		{"application/xml", XML},
		{"text/xml", XML},
		{"text/plain", Unsupported},
		{"application/octet-stream", Unsupported},
		{"text/html", Unsupported},
		{"X-somereal/badcontenttype", Unsupported},
	}

	for _, test := range tests {

		mime.FromString(test.input)

		if mime != test.want {
			t.Errorf("Content-type key not correct %d must be %d", mime, test.want)
		}

	}

}

func TestMimeString(t *testing.T) {

	var mime MimeType

	var tests = []struct {
		want  string
		input MimeType
	}{
		{"application/json", JSON},
		{"application/xml", XML},
		{"application/octet-stream", Unsupported},
	}

	for _, test := range tests {

		mime = test.input

		if mime.String() != test.want {
			t.Errorf("Content-type return string key not correct %s must be %s", mime.String(), test.want)
		}

	}

}

func TestMimeDetectContentType(t *testing.T) {

	var mime MimeType

	ct := "application/octet-stream"
	payload := []byte(`<?xml version="1.0" encoding="UTF-8"?><books>Hello, i'am a test string</books>`)

	mime.DetectContentType(ct, payload)

	if mime != XML {
		t.Errorf("Content-type return string key not correct %s must be application/xml", mime.String())
	}
}
