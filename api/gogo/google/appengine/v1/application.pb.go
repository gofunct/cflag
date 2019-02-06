// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: google/appengine/v1/application.proto

package google_appengine_v1

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "go.pedge.io/pb/gogo/google/api"
import google_protobuf1 "github.com/gogo/protobuf/types"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// An Application resource contains the top-level configuration of an App
// Engine application.
type Application struct {
	// Full path to the Application resource in the API.
	// Example: `apps/myapp`.
	//
	// @OutputOnly
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Identifier of the Application resource. This identifier is equivalent
	// to the project ID of the Google Cloud Platform project where you want to
	// deploy your application.
	// Example: `myapp`.
	Id string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	// HTTP path dispatch rules for requests to the application that do not
	// explicitly target a service or version. Rules are order-dependent.
	//
	// @OutputOnly
	DispatchRules []*UrlDispatchRule `protobuf:"bytes,3,rep,name=dispatch_rules,json=dispatchRules" json:"dispatch_rules,omitempty"`
	// Google Apps authentication domain that controls which users can access
	// this application.
	//
	// Defaults to open access for any Google Account.
	AuthDomain string `protobuf:"bytes,6,opt,name=auth_domain,json=authDomain,proto3" json:"auth_domain,omitempty"`
	// Location from which this application will be run. Application instances
	// will run out of data centers in the chosen location, which is also where
	// all of the application's end user content is stored.
	//
	// Defaults to `us-central`.
	//
	// Options are:
	//
	// `us-central` - Central US
	//
	// `europe-west` - Western Europe
	//
	// `us-east1` - Eastern US
	LocationId string `protobuf:"bytes,7,opt,name=location_id,json=locationId,proto3" json:"location_id,omitempty"`
	// Google Cloud Storage bucket that can be used for storing files
	// associated with this application. This bucket is associated with the
	// application and can be used by the gcloud deployment commands.
	//
	// @OutputOnly
	CodeBucket string `protobuf:"bytes,8,opt,name=code_bucket,json=codeBucket,proto3" json:"code_bucket,omitempty"`
	// Cookie expiration policy for this application.
	//
	// @OutputOnly
	DefaultCookieExpiration *google_protobuf1.Duration `protobuf:"bytes,9,opt,name=default_cookie_expiration,json=defaultCookieExpiration" json:"default_cookie_expiration,omitempty"`
	// Hostname used to reach this application, as resolved by App Engine.
	//
	// @OutputOnly
	DefaultHostname string `protobuf:"bytes,11,opt,name=default_hostname,json=defaultHostname,proto3" json:"default_hostname,omitempty"`
	// Google Cloud Storage bucket that can be used by this application to store
	// content.
	//
	// @OutputOnly
	DefaultBucket string `protobuf:"bytes,12,opt,name=default_bucket,json=defaultBucket,proto3" json:"default_bucket,omitempty"`
}

func (m *Application) Reset()                    { *m = Application{} }
func (m *Application) String() string            { return proto.CompactTextString(m) }
func (*Application) ProtoMessage()               {}
func (*Application) Descriptor() ([]byte, []int) { return fileDescriptorApplication, []int{0} }

func (m *Application) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Application) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Application) GetDispatchRules() []*UrlDispatchRule {
	if m != nil {
		return m.DispatchRules
	}
	return nil
}

func (m *Application) GetAuthDomain() string {
	if m != nil {
		return m.AuthDomain
	}
	return ""
}

func (m *Application) GetLocationId() string {
	if m != nil {
		return m.LocationId
	}
	return ""
}

func (m *Application) GetCodeBucket() string {
	if m != nil {
		return m.CodeBucket
	}
	return ""
}

func (m *Application) GetDefaultCookieExpiration() *google_protobuf1.Duration {
	if m != nil {
		return m.DefaultCookieExpiration
	}
	return nil
}

func (m *Application) GetDefaultHostname() string {
	if m != nil {
		return m.DefaultHostname
	}
	return ""
}

func (m *Application) GetDefaultBucket() string {
	if m != nil {
		return m.DefaultBucket
	}
	return ""
}

// Rules to match an HTTP request and dispatch that request to a service.
type UrlDispatchRule struct {
	// Domain name to match against. The wildcard "`*`" is supported if
	// specified before a period: "`*.`".
	//
	// Defaults to matching all domains: "`*`".
	Domain string `protobuf:"bytes,1,opt,name=domain,proto3" json:"domain,omitempty"`
	// Pathname within the host. Must start with a "`/`". A
	// single "`*`" can be included at the end of the path. The sum
	// of the lengths of the domain and path may not exceed 100
	// characters.
	Path string `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
	// Resource ID of a service in this application that should
	// serve the matched request. The service must already
	// exist. Example: `default`.
	Service string `protobuf:"bytes,3,opt,name=service,proto3" json:"service,omitempty"`
}

func (m *UrlDispatchRule) Reset()                    { *m = UrlDispatchRule{} }
func (m *UrlDispatchRule) String() string            { return proto.CompactTextString(m) }
func (*UrlDispatchRule) ProtoMessage()               {}
func (*UrlDispatchRule) Descriptor() ([]byte, []int) { return fileDescriptorApplication, []int{1} }

func (m *UrlDispatchRule) GetDomain() string {
	if m != nil {
		return m.Domain
	}
	return ""
}

func (m *UrlDispatchRule) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *UrlDispatchRule) GetService() string {
	if m != nil {
		return m.Service
	}
	return ""
}

func init() {
	proto.RegisterType((*Application)(nil), "google.appengine.v1.Application")
	proto.RegisterType((*UrlDispatchRule)(nil), "google.appengine.v1.UrlDispatchRule")
}

func init() { proto.RegisterFile("google/appengine/v1/application.proto", fileDescriptorApplication) }

var fileDescriptorApplication = []byte{
	// 377 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0xdf, 0x6a, 0xc2, 0x30,
	0x18, 0xc5, 0xa9, 0x1d, 0x3a, 0xd3, 0xf9, 0x87, 0x0c, 0x66, 0x94, 0xb1, 0x89, 0x4c, 0x70, 0x17,
	0x6b, 0xd1, 0x3d, 0xc1, 0x9c, 0x83, 0x8d, 0xdd, 0x48, 0x41, 0x76, 0x59, 0x62, 0x13, 0x6d, 0xb0,
	0x36, 0xa5, 0x4d, 0x65, 0xcf, 0xb0, 0xa7, 0x1e, 0xfd, 0x92, 0x3a, 0x19, 0xde, 0xe5, 0xfb, 0xf5,
	0x9c, 0xe4, 0xeb, 0x39, 0x68, 0xbc, 0x95, 0x72, 0x1b, 0x73, 0x8f, 0xa6, 0x29, 0x4f, 0xb6, 0x22,
	0xe1, 0xde, 0x61, 0x5a, 0x0e, 0xb1, 0x08, 0xa9, 0x12, 0x32, 0x71, 0xd3, 0x4c, 0x2a, 0x89, 0xaf,
	0xb5, 0xcc, 0x3d, 0xca, 0xdc, 0xc3, 0x74, 0x70, 0x7b, 0xf4, 0x0a, 0x8f, 0x26, 0x89, 0x54, 0xe0,
	0xc8, 0xb5, 0x65, 0x70, 0x67, 0xbe, 0xc2, 0xb4, 0x2e, 0x36, 0x1e, 0x2b, 0xb2, 0x93, 0x2b, 0x47,
	0x3f, 0x36, 0x72, 0x5e, 0xfe, 0x1e, 0xc2, 0x18, 0x5d, 0x24, 0x74, 0xcf, 0x89, 0x35, 0xb4, 0x26,
	0x4d, 0x1f, 0xce, 0xb8, 0x8d, 0x6a, 0x82, 0x91, 0x1a, 0x90, 0x9a, 0x60, 0xf8, 0x13, 0xb5, 0x99,
	0xc8, 0x53, 0xaa, 0xc2, 0x28, 0xc8, 0x8a, 0x98, 0xe7, 0xc4, 0x1e, 0xda, 0x13, 0x67, 0xf6, 0xe0,
	0x9e, 0xd9, 0xcf, 0x5d, 0x65, 0xf1, 0xc2, 0xa8, 0xfd, 0x22, 0xe6, 0x7e, 0x8b, 0x9d, 0x4c, 0x39,
	0xbe, 0x47, 0x0e, 0x2d, 0x54, 0x14, 0x30, 0xb9, 0xa7, 0x22, 0x21, 0x75, 0x78, 0x05, 0x95, 0x68,
	0x01, 0xa4, 0x14, 0xc4, 0x52, 0x6f, 0x17, 0x08, 0x46, 0x1a, 0x5a, 0x50, 0xa1, 0x0f, 0x56, 0x0a,
	0x42, 0xc9, 0x78, 0xb0, 0x2e, 0xc2, 0x1d, 0x57, 0xe4, 0x52, 0x0b, 0x4a, 0x34, 0x07, 0x82, 0x57,
	0xa8, 0xcf, 0xf8, 0x86, 0x16, 0xb1, 0x0a, 0x42, 0x29, 0x77, 0x82, 0x07, 0xfc, 0x3b, 0x15, 0x3a,
	0x06, 0xd2, 0x1c, 0x5a, 0x13, 0x67, 0xd6, 0xaf, 0x56, 0xaf, 0x72, 0x72, 0x17, 0x26, 0x27, 0xbf,
	0x67, 0xbc, 0xaf, 0x60, 0x7d, 0x3b, 0x3a, 0xf1, 0x23, 0xea, 0x56, 0xd7, 0x46, 0x32, 0x57, 0x10,
	0x9b, 0x03, 0x8f, 0x77, 0x0c, 0x7f, 0x37, 0x18, 0x8f, 0x51, 0xbb, 0x92, 0x9a, 0x2d, 0xaf, 0x40,
	0xd8, 0x32, 0x54, 0x2f, 0x3a, 0xfa, 0x42, 0x9d, 0x7f, 0x69, 0xe1, 0x1b, 0x54, 0x37, 0xc9, 0xe8,
	0x46, 0xcc, 0x54, 0xf6, 0x94, 0x52, 0x15, 0x99, 0x56, 0xe0, 0x8c, 0x09, 0x6a, 0xe4, 0x3c, 0x3b,
	0x88, 0x90, 0x13, 0x1b, 0x70, 0x35, 0xce, 0x9f, 0x50, 0x2f, 0x94, 0xfb, 0x73, 0xf5, 0xcc, 0xbb,
	0x27, 0xed, 0x2f, 0xcb, 0x9f, 0x5f, 0x5a, 0xeb, 0x3a, 0xa4, 0xf0, 0xfc, 0x1b, 0x00, 0x00, 0xff,
	0xff, 0xab, 0x99, 0xbe, 0xe9, 0x97, 0x02, 0x00, 0x00,
}