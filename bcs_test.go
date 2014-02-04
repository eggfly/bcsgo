package bcsgo

import (
	"fmt"
	"testing"
)

var bcs = NewBCS("vYlphQiwbhVz67jjW48ddY3C", "mkfr0AYygGjgm4MIC7KBc7qzFOtz9Nha")

func TestSign(t *testing.T) {
	url := bcs.Sign("GET", "", "/", "", "", "")
	url_ex := "http://bcs.duapp.com//?sign=MBO:vYlphQiwbhVz67jjW48ddY3C:yf27Oy6JVtK6nxRtIASKX6H%2BR4I%3D"
	if url != url_ex {
		t.Fail()
	}
}

func TestNewAndDeleteBucketAndACL(t *testing.T) {
	bucketName := "testsml2"
	newBucket := bcs.Bucket(bucketName)
	bucketErr := newBucket.CreateWithACL(ACL_PUBLIC_READ)
	if bucketErr != nil {
		fmt.Println(bucketErr)
		t.Fail()
	}

	bucketACL, bucketACLErr := newBucket.GetACL()
	expectedBucketACL := fmt.Sprintf(`{"statements":[{"action":["*"],"effect":"allow","resource":["testsml2\/"],"user":["psp:egg90"]},{"action":["get_object"],"effect":"allow","resource":["%s\/"],"user":["*"]}]}`, bucketName)
	if bucketACLErr != nil {
		fmt.Println(bucketACLErr)
		t.Fail()
	}
	if bucketACL != expectedBucketACL {
		fmt.Println(bucketACL)
		fmt.Println(expectedBucketACL)
		t.Fail()
	}

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
	bcssdk := bcs.Bucket("bcssdk")
	objects, e := bcssdk.ListObjects("", 0, 5)
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
	bcssdk := bcs.Bucket("bcssdk")
	acl, aclErr := bcssdk.GetACL()
	if aclErr != nil || acl == "" {
		t.Fail()
	}
	putErr := bcssdk.SetACL(ACL_PUBLIC_READ)
	if putErr != nil {
		t.Fail()
	}
}

func PutAndDeleteObjectTest(t *testing.T) {
	bcssdk := bcs.Bucket("bcssdk")
	path := "/testDir/test.txt"
	testObj := bcssdk.Object(path)
	testObj, err := testObj.PutFile("test.txt", ACL_PUBLIC_READ)
	if (err != nil) || testObj.AbsolutePath != path {
		t.Fail()
	}

	deleteErr := testObj.Delete()
	if deleteErr != nil {
		t.Fail()
	}
}
