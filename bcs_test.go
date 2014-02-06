package bcsgo

import (
	"testing"
)

var ak = "zaTGAk9k6qoRaVoVcTCRGbjZ"
var sk = "r7ay1xOM12s4afPUqRZ9f53su8OF6lwj"
var bcs = NewBCS(ak, sk)

func init() {
	// DEBUG = true
}

func TestBCSListBuckets(t *testing.T) {
	buckets, e := bcs.ListBuckets()
	if e != nil {
		t.Error(e)
	}
	if buckets == nil {
		t.Error("buckets list is nil")
	}
}

func TestBCSNewBucket(t *testing.T) {
	bucket := bcs.Bucket("mockBucket")
	if bucket == nil {
		t.Error("new bucket shouldn't be nil")
	}
}
