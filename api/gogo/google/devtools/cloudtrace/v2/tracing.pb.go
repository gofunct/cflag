// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: google/devtools/cloudtrace/v2/tracing.proto

package google_devtools_cloudtrace_v2

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "go.pedge.io/pb/gogo/google/api"
import _ "github.com/gogo/protobuf/types"
import _ "github.com/gogo/protobuf/types"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// The request message for the `BatchWriteSpans` method.
type BatchWriteSpansRequest struct {
	// Required. The name of the project where the spans belong. The format is
	// `projects/[PROJECT_ID]`.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// A list of new spans. The span names must not match existing
	// spans, or the results are undefined.
	Spans []*Span `protobuf:"bytes,2,rep,name=spans" json:"spans,omitempty"`
}

func (m *BatchWriteSpansRequest) Reset()                    { *m = BatchWriteSpansRequest{} }
func (m *BatchWriteSpansRequest) String() string            { return proto.CompactTextString(m) }
func (*BatchWriteSpansRequest) ProtoMessage()               {}
func (*BatchWriteSpansRequest) Descriptor() ([]byte, []int) { return fileDescriptorTracing, []int{0} }

func (m *BatchWriteSpansRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *BatchWriteSpansRequest) GetSpans() []*Span {
	if m != nil {
		return m.Spans
	}
	return nil
}

func init() {
	proto.RegisterType((*BatchWriteSpansRequest)(nil), "google.devtools.cloudtrace.v2.BatchWriteSpansRequest")
}

func init() { proto.RegisterFile("google/devtools/cloudtrace/v2/tracing.proto", fileDescriptorTracing) }

var fileDescriptorTracing = []byte{
	// 369 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x90, 0xcf, 0x4a, 0xc3, 0x40,
	0x10, 0xc6, 0x49, 0xfc, 0x03, 0xae, 0x82, 0xb0, 0x60, 0x29, 0x51, 0xb1, 0x46, 0xc1, 0xb6, 0xca,
	0x2e, 0x44, 0x3c, 0x58, 0xf0, 0xd2, 0x22, 0x5e, 0x4b, 0x2b, 0xf5, 0xd2, 0xcb, 0x36, 0x1d, 0x63,
	0xa4, 0xc9, 0xc6, 0xec, 0x36, 0x20, 0xe2, 0xc5, 0x8b, 0x0f, 0xa0, 0x4f, 0xe1, 0xc5, 0xf7, 0xf0,
	0xea, 0x2b, 0xf8, 0x20, 0xb2, 0xbb, 0x89, 0x42, 0x29, 0x6d, 0x6f, 0xbb, 0x33, 0xbf, 0x99, 0xf9,
	0xbe, 0x0f, 0x1d, 0x07, 0x9c, 0x07, 0x23, 0xa0, 0x43, 0xc8, 0x24, 0xe7, 0x23, 0x41, 0xfd, 0x11,
	0x1f, 0x0f, 0x65, 0xca, 0x7c, 0xa0, 0x99, 0x47, 0xd5, 0x23, 0x8c, 0x03, 0x92, 0xa4, 0x5c, 0x72,
	0xbc, 0x6b, 0x60, 0x52, 0xc0, 0xe4, 0x1f, 0x26, 0x99, 0xe7, 0xec, 0xe4, 0xbb, 0x58, 0x12, 0x52,
	0x16, 0xc7, 0x5c, 0x32, 0x19, 0xf2, 0x58, 0x98, 0x61, 0xa7, 0x36, 0xff, 0x12, 0xe4, 0xe8, 0x76,
	0x8e, 0xea, 0xdf, 0x60, 0x7c, 0x4b, 0x21, 0x4a, 0xe4, 0x63, 0xde, 0xdc, 0x9b, 0x6c, 0xca, 0x30,
	0x02, 0x21, 0x59, 0x94, 0x18, 0xc0, 0x0d, 0x50, 0xa9, 0xc9, 0xa4, 0x7f, 0x77, 0x93, 0x86, 0x12,
	0xba, 0x09, 0x8b, 0x45, 0x07, 0x1e, 0xc6, 0x20, 0x24, 0xc6, 0x68, 0x39, 0x66, 0x11, 0x94, 0xad,
	0x8a, 0x55, 0x5d, 0xeb, 0xe8, 0x37, 0x3e, 0x47, 0x2b, 0x42, 0x31, 0x65, 0xbb, 0xb2, 0x54, 0x5d,
	0xf7, 0x0e, 0xc8, 0x4c, 0x8f, 0x44, 0xed, 0xeb, 0x98, 0x09, 0xef, 0xd3, 0x46, 0x1b, 0xd7, 0xaa,
	0xd1, 0x85, 0x34, 0x0b, 0x7d, 0xc0, 0xef, 0x16, 0xda, 0x9c, 0x38, 0x8d, 0xcf, 0xe6, 0x2c, 0x9c,
	0x2e, 0xd5, 0x29, 0x15, 0x63, 0x85, 0x4d, 0x72, 0xa9, 0x32, 0x70, 0xbd, 0x97, 0xef, 0x9f, 0x37,
	0xfb, 0xc4, 0x3d, 0x52, 0x99, 0x3d, 0x29, 0x07, 0x17, 0x49, 0xca, 0xef, 0xc1, 0x97, 0x82, 0xd6,
	0x9f, 0x4d, 0x8a, 0xa2, 0x31, 0xf8, 0x5b, 0xda, 0xb0, 0xea, 0xf8, 0xd5, 0x42, 0xa8, 0x95, 0x02,
	0x33, 0x27, 0xf0, 0x22, 0x16, 0x9d, 0x45, 0x20, 0x97, 0x6a, 0x31, 0x35, 0xf7, 0x70, 0x9a, 0x98,
	0x5c, 0x8b, 0x52, 0xa5, 0xe3, 0x6a, 0x58, 0xf5, 0xa6, 0x8f, 0xf6, 0x7d, 0x1e, 0xcd, 0x5e, 0xdd,
	0xd4, 0x99, 0x86, 0x71, 0xd0, 0x56, 0xce, 0xdb, 0xd6, 0x87, 0xbd, 0x75, 0x65, 0xf8, 0x96, 0xc2,
	0x88, 0x4e, 0x9c, 0xf4, 0xbc, 0xaf, 0xa2, 0xde, 0xd7, 0xf5, 0xbe, 0xae, 0xf7, 0x7b, 0xde, 0x60,
	0x55, 0x47, 0x76, 0xfa, 0x1b, 0x00, 0x00, 0xff, 0xff, 0x33, 0x0f, 0xc2, 0xf2, 0xdb, 0x02, 0x00,
	0x00,
}