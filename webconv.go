package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/clbanning/mxj"
	"io"
	"log"
	"net/http"
)

var (
	ErrConversion  = errors.New("webconv: conversion error, possible mailformed payload check logs")
	ErrUnsupported = errors.New("webconv: unsupported format")
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
func (c *Converter) MakeReply(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", c.ResultContentType.String())

	if err == ErrConversion {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else if err == ErrUnsupported {
		http.Error(w, "Unsupported format: conventer only accepts XML or JSON", http.StatusUnsupportedMediaType)

	} else {
		io.Copy(w, &c.Result)

	}
}

// SwapFormat fill Input.Result
// if input data is Unsupported they return raw data back
// (TODO: maybe is better to return 415 Unsupported Media Type?)
func (c *Converter) SwapFormat() error {

	switch c.ContentType {

	case JSON:
		return c.json2xml()

	case XML:
		return c.xml2json()

	default:
		return ErrUnsupported

	}

}

// xml2json converts XML in Payload and saves JSON to Result
func (c *Converter) xml2json() error {

	c.ResultContentType = JSON

	mv, err := mxj.NewMapXml(c.Payload.Bytes())
	if err != nil {
		log.Println(err)
		return ErrConversion
	}

	js, err := json.Marshal(mv)
	if err != nil {
		log.Println(err)
		return ErrConversion
	}

	c.Result.Write(js)

	return nil

}

// json2xml converts JSON in Payload and saves XML to Result
func (c *Converter) json2xml() error {

	c.ResultContentType = XML

	var js map[string]interface{}
	if err := json.Unmarshal(c.Payload.Bytes(), &js); err != nil {
		log.Println(err)
		return ErrConversion
	}

	mv := mxj.Map(js)
	xml, err := mv.XmlIndent("", "  ")
	if err != nil {
		log.Println(err)
		return ErrConversion
	}

	c.Result.WriteString(`<?xml version="1.0" encoding="UTF-8"?>` + "\n")
	c.Result.Write(xml)

	return nil

}
