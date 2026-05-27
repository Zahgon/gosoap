package gosoap

import (
	"net/http"
	"sync"
	"time"
)

type SoapParams interface{}

// HeaderParams holds params specific to the header
type HeaderParams map[string]interface{}

// Params type is used to set the params in soap request
type Params map[string]interface{}
type ArrayParams [][2]interface{}
type SliceParams []interface{}

type DumpLogger interface {
	LogRequest(method string, dump []byte)
	LogResponse(method string, dump []byte)
}

type fmtLogger struct{}

func (l *fmtLogger) LogRequest(method string, dump []byte) { _ = "STUB: not implemented"; return }

func (l *fmtLogger) LogResponse(method string, dump []byte) { _ = "STUB: not implemented"; return }

// Config config the Client
type Config struct {
	Dump   bool
	Logger DumpLogger
}

// SoapClient return new *Client to handle the requests with the WSDL
func SoapClient(wsdl string, httpClient *http.Client) (*Client, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SoapClientWithConfig return new *Client to handle the requests with the WSDL
func SoapClientWithConfig(wsdl string, httpClient *http.Client, config *Config) (*Client, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Client struct hold all the information about WSDL,
// request and response of the server
type Client struct {
	HTTPClient   *http.Client
	AutoAction   bool
	URL          string
	HeaderName   string
	HeaderParams SoapParams
	Definitions  *wsdlDefinitions
	// Must be set before first request otherwise has no effect, minimum is 15 minutes.
	RefreshDefinitionsAfter time.Duration
	Username                string
	Password                string

	once                 sync.Once
	definitionsErr       error
	onRequest            sync.WaitGroup
	onDefinitionsRefresh sync.WaitGroup
	wsdl                 string
	config               *Config
}

// Call call's the method m with Params p
func (c *Client) Call(m string, p SoapParams) (res *Response, err error) {
	_ = "STUB: not implemented"
	return nil, nil

	// CallByStruct call's by struct
}

func (c *Client) CallByStruct(s RequestStruct) (res *Response, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Client) waitAndRefreshDefinitions(d time.Duration) { _ = "STUB: not implemented"; return }

func (c *Client) initWsdl() { _ = "STUB: not implemented"; return }

// SetWSDL set WSDL url
func (c *Client) SetWSDL(wsdl string) { _ = "STUB: not implemented"; return }

// Do Process Soap Request
func (c *Client) Do(req *Request) (res *Response, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// 15 minute to prevent abuse.

// err = xml.Unmarshal(b, &soap)
// error: xml: encoding "ISO-8859-1" declared but Decoder.CharsetReader is nil
// https://stackoverflow.com/questions/6002619/unmarshal-an-iso-8859-1-xml-input-in-go
// https://github.com/golang/go/issues/8937

type process struct {
	Client     *Client
	Request    *Request
	SoapAction string
	Payload    []byte
}

// doRequest makes new request to the server using the c.Method, c.URL and the body.
// body is enveloped in Do method
func (p *process) doRequest(url string) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (p *process) httpClient() *http.Client { _ = "STUB: not implemented"; return nil }

// ErrorWithPayload error payload schema
type ErrorWithPayload struct {
	error
	Payload []byte
}

// GetPayloadFromError returns the payload of a ErrorWithPayload
func GetPayloadFromError(err error) []byte { _ = "STUB: not implemented"; return nil }

// SoapEnvelope struct
type SoapEnvelope struct {
	XMLName struct{} `xml:"Envelope"`
	Header  SoapHeader
	Body    SoapBody
}

// SoapHeader struct
type SoapHeader struct {
	XMLName  struct{} `xml:"Header"`
	Contents []byte   `xml:",innerxml"`
}

// SoapBody struct
type SoapBody struct {
	XMLName  struct{} `xml:"Body"`
	Contents []byte   `xml:",innerxml"`
}
