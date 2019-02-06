// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: google/cloud/vision/v1p1beta1/geometry.proto

package google_cloud_vision_v1p1beta1

import proto "github.com/gogo/protobuf/proto"
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
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// A vertex represents a 2D point in the image.
// NOTE: the vertex coordinates are in the same scale as the original image.
type Vertex struct {
	// X coordinate.
	X int32 `protobuf:"varint,1,opt,name=x,proto3" json:"x,omitempty"`
	// Y coordinate.
	Y int32 `protobuf:"varint,2,opt,name=y,proto3" json:"y,omitempty"`
}

func (m *Vertex) Reset()                    { *m = Vertex{} }
func (m *Vertex) String() string            { return proto.CompactTextString(m) }
func (*Vertex) ProtoMessage()               {}
func (*Vertex) Descriptor() ([]byte, []int) { return fileDescriptorGeometry, []int{0} }

func (m *Vertex) GetX() int32 {
	if m != nil {
		return m.X
	}
	return 0
}

func (m *Vertex) GetY() int32 {
	if m != nil {
		return m.Y
	}
	return 0
}

// A bounding polygon for the detected image annotation.
type BoundingPoly struct {
	// The bounding polygon vertices.
	Vertices []*Vertex `protobuf:"bytes,1,rep,name=vertices" json:"vertices,omitempty"`
}

func (m *BoundingPoly) Reset()                    { *m = BoundingPoly{} }
func (m *BoundingPoly) String() string            { return proto.CompactTextString(m) }
func (*BoundingPoly) ProtoMessage()               {}
func (*BoundingPoly) Descriptor() ([]byte, []int) { return fileDescriptorGeometry, []int{1} }

func (m *BoundingPoly) GetVertices() []*Vertex {
	if m != nil {
		return m.Vertices
	}
	return nil
}

// A 3D position in the image, used primarily for Face detection landmarks.
// A valid Position must have both x and y coordinates.
// The position coordinates are in the same scale as the original image.
type Position struct {
	// X coordinate.
	X float32 `protobuf:"fixed32,1,opt,name=x,proto3" json:"x,omitempty"`
	// Y coordinate.
	Y float32 `protobuf:"fixed32,2,opt,name=y,proto3" json:"y,omitempty"`
	// Z coordinate (or depth).
	Z float32 `protobuf:"fixed32,3,opt,name=z,proto3" json:"z,omitempty"`
}

func (m *Position) Reset()                    { *m = Position{} }
func (m *Position) String() string            { return proto.CompactTextString(m) }
func (*Position) ProtoMessage()               {}
func (*Position) Descriptor() ([]byte, []int) { return fileDescriptorGeometry, []int{2} }

func (m *Position) GetX() float32 {
	if m != nil {
		return m.X
	}
	return 0
}

func (m *Position) GetY() float32 {
	if m != nil {
		return m.Y
	}
	return 0
}

func (m *Position) GetZ() float32 {
	if m != nil {
		return m.Z
	}
	return 0
}

func init() {
	proto.RegisterType((*Vertex)(nil), "google.cloud.vision.v1p1beta1.Vertex")
	proto.RegisterType((*BoundingPoly)(nil), "google.cloud.vision.v1p1beta1.BoundingPoly")
	proto.RegisterType((*Position)(nil), "google.cloud.vision.v1p1beta1.Position")
}

func init() {
	proto.RegisterFile("google/cloud/vision/v1p1beta1/geometry.proto", fileDescriptorGeometry)
}

var fileDescriptorGeometry = []byte{
	// 215 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0xb1, 0x4b, 0xc4, 0x30,
	0x14, 0x87, 0x79, 0x3d, 0x3c, 0x8e, 0x78, 0x2e, 0x9d, 0xb2, 0x08, 0x67, 0x51, 0xb8, 0x41, 0x12,
	0xaa, 0x82, 0xb3, 0x5d, 0x5c, 0x63, 0x07, 0x77, 0xaf, 0xf7, 0x08, 0x81, 0x5e, 0x5e, 0x49, 0xd3,
	0xd2, 0xf4, 0x2f, 0x77, 0x94, 0x36, 0x52, 0x71, 0xe9, 0xf8, 0xc1, 0xf7, 0xe3, 0x83, 0x1f, 0x7b,
	0xd4, 0x44, 0xba, 0x46, 0x59, 0xd5, 0xd4, 0x9d, 0x65, 0x6f, 0x5a, 0x43, 0x56, 0xf6, 0x79, 0x93,
	0x9f, 0xd0, 0x7f, 0xe5, 0x52, 0x23, 0x5d, 0xd0, 0xbb, 0x20, 0x1a, 0x47, 0x9e, 0xd2, 0xdb, 0x68,
	0x8b, 0xd9, 0x16, 0xd1, 0x16, 0x8b, 0x9d, 0xdd, 0xb3, 0xed, 0x27, 0x3a, 0x8f, 0x43, 0xba, 0x67,
	0x30, 0x70, 0x38, 0xc0, 0xf1, 0xaa, 0x84, 0x99, 0x02, 0x4f, 0x22, 0x85, 0xec, 0x83, 0xed, 0x0b,
	0xea, 0xec, 0xd9, 0x58, 0xad, 0xa8, 0x0e, 0xe9, 0x1b, 0xdb, 0xf5, 0xe8, 0xbc, 0xa9, 0xb0, 0xe5,
	0x70, 0xd8, 0x1c, 0xaf, 0x9f, 0x1e, 0xc4, 0x6a, 0x47, 0xc4, 0x48, 0xb9, 0xcc, 0xb2, 0x17, 0xb6,
	0x53, 0xd4, 0x1a, 0x6f, 0xc8, 0xfe, 0xa5, 0x93, 0x7f, 0xe9, 0xa4, 0x84, 0x30, 0xd1, 0xc8, 0x37,
	0x91, 0xc6, 0xe2, 0x95, 0xdd, 0x55, 0x74, 0x59, 0x6f, 0x15, 0x37, 0xef, 0xbf, 0x17, 0xa8, 0xe9,
	0x01, 0x05, 0xdf, 0x00, 0xa7, 0xed, 0xfc, 0xc6, 0xf3, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xc2,
	0x41, 0x81, 0x2c, 0x3d, 0x01, 0x00, 0x00,
}