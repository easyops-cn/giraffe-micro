// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: code.proto

package status

import (
	fmt "fmt"
	math "math"

	proto "github.com/gogo/protobuf/proto"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type Code int32

const (
	// Not an error; returned on success
	//
	// HTTP Mapping: 200 OK
	Code_OK Code = 0
	// The operation was cancelled, typically by the caller.
	//
	// HTTP Mapping: 499 Client Closed Request
	Code_CANCELLED Code = 1
	// Unknown error.  For example, this error may be returned when
	// a `Status` value received from another address space belongs to
	// an error space that is not known in this address space.  Also
	// errors raised by APIs that do not return enough error information
	// may be converted to this error.
	//
	// HTTP Mapping: 500 Internal Server Error
	Code_UNKNOWN Code = 2
	// The client specified an invalid argument.  Note that this differs
	// from `FAILED_PRECONDITION`.  `INVALID_ARGUMENT` indicates arguments
	// that are problematic regardless of the state of the system
	// (e.g., a malformed file name).
	//
	// HTTP Mapping: 400 Bad Request
	Code_INVALID_ARGUMENT Code = 100000
	// The deadline expired before the operation could complete. For operations
	// that change the state of the system, this error may be returned
	// even if the operation has completed successfully.  For example, a
	// successful response from a server could have been delayed long
	// enough for the deadline to expire.
	//
	// HTTP Mapping: 504 Gateway Timeout
	Code_DEADLINE_EXCEEDED Code = 100014
	// Some requested entity (e.g., file or directory) was not found.
	//
	// Note to server developers: if a request is denied for an entire class
	// of users, such as gradual feature rollout or undocumented whitelist,
	// `NOT_FOUND` may be used. If a request is denied for some users within
	// a class of users, such as user-based access control, `PERMISSION_DENIED`
	// must be used.
	//
	// HTTP Mapping: 404 Not Found
	Code_NOT_FOUND Code = 100005
	// The entity that a client attempted to create (e.g., file or directory)
	// already exists.
	//
	// HTTP Mapping: 409 Conflict
	Code_ALREADY_EXISTS Code = 100007
	// The caller does not have permission to execute the specified
	// operation. `PERMISSION_DENIED` must not be used for rejections
	// caused by exhausting some resource (use `RESOURCE_EXHAUSTED`
	// instead for those errors). `PERMISSION_DENIED` must not be
	// used if the caller can not be identified (use `UNAUTHENTICATED`
	// instead for those errors). This error code does not imply the
	// request is valid or the requested entity exists or satisfies
	// other pre-conditions.
	//
	// HTTP Mapping: 403 Forbidden
	Code_PERMISSION_DENIED Code = 100004
	// The request does not have valid authentication credentials for the
	// operation.
	//
	// HTTP Mapping: 401 Unauthorized
	Code_UNAUTHENTICATED Code = 100003
	// Some resource has been exhausted, perhaps a per-user quota, or
	// perhaps the entire file system is out of space.
	//
	// HTTP Mapping: 429 Too Many Requests
	Code_RESOURCE_EXHAUSTED Code = 100008
	// The operation was rejected because the system is not in a state
	// required for the operation's execution.  For example, the directory
	// to be deleted is non-empty, an rmdir operation is applied to
	// a non-directory, etc.
	//
	// Service implementors can use the following guidelines to decide
	// between `FAILED_PRECONDITION`, `ABORTED`, and `UNAVAILABLE`:
	//  (a) Use `UNAVAILABLE` if the client can retry just the failing call.
	//  (b) Use `ABORTED` if the client should retry at a higher level
	//      (e.g., when a client-specified test-and-set fails, indicating the
	//      client should restart a read-modify-write sequence).
	//  (c) Use `FAILED_PRECONDITION` if the client should not retry until
	//      the system state has been explicitly fixed.  E.g., if an "rmdir"
	//      fails because the directory is non-empty, `FAILED_PRECONDITION`
	//      should be returned since the client should not retry unless
	//      the files are deleted from the directory.
	//
	// HTTP Mapping: 400 Bad Request
	Code_FAILED_PRECONDITION Code = 100001
	// The operation was aborted, typically due to a concurrency issue such as
	// a sequencer check failure or transaction abort.
	//
	// See the guidelines above for deciding between `FAILED_PRECONDITION`,
	// `ABORTED`, and `UNAVAILABLE`.
	//
	// HTTP Mapping: 409 Conflict
	Code_ABORTED Code = 100006
	// The operation was attempted past the valid range.  E.g., seeking or
	// reading past end-of-file.
	//
	// Unlike `INVALID_ARGUMENT`, this error indicates a problem that may
	// be fixed if the system state changes. For example, a 32-bit file
	// system will generate `INVALID_ARGUMENT` if asked to read at an
	// offset that is not in the range [0,2^32-1], but it will generate
	// `OUT_OF_RANGE` if asked to read from an offset past the current
	// file size.
	//
	// There is a fair bit of overlap between `FAILED_PRECONDITION` and
	// `OUT_OF_RANGE`.  We recommend using `OUT_OF_RANGE` (the more specific
	// error) when it applies so that callers who are iterating through
	// a space can easily look for an `OUT_OF_RANGE` error to detect when
	// they are done.
	//
	// HTTP Mapping: 400 Bad Request
	Code_OUT_OF_RANGE Code = 100002
	// The operation is not implemented or is not supported/enabled in this
	// service.
	//
	// HTTP Mapping: 501 Not Implemented
	Code_UNIMPLEMENTED Code = 100012
	// Internal errors.  This means that some invariants expected by the
	// underlying system have been broken.  This error code is reserved
	// for serious errors.
	//
	// HTTP Mapping: 500 Internal Server Error
	Code_INTERNAL Code = 100011
	// The service is currently unavailable.  This is most likely a
	// transient condition, which can be corrected by retrying with
	// a backoff.
	//
	// See the guidelines above for deciding between `FAILED_PRECONDITION`,
	// `ABORTED`, and `UNAVAILABLE`.
	//
	// HTTP Mapping: 503 Service Unavailable
	Code_UNAVAILABLE Code = 100013
	// Unrecoverable data loss or corruption.
	//
	// HTTP Mapping: 500 Internal Server Error
	Code_DATA_LOSS Code = 100009
)

var Code_name = map[int32]string{
	0:      "OK",
	1:      "CANCELLED",
	2:      "UNKNOWN",
	100000: "INVALID_ARGUMENT",
	100014: "DEADLINE_EXCEEDED",
	100005: "NOT_FOUND",
	100007: "ALREADY_EXISTS",
	100004: "PERMISSION_DENIED",
	100003: "UNAUTHENTICATED",
	100008: "RESOURCE_EXHAUSTED",
	100001: "FAILED_PRECONDITION",
	100006: "ABORTED",
	100002: "OUT_OF_RANGE",
	100012: "UNIMPLEMENTED",
	100011: "INTERNAL",
	100013: "UNAVAILABLE",
	100009: "DATA_LOSS",
}

var Code_value = map[string]int32{
	"OK":                  0,
	"CANCELLED":           1,
	"UNKNOWN":             2,
	"INVALID_ARGUMENT":    100000,
	"DEADLINE_EXCEEDED":   100014,
	"NOT_FOUND":           100005,
	"ALREADY_EXISTS":      100007,
	"PERMISSION_DENIED":   100004,
	"UNAUTHENTICATED":     100003,
	"RESOURCE_EXHAUSTED":  100008,
	"FAILED_PRECONDITION": 100001,
	"ABORTED":             100006,
	"OUT_OF_RANGE":        100002,
	"UNIMPLEMENTED":       100012,
	"INTERNAL":            100011,
	"UNAVAILABLE":         100013,
	"DATA_LOSS":           100009,
}

func (x Code) String() string {
	return proto.EnumName(Code_name, int32(x))
}

func (Code) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_6e9b0151640170c3, []int{0}
}

func init() {
	proto.RegisterEnum("status.Code", Code_name, Code_value)
}

func init() { proto.RegisterFile("code.proto", fileDescriptor_6e9b0151640170c3) }

var fileDescriptor_6e9b0151640170c3 = []byte{
	// 343 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x24, 0x90, 0x4b, 0x6e, 0x53, 0x41,
	0x10, 0x45, 0x79, 0x16, 0xea, 0x90, 0x4a, 0x9c, 0x54, 0x2a, 0x7c, 0x27, 0x5e, 0x00, 0x03, 0x26,
	0xac, 0xa0, 0xfc, 0xba, 0x9c, 0xb4, 0xd2, 0xae, 0xb6, 0xfa, 0x13, 0xc2, 0xa8, 0x05, 0x24, 0xe3,
	0x87, 0x88, 0x59, 0x46, 0xef, 0x81, 0xff, 0x47, 0x7c, 0xd6, 0xc1, 0x10, 0x89, 0x0d, 0x20, 0xaf,
	0x04, 0x3d, 0x3c, 0xbe, 0xf7, 0xea, 0xea, 0x1c, 0x80, 0x17, 0xc3, 0xe5, 0xd5, 0xa3, 0x97, 0xaf,
	0x86, 0xf5, 0x40, 0xe6, 0x7a, 0xfd, 0x6c, 0xfd, 0xfa, 0xfa, 0xe1, 0x9f, 0x09, 0xdc, 0xec, 0x87,
	0xcb, 0x2b, 0x32, 0x30, 0x09, 0x67, 0x78, 0x83, 0xa6, 0xb0, 0xdb, 0xb3, 0xf6, 0xe2, 0xbd, 0x58,
	0xec, 0x68, 0x0f, 0x76, 0x8a, 0x9e, 0x69, 0x78, 0xa2, 0x38, 0xa1, 0xbb, 0x80, 0x4e, 0xcf, 0xd9,
	0x3b, 0x5b, 0x39, 0x9e, 0x94, 0xa5, 0x68, 0xc6, 0x37, 0xcd, 0xd0, 0x3d, 0x38, 0xb2, 0xc2, 0xd6,
	0x3b, 0x95, 0x2a, 0x17, 0xbd, 0x88, 0x15, 0x8b, 0x3f, 0x9b, 0xa1, 0x43, 0xd8, 0xd5, 0x90, 0xeb,
	0x22, 0x14, 0xb5, 0xf8, 0xb1, 0x19, 0xba, 0x0d, 0x07, 0xec, 0xa3, 0xb0, 0x7d, 0x5a, 0xe5, 0xc2,
	0xa5, 0x9c, 0xf0, 0xf3, 0x76, 0xbf, 0x92, 0xb8, 0x74, 0x29, 0xb9, 0xa0, 0xd5, 0x8a, 0x3a, 0xb1,
	0xf8, 0xa1, 0x19, 0xba, 0x03, 0x87, 0x45, 0xb9, 0xe4, 0x53, 0xd1, 0xec, 0x7a, 0xce, 0x62, 0xf1,
	0x7d, 0x33, 0x74, 0x1f, 0x28, 0x4a, 0x0a, 0x25, 0xf6, 0xe3, 0xdf, 0x29, 0x97, 0x34, 0x26, 0x5f,
	0x9a, 0xa1, 0x07, 0x70, 0xbc, 0x60, 0xe7, 0xc5, 0xd6, 0x55, 0x94, 0x3e, 0xa8, 0x75, 0xd9, 0x05,
	0xc5, 0xb7, 0xcd, 0xd0, 0x14, 0x76, 0x78, 0x1e, 0xe2, 0xd8, 0xfc, 0xd4, 0x0c, 0x11, 0xec, 0x87,
	0x92, 0x6b, 0x58, 0xd4, 0xc8, 0x7a, 0x22, 0xf8, 0xae, 0x19, 0x3a, 0x86, 0x69, 0x51, 0xb7, 0x5c,
	0x79, 0x19, 0xd1, 0xc4, 0xe2, 0xf7, 0x66, 0xe8, 0x00, 0x6e, 0x39, 0xcd, 0x12, 0x95, 0x3d, 0x7e,
	0x6b, 0x86, 0x8e, 0x60, 0xaf, 0x28, 0x9f, 0xb3, 0xf3, 0x3c, 0xf7, 0x82, 0x3f, 0xb6, 0x98, 0x96,
	0x33, 0x57, 0x1f, 0x52, 0xc2, 0xaf, 0xcd, 0xcc, 0xf7, 0x7f, 0x6d, 0x66, 0xdd, 0xef, 0xcd, 0xac,
	0xfb, 0xbb, 0x99, 0x75, 0xcf, 0xcd, 0x7f, 0xe5, 0x8f, 0xff, 0x05, 0x00, 0x00, 0xff, 0xff, 0x89,
	0x33, 0xb9, 0x5f, 0x80, 0x01, 0x00, 0x00,
}
