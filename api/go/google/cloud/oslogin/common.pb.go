// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/oslogin/common.proto

package google_cloud_oslogin_common

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "go.pedge.io/pb/go/google/api"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// The POSIX account information associated with a Google account.
type PosixAccount struct {
	// Only one POSIX account can be marked as primary.
	Primary bool `protobuf:"varint,1,opt,name=primary" json:"primary,omitempty"`
	// The username of the POSIX account.
	Username string `protobuf:"bytes,2,opt,name=username" json:"username,omitempty"`
	// The user ID.
	Uid int64 `protobuf:"varint,3,opt,name=uid" json:"uid,omitempty"`
	// The default group ID.
	Gid int64 `protobuf:"varint,4,opt,name=gid" json:"gid,omitempty"`
	// The path to the home directory for this account.
	HomeDirectory string `protobuf:"bytes,5,opt,name=home_directory,json=homeDirectory" json:"home_directory,omitempty"`
	// The path to the logic shell for this account.
	Shell string `protobuf:"bytes,6,opt,name=shell" json:"shell,omitempty"`
	// The GECOS (user information) entry for this account.
	Gecos string `protobuf:"bytes,7,opt,name=gecos" json:"gecos,omitempty"`
	// System identifier for which account the username or uid applies to.
	// By default, the empty value is used.
	SystemId string `protobuf:"bytes,8,opt,name=system_id,json=systemId" json:"system_id,omitempty"`
	// Output only. A POSIX account identifier.
	AccountId string `protobuf:"bytes,9,opt,name=account_id,json=accountId" json:"account_id,omitempty"`
}

func (m *PosixAccount) Reset()                    { *m = PosixAccount{} }
func (m *PosixAccount) String() string            { return proto.CompactTextString(m) }
func (*PosixAccount) ProtoMessage()               {}
func (*PosixAccount) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *PosixAccount) GetPrimary() bool {
	if m != nil {
		return m.Primary
	}
	return false
}

func (m *PosixAccount) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *PosixAccount) GetUid() int64 {
	if m != nil {
		return m.Uid
	}
	return 0
}

func (m *PosixAccount) GetGid() int64 {
	if m != nil {
		return m.Gid
	}
	return 0
}

func (m *PosixAccount) GetHomeDirectory() string {
	if m != nil {
		return m.HomeDirectory
	}
	return ""
}

func (m *PosixAccount) GetShell() string {
	if m != nil {
		return m.Shell
	}
	return ""
}

func (m *PosixAccount) GetGecos() string {
	if m != nil {
		return m.Gecos
	}
	return ""
}

func (m *PosixAccount) GetSystemId() string {
	if m != nil {
		return m.SystemId
	}
	return ""
}

func (m *PosixAccount) GetAccountId() string {
	if m != nil {
		return m.AccountId
	}
	return ""
}

// The SSH public key information associated with a Google account.
type SshPublicKey struct {
	// Public key text in SSH format, defined by
	// <a href="https://www.ietf.org/rfc/rfc4253.txt" target="_blank">RFC4253</a>
	// section 6.6.
	Key string `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	// An expiration time in microseconds since epoch.
	ExpirationTimeUsec int64 `protobuf:"varint,2,opt,name=expiration_time_usec,json=expirationTimeUsec" json:"expiration_time_usec,omitempty"`
	// Output only. The SHA-256 fingerprint of the SSH public key.
	Fingerprint string `protobuf:"bytes,3,opt,name=fingerprint" json:"fingerprint,omitempty"`
}

func (m *SshPublicKey) Reset()                    { *m = SshPublicKey{} }
func (m *SshPublicKey) String() string            { return proto.CompactTextString(m) }
func (*SshPublicKey) ProtoMessage()               {}
func (*SshPublicKey) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *SshPublicKey) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *SshPublicKey) GetExpirationTimeUsec() int64 {
	if m != nil {
		return m.ExpirationTimeUsec
	}
	return 0
}

func (m *SshPublicKey) GetFingerprint() string {
	if m != nil {
		return m.Fingerprint
	}
	return ""
}

func init() {
	proto.RegisterType((*PosixAccount)(nil), "google.cloud.oslogin.common.PosixAccount")
	proto.RegisterType((*SshPublicKey)(nil), "google.cloud.oslogin.common.SshPublicKey")
}

func init() { proto.RegisterFile("google/cloud/oslogin/common.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 344 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0x5f, 0xcb, 0xd3, 0x30,
	0x14, 0xc6, 0xa9, 0xf5, 0x7d, 0xdf, 0x35, 0x4e, 0x91, 0xb0, 0x8b, 0xb0, 0x29, 0xd6, 0x81, 0xb0,
	0xab, 0x56, 0xf0, 0x13, 0x38, 0xbc, 0x19, 0x0a, 0x8e, 0xaa, 0xd7, 0xa5, 0x4b, 0x8e, 0xdd, 0xc1,
	0x26, 0xa7, 0x24, 0x29, 0xac, 0x9f, 0xdc, 0x5b, 0x49, 0xb2, 0xa9, 0x57, 0xef, 0x5d, 0x9e, 0xdf,
	0xf3, 0x9c, 0x70, 0xfe, 0xb0, 0xb7, 0x3d, 0x51, 0x3f, 0x40, 0x2d, 0x07, 0x9a, 0x54, 0x4d, 0x6e,
	0xa0, 0x1e, 0x4d, 0x2d, 0x49, 0x6b, 0x32, 0xd5, 0x68, 0xc9, 0x13, 0xdf, 0xa4, 0x48, 0x15, 0x23,
	0xd5, 0x35, 0x52, 0xa5, 0xc8, 0xfa, 0xd5, 0xb5, 0xbe, 0x1b, 0xb1, 0xee, 0x8c, 0x21, 0xdf, 0x79,
	0x24, 0xe3, 0x52, 0xe9, 0xf6, 0x77, 0xc6, 0x96, 0x47, 0x72, 0x78, 0xf9, 0x28, 0x25, 0x4d, 0xc6,
	0x73, 0xc1, 0x1e, 0x46, 0x8b, 0xba, 0xb3, 0xb3, 0xc8, 0xca, 0x6c, 0xb7, 0x68, 0x6e, 0x92, 0xaf,
	0xd9, 0x62, 0x72, 0x60, 0x4d, 0xa7, 0x41, 0x3c, 0x29, 0xb3, 0x5d, 0xd1, 0xfc, 0xd5, 0xfc, 0x25,
	0xcb, 0x27, 0x54, 0x22, 0x2f, 0xb3, 0x5d, 0xde, 0x84, 0x67, 0x20, 0x3d, 0x2a, 0xf1, 0x34, 0x91,
	0x1e, 0x15, 0x7f, 0xc7, 0x5e, 0x9c, 0x49, 0x43, 0xab, 0xd0, 0x82, 0xf4, 0x64, 0x67, 0x71, 0x17,
	0x7f, 0x79, 0x1e, 0xe8, 0xa7, 0x1b, 0xe4, 0x2b, 0x76, 0xe7, 0xce, 0x30, 0x0c, 0xe2, 0x3e, 0xba,
	0x49, 0x04, 0xda, 0x83, 0x24, 0x27, 0x1e, 0x12, 0x8d, 0x82, 0x6f, 0x58, 0xe1, 0x66, 0xe7, 0x41,
	0xb7, 0xa8, 0xc4, 0x22, 0xf5, 0x94, 0xc0, 0x41, 0xf1, 0xd7, 0x8c, 0x75, 0x69, 0xa8, 0xe0, 0x16,
	0xd1, 0x2d, 0xae, 0xe4, 0xa0, 0xb6, 0x9e, 0x2d, 0xbf, 0xb9, 0xf3, 0x71, 0x3a, 0x0d, 0x28, 0x3f,
	0xc3, 0x1c, 0x1a, 0xfe, 0x05, 0x69, 0xe8, 0xa2, 0x09, 0x4f, 0xfe, 0x9e, 0xad, 0xe0, 0x32, 0xa2,
	0x8d, 0x0b, 0x6b, 0x3d, 0x6a, 0x68, 0x27, 0x07, 0x32, 0x0e, 0x9f, 0x37, 0xfc, 0x9f, 0xf7, 0x1d,
	0x35, 0xfc, 0x70, 0x20, 0x79, 0xc9, 0x9e, 0xfd, 0x44, 0xd3, 0x83, 0x1d, 0x2d, 0x1a, 0x1f, 0xd7,
	0x51, 0x34, 0xff, 0xa3, 0x7d, 0xcd, 0xde, 0x48, 0xd2, 0xd5, 0x23, 0x07, 0xdb, 0x2f, 0xbf, 0xba,
	0x2f, 0x41, 0x1f, 0xc3, 0x81, 0x4e, 0xf7, 0xf1, 0x4e, 0x1f, 0xfe, 0x04, 0x00, 0x00, 0xff, 0xff,
	0x74, 0xb6, 0x31, 0x83, 0x07, 0x02, 0x00, 0x00,
}