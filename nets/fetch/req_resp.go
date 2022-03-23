package fetch

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/cnych/stardust/timex"
)

func Fetch(req *Request) (*http.Response, error) {

	if req == nil {
		return nil, fmt.Errorf("request is nil")
	}
	if req.Url == "" {
		return nil, fmt.Errorf("request url is nil")
	}
	if req.Method == "" {
		req.Method = http.MethodGet
	}

	var body io.Reader
	if req.Body != "" {
		body = strings.NewReader(req.Body)
	}

	newRequest, err := http.NewRequest(req.Method, req.Url, body)
	if err != nil {
		return nil, err
	}

	header := http.Header{}

	for k, v := range req.Header {
		header.Add(k, v)
	}

	newRequest.Header = header

	var timeout time.Duration

	if req.Timeout > 0 {
		timeout = timex.DurationMS(req.Timeout)
	}

	client := http.Client{Timeout: timeout}

	newResponse, err := client.Do(newRequest)
	if err != nil {
		return nil, err
	}

	return newResponse, nil
}

func Retry(retries int, f func() error) error {

	err := f()

	if err == nil {
		return nil
	}
	if retries <= 0 {
		return err
	}
	for i := 0; i < retries; i++ {
		err = f()
		if err == nil {
			return nil
		}
		if retries <= 0 {
			return err
		}
		for i := 0; i < retries; i++ {
			err = f()
			if err == nil {
				return nil
			}
		}

	}
	return err
}
