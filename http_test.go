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
	srv_dummy := &Server{}

	if !reflect.DeepEqual(srv, srv_dummy) {
		t.Errorf("Webconv() wrong return %d != %d", srv, srv_dummy)
	}
}

// TestWebServer send json and checks server reply with xml
func TestServeHTTP(t *testing.T) {

	var payload bytes.Buffer
	ct := "application/json"

	payload.WriteString(`{"books": "Hello, i'am a test string"}`)
	check_xml := []byte(`<?xml version="1.0" encoding="UTF-8"?>` + "\n" + `<books>Hello, i'am a test string</books>`)

	serv := NewServer()
	ts := httptest.NewServer(serv)
	defer ts.Close()

	res, err := http.Post(ts.URL, ct, &payload)
	if err != nil {
		t.Fatal(err)
	}

	contenttype := res.Header.Get("Content-type")
	result, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		t.Fatal(err)
	}

	if contenttype != "application/xml" {
		t.Error("Content-type key incorrect")
	}

	if string(result) != string(check_xml) {
		t.Error("Recieved payload data wrong")
	}

}

func BenchmarkParallelTestServeHTTP(b *testing.B) {

	serv := NewServer()
	ts := httptest.NewServer(serv)
	defer ts.Close()

	b.RunParallel(func(pb *testing.PB) {

		for pb.Next() {
			ct := "application/json"
			payload := bytes.NewReader([]byte(`{"books": "Hello, i'am a test string"}`))

			res, err := http.Post(ts.URL, ct, payload)
			if err != nil {
				b.Fatal(err)
			}
			res.Body.Close()
		}
	})

}
