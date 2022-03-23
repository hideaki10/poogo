package fetch

import "net/http"

type Request struct {
	Method  string            `json:"method"`
	Url     string            `json:"url"`
	Header  map[string]string `json:"header"`
	Body    string            `json:"body,omitempty"` // omitempty is for json marshal
	Retries int               `json:"retries"`
	Timeout int64             `json:"timeout"`
}

func NewGetRequest(url string, retries int, timeout int64, header map[string]string) *Request {
	return &Request{
		Method:  http.MethodGet,
		Url:     url,
		Header:  header,
		Retries: retries,
		Timeout: timeout,
	}
}

func DefaultClient(url string, header map[string]string) *Request {
	return NewGetRequest(url, 3, 3000, header)
}

type Response struct {
	StatusCode int               `json:"status_code"`
	URL        string            `json:"url"`
	Header     map[string]string `json:"header"`
	Body       string            `json:"body",omitempty` // omitempty is for json marshal
}
