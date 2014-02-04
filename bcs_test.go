package bcsgo

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

var ak = "zaTGAk9k6qoRaVoVcTCRGbjZ"
var sk = "r7ay1xOM12s4afPUqRZ9f53su8OF6lwj"
var bcs = NewBCS(ak, sk)

var sessionBucketName = randomGlobalBucketName(0)

func randomGlobalBucketName(index int) string {
	return ak[:2] + "-sdk-" + strconv.FormatInt(time.Now().Unix(), 10)[5:] + "-" + strconv.Itoa(index)
}

func init() {
	// DEBUG = true
}

// test function must starts with "Test"
func TestSign(t *testing.T) {
	// url := bcs.Sign("GET", "", "/", "", "", "")
	// url_ex := "http://bcs.duapp.com//?sign=MBO:vYlphQiwbhVz67jjW48ddY3C:yf27Oy6JVtK6nxRtIASKX6H%2BR4I%3D"
	// if url != url_ex {
	// 	t.Fail()
	// }
}

func TestSimpleCreateBucket(t *testing.T) {
	newBucket := bcs.Bucket(sessionBucketName)
	err := newBucket.Create()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
}

func TestNewAndDeleteBucketAndACL(t *testing.T) {
	bucketName := randomGlobalBucketName(1)
	newBucket := bcs.Bucket(bucketName)
	bucketErr := newBucket.CreateWithACL(ACL_PUBLIC_READ)
	if bucketErr != nil {
		fmt.Println(bucketErr)
		t.Fail()
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
		fmt.Println(bucketErr)
		t.Fail()
	}
}

func TestNewBucketWithInvalidName(t *testing.T) {
	newBucket := bcs.Bucket("testErrorBucket")
	bucketErr := newBucket.Create()
	// It shall be failed.
	if bucketErr == nil {
		t.Fail()
	}
}

func TestListBuckets(t *testing.T) {
	buckets, e := bcs.ListBuckets()
	if e != nil {
		t.Fail()
	}
	if buckets == nil {
		t.Fail()
	}
}

func TestListObjects(t *testing.T) {
	// todo prefix
	bucket := bcs.Bucket(sessionBucketName)
	objects, e := bucket.ListObjects("", 0, 5)
	if e != nil {
		t.Fail()
	}
	for _, pObject := range objects.Objects {
		if pObject == nil {
			t.Fail()
		}
	}
}

func TestBucketACL(t *testing.T) {
	bucket := bcs.Bucket(sessionBucketName)
	acl, aclErr := bucket.GetACL()
	if aclErr != nil || acl == "" {
		t.Fail()
	}
	putErr := bucket.SetACL(ACL_PUBLIC_READ)
	if putErr != nil {
		t.Fail()
	}
}

func TestPutAndDeleteObject(t *testing.T) {
	bucket := bcs.Bucket(sessionBucketName)
	path := "/testDir/test.txt"
	testObj := bucket.Object(path)
	testObj, err := testObj.PutFile("test.txt", ACL_PUBLIC_READ)
	if (err != nil) || testObj.AbsolutePath != path {
		t.Fail()
	}

	deleteErr := testObj.Delete()
	if deleteErr != nil {
		t.Fail()
	}
}

func TestFinallyDeleteSessionBucket(t *testing.T) {
	bucket := bcs.Bucket(sessionBucketName)
	err := bucket.Delete()
	if err != nil {
		t.Fail()
	}
}
