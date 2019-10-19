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
	srvDummy := &http.Server{}

	if reflect.TypeOf(srv) != reflect.TypeOf(srvDummy) {
		t.Errorf("TestNewServer() wrong return %T must be %T type", reflect.TypeOf(srv), reflect.TypeOf(srvDummy))
	}
}

// TestWebServer send json and checks server reply with xml
func TestServeHTTP(t *testing.T) {

	var payload bytes.Buffer
	ct := "application/json"

	payload.WriteString(`{"books": "Hello, i'am a test string"}`)
	checkXML := []byte(`<?xml version="1.0" encoding="UTF-8"?>` + "\n" + `<books>Hello, i'am a test string</books>`)

	ts := httptest.NewServer(http.HandlerFunc(WebconvHandler))
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

	if string(result) != string(checkXML) {
		t.Error("Received payload data wrong")
	}

	res, err = http.Get(ts.URL)
	if err != nil {
		t.Fatal(err)
	}

	contenttype = res.Header.Get("Content-type")
	result, err = ioutil.ReadAll(res.Body)

	if err != nil {
		t.Fatal(err)
	} else {
		defer res.Body.Close()
	}

	if res.StatusCode != 405 {
		t.Error("Wrong response code for GET request")

	}

}

func BenchmarkParallelTestServeHTTP(b *testing.B) {

	h := http.HandlerFunc(WebconvHandler)
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
