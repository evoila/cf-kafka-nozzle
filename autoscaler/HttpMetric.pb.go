// Code generated by protoc-gen-go.
// source: HttpMetric.proto
// DO NOT EDIT!

/*
Package autoscaler is a generated protocol buffer package.

It is generated from these files:
	HttpMetric.proto

It has these top-level messages:
	ProtoHttpMetric
*/
package autoscaler

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ProtoHttpMetric struct {
	Timestamp   int64  `protobuf:"varint,1,opt,name=timestamp" json:"timestamp,omitempty"`
	MetricName  string `protobuf:"bytes,2,opt,name=metricName" json:"metricName,omitempty"`
	AppId       string `protobuf:"bytes,3,opt,name=appId" json:"appId,omitempty"`
	Requests    int32  `protobuf:"varint,4,opt,name=requests" json:"requests,omitempty"`
	Latency     int32  `protobuf:"varint,5,opt,name=latency" json:"latency,omitempty"`
	Description string `protobuf:"bytes,6,opt,name=description" json:"description,omitempty"`
}

func (m *ProtoHttpMetric) Reset()                    { *m = ProtoHttpMetric{} }
func (m *ProtoHttpMetric) String() string            { return proto.CompactTextString(m) }
func (*ProtoHttpMetric) ProtoMessage()               {}
func (*ProtoHttpMetric) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ProtoHttpMetric) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *ProtoHttpMetric) GetMetricName() string {
	if m != nil {
		return m.MetricName
	}
	return ""
}

func (m *ProtoHttpMetric) GetAppId() string {
	if m != nil {
		return m.AppId
	}
	return ""
}

func (m *ProtoHttpMetric) GetRequests() int32 {
	if m != nil {
		return m.Requests
	}
	return 0
}

func (m *ProtoHttpMetric) GetLatency() int32 {
	if m != nil {
		return m.Latency
	}
	return 0
}

func (m *ProtoHttpMetric) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func init() {
	proto.RegisterType((*ProtoHttpMetric)(nil), "autoscaler.ProtoHttpMetric")
}

func init() { proto.RegisterFile("HttpMetric.proto", fileDescriptor0) }

var fileDescriptor1 = []byte{
	// 222 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x8f, 0x31, 0x4b, 0x04, 0x31,
	0x10, 0x85, 0xc9, 0x9d, 0x7b, 0x7a, 0x63, 0xa1, 0x04, 0x8b, 0x9c, 0x88, 0x04, 0xab, 0xc5, 0x62,
	0x1b, 0xff, 0xc1, 0x55, 0x5a, 0x28, 0xc7, 0x36, 0xd6, 0xb9, 0xec, 0x1c, 0x04, 0x36, 0x9b, 0x31,
	0x99, 0x2d, 0xfc, 0x4b, 0xfe, 0x07, 0xff, 0x9b, 0x98, 0x55, 0xb3, 0x20, 0xd7, 0xcd, 0xfb, 0xde,
	0x9b, 0x07, 0x0f, 0x2e, 0x1f, 0x99, 0xe9, 0x19, 0x39, 0x3a, 0xdb, 0x50, 0x0c, 0x1c, 0x24, 0x98,
	0x91, 0x43, 0xb2, 0xa6, 0xc7, 0x78, 0xf7, 0x29, 0xe0, 0x62, 0xf7, 0x4d, 0x4b, 0x4a, 0xde, 0xc0,
	0x9a, 0x9d, 0xc7, 0xc4, 0xc6, 0x93, 0x12, 0x5a, 0xd4, 0xcb, 0xb6, 0x00, 0x79, 0x0b, 0xe0, 0x73,
	0xee, 0xc5, 0x78, 0x54, 0x0b, 0x2d, 0xea, 0x75, 0x3b, 0x23, 0xf2, 0x0a, 0x2a, 0x43, 0xf4, 0xd4,
	0xa9, 0x65, 0xb6, 0x26, 0x21, 0xaf, 0xe1, 0x2c, 0xe2, 0xdb, 0x88, 0x89, 0x93, 0x3a, 0xd1, 0xa2,
	0xae, 0xda, 0x3f, 0x2d, 0x15, 0x9c, 0xf6, 0x86, 0x71, 0xb0, 0xef, 0xaa, 0xca, 0xd6, 0xaf, 0x94,
	0x1a, 0xce, 0x3b, 0x4c, 0x36, 0x3a, 0x62, 0x17, 0x06, 0xb5, 0xca, 0x8d, 0x73, 0xb4, 0xb5, 0xb0,
	0xe9, 0xb0, 0xb1, 0x87, 0xa6, 0x6c, 0x9a, 0x56, 0xee, 0xc7, 0xc3, 0x76, 0xb3, 0xfb, 0xb9, 0xca,
	0xb8, 0xd7, 0x68, 0x88, 0x30, 0x7e, 0x2c, 0xee, 0xff, 0xbd, 0x1d, 0x0d, 0xef, 0x57, 0xb9, 0xf1,
	0xe1, 0x2b, 0x00, 0x00, 0xff, 0xff, 0xec, 0x22, 0xf6, 0x8a, 0x4b, 0x01, 0x00, 0x00,
}
