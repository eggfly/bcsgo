package bcs

import (
// "fmt"
)

type Bucket struct {
	bcs  *BCS
	Name string
}

func newBucket(bcs *BCS, bucketName string) *Bucket {
	return &Bucket{bcs, bucketName}
}
