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
	HEADER_ACL = "X-Bs-Acl"
)

var DEBUG bool = false
var DEBUG_REQUEST_BODY = false
