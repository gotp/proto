// Code generated by protoc-gen-go. DO NOT EDIT.
// source: retcode.proto

package gotp

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type RetcodeType int32

const (
	//------------ common ---------------
	RetcodeType_RetSuccess RetcodeType = 0
	//------------ system 4xxxxx ---------------
	RetcodeType_RetUnknowError RetcodeType = 400001
	RetcodeType_RetBadRequest  RetcodeType = 400400
	RetcodeType_RetNotFound    RetcodeType = 400404
	RetcodeType_RetServerError RetcodeType = 400504
	//------------- access 5xxxxx --------
	RetcodeType_RetServerUnreachable RetcodeType = 500504
)

var RetcodeType_name = map[int32]string{
	0:      "RetSuccess",
	400001: "RetUnknowError",
	400400: "RetBadRequest",
	400404: "RetNotFound",
	400504: "RetServerError",
	500504: "RetServerUnreachable",
}

var RetcodeType_value = map[string]int32{
	"RetSuccess":           0,
	"RetUnknowError":       400001,
	"RetBadRequest":        400400,
	"RetNotFound":          400404,
	"RetServerError":       400504,
	"RetServerUnreachable": 500504,
}

func (x RetcodeType) String() string {
	return proto.EnumName(RetcodeType_name, int32(x))
}

func (RetcodeType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_523b819fd1de7ce3, []int{0}
}

func init() {
	proto.RegisterEnum("gotp.RetcodeType", RetcodeType_name, RetcodeType_value)
}

func init() { proto.RegisterFile("retcode.proto", fileDescriptor_523b819fd1de7ce3) }

var fileDescriptor_523b819fd1de7ce3 = []byte{
	// 172 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2d, 0x4a, 0x2d, 0x49,
	0xce, 0x4f, 0x49, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x49, 0xcf, 0x2f, 0x29, 0xd0,
	0xea, 0x65, 0xe4, 0xe2, 0x0e, 0x82, 0x88, 0x87, 0x54, 0x16, 0xa4, 0x0a, 0xf1, 0x71, 0x71, 0x05,
	0xa5, 0x96, 0x04, 0x97, 0x26, 0x27, 0xa7, 0x16, 0x17, 0x0b, 0x30, 0x08, 0x89, 0x70, 0xf1, 0x05,
	0xa5, 0x96, 0x84, 0xe6, 0x65, 0xe7, 0xe5, 0x97, 0xbb, 0x16, 0x15, 0xe5, 0x17, 0x09, 0x34, 0x6e,
	0x95, 0x10, 0x12, 0xe6, 0xe2, 0x0d, 0x4a, 0x2d, 0x71, 0x4a, 0x4c, 0x09, 0x4a, 0x2d, 0x2c, 0x4d,
	0x2d, 0x2e, 0x11, 0x98, 0xb0, 0x43, 0x42, 0x48, 0x10, 0x6c, 0x92, 0x5f, 0x7e, 0x89, 0x5b, 0x7e,
	0x69, 0x5e, 0x8a, 0xc0, 0x94, 0x1d, 0x12, 0x50, 0xdd, 0xc1, 0xa9, 0x45, 0x65, 0xa9, 0x45, 0x10,
	0xdd, 0x3f, 0x76, 0x48, 0x08, 0x49, 0x71, 0x89, 0xc0, 0x45, 0x43, 0xf3, 0x8a, 0x52, 0x13, 0x93,
	0x33, 0x12, 0x93, 0x72, 0x52, 0x05, 0x66, 0x1c, 0x93, 0x4b, 0x62, 0x03, 0x3b, 0xce, 0x18, 0x10,
	0x00, 0x00, 0xff, 0xff, 0x6c, 0x81, 0x51, 0xbd, 0xad, 0x00, 0x00, 0x00,
}
