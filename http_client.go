package bcsgo

import (
	"errors"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"strconv"
	"time"
)

type HttpClient struct {
	client *http.Client
}

func NewHttpClient() *HttpClient {
	return &HttpClient{&http.Client{}}
}

func (this *HttpClient) Get(url string) (*http.Response, []byte, error) {
	return this.createAndDoRequestForResult(GET, url, nil, nil)
}
func (this *HttpClient) Put(url string, data io.Reader, size int64, modifyHeader func(header *http.Header)) (*http.Response, []byte, error) {
	customRequest := func(req *http.Request) {
		req.ContentLength = size
		if modifyHeader != nil {
			modifyHeader(&req.Header)
		}
	}
	return this.createAndDoRequestForResult(PUT, url, data, customRequest)
}
func (this *HttpClient) Head(url string) (*http.Response, []byte, error) {
	return this.createAndDoRequestForResult(HEAD, url, nil, nil)
}
func (this *HttpClient) Delete(url string) (*http.Response, []byte, error) {
	return this.createAndDoRequestForResult(DELETE, url, nil, nil)
}
func (this *HttpClient) dumpRequest(req *http.Request) {
	if DEBUG {
		dump, dumpErr := httputil.DumpRequest(req, DEBUG_REQUEST_BODY)
		log.Println(string(dump), dumpErr)
	}
}
func (this *HttpClient) createAndDoRequestForResult(method string, url string, data io.Reader, customRequest func(*http.Request)) (*http.Response, []byte, error) {
	req, err := http.NewRequest(method, url, data)
	if err != nil {
		return nil, nil, err
	}
	if customRequest != nil {
		customRequest(req)
	}
	this.dumpRequest(req)

	var oldTime time.Time
	if DEBUG {
		oldTime = time.Now()
	}
	resp, err := this.client.Do(req)
	if DEBUG {
		log.Println(time.Now().Sub(oldTime))
	}
	respData, err := this.handleResponseContent(resp, err)
	return resp, respData, err
}
func (this *HttpClient) handleResponseContent(resp *http.Response, err error) ([]byte, error) {
	if DEBUG {
		dump, dumpErr := httputil.DumpResponse(resp, true)
		log.Println(string(dump), dumpErr)
	}
	if err != nil {
		return nil, err
	} else {
		return readAllResponseBodyWithError(resp)
	}
}

func mergeResponseError(err error, resp *http.Response) error {
	if err != nil {
		return err
	} else if resp.StatusCode != http.StatusOK {
		return errors.New("request not ok, status: " + strconv.Itoa(resp.StatusCode) +
			", body: " + string(readAllResponseBodyIgnoreError(resp)))
	} else {
		return nil
	}
}

func readAllResponseBodyIgnoreError(resp *http.Response) []byte {
	respData, _ := readAllResponseBodyWithError(resp)
	return respData
}

func readAllResponseBodyWithError(resp *http.Response) ([]byte, error) {
	// defer resp.Body.Close()
	respData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}
	return respData, err
}
