// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/functions/v1beta2/functions.proto

package google_cloud_functions_v1beta2

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "go.pedge.io/pb/go/google/api"
import _ "go.pedge.io/pb/go/google/api"
import _ "go.pedge.io/pb/go/google/longrunning"
import google_protobuf3 "github.com/golang/protobuf/ptypes/duration"
import google_protobuf4 "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Describes the current stage of a deployment.
type CloudFunctionStatus int32

const (
	// Status not specified.
	CloudFunctionStatus_STATUS_UNSPECIFIED CloudFunctionStatus = 0
	// Successfully deployed.
	CloudFunctionStatus_READY CloudFunctionStatus = 1
	// Not deployed correctly - behavior is undefined. The item should be updated
	// or deleted to move it out of this state.
	CloudFunctionStatus_FAILED CloudFunctionStatus = 2
	// Creation or update in progress.
	CloudFunctionStatus_DEPLOYING CloudFunctionStatus = 3
	// Deletion in progress.
	CloudFunctionStatus_DELETING CloudFunctionStatus = 4
)

var CloudFunctionStatus_name = map[int32]string{
	0: "STATUS_UNSPECIFIED",
	1: "READY",
	2: "FAILED",
	3: "DEPLOYING",
	4: "DELETING",
}
var CloudFunctionStatus_value = map[string]int32{
	"STATUS_UNSPECIFIED": 0,
	"READY":              1,
	"FAILED":             2,
	"DEPLOYING":          3,
	"DELETING":           4,
}

func (x CloudFunctionStatus) String() string {
	return proto.EnumName(CloudFunctionStatus_name, int32(x))
}
func (CloudFunctionStatus) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// Describes a Cloud Function that contains user computation executed in
// response to an event. It encapsulate function and triggers configurations.
type CloudFunction struct {
	// A user-defined name of the function. Function names must be unique
	// globally and match pattern `projects/*/locations/*/functions/*`
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// The location of the function source code.
	//
	// Types that are valid to be assigned to SourceCode:
	//	*CloudFunction_SourceArchiveUrl
	//	*CloudFunction_SourceRepository
	SourceCode isCloudFunction_SourceCode `protobuf_oneof:"source_code"`
	// An event that triggers the function.
	//
	// Types that are valid to be assigned to Trigger:
	//	*CloudFunction_HttpsTrigger
	//	*CloudFunction_EventTrigger
	Trigger isCloudFunction_Trigger `protobuf_oneof:"trigger"`
	// Output only. Status of the function deployment.
	Status CloudFunctionStatus `protobuf:"varint,7,opt,name=status,enum=google.cloud.functions.v1beta2.CloudFunctionStatus" json:"status,omitempty"`
	// Output only. Name of the most recent operation modifying the function. If
	// the function status is `DEPLOYING` or `DELETING`, then it points to the
	// active operation.
	LatestOperation string `protobuf:"bytes,8,opt,name=latest_operation,json=latestOperation" json:"latest_operation,omitempty"`
	// The name of the function (as defined in source code) that will be
	// executed. Defaults to the resource name suffix, if not specified. For
	// backward compatibility, if function with given name is not found, then the
	// system will try to use function named "function".
	// For Node.js this is name of a function exported by the module specified
	// in `source_location`.
	EntryPoint string `protobuf:"bytes,9,opt,name=entry_point,json=entryPoint" json:"entry_point,omitempty"`
	// The function execution timeout. Execution is considered failed and
	// can be terminated if the function is not completed at the end of the
	// timeout period. Defaults to 60 seconds.
	Timeout *google_protobuf3.Duration `protobuf:"bytes,10,opt,name=timeout" json:"timeout,omitempty"`
	// The amount of memory in MB available for a function.
	// Defaults to 256MB.
	AvailableMemoryMb int32 `protobuf:"varint,11,opt,name=available_memory_mb,json=availableMemoryMb" json:"available_memory_mb,omitempty"`
	// Output only. The service account of the function.
	ServiceAccount string `protobuf:"bytes,13,opt,name=service_account,json=serviceAccount" json:"service_account,omitempty"`
	// Output only. The last update timestamp of a Cloud Function.
	UpdateTime *google_protobuf4.Timestamp `protobuf:"bytes,15,opt,name=update_time,json=updateTime" json:"update_time,omitempty"`
}

func (m *CloudFunction) Reset()                    { *m = CloudFunction{} }
func (m *CloudFunction) String() string            { return proto.CompactTextString(m) }
func (*CloudFunction) ProtoMessage()               {}
func (*CloudFunction) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type isCloudFunction_SourceCode interface {
	isCloudFunction_SourceCode()
}
type isCloudFunction_Trigger interface {
	isCloudFunction_Trigger()
}

type CloudFunction_SourceArchiveUrl struct {
	SourceArchiveUrl string `protobuf:"bytes,14,opt,name=source_archive_url,json=sourceArchiveUrl,oneof"`
}
type CloudFunction_SourceRepository struct {
	SourceRepository *SourceRepository `protobuf:"bytes,3,opt,name=source_repository,json=sourceRepository,oneof"`
}
type CloudFunction_HttpsTrigger struct {
	HttpsTrigger *HTTPSTrigger `protobuf:"bytes,6,opt,name=https_trigger,json=httpsTrigger,oneof"`
}
type CloudFunction_EventTrigger struct {
	EventTrigger *EventTrigger `protobuf:"bytes,12,opt,name=event_trigger,json=eventTrigger,oneof"`
}

func (*CloudFunction_SourceArchiveUrl) isCloudFunction_SourceCode() {}
func (*CloudFunction_SourceRepository) isCloudFunction_SourceCode() {}
func (*CloudFunction_HttpsTrigger) isCloudFunction_Trigger()        {}
func (*CloudFunction_EventTrigger) isCloudFunction_Trigger()        {}

func (m *CloudFunction) GetSourceCode() isCloudFunction_SourceCode {
	if m != nil {
		return m.SourceCode
	}
	return nil
}
func (m *CloudFunction) GetTrigger() isCloudFunction_Trigger {
	if m != nil {
		return m.Trigger
	}
	return nil
}

func (m *CloudFunction) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CloudFunction) GetSourceArchiveUrl() string {
	if x, ok := m.GetSourceCode().(*CloudFunction_SourceArchiveUrl); ok {
		return x.SourceArchiveUrl
	}
	return ""
}

func (m *CloudFunction) GetSourceRepository() *SourceRepository {
	if x, ok := m.GetSourceCode().(*CloudFunction_SourceRepository); ok {
		return x.SourceRepository
	}
	return nil
}

func (m *CloudFunction) GetHttpsTrigger() *HTTPSTrigger {
	if x, ok := m.GetTrigger().(*CloudFunction_HttpsTrigger); ok {
		return x.HttpsTrigger
	}
	return nil
}

func (m *CloudFunction) GetEventTrigger() *EventTrigger {
	if x, ok := m.GetTrigger().(*CloudFunction_EventTrigger); ok {
		return x.EventTrigger
	}
	return nil
}

func (m *CloudFunction) GetStatus() CloudFunctionStatus {
	if m != nil {
		return m.Status
	}
	return CloudFunctionStatus_STATUS_UNSPECIFIED
}

func (m *CloudFunction) GetLatestOperation() string {
	if m != nil {
		return m.LatestOperation
	}
	return ""
}

func (m *CloudFunction) GetEntryPoint() string {
	if m != nil {
		return m.EntryPoint
	}
	return ""
}

func (m *CloudFunction) GetTimeout() *google_protobuf3.Duration {
	if m != nil {
		return m.Timeout
	}
	return nil
}

func (m *CloudFunction) GetAvailableMemoryMb() int32 {
	if m != nil {
		return m.AvailableMemoryMb
	}
	return 0
}

func (m *CloudFunction) GetServiceAccount() string {
	if m != nil {
		return m.ServiceAccount
	}
	return ""
}

func (m *CloudFunction) GetUpdateTime() *google_protobuf4.Timestamp {
	if m != nil {
		return m.UpdateTime
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*CloudFunction) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _CloudFunction_OneofMarshaler, _CloudFunction_OneofUnmarshaler, _CloudFunction_OneofSizer, []interface{}{
		(*CloudFunction_SourceArchiveUrl)(nil),
		(*CloudFunction_SourceRepository)(nil),
		(*CloudFunction_HttpsTrigger)(nil),
		(*CloudFunction_EventTrigger)(nil),
	}
}

func _CloudFunction_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*CloudFunction)
	// source_code
	switch x := m.SourceCode.(type) {
	case *CloudFunction_SourceArchiveUrl:
		b.EncodeVarint(14<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.SourceArchiveUrl)
	case *CloudFunction_SourceRepository:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.SourceRepository); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("CloudFunction.SourceCode has unexpected type %T", x)
	}
	// trigger
	switch x := m.Trigger.(type) {
	case *CloudFunction_HttpsTrigger:
		b.EncodeVarint(6<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.HttpsTrigger); err != nil {
			return err
		}
	case *CloudFunction_EventTrigger:
		b.EncodeVarint(12<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.EventTrigger); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("CloudFunction.Trigger has unexpected type %T", x)
	}
	return nil
}

func _CloudFunction_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*CloudFunction)
	switch tag {
	case 14: // source_code.source_archive_url
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.SourceCode = &CloudFunction_SourceArchiveUrl{x}
		return true, err
	case 3: // source_code.source_repository
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(SourceRepository)
		err := b.DecodeMessage(msg)
		m.SourceCode = &CloudFunction_SourceRepository{msg}
		return true, err
	case 6: // trigger.https_trigger
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(HTTPSTrigger)
		err := b.DecodeMessage(msg)
		m.Trigger = &CloudFunction_HttpsTrigger{msg}
		return true, err
	case 12: // trigger.event_trigger
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(EventTrigger)
		err := b.DecodeMessage(msg)
		m.Trigger = &CloudFunction_EventTrigger{msg}
		return true, err
	default:
		return false, nil
	}
}

func _CloudFunction_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*CloudFunction)
	// source_code
	switch x := m.SourceCode.(type) {
	case *CloudFunction_SourceArchiveUrl:
		n += proto.SizeVarint(14<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.SourceArchiveUrl)))
		n += len(x.SourceArchiveUrl)
	case *CloudFunction_SourceRepository:
		s := proto.Size(x.SourceRepository)
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	// trigger
	switch x := m.Trigger.(type) {
	case *CloudFunction_HttpsTrigger:
		s := proto.Size(x.HttpsTrigger)
		n += proto.SizeVarint(6<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *CloudFunction_EventTrigger:
		s := proto.Size(x.EventTrigger)
		n += proto.SizeVarint(12<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// Describes HTTPSTrigger, could be used to connect web hooks to function.
type HTTPSTrigger struct {
	// Output only. The deployed url for the function.
	Url string `protobuf:"bytes,1,opt,name=url" json:"url,omitempty"`
}

func (m *HTTPSTrigger) Reset()                    { *m = HTTPSTrigger{} }
func (m *HTTPSTrigger) String() string            { return proto.CompactTextString(m) }
func (*HTTPSTrigger) ProtoMessage()               {}
func (*HTTPSTrigger) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *HTTPSTrigger) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

// Describes EventTrigger, used to request events be sent from another
// service.
type EventTrigger struct {
	// `event_type` names contain the service that is sending an event and the
	// kind of event that was fired. Must be of the form
	// `providers/*/eventTypes/*` e.g. Directly handle a Message published to
	// Google Cloud Pub/Sub `providers/cloud.pubsub/eventTypes/topic.publish`
	//
	//      Handle an object changing in Google Cloud Storage
	//      `providers/cloud.storage/eventTypes/object.change`
	//
	//      Handle a write to the Firebase Realtime Database
	//      `providers/firebase.database/eventTypes/data.write`
	EventType string `protobuf:"bytes,1,opt,name=event_type,json=eventType" json:"event_type,omitempty"`
	// Which instance of the source's service should send events. E.g. for Pub/Sub
	// this would be a Pub/Sub topic at `projects/*/topics/*`. For Google Cloud
	// Storage this would be a bucket at `projects/*/buckets/*`. For any source
	// that only supports one instance per-project, this should be the name of the
	// project (`projects/*`)
	Resource string `protobuf:"bytes,2,opt,name=resource" json:"resource,omitempty"`
}

func (m *EventTrigger) Reset()                    { *m = EventTrigger{} }
func (m *EventTrigger) String() string            { return proto.CompactTextString(m) }
func (*EventTrigger) ProtoMessage()               {}
func (*EventTrigger) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *EventTrigger) GetEventType() string {
	if m != nil {
		return m.EventType
	}
	return ""
}

func (m *EventTrigger) GetResource() string {
	if m != nil {
		return m.Resource
	}
	return ""
}

// Describes the location of the function source in a remote repository.
type SourceRepository struct {
	// URL to the hosted repository where the function is defined. Only paths in
	// https://source.developers.google.com domain are supported. The path should
	// contain the name of the repository.
	RepositoryUrl string `protobuf:"bytes,1,opt,name=repository_url,json=repositoryUrl" json:"repository_url,omitempty"`
	// The path within the repository where the function is defined. The path
	// should point to the directory where Cloud Functions files are located. Use
	// "/" if the function is defined directly in the root directory of a
	// repository.
	SourcePath string `protobuf:"bytes,2,opt,name=source_path,json=sourcePath" json:"source_path,omitempty"`
	// The version of a function. Defaults to the latest version of the master
	// branch.
	//
	// Types that are valid to be assigned to Version:
	//	*SourceRepository_Branch
	//	*SourceRepository_Tag
	//	*SourceRepository_Revision
	Version isSourceRepository_Version `protobuf_oneof:"version"`
	// Output only. The id of the revision that was resolved at the moment of
	// function creation or update. For example when a user deployed from a
	// branch, it will be the revision id of the latest change on this branch at
	// that time. If user deployed from revision then this value will be always
	// equal to the revision specified by the user.
	DeployedRevision string `protobuf:"bytes,6,opt,name=deployed_revision,json=deployedRevision" json:"deployed_revision,omitempty"`
}

func (m *SourceRepository) Reset()                    { *m = SourceRepository{} }
func (m *SourceRepository) String() string            { return proto.CompactTextString(m) }
func (*SourceRepository) ProtoMessage()               {}
func (*SourceRepository) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type isSourceRepository_Version interface {
	isSourceRepository_Version()
}

type SourceRepository_Branch struct {
	Branch string `protobuf:"bytes,3,opt,name=branch,oneof"`
}
type SourceRepository_Tag struct {
	Tag string `protobuf:"bytes,4,opt,name=tag,oneof"`
}
type SourceRepository_Revision struct {
	Revision string `protobuf:"bytes,5,opt,name=revision,oneof"`
}

func (*SourceRepository_Branch) isSourceRepository_Version()   {}
func (*SourceRepository_Tag) isSourceRepository_Version()      {}
func (*SourceRepository_Revision) isSourceRepository_Version() {}

func (m *SourceRepository) GetVersion() isSourceRepository_Version {
	if m != nil {
		return m.Version
	}
	return nil
}

func (m *SourceRepository) GetRepositoryUrl() string {
	if m != nil {
		return m.RepositoryUrl
	}
	return ""
}

func (m *SourceRepository) GetSourcePath() string {
	if m != nil {
		return m.SourcePath
	}
	return ""
}

func (m *SourceRepository) GetBranch() string {
	if x, ok := m.GetVersion().(*SourceRepository_Branch); ok {
		return x.Branch
	}
	return ""
}

func (m *SourceRepository) GetTag() string {
	if x, ok := m.GetVersion().(*SourceRepository_Tag); ok {
		return x.Tag
	}
	return ""
}

func (m *SourceRepository) GetRevision() string {
	if x, ok := m.GetVersion().(*SourceRepository_Revision); ok {
		return x.Revision
	}
	return ""
}

func (m *SourceRepository) GetDeployedRevision() string {
	if m != nil {
		return m.DeployedRevision
	}
	return ""
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*SourceRepository) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _SourceRepository_OneofMarshaler, _SourceRepository_OneofUnmarshaler, _SourceRepository_OneofSizer, []interface{}{
		(*SourceRepository_Branch)(nil),
		(*SourceRepository_Tag)(nil),
		(*SourceRepository_Revision)(nil),
	}
}

func _SourceRepository_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*SourceRepository)
	// version
	switch x := m.Version.(type) {
	case *SourceRepository_Branch:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.Branch)
	case *SourceRepository_Tag:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.Tag)
	case *SourceRepository_Revision:
		b.EncodeVarint(5<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.Revision)
	case nil:
	default:
		return fmt.Errorf("SourceRepository.Version has unexpected type %T", x)
	}
	return nil
}

func _SourceRepository_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*SourceRepository)
	switch tag {
	case 3: // version.branch
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Version = &SourceRepository_Branch{x}
		return true, err
	case 4: // version.tag
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Version = &SourceRepository_Tag{x}
		return true, err
	case 5: // version.revision
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Version = &SourceRepository_Revision{x}
		return true, err
	default:
		return false, nil
	}
}

func _SourceRepository_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*SourceRepository)
	// version
	switch x := m.Version.(type) {
	case *SourceRepository_Branch:
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.Branch)))
		n += len(x.Branch)
	case *SourceRepository_Tag:
		n += proto.SizeVarint(4<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.Tag)))
		n += len(x.Tag)
	case *SourceRepository_Revision:
		n += proto.SizeVarint(5<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.Revision)))
		n += len(x.Revision)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// Request for the `CreateFunction` method.
type CreateFunctionRequest struct {
	// The project and location in which the function should be created, specified
	// in the format `projects/*/locations/*`
	Location string `protobuf:"bytes,1,opt,name=location" json:"location,omitempty"`
	// Function to be created.
	Function *CloudFunction `protobuf:"bytes,2,opt,name=function" json:"function,omitempty"`
}

func (m *CreateFunctionRequest) Reset()                    { *m = CreateFunctionRequest{} }
func (m *CreateFunctionRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateFunctionRequest) ProtoMessage()               {}
func (*CreateFunctionRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *CreateFunctionRequest) GetLocation() string {
	if m != nil {
		return m.Location
	}
	return ""
}

func (m *CreateFunctionRequest) GetFunction() *CloudFunction {
	if m != nil {
		return m.Function
	}
	return nil
}

// Request for the `UpdateFunction` method.
type UpdateFunctionRequest struct {
	// The name of the function to be updated.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// New version of the function.
	Function *CloudFunction `protobuf:"bytes,2,opt,name=function" json:"function,omitempty"`
}

func (m *UpdateFunctionRequest) Reset()                    { *m = UpdateFunctionRequest{} }
func (m *UpdateFunctionRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdateFunctionRequest) ProtoMessage()               {}
func (*UpdateFunctionRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *UpdateFunctionRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UpdateFunctionRequest) GetFunction() *CloudFunction {
	if m != nil {
		return m.Function
	}
	return nil
}

// Request for the `GetFunction` method.
type GetFunctionRequest struct {
	// The name of the function which details should be obtained.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *GetFunctionRequest) Reset()                    { *m = GetFunctionRequest{} }
func (m *GetFunctionRequest) String() string            { return proto.CompactTextString(m) }
func (*GetFunctionRequest) ProtoMessage()               {}
func (*GetFunctionRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *GetFunctionRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// Request for the `ListFunctions` method.
type ListFunctionsRequest struct {
	// The project and location from which the function should be listed,
	// specified in the format `projects/*/locations/*`
	// If you want to list functions in all locations, use "-" in place of a
	// location.
	Location string `protobuf:"bytes,1,opt,name=location" json:"location,omitempty"`
	// Maximum number of functions to return per call.
	PageSize int32 `protobuf:"varint,2,opt,name=page_size,json=pageSize" json:"page_size,omitempty"`
	// The value returned by the last
	// `ListFunctionsResponse`; indicates that
	// this is a continuation of a prior `ListFunctions` call, and that the
	// system should return the next page of data.
	PageToken string `protobuf:"bytes,3,opt,name=page_token,json=pageToken" json:"page_token,omitempty"`
}

func (m *ListFunctionsRequest) Reset()                    { *m = ListFunctionsRequest{} }
func (m *ListFunctionsRequest) String() string            { return proto.CompactTextString(m) }
func (*ListFunctionsRequest) ProtoMessage()               {}
func (*ListFunctionsRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *ListFunctionsRequest) GetLocation() string {
	if m != nil {
		return m.Location
	}
	return ""
}

func (m *ListFunctionsRequest) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *ListFunctionsRequest) GetPageToken() string {
	if m != nil {
		return m.PageToken
	}
	return ""
}

// Response for the `ListFunctions` method.
type ListFunctionsResponse struct {
	// The functions that match the request.
	Functions []*CloudFunction `protobuf:"bytes,1,rep,name=functions" json:"functions,omitempty"`
	// If not empty, indicates that there may be more functions that match
	// the request; this value should be passed in a new
	// [google.cloud.functions.v1beta2.ListFunctionsRequest][]
	// to get more functions.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken" json:"next_page_token,omitempty"`
}

func (m *ListFunctionsResponse) Reset()                    { *m = ListFunctionsResponse{} }
func (m *ListFunctionsResponse) String() string            { return proto.CompactTextString(m) }
func (*ListFunctionsResponse) ProtoMessage()               {}
func (*ListFunctionsResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *ListFunctionsResponse) GetFunctions() []*CloudFunction {
	if m != nil {
		return m.Functions
	}
	return nil
}

func (m *ListFunctionsResponse) GetNextPageToken() string {
	if m != nil {
		return m.NextPageToken
	}
	return ""
}

// Request for the `DeleteFunction` method.
type DeleteFunctionRequest struct {
	// The name of the function which should be deleted.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *DeleteFunctionRequest) Reset()                    { *m = DeleteFunctionRequest{} }
func (m *DeleteFunctionRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteFunctionRequest) ProtoMessage()               {}
func (*DeleteFunctionRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *DeleteFunctionRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// Request for the `CallFunction` method.
type CallFunctionRequest struct {
	// The name of the function to be called.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// Input to be passed to the function.
	Data string `protobuf:"bytes,2,opt,name=data" json:"data,omitempty"`
}

func (m *CallFunctionRequest) Reset()                    { *m = CallFunctionRequest{} }
func (m *CallFunctionRequest) String() string            { return proto.CompactTextString(m) }
func (*CallFunctionRequest) ProtoMessage()               {}
func (*CallFunctionRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *CallFunctionRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CallFunctionRequest) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

// Response of `CallFunction` method.
type CallFunctionResponse struct {
	// Execution id of function invocation.
	ExecutionId string `protobuf:"bytes,1,opt,name=execution_id,json=executionId" json:"execution_id,omitempty"`
	// Result populated for successful execution of synchronous function. Will
	// not be populated if function does not return a result through context.
	Result string `protobuf:"bytes,2,opt,name=result" json:"result,omitempty"`
	// Either system or user-function generated error. Set if execution
	// was not successful.
	Error string `protobuf:"bytes,3,opt,name=error" json:"error,omitempty"`
}

func (m *CallFunctionResponse) Reset()                    { *m = CallFunctionResponse{} }
func (m *CallFunctionResponse) String() string            { return proto.CompactTextString(m) }
func (*CallFunctionResponse) ProtoMessage()               {}
func (*CallFunctionResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *CallFunctionResponse) GetExecutionId() string {
	if m != nil {
		return m.ExecutionId
	}
	return ""
}

func (m *CallFunctionResponse) GetResult() string {
	if m != nil {
		return m.Result
	}
	return ""
}

func (m *CallFunctionResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func init() {
	proto.RegisterType((*CloudFunction)(nil), "google.cloud.functions.v1beta2.CloudFunction")
	proto.RegisterType((*HTTPSTrigger)(nil), "google.cloud.functions.v1beta2.HTTPSTrigger")
	proto.RegisterType((*EventTrigger)(nil), "google.cloud.functions.v1beta2.EventTrigger")
	proto.RegisterType((*SourceRepository)(nil), "google.cloud.functions.v1beta2.SourceRepository")
	proto.RegisterType((*CreateFunctionRequest)(nil), "google.cloud.functions.v1beta2.CreateFunctionRequest")
	proto.RegisterType((*UpdateFunctionRequest)(nil), "google.cloud.functions.v1beta2.UpdateFunctionRequest")
	proto.RegisterType((*GetFunctionRequest)(nil), "google.cloud.functions.v1beta2.GetFunctionRequest")
	proto.RegisterType((*ListFunctionsRequest)(nil), "google.cloud.functions.v1beta2.ListFunctionsRequest")
	proto.RegisterType((*ListFunctionsResponse)(nil), "google.cloud.functions.v1beta2.ListFunctionsResponse")
	proto.RegisterType((*DeleteFunctionRequest)(nil), "google.cloud.functions.v1beta2.DeleteFunctionRequest")
	proto.RegisterType((*CallFunctionRequest)(nil), "google.cloud.functions.v1beta2.CallFunctionRequest")
	proto.RegisterType((*CallFunctionResponse)(nil), "google.cloud.functions.v1beta2.CallFunctionResponse")
	proto.RegisterEnum("google.cloud.functions.v1beta2.CloudFunctionStatus", CloudFunctionStatus_name, CloudFunctionStatus_value)
}

func init() { proto.RegisterFile("google/cloud/functions/v1beta2/functions.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 1185 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x56, 0x4d, 0x73, 0xdb, 0x44,
	0x18, 0xae, 0x92, 0xd8, 0xb5, 0x5f, 0x7f, 0xc4, 0xdd, 0xd6, 0x1d, 0x61, 0x5a, 0x6a, 0xc4, 0x00,
	0xc1, 0x05, 0x1b, 0xdc, 0x00, 0x33, 0xa5, 0x65, 0x48, 0x62, 0xa7, 0xf5, 0x34, 0x6d, 0x3d, 0xb2,
	0x73, 0xe8, 0x49, 0xb3, 0x96, 0xb7, 0xb6, 0x40, 0xd6, 0x0a, 0x69, 0xe5, 0xa9, 0xcb, 0x94, 0x03,
	0x67, 0x6e, 0xfc, 0x83, 0x9e, 0x19, 0x7e, 0x00, 0x33, 0xdc, 0xb9, 0x73, 0x65, 0x38, 0xf1, 0x43,
	0x98, 0x5d, 0xad, 0x64, 0x39, 0x31, 0x28, 0xc9, 0x70, 0xd3, 0x3e, 0xef, 0xd7, 0xb3, 0xfb, 0x3e,
	0xfb, 0x6a, 0xa1, 0x39, 0xa1, 0x74, 0x62, 0x93, 0x96, 0x69, 0xd3, 0x60, 0xdc, 0x7a, 0x1e, 0x38,
	0x26, 0xb3, 0xa8, 0xe3, 0xb7, 0xe6, 0x9f, 0x8c, 0x08, 0xc3, 0xed, 0x25, 0xd2, 0x74, 0x3d, 0xca,
	0x28, 0x7a, 0x2b, 0xf4, 0x6f, 0x0a, 0xff, 0xe6, 0xd2, 0x2a, 0xfd, 0x6b, 0x37, 0x64, 0x3e, 0xec,
	0x5a, 0x2d, 0xec, 0x38, 0x94, 0xe1, 0x44, 0x74, 0xad, 0x9a, 0xb4, 0x06, 0x6c, 0x2a, 0xe1, 0x56,
	0x0a, 0x09, 0xea, 0x12, 0x6f, 0x25, 0xcf, 0x3b, 0x32, 0xc0, 0xa6, 0xce, 0xc4, 0x0b, 0x1c, 0xc7,
	0x72, 0x26, 0xa7, 0x9d, 0x24, 0xd5, 0x96, 0x58, 0x8d, 0x82, 0xe7, 0xad, 0x71, 0x10, 0x3a, 0x48,
	0xfb, 0xad, 0x93, 0x76, 0x66, 0xcd, 0x88, 0xcf, 0xf0, 0xcc, 0x0d, 0x1d, 0xb4, 0x3f, 0x33, 0x50,
	0x3a, 0xe0, 0x94, 0x0e, 0x25, 0x23, 0x84, 0x60, 0xcb, 0xc1, 0x33, 0xa2, 0x2a, 0x75, 0x65, 0x27,
	0xaf, 0x8b, 0x6f, 0xd4, 0x04, 0xe4, 0xd3, 0xc0, 0x33, 0x89, 0x81, 0x3d, 0x73, 0x6a, 0xcd, 0x89,
	0x11, 0x78, 0xb6, 0x5a, 0xe6, 0x1e, 0x0f, 0x2f, 0xe9, 0x95, 0xd0, 0xb6, 0x17, 0x9a, 0x8e, 0x3d,
	0x1b, 0x19, 0x70, 0x45, 0xfa, 0x7b, 0xc4, 0xa5, 0xbe, 0xc5, 0xa8, 0xb7, 0x50, 0x37, 0xeb, 0xca,
	0x4e, 0xa1, 0xfd, 0x71, 0xf3, 0xbf, 0x4f, 0xb7, 0x39, 0x10, 0x81, 0x7a, 0x1c, 0xb7, 0x2c, 0xb0,
	0xc4, 0xd0, 0x00, 0x4a, 0x53, 0xc6, 0x5c, 0xdf, 0x60, 0x9e, 0x35, 0x99, 0x10, 0x4f, 0xcd, 0x8a,
	0xe4, 0x1f, 0xa6, 0x25, 0x7f, 0x38, 0x1c, 0xf6, 0x07, 0xc3, 0x30, 0xe6, 0xa1, 0xa2, 0x17, 0x45,
	0x12, 0xb9, 0xe6, 0x49, 0xc9, 0x9c, 0x38, 0x2c, 0x4e, 0x5a, 0x3c, 0x5b, 0xd2, 0x2e, 0x0f, 0x4a,
	0x24, 0x25, 0x89, 0x35, 0x7a, 0x04, 0x59, 0x9f, 0x61, 0x16, 0xf8, 0xea, 0xe5, 0xba, 0xb2, 0x53,
	0x6e, 0xdf, 0x49, 0xcb, 0xb6, 0xd2, 0x8d, 0x81, 0x08, 0xd5, 0x65, 0x0a, 0xf4, 0x01, 0x54, 0x6c,
	0xcc, 0x88, 0xcf, 0x8c, 0x58, 0x09, 0x6a, 0x4e, 0xf4, 0x69, 0x3b, 0xc4, 0x9f, 0x46, 0x30, 0xba,
	0x05, 0x05, 0xe2, 0x30, 0x6f, 0x61, 0xb8, 0xd4, 0x72, 0x98, 0x9a, 0x17, 0x5e, 0x20, 0xa0, 0x3e,
	0x47, 0xd0, 0x1d, 0xb8, 0xcc, 0xc5, 0x40, 0x03, 0xa6, 0x82, 0xd8, 0xe7, 0x1b, 0x11, 0xb3, 0x48,
	0x2c, 0xcd, 0x8e, 0x14, 0x93, 0x1e, 0x79, 0xa2, 0x26, 0x5c, 0xc5, 0x73, 0x6c, 0xd9, 0x78, 0x64,
	0x13, 0x63, 0x46, 0x66, 0xd4, 0x5b, 0x18, 0xb3, 0x91, 0x5a, 0xa8, 0x2b, 0x3b, 0x19, 0xfd, 0x4a,
	0x6c, 0x7a, 0x2c, 0x2c, 0x8f, 0x47, 0xe8, 0x7d, 0xd8, 0xf6, 0x89, 0x37, 0xb7, 0xb8, 0x72, 0x4c,
	0x93, 0x06, 0x0e, 0x53, 0x4b, 0x82, 0x49, 0x59, 0xc2, 0x7b, 0x21, 0x8a, 0xbe, 0x80, 0x42, 0xe0,
	0x8e, 0x31, 0x23, 0x06, 0x2f, 0xa5, 0x6e, 0x0b, 0x46, 0xb5, 0x53, 0x8c, 0x86, 0x91, 0x7c, 0x75,
	0x08, 0xdd, 0x39, 0xb0, 0x5f, 0x82, 0x82, 0x94, 0x9b, 0x49, 0xc7, 0x64, 0x3f, 0x0f, 0x97, 0x65,
	0x07, 0xb5, 0x3a, 0x14, 0x93, 0x2d, 0x47, 0x15, 0xd8, 0xe4, 0xca, 0x0d, 0xb5, 0xcd, 0x3f, 0xb5,
	0x1e, 0x14, 0x93, 0xfd, 0x43, 0x37, 0x01, 0xa4, 0x08, 0x16, 0x6e, 0x74, 0x09, 0xf2, 0x61, 0x47,
	0x17, 0x2e, 0x41, 0x35, 0xc8, 0x79, 0x24, 0x2c, 0xa6, 0x6e, 0x08, 0x63, 0xbc, 0xd6, 0xfe, 0x52,
	0xa0, 0x72, 0x52, 0xbd, 0xe8, 0x5d, 0x28, 0x2f, 0xef, 0x80, 0xb1, 0x2c, 0x5e, 0x5a, 0xa2, 0xfc,
	0xc6, 0xdc, 0x8a, 0xb7, 0xe0, 0x62, 0x36, 0x95, 0xa9, 0x21, 0x84, 0xfa, 0x98, 0x4d, 0x91, 0x0a,
	0xd9, 0x91, 0x87, 0x1d, 0x73, 0x2a, 0xee, 0x11, 0xbf, 0x76, 0x72, 0x8d, 0x10, 0x6c, 0x32, 0x3c,
	0x51, 0xb7, 0x24, 0xcc, 0x17, 0xe8, 0x06, 0xa7, 0x39, 0xb7, 0x7c, 0x2e, 0x90, 0x8c, 0x34, 0xc4,
	0x08, 0xba, 0x0d, 0x57, 0xc6, 0xc4, 0xb5, 0xe9, 0x82, 0x8c, 0x8d, 0xd8, 0x2d, 0x2b, 0x4a, 0x56,
	0x22, 0x83, 0x2e, 0x71, 0x7e, 0x9a, 0x73, 0xe2, 0xf1, 0x4f, 0xed, 0x7b, 0xa8, 0x1e, 0x78, 0x04,
	0x33, 0x12, 0xc9, 0x53, 0x27, 0xdf, 0x06, 0xc4, 0x67, 0xfc, 0x54, 0x6c, 0x6a, 0x86, 0x7a, 0x0c,
	0xb7, 0x17, 0xaf, 0x51, 0x0f, 0x72, 0x91, 0xc8, 0xc5, 0xb6, 0x0a, 0xed, 0x8f, 0xce, 0x75, 0x05,
	0xf4, 0x38, 0x5c, 0x9b, 0x43, 0xf5, 0x58, 0x74, 0xfd, 0x64, 0xfd, 0x75, 0x33, 0xeb, 0x7f, 0xac,
	0xbb, 0x03, 0xe8, 0x01, 0x61, 0x67, 0x28, 0xaa, 0x39, 0x70, 0xed, 0xc8, 0xf2, 0x63, 0x57, 0xff,
	0x2c, 0x07, 0xf4, 0x26, 0xe4, 0x5d, 0x3c, 0x21, 0x86, 0x6f, 0xbd, 0x0c, 0x35, 0x95, 0xd1, 0x73,
	0x1c, 0x18, 0x58, 0x2f, 0x09, 0x97, 0xa3, 0x30, 0x32, 0xfa, 0x0d, 0x71, 0xc2, 0xd6, 0xeb, 0xc2,
	0x7d, 0xc8, 0x01, 0xed, 0x47, 0x05, 0xaa, 0x27, 0x0a, 0xfa, 0x2e, 0x75, 0x7c, 0x82, 0x1e, 0x41,
	0x3e, 0xde, 0xa0, 0xaa, 0xd4, 0x37, 0xcf, 0xbf, 0xff, 0x65, 0x3c, 0x7a, 0x0f, 0xb6, 0x1d, 0xf2,
	0x82, 0x19, 0x09, 0x2a, 0xa1, 0x42, 0x4b, 0x1c, 0xee, 0xc7, 0x74, 0x6e, 0x43, 0xb5, 0x43, 0x6c,
	0x72, 0xa6, 0x06, 0x69, 0xf7, 0xe1, 0xea, 0x01, 0xb6, 0xed, 0xb3, 0xf4, 0x12, 0xc1, 0xd6, 0x18,
	0x33, 0x2c, 0x8b, 0x8a, 0x6f, 0x6d, 0x02, 0xd7, 0x56, 0xc3, 0xe5, 0xc6, 0xdf, 0x86, 0x22, 0x79,
	0x41, 0xcc, 0x80, 0x83, 0x86, 0x35, 0x96, 0x79, 0x0a, 0x31, 0xd6, 0x1b, 0xa3, 0xeb, 0x90, 0xf5,
	0x88, 0x1f, 0xd8, 0x4c, 0x26, 0x94, 0x2b, 0x74, 0x0d, 0x32, 0xc4, 0xf3, 0xa8, 0x27, 0xcf, 0x39,
	0x5c, 0x34, 0x30, 0x5c, 0x5d, 0x33, 0x93, 0xd1, 0x75, 0x40, 0x83, 0xe1, 0xde, 0xf0, 0x78, 0x60,
	0x1c, 0x3f, 0x19, 0xf4, 0xbb, 0x07, 0xbd, 0xc3, 0x5e, 0xb7, 0x53, 0xb9, 0x84, 0xf2, 0x90, 0xd1,
	0xbb, 0x7b, 0x9d, 0x67, 0x15, 0x05, 0x01, 0x64, 0x0f, 0xf7, 0x7a, 0x47, 0xdd, 0x4e, 0x65, 0x03,
	0x95, 0x20, 0xdf, 0xe9, 0xf6, 0x8f, 0x9e, 0x3e, 0xeb, 0x3d, 0x79, 0x50, 0xd9, 0x44, 0x45, 0xc8,
	0x75, 0xba, 0x47, 0xdd, 0x21, 0x5f, 0x6d, 0xb5, 0x7f, 0xcf, 0x41, 0x75, 0xa5, 0x86, 0x3f, 0x08,
	0xc7, 0x23, 0xfa, 0x55, 0x81, 0xd2, 0x4a, 0x83, 0xd1, 0x6e, 0x5a, 0x17, 0xd7, 0x09, 0xb0, 0xf6,
	0xe9, 0x39, 0xa3, 0xc2, 0xc3, 0xd4, 0xee, 0xfd, 0xf0, 0xc7, 0xdf, 0x3f, 0x6d, 0x7c, 0x86, 0x76,
	0xe3, 0x77, 0xca, 0x77, 0x91, 0x6e, 0xef, 0xbb, 0x1e, 0xfd, 0x9a, 0x98, 0xcc, 0x6f, 0x35, 0x5a,
	0x11, 0xe6, 0xb7, 0x1a, 0xaf, 0x96, 0x6f, 0x1b, 0xf4, 0xb3, 0x02, 0x85, 0xc4, 0xc5, 0x41, 0xed,
	0x34, 0x12, 0xa7, 0x6f, 0x59, 0xed, 0x7c, 0xa2, 0xd5, 0xee, 0x0a, 0xc2, 0xbb, 0xa8, 0xbd, 0x24,
	0xcc, 0x15, 0xf4, 0x2f, 0x64, 0x13, 0xef, 0xb0, 0xc6, 0x2b, 0xf4, 0x8b, 0x02, 0xe5, 0xd5, 0xf9,
	0x86, 0x52, 0x8f, 0x6d, 0xed, 0x3c, 0xac, 0xdd, 0x8c, 0xc2, 0x12, 0x8f, 0xb7, 0x66, 0xfc, 0x6f,
	0xd6, 0x0e, 0x05, 0xc9, 0xaf, 0xb4, 0x0b, 0x9d, 0xea, 0xdd, 0x78, 0x2e, 0xf1, 0xf3, 0x2d, 0xaf,
	0x0e, 0xc4, 0x74, 0xc2, 0x6b, 0x07, 0x68, 0x1a, 0xe1, 0x8e, 0x20, 0xfc, 0x65, 0xed, 0x02, 0xa7,
	0x9a, 0xa0, 0xfb, 0x5a, 0x81, 0xf2, 0xea, 0x78, 0x48, 0xa7, 0xbb, 0x76, 0x9c, 0xa4, 0xd1, 0x95,
	0x22, 0x68, 0x5c, 0x44, 0x04, 0xbf, 0x29, 0x50, 0x4c, 0xce, 0x15, 0x94, 0xfe, 0x60, 0x3b, 0x3d,
	0xc4, 0x6a, 0xbb, 0xe7, 0x0b, 0x92, 0xb7, 0x6d, 0x5f, 0xf0, 0xbe, 0xa7, 0x7d, 0x7e, 0x81, 0x63,
	0x36, 0xb1, 0x6d, 0xdf, 0x55, 0x1a, 0xfb, 0xf7, 0x40, 0x33, 0xe9, 0x2c, 0xa5, 0xfc, 0x7e, 0x39,
	0xbe, 0xea, 0x7d, 0xfe, 0xb4, 0xea, 0x2b, 0xaf, 0x37, 0x36, 0x1f, 0x1c, 0x1c, 0x8e, 0xb2, 0xe2,
	0xa5, 0x75, 0xe7, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xa4, 0x22, 0xc1, 0xe5, 0x33, 0x0d, 0x00,
	0x00,
}