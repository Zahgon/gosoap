package gosoap

// Response Soap Response
type Response struct {
	Body    []byte
	Header  []byte
	Payload []byte
}

// FaultError implements error interface
type FaultError struct {
	fault *Fault
}

func (e FaultError) Error() string { _ = "STUB: not implemented"; return "" }

// IsFault returns whether the given error is a fault error or not.
//
// IsFault will return false when the error could not be typecasted to FaultError, because
// every fault error should have it's dynamic type as FaultError.
func IsFault(err error) bool { _ = "STUB: not implemented"; return false }

// Unmarshal get the body and unmarshal into the interface
func (r *Response) Unmarshal(v interface{}) error { _ = "STUB: not implemented"; return nil }
