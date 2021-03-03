// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/rpc/error_details.proto

package errdetails

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

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
// number of retries have been reached or a maximum retry delay cap has been
// reached.
type RetryInfo struct {
	// Clients should wait at least this long between retrying the same request.
	RetryDelay           *duration.Duration `protobuf:"bytes,1,opt,name=retry_delay,json=retryDelay,proto3" json:"retry_delay,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *RetryInfo) Reset()         { *m = RetryInfo{} }
func (m *RetryInfo) String() string { return proto.CompactTextString(m) }
func (*RetryInfo) ProtoMessage()    {}
func (*RetryInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_851816e4d6b6361a, []int{0}
}

func (m *RetryInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RetryInfo.Unmarshal(m, b)
}
func (m *RetryInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RetryInfo.Marshal(b, m, deterministic)
}
func (m *RetryInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RetryInfo.Merge(m, src)
}
func (m *RetryInfo) XXX_Size() int {
	return xxx_messageInfo_RetryInfo.Size(m)
}
func (m *RetryInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_RetryInfo.DiscardUnknown(m)
}

var xxx_messageInfo_RetryInfo proto.InternalMessageInfo

func (m *RetryInfo) GetRetryDelay() *duration.Duration {
	if m != nil {
		return m.RetryDelay
	}
	return nil
}

// Describes additional debugging info.
type DebugInfo struct {
	// The stack trace entries indicating where the error occurred.
	StackEntries []string `protobuf:"bytes,1,rep,name=stack_entries,json=stackEntries,proto3" json:"stack_entries,omitempty"`
	// Additional debugging information provided by the server.
	Detail               string   `protobuf:"bytes,2,opt,name=detail,proto3" json:"detail,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DebugInfo) Reset()         { *m = DebugInfo{} }
func (m *DebugInfo) String() string { return proto.CompactTextString(m) }
func (*DebugInfo) ProtoMessage()    {}
func (*DebugInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_851816e4d6b6361a, []int{1}
}

func (m *DebugInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DebugInfo.Unmarshal(m, b)
}
func (m *DebugInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DebugInfo.Marshal(b, m, deterministic)
}
func (m *DebugInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DebugInfo.Merge(m, src)
}
func (m *DebugInfo) XXX_Size() int {
	return xxx_messageInfo_DebugInfo.Size(m)
}
func (m *DebugInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_DebugInfo.DiscardUnknown(m)
}

var xxx_messageInfo_DebugInfo proto.InternalMessageInfo

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
// Also see RetryInfo and Help types for other details about handling a
// quota failure.
type QuotaFailure struct {
	// Describes all quota violations.
	Violations           []*QuotaFailure_Violation `protobuf:"bytes,1,rep,name=violations,proto3" json:"violations,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *QuotaFailure) Reset()         { *m = QuotaFailure{} }
func (m *QuotaFailure) String() string { return proto.CompactTextString(m) }
func (*QuotaFailure) ProtoMessage()    {}
func (*QuotaFailure) Descriptor() ([]byte, []int) {
	return fileDescriptor_851816e4d6b6361a, []int{2}
}

func (m *QuotaFailure) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QuotaFailure.Unmarshal(m, b)
}
func (m *QuotaFailure) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QuotaFailure.Marshal(b, m, deterministic)
}
func (m *QuotaFailure) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QuotaFailure.Merge(m, src)
}
func (m *QuotaFailure) XXX_Size() int {
	return xxx_messageInfo_QuotaFailure.Size(m)
}
func (m *QuotaFailure) XXX_DiscardUnknown() {
	xxx_messageInfo_QuotaFailure.DiscardUnknown(m)
}

var xxx_messageInfo_QuotaFailure proto.InternalMessageInfo

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
	Description          string   `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QuotaFailure_Violation) Reset()         { *m = QuotaFailure_Violation{} }
func (m *QuotaFailure_Violation) String() string { return proto.CompactTextString(m) }
func (*QuotaFailure_Violation) ProtoMessage()    {}
func (*QuotaFailure_Violation) Descriptor() ([]byte, []int) {
	return fileDescriptor_851816e4d6b6361a, []int{2, 0}
}

func (m *QuotaFailure_Violation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QuotaFailure_Violation.Unmarshal(m, b)
}
func (m *QuotaFailure_Violation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QuotaFailure_Violation.Marshal(b, m, deterministic)
}
func (m *QuotaFailure_Violation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QuotaFailure_Violation.Merge(m, src)
}
func (m *QuotaFailure_Violation) XXX_Size() int {
	return xxx_messageInfo_QuotaFailure_Violation.Size(m)
}
func (m *QuotaFailure_Violation) XXX_DiscardUnknown() {
	xxx_messageInfo_QuotaFailure_Violation.DiscardUnknown(m)
}

var xxx_messageInfo_QuotaFailure_Violation proto.InternalMessageInfo

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

// Describes the cause of the error with structured details.
//
// Example of an error when contacting the "pubsub.googleapis.com" API when it
// is not enabled:
//     { "type":   "API_DISABLED"
//       "domain": "googleapis.com"
//       "metadata": {
//         "resource": "projects/123",
//         "service": "pubsub.googleapis.com"
//       }
//     }
// This response indicates that the pubsub.googleapis.com API is not enabled.
//
// Example of an error that is returned when attempting to create a Spanner
// instance in a region that is out of stock:
//     { "type":   "STOCKOUT"
//       "domain": "spanner.googleapis.com",
//       "metadata": {
//         "availableRegions": ""us-central1,us-east2"
//       }
//     }
//
type ErrorInfo struct {
	// The reason of the error. This is a constant value that identifies the
	// proximate cause of the error. Error reasons are unique within a particular
	// domain of errors. This should be at most 63 characters and match
	// /[A-Z0-9_]+/.
	Reason string `protobuf:"bytes,1,opt,name=reason,proto3" json:"reason,omitempty"`
	// The logical grouping to which the "reason" belongs.  Often "domain" will
	// contain the registered service name of the tool or product that is the
	// source of the error. Example: "pubsub.googleapis.com". If the error is
	// common across many APIs, the first segment of the example above will be
	// omitted.  The value will be, "googleapis.com".
	Domain string `protobuf:"bytes,2,opt,name=domain,proto3" json:"domain,omitempty"`
	// Additional structured details about this error.
	//
	// Keys should match /[a-zA-Z0-9-_]/ and be limited to 64 characters in
	// length. When identifying the current value of an exceeded limit, the units
	// should be contained in the key, not the value.  For example, rather than
	// {"instanceLimit": "100/request"}, should be returned as,
	// {"instanceLimitPerRequest": "100"}, if the client exceeds the number of
	// instances that can be created in a single (batch) request.
	Metadata             map[string]string `protobuf:"bytes,3,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ErrorInfo) Reset()         { *m = ErrorInfo{} }
func (m *ErrorInfo) String() string { return proto.CompactTextString(m) }
func (*ErrorInfo) ProtoMessage()    {}
func (*ErrorInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_851816e4d6b6361a, []int{3}
}

func (m *ErrorInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ErrorInfo.Unmarshal(m, b)
}
func (m *ErrorInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ErrorInfo.Marshal(b, m, deterministic)
}
func (m *ErrorInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ErrorInfo.Merge(m, src)
}
func (m *ErrorInfo) XXX_Size() int {
	return xxx_messageInfo_ErrorInfo.Size(m)
}
func (m *ErrorInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ErrorInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ErrorInfo proto.InternalMessageInfo

func (m *ErrorInfo) GetReason() string {
	if m != nil {
		return m.Reason
	}
	return ""
}

func (m *ErrorInfo) GetDomain() string {
	if m != nil {
		return m.Domain
	}
	return ""
}

func (m *ErrorInfo) GetMetadata() map[string]string {
	if m != nil {
		return m.Metadata
	}
	return nil
}

// Describes what preconditions have failed.
//
// For example, if an RPC failed because it required the Terms of Service to be
// acknowledged, it could list the terms of service violation in the
// PreconditionFailure message.
type PreconditionFailure struct {
	// Describes all precondition violations.
	Violations           []*PreconditionFailure_Violation `protobuf:"bytes,1,rep,name=violations,proto3" json:"violations,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                         `json:"-"`
	XXX_unrecognized     []byte                           `json:"-"`
	XXX_sizecache        int32                            `json:"-"`
}

func (m *PreconditionFailure) Reset()         { *m = PreconditionFailure{} }
func (m *PreconditionFailure) String() string { return proto.CompactTextString(m) }
func (*PreconditionFailure) ProtoMessage()    {}
func (*PreconditionFailure) Descriptor() ([]byte, []int) {
	return fileDescriptor_851816e4d6b6361a, []int{4}
}

func (m *PreconditionFailure) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PreconditionFailure.Unmarshal(m, b)
}
func (m *PreconditionFailure) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PreconditionFailure.Marshal(b, m, deterministic)
}
func (m *PreconditionFailure) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PreconditionFailure.Merge(m, src)
}
func (m *PreconditionFailure) XXX_Size() int {
	return xxx_messageInfo_PreconditionFailure.Size(m)
}
func (m *PreconditionFailure) XXX_DiscardUnknown() {
	xxx_messageInfo_PreconditionFailure.DiscardUnknown(m)
}

var xxx_messageInfo_PreconditionFailure proto.InternalMessageInfo

func (m *PreconditionFailure) GetViolations() []*PreconditionFailure_Violation {
	if m != nil {
		return m.Violations
	}
	return nil
}

// A message type used to describe a single precondition failure.
type PreconditionFailure_Violation struct {
	// The type of PreconditionFailure. We recommend using a service-specific
	// enum type to define the supported precondition violation subjects. For
	// example, "TOS" for "Terms of Service violation".
	Type string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	// The subject, relative to the type, that failed.
	// For example, "google.com/cloud" relative to the "TOS" type would indicate
	// which terms of service is being referenced.
	Subject string `protobuf:"bytes,2,opt,name=subject,proto3" json:"subject,omitempty"`
	// A description of how the precondition failed. Developers can use this
	// description to understand how to fix the failure.
	//
	// For example: "Terms of service not accepted".
	Description          string   `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PreconditionFailure_Violation) Reset()         { *m = PreconditionFailure_Violation{} }
func (m *PreconditionFailure_Violation) String() string { return proto.CompactTextString(m) }
func (*PreconditionFailure_Violation) ProtoMessage()    {}
func (*PreconditionFailure_Violation) Descriptor() ([]byte, []int) {
	return fileDescriptor_851816e4d6b6361a, []int{4, 0}
}

func (m *PreconditionFailure_Violation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PreconditionFailure_Violation.Unmarshal(m, b)
}
func (m *PreconditionFailure_Violation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PreconditionFailure_Violation.Marshal(b, m, deterministic)
}
func (m *PreconditionFailure_Violation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PreconditionFailure_Violation.Merge(m, src)
}
func (m *PreconditionFailure_Violation) XXX_Size() int {
	return xxx_messageInfo_PreconditionFailure_Violation.Size(m)
}
func (m *PreconditionFailure_Violation) XXX_DiscardUnknown() {
	xxx_messageInfo_PreconditionFailure_Violation.DiscardUnknown(m)
}

var xxx_messageInfo_PreconditionFailure_Violation proto.InternalMessageInfo

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
	FieldViolations      []*BadRequest_FieldViolation `protobuf:"bytes,1,rep,name=field_violations,json=fieldViolations,proto3" json:"field_violations,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *BadRequest) Reset()         { *m = BadRequest{} }
func (m *BadRequest) String() string { return proto.CompactTextString(m) }
func (*BadRequest) ProtoMessage()    {}
func (*BadRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_851816e4d6b6361a, []int{5}
}

func (m *BadRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BadRequest.Unmarshal(m, b)
}
func (m *BadRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BadRequest.Marshal(b, m, deterministic)
}
func (m *BadRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BadRequest.Merge(m, src)
}
func (m *BadRequest) XXX_Size() int {
	return xxx_messageInfo_BadRequest.Size(m)
}
func (m *BadRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BadRequest.DiscardUnknown(m)
}

var xxx_messageInfo_BadRequest proto.InternalMessageInfo

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
	Description          string   `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BadRequest_FieldViolation) Reset()         { *m = BadRequest_FieldViolation{} }
func (m *BadRequest_FieldViolation) String() string { return proto.CompactTextString(m) }
func (*BadRequest_FieldViolation) ProtoMessage()    {}
func (*BadRequest_FieldViolation) Descriptor() ([]byte, []int) {
	return fileDescriptor_851816e4d6b6361a, []int{5, 0}
}

func (m *BadRequest_FieldViolation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BadRequest_FieldViolation.Unmarshal(m, b)
}
func (m *BadRequest_FieldViolation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BadRequest_FieldViolation.Marshal(b, m, deterministic)
}
func (m *BadRequest_FieldViolation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BadRequest_FieldViolation.Merge(m, src)
}
func (m *BadRequest_FieldViolation) XXX_Size() int {
	return xxx_messageInfo_BadRequest_FieldViolation.Size(m)
}
func (m *BadRequest_FieldViolation) XXX_DiscardUnknown() {
	xxx_messageInfo_BadRequest_FieldViolation.DiscardUnknown(m)
}

var xxx_messageInfo_BadRequest_FieldViolation proto.InternalMessageInfo

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
	ServingData          string   `protobuf:"bytes,2,opt,name=serving_data,json=servingData,proto3" json:"serving_data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestInfo) Reset()         { *m = RequestInfo{} }
func (m *RequestInfo) String() string { return proto.CompactTextString(m) }
func (*RequestInfo) ProtoMessage()    {}
func (*RequestInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_851816e4d6b6361a, []int{6}
}

func (m *RequestInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestInfo.Unmarshal(m, b)
}
func (m *RequestInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestInfo.Marshal(b, m, deterministic)
}
func (m *RequestInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestInfo.Merge(m, src)
}
func (m *RequestInfo) XXX_Size() int {
	return xxx_messageInfo_RequestInfo.Size(m)
}
func (m *RequestInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestInfo.DiscardUnknown(m)
}

var xxx_messageInfo_RequestInfo proto.InternalMessageInfo

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
	Description          string   `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResourceInfo) Reset()         { *m = ResourceInfo{} }
func (m *ResourceInfo) String() string { return proto.CompactTextString(m) }
func (*ResourceInfo) ProtoMessage()    {}
func (*ResourceInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_851816e4d6b6361a, []int{7}
}

func (m *ResourceInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResourceInfo.Unmarshal(m, b)
}
func (m *ResourceInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResourceInfo.Marshal(b, m, deterministic)
}
func (m *ResourceInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResourceInfo.Merge(m, src)
}
func (m *ResourceInfo) XXX_Size() int {
	return xxx_messageInfo_ResourceInfo.Size(m)
}
func (m *ResourceInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ResourceInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ResourceInfo proto.InternalMessageInfo

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
	Links                []*Help_Link `protobuf:"bytes,1,rep,name=links,proto3" json:"links,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Help) Reset()         { *m = Help{} }
func (m *Help) String() string { return proto.CompactTextString(m) }
func (*Help) ProtoMessage()    {}
func (*Help) Descriptor() ([]byte, []int) {
	return fileDescriptor_851816e4d6b6361a, []int{8}
}

func (m *Help) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Help.Unmarshal(m, b)
}
func (m *Help) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Help.Marshal(b, m, deterministic)
}
func (m *Help) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Help.Merge(m, src)
}
func (m *Help) XXX_Size() int {
	return xxx_messageInfo_Help.Size(m)
}
func (m *Help) XXX_DiscardUnknown() {
	xxx_messageInfo_Help.DiscardUnknown(m)
}

var xxx_messageInfo_Help proto.InternalMessageInfo

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
	Url                  string   `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Help_Link) Reset()         { *m = Help_Link{} }
func (m *Help_Link) String() string { return proto.CompactTextString(m) }
func (*Help_Link) ProtoMessage()    {}
func (*Help_Link) Descriptor() ([]byte, []int) {
	return fileDescriptor_851816e4d6b6361a, []int{8, 0}
}

func (m *Help_Link) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Help_Link.Unmarshal(m, b)
}
func (m *Help_Link) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Help_Link.Marshal(b, m, deterministic)
}
func (m *Help_Link) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Help_Link.Merge(m, src)
}
func (m *Help_Link) XXX_Size() int {
	return xxx_messageInfo_Help_Link.Size(m)
}
func (m *Help_Link) XXX_DiscardUnknown() {
	xxx_messageInfo_Help_Link.DiscardUnknown(m)
}

var xxx_messageInfo_Help_Link proto.InternalMessageInfo

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
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LocalizedMessage) Reset()         { *m = LocalizedMessage{} }
func (m *LocalizedMessage) String() string { return proto.CompactTextString(m) }
func (*LocalizedMessage) ProtoMessage()    {}
func (*LocalizedMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_851816e4d6b6361a, []int{9}
}

func (m *LocalizedMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LocalizedMessage.Unmarshal(m, b)
}
func (m *LocalizedMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LocalizedMessage.Marshal(b, m, deterministic)
}
func (m *LocalizedMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LocalizedMessage.Merge(m, src)
}
func (m *LocalizedMessage) XXX_Size() int {
	return xxx_messageInfo_LocalizedMessage.Size(m)
}
func (m *LocalizedMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_LocalizedMessage.DiscardUnknown(m)
}

var xxx_messageInfo_LocalizedMessage proto.InternalMessageInfo

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
	proto.RegisterType((*ErrorInfo)(nil), "google.rpc.ErrorInfo")
	proto.RegisterMapType((map[string]string)(nil), "google.rpc.ErrorInfo.MetadataEntry")
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

func init() {
	proto.RegisterFile("google/rpc/error_details.proto", fileDescriptor_851816e4d6b6361a)
}

var fileDescriptor_851816e4d6b6361a = []byte{
	// 674 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0x95, 0x93, 0xb4, 0xe0, 0x49, 0x5a, 0x8a, 0x81, 0x2a, 0x44, 0x02, 0x05, 0x57, 0x48, 0x45,
	0x48, 0x8e, 0x54, 0x2e, 0xa8, 0x3d, 0x54, 0x0a, 0xe9, 0x97, 0xd4, 0x42, 0xb0, 0x10, 0x07, 0x38,
	0x44, 0x1b, 0x7b, 0x62, 0x2d, 0x71, 0xbc, 0x66, 0xbd, 0x2e, 0x0a, 0xbf, 0x82, 0x3b, 0x37, 0x4e,
	0xfc, 0x05, 0x0e, 0xfc, 0x37, 0xb4, 0x5f, 0x89, 0xdb, 0x14, 0xc4, 0x6d, 0xdf, 0xec, 0xdb, 0xb7,
	0xf3, 0x9e, 0x66, 0x17, 0x1e, 0x27, 0x8c, 0x25, 0x29, 0xf6, 0x78, 0x1e, 0xf5, 0x90, 0x73, 0xc6,
	0x47, 0x31, 0x0a, 0x42, 0xd3, 0x22, 0xc8, 0x39, 0x13, 0xcc, 0x03, 0xbd, 0x1f, 0xf0, 0x3c, 0xea,
	0x58, 0xae, 0xda, 0x19, 0x97, 0x93, 0x5e, 0x5c, 0x72, 0x22, 0x28, 0xcb, 0x34, 0xd7, 0x3f, 0x01,
	0x37, 0x44, 0xc1, 0xe7, 0x67, 0xd9, 0x84, 0x79, 0xfb, 0xd0, 0xe4, 0x12, 0x8c, 0x62, 0x4c, 0xc9,
	0xbc, 0xed, 0x74, 0x9d, 0xdd, 0xe6, 0xde, 0xc3, 0xc0, 0xc8, 0x59, 0x89, 0x60, 0x60, 0x24, 0x42,
	0x50, 0xec, 0x81, 0x24, 0xfb, 0xa7, 0xe0, 0x0e, 0x70, 0x5c, 0x26, 0x4a, 0x68, 0x07, 0x36, 0x0a,
	0x41, 0xa2, 0xe9, 0x08, 0x33, 0xc1, 0x29, 0x16, 0x6d, 0xa7, 0x5b, 0xdf, 0x75, 0xc3, 0x96, 0x2a,
	0x1e, 0xe9, 0x9a, 0xb7, 0x0d, 0xeb, 0xba, 0xef, 0x76, 0xad, 0xeb, 0xec, 0xba, 0xa1, 0x41, 0xfe,
	0x77, 0x07, 0x5a, 0x6f, 0x4b, 0x26, 0xc8, 0x31, 0xa1, 0x69, 0xc9, 0xd1, 0xeb, 0x03, 0x5c, 0x52,
	0x96, 0xaa, 0x3b, 0xb5, 0x54, 0x73, 0xcf, 0x0f, 0x96, 0x26, 0x83, 0x2a, 0x3b, 0x78, 0x6f, 0xa9,
	0x61, 0xe5, 0x54, 0xe7, 0x04, 0xdc, 0xc5, 0x86, 0xd7, 0x86, 0x5b, 0x45, 0x39, 0xfe, 0x84, 0x91,
	0x50, 0x1e, 0xdd, 0xd0, 0x42, 0xaf, 0x0b, 0xcd, 0x18, 0x8b, 0x88, 0xd3, 0x5c, 0x12, 0x4d, 0x63,
	0xd5, 0x92, 0xff, 0xcb, 0x01, 0xf7, 0x48, 0x86, 0xae, 0x8c, 0x6e, 0xc3, 0x3a, 0x47, 0x52, 0xb0,
	0xcc, 0x08, 0x19, 0xa4, 0xbc, 0xb1, 0x19, 0xa1, 0xd9, 0xc2, 0x9b, 0x42, 0xde, 0x21, 0xdc, 0x9e,
	0xa1, 0x20, 0x31, 0x11, 0xa4, 0x5d, 0x57, 0x46, 0x76, 0xaa, 0x46, 0x16, 0xc2, 0xc1, 0x85, 0x61,
	0xc9, 0xb0, 0xe6, 0xe1, 0xe2, 0x50, 0xe7, 0x00, 0x36, 0xae, 0x6c, 0x79, 0x5b, 0x50, 0x9f, 0xe2,
	0xdc, 0x5c, 0x2f, 0x97, 0xde, 0x7d, 0x58, 0xbb, 0x24, 0x69, 0x89, 0xe6, 0x6a, 0x0d, 0xf6, 0x6b,
	0x2f, 0x1d, 0xff, 0xb7, 0x03, 0xf7, 0x86, 0x1c, 0x23, 0x96, 0xc5, 0x54, 0x9a, 0xb1, 0x01, 0x9f,
	0xdd, 0x10, 0xf0, 0xb3, 0x6a, 0x5f, 0x37, 0x1c, 0xfa, 0x4b, 0xce, 0x1f, 0xab, 0x39, 0x7b, 0xd0,
	0x10, 0xf3, 0x1c, 0x4d, 0x73, 0x6a, 0x5d, 0xcd, 0xbe, 0xf6, 0xcf, 0xec, 0xeb, 0xab, 0xd9, 0xff,
	0x74, 0x00, 0xfa, 0x24, 0x0e, 0xf1, 0x73, 0x89, 0x85, 0xf0, 0x86, 0xb0, 0x35, 0xa1, 0x98, 0xc6,
	0xa3, 0x95, 0xe6, 0x9f, 0x56, 0x9b, 0x5f, 0x9e, 0x08, 0x8e, 0x25, 0x7d, 0xd9, 0xf8, 0x9d, 0xc9,
	0x15, 0x5c, 0x74, 0x4e, 0x61, 0xf3, 0x2a, 0x45, 0x86, 0xa9, 0x48, 0xc6, 0x83, 0x06, 0xff, 0x31,
	0x26, 0x6f, 0xa0, 0x69, 0x2e, 0x55, 0x73, 0xf2, 0x08, 0x80, 0x6b, 0x38, 0xa2, 0x56, 0xcb, 0x35,
	0x95, 0xb3, 0xd8, 0x7b, 0x02, 0xad, 0x02, 0xf9, 0x25, 0xcd, 0x92, 0x91, 0x1a, 0x0d, 0x23, 0x68,
	0x6a, 0x03, 0x22, 0x88, 0xff, 0xcd, 0x81, 0x56, 0x88, 0x05, 0x2b, 0x79, 0x84, 0xf6, 0x8d, 0x71,
	0x83, 0x47, 0x95, 0x94, 0x5b, 0xb6, 0xf8, 0x4e, 0xa6, 0x5d, 0x25, 0x65, 0x64, 0x66, 0x67, 0x62,
	0x41, 0x7a, 0x4d, 0x66, 0x28, 0x3d, 0xb2, 0x2f, 0x19, 0x72, 0x13, 0xb9, 0x06, 0xd7, 0x3d, 0x36,
	0x56, 0x3d, 0x32, 0x68, 0x9c, 0x62, 0x9a, 0x7b, 0xcf, 0x61, 0x2d, 0xa5, 0xd9, 0xd4, 0x86, 0xff,
	0xa0, 0x1a, 0xbe, 0x24, 0x04, 0xe7, 0x34, 0x9b, 0x86, 0x9a, 0xd3, 0xd9, 0x87, 0x86, 0x84, 0xd7,
	0xe5, 0x9d, 0x15, 0x79, 0x39, 0xd9, 0x25, 0xb7, 0x9f, 0x83, 0x5c, 0xfa, 0x03, 0xd8, 0x3a, 0x67,
	0x11, 0x49, 0xe9, 0x57, 0x8c, 0x2f, 0xb0, 0x28, 0x48, 0x82, 0xf2, 0xa5, 0xa5, 0xb2, 0x66, 0xfd,
	0x1b, 0x24, 0xe7, 0x6c, 0xa6, 0x29, 0x76, 0xce, 0x0c, 0xec, 0xa7, 0xb0, 0x19, 0xb1, 0x59, 0xa5,
	0xc9, 0xfe, 0x5d, 0xf5, 0xee, 0x06, 0xfa, 0x13, 0x1d, 0xca, 0x6f, 0x6e, 0xe8, 0x7c, 0x38, 0x34,
	0x84, 0x84, 0xa5, 0x24, 0x4b, 0x02, 0xc6, 0x93, 0x5e, 0x82, 0x99, 0xfa, 0x04, 0x7b, 0x7a, 0x8b,
	0xe4, 0xb4, 0xb0, 0x9f, 0xb0, 0xf9, 0x81, 0x0f, 0x96, 0xcb, 0x1f, 0xb5, 0x7a, 0x38, 0x7c, 0x35,
	0x5e, 0x57, 0x27, 0x5e, 0xfc, 0x09, 0x00, 0x00, 0xff, 0xff, 0x2b, 0x8c, 0x9a, 0x65, 0xb5, 0x05,
	0x00, 0x00,
}
