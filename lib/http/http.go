package http

import (
	"bytes"
	"crypto/tls"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

type RequestOption struct {
	ProxyUrl string
	Headers  map[string]string
	Timeout  int
	tr       *http.Transport
}

func NewRequestOption(proxyUrl string, timeout int, fxs ...func(opt *RequestOption)) *RequestOption {
	opt := &RequestOption{}
	for _, fx := range fxs {
		fx(opt)
	}
	opt.ProxyUrl = proxyUrl
	if timeout == 0 {
		opt.Timeout = 5
	}
	var proxy func(*http.Request) (*url.URL, error)
	if opt.ProxyUrl != "" {
		tmp, _ := url.Parse(opt.ProxyUrl)
		proxy = http.ProxyURL(tmp)
	}
	opt.tr = &http.Transport{
		Proxy:           proxy,
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	return opt
}

func setRequest(r *http.Request, opt *RequestOption, header map[string]string) {
	if opt.Headers != nil {
		for k, v := range opt.Headers {
			r.Header.Set(k, v)
		}
	}
	if header != nil {
		for k, v := range header {
			r.Header.Set(k, v)
		}
	}
}

func (opt *RequestOption) HttpGet(reqUrl string, header map[string]string, timeout int,
	fxs ...func(opt *RequestOption)) (*http.Response, []byte, error) {
	if timeout == 0 {
		timeout = opt.Timeout
	}
	req, _ := http.NewRequest("GET", reqUrl, nil)
	setRequest(req, opt, header)
	client := &http.Client{
		Transport: opt.tr,
		Timeout:   time.Second * time.Duration(timeout), //超时时间
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, nil, err
	}
	if resp.Body != nil {
		defer resp.Body.Close()
	}
	body, _ := ioutil.ReadAll(resp.Body)
	return resp, body, nil
}

func (opt *RequestOption) HttpPost(reqUrl string, reqData []byte, timeout int) (*http.Response,
	[]byte, error) {
	if timeout == 0 {
		timeout = opt.Timeout
	}
	req, _ := http.NewRequest("POST", reqUrl, bytes.NewReader(reqData))
	setRequest(req, opt, nil)
	client := &http.Client{
		Transport: opt.tr,
		Timeout:   time.Second * time.Duration(timeout), //超时时间
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, nil, err
	}
	if resp.Body != nil {
		defer resp.Body.Close()
	}
	body, _ := ioutil.ReadAll(resp.Body)
	return resp, body, nil
}
