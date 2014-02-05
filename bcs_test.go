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
		t.Error(err)
	}
}

func TestNewAndDeleteBucketAndACL(t *testing.T) {
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

func TestNewBucketWithInvalidName(t *testing.T) {
	newBucket := bcs.Bucket("testErrorBucket")
	bucketErr := newBucket.Create()
	// It shall be failed.
	if bucketErr == nil {
		t.Error("create bucket with invaid name should failed")
	}
}

func TestListBuckets(t *testing.T) {
	buckets, e := bcs.ListBuckets()
	if e != nil {
		t.Error(e)
	}
	if buckets == nil {
		t.Error("buckets list is nil")
	}
}

func TestListObjects(t *testing.T) {
	// todo prefix
	bucket := bcs.Bucket(sessionBucketName)
	objects, e := bucket.ListObjects("", 0, 5)
	if e != nil {
		t.Error("object list shouldn't be nil")
	}
	for _, pObject := range objects.Objects {
		if pObject == nil {
			t.Error("object should not be nil")
		}
	}
}

func TestBucketACL(t *testing.T) {
	bucket := bcs.Bucket(sessionBucketName)
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
}

func TestPutAndDeleteObject(t *testing.T) {
	// todo file name with blank char
	bucket := bcs.Bucket(sessionBucketName)
	path := "/testDir/testwithblank.txt"
	testObj := bucket.Object(path)
	testObj, err := testObj.PutFile("test.txt", ACL_PUBLIC_READ)
	if err != nil {
		t.Error(err)
	}
	if testObj.AbsolutePath != path {
		t.Error("testObj.AbsolutePath != path", testObj.AbsolutePath, path)
	}

	expectedPublicLink := fmt.Sprintf("%s/%s%s", BCS_HOST, sessionBucketName, path)
	publicLink := testObj.PublicLink()
	if expectedPublicLink != publicLink {
		t.Error("expectedPublicLink != publicLink", expectedPublicLink, publicLink)
	}

	deleteErr := testObj.Delete()
	if deleteErr != nil {
		t.Error(deleteErr)
	}
}

func TestFinallyDeleteSessionBucket(t *testing.T) {
	bucket := bcs.Bucket(sessionBucketName)
	err := bucket.Delete()
	if err != nil {
		t.Error(err)
	}
}
