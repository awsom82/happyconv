package main

import (
	"bytes"
	"encoding/json"
	"github.com/clbanning/mxj"
	"io"
	"log"
	"net/http"
)

type Converter struct {
	Payload           bytes.Buffer
	ContentType       MimeType
	Result            bytes.Buffer
	ResultContentType MimeType
}

func NewConv() *Converter {
	return &Converter{}
}

// CopyInput creates copy of received data and sets content-type
func (c *Converter) CopyInput(r *http.Request) {

	//this is not necessary, but its better keep payload
	// in struct, for feature purposes (saving to db, etc.)
	io.Copy(&c.Payload, r.Body)

	ct := r.Header.Get("Content-type")
	c.ContentType.DetectContentType(ct, c.Payload.Bytes())

}

// MakeReply returns forms output reply from webserver
func (c *Converter) MakeReply(w http.ResponseWriter) {
	w.Header().Set("Content-Type", c.ResultContentType.String())

	if c.ContentType == Unsupported {
		http.Error(w, "Unsupported format: conventer only accepts XML or JSON", http.StatusUnsupportedMediaType)

	} else {
		io.Copy(w, &c.Result)

	}
}

// SwapFormat fill Input.Result
// if input data is Unsupported they return raw data back
// (TODO: maybe is better to return 415 Unsupported Media Type?)
func (c *Converter) SwapFormat() {

	switch c.ContentType {

	case JSON:
		c.json2xml()

	case XML:
		c.xml2json()

	default:

	}

}

// xml2json converts XML in Payload and saves JSON to Result
func (c *Converter) xml2json() {

	c.ResultContentType = JSON

	mv, err := mxj.NewMapXml(c.Payload.Bytes())
	if err != nil {
		log.Println(err)
	}
	var js []byte
	js, err = json.Marshal(mv)
	if err != nil {
		log.Println(err)
	}

	c.Result.Write(js)

}

// json2xml converts JSON in Payload and saves XML to Result
func (c *Converter) json2xml() {

	c.ResultContentType = XML

	var js map[string]interface{}
	if err := json.Unmarshal(c.Payload.Bytes(), &js); err != nil {
		log.Println(err)
	}

	mv := mxj.Map(js)
	xml, err := mv.XmlIndent("", "  ")
	if err != nil {
		log.Println(err)
	}

	c.Result.WriteString(`<?xml version="1.0" encoding="UTF-8"?>` + "\n")
	c.Result.Write(xml)

}
