package bcsgo

const (
	GET    = "GET"
	POST   = "POST"
	PUT    = "PUT"
	HEAD   = "HEAD"
	DELETE = "DELETE"
)

const (
	BCS_HOST = "http://bcs.duapp.com"
)

const (
	ACL_PRIVATE           = "private"
	ACL_PUBLIC_READ       = "public-read"
	ACL_PUBLIC_WRITE      = "public-write"
	ACL_PUBLIC_READ_WRITE = "public-read-write"
	ACL_PUBLIC_CONTROL    = "public-control"
)

const (
	HEADER_COPY_SOURCE = "x-bs-copy-source"
	HEADER_ACL         = "X-Bs-Acl"
	HEADER_VERSION     = "X-Bs-Version"
	HEADER_FILESIZE    = "X-Bs-File-Size"
	HEADER_ETAG        = "Etag"
	HEADER_CONTENT_MD5 = "Content-Md5"
)

var DEBUG bool = false
var DEBUG_REQUEST_BODY = false
