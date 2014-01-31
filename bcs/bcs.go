package bcs

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"net/url"
	"strings"
)

const BCS_HOST = "http://bcs.duapp.com"

type BCS struct {
	ak, sk string
}

func NewBCS(ak, sk string) *BCS {
	return &BCS{ak, sk}
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
		"/"+url.QueryEscape(o[1:]),
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
