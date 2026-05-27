package gosoap

// Request Soap Request
type Request struct {
	Method string
	Params SoapParams
}

func NewRequest(m string, p SoapParams) *Request { _ = "STUB: not implemented"; return nil }

// RequestStruct soap request interface
type RequestStruct interface {
	SoapBuildRequest() *Request
}

// NewRequestByStruct create a new request using builder
func NewRequestByStruct(s RequestStruct) (*Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
