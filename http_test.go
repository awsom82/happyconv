package main

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestNewServer(t *testing.T) {

	srv := NewServer()
	srv_dummy := &http.Server{}
	//
	// if !reflect.DeepEqual(srv, srv_dummy) {
	//   t.Errorf("Webconv() wrong return %d != %d", srv, srv_dummy)
	// }

	if reflect.TypeOf(srv).String() != reflect.TypeOf(srv_dummy).String() {
		t.Errorf("TestNewServer() wrong return %s must be %s type", reflect.TypeOf(srv).String(), reflect.TypeOf(srv_dummy).String())
	}
}

// TestWebServer send json and checks server reply with xml
func TestServeHTTP(t *testing.T) {

	var payload bytes.Buffer
	ct := "application/json"

	payload.WriteString(`{"books": "Hello, i'am a test string"}`)
	check_xml := []byte(`<?xml version="1.0" encoding="UTF-8"?>` + "\n" + `<books>Hello, i'am a test string</books>`)

	ts := httptest.NewServer(http.HandlerFunc(WebconvHadler))
	defer ts.Close()

	res, err := http.Post(ts.URL, ct, &payload)
	if err != nil {
		t.Fatal(err)
	}

	contenttype := res.Header.Get("Content-type")
	result, err := ioutil.ReadAll(res.Body)

	if err != nil {
		t.Fatal(err)
	} else {
		defer res.Body.Close()
	}

	if contenttype != "application/xml" {
		t.Error("Content-type key incorrect")
	}

	if string(result) != string(check_xml) {
		t.Error("Recieved payload data wrong")
	}

}

func BenchmarkParallelTestServeHTTP(b *testing.B) {

	h := http.HandlerFunc(WebconvHadler)
	ts := httptest.NewServer(h)
	defer ts.Close()

	b.RunParallel(func(pb *testing.PB) {

		for pb.Next() {
			ct := "application/json"
			payload := bytes.NewReader([]byte(`{"books": "Hello, i'am a test string"}`))

			res, err := http.Post(ts.URL, ct, payload)
			if err != nil {
				b.Fatal(err)
			} else {
				res.Body.Close()
			}

		}
	})

}
