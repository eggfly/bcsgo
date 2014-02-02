package bcsgo

import (
// "fmt"
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
