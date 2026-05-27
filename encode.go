package gosoap

import (
	"encoding/xml"
)

var (
	soapPrefix                            = "soap"
	customEnvelopeAttrs map[string]string = nil
)

// SetCustomEnvelope define customizated envelope
func SetCustomEnvelope(prefix string, attrs map[string]string) { _ = "STUB: not implemented"; return }

// MarshalXML envelope the body and encode to xml
func (c process) MarshalXML(e *xml.Encoder, _ xml.StartElement) error {
	_ = "STUB: not implemented"
	return nil

	//start envelope
}

//end envelope

type tokenData struct {
	data []xml.Token
}

func (tokens *tokenData) recursiveEncode(hm interface{}) { _ = "STUB: not implemented"; return }

func (tokens *tokenData) startEnvelope() { _ = "STUB: not implemented"; return }

func (tokens *tokenData) endEnvelope() { _ = "STUB: not implemented"; return }

func (tokens *tokenData) startHeader(m, n string) { _ = "STUB: not implemented"; return }

func (tokens *tokenData) endHeader(m string) { _ = "STUB: not implemented"; return }

func (tokens *tokenData) startBody(m, n string) error { _ = "STUB: not implemented"; return nil }

// endToken close body of the envelope
func (tokens *tokenData) endBody(m string) { _ = "STUB: not implemented"; return }
