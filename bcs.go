package bcsgo

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
)

type BCS struct {
	ak, sk     string
	httpClient *HttpClient
}

func NewBCS(ak, sk string) *BCS {
	return &BCS{ak, sk, NewHttpClient()}
}

func (this *BCS) ListBuckets() ([]*Bucket, error) {
	link := this.getUrl()
	_, data, err := this.httpClient.Get(link)
	if err != nil {
		return nil, err
	} else {
		list := []*Bucket{}
		err := json.Unmarshal(data, &list)
		for i, _ := range list {
			list[i].bcs = this
		}
		return list, err
	}
}
func (this *BCS) getUrl() string {
	return this.simpleSign(GET, "", "/")
}
func (this *BCS) Bucket(bucketName string) *Bucket {
	return &Bucket{this, bucketName}
}
func (this *BCS) simpleSign(m, b, o string) string {
	return this.Sign(m, b, o, "", "", "")
}
func (this *BCS) Sign(m, b, o, t, i, s string) string {
	flag := ""
	ss := ""
	flag += "M"
	ss += "Method=" + m + "\n"
	flag += "B"
	ss += "Bucket=" + b + "\n"
	flag += "O"
	ss += "Object=" + o + "\n"
	if t != "" {
		flag += "T"
		ss += "Time=" + t + "\n"
	}
	if i != "" {
		flag += "I"
		ss += "Ip=" + i + "\n"
	}
	if s != "" {
		flag += "S"
		ss += "Size=" + s + "\n"
	}
	ss = flag + "\n" + ss
	h := func(sk, body string) string {
		hash := hmac.New(sha1.New, []byte(sk))
		hash.Write([]byte(body))
		digest := hash.Sum(nil)
		sign := base64.StdEncoding.EncodeToString(digest)
		sign = strings.TrimSpace(sign)
		sign = url.QueryEscape(sign)
		return sign
	}
	sign := h(this.sk, ss)
	url := fmt.Sprintf(
		"%s/%s%s?sign=%s:%s:%s",
		BCS_HOST,
		b,
		//"/"+url.QueryEscape(o[1:]),
		url.QueryEscape(o),
		flag,
		this.ak,
		sign)
	if t != "" {
		url += "&time=" + t
	}
	if i != "" {
		url += "&ip=" + i
	}
	if s != "" {
		url += "&size=" + s
	}
	return url
}
