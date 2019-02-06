// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/bigtable/admin/v2/instance.proto

package google_bigtable_admin_v2

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "go.pedge.io/pb/go/google/api"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Possible states of an instance.
type Instance_State int32

const (
	// The state of the instance could not be determined.
	Instance_STATE_NOT_KNOWN Instance_State = 0
	// The instance has been successfully created and can serve requests
	// to its tables.
	Instance_READY Instance_State = 1
	// The instance is currently being created, and may be destroyed
	// if the creation process encounters an error.
	Instance_CREATING Instance_State = 2
)

var Instance_State_name = map[int32]string{
	0: "STATE_NOT_KNOWN",
	1: "READY",
	2: "CREATING",
}
var Instance_State_value = map[string]int32{
	"STATE_NOT_KNOWN": 0,
	"READY":           1,
	"CREATING":        2,
}

func (x Instance_State) String() string {
	return proto.EnumName(Instance_State_name, int32(x))
}
func (Instance_State) EnumDescriptor() ([]byte, []int) { return fileDescriptor3, []int{0, 0} }

// The type of the instance.
type Instance_Type int32

const (
	// The type of the instance is unspecified. If set when creating an
	// instance, a `PRODUCTION` instance will be created. If set when updating
	// an instance, the type will be left unchanged.
	Instance_TYPE_UNSPECIFIED Instance_Type = 0
	// An instance meant for production use. `serve_nodes` must be set
	// on the cluster.
	Instance_PRODUCTION Instance_Type = 1
	// The instance is meant for development and testing purposes only; it has
	// no performance or uptime guarantees and is not covered by SLA.
	// After a development instance is created, it can be upgraded by
	// updating the instance to type `PRODUCTION`. An instance created
	// as a production instance cannot be changed to a development instance.
	// When creating a development instance, `serve_nodes` on the cluster must
	// not be set.
	Instance_DEVELOPMENT Instance_Type = 2
)

var Instance_Type_name = map[int32]string{
	0: "TYPE_UNSPECIFIED",
	1: "PRODUCTION",
	2: "DEVELOPMENT",
}
var Instance_Type_value = map[string]int32{
	"TYPE_UNSPECIFIED": 0,
	"PRODUCTION":       1,
	"DEVELOPMENT":      2,
}

func (x Instance_Type) String() string {
	return proto.EnumName(Instance_Type_name, int32(x))
}
func (Instance_Type) EnumDescriptor() ([]byte, []int) { return fileDescriptor3, []int{0, 1} }

// Possible states of a cluster.
type Cluster_State int32

const (
	// The state of the cluster could not be determined.
	Cluster_STATE_NOT_KNOWN Cluster_State = 0
	// The cluster has been successfully created and is ready to serve requests.
	Cluster_READY Cluster_State = 1
	// The cluster is currently being created, and may be destroyed
	// if the creation process encounters an error.
	// A cluster may not be able to serve requests while being created.
	Cluster_CREATING Cluster_State = 2
	// The cluster is currently being resized, and may revert to its previous
	// node count if the process encounters an error.
	// A cluster is still capable of serving requests while being resized,
	// but may exhibit performance as if its number of allocated nodes is
	// between the starting and requested states.
	Cluster_RESIZING Cluster_State = 3
	// The cluster has no backing nodes. The data (tables) still
	// exist, but no operations can be performed on the cluster.
	Cluster_DISABLED Cluster_State = 4
)

var Cluster_State_name = map[int32]string{
	0: "STATE_NOT_KNOWN",
	1: "READY",
	2: "CREATING",
	3: "RESIZING",
	4: "DISABLED",
}
var Cluster_State_value = map[string]int32{
	"STATE_NOT_KNOWN": 0,
	"READY":           1,
	"CREATING":        2,
	"RESIZING":        3,
	"DISABLED":        4,
}

func (x Cluster_State) String() string {
	return proto.EnumName(Cluster_State_name, int32(x))
}
func (Cluster_State) EnumDescriptor() ([]byte, []int) { return fileDescriptor3, []int{1, 0} }

// A collection of Bigtable [Tables][google.bigtable.admin.v2.Table] and
// the resources that serve them.
// All tables in an instance are served from a single
// [Cluster][google.bigtable.admin.v2.Cluster].
type Instance struct {
	// (`OutputOnly`)
	// The unique name of the instance. Values are of the form
	// `projects/<project>/instances/[a-z][a-z0-9\\-]+[a-z0-9]`.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// The descriptive name for this instance as it appears in UIs.
	// Can be changed at any time, but should be kept globally unique
	// to avoid confusion.
	DisplayName string `protobuf:"bytes,2,opt,name=display_name,json=displayName" json:"display_name,omitempty"`
	// (`OutputOnly`)
	// The current state of the instance.
	State Instance_State `protobuf:"varint,3,opt,name=state,enum=google.bigtable.admin.v2.Instance_State" json:"state,omitempty"`
	// The type of the instance. Defaults to `PRODUCTION`.
	Type Instance_Type `protobuf:"varint,4,opt,name=type,enum=google.bigtable.admin.v2.Instance_Type" json:"type,omitempty"`
}

func (m *Instance) Reset()                    { *m = Instance{} }
func (m *Instance) String() string            { return proto.CompactTextString(m) }
func (*Instance) ProtoMessage()               {}
func (*Instance) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

func (m *Instance) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Instance) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

func (m *Instance) GetState() Instance_State {
	if m != nil {
		return m.State
	}
	return Instance_STATE_NOT_KNOWN
}

func (m *Instance) GetType() Instance_Type {
	if m != nil {
		return m.Type
	}
	return Instance_TYPE_UNSPECIFIED
}

// A resizable group of nodes in a particular cloud location, capable
// of serving all [Tables][google.bigtable.admin.v2.Table] in the parent
// [Instance][google.bigtable.admin.v2.Instance].
type Cluster struct {
	// (`OutputOnly`)
	// The unique name of the cluster. Values are of the form
	// `projects/<project>/instances/<instance>/clusters/[a-z][-a-z0-9]*`.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// (`CreationOnly`)
	// The location where this cluster's nodes and storage reside. For best
	// performance, clients should be located as close as possible to this cluster.
	// Currently only zones are supported, so values should be of the form
	// `projects/<project>/locations/<zone>`.
	Location string `protobuf:"bytes,2,opt,name=location" json:"location,omitempty"`
	// (`OutputOnly`)
	// The current state of the cluster.
	State Cluster_State `protobuf:"varint,3,opt,name=state,enum=google.bigtable.admin.v2.Cluster_State" json:"state,omitempty"`
	// The number of nodes allocated to this cluster. More nodes enable higher
	// throughput and more consistent performance.
	ServeNodes int32 `protobuf:"varint,4,opt,name=serve_nodes,json=serveNodes" json:"serve_nodes,omitempty"`
	// (`CreationOnly`)
	// The type of storage used by this cluster to serve its
	// parent instance's tables, unless explicitly overridden.
	DefaultStorageType StorageType `protobuf:"varint,5,opt,name=default_storage_type,json=defaultStorageType,enum=google.bigtable.admin.v2.StorageType" json:"default_storage_type,omitempty"`
}

func (m *Cluster) Reset()                    { *m = Cluster{} }
func (m *Cluster) String() string            { return proto.CompactTextString(m) }
func (*Cluster) ProtoMessage()               {}
func (*Cluster) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{1} }

func (m *Cluster) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Cluster) GetLocation() string {
	if m != nil {
		return m.Location
	}
	return ""
}

func (m *Cluster) GetState() Cluster_State {
	if m != nil {
		return m.State
	}
	return Cluster_STATE_NOT_KNOWN
}

func (m *Cluster) GetServeNodes() int32 {
	if m != nil {
		return m.ServeNodes
	}
	return 0
}

func (m *Cluster) GetDefaultStorageType() StorageType {
	if m != nil {
		return m.DefaultStorageType
	}
	return StorageType_STORAGE_TYPE_UNSPECIFIED
}

func init() {
	proto.RegisterType((*Instance)(nil), "google.bigtable.admin.v2.Instance")
	proto.RegisterType((*Cluster)(nil), "google.bigtable.admin.v2.Cluster")
	proto.RegisterEnum("google.bigtable.admin.v2.Instance_State", Instance_State_name, Instance_State_value)
	proto.RegisterEnum("google.bigtable.admin.v2.Instance_Type", Instance_Type_name, Instance_Type_value)
	proto.RegisterEnum("google.bigtable.admin.v2.Cluster_State", Cluster_State_name, Cluster_State_value)
}

func init() { proto.RegisterFile("google/bigtable/admin/v2/instance.proto", fileDescriptor3) }

var fileDescriptor3 = []byte{
	// 486 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0x51, 0x6b, 0x9b, 0x5e,
	0x18, 0xc6, 0xab, 0x4d, 0xfe, 0xff, 0xf4, 0x4d, 0xd7, 0xca, 0x59, 0x2f, 0x42, 0x28, 0x5b, 0x27,
	0x94, 0xe6, 0x4a, 0xc1, 0xb1, 0xab, 0xd1, 0x81, 0xd1, 0xb3, 0x22, 0xeb, 0xd4, 0xa9, 0x4d, 0xe9,
	0x08, 0xc8, 0x49, 0x3c, 0x0b, 0x82, 0x7a, 0x5c, 0x3c, 0x09, 0xe4, 0x2b, 0xed, 0x7a, 0x5f, 0x61,
	0x37, 0xfb, 0x54, 0xc3, 0xa3, 0x8e, 0x15, 0x96, 0x31, 0x76, 0xe7, 0xfb, 0x9e, 0xdf, 0xf3, 0x3e,
	0xc7, 0xe7, 0xbc, 0x70, 0xb5, 0x62, 0x6c, 0x95, 0x51, 0x7d, 0x91, 0xae, 0x38, 0x59, 0x64, 0x54,
	0x27, 0x49, 0x9e, 0x16, 0xfa, 0xd6, 0xd0, 0xd3, 0xa2, 0xe2, 0xa4, 0x58, 0x52, 0xad, 0x5c, 0x33,
	0xce, 0xd0, 0xa8, 0x01, 0xb5, 0x0e, 0xd4, 0x04, 0xa8, 0x6d, 0x8d, 0xf1, 0x79, 0x3b, 0x82, 0x94,
	0xa9, 0x4e, 0x8a, 0x82, 0x71, 0xc2, 0x53, 0x56, 0x54, 0x8d, 0x6e, 0x7c, 0xb9, 0xd7, 0x60, 0xc9,
	0xf2, 0x9c, 0x15, 0x0d, 0xa6, 0x7e, 0x95, 0x61, 0xe0, 0xb4, 0x8e, 0x08, 0x41, 0xaf, 0x20, 0x39,
	0x1d, 0x49, 0x17, 0xd2, 0xe4, 0x28, 0x10, 0xdf, 0xe8, 0x05, 0x1c, 0x27, 0x69, 0x55, 0x66, 0x64,
	0x17, 0x8b, 0x33, 0x59, 0x9c, 0x0d, 0xdb, 0x9e, 0x5b, 0x23, 0x6f, 0xa0, 0x5f, 0x71, 0xc2, 0xe9,
	0xe8, 0xf0, 0x42, 0x9a, 0x9c, 0x18, 0x13, 0x6d, 0xdf, 0x95, 0xb5, 0xce, 0x49, 0x0b, 0x6b, 0x3e,
	0x68, 0x64, 0xe8, 0x35, 0xf4, 0xf8, 0xae, 0xa4, 0xa3, 0x9e, 0x90, 0x5f, 0xfd, 0x85, 0x3c, 0xda,
	0x95, 0x34, 0x10, 0x22, 0xf5, 0x15, 0xf4, 0xc5, 0x30, 0xf4, 0x14, 0x4e, 0xc3, 0xc8, 0x8c, 0x70,
	0xec, 0x7a, 0x51, 0xfc, 0xce, 0xf5, 0xee, 0x5d, 0xe5, 0x00, 0x1d, 0x41, 0x3f, 0xc0, 0xa6, 0xfd,
	0xa0, 0x48, 0xe8, 0x18, 0x06, 0x56, 0x80, 0xcd, 0xc8, 0x71, 0x6f, 0x14, 0x59, 0xbd, 0x86, 0x5e,
	0x3d, 0x04, 0x9d, 0x81, 0x12, 0x3d, 0xf8, 0x38, 0xbe, 0x73, 0x43, 0x1f, 0x5b, 0xce, 0x5b, 0x07,
	0xdb, 0xca, 0x01, 0x3a, 0x01, 0xf0, 0x03, 0xcf, 0xbe, 0xb3, 0x22, 0xc7, 0x73, 0x15, 0x09, 0x9d,
	0xc2, 0xd0, 0xc6, 0x33, 0x7c, 0xeb, 0xf9, 0xef, 0xb1, 0x1b, 0x29, 0xb2, 0xfa, 0x4d, 0x86, 0xff,
	0xad, 0x6c, 0x53, 0x71, 0xba, 0xfe, 0x6d, 0x6a, 0x63, 0x18, 0x64, 0x6c, 0x29, 0x1e, 0xa4, 0x4d,
	0xec, 0x67, 0x8d, 0xae, 0x1f, 0xc7, 0xf5, 0x87, 0xff, 0x6d, 0x1d, 0x1e, 0xa7, 0xf5, 0x1c, 0x86,
	0x15, 0x5d, 0x6f, 0x69, 0x5c, 0xb0, 0x84, 0x56, 0x22, 0xb4, 0x7e, 0x00, 0xa2, 0xe5, 0xd6, 0x1d,
	0x74, 0x0f, 0x67, 0x09, 0xfd, 0x44, 0x36, 0x19, 0x8f, 0x2b, 0xce, 0xd6, 0x64, 0x45, 0x63, 0x11,
	0x6f, 0x5f, 0xd8, 0x5d, 0xee, 0xb7, 0x0b, 0x1b, 0x5a, 0x84, 0x8b, 0xda, 0x11, 0xbf, 0xf4, 0xd4,
	0x0f, 0xff, 0x14, 0x75, 0x5d, 0x05, 0x38, 0x74, 0x3e, 0xd6, 0xd5, 0x61, 0x5d, 0xd9, 0x4e, 0x68,
	0x4e, 0x6f, 0xb1, 0xad, 0xf4, 0xa6, 0x9f, 0xe1, 0x7c, 0xc9, 0xf2, 0xbd, 0x57, 0x9a, 0x3e, 0xe9,
	0x9e, 0xdc, 0xaf, 0xb7, 0xd5, 0x97, 0xbe, 0xc8, 0xcf, 0x6e, 0x1a, 0xd6, 0xca, 0xd8, 0x26, 0xd1,
	0xa6, 0x9d, 0xc2, 0x14, 0x8a, 0x99, 0xf1, 0xbd, 0x03, 0xe6, 0x02, 0x98, 0x77, 0xc0, 0x5c, 0x00,
	0xf3, 0x99, 0xb1, 0xf8, 0x4f, 0x2c, 0xfe, 0xcb, 0x1f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x31, 0x30,
	0x1b, 0x63, 0x82, 0x03, 0x00, 0x00,
}