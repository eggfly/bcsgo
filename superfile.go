package bcsgo

import (
	// "encoding/json"
	"fmt"
	// "io"
	// "net/http"
	"strings"
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
	parts := make([]string, 0)
	for i, item := range this.Objects {
		parts = append(parts, fmt.Sprintf(`"part_%d": {"url": "%s", "etag":"%s"}`, i, item.refStr(), item.ContentMD5))
	}
	partsStr := strings.Join(parts, ",")
	meta := fmt.Sprintf(`{"object_list": {%s}}`, partsStr)
	reader := strings.NewReader(meta)
	resp, _, err := this.bucket.bcs.httpClient.Put(link, reader, int64(len(meta)), nil)
	if err == nil {
		this.ContentMD5 = resp.Header.Get("Content-MD5")
		this.VersionKey = resp.Header.Get("X-Bs-Version")
	}
	return err
}
