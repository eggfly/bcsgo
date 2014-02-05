package bcsgo

import (
	// "encoding/json"
	"net/http"
)

type Superfile struct {
	Object
	Objects []*Object
}

func (this *Superfile) putSuperfileUrl() string {
	return this.putUrl() + "&superfile=1"
}
func (this *Superfile) Put() error {
	link := this.putSuperfileUrl()
	modifyHeader := func(header *http.Header) {
	}
	resp, _, err := this.bucket.bcs.httpClient.Put(link, nil, 0, modifyHeader)
	if err == nil {
		this.ContentMD5 = resp.Header.Get("Content-MD5")
		this.VersionKey = resp.Header.Get("X-Bs-Version")
	}
	return err
}
