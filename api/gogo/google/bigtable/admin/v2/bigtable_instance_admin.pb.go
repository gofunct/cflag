// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: google/bigtable/admin/v2/bigtable_instance_admin.proto

package google_bigtable_admin_v2

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "go.pedge.io/pb/gogo/google/api"
import _ "go.pedge.io/pb/gogo/google/longrunning"
import _ "github.com/gogo/protobuf/types"
import google_protobuf1 "github.com/gogo/protobuf/types"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// Request message for BigtableInstanceAdmin.CreateInstance.
type CreateInstanceRequest struct {
	// The unique name of the project in which to create the new instance.
	// Values are of the form `projects/<project>`.
	Parent string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	// The ID to be used when referring to the new instance within its project,
	// e.g., just `myinstance` rather than
	// `projects/myproject/instances/myinstance`.
	InstanceId string `protobuf:"bytes,2,opt,name=instance_id,json=instanceId,proto3" json:"instance_id,omitempty"`
	// The instance to create.
	// Fields marked `OutputOnly` must be left blank.
	Instance *Instance `protobuf:"bytes,3,opt,name=instance" json:"instance,omitempty"`
	// The clusters to be created within the instance, mapped by desired
	// cluster ID, e.g., just `mycluster` rather than
	// `projects/myproject/instances/myinstance/clusters/mycluster`.
	// Fields marked `OutputOnly` must be left blank.
	// Currently exactly one cluster must be specified.
	Clusters map[string]*Cluster `protobuf:"bytes,4,rep,name=clusters" json:"clusters,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *CreateInstanceRequest) Reset()         { *m = CreateInstanceRequest{} }
func (m *CreateInstanceRequest) String() string { return proto.CompactTextString(m) }
func (*CreateInstanceRequest) ProtoMessage()    {}
func (*CreateInstanceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptorBigtableInstanceAdmin, []int{0}
}

func (m *CreateInstanceRequest) GetParent() string {
	if m != nil {
		return m.Parent
	}
	return ""
}

func (m *CreateInstanceRequest) GetInstanceId() string {
	if m != nil {
		return m.InstanceId
	}
	return ""
}

func (m *CreateInstanceRequest) GetInstance() *Instance {
	if m != nil {
		return m.Instance
	}
	return nil
}

func (m *CreateInstanceRequest) GetClusters() map[string]*Cluster {
	if m != nil {
		return m.Clusters
	}
	return nil
}

// Request message for BigtableInstanceAdmin.GetInstance.
type GetInstanceRequest struct {
	// The unique name of the requested instance. Values are of the form
	// `projects/<project>/instances/<instance>`.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (m *GetInstanceRequest) Reset()         { *m = GetInstanceRequest{} }
func (m *GetInstanceRequest) String() string { return proto.CompactTextString(m) }
func (*GetInstanceRequest) ProtoMessage()    {}
func (*GetInstanceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptorBigtableInstanceAdmin, []int{1}
}

func (m *GetInstanceRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// Request message for BigtableInstanceAdmin.ListInstances.
type ListInstancesRequest struct {
	// The unique name of the project for which a list of instances is requested.
	// Values are of the form `projects/<project>`.
	Parent string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	// The value of `next_page_token` returned by a previous call.
	PageToken string `protobuf:"bytes,2,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
}

func (m *ListInstancesRequest) Reset()         { *m = ListInstancesRequest{} }
func (m *ListInstancesRequest) String() string { return proto.CompactTextString(m) }
func (*ListInstancesRequest) ProtoMessage()    {}
func (*ListInstancesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptorBigtableInstanceAdmin, []int{2}
}

func (m *ListInstancesRequest) GetParent() string {
	if m != nil {
		return m.Parent
	}
	return ""
}

func (m *ListInstancesRequest) GetPageToken() string {
	if m != nil {
		return m.PageToken
	}
	return ""
}

// Response message for BigtableInstanceAdmin.ListInstances.
type ListInstancesResponse struct {
	// The list of requested instances.
	Instances []*Instance `protobuf:"bytes,1,rep,name=instances" json:"instances,omitempty"`
	// Locations from which Instance information could not be retrieved,
	// due to an outage or some other transient condition.
	// Instances whose Clusters are all in one of the failed locations
	// may be missing from `instances`, and Instances with at least one
	// Cluster in a failed location may only have partial information returned.
	FailedLocations []string `protobuf:"bytes,2,rep,name=failed_locations,json=failedLocations" json:"failed_locations,omitempty"`
	// Set if not all instances could be returned in a single response.
	// Pass this value to `page_token` in another request to get the next
	// page of results.
	NextPageToken string `protobuf:"bytes,3,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
}

func (m *ListInstancesResponse) Reset()         { *m = ListInstancesResponse{} }
func (m *ListInstancesResponse) String() string { return proto.CompactTextString(m) }
func (*ListInstancesResponse) ProtoMessage()    {}
func (*ListInstancesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptorBigtableInstanceAdmin, []int{3}
}

func (m *ListInstancesResponse) GetInstances() []*Instance {
	if m != nil {
		return m.Instances
	}
	return nil
}

func (m *ListInstancesResponse) GetFailedLocations() []string {
	if m != nil {
		return m.FailedLocations
	}
	return nil
}

func (m *ListInstancesResponse) GetNextPageToken() string {
	if m != nil {
		return m.NextPageToken
	}
	return ""
}

// Request message for BigtableInstanceAdmin.DeleteInstance.
type DeleteInstanceRequest struct {
	// The unique name of the instance to be deleted.
	// Values are of the form `projects/<project>/instances/<instance>`.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (m *DeleteInstanceRequest) Reset()         { *m = DeleteInstanceRequest{} }
func (m *DeleteInstanceRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteInstanceRequest) ProtoMessage()    {}
func (*DeleteInstanceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptorBigtableInstanceAdmin, []int{4}
}

func (m *DeleteInstanceRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// Request message for BigtableInstanceAdmin.CreateCluster.
type CreateClusterRequest struct {
	// The unique name of the instance in which to create the new cluster.
	// Values are of the form
	// `projects/<project>/instances/<instance>`.
	Parent string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	// The ID to be used when referring to the new cluster within its instance,
	// e.g., just `mycluster` rather than
	// `projects/myproject/instances/myinstance/clusters/mycluster`.
	ClusterId string `protobuf:"bytes,2,opt,name=cluster_id,json=clusterId,proto3" json:"cluster_id,omitempty"`
	// The cluster to be created.
	// Fields marked `OutputOnly` must be left blank.
	Cluster *Cluster `protobuf:"bytes,3,opt,name=cluster" json:"cluster,omitempty"`
}

func (m *CreateClusterRequest) Reset()         { *m = CreateClusterRequest{} }
func (m *CreateClusterRequest) String() string { return proto.CompactTextString(m) }
func (*CreateClusterRequest) ProtoMessage()    {}
func (*CreateClusterRequest) Descriptor() ([]byte, []int) {
	return fileDescriptorBigtableInstanceAdmin, []int{5}
}

func (m *CreateClusterRequest) GetParent() string {
	if m != nil {
		return m.Parent
	}
	return ""
}

func (m *CreateClusterRequest) GetClusterId() string {
	if m != nil {
		return m.ClusterId
	}
	return ""
}

func (m *CreateClusterRequest) GetCluster() *Cluster {
	if m != nil {
		return m.Cluster
	}
	return nil
}

// Request message for BigtableInstanceAdmin.GetCluster.
type GetClusterRequest struct {
	// The unique name of the requested cluster. Values are of the form
	// `projects/<project>/instances/<instance>/clusters/<cluster>`.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (m *GetClusterRequest) Reset()         { *m = GetClusterRequest{} }
func (m *GetClusterRequest) String() string { return proto.CompactTextString(m) }
func (*GetClusterRequest) ProtoMessage()    {}
func (*GetClusterRequest) Descriptor() ([]byte, []int) {
	return fileDescriptorBigtableInstanceAdmin, []int{6}
}

func (m *GetClusterRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// Request message for BigtableInstanceAdmin.ListClusters.
type ListClustersRequest struct {
	// The unique name of the instance for which a list of clusters is requested.
	// Values are of the form `projects/<project>/instances/<instance>`.
	// Use `<instance> = '-'` to list Clusters for all Instances in a project,
	// e.g., `projects/myproject/instances/-`.
	Parent string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	// The value of `next_page_token` returned by a previous call.
	PageToken string `protobuf:"bytes,2,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
}

func (m *ListClustersRequest) Reset()         { *m = ListClustersRequest{} }
func (m *ListClustersRequest) String() string { return proto.CompactTextString(m) }
func (*ListClustersRequest) ProtoMessage()    {}
func (*ListClustersRequest) Descriptor() ([]byte, []int) {
	return fileDescriptorBigtableInstanceAdmin, []int{7}
}

func (m *ListClustersRequest) GetParent() string {
	if m != nil {
		return m.Parent
	}
	return ""
}

func (m *ListClustersRequest) GetPageToken() string {
	if m != nil {
		return m.PageToken
	}
	return ""
}

// Response message for BigtableInstanceAdmin.ListClusters.
type ListClustersResponse struct {
	// The list of requested clusters.
	Clusters []*Cluster `protobuf:"bytes,1,rep,name=clusters" json:"clusters,omitempty"`
	// Locations from which Cluster information could not be retrieved,
	// due to an outage or some other transient condition.
	// Clusters from these locations may be missing from `clusters`,
	// or may only have partial information returned.
	FailedLocations []string `protobuf:"bytes,2,rep,name=failed_locations,json=failedLocations" json:"failed_locations,omitempty"`
	// Set if not all clusters could be returned in a single response.
	// Pass this value to `page_token` in another request to get the next
	// page of results.
	NextPageToken string `protobuf:"bytes,3,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
}

func (m *ListClustersResponse) Reset()         { *m = ListClustersResponse{} }
func (m *ListClustersResponse) String() string { return proto.CompactTextString(m) }
func (*ListClustersResponse) ProtoMessage()    {}
func (*ListClustersResponse) Descriptor() ([]byte, []int) {
	return fileDescriptorBigtableInstanceAdmin, []int{8}
}

func (m *ListClustersResponse) GetClusters() []*Cluster {
	if m != nil {
		return m.Clusters
	}
	return nil
}

func (m *ListClustersResponse) GetFailedLocations() []string {
	if m != nil {
		return m.FailedLocations
	}
	return nil
}

func (m *ListClustersResponse) GetNextPageToken() string {
	if m != nil {
		return m.NextPageToken
	}
	return ""
}

// Request message for BigtableInstanceAdmin.DeleteCluster.
type DeleteClusterRequest struct {
	// The unique name of the cluster to be deleted. Values are of the form
	// `projects/<project>/instances/<instance>/clusters/<cluster>`.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (m *DeleteClusterRequest) Reset()         { *m = DeleteClusterRequest{} }
func (m *DeleteClusterRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteClusterRequest) ProtoMessage()    {}
func (*DeleteClusterRequest) Descriptor() ([]byte, []int) {
	return fileDescriptorBigtableInstanceAdmin, []int{9}
}

func (m *DeleteClusterRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// The metadata for the Operation returned by CreateInstance.
type CreateInstanceMetadata struct {
	// The request that prompted the initiation of this CreateInstance operation.
	OriginalRequest *CreateInstanceRequest `protobuf:"bytes,1,opt,name=original_request,json=originalRequest" json:"original_request,omitempty"`
	// The time at which the original request was received.
	RequestTime *google_protobuf1.Timestamp `protobuf:"bytes,2,opt,name=request_time,json=requestTime" json:"request_time,omitempty"`
	// The time at which the operation failed or was completed successfully.
	FinishTime *google_protobuf1.Timestamp `protobuf:"bytes,3,opt,name=finish_time,json=finishTime" json:"finish_time,omitempty"`
}

func (m *CreateInstanceMetadata) Reset()         { *m = CreateInstanceMetadata{} }
func (m *CreateInstanceMetadata) String() string { return proto.CompactTextString(m) }
func (*CreateInstanceMetadata) ProtoMessage()    {}
func (*CreateInstanceMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptorBigtableInstanceAdmin, []int{10}
}

func (m *CreateInstanceMetadata) GetOriginalRequest() *CreateInstanceRequest {
	if m != nil {
		return m.OriginalRequest
	}
	return nil
}

func (m *CreateInstanceMetadata) GetRequestTime() *google_protobuf1.Timestamp {
	if m != nil {
		return m.RequestTime
	}
	return nil
}

func (m *CreateInstanceMetadata) GetFinishTime() *google_protobuf1.Timestamp {
	if m != nil {
		return m.FinishTime
	}
	return nil
}

// The metadata for the Operation returned by CreateCluster.
type CreateClusterMetadata struct {
	// The request that prompted the initiation of this CreateCluster operation.
	OriginalRequest *CreateClusterRequest `protobuf:"bytes,1,opt,name=original_request,json=originalRequest" json:"original_request,omitempty"`
	// The time at which the original request was received.
	RequestTime *google_protobuf1.Timestamp `protobuf:"bytes,2,opt,name=request_time,json=requestTime" json:"request_time,omitempty"`
	// The time at which the operation failed or was completed successfully.
	FinishTime *google_protobuf1.Timestamp `protobuf:"bytes,3,opt,name=finish_time,json=finishTime" json:"finish_time,omitempty"`
}

func (m *CreateClusterMetadata) Reset()         { *m = CreateClusterMetadata{} }
func (m *CreateClusterMetadata) String() string { return proto.CompactTextString(m) }
func (*CreateClusterMetadata) ProtoMessage()    {}
func (*CreateClusterMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptorBigtableInstanceAdmin, []int{11}
}

func (m *CreateClusterMetadata) GetOriginalRequest() *CreateClusterRequest {
	if m != nil {
		return m.OriginalRequest
	}
	return nil
}

func (m *CreateClusterMetadata) GetRequestTime() *google_protobuf1.Timestamp {
	if m != nil {
		return m.RequestTime
	}
	return nil
}

func (m *CreateClusterMetadata) GetFinishTime() *google_protobuf1.Timestamp {
	if m != nil {
		return m.FinishTime
	}
	return nil
}

// The metadata for the Operation returned by UpdateCluster.
type UpdateClusterMetadata struct {
	// The request that prompted the initiation of this UpdateCluster operation.
	OriginalRequest *Cluster `protobuf:"bytes,1,opt,name=original_request,json=originalRequest" json:"original_request,omitempty"`
	// The time at which the original request was received.
	RequestTime *google_protobuf1.Timestamp `protobuf:"bytes,2,opt,name=request_time,json=requestTime" json:"request_time,omitempty"`
	// The time at which the operation failed or was completed successfully.
	FinishTime *google_protobuf1.Timestamp `protobuf:"bytes,3,opt,name=finish_time,json=finishTime" json:"finish_time,omitempty"`
}

func (m *UpdateClusterMetadata) Reset()         { *m = UpdateClusterMetadata{} }
func (m *UpdateClusterMetadata) String() string { return proto.CompactTextString(m) }
func (*UpdateClusterMetadata) ProtoMessage()    {}
func (*UpdateClusterMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptorBigtableInstanceAdmin, []int{12}
}

func (m *UpdateClusterMetadata) GetOriginalRequest() *Cluster {
	if m != nil {
		return m.OriginalRequest
	}
	return nil
}

func (m *UpdateClusterMetadata) GetRequestTime() *google_protobuf1.Timestamp {
	if m != nil {
		return m.RequestTime
	}
	return nil
}

func (m *UpdateClusterMetadata) GetFinishTime() *google_protobuf1.Timestamp {
	if m != nil {
		return m.FinishTime
	}
	return nil
}

func init() {
	proto.RegisterType((*CreateInstanceRequest)(nil), "google.bigtable.admin.v2.CreateInstanceRequest")
	proto.RegisterType((*GetInstanceRequest)(nil), "google.bigtable.admin.v2.GetInstanceRequest")
	proto.RegisterType((*ListInstancesRequest)(nil), "google.bigtable.admin.v2.ListInstancesRequest")
	proto.RegisterType((*ListInstancesResponse)(nil), "google.bigtable.admin.v2.ListInstancesResponse")
	proto.RegisterType((*DeleteInstanceRequest)(nil), "google.bigtable.admin.v2.DeleteInstanceRequest")
	proto.RegisterType((*CreateClusterRequest)(nil), "google.bigtable.admin.v2.CreateClusterRequest")
	proto.RegisterType((*GetClusterRequest)(nil), "google.bigtable.admin.v2.GetClusterRequest")
	proto.RegisterType((*ListClustersRequest)(nil), "google.bigtable.admin.v2.ListClustersRequest")
	proto.RegisterType((*ListClustersResponse)(nil), "google.bigtable.admin.v2.ListClustersResponse")
	proto.RegisterType((*DeleteClusterRequest)(nil), "google.bigtable.admin.v2.DeleteClusterRequest")
	proto.RegisterType((*CreateInstanceMetadata)(nil), "google.bigtable.admin.v2.CreateInstanceMetadata")
	proto.RegisterType((*CreateClusterMetadata)(nil), "google.bigtable.admin.v2.CreateClusterMetadata")
	proto.RegisterType((*UpdateClusterMetadata)(nil), "google.bigtable.admin.v2.UpdateClusterMetadata")
}

func init() {
	proto.RegisterFile("google/bigtable/admin/v2/bigtable_instance_admin.proto", fileDescriptorBigtableInstanceAdmin)
}

var fileDescriptorBigtableInstanceAdmin = []byte{
	// 1011 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x56, 0x41, 0x6f, 0x1b, 0x45,
	0x14, 0xd6, 0xd8, 0xa5, 0x34, 0x6f, 0xeb, 0x24, 0x0c, 0x71, 0x64, 0x2d, 0x2d, 0x4d, 0xb7, 0x52,
	0xeb, 0xba, 0x61, 0x57, 0x2c, 0x88, 0xa0, 0x54, 0x46, 0x90, 0x50, 0x45, 0x91, 0x52, 0x11, 0x59,
	0xa5, 0x52, 0x51, 0x84, 0x35, 0xb1, 0x27, 0x66, 0xe9, 0x7a, 0x76, 0xd9, 0x1d, 0x47, 0x44, 0xa8,
	0x1c, 0x10, 0xe2, 0x50, 0x09, 0x0e, 0x70, 0x44, 0x9c, 0xb8, 0x20, 0xc4, 0x2f, 0x81, 0x23, 0x47,
	0x4e, 0x48, 0x9c, 0x11, 0x3f, 0x01, 0xcd, 0xce, 0xcc, 0xda, 0xeb, 0xac, 0xbd, 0x6b, 0x21, 0xa4,
	0xde, 0xec, 0x37, 0xef, 0xbd, 0xf9, 0xde, 0xf7, 0xbe, 0x7d, 0x6f, 0xe0, 0x8d, 0x41, 0x10, 0x0c,
	0x7c, 0xea, 0x1c, 0x7b, 0x03, 0x4e, 0x8e, 0x7d, 0xea, 0x90, 0xfe, 0xd0, 0x63, 0xce, 0xa9, 0x9b,
	0x5a, 0xba, 0x1e, 0x8b, 0x39, 0x61, 0x3d, 0xda, 0x4d, 0x8e, 0xec, 0x30, 0x0a, 0x78, 0x80, 0x1b,
	0x32, 0xce, 0xd6, 0x5e, 0xb6, 0x3c, 0x3c, 0x75, 0xcd, 0x2b, 0x2a, 0x23, 0x09, 0x3d, 0x87, 0x30,
	0x16, 0x70, 0xc2, 0xbd, 0x80, 0xc5, 0x32, 0xce, 0xbc, 0x35, 0xf3, 0x3e, 0x7d, 0x8d, 0x72, 0xbc,
	0xa1, 0x1c, 0xfd, 0x80, 0x0d, 0xa2, 0x11, 0x63, 0x1e, 0x1b, 0x38, 0x41, 0x48, 0xa3, 0x4c, 0xb6,
	0x97, 0x94, 0x53, 0xf2, 0xef, 0x78, 0x74, 0xe2, 0xd0, 0x61, 0xc8, 0xcf, 0xd4, 0xe1, 0xb5, 0xe9,
	0x43, 0xee, 0x0d, 0x69, 0xcc, 0xc9, 0x30, 0x94, 0x0e, 0xd6, 0x6f, 0x15, 0xa8, 0xef, 0x46, 0x94,
	0x70, 0xba, 0xaf, 0xee, 0xee, 0xd0, 0x4f, 0x46, 0x34, 0xe6, 0x78, 0x1d, 0x2e, 0x86, 0x24, 0xa2,
	0x8c, 0x37, 0xd0, 0x06, 0x6a, 0x2e, 0x75, 0xd4, 0x3f, 0x7c, 0x0d, 0x8c, 0x94, 0x0d, 0xaf, 0xdf,
	0xa8, 0x24, 0x87, 0xa0, 0x4d, 0xfb, 0x7d, 0xfc, 0x16, 0x5c, 0xd2, 0xff, 0x1a, 0xd5, 0x0d, 0xd4,
	0x34, 0x5c, 0xcb, 0x9e, 0xc5, 0x94, 0x9d, 0xde, 0x9a, 0xc6, 0xe0, 0x47, 0x70, 0xa9, 0xe7, 0x8f,
	0x62, 0x4e, 0xa3, 0xb8, 0x71, 0x61, 0xa3, 0xda, 0x34, 0xdc, 0xf6, 0xec, 0xf8, 0x5c, 0xec, 0xf6,
	0xae, 0x8a, 0xbf, 0xc7, 0x78, 0x74, 0xd6, 0x49, 0xd3, 0x99, 0x1f, 0x42, 0x2d, 0x73, 0x84, 0x57,
	0xa1, 0xfa, 0x98, 0x9e, 0xa9, 0x0a, 0xc5, 0x4f, 0xbc, 0x05, 0xcf, 0x9d, 0x12, 0x7f, 0x44, 0x93,
	0xc2, 0x0c, 0xf7, 0xfa, 0x9c, 0xab, 0x65, 0xa6, 0x8e, 0xf4, 0xdf, 0xae, 0xbc, 0x89, 0xac, 0x26,
	0xe0, 0x3d, 0xca, 0xa7, 0x99, 0xc4, 0x70, 0x81, 0x91, 0x21, 0x55, 0xb7, 0x24, 0xbf, 0xad, 0xfb,
	0xb0, 0x76, 0xe0, 0xc5, 0xa9, 0x6b, 0x5c, 0xc4, 0xfa, 0x55, 0x80, 0x90, 0x0c, 0x68, 0x97, 0x07,
	0x8f, 0x29, 0x53, 0xa4, 0x2f, 0x09, 0xcb, 0x03, 0x61, 0xb0, 0x7e, 0x41, 0x50, 0x9f, 0xca, 0x17,
	0x87, 0x01, 0x8b, 0x29, 0x7e, 0x1b, 0x96, 0x34, 0xb3, 0x71, 0x03, 0x25, 0x74, 0x96, 0x69, 0xc7,
	0x38, 0x08, 0xdf, 0x86, 0xd5, 0x13, 0xe2, 0xf9, 0xb4, 0xdf, 0xf5, 0x83, 0x9e, 0x94, 0x5e, 0xa3,
	0xb2, 0x51, 0x6d, 0x2e, 0x75, 0x56, 0xa4, 0xfd, 0x40, 0x9b, 0xf1, 0x4d, 0x58, 0x61, 0xf4, 0x53,
	0xde, 0x9d, 0x80, 0x5a, 0x4d, 0xa0, 0xd6, 0x84, 0xf9, 0x30, 0x85, 0x7b, 0x07, 0xea, 0xef, 0x52,
	0x9f, 0x9e, 0x17, 0x5d, 0x1e, 0x55, 0x4f, 0x11, 0xac, 0xc9, 0x36, 0x6b, 0xc6, 0x8b, 0xb9, 0x52,
	0x1d, 0x1f, 0x0b, 0x74, 0x49, 0x59, 0xf6, 0xfb, 0xf8, 0x2e, 0x3c, 0xaf, 0xfe, 0x28, 0x79, 0x96,
	0xe8, 0xb1, 0x8e, 0xb0, 0x6e, 0xc1, 0x0b, 0x7b, 0x94, 0x4f, 0x01, 0xc9, 0x43, 0x7d, 0x00, 0x2f,
	0x8a, 0x86, 0x68, 0xb9, 0xfd, 0xc7, 0xfe, 0xfe, 0x84, 0xa4, 0x5e, 0xc6, 0xe9, 0x54, 0x7b, 0xdb,
	0x13, 0x1f, 0x8b, 0xec, 0x6e, 0x89, 0x6a, 0xd2, 0x90, 0xff, 0xa3, 0xb7, 0x2d, 0x58, 0x93, 0xbd,
	0x2d, 0x41, 0xd2, 0x3f, 0x08, 0xd6, 0xb3, 0x5f, 0xf0, 0x7d, 0xca, 0x49, 0x9f, 0x70, 0x82, 0x3f,
	0x80, 0xd5, 0x20, 0xf2, 0x06, 0x1e, 0x23, 0x7e, 0x37, 0x92, 0x29, 0x92, 0x50, 0xc3, 0x75, 0x16,
	0x9c, 0x06, 0x9d, 0x15, 0x9d, 0x48, 0x43, 0x69, 0xc3, 0x65, 0x95, 0xb2, 0x2b, 0xe6, 0xa1, 0xfa,
	0xd4, 0x4d, 0x9d, 0x57, 0x0f, 0x4b, 0xfb, 0x81, 0x1e, 0x96, 0x1d, 0x43, 0xf9, 0x0b, 0x0b, 0xbe,
	0x0b, 0xc6, 0x89, 0xc7, 0xbc, 0xf8, 0x23, 0x19, 0x5d, 0x2d, 0x8c, 0x06, 0xe9, 0x2e, 0x0c, 0xd6,
	0xdf, 0x48, 0x0f, 0x5c, 0xc5, 0x4f, 0x5a, 0xf1, 0xa3, 0x99, 0x15, 0xdb, 0x45, 0x15, 0x67, 0xa9,
	0x7e, 0xb6, 0x0a, 0xfe, 0x13, 0x41, 0xfd, 0xfd, 0xb0, 0x9f, 0x53, 0xf0, 0xc1, 0xcc, 0x82, 0x4b,
	0x68, 0xf8, 0x59, 0xaa, 0xd1, 0xfd, 0xc3, 0x80, 0xfa, 0x8e, 0x82, 0xaa, 0xd5, 0xf7, 0x8e, 0x40,
	0x8c, 0xbf, 0x41, 0xb0, 0x9c, 0x55, 0x25, 0x5e, 0x54, 0xbf, 0xe6, 0x55, 0x1d, 0x30, 0xf1, 0x0e,
	0xb0, 0xdf, 0xd3, 0xef, 0x00, 0x6b, 0xf3, 0x8b, 0xdf, 0xff, 0xfa, 0xae, 0x72, 0xd3, 0xba, 0x2e,
	0x5e, 0x10, 0x9f, 0xc9, 0x79, 0xd2, 0x0e, 0xa3, 0xe0, 0x63, 0xda, 0xe3, 0xb1, 0xd3, 0x7a, 0x92,
	0xbe, 0x2a, 0xe2, 0x6d, 0xd4, 0xc2, 0x4f, 0x11, 0x18, 0x13, 0x3b, 0x0a, 0x6f, 0xce, 0x46, 0x73,
	0x7e, 0x95, 0x99, 0x25, 0x56, 0x87, 0x75, 0x3b, 0xc1, 0x73, 0x03, 0x4b, 0x3c, 0xe2, 0x3b, 0x9f,
	0x40, 0x33, 0x06, 0xe3, 0xb4, 0x9e, 0xe0, 0xef, 0x11, 0xd4, 0x32, 0x6b, 0x0b, 0xcf, 0x91, 0x7a,
	0xde, 0xbe, 0x34, 0x9d, 0xd2, 0xfe, 0x72, 0x60, 0x4e, 0xa1, 0x9b, 0xc7, 0x16, 0xfe, 0x0a, 0xc1,
	0xb2, 0x54, 0x6e, 0xca, 0x56, 0x89, 0xfa, 0x4b, 0x71, 0xa4, 0x7a, 0x66, 0x16, 0x73, 0x24, 0x7a,
	0xf6, 0x25, 0x82, 0xe5, 0xec, 0xbe, 0x9c, 0x27, 0xa2, 0xdc, 0xcd, 0x6a, 0xae, 0x9f, 0x93, 0xf2,
	0x3d, 0xf1, 0x4e, 0xd4, 0x7c, 0xb4, 0x4a, 0x74, 0xeb, 0x07, 0x04, 0xb5, 0xcc, 0xbc, 0xc1, 0x0b,
	0x0e, 0xa6, 0x22, 0x25, 0xb7, 0x13, 0x2c, 0x5b, 0xd6, 0x66, 0x7e, 0x6f, 0x32, 0x68, 0x1c, 0xbd,
	0xc3, 0xb6, 0xf5, 0x6e, 0xc6, 0xdf, 0x22, 0x80, 0xf1, 0x72, 0xc6, 0x77, 0xe6, 0x2a, 0x7b, 0x0a,
	0x59, 0xf1, 0xc4, 0xb1, 0x5e, 0x4f, 0xd0, 0xd9, 0x78, 0xb3, 0x88, 0xa9, 0x14, 0x9a, 0x20, 0xed,
	0x47, 0x04, 0x97, 0x27, 0x37, 0x37, 0x7e, 0x65, 0xbe, 0x62, 0xa7, 0x1e, 0x0c, 0xa6, 0x5d, 0xd6,
	0x5d, 0xe9, 0x3b, 0x8b, 0xb2, 0x24, 0x87, 0x62, 0x2a, 0xd4, 0x32, 0x43, 0x1a, 0x17, 0x13, 0x52,
	0xd4, 0xcd, 0xad, 0x04, 0xc9, 0xab, 0xe6, 0x42, 0x7c, 0x09, 0xb9, 0x7f, 0x8d, 0xa0, 0x96, 0x79,
	0x42, 0xcc, 0xd3, 0x59, 0xde, 0x5b, 0x63, 0xa6, 0xd8, 0x15, 0x39, 0xad, 0x85, 0x20, 0xed, 0x7c,
	0x0e, 0x57, 0x7a, 0xc1, 0x70, 0x26, 0x84, 0x1d, 0x33, 0x77, 0xf4, 0x1f, 0x8a, 0xab, 0x0f, 0xd1,
	0xcf, 0x95, 0x97, 0xf7, 0x64, 0xe0, 0xae, 0x1f, 0x8c, 0xfa, 0xb6, 0xf6, 0xb5, 0x13, 0x27, 0xfb,
	0xa1, 0xfb, 0xab, 0x76, 0x38, 0x4a, 0x1c, 0x8e, 0xb4, 0xc3, 0x51, 0xe2, 0x70, 0xf4, 0xd0, 0x3d,
	0xbe, 0x98, 0x54, 0xf1, 0xda, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x79, 0x3d, 0xe4, 0xaa, 0xa8,
	0x0e, 0x00, 0x00,
}