// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/audit/audit_log.proto

package google_cloud_audit

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "go.pedge.io/pb/go/google/api"
import google_protobuf1 "github.com/golang/protobuf/ptypes/any"
import google_protobuf2 "github.com/golang/protobuf/ptypes/struct"
import google_rpc "go.pedge.io/pb/go/google/rpc"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Common audit log format for Google Cloud Platform API operations.
type AuditLog struct {
	// The name of the API service performing the operation. For example,
	// `"datastore.googleapis.com"`.
	ServiceName string `protobuf:"bytes,7,opt,name=service_name,json=serviceName" json:"service_name,omitempty"`
	// The name of the service method or operation.
	// For API calls, this should be the name of the API method.
	// For example,
	//
	//     "google.datastore.v1.Datastore.RunQuery"
	//     "google.logging.v1.LoggingService.DeleteLog"
	MethodName string `protobuf:"bytes,8,opt,name=method_name,json=methodName" json:"method_name,omitempty"`
	// The resource or collection that is the target of the operation.
	// The name is a scheme-less URI, not including the API service name.
	// For example:
	//
	//     "shelves/SHELF_ID/books"
	//     "shelves/SHELF_ID/books/BOOK_ID"
	ResourceName string `protobuf:"bytes,11,opt,name=resource_name,json=resourceName" json:"resource_name,omitempty"`
	// The number of items returned from a List or Query API method,
	// if applicable.
	NumResponseItems int64 `protobuf:"varint,12,opt,name=num_response_items,json=numResponseItems" json:"num_response_items,omitempty"`
	// The status of the overall operation.
	Status *google_rpc.Status `protobuf:"bytes,2,opt,name=status" json:"status,omitempty"`
	// Authentication information.
	AuthenticationInfo *AuthenticationInfo `protobuf:"bytes,3,opt,name=authentication_info,json=authenticationInfo" json:"authentication_info,omitempty"`
	// Authorization information. If there are multiple
	// resources or permissions involved, then there is
	// one AuthorizationInfo element for each {resource, permission} tuple.
	AuthorizationInfo []*AuthorizationInfo `protobuf:"bytes,9,rep,name=authorization_info,json=authorizationInfo" json:"authorization_info,omitempty"`
	// Metadata about the operation.
	RequestMetadata *RequestMetadata `protobuf:"bytes,4,opt,name=request_metadata,json=requestMetadata" json:"request_metadata,omitempty"`
	// The operation request. This may not include all request parameters,
	// such as those that are too large, privacy-sensitive, or duplicated
	// elsewhere in the log record.
	// It should never include user-generated data, such as file contents.
	// When the JSON object represented here has a proto equivalent, the proto
	// name will be indicated in the `@type` property.
	Request *google_protobuf2.Struct `protobuf:"bytes,16,opt,name=request" json:"request,omitempty"`
	// The operation response. This may not include all response elements,
	// such as those that are too large, privacy-sensitive, or duplicated
	// elsewhere in the log record.
	// It should never include user-generated data, such as file contents.
	// When the JSON object represented here has a proto equivalent, the proto
	// name will be indicated in the `@type` property.
	Response *google_protobuf2.Struct `protobuf:"bytes,17,opt,name=response" json:"response,omitempty"`
	// Other service-specific data about the request, response, and other
	// activities.
	ServiceData *google_protobuf1.Any `protobuf:"bytes,15,opt,name=service_data,json=serviceData" json:"service_data,omitempty"`
}

func (m *AuditLog) Reset()                    { *m = AuditLog{} }
func (m *AuditLog) String() string            { return proto.CompactTextString(m) }
func (*AuditLog) ProtoMessage()               {}
func (*AuditLog) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *AuditLog) GetServiceName() string {
	if m != nil {
		return m.ServiceName
	}
	return ""
}

func (m *AuditLog) GetMethodName() string {
	if m != nil {
		return m.MethodName
	}
	return ""
}

func (m *AuditLog) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *AuditLog) GetNumResponseItems() int64 {
	if m != nil {
		return m.NumResponseItems
	}
	return 0
}

func (m *AuditLog) GetStatus() *google_rpc.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *AuditLog) GetAuthenticationInfo() *AuthenticationInfo {
	if m != nil {
		return m.AuthenticationInfo
	}
	return nil
}

func (m *AuditLog) GetAuthorizationInfo() []*AuthorizationInfo {
	if m != nil {
		return m.AuthorizationInfo
	}
	return nil
}

func (m *AuditLog) GetRequestMetadata() *RequestMetadata {
	if m != nil {
		return m.RequestMetadata
	}
	return nil
}

func (m *AuditLog) GetRequest() *google_protobuf2.Struct {
	if m != nil {
		return m.Request
	}
	return nil
}

func (m *AuditLog) GetResponse() *google_protobuf2.Struct {
	if m != nil {
		return m.Response
	}
	return nil
}

func (m *AuditLog) GetServiceData() *google_protobuf1.Any {
	if m != nil {
		return m.ServiceData
	}
	return nil
}

// Authentication information for the operation.
type AuthenticationInfo struct {
	// The email address of the authenticated user making the request.
	PrincipalEmail string `protobuf:"bytes,1,opt,name=principal_email,json=principalEmail" json:"principal_email,omitempty"`
}

func (m *AuthenticationInfo) Reset()                    { *m = AuthenticationInfo{} }
func (m *AuthenticationInfo) String() string            { return proto.CompactTextString(m) }
func (*AuthenticationInfo) ProtoMessage()               {}
func (*AuthenticationInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *AuthenticationInfo) GetPrincipalEmail() string {
	if m != nil {
		return m.PrincipalEmail
	}
	return ""
}

// Authorization information for the operation.
type AuthorizationInfo struct {
	// The resource being accessed, as a REST-style string. For example:
	//
	//     bigquery.googlapis.com/projects/PROJECTID/datasets/DATASETID
	Resource string `protobuf:"bytes,1,opt,name=resource" json:"resource,omitempty"`
	// The required IAM permission.
	Permission string `protobuf:"bytes,2,opt,name=permission" json:"permission,omitempty"`
	// Whether or not authorization for `resource` and `permission`
	// was granted.
	Granted bool `protobuf:"varint,3,opt,name=granted" json:"granted,omitempty"`
}

func (m *AuthorizationInfo) Reset()                    { *m = AuthorizationInfo{} }
func (m *AuthorizationInfo) String() string            { return proto.CompactTextString(m) }
func (*AuthorizationInfo) ProtoMessage()               {}
func (*AuthorizationInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *AuthorizationInfo) GetResource() string {
	if m != nil {
		return m.Resource
	}
	return ""
}

func (m *AuthorizationInfo) GetPermission() string {
	if m != nil {
		return m.Permission
	}
	return ""
}

func (m *AuthorizationInfo) GetGranted() bool {
	if m != nil {
		return m.Granted
	}
	return false
}

// Metadata about the request.
type RequestMetadata struct {
	// The IP address of the caller.
	CallerIp string `protobuf:"bytes,1,opt,name=caller_ip,json=callerIp" json:"caller_ip,omitempty"`
	// The user agent of the caller.
	// This information is not authenticated and should be treated accordingly.
	// For example:
	//
	// +   `google-api-python-client/1.4.0`:
	//     The request was made by the Google API client for Python.
	// +   `Cloud SDK Command Line Tool apitools-client/1.0 gcloud/0.9.62`:
	//     The request was made by the Google Cloud SDK CLI (gcloud).
	// +   `AppEngine-Google; (+http://code.google.com/appengine; appid: s~my-project`:
	//     The request was made from the `my-project` App Engine app.
	CallerSuppliedUserAgent string `protobuf:"bytes,2,opt,name=caller_supplied_user_agent,json=callerSuppliedUserAgent" json:"caller_supplied_user_agent,omitempty"`
}

func (m *RequestMetadata) Reset()                    { *m = RequestMetadata{} }
func (m *RequestMetadata) String() string            { return proto.CompactTextString(m) }
func (*RequestMetadata) ProtoMessage()               {}
func (*RequestMetadata) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *RequestMetadata) GetCallerIp() string {
	if m != nil {
		return m.CallerIp
	}
	return ""
}

func (m *RequestMetadata) GetCallerSuppliedUserAgent() string {
	if m != nil {
		return m.CallerSuppliedUserAgent
	}
	return ""
}

func init() {
	proto.RegisterType((*AuditLog)(nil), "google.cloud.audit.AuditLog")
	proto.RegisterType((*AuthenticationInfo)(nil), "google.cloud.audit.AuthenticationInfo")
	proto.RegisterType((*AuthorizationInfo)(nil), "google.cloud.audit.AuthorizationInfo")
	proto.RegisterType((*RequestMetadata)(nil), "google.cloud.audit.RequestMetadata")
}

func init() { proto.RegisterFile("google/cloud/audit/audit_log.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 550 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x93, 0x5f, 0x6f, 0xd3, 0x30,
	0x14, 0xc5, 0x55, 0x3a, 0xad, 0xed, 0x6d, 0x47, 0x5b, 0x83, 0x68, 0x28, 0x13, 0x94, 0x4e, 0x40,
	0x41, 0x28, 0x15, 0xdb, 0x03, 0x0f, 0x88, 0x87, 0x4e, 0xf0, 0x50, 0x09, 0xa6, 0x29, 0x05, 0xf1,
	0x18, 0x79, 0x89, 0x9b, 0x59, 0x24, 0xb6, 0xf1, 0x1f, 0xa4, 0xf1, 0x9d, 0xf9, 0x0e, 0xa8, 0x37,
	0x4e, 0xd9, 0x9a, 0xc1, 0xcb, 0x24, 0x9f, 0xf3, 0xbb, 0xd7, 0xde, 0xc9, 0x29, 0x4c, 0x33, 0x29,
	0xb3, 0x9c, 0xcd, 0x93, 0x5c, 0xba, 0x74, 0x4e, 0x5d, 0xca, 0x6d, 0xf9, 0x37, 0xce, 0x65, 0x16,
	0x2a, 0x2d, 0xad, 0x24, 0xa4, 0x64, 0x42, 0x64, 0x42, 0x74, 0xc7, 0x87, 0x7e, 0x8e, 0x2a, 0x3e,
	0xa7, 0x42, 0x48, 0x4b, 0x2d, 0x97, 0xc2, 0x94, 0x13, 0xe3, 0x87, 0xde, 0xc5, 0xd3, 0x85, 0x5b,
	0xcf, 0xa9, 0xb8, 0xf2, 0xd6, 0xe1, 0xae, 0x65, 0xac, 0x76, 0x89, 0xf5, 0xee, 0xc8, 0xbb, 0x5a,
	0x25, 0x73, 0x63, 0xa9, 0x75, 0x7e, 0xe3, 0xf4, 0xf7, 0x1e, 0xb4, 0x17, 0x9b, 0x9b, 0x3f, 0xc9,
	0x8c, 0x3c, 0x85, 0x9e, 0x61, 0xfa, 0x27, 0x4f, 0x58, 0x2c, 0x68, 0xc1, 0x82, 0xd6, 0xa4, 0x31,
	0xeb, 0x44, 0x5d, 0xaf, 0x9d, 0xd1, 0x82, 0x91, 0x27, 0xd0, 0x2d, 0x98, 0xbd, 0x94, 0x69, 0x49,
	0xb4, 0x91, 0x80, 0x52, 0x42, 0xe0, 0x08, 0x0e, 0x34, 0x33, 0xd2, 0xe9, 0x6a, 0x49, 0x17, 0x91,
	0x5e, 0x25, 0x22, 0xf4, 0x1a, 0x88, 0x70, 0x45, 0xac, 0x99, 0x51, 0x52, 0x18, 0x16, 0x73, 0xcb,
	0x0a, 0x13, 0xf4, 0x26, 0x8d, 0x59, 0x33, 0x1a, 0x08, 0x57, 0x44, 0xde, 0x58, 0x6e, 0x74, 0xf2,
	0x0a, 0xf6, 0xcb, 0x37, 0x07, 0x77, 0x26, 0x8d, 0x59, 0xf7, 0x98, 0x84, 0x3e, 0x38, 0xad, 0x92,
	0x70, 0x85, 0x4e, 0xe4, 0x09, 0xf2, 0x0d, 0xee, 0x51, 0x67, 0x2f, 0x99, 0xb0, 0x3c, 0xc1, 0xe8,
	0x62, 0x2e, 0xd6, 0x32, 0x68, 0xe2, 0xe0, 0xf3, 0xb0, 0x9e, 0x78, 0xb8, 0xb8, 0x81, 0x2f, 0xc5,
	0x5a, 0x46, 0x84, 0xd6, 0x34, 0xf2, 0x05, 0x50, 0x95, 0x9a, 0xff, 0xba, 0xb6, 0xb7, 0x33, 0x69,
	0xce, 0xba, 0xc7, 0xcf, 0xfe, 0xb5, 0x77, 0x4b, 0xe3, 0xda, 0x21, 0xdd, 0x95, 0xc8, 0x19, 0x0c,
	0x34, 0xfb, 0xe1, 0x98, 0xb1, 0x71, 0xc1, 0x2c, 0x4d, 0xa9, 0xa5, 0xc1, 0x1e, 0xbe, 0xf5, 0xe8,
	0xb6, 0x9d, 0x51, 0xc9, 0x7e, 0xf6, 0x68, 0xd4, 0xd7, 0x37, 0x05, 0xf2, 0x06, 0x5a, 0x5e, 0x0a,
	0x06, 0xb8, 0x66, 0x54, 0xad, 0xa9, 0x7a, 0x11, 0xae, 0xb0, 0x17, 0x51, 0xc5, 0x91, 0x13, 0x68,
	0x57, 0xdf, 0x21, 0x18, 0xfe, 0x7f, 0x66, 0x0b, 0x92, 0xb7, 0x7f, 0x9b, 0x82, 0x6f, 0xee, 0xe3,
	0xe0, 0xfd, 0xda, 0xe0, 0x42, 0x5c, 0x6d, 0xfb, 0xf3, 0x81, 0x5a, 0x3a, 0x7d, 0x0f, 0xa4, 0x1e,
	0x38, 0x79, 0x01, 0x7d, 0xa5, 0xb9, 0x48, 0xb8, 0xa2, 0x79, 0xcc, 0x0a, 0xca, 0xf3, 0xa0, 0x81,
	0xb5, 0xb9, 0xbb, 0x95, 0x3f, 0x6e, 0xd4, 0x29, 0x87, 0x61, 0x2d, 0x57, 0x32, 0xc6, 0xff, 0x00,
	0xdb, 0xe5, 0xc7, 0xb6, 0x67, 0xf2, 0x18, 0x40, 0x31, 0x5d, 0x70, 0x63, 0xb8, 0x14, 0xd8, 0x9f,
	0x4e, 0x74, 0x4d, 0x21, 0x01, 0xb4, 0x32, 0x4d, 0x85, 0x65, 0x29, 0x76, 0xa4, 0x1d, 0x55, 0xc7,
	0xe9, 0x77, 0xe8, 0xef, 0xc4, 0x4d, 0x1e, 0x41, 0x27, 0xa1, 0x79, 0xce, 0x74, 0xcc, 0x55, 0x75,
	0x53, 0x29, 0x2c, 0x15, 0x79, 0x07, 0x63, 0x6f, 0x1a, 0xa7, 0x54, 0xce, 0x59, 0x1a, 0x3b, 0xc3,
	0x74, 0x4c, 0x33, 0x26, 0xac, 0xbf, 0x79, 0x54, 0x12, 0x2b, 0x0f, 0x7c, 0x35, 0x4c, 0x2f, 0x36,
	0xf6, 0xe9, 0x4b, 0x78, 0x90, 0xc8, 0xe2, 0x96, 0x4f, 0x7e, 0x7a, 0x50, 0xfd, 0x3a, 0xcf, 0x37,
	0x99, 0x9e, 0x37, 0x2e, 0xf6, 0x31, 0xdc, 0x93, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xb6, 0x54,
	0x25, 0x46, 0x62, 0x04, 0x00, 0x00,
}