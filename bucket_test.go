package bcsgo

import (
	"strconv"
	"testing"
	"time"
)

func randomGlobalBucketName(index int) string {
	return ak[:2] + "-sdk-" + strconv.FormatInt(time.Now().Unix(), 10)[5:] + "-" + strconv.Itoa(index)
}

func createBucketTempForTest(t *testing.T) *Bucket {
	bucketName := randomGlobalBucketName(1)
	bucket := bcs.Bucket(bucketName)
	err := bucket.Create()
	if err != nil {
		t.Error(err)
	}
	return bucket
}

func deleteBucketForTest(t *testing.T, bucket *Bucket) {
	deleteErr := bucket.Delete()
	if deleteErr != nil {
		t.Error(deleteErr)
	}
}

func TestBucketCreateWithACLAndDelete(t *testing.T) {
	bucketName := randomGlobalBucketName(1)
	newBucket := bcs.Bucket(bucketName)
	bucketErr := newBucket.CreateWithACL(ACL_PUBLIC_READ)
	if bucketErr != nil {
		t.Error(bucketErr)
	}

	// bucketACL, bucketACLErr := newBucket.GetACL()
	// expectedBucketACL := fmt.Sprintf(`{"statements":[{"action":["*"],"effect":"allow","resource":["testsml2\/"],"user":["psp:egg90"]},{"action":["get_object"],"effect":"allow","resource":["%s\/"],"user":["*"]}]}`, bucketName)
	// if bucketACLErr != nil {
	// 	fmt.Println(bucketACLErr)
	// 	t.Fail()
	// }
	// if bucketACL != expectedBucketACL {
	// 	fmt.Println(bucketACL)
	// 	fmt.Println(expectedBucketACL)
	// 	t.Fail()
	// }

	bucketErr = newBucket.Delete()
	if bucketErr != nil {
		t.Error(bucketErr)
	}
}

func TestBucketCreateWithInvalidName(t *testing.T) {
	newBucket := bcs.Bucket("testErrorBucket")
	bucketErr := newBucket.Create()
	// It shall be failed.
	if bucketErr == nil {
		t.Error("create bucket with invaid name should failed")
	}
}

func TestBucketListObjects(t *testing.T) {
	bucket := createBucketTempForTest(t)

	// todo prefix
	objects, e := bucket.ListObjects("", 0, 5)
	if e != nil {
		t.Error("object list shouldn't be nil")
	}
	for _, pObject := range objects.Objects {
		if pObject == nil {
			t.Error("object should not be nil")
		}
	}

	deleteBucketForTest(t, bucket)
}

func TestBucketACL(t *testing.T) {
	bucket := createBucketTempForTest(t)

	acl, aclErr := bucket.GetACL()
	if aclErr != nil {
		t.Error(aclErr)
	}
	if acl == "" {
		t.Error("acl string shouldn't be nil")
	}
	putErr := bucket.SetACL(ACL_PUBLIC_READ)
	if putErr != nil {
		t.Error(putErr)
	}

	deleteBucketForTest(t, bucket)
	t.Fail()
}
