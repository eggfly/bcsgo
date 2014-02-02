package bcsgo

import (
	// "encoding/json"
	"net/http"
	"os"
	"strconv"
)

type Object struct {
	bucket       *Bucket
	VersionKey   string `json:"version_key"`
	AbsolutePath string `json:"object"`
	Superfile    string `json:"superfile"`
	Size         uint64 `json:"size,string"`
	ParentDir    string `json:"parent_dir"`
	IsDir        string `json:"is_dir"`
	MDatetime    string `json:"mdatetime"`
	RefKey       string `json:"ref_key"`
	ContentMD5   string `json:"content_md5"`
}

func (this *Object) putUrl() string {
	return this.bucket.bcs.simpleSign(PUT, this.bucket.Name, this.AbsolutePath)
}

func (this *Object) PutFile(localFile string) (*Object, error) {
	link := this.putUrl()
	file, err := os.Open(localFile)
	if err != nil {
		return nil, err
	}
	resp, _, err := this.bucket.bcs.httpClient.Put(link, file)
	if err != nil || resp.StatusCode != http.StatusOK {
		return nil, err
	} else {
		this.ContentMD5 = resp.Header["Content-Md5"][0]
		this.VersionKey = resp.Header["X-Bs-Version"][0]
		this.Size, _ = strconv.ParseUint(resp.Header["X-Bs-File-Size"][0], 10, 64)
		return this, err
	}
}
