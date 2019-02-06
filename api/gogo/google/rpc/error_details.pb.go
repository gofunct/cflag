// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: google/rpc/error_details.proto

package google_rpc

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/gogo/protobuf/types"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Describes when the clients can retry a failed request. Clients could ignore
// the recommendation here or retry when this information is missing from error
// responses.
//
// It's always recommended that clients should use exponential backoff when
// retrying.
//
// Clients should wait until `retry_delay` amount of time has passed since
// receiving the error response before retrying.  If retrying requests also
// fail, clients should use an exponential backoff scheme to gradually increase
// the delay between retries based on `retry_delay`, until either a maximum
// number of retires have been reached or a maximum retry delay cap has been
// reached.
type RetryInfo struct {
	// Clients should wait at least this long between retrying the same request.
	RetryDelay *google_protobuf.Duration `protobuf:"bytes,1,opt,name=retry_delay,json=retryDelay" json:"retry_delay,omitempty"`
}

func (m *RetryInfo) Reset()                    { *m = RetryInfo{} }
func (m *RetryInfo) String() string            { return proto.CompactTextString(m) }
func (*RetryInfo) ProtoMessage()               {}
func (*RetryInfo) Descriptor() ([]byte, []int) { return fileDescriptorErrorDetails, []int{0} }

func (m *RetryInfo) GetRetryDelay() *google_protobuf.Duration {
	if m != nil {
		return m.RetryDelay
	}
	return nil
}

// Describes additional debugging info.
type DebugInfo struct {
	// The stack trace entries indicating where the error occurred.
	StackEntries []string `protobuf:"bytes,1,rep,name=stack_entries,json=stackEntries" json:"stack_entries,omitempty"`
	// Additional debugging information provided by the server.
	Detail string `protobuf:"bytes,2,opt,name=detail,proto3" json:"detail,omitempty"`
}

func (m *DebugInfo) Reset()                    { *m = DebugInfo{} }
func (m *DebugInfo) String() string            { return proto.CompactTextString(m) }
func (*DebugInfo) ProtoMessage()               {}
func (*DebugInfo) Descriptor() ([]byte, []int) { return fileDescriptorErrorDetails, []int{1} }

func (m *DebugInfo) GetStackEntries() []string {
	if m != nil {
		return m.StackEntries
	}
	return nil
}

func (m *DebugInfo) GetDetail() string {
	if m != nil {
		return m.Detail
	}
	return ""
}

// Describes how a quota check failed.
//
// For example if a daily limit was exceeded for the calling project,
// a service could respond with a QuotaFailure detail containing the project
// id and the description of the quota limit that was exceeded.  If the
// calling project hasn't enabled the service in the developer console, then
// a service could respond with the project id and set `service_disabled`
// to true.
//
// Also see RetryDetail and Help types for other details about handling a
// quota failure.
type QuotaFailure struct {
	// Describes all quota violations.
	Violations []*QuotaFailure_Violation `protobuf:"bytes,1,rep,name=violations" json:"violations,omitempty"`
}

func (m *QuotaFailure) Reset()                    { *m = QuotaFailure{} }
func (m *QuotaFailure) String() string            { return proto.CompactTextString(m) }
func (*QuotaFailure) ProtoMessage()               {}
func (*QuotaFailure) Descriptor() ([]byte, []int) { return fileDescriptorErrorDetails, []int{2} }

func (m *QuotaFailure) GetViolations() []*QuotaFailure_Violation {
	if m != nil {
		return m.Violations
	}
	return nil
}

// A message type used to describe a single quota violation.  For example, a
// daily quota or a custom quota that was exceeded.
type QuotaFailure_Violation struct {
	// The subject on which the quota check failed.
	// For example, "clientip:<ip address of client>" or "project:<Google
	// developer project id>".
	Subject string `protobuf:"bytes,1,opt,name=subject,proto3" json:"subject,omitempty"`
	// A description of how the quota check failed. Clients can use this
	// description to find more about the quota configuration in the service's
	// public documentation, or find the relevant quota limit to adjust through
	// developer console.
	//
	// For example: "Service disabled" or "Daily Limit for read operations
	// exceeded".
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
}

func (m *QuotaFailure_Violation) Reset()         { *m = QuotaFailure_Violation{} }
func (m *QuotaFailure_Violation) String() string { return proto.CompactTextString(m) }
func (*QuotaFailure_Violation) ProtoMessage()    {}
func (*QuotaFailure_Violation) Descriptor() ([]byte, []int) {
	return fileDescriptorErrorDetails, []int{2, 0}
}

func (m *QuotaFailure_Violation) GetSubject() string {
	if m != nil {
		return m.Subject
	}
	return ""
}

func (m *QuotaFailure_Violation) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

// Describes what preconditions have failed.
//
// For example, if an RPC failed because it required the Terms of Service to be
// acknowledged, it could list the terms of service violation in the
// PreconditionFailure message.
type PreconditionFailure struct {
	// Describes all precondition violations.
	Violations []*PreconditionFailure_Violation `protobuf:"bytes,1,rep,name=violations" json:"violations,omitempty"`
}

func (m *PreconditionFailure) Reset()                    { *m = PreconditionFailure{} }
func (m *PreconditionFailure) String() string            { return proto.CompactTextString(m) }
func (*PreconditionFailure) ProtoMessage()               {}
func (*PreconditionFailure) Descriptor() ([]byte, []int) { return fileDescriptorErrorDetails, []int{3} }

func (m *PreconditionFailure) GetViolations() []*PreconditionFailure_Violation {
	if m != nil {
		return m.Violations
	}
	return nil
}

// A message type used to describe a single precondition failure.
type PreconditionFailure_Violation struct {
	// The type of PreconditionFailure. We recommend using a service-specific
	// enum type to define the supported precondition violation types. For
	// example, "TOS" for "Terms of Service violation".
	Type string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	// The subject, relative to the type, that failed.
	// For example, "google.com/cloud" relative to the "TOS" type would
	// indicate which terms of service is being referenced.
	Subject string `protobuf:"bytes,2,opt,name=subject,proto3" json:"subject,omitempty"`
	// A description of how the precondition failed. Developers can use this
	// description to understand how to fix the failure.
	//
	// For example: "Terms of service not accepted".
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
}

func (m *PreconditionFailure_Violation) Reset()         { *m = PreconditionFailure_Violation{} }
func (m *PreconditionFailure_Violation) String() string { return proto.CompactTextString(m) }
func (*PreconditionFailure_Violation) ProtoMessage()    {}
func (*PreconditionFailure_Violation) Descriptor() ([]byte, []int) {
	return fileDescriptorErrorDetails, []int{3, 0}
}

func (m *PreconditionFailure_Violation) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *PreconditionFailure_Violation) GetSubject() string {
	if m != nil {
		return m.Subject
	}
	return ""
}

func (m *PreconditionFailure_Violation) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

// Describes violations in a client request. This error type focuses on the
// syntactic aspects of the request.
type BadRequest struct {
	// Describes all violations in a client request.
	FieldViolations []*BadRequest_FieldViolation `protobuf:"bytes,1,rep,name=field_violations,json=fieldViolations" json:"field_violations,omitempty"`
}

func (m *BadRequest) Reset()                    { *m = BadRequest{} }
func (m *BadRequest) String() string            { return proto.CompactTextString(m) }
func (*BadRequest) ProtoMessage()               {}
func (*BadRequest) Descriptor() ([]byte, []int) { return fileDescriptorErrorDetails, []int{4} }

func (m *BadRequest) GetFieldViolations() []*BadRequest_FieldViolation {
	if m != nil {
		return m.FieldViolations
	}
	return nil
}

// A message type used to describe a single bad request field.
type BadRequest_FieldViolation struct {
	// A path leading to a field in the request body. The value will be a
	// sequence of dot-separated identifiers that identify a protocol buffer
	// field. E.g., "field_violations.field" would identify this field.
	Field string `protobuf:"bytes,1,opt,name=field,proto3" json:"field,omitempty"`
	// A description of why the request element is bad.
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
}

func (m *BadRequest_FieldViolation) Reset()         { *m = BadRequest_FieldViolation{} }
func (m *BadRequest_FieldViolation) String() string { return proto.CompactTextString(m) }
func (*BadRequest_FieldViolation) ProtoMessage()    {}
func (*BadRequest_FieldViolation) Descriptor() ([]byte, []int) {
	return fileDescriptorErrorDetails, []int{4, 0}
}

func (m *BadRequest_FieldViolation) GetField() string {
	if m != nil {
		return m.Field
	}
	return ""
}

func (m *BadRequest_FieldViolation) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

// Contains metadata about the request that clients can attach when filing a bug
// or providing other forms of feedback.
type RequestInfo struct {
	// An opaque string that should only be interpreted by the service generating
	// it. For example, it can be used to identify requests in the service's logs.
	RequestId string `protobuf:"bytes,1,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	// Any data that was used to serve this request. For example, an encrypted
	// stack trace that can be sent back to the service provider for debugging.
	ServingData string `protobuf:"bytes,2,opt,name=serving_data,json=servingData,proto3" json:"serving_data,omitempty"`
}

func (m *RequestInfo) Reset()                    { *m = RequestInfo{} }
func (m *RequestInfo) String() string            { return proto.CompactTextString(m) }
func (*RequestInfo) ProtoMessage()               {}
func (*RequestInfo) Descriptor() ([]byte, []int) { return fileDescriptorErrorDetails, []int{5} }

func (m *RequestInfo) GetRequestId() string {
	if m != nil {
		return m.RequestId
	}
	return ""
}

func (m *RequestInfo) GetServingData() string {
	if m != nil {
		return m.ServingData
	}
	return ""
}

// Describes the resource that is being accessed.
type ResourceInfo struct {
	// A name for the type of resource being accessed, e.g. "sql table",
	// "cloud storage bucket", "file", "Google calendar"; or the type URL
	// of the resource: e.g. "type.googleapis.com/google.pubsub.v1.Topic".
	ResourceType string `protobuf:"bytes,1,opt,name=resource_type,json=resourceType,proto3" json:"resource_type,omitempty"`
	// The name of the resource being accessed.  For example, a shared calendar
	// name: "example.com_4fghdhgsrgh@group.calendar.google.com", if the current
	// error is [google.rpc.Code.PERMISSION_DENIED][google.rpc.Code.PERMISSION_DENIED].
	ResourceName string `protobuf:"bytes,2,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// The owner of the resource (optional).
	// For example, "user:<owner email>" or "project:<Google developer project
	// id>".
	Owner string `protobuf:"bytes,3,opt,name=owner,proto3" json:"owner,omitempty"`
	// Describes what error is encountered when accessing this resource.
	// For example, updating a cloud project may require the `writer` permission
	// on the developer console project.
	Description string `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
}

func (m *ResourceInfo) Reset()                    { *m = ResourceInfo{} }
func (m *ResourceInfo) String() string            { return proto.CompactTextString(m) }
func (*ResourceInfo) ProtoMessage()               {}
func (*ResourceInfo) Descriptor() ([]byte, []int) { return fileDescriptorErrorDetails, []int{6} }

func (m *ResourceInfo) GetResourceType() string {
	if m != nil {
		return m.ResourceType
	}
	return ""
}

func (m *ResourceInfo) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *ResourceInfo) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *ResourceInfo) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

// Provides links to documentation or for performing an out of band action.
//
// For example, if a quota check failed with an error indicating the calling
// project hasn't enabled the accessed service, this can contain a URL pointing
// directly to the right place in the developer console to flip the bit.
type Help struct {
	// URL(s) pointing to additional information on handling the current error.
	Links []*Help_Link `protobuf:"bytes,1,rep,name=links" json:"links,omitempty"`
}

func (m *Help) Reset()                    { *m = Help{} }
func (m *Help) String() string            { return proto.CompactTextString(m) }
func (*Help) ProtoMessage()               {}
func (*Help) Descriptor() ([]byte, []int) { return fileDescriptorErrorDetails, []int{7} }

func (m *Help) GetLinks() []*Help_Link {
	if m != nil {
		return m.Links
	}
	return nil
}

// Describes a URL link.
type Help_Link struct {
	// Describes what the link offers.
	Description string `protobuf:"bytes,1,opt,name=description,proto3" json:"description,omitempty"`
	// The URL of the link.
	Url string `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
}

func (m *Help_Link) Reset()                    { *m = Help_Link{} }
func (m *Help_Link) String() string            { return proto.CompactTextString(m) }
func (*Help_Link) ProtoMessage()               {}
func (*Help_Link) Descriptor() ([]byte, []int) { return fileDescriptorErrorDetails, []int{7, 0} }

func (m *Help_Link) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Help_Link) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

// Provides a localized error message that is safe to return to the user
// which can be attached to an RPC error.
type LocalizedMessage struct {
	// The locale used following the specification defined at
	// http://www.rfc-editor.org/rfc/bcp/bcp47.txt.
	// Examples are: "en-US", "fr-CH", "es-MX"
	Locale string `protobuf:"bytes,1,opt,name=locale,proto3" json:"locale,omitempty"`
	// The localized error message in the above locale.
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (m *LocalizedMessage) Reset()                    { *m = LocalizedMessage{} }
func (m *LocalizedMessage) String() string            { return proto.CompactTextString(m) }
func (*LocalizedMessage) ProtoMessage()               {}
func (*LocalizedMessage) Descriptor() ([]byte, []int) { return fileDescriptorErrorDetails, []int{8} }

func (m *LocalizedMessage) GetLocale() string {
	if m != nil {
		return m.Locale
	}
	return ""
}

func (m *LocalizedMessage) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*RetryInfo)(nil), "google.rpc.RetryInfo")
	proto.RegisterType((*DebugInfo)(nil), "google.rpc.DebugInfo")
	proto.RegisterType((*QuotaFailure)(nil), "google.rpc.QuotaFailure")
	proto.RegisterType((*QuotaFailure_Violation)(nil), "google.rpc.QuotaFailure.Violation")
	proto.RegisterType((*PreconditionFailure)(nil), "google.rpc.PreconditionFailure")
	proto.RegisterType((*PreconditionFailure_Violation)(nil), "google.rpc.PreconditionFailure.Violation")
	proto.RegisterType((*BadRequest)(nil), "google.rpc.BadRequest")
	proto.RegisterType((*BadRequest_FieldViolation)(nil), "google.rpc.BadRequest.FieldViolation")
	proto.RegisterType((*RequestInfo)(nil), "google.rpc.RequestInfo")
	proto.RegisterType((*ResourceInfo)(nil), "google.rpc.ResourceInfo")
	proto.RegisterType((*Help)(nil), "google.rpc.Help")
	proto.RegisterType((*Help_Link)(nil), "google.rpc.Help.Link")
	proto.RegisterType((*LocalizedMessage)(nil), "google.rpc.LocalizedMessage")
}

func init() { proto.RegisterFile("google/rpc/error_details.proto", fileDescriptorErrorDetails) }

var fileDescriptorErrorDetails = []byte{
	// 564 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0x95, 0x93, 0xb4, 0xc8, 0x93, 0x50, 0xc2, 0xf2, 0xa1, 0x10, 0x09, 0x14, 0x8c, 0x90, 0x82,
	0x2a, 0x39, 0x52, 0xb9, 0xf5, 0x18, 0xdc, 0x36, 0x91, 0x0a, 0x04, 0x0b, 0x71, 0xe1, 0x60, 0x6d,
	0xec, 0x49, 0xb4, 0xc4, 0xf1, 0x9a, 0xf5, 0xba, 0x28, 0xfc, 0x0a, 0xee, 0xdc, 0x38, 0xf1, 0x27,
	0xf8, 0x6f, 0x68, 0xed, 0xdd, 0x66, 0xdb, 0x14, 0xd4, 0x9b, 0xdf, 0xdb, 0xb7, 0xcf, 0x6f, 0x46,
	0xb3, 0x03, 0xcf, 0x96, 0x9c, 0x2f, 0x53, 0x1c, 0x89, 0x3c, 0x1e, 0xa1, 0x10, 0x5c, 0x44, 0x09,
	0x4a, 0xca, 0xd2, 0xc2, 0xcf, 0x05, 0x97, 0x9c, 0x40, 0x7d, 0xee, 0x8b, 0x3c, 0xee, 0x1b, 0x6d,
	0x75, 0x32, 0x2f, 0x17, 0xa3, 0xa4, 0x14, 0x54, 0x32, 0x9e, 0xd5, 0x5a, 0xef, 0x0c, 0xdc, 0x10,
	0xa5, 0xd8, 0x4c, 0xb3, 0x05, 0x27, 0xc7, 0xd0, 0x16, 0x0a, 0x44, 0x09, 0xa6, 0x74, 0xd3, 0x73,
	0x06, 0xce, 0xb0, 0x7d, 0xf4, 0xc4, 0xd7, 0x76, 0xc6, 0xc2, 0x0f, 0xb4, 0x45, 0x08, 0x95, 0x3a,
	0x50, 0x62, 0x6f, 0x02, 0x6e, 0x80, 0xf3, 0x72, 0x59, 0x19, 0xbd, 0x80, 0xbb, 0x85, 0xa4, 0xf1,
	0x2a, 0xc2, 0x4c, 0x0a, 0x86, 0x45, 0xcf, 0x19, 0x34, 0x87, 0x6e, 0xd8, 0xa9, 0xc8, 0x93, 0x9a,
	0x23, 0x8f, 0x61, 0xbf, 0xce, 0xdd, 0x6b, 0x0c, 0x9c, 0xa1, 0x1b, 0x6a, 0xe4, 0xfd, 0x74, 0xa0,
	0xf3, 0xa1, 0xe4, 0x92, 0x9e, 0x52, 0x96, 0x96, 0x02, 0xc9, 0x18, 0xe0, 0x82, 0xf1, 0xb4, 0xfa,
	0x67, 0x6d, 0xd5, 0x3e, 0xf2, 0xfc, 0x6d, 0x91, 0xbe, 0xad, 0xf6, 0x3f, 0x19, 0x69, 0x68, 0xdd,
	0xea, 0x9f, 0x81, 0x7b, 0x79, 0x40, 0x7a, 0x70, 0xa7, 0x28, 0xe7, 0x5f, 0x30, 0x96, 0x55, 0x8d,
	0x6e, 0x68, 0x20, 0x19, 0x40, 0x3b, 0xc1, 0x22, 0x16, 0x2c, 0x57, 0x42, 0x1d, 0xcc, 0xa6, 0xbc,
	0x3f, 0x0e, 0x3c, 0x98, 0x09, 0x8c, 0x79, 0x96, 0x30, 0x45, 0x98, 0x90, 0xd3, 0x1b, 0x42, 0xbe,
	0xb2, 0x43, 0xde, 0x70, 0xe9, 0x1f, 0x59, 0x3f, 0xdb, 0x59, 0x09, 0xb4, 0xe4, 0x26, 0x47, 0x1d,
	0xb4, 0xfa, 0xb6, 0xf3, 0x37, 0xfe, 0x9b, 0xbf, 0xb9, 0x9b, 0xff, 0xb7, 0x03, 0x30, 0xa6, 0x49,
	0x88, 0x5f, 0x4b, 0x2c, 0x24, 0x99, 0x41, 0x77, 0xc1, 0x30, 0x4d, 0xa2, 0x9d, 0xf0, 0x2f, 0xed,
	0xf0, 0xdb, 0x1b, 0xfe, 0xa9, 0x92, 0x6f, 0x83, 0xdf, 0x5b, 0x5c, 0xc1, 0x45, 0x7f, 0x02, 0x07,
	0x57, 0x25, 0xe4, 0x21, 0xec, 0x55, 0x22, 0x5d, 0x43, 0x0d, 0x6e, 0xd1, 0xea, 0xf7, 0xd0, 0xd6,
	0x3f, 0xad, 0x86, 0xea, 0x29, 0x80, 0xa8, 0x61, 0xc4, 0x8c, 0x97, 0xab, 0x99, 0x69, 0x42, 0x9e,
	0x43, 0xa7, 0x40, 0x71, 0xc1, 0xb2, 0x65, 0x94, 0x50, 0x49, 0x8d, 0xa1, 0xe6, 0x02, 0x2a, 0xa9,
	0xf7, 0xc3, 0x81, 0x4e, 0x88, 0x05, 0x2f, 0x45, 0x8c, 0x66, 0x4e, 0x85, 0xc6, 0x91, 0xd5, 0xe5,
	0x8e, 0x21, 0x3f, 0xaa, 0x6e, 0xdb, 0xa2, 0x8c, 0xae, 0x51, 0x3b, 0x5f, 0x8a, 0xde, 0xd1, 0x35,
	0xaa, 0x1a, 0xf9, 0xb7, 0x0c, 0x85, 0x6e, 0x79, 0x0d, 0xae, 0xd7, 0xd8, 0xda, 0xad, 0x91, 0x43,
	0x6b, 0x82, 0x69, 0x4e, 0x0e, 0x61, 0x2f, 0x65, 0xd9, 0xca, 0x34, 0xff, 0x91, 0xdd, 0x7c, 0x25,
	0xf0, 0xcf, 0x59, 0xb6, 0x0a, 0x6b, 0x4d, 0xff, 0x18, 0x5a, 0x0a, 0x5e, 0xb7, 0x77, 0x76, 0xec,
	0x49, 0x17, 0x9a, 0xa5, 0x30, 0x0f, 0x4c, 0x7d, 0x7a, 0x01, 0x74, 0xcf, 0x79, 0x4c, 0x53, 0xf6,
	0x1d, 0x93, 0xb7, 0x58, 0x14, 0x74, 0x89, 0xea, 0x25, 0xa6, 0x8a, 0x33, 0xf5, 0x6b, 0xa4, 0xe6,
	0x6c, 0x5d, 0x4b, 0xcc, 0x9c, 0x69, 0x38, 0x3e, 0x84, 0x83, 0x98, 0xaf, 0xad, 0x90, 0xe3, 0xfb,
	0x27, 0x6a, 0x13, 0x05, 0xf5, 0x22, 0x9a, 0xa9, 0x55, 0x31, 0x73, 0x7e, 0x35, 0x9a, 0xe1, 0xec,
	0xcd, 0x7c, 0xbf, 0xda, 0x1c, 0xaf, 0xff, 0x06, 0x00, 0x00, 0xff, 0xff, 0xe8, 0xd5, 0x62, 0xee,
	0xb8, 0x04, 0x00, 0x00,
}