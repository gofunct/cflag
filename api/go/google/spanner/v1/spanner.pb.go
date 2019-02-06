// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/spanner/v1/spanner.proto

package google_spanner_v1

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "go.pedge.io/pb/go/google/api"
import _ "github.com/golang/protobuf/ptypes/empty"
import google_protobuf1 "github.com/golang/protobuf/ptypes/struct"
import google_protobuf3 "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Mode in which the query must be processed.
type ExecuteSqlRequest_QueryMode int32

const (
	// The default mode where only the query result, without any information
	// about the query plan is returned.
	ExecuteSqlRequest_NORMAL ExecuteSqlRequest_QueryMode = 0
	// This mode returns only the query plan, without any result rows or
	// execution statistics information.
	ExecuteSqlRequest_PLAN ExecuteSqlRequest_QueryMode = 1
	// This mode returns both the query plan and the execution statistics along
	// with the result rows.
	ExecuteSqlRequest_PROFILE ExecuteSqlRequest_QueryMode = 2
)

var ExecuteSqlRequest_QueryMode_name = map[int32]string{
	0: "NORMAL",
	1: "PLAN",
	2: "PROFILE",
}
var ExecuteSqlRequest_QueryMode_value = map[string]int32{
	"NORMAL":  0,
	"PLAN":    1,
	"PROFILE": 2,
}

func (x ExecuteSqlRequest_QueryMode) String() string {
	return proto.EnumName(ExecuteSqlRequest_QueryMode_name, int32(x))
}
func (ExecuteSqlRequest_QueryMode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor4, []int{6, 0}
}

// The request for [CreateSession][google.spanner.v1.Spanner.CreateSession].
type CreateSessionRequest struct {
	// Required. The database in which the new session is created.
	Database string `protobuf:"bytes,1,opt,name=database" json:"database,omitempty"`
	// The session to create.
	Session *Session `protobuf:"bytes,2,opt,name=session" json:"session,omitempty"`
}

func (m *CreateSessionRequest) Reset()                    { *m = CreateSessionRequest{} }
func (m *CreateSessionRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateSessionRequest) ProtoMessage()               {}
func (*CreateSessionRequest) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{0} }

func (m *CreateSessionRequest) GetDatabase() string {
	if m != nil {
		return m.Database
	}
	return ""
}

func (m *CreateSessionRequest) GetSession() *Session {
	if m != nil {
		return m.Session
	}
	return nil
}

// A session in the Cloud Spanner API.
type Session struct {
	// The name of the session. This is always system-assigned; values provided
	// when creating a session are ignored.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// The labels for the session.
	//
	//  * Label keys must be between 1 and 63 characters long and must conform to
	//    the following regular expression: `[a-z]([-a-z0-9]*[a-z0-9])?`.
	//  * Label values must be between 0 and 63 characters long and must conform
	//    to the regular expression `([a-z]([-a-z0-9]*[a-z0-9])?)?`.
	//  * No more than 64 labels can be associated with a given session.
	//
	// See https://goo.gl/xmQnxf for more information on and examples of labels.
	Labels map[string]string `protobuf:"bytes,2,rep,name=labels" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// Output only. The timestamp when the session is created.
	CreateTime *google_protobuf3.Timestamp `protobuf:"bytes,3,opt,name=create_time,json=createTime" json:"create_time,omitempty"`
	// Output only. The approximate timestamp when the session is last used. It is
	// typically earlier than the actual last use time.
	ApproximateLastUseTime *google_protobuf3.Timestamp `protobuf:"bytes,4,opt,name=approximate_last_use_time,json=approximateLastUseTime" json:"approximate_last_use_time,omitempty"`
}

func (m *Session) Reset()                    { *m = Session{} }
func (m *Session) String() string            { return proto.CompactTextString(m) }
func (*Session) ProtoMessage()               {}
func (*Session) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{1} }

func (m *Session) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Session) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *Session) GetCreateTime() *google_protobuf3.Timestamp {
	if m != nil {
		return m.CreateTime
	}
	return nil
}

func (m *Session) GetApproximateLastUseTime() *google_protobuf3.Timestamp {
	if m != nil {
		return m.ApproximateLastUseTime
	}
	return nil
}

// The request for [GetSession][google.spanner.v1.Spanner.GetSession].
type GetSessionRequest struct {
	// Required. The name of the session to retrieve.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *GetSessionRequest) Reset()                    { *m = GetSessionRequest{} }
func (m *GetSessionRequest) String() string            { return proto.CompactTextString(m) }
func (*GetSessionRequest) ProtoMessage()               {}
func (*GetSessionRequest) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{2} }

func (m *GetSessionRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// The request for [ListSessions][google.spanner.v1.Spanner.ListSessions].
type ListSessionsRequest struct {
	// Required. The database in which to list sessions.
	Database string `protobuf:"bytes,1,opt,name=database" json:"database,omitempty"`
	// Number of sessions to be returned in the response. If 0 or less, defaults
	// to the server's maximum allowed page size.
	PageSize int32 `protobuf:"varint,2,opt,name=page_size,json=pageSize" json:"page_size,omitempty"`
	// If non-empty, `page_token` should contain a
	// [next_page_token][google.spanner.v1.ListSessionsResponse.next_page_token] from a previous
	// [ListSessionsResponse][google.spanner.v1.ListSessionsResponse].
	PageToken string `protobuf:"bytes,3,opt,name=page_token,json=pageToken" json:"page_token,omitempty"`
	// An expression for filtering the results of the request. Filter rules are
	// case insensitive. The fields eligible for filtering are:
	//
	//   * `labels.key` where key is the name of a label
	//
	// Some examples of using filters are:
	//
	//   * `labels.env:*` --> The session has the label "env".
	//   * `labels.env:dev` --> The session has the label "env" and the value of
	//                        the label contains the string "dev".
	Filter string `protobuf:"bytes,4,opt,name=filter" json:"filter,omitempty"`
}

func (m *ListSessionsRequest) Reset()                    { *m = ListSessionsRequest{} }
func (m *ListSessionsRequest) String() string            { return proto.CompactTextString(m) }
func (*ListSessionsRequest) ProtoMessage()               {}
func (*ListSessionsRequest) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{3} }

func (m *ListSessionsRequest) GetDatabase() string {
	if m != nil {
		return m.Database
	}
	return ""
}

func (m *ListSessionsRequest) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *ListSessionsRequest) GetPageToken() string {
	if m != nil {
		return m.PageToken
	}
	return ""
}

func (m *ListSessionsRequest) GetFilter() string {
	if m != nil {
		return m.Filter
	}
	return ""
}

// The response for [ListSessions][google.spanner.v1.Spanner.ListSessions].
type ListSessionsResponse struct {
	// The list of requested sessions.
	Sessions []*Session `protobuf:"bytes,1,rep,name=sessions" json:"sessions,omitempty"`
	// `next_page_token` can be sent in a subsequent
	// [ListSessions][google.spanner.v1.Spanner.ListSessions] call to fetch more of the matching
	// sessions.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken" json:"next_page_token,omitempty"`
}

func (m *ListSessionsResponse) Reset()                    { *m = ListSessionsResponse{} }
func (m *ListSessionsResponse) String() string            { return proto.CompactTextString(m) }
func (*ListSessionsResponse) ProtoMessage()               {}
func (*ListSessionsResponse) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{4} }

func (m *ListSessionsResponse) GetSessions() []*Session {
	if m != nil {
		return m.Sessions
	}
	return nil
}

func (m *ListSessionsResponse) GetNextPageToken() string {
	if m != nil {
		return m.NextPageToken
	}
	return ""
}

// The request for [DeleteSession][google.spanner.v1.Spanner.DeleteSession].
type DeleteSessionRequest struct {
	// Required. The name of the session to delete.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *DeleteSessionRequest) Reset()                    { *m = DeleteSessionRequest{} }
func (m *DeleteSessionRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteSessionRequest) ProtoMessage()               {}
func (*DeleteSessionRequest) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{5} }

func (m *DeleteSessionRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// The request for [ExecuteSql][google.spanner.v1.Spanner.ExecuteSql] and
// [ExecuteStreamingSql][google.spanner.v1.Spanner.ExecuteStreamingSql].
type ExecuteSqlRequest struct {
	// Required. The session in which the SQL query should be performed.
	Session string `protobuf:"bytes,1,opt,name=session" json:"session,omitempty"`
	// The transaction to use. If none is provided, the default is a
	// temporary read-only transaction with strong concurrency.
	Transaction *TransactionSelector `protobuf:"bytes,2,opt,name=transaction" json:"transaction,omitempty"`
	// Required. The SQL query string.
	Sql string `protobuf:"bytes,3,opt,name=sql" json:"sql,omitempty"`
	// The SQL query string can contain parameter placeholders. A parameter
	// placeholder consists of `'@'` followed by the parameter
	// name. Parameter names consist of any combination of letters,
	// numbers, and underscores.
	//
	// Parameters can appear anywhere that a literal value is expected.  The same
	// parameter name can be used more than once, for example:
	//   `"WHERE id > @msg_id AND id < @msg_id + 100"`
	//
	// It is an error to execute an SQL query with unbound parameters.
	//
	// Parameter values are specified using `params`, which is a JSON
	// object whose keys are parameter names, and whose values are the
	// corresponding parameter values.
	Params *google_protobuf1.Struct `protobuf:"bytes,4,opt,name=params" json:"params,omitempty"`
	// It is not always possible for Cloud Spanner to infer the right SQL type
	// from a JSON value.  For example, values of type `BYTES` and values
	// of type `STRING` both appear in [params][google.spanner.v1.ExecuteSqlRequest.params] as JSON strings.
	//
	// In these cases, `param_types` can be used to specify the exact
	// SQL type for some or all of the SQL query parameters. See the
	// definition of [Type][google.spanner.v1.Type] for more information
	// about SQL types.
	ParamTypes map[string]*Type `protobuf:"bytes,5,rep,name=param_types,json=paramTypes" json:"param_types,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// If this request is resuming a previously interrupted SQL query
	// execution, `resume_token` should be copied from the last
	// [PartialResultSet][google.spanner.v1.PartialResultSet] yielded before the interruption. Doing this
	// enables the new SQL query execution to resume where the last one left
	// off. The rest of the request parameters must exactly match the
	// request that yielded this token.
	ResumeToken []byte `protobuf:"bytes,6,opt,name=resume_token,json=resumeToken,proto3" json:"resume_token,omitempty"`
	// Used to control the amount of debugging information returned in
	// [ResultSetStats][google.spanner.v1.ResultSetStats].
	QueryMode ExecuteSqlRequest_QueryMode `protobuf:"varint,7,opt,name=query_mode,json=queryMode,enum=google.spanner.v1.ExecuteSqlRequest_QueryMode" json:"query_mode,omitempty"`
}

func (m *ExecuteSqlRequest) Reset()                    { *m = ExecuteSqlRequest{} }
func (m *ExecuteSqlRequest) String() string            { return proto.CompactTextString(m) }
func (*ExecuteSqlRequest) ProtoMessage()               {}
func (*ExecuteSqlRequest) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{6} }

func (m *ExecuteSqlRequest) GetSession() string {
	if m != nil {
		return m.Session
	}
	return ""
}

func (m *ExecuteSqlRequest) GetTransaction() *TransactionSelector {
	if m != nil {
		return m.Transaction
	}
	return nil
}

func (m *ExecuteSqlRequest) GetSql() string {
	if m != nil {
		return m.Sql
	}
	return ""
}

func (m *ExecuteSqlRequest) GetParams() *google_protobuf1.Struct {
	if m != nil {
		return m.Params
	}
	return nil
}

func (m *ExecuteSqlRequest) GetParamTypes() map[string]*Type {
	if m != nil {
		return m.ParamTypes
	}
	return nil
}

func (m *ExecuteSqlRequest) GetResumeToken() []byte {
	if m != nil {
		return m.ResumeToken
	}
	return nil
}

func (m *ExecuteSqlRequest) GetQueryMode() ExecuteSqlRequest_QueryMode {
	if m != nil {
		return m.QueryMode
	}
	return ExecuteSqlRequest_NORMAL
}

// The request for [Read][google.spanner.v1.Spanner.Read] and
// [StreamingRead][google.spanner.v1.Spanner.StreamingRead].
type ReadRequest struct {
	// Required. The session in which the read should be performed.
	Session string `protobuf:"bytes,1,opt,name=session" json:"session,omitempty"`
	// The transaction to use. If none is provided, the default is a
	// temporary read-only transaction with strong concurrency.
	Transaction *TransactionSelector `protobuf:"bytes,2,opt,name=transaction" json:"transaction,omitempty"`
	// Required. The name of the table in the database to be read.
	Table string `protobuf:"bytes,3,opt,name=table" json:"table,omitempty"`
	// If non-empty, the name of an index on [table][google.spanner.v1.ReadRequest.table]. This index is
	// used instead of the table primary key when interpreting [key_set][google.spanner.v1.ReadRequest.key_set]
	// and sorting result rows. See [key_set][google.spanner.v1.ReadRequest.key_set] for further information.
	Index string `protobuf:"bytes,4,opt,name=index" json:"index,omitempty"`
	// The columns of [table][google.spanner.v1.ReadRequest.table] to be returned for each row matching
	// this request.
	Columns []string `protobuf:"bytes,5,rep,name=columns" json:"columns,omitempty"`
	// Required. `key_set` identifies the rows to be yielded. `key_set` names the
	// primary keys of the rows in [table][google.spanner.v1.ReadRequest.table] to be yielded, unless [index][google.spanner.v1.ReadRequest.index]
	// is present. If [index][google.spanner.v1.ReadRequest.index] is present, then [key_set][google.spanner.v1.ReadRequest.key_set] instead names
	// index keys in [index][google.spanner.v1.ReadRequest.index].
	//
	// Rows are yielded in table primary key order (if [index][google.spanner.v1.ReadRequest.index] is empty)
	// or index key order (if [index][google.spanner.v1.ReadRequest.index] is non-empty).
	//
	// It is not an error for the `key_set` to name rows that do not
	// exist in the database. Read yields nothing for nonexistent rows.
	KeySet *KeySet `protobuf:"bytes,6,opt,name=key_set,json=keySet" json:"key_set,omitempty"`
	// If greater than zero, only the first `limit` rows are yielded. If `limit`
	// is zero, the default is no limit.
	Limit int64 `protobuf:"varint,8,opt,name=limit" json:"limit,omitempty"`
	// If this request is resuming a previously interrupted read,
	// `resume_token` should be copied from the last
	// [PartialResultSet][google.spanner.v1.PartialResultSet] yielded before the interruption. Doing this
	// enables the new read to resume where the last read left off. The
	// rest of the request parameters must exactly match the request
	// that yielded this token.
	ResumeToken []byte `protobuf:"bytes,9,opt,name=resume_token,json=resumeToken,proto3" json:"resume_token,omitempty"`
}

func (m *ReadRequest) Reset()                    { *m = ReadRequest{} }
func (m *ReadRequest) String() string            { return proto.CompactTextString(m) }
func (*ReadRequest) ProtoMessage()               {}
func (*ReadRequest) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{7} }

func (m *ReadRequest) GetSession() string {
	if m != nil {
		return m.Session
	}
	return ""
}

func (m *ReadRequest) GetTransaction() *TransactionSelector {
	if m != nil {
		return m.Transaction
	}
	return nil
}

func (m *ReadRequest) GetTable() string {
	if m != nil {
		return m.Table
	}
	return ""
}

func (m *ReadRequest) GetIndex() string {
	if m != nil {
		return m.Index
	}
	return ""
}

func (m *ReadRequest) GetColumns() []string {
	if m != nil {
		return m.Columns
	}
	return nil
}

func (m *ReadRequest) GetKeySet() *KeySet {
	if m != nil {
		return m.KeySet
	}
	return nil
}

func (m *ReadRequest) GetLimit() int64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *ReadRequest) GetResumeToken() []byte {
	if m != nil {
		return m.ResumeToken
	}
	return nil
}

// The request for [BeginTransaction][google.spanner.v1.Spanner.BeginTransaction].
type BeginTransactionRequest struct {
	// Required. The session in which the transaction runs.
	Session string `protobuf:"bytes,1,opt,name=session" json:"session,omitempty"`
	// Required. Options for the new transaction.
	Options *TransactionOptions `protobuf:"bytes,2,opt,name=options" json:"options,omitempty"`
}

func (m *BeginTransactionRequest) Reset()                    { *m = BeginTransactionRequest{} }
func (m *BeginTransactionRequest) String() string            { return proto.CompactTextString(m) }
func (*BeginTransactionRequest) ProtoMessage()               {}
func (*BeginTransactionRequest) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{8} }

func (m *BeginTransactionRequest) GetSession() string {
	if m != nil {
		return m.Session
	}
	return ""
}

func (m *BeginTransactionRequest) GetOptions() *TransactionOptions {
	if m != nil {
		return m.Options
	}
	return nil
}

// The request for [Commit][google.spanner.v1.Spanner.Commit].
type CommitRequest struct {
	// Required. The session in which the transaction to be committed is running.
	Session string `protobuf:"bytes,1,opt,name=session" json:"session,omitempty"`
	// Required. The transaction in which to commit.
	//
	// Types that are valid to be assigned to Transaction:
	//	*CommitRequest_TransactionId
	//	*CommitRequest_SingleUseTransaction
	Transaction isCommitRequest_Transaction `protobuf_oneof:"transaction"`
	// The mutations to be executed when this transaction commits. All
	// mutations are applied atomically, in the order they appear in
	// this list.
	Mutations []*Mutation `protobuf:"bytes,4,rep,name=mutations" json:"mutations,omitempty"`
}

func (m *CommitRequest) Reset()                    { *m = CommitRequest{} }
func (m *CommitRequest) String() string            { return proto.CompactTextString(m) }
func (*CommitRequest) ProtoMessage()               {}
func (*CommitRequest) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{9} }

type isCommitRequest_Transaction interface {
	isCommitRequest_Transaction()
}

type CommitRequest_TransactionId struct {
	TransactionId []byte `protobuf:"bytes,2,opt,name=transaction_id,json=transactionId,proto3,oneof"`
}
type CommitRequest_SingleUseTransaction struct {
	SingleUseTransaction *TransactionOptions `protobuf:"bytes,3,opt,name=single_use_transaction,json=singleUseTransaction,oneof"`
}

func (*CommitRequest_TransactionId) isCommitRequest_Transaction()        {}
func (*CommitRequest_SingleUseTransaction) isCommitRequest_Transaction() {}

func (m *CommitRequest) GetTransaction() isCommitRequest_Transaction {
	if m != nil {
		return m.Transaction
	}
	return nil
}

func (m *CommitRequest) GetSession() string {
	if m != nil {
		return m.Session
	}
	return ""
}

func (m *CommitRequest) GetTransactionId() []byte {
	if x, ok := m.GetTransaction().(*CommitRequest_TransactionId); ok {
		return x.TransactionId
	}
	return nil
}

func (m *CommitRequest) GetSingleUseTransaction() *TransactionOptions {
	if x, ok := m.GetTransaction().(*CommitRequest_SingleUseTransaction); ok {
		return x.SingleUseTransaction
	}
	return nil
}

func (m *CommitRequest) GetMutations() []*Mutation {
	if m != nil {
		return m.Mutations
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*CommitRequest) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _CommitRequest_OneofMarshaler, _CommitRequest_OneofUnmarshaler, _CommitRequest_OneofSizer, []interface{}{
		(*CommitRequest_TransactionId)(nil),
		(*CommitRequest_SingleUseTransaction)(nil),
	}
}

func _CommitRequest_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*CommitRequest)
	// transaction
	switch x := m.Transaction.(type) {
	case *CommitRequest_TransactionId:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		b.EncodeRawBytes(x.TransactionId)
	case *CommitRequest_SingleUseTransaction:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.SingleUseTransaction); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("CommitRequest.Transaction has unexpected type %T", x)
	}
	return nil
}

func _CommitRequest_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*CommitRequest)
	switch tag {
	case 2: // transaction.transaction_id
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeRawBytes(true)
		m.Transaction = &CommitRequest_TransactionId{x}
		return true, err
	case 3: // transaction.single_use_transaction
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(TransactionOptions)
		err := b.DecodeMessage(msg)
		m.Transaction = &CommitRequest_SingleUseTransaction{msg}
		return true, err
	default:
		return false, nil
	}
}

func _CommitRequest_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*CommitRequest)
	// transaction
	switch x := m.Transaction.(type) {
	case *CommitRequest_TransactionId:
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.TransactionId)))
		n += len(x.TransactionId)
	case *CommitRequest_SingleUseTransaction:
		s := proto.Size(x.SingleUseTransaction)
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// The response for [Commit][google.spanner.v1.Spanner.Commit].
type CommitResponse struct {
	// The Cloud Spanner timestamp at which the transaction committed.
	CommitTimestamp *google_protobuf3.Timestamp `protobuf:"bytes,1,opt,name=commit_timestamp,json=commitTimestamp" json:"commit_timestamp,omitempty"`
}

func (m *CommitResponse) Reset()                    { *m = CommitResponse{} }
func (m *CommitResponse) String() string            { return proto.CompactTextString(m) }
func (*CommitResponse) ProtoMessage()               {}
func (*CommitResponse) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{10} }

func (m *CommitResponse) GetCommitTimestamp() *google_protobuf3.Timestamp {
	if m != nil {
		return m.CommitTimestamp
	}
	return nil
}

// The request for [Rollback][google.spanner.v1.Spanner.Rollback].
type RollbackRequest struct {
	// Required. The session in which the transaction to roll back is running.
	Session string `protobuf:"bytes,1,opt,name=session" json:"session,omitempty"`
	// Required. The transaction to roll back.
	TransactionId []byte `protobuf:"bytes,2,opt,name=transaction_id,json=transactionId,proto3" json:"transaction_id,omitempty"`
}

func (m *RollbackRequest) Reset()                    { *m = RollbackRequest{} }
func (m *RollbackRequest) String() string            { return proto.CompactTextString(m) }
func (*RollbackRequest) ProtoMessage()               {}
func (*RollbackRequest) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{11} }

func (m *RollbackRequest) GetSession() string {
	if m != nil {
		return m.Session
	}
	return ""
}

func (m *RollbackRequest) GetTransactionId() []byte {
	if m != nil {
		return m.TransactionId
	}
	return nil
}

func init() {
	proto.RegisterType((*CreateSessionRequest)(nil), "google.spanner.v1.CreateSessionRequest")
	proto.RegisterType((*Session)(nil), "google.spanner.v1.Session")
	proto.RegisterType((*GetSessionRequest)(nil), "google.spanner.v1.GetSessionRequest")
	proto.RegisterType((*ListSessionsRequest)(nil), "google.spanner.v1.ListSessionsRequest")
	proto.RegisterType((*ListSessionsResponse)(nil), "google.spanner.v1.ListSessionsResponse")
	proto.RegisterType((*DeleteSessionRequest)(nil), "google.spanner.v1.DeleteSessionRequest")
	proto.RegisterType((*ExecuteSqlRequest)(nil), "google.spanner.v1.ExecuteSqlRequest")
	proto.RegisterType((*ReadRequest)(nil), "google.spanner.v1.ReadRequest")
	proto.RegisterType((*BeginTransactionRequest)(nil), "google.spanner.v1.BeginTransactionRequest")
	proto.RegisterType((*CommitRequest)(nil), "google.spanner.v1.CommitRequest")
	proto.RegisterType((*CommitResponse)(nil), "google.spanner.v1.CommitResponse")
	proto.RegisterType((*RollbackRequest)(nil), "google.spanner.v1.RollbackRequest")
	proto.RegisterEnum("google.spanner.v1.ExecuteSqlRequest_QueryMode", ExecuteSqlRequest_QueryMode_name, ExecuteSqlRequest_QueryMode_value)
}

func init() { proto.RegisterFile("google/spanner/v1/spanner.proto", fileDescriptor4) }

var fileDescriptor4 = []byte{
	// 1387 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x56, 0xcd, 0x6f, 0x13, 0x47,
	0x14, 0x67, 0x9d, 0xc4, 0x8e, 0x9f, 0xf3, 0xc5, 0x90, 0x06, 0x63, 0x28, 0x98, 0xa5, 0x90, 0xc8,
	0x52, 0xed, 0x26, 0x45, 0x15, 0xb8, 0x1f, 0x40, 0xc0, 0x40, 0x84, 0x43, 0xcc, 0x3a, 0x80, 0xd4,
	0x52, 0x59, 0x63, 0x7b, 0x48, 0xb7, 0xd9, 0xaf, 0xec, 0x8c, 0xa3, 0x98, 0x8a, 0x4b, 0xa5, 0x9e,
	0x7a, 0x69, 0xe9, 0xa1, 0x87, 0xf6, 0xd6, 0x9e, 0x2a, 0xee, 0xbd, 0xf5, 0x0f, 0xe8, 0xb5, 0xff,
	0x42, 0xff, 0x8b, 0x5e, 0xaa, 0xf9, 0x72, 0x36, 0xf6, 0xc6, 0x09, 0x72, 0xd5, 0xd3, 0xce, 0xbc,
	0xf7, 0x66, 0xde, 0x6f, 0x7f, 0xef, 0xcd, 0xfc, 0x06, 0x2e, 0x6c, 0xf9, 0xfe, 0x96, 0x43, 0x4a,
	0x34, 0xc0, 0x9e, 0x47, 0xc2, 0xd2, 0xee, 0xb2, 0x1e, 0x16, 0x83, 0xd0, 0x67, 0x3e, 0x3a, 0x29,
	0x03, 0x8a, 0xda, 0xba, 0xbb, 0x9c, 0x3b, 0xa7, 0xd6, 0xe0, 0xc0, 0x2e, 0x61, 0xcf, 0xf3, 0x19,
	0x66, 0xb6, 0xef, 0x51, 0xb9, 0x20, 0x77, 0x56, 0x79, 0xc5, 0xac, 0xd9, 0x79, 0x5e, 0x22, 0x6e,
	0xc0, 0xba, 0xca, 0x79, 0xae, 0xdf, 0x49, 0x59, 0xd8, 0x69, 0x31, 0xe5, 0xbd, 0xd0, 0xef, 0x65,
	0xb6, 0x4b, 0x28, 0xc3, 0x6e, 0xd0, 0xb7, 0x3c, 0x82, 0x76, 0x9b, 0x74, 0x75, 0xe6, 0xfc, 0xa0,
	0xd7, 0xed, 0x48, 0x70, 0x2a, 0xc2, 0x1c, 0x8c, 0x08, 0x09, 0xed, 0x38, 0xac, 0x41, 0x89, 0x06,
	0x71, 0x69, 0x30, 0x86, 0x85, 0xd8, 0xa3, 0xb8, 0x15, 0xd9, 0x28, 0x06, 0x08, 0xeb, 0x06, 0x44,
	0x7a, 0xcd, 0x2f, 0x60, 0xfe, 0x76, 0x48, 0x30, 0x23, 0x75, 0x42, 0xa9, 0xed, 0x7b, 0x16, 0xd9,
	0xe9, 0x10, 0xca, 0x50, 0x0e, 0x26, 0xdb, 0x98, 0xe1, 0x26, 0xa6, 0x24, 0x6b, 0xe4, 0x8d, 0xa5,
	0xb4, 0xd5, 0x9b, 0xa3, 0xab, 0x90, 0xa2, 0x32, 0x3a, 0x9b, 0xc8, 0x1b, 0x4b, 0x99, 0x95, 0x5c,
	0x71, 0x80, 0xf9, 0xa2, 0xde, 0x4f, 0x87, 0x9a, 0xaf, 0x13, 0x90, 0x52, 0x46, 0x84, 0x60, 0xdc,
	0xc3, 0xae, 0xde, 0x59, 0x8c, 0xd1, 0x27, 0x90, 0x74, 0x70, 0x93, 0x38, 0x34, 0x9b, 0xc8, 0x8f,
	0x2d, 0x65, 0x56, 0xae, 0x1c, 0xbe, 0x69, 0xb1, 0x2a, 0x02, 0x2b, 0x1e, 0x0b, 0xbb, 0x96, 0x5a,
	0x85, 0x3e, 0x84, 0x4c, 0x4b, 0xfc, 0x49, 0x83, 0x97, 0x22, 0x3b, 0x76, 0x10, 0x99, 0xae, 0x53,
	0x71, 0x53, 0xd7, 0xc9, 0x02, 0x19, 0xce, 0x0d, 0xe8, 0x31, 0x9c, 0xc1, 0x41, 0x10, 0xfa, 0x7b,
	0xb6, 0xcb, 0x77, 0x70, 0x30, 0x65, 0x8d, 0x0e, 0x55, 0x5b, 0x8d, 0x1f, 0xb9, 0xd5, 0x42, 0x64,
	0x71, 0x15, 0x53, 0xf6, 0x98, 0x8a, 0x6d, 0x73, 0xd7, 0x21, 0x13, 0x81, 0x8a, 0xe6, 0x60, 0x6c,
	0x9b, 0x74, 0xd5, 0x5f, 0xf3, 0x21, 0x9a, 0x87, 0x89, 0x5d, 0xec, 0x74, 0x88, 0x20, 0x32, 0x6d,
	0xc9, 0x49, 0x39, 0x71, 0xcd, 0x30, 0x17, 0xe1, 0xe4, 0x3d, 0xc2, 0xfa, 0xaa, 0x12, 0xc3, 0x9b,
	0xf9, 0x8d, 0x01, 0xa7, 0xaa, 0x36, 0xd5, 0xa1, 0xf4, 0x38, 0x15, 0x3c, 0x0b, 0xe9, 0x00, 0x6f,
	0x91, 0x06, 0xb5, 0x5f, 0xc8, 0xd4, 0x13, 0xd6, 0x24, 0x37, 0xd4, 0xed, 0x17, 0x04, 0xbd, 0x0d,
	0x20, 0x9c, 0xcc, 0xdf, 0x26, 0x9e, 0xe0, 0x31, 0x6d, 0x89, 0xf0, 0x4d, 0x6e, 0x40, 0x0b, 0x90,
	0x7c, 0x6e, 0x3b, 0x8c, 0x84, 0x82, 0x97, 0xb4, 0xa5, 0x66, 0xe6, 0x2e, 0xcc, 0x1f, 0x84, 0x41,
	0x03, 0xdf, 0xa3, 0x04, 0x7d, 0x00, 0x93, 0xaa, 0x05, 0x68, 0xd6, 0x10, 0x95, 0x1d, 0xd6, 0x2e,
	0xbd, 0x58, 0x74, 0x05, 0x66, 0x3d, 0xb2, 0xc7, 0x1a, 0x11, 0x2c, 0x92, 0xa4, 0x69, 0x6e, 0xae,
	0x69, 0x3c, 0x66, 0x01, 0xe6, 0xef, 0x10, 0x87, 0x0c, 0x74, 0x70, 0x1c, 0x57, 0xdf, 0x8e, 0xc3,
	0xc9, 0xca, 0x1e, 0x69, 0x75, 0x18, 0xa9, 0xef, 0x38, 0x3a, 0x32, 0xbb, 0xdf, 0xcf, 0x32, 0x58,
	0x4f, 0xd1, 0x7d, 0xc8, 0x44, 0x0e, 0x94, 0xea, 0xf6, 0xb8, 0xc6, 0xdc, 0xdc, 0x8f, 0xaa, 0x13,
	0x87, 0xb4, 0x98, 0x1f, 0x5a, 0xd1, 0xa5, 0xbc, 0xf4, 0x74, 0xc7, 0x51, 0x6c, 0xf2, 0x21, 0x2a,
	0x41, 0x32, 0xc0, 0x21, 0x76, 0xa9, 0xea, 0xaf, 0xd3, 0x03, 0xfd, 0x55, 0x17, 0x17, 0x8e, 0xa5,
	0xc2, 0xd0, 0x63, 0xc8, 0x88, 0x51, 0x83, 0x1f, 0x5f, 0x9a, 0x9d, 0x10, 0x5c, 0x5e, 0x8d, 0x01,
	0x33, 0xf0, 0x87, 0xc5, 0x1a, 0x5f, 0xb7, 0xc9, 0x97, 0xc9, 0x33, 0x03, 0x41, 0xcf, 0x80, 0x2e,
	0xc2, 0x14, 0xbf, 0x58, 0x5c, 0x4d, 0x72, 0x32, 0x6f, 0x2c, 0x4d, 0x59, 0x19, 0x69, 0x93, 0x25,
	0x5f, 0x07, 0xd8, 0xe9, 0x90, 0xb0, 0xdb, 0x70, 0xfd, 0x36, 0xc9, 0xa6, 0xf2, 0xc6, 0xd2, 0xcc,
	0x4a, 0xf1, 0x58, 0x89, 0x1f, 0xf1, 0x65, 0xeb, 0x7e, 0x9b, 0x58, 0xe9, 0x1d, 0x3d, 0xcc, 0x3d,
	0x81, 0xd9, 0x3e, 0x40, 0x31, 0x27, 0xe3, 0xdd, 0xe8, 0xc9, 0x88, 0xb0, 0x13, 0x25, 0xbd, 0x1b,
	0x90, 0xe8, 0x91, 0x29, 0x42, 0xba, 0x97, 0x0f, 0x01, 0x24, 0x1f, 0x6e, 0x58, 0xeb, 0xb7, 0xaa,
	0x73, 0x27, 0xd0, 0x24, 0x8c, 0xd7, 0xaa, 0xb7, 0x1e, 0xce, 0x19, 0x28, 0x03, 0xa9, 0x9a, 0xb5,
	0x71, 0x77, 0xad, 0x5a, 0x99, 0x4b, 0x98, 0xbf, 0x26, 0x20, 0x63, 0x11, 0xdc, 0xfe, 0x3f, 0xfb,
	0x60, 0x1e, 0x26, 0x18, 0x6e, 0x3a, 0x44, 0x75, 0x82, 0x9c, 0x70, 0xab, 0xed, 0xb5, 0xc9, 0x9e,
	0x3a, 0x52, 0x72, 0xc2, 0xf1, 0xb4, 0x7c, 0xa7, 0xe3, 0x7a, 0xb2, 0xd8, 0x69, 0x4b, 0x4f, 0xd1,
	0x0a, 0xa4, 0xb6, 0x49, 0x97, 0x2b, 0x81, 0x28, 0x57, 0x66, 0xe5, 0x4c, 0x0c, 0x96, 0x07, 0xa4,
	0x5b, 0x27, 0xcc, 0x4a, 0x6e, 0x8b, 0x2f, 0xcf, 0xe1, 0xd8, 0xae, 0xcd, 0xb2, 0x93, 0x79, 0x63,
	0x69, 0xcc, 0x92, 0x93, 0x81, 0xea, 0xa7, 0x07, 0xaa, 0x6f, 0x32, 0x38, 0xbd, 0x4a, 0xb6, 0x6c,
	0x2f, 0xf2, 0x6f, 0x47, 0x33, 0x76, 0x03, 0x52, 0x7e, 0x20, 0xb4, 0x56, 0xb1, 0x75, 0x79, 0x38,
	0x5b, 0x1b, 0x32, 0xd8, 0xd2, 0xab, 0xcc, 0x7f, 0x0c, 0x98, 0xbe, 0xed, 0xbb, 0xae, 0xcd, 0x8e,
	0x4e, 0xb6, 0x08, 0x33, 0x11, 0x8e, 0x1b, 0x76, 0x5b, 0xe4, 0x9c, 0xba, 0x7f, 0xc2, 0x9a, 0x8e,
	0xd8, 0xd7, 0xda, 0xe8, 0x73, 0x58, 0xa0, 0xb6, 0xb7, 0xe5, 0x10, 0x79, 0xb9, 0x47, 0x4a, 0x3a,
	0xf6, 0x06, 0x20, 0xef, 0x9f, 0xb0, 0xe6, 0xe5, 0x36, 0xfc, 0x9e, 0x8f, 0x14, 0xf7, 0x3a, 0xa4,
	0xb5, 0x8a, 0xf3, 0x53, 0xcd, 0xcf, 0xe7, 0xd9, 0x98, 0x1d, 0xd7, 0x55, 0x8c, 0xb5, 0x1f, 0xbd,
	0x3a, 0x7d, 0xa0, 0xc3, 0xcc, 0xa7, 0x30, 0xa3, 0x7f, 0x5e, 0x5d, 0xa3, 0x15, 0x98, 0x6b, 0x09,
	0x4b, 0xa3, 0xf7, 0xd2, 0x10, 0x34, 0x0c, 0x17, 0xa6, 0x59, 0xb9, 0xa6, 0x67, 0x30, 0x2d, 0x98,
	0xb5, 0x7c, 0xc7, 0x69, 0xe2, 0xd6, 0xf6, 0xd1, 0xbc, 0x5e, 0x8e, 0xe7, 0xb5, 0x8f, 0xd5, 0x95,
	0x57, 0x33, 0x90, 0xaa, 0xcb, 0xdf, 0x43, 0x3f, 0xf1, 0xb2, 0x45, 0x1f, 0x14, 0x68, 0x31, 0x86,
	0x81, 0xb8, 0x27, 0x47, 0x6e, 0x88, 0x2c, 0x98, 0x95, 0xaf, 0xff, 0xfa, 0xfb, 0x87, 0xc4, 0x0d,
	0xb3, 0xcc, 0x9f, 0x2f, 0x5f, 0x69, 0x1d, 0xfb, 0x38, 0x08, 0xfd, 0x2f, 0x49, 0x8b, 0xd1, 0x52,
	0xa1, 0x64, 0x7b, 0x94, 0x61, 0xaf, 0x45, 0xf8, 0x58, 0xfb, 0x69, 0xa9, 0xf0, 0xb2, 0xa4, 0x05,
	0xa5, 0x6c, 0x14, 0xd0, 0x77, 0x06, 0xc0, 0xbe, 0xaa, 0xa2, 0x77, 0x62, 0x32, 0x0e, 0x88, 0xee,
	0x50, 0x5c, 0x37, 0x05, 0xae, 0x32, 0xba, 0x26, 0x70, 0x71, 0x8d, 0x39, 0x06, 0xa6, 0x1e, 0xa4,
	0x52, 0xe1, 0x25, 0xfa, 0xc5, 0x80, 0xa9, 0xa8, 0x6e, 0xa2, 0xb8, 0x6b, 0x25, 0x46, 0xdf, 0x73,
	0x8b, 0x47, 0xc6, 0xc9, 0xce, 0x31, 0x57, 0x05, 0xc6, 0x8f, 0xd0, 0x08, 0xdc, 0xa1, 0x57, 0x06,
	0x4c, 0x1f, 0x50, 0xd9, 0xd8, 0xb2, 0xc6, 0xe9, 0x70, 0x6e, 0x61, 0xa0, 0x3d, 0x2b, 0xfc, 0x95,
	0xad, 0xa9, 0x2b, 0x8c, 0x44, 0x1d, 0xec, 0x4b, 0x4e, 0x6c, 0x35, 0x07, 0x14, 0x29, 0x77, 0x2e,
	0x26, 0xca, 0x12, 0x0f, 0xeb, 0x3a, 0x61, 0xe6, 0x23, 0x01, 0xea, 0x81, 0x79, 0x57, 0x80, 0x52,
	0xc9, 0xde, 0x10, 0x57, 0x99, 0xf4, 0x92, 0xf2, 0x9e, 0xfb, 0xc3, 0x80, 0x53, 0x1a, 0x06, 0x0b,
	0x09, 0x76, 0x6d, 0x6f, 0xeb, 0xf8, 0x70, 0x2f, 0xc5, 0x44, 0xd5, 0x70, 0xc8, 0x6c, 0xec, 0xec,
	0xa3, 0xfe, 0x54, 0xa0, 0xde, 0x34, 0x37, 0xfe, 0x0b, 0xd4, 0x11, 0x8c, 0x65, 0xa3, 0xf0, 0x9e,
	0x81, 0xbe, 0x37, 0x60, 0x9c, 0xcb, 0x24, 0x3a, 0x1f, 0x4b, 0x5d, 0x4f, 0x3f, 0x8f, 0xa0, 0xf6,
	0x81, 0x00, 0x59, 0x31, 0x6f, 0x8e, 0x02, 0x32, 0x24, 0xb8, 0xcd, 0x49, 0x7d, 0x6d, 0xc0, 0x74,
	0x0f, 0xe9, 0xb1, 0xc0, 0x1d, 0x8b, 0xc8, 0x4d, 0x81, 0xf1, 0xa1, 0xb9, 0x36, 0x0a, 0x46, 0x1a,
	0xc5, 0x25, 0x29, 0xfc, 0xdd, 0x80, 0xb9, 0x7e, 0x0d, 0x45, 0x85, 0x18, 0x44, 0x87, 0x08, 0x6d,
	0xee, 0xfc, 0x70, 0x61, 0x32, 0x9f, 0x0a, 0xe0, 0x8f, 0xcc, 0xea, 0x28, 0xc0, 0x9b, 0x7d, 0xc9,
	0x39, 0xd1, 0x3f, 0x1b, 0x90, 0x94, 0x4a, 0x84, 0xf2, 0x71, 0x17, 0x79, 0x54, 0xa1, 0x73, 0x17,
	0x87, 0x44, 0xa8, 0xcb, 0x68, 0x5d, 0x00, 0xbd, 0x67, 0xae, 0x8e, 0x02, 0x54, 0x8a, 0x1a, 0x87,
	0xf7, 0xa3, 0x01, 0x93, 0x5a, 0xcf, 0x90, 0x19, 0xd7, 0x02, 0x07, 0xc5, 0xee, 0xd0, 0xdb, 0x68,
	0x43, 0xe0, 0x5a, 0x33, 0xef, 0x8c, 0xd4, 0x9d, 0x2a, 0x59, 0xd9, 0x28, 0xac, 0x7e, 0x06, 0x6f,
	0xb5, 0x7c, 0x77, 0x10, 0xd1, 0xea, 0x94, 0x92, 0xca, 0x1a, 0x07, 0x50, 0x33, 0x7e, 0x4b, 0x9c,
	0xbe, 0x27, 0x63, 0x6e, 0x3b, 0x7e, 0xa7, 0x5d, 0x54, 0xde, 0xe2, 0x93, 0xe5, 0x3f, 0xb5, 0xe7,
	0x99, 0xf0, 0x3c, 0x53, 0x9e, 0x67, 0x4f, 0x96, 0x9b, 0x49, 0x81, 0xfe, 0xfd, 0x7f, 0x03, 0x00,
	0x00, 0xff, 0xff, 0x9c, 0xd8, 0xb6, 0x06, 0x13, 0x11, 0x00, 0x00,
}