package bcsgo

import (
	"fmt"
	"os"
	"strings"
	"testing"
	"time"
)

var bucketForObjectTest *Bucket

const (
	_LARGER_NAME = "256KB.data"
	_TEST_NAME   = "test.txt"
)

func createTestFile(filename string, size int) {
	file, _ := os.Create(filename)
	file.WriteString(strings.Repeat(" ", size))
	file.Close()
}
func deleteTestFile(filename string) {
	os.Remove(filename)
}

func TestObjectInit(t *testing.T) {
	bucket := createBucketTempForTest(t)
	bucketForObjectTest = bucket

	createTestFile(_LARGER_NAME, 256*1024)
	createTestFile(_TEST_NAME, 1024)
}

func TestObjectPutAndListAndHeadAndDeleteObject(t *testing.T) {
	bucket := bucketForObjectTest
	// todo file name with blank char
	path := "/testDir/testwithblank.txt"
	testObj := bucket.Object(path)
	testObj, err := testObj.PutFileWithACL(_TEST_NAME, ACL_PUBLIC_READ)
	if err != nil {
		t.Error(err)
	}
	if testObj.AbsolutePath != path {
		t.Error("testObj.AbsolutePath != path", testObj.AbsolutePath, path)
	}

	// expectedPublicLink := fmt.Sprintf("%s/%s%s", BCS_HOST, bucket.Name, path)
	// publicLink := testObj.PublicLink()
	// if expectedPublicLink != publicLink {
	// 	t.Error("expectedPublicLink != publicLink", expectedPublicLink, publicLink)
	// }

	headErr := testObj.Head()
	if headErr != nil {
		t.Error(headErr)
	}
	if testObj.ContentMD5 == "" || testObj.VersionKey == "" {
		t.Error("Info after HEAD is not ok!")
	}

	listObjectsTest := func(prefix string, start, limit, expectedCount int) {
		objects, e := bucket.ListObjects(prefix, start, limit)
		if e != nil {
			t.Error("object list shouldn't be nil")
		}
		for _, pObject := range objects.Objects {
			if pObject == nil {
				t.Error("object should not be nil")
			}
		}
		resultCount := len(objects.Objects)
		if expectedCount != resultCount {
			t.Error(fmt.Sprintf(`expectedCount != result, expectedCount = %d, resultCount = %d`, expectedCount, resultCount))
		}
	}

	listObjectsTest("", 0, 100, 1)
	listObjectsTest("/", 1, 200, 0)
	listObjectsTest("/testDir/testwithblank.txt", 0, 1, 1)
	listObjectsTest("/testDir/testwithblank.txt!", 0, 2, 0)

	deleteErr := testObj.Delete()
	if deleteErr != nil {
		t.Error(deleteErr)
	}
}

func TestObjectLargerSingleFileAndACL(t *testing.T) {
	bucket := bucketForObjectTest
	obj := bucket.Object("/larger.data")
	obj, err := obj.PutFile(_LARGER_NAME)
	if err != nil {
		t.Error(err)
	}

	acl, aclErr := obj.GetACL()
	if aclErr != nil {
		t.Error(aclErr)
	}
	if acl == "" {
		t.Error("acl string shouldn't be nil")
	}

	setACLCheckError := func(acl string) {
		putErr := obj.SetACL(acl)
		if putErr != nil {
			t.Error(putErr)
		}
	}

	setACLCheckError(ACL_PUBLIC_CONTROL)
	setACLCheckError(ACL_PUBLIC_READ)
	setACLCheckError(ACL_PUBLIC_WRITE)
	setACLCheckError(ACL_PUBLIC_READ_WRITE)
	setACLCheckError(ACL_PRIVATE)

	deleteErr := obj.Delete()
	if deleteErr != nil {
		t.Error(deleteErr)
	}
}

func TestObjectFinalize(t *testing.T) {
	time.Sleep(500 * time.Millisecond)
	deleteBucketForTest(t, bucketForObjectTest)
	bucketForObjectTest = nil
	deleteTestFile(_LARGER_NAME)
	deleteTestFile(_TEST_NAME)
}
