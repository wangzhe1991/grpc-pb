// Code generated by protoc-gen-go-errors. DO NOT EDIT.

package errorx

import (
	fmt "fmt"
	gerr "github.com/wangzhe1991/protoc-gen-go-errorx/gerr"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
const _ = gerr.SupportPackageIsVersion1

var bizErrorCodeMap map[string]int = map[string]int{

	"TestErrorReason_OK":           100002,
	"TestErrorReason_Err":          100003,
	"TestErrorReason_Unknown":      100004,
	"TestErrorReason_DB":           100005,
	"TestErrorReason_Redis":        100006,
	"TestErrorReason_GRPC":         100007,
	"TestErrorReason_NotFound":     100008,
	"TestErrorReason_NoPermission": 100009,
	"TestErrorReason_Existed":      100010,
	"TestErrorReason_ParamInvalid": 100011,
}

var OK = gerr.New(200, 100002, "TestErrorReason_OK", "ok")
var Err = gerr.New(500, 100003, "TestErrorReason_Err", "internal server failed")
var Unknown = gerr.New(500, 100004, "TestErrorReason_Unknown", "unknown failed")
var DB = gerr.New(500, 100005, "TestErrorReason_DB", "internal db failed")
var Redis = gerr.New(500, 100006, "TestErrorReason_Redis", "internal redis failed")
var GRPC = gerr.New(500, 100007, "TestErrorReason_GRPC", "internal grpc failed")
var NotFound = gerr.New(400, 100008, "TestErrorReason_NotFound", "not found")
var NoPermission = gerr.New(400, 100009, "TestErrorReason_NoPermission", "no permission")
var Existed = gerr.New(400, 100010, "TestErrorReason_Existed", "already existed")
var ParamInvalid = gerr.New(400, 100011, "TestErrorReason_ParamInvalid", "invalid param")

// IsOK ok
func IsOK(err error) bool {
	if err == nil {
		return false
	}
	e := gerr.FromError(err)
	return e.Reason == "TestErrorReason_OK" && e.HttpCode == 200
}

// ErrorOK ok
func ErrorOK(format string, args ...interface{}) *gerr.Error {
	return gerr.New(200, 100002, "TestErrorReason_OK", fmt.Sprintf(format, args...))
}

// IsErr internal server failed
func IsErr(err error) bool {
	if err == nil {
		return false
	}
	e := gerr.FromError(err)
	return e.Reason == "TestErrorReason_Err" && e.HttpCode == 500
}

// ErrorErr internal server failed
func ErrorErr(format string, args ...interface{}) *gerr.Error {
	return gerr.New(500, 100003, "TestErrorReason_Err", fmt.Sprintf(format, args...))
}

// IsUnknown unknown failed
func IsUnknown(err error) bool {
	if err == nil {
		return false
	}
	e := gerr.FromError(err)
	return e.Reason == "TestErrorReason_Unknown" && e.HttpCode == 500
}

// ErrorUnknown unknown failed
func ErrorUnknown(format string, args ...interface{}) *gerr.Error {
	return gerr.New(500, 100004, "TestErrorReason_Unknown", fmt.Sprintf(format, args...))
}

// IsDB internal db failed
func IsDB(err error) bool {
	if err == nil {
		return false
	}
	e := gerr.FromError(err)
	return e.Reason == "TestErrorReason_DB" && e.HttpCode == 500
}

// ErrorDB internal db failed
func ErrorDB(format string, args ...interface{}) *gerr.Error {
	return gerr.New(500, 100005, "TestErrorReason_DB", fmt.Sprintf(format, args...))
}

// IsRedis internal redis failed
func IsRedis(err error) bool {
	if err == nil {
		return false
	}
	e := gerr.FromError(err)
	return e.Reason == "TestErrorReason_Redis" && e.HttpCode == 500
}

// ErrorRedis internal redis failed
func ErrorRedis(format string, args ...interface{}) *gerr.Error {
	return gerr.New(500, 100006, "TestErrorReason_Redis", fmt.Sprintf(format, args...))
}

// IsGRPC internal grpc failed
func IsGRPC(err error) bool {
	if err == nil {
		return false
	}
	e := gerr.FromError(err)
	return e.Reason == "TestErrorReason_GRPC" && e.HttpCode == 500
}

// ErrorGRPC internal grpc failed
func ErrorGRPC(format string, args ...interface{}) *gerr.Error {
	return gerr.New(500, 100007, "TestErrorReason_GRPC", fmt.Sprintf(format, args...))
}

// IsNotFound not found
func IsNotFound(err error) bool {
	if err == nil {
		return false
	}
	e := gerr.FromError(err)
	return e.Reason == "TestErrorReason_NotFound" && e.HttpCode == 400
}

// ErrorNotFound not found
func ErrorNotFound(format string, args ...interface{}) *gerr.Error {
	return gerr.New(400, 100008, "TestErrorReason_NotFound", fmt.Sprintf(format, args...))
}

// IsNoPermission no permission
func IsNoPermission(err error) bool {
	if err == nil {
		return false
	}
	e := gerr.FromError(err)
	return e.Reason == "TestErrorReason_NoPermission" && e.HttpCode == 400
}

// ErrorNoPermission no permission
func ErrorNoPermission(format string, args ...interface{}) *gerr.Error {
	return gerr.New(400, 100009, "TestErrorReason_NoPermission", fmt.Sprintf(format, args...))
}

// IsExisted already existed
func IsExisted(err error) bool {
	if err == nil {
		return false
	}
	e := gerr.FromError(err)
	return e.Reason == "TestErrorReason_Existed" && e.HttpCode == 400
}

// ErrorExisted already existed
func ErrorExisted(format string, args ...interface{}) *gerr.Error {
	return gerr.New(400, 100010, "TestErrorReason_Existed", fmt.Sprintf(format, args...))
}

// IsParamInvalid invalid param
func IsParamInvalid(err error) bool {
	if err == nil {
		return false
	}
	e := gerr.FromError(err)
	return e.Reason == "TestErrorReason_ParamInvalid" && e.HttpCode == 400
}

// ErrorParamInvalid invalid param
func ErrorParamInvalid(format string, args ...interface{}) *gerr.Error {
	return gerr.New(400, 100011, "TestErrorReason_ParamInvalid", fmt.Sprintf(format, args...))
}

// BizErrorCode 获取业务code编码
func BizErrorCode(err error) int {
	if err == nil {
		return 0
	}
	e := gerr.FromError(err)
	return bizErrorCodeMap[e.Reason]
}