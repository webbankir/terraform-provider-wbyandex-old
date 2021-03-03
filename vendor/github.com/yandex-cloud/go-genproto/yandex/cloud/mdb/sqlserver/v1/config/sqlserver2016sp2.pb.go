// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.6.1
// source: yandex/cloud/mdb/sqlserver/v1/config/sqlserver2016sp2.proto

package sqlserver

import (
	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	_ "github.com/yandex-cloud/go-genproto/yandex/cloud"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// SQL Server 2016 SP2 Standard edition supported configuration options are listed here.
//
// Detailed description for each set of options is available in [SQL Server documentation](https://docs.microsoft.com/en-us/sql/database-engine/configure-windows/server-configuration-options-sql-server?view=sql-server-2016).
//
// Any options that are not listed here are not supported.
type SQLServerConfig2016Sp2Std struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Limits the number of processors to use in parallel plan execution per task.
	//
	// See in-depth description in [SQL Server documentation](https://docs.microsoft.com/en-us/sql/database-engine/configure-windows/configure-the-max-degree-of-parallelism-server-configuration-option?view=sql-server-2016).
	MaxDegreeOfParallelism *wrappers.Int64Value `protobuf:"bytes,1,opt,name=max_degree_of_parallelism,json=maxDegreeOfParallelism,proto3" json:"max_degree_of_parallelism,omitempty"`
	// Specifies the threshold at which SQL Server creates and runs parallel plans for queries.
	//
	// SQL Server creates and runs a parallel plan for a query only when the estimated cost to run a serial plan for the same query is higher than the value of the option.
	//
	// See in-depth description in [SQL Server documentation](https://docs.microsoft.com/en-us/sql/database-engine/configure-windows/configure-the-cost-threshold-for-parallelism-server-configuration-option?view=sql-server-2016).
	CostThresholdForParallelism *wrappers.Int64Value `protobuf:"bytes,2,opt,name=cost_threshold_for_parallelism,json=costThresholdForParallelism,proto3" json:"cost_threshold_for_parallelism,omitempty"`
	// Describes how to configure login auditing to monitor SQL Server Database Engine login activity.
	// Possible values:
	// * 0 - do not log login attempts,
	// * 1 - log only failed login attempts,
	// * 2 - log only successful login attempts (not recommended),
	// * 3 - log all login attempts (not recommended).
	//
	// See in-depth description in [SQL Server documentation](https://docs.microsoft.com/en-us/sql/ssms/configure-login-auditing-sql-server-management-studio?view=sql-server-2016).
	AuditLevel *wrappers.Int64Value `protobuf:"bytes,3,opt,name=audit_level,json=auditLevel,proto3" json:"audit_level,omitempty"`
	// Manages the fill factor server configuration option.
	// When an index is created or rebuilt the fill factor determines the percentage of space on each index leaf-level page to be filled with data, reserving the rest as free space for future growth.
	//
	// Values 0 and 100 mean full page usage (no space reserved).
	//
	// See in-depth description in [SQL Server documentation](https://docs.microsoft.com/en-us/sql/database-engine/configure-windows/configure-the-fill-factor-server-configuration-option?view=sql-server-2016).
	FillFactorPercent *wrappers.Int64Value `protobuf:"bytes,4,opt,name=fill_factor_percent,json=fillFactorPercent,proto3" json:"fill_factor_percent,omitempty"`
	// Determines whether plans should be cached only after second execution.
	// Allows to avoid SQL cache bloat because of single-use plans.
	//
	// See in-depth description in [SQL Server documentation](https://docs.microsoft.com/en-us/sql/database-engine/configure-windows/optimize-for-ad-hoc-workloads-server-configuration-option?view=sql-server-2016).
	OptimizeForAdHocWorkloads *wrappers.BoolValue `protobuf:"bytes,5,opt,name=optimize_for_ad_hoc_workloads,json=optimizeForAdHocWorkloads,proto3" json:"optimize_for_ad_hoc_workloads,omitempty"`
}

func (x *SQLServerConfig2016Sp2Std) Reset() {
	*x = SQLServerConfig2016Sp2Std{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_mdb_sqlserver_v1_config_sqlserver2016sp2_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SQLServerConfig2016Sp2Std) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SQLServerConfig2016Sp2Std) ProtoMessage() {}

func (x *SQLServerConfig2016Sp2Std) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_mdb_sqlserver_v1_config_sqlserver2016sp2_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SQLServerConfig2016Sp2Std.ProtoReflect.Descriptor instead.
func (*SQLServerConfig2016Sp2Std) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_mdb_sqlserver_v1_config_sqlserver2016sp2_proto_rawDescGZIP(), []int{0}
}

func (x *SQLServerConfig2016Sp2Std) GetMaxDegreeOfParallelism() *wrappers.Int64Value {
	if x != nil {
		return x.MaxDegreeOfParallelism
	}
	return nil
}

func (x *SQLServerConfig2016Sp2Std) GetCostThresholdForParallelism() *wrappers.Int64Value {
	if x != nil {
		return x.CostThresholdForParallelism
	}
	return nil
}

func (x *SQLServerConfig2016Sp2Std) GetAuditLevel() *wrappers.Int64Value {
	if x != nil {
		return x.AuditLevel
	}
	return nil
}

func (x *SQLServerConfig2016Sp2Std) GetFillFactorPercent() *wrappers.Int64Value {
	if x != nil {
		return x.FillFactorPercent
	}
	return nil
}

func (x *SQLServerConfig2016Sp2Std) GetOptimizeForAdHocWorkloads() *wrappers.BoolValue {
	if x != nil {
		return x.OptimizeForAdHocWorkloads
	}
	return nil
}

type SQLServerConfigSet2016Sp2Std struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Effective settings for an SQL Server 2016 SP2 cluster (a combination of settings defined in [user_config] and [default_config]).
	EffectiveConfig *SQLServerConfig2016Sp2Std `protobuf:"bytes,1,opt,name=effective_config,json=effectiveConfig,proto3" json:"effective_config,omitempty"`
	// User-defined settings for an SQL Server 2016 SP2 cluster.
	UserConfig *SQLServerConfig2016Sp2Std `protobuf:"bytes,2,opt,name=user_config,json=userConfig,proto3" json:"user_config,omitempty"`
	// Default configuration for an SQL Server 2016 SP2 cluster.
	DefaultConfig *SQLServerConfig2016Sp2Std `protobuf:"bytes,3,opt,name=default_config,json=defaultConfig,proto3" json:"default_config,omitempty"`
}

func (x *SQLServerConfigSet2016Sp2Std) Reset() {
	*x = SQLServerConfigSet2016Sp2Std{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_mdb_sqlserver_v1_config_sqlserver2016sp2_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SQLServerConfigSet2016Sp2Std) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SQLServerConfigSet2016Sp2Std) ProtoMessage() {}

func (x *SQLServerConfigSet2016Sp2Std) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_mdb_sqlserver_v1_config_sqlserver2016sp2_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SQLServerConfigSet2016Sp2Std.ProtoReflect.Descriptor instead.
func (*SQLServerConfigSet2016Sp2Std) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_mdb_sqlserver_v1_config_sqlserver2016sp2_proto_rawDescGZIP(), []int{1}
}

func (x *SQLServerConfigSet2016Sp2Std) GetEffectiveConfig() *SQLServerConfig2016Sp2Std {
	if x != nil {
		return x.EffectiveConfig
	}
	return nil
}

func (x *SQLServerConfigSet2016Sp2Std) GetUserConfig() *SQLServerConfig2016Sp2Std {
	if x != nil {
		return x.UserConfig
	}
	return nil
}

func (x *SQLServerConfigSet2016Sp2Std) GetDefaultConfig() *SQLServerConfig2016Sp2Std {
	if x != nil {
		return x.DefaultConfig
	}
	return nil
}

// SQL Server 2016 SP2 Enterprise edition supported configuration options are listed here.
//
// Detailed description for each set of options is available in [SQL Server documentation](https://docs.microsoft.com/en-us/sql/database-engine/configure-windows/server-configuration-options-sql-server?view=sql-server-2016).
//
// Any options that are not listed here are not supported.
type SQLServerConfig2016Sp2Ent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Limits the number of processors to use in parallel plan execution per task.
	//
	// See in-depth description in [SQL Server documentation](https://docs.microsoft.com/en-us/sql/database-engine/configure-windows/configure-the-max-degree-of-parallelism-server-configuration-option?view=sql-server-2016).
	MaxDegreeOfParallelism *wrappers.Int64Value `protobuf:"bytes,1,opt,name=max_degree_of_parallelism,json=maxDegreeOfParallelism,proto3" json:"max_degree_of_parallelism,omitempty"`
	// Specifies the threshold at which SQL Server creates and runs parallel plans for queries.
	//
	// SQL Server creates and runs a parallel plan for a query only when the estimated cost to run a serial plan for the same query is higher than the value of the option.
	//
	// See in-depth description in [SQL Server documentation](https://docs.microsoft.com/en-us/sql/database-engine/configure-windows/configure-the-cost-threshold-for-parallelism-server-configuration-option?view=sql-server-2016).
	CostThresholdForParallelism *wrappers.Int64Value `protobuf:"bytes,2,opt,name=cost_threshold_for_parallelism,json=costThresholdForParallelism,proto3" json:"cost_threshold_for_parallelism,omitempty"`
	// Describes how to configure login auditing to monitor SQL Server Database Engine login activity.
	// Possible values:
	// * 0 - do not log login attempts,
	// * 1 - log only failed login attempts,
	// * 2 - log only successful login attempts (not recommended),
	// * 3 - log all login attempts (not recommended).
	//
	// See in-depth description in [SQL Server documentation](https://docs.microsoft.com/en-us/sql/ssms/configure-login-auditing-sql-server-management-studio?view=sql-server-2016).
	AuditLevel *wrappers.Int64Value `protobuf:"bytes,3,opt,name=audit_level,json=auditLevel,proto3" json:"audit_level,omitempty"`
	// Manages the fill factor server configuration option.
	// When an index is created or rebuilt the fill factor determines the percentage of space on each index leaf-level page to be filled with data, reserving the rest as free space for future growth.
	//
	// Values 0 and 100 mean full page usage (no space reserved).
	//
	// See in-depth description in [SQL Server documentation](https://docs.microsoft.com/en-us/sql/database-engine/configure-windows/configure-the-fill-factor-server-configuration-option?view=sql-server-2016).
	FillFactorPercent *wrappers.Int64Value `protobuf:"bytes,4,opt,name=fill_factor_percent,json=fillFactorPercent,proto3" json:"fill_factor_percent,omitempty"`
	// Determines whether plans should be cached only after second execution.
	// Allows to avoid SQL cache bloat because of single-use plans.
	//
	// See in-depth description in [SQL Server documentation](https://docs.microsoft.com/en-us/sql/database-engine/configure-windows/optimize-for-ad-hoc-workloads-server-configuration-option?view=sql-server-2016).
	OptimizeForAdHocWorkloads *wrappers.BoolValue `protobuf:"bytes,5,opt,name=optimize_for_ad_hoc_workloads,json=optimizeForAdHocWorkloads,proto3" json:"optimize_for_ad_hoc_workloads,omitempty"`
}

func (x *SQLServerConfig2016Sp2Ent) Reset() {
	*x = SQLServerConfig2016Sp2Ent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_mdb_sqlserver_v1_config_sqlserver2016sp2_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SQLServerConfig2016Sp2Ent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SQLServerConfig2016Sp2Ent) ProtoMessage() {}

func (x *SQLServerConfig2016Sp2Ent) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_mdb_sqlserver_v1_config_sqlserver2016sp2_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SQLServerConfig2016Sp2Ent.ProtoReflect.Descriptor instead.
func (*SQLServerConfig2016Sp2Ent) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_mdb_sqlserver_v1_config_sqlserver2016sp2_proto_rawDescGZIP(), []int{2}
}

func (x *SQLServerConfig2016Sp2Ent) GetMaxDegreeOfParallelism() *wrappers.Int64Value {
	if x != nil {
		return x.MaxDegreeOfParallelism
	}
	return nil
}

func (x *SQLServerConfig2016Sp2Ent) GetCostThresholdForParallelism() *wrappers.Int64Value {
	if x != nil {
		return x.CostThresholdForParallelism
	}
	return nil
}

func (x *SQLServerConfig2016Sp2Ent) GetAuditLevel() *wrappers.Int64Value {
	if x != nil {
		return x.AuditLevel
	}
	return nil
}

func (x *SQLServerConfig2016Sp2Ent) GetFillFactorPercent() *wrappers.Int64Value {
	if x != nil {
		return x.FillFactorPercent
	}
	return nil
}

func (x *SQLServerConfig2016Sp2Ent) GetOptimizeForAdHocWorkloads() *wrappers.BoolValue {
	if x != nil {
		return x.OptimizeForAdHocWorkloads
	}
	return nil
}

type SQLServerConfigSet2016Sp2Ent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Effective settings for an SQL Server 2016 SP2 cluster (a combination of settings defined in [user_config] and [default_config]).
	EffectiveConfig *SQLServerConfig2016Sp2Ent `protobuf:"bytes,1,opt,name=effective_config,json=effectiveConfig,proto3" json:"effective_config,omitempty"`
	// User-defined settings for an SQL Server 2016 SP2 cluster.
	UserConfig *SQLServerConfig2016Sp2Ent `protobuf:"bytes,2,opt,name=user_config,json=userConfig,proto3" json:"user_config,omitempty"`
	// Default configuration for an SQL Server 2016 SP2 cluster.
	DefaultConfig *SQLServerConfig2016Sp2Ent `protobuf:"bytes,3,opt,name=default_config,json=defaultConfig,proto3" json:"default_config,omitempty"`
}

func (x *SQLServerConfigSet2016Sp2Ent) Reset() {
	*x = SQLServerConfigSet2016Sp2Ent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_mdb_sqlserver_v1_config_sqlserver2016sp2_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SQLServerConfigSet2016Sp2Ent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SQLServerConfigSet2016Sp2Ent) ProtoMessage() {}

func (x *SQLServerConfigSet2016Sp2Ent) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_mdb_sqlserver_v1_config_sqlserver2016sp2_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SQLServerConfigSet2016Sp2Ent.ProtoReflect.Descriptor instead.
func (*SQLServerConfigSet2016Sp2Ent) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_mdb_sqlserver_v1_config_sqlserver2016sp2_proto_rawDescGZIP(), []int{3}
}

func (x *SQLServerConfigSet2016Sp2Ent) GetEffectiveConfig() *SQLServerConfig2016Sp2Ent {
	if x != nil {
		return x.EffectiveConfig
	}
	return nil
}

func (x *SQLServerConfigSet2016Sp2Ent) GetUserConfig() *SQLServerConfig2016Sp2Ent {
	if x != nil {
		return x.UserConfig
	}
	return nil
}

func (x *SQLServerConfigSet2016Sp2Ent) GetDefaultConfig() *SQLServerConfig2016Sp2Ent {
	if x != nil {
		return x.DefaultConfig
	}
	return nil
}

var File_yandex_cloud_mdb_sqlserver_v1_config_sqlserver2016sp2_proto protoreflect.FileDescriptor

var file_yandex_cloud_mdb_sqlserver_v1_config_sqlserver2016sp2_proto_rawDesc = []byte{
	0x0a, 0x3b, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x6d,
	0x64, 0x62, 0x2f, 0x73, 0x71, 0x6c, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x73, 0x71, 0x6c, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x32, 0x30, 0x31, 0x36, 0x73, 0x70, 0x32, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x24, 0x79,
	0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6d, 0x64, 0x62, 0x2e,
	0x73, 0x71, 0x6c, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xe9, 0x03, 0x0a, 0x19, 0x53, 0x51, 0x4c, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x32, 0x30, 0x31, 0x36, 0x73, 0x70, 0x32, 0x73, 0x74, 0x64,
	0x12, 0x60, 0x0a, 0x19, 0x6d, 0x61, 0x78, 0x5f, 0x64, 0x65, 0x67, 0x72, 0x65, 0x65, 0x5f, 0x6f,
	0x66, 0x5f, 0x70, 0x61, 0x72, 0x61, 0x6c, 0x6c, 0x65, 0x6c, 0x69, 0x73, 0x6d, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x42, 0x08, 0xfa, 0xc7, 0x31, 0x04, 0x31, 0x2d, 0x39, 0x39, 0x52, 0x16, 0x6d, 0x61, 0x78, 0x44,
	0x65, 0x67, 0x72, 0x65, 0x65, 0x4f, 0x66, 0x50, 0x61, 0x72, 0x61, 0x6c, 0x6c, 0x65, 0x6c, 0x69,
	0x73, 0x6d, 0x12, 0x6d, 0x0a, 0x1e, 0x63, 0x6f, 0x73, 0x74, 0x5f, 0x74, 0x68, 0x72, 0x65, 0x73,
	0x68, 0x6f, 0x6c, 0x64, 0x5f, 0x66, 0x6f, 0x72, 0x5f, 0x70, 0x61, 0x72, 0x61, 0x6c, 0x6c, 0x65,
	0x6c, 0x69, 0x73, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74,
	0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x0b, 0xfa, 0xc7, 0x31, 0x07, 0x35, 0x2d, 0x33,
	0x32, 0x37, 0x36, 0x37, 0x52, 0x1b, 0x63, 0x6f, 0x73, 0x74, 0x54, 0x68, 0x72, 0x65, 0x73, 0x68,
	0x6f, 0x6c, 0x64, 0x46, 0x6f, 0x72, 0x50, 0x61, 0x72, 0x61, 0x6c, 0x6c, 0x65, 0x6c, 0x69, 0x73,
	0x6d, 0x12, 0x45, 0x0a, 0x0b, 0x61, 0x75, 0x64, 0x69, 0x74, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x42, 0x07, 0xfa, 0xc7, 0x31, 0x03, 0x30, 0x2d, 0x33, 0x52, 0x0a, 0x61, 0x75,
	0x64, 0x69, 0x74, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x56, 0x0a, 0x13, 0x66, 0x69, 0x6c, 0x6c,
	0x5f, 0x66, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x5f, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x42, 0x09, 0xfa, 0xc7, 0x31, 0x05, 0x30, 0x2d, 0x31, 0x30, 0x30, 0x52, 0x11, 0x66,
	0x69, 0x6c, 0x6c, 0x46, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74,
	0x12, 0x5c, 0x0a, 0x1d, 0x6f, 0x70, 0x74, 0x69, 0x6d, 0x69, 0x7a, 0x65, 0x5f, 0x66, 0x6f, 0x72,
	0x5f, 0x61, 0x64, 0x5f, 0x68, 0x6f, 0x63, 0x5f, 0x77, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64,
	0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x52, 0x19, 0x6f, 0x70, 0x74, 0x69, 0x6d, 0x69, 0x7a, 0x65, 0x46, 0x6f, 0x72,
	0x41, 0x64, 0x48, 0x6f, 0x63, 0x57, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x73, 0x22, 0xd4,
	0x02, 0x0a, 0x1c, 0x53, 0x51, 0x4c, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x53, 0x65, 0x74, 0x32, 0x30, 0x31, 0x36, 0x73, 0x70, 0x32, 0x73, 0x74, 0x64, 0x12,
	0x6a, 0x0a, 0x10, 0x65, 0x66, 0x66, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3f, 0x2e, 0x79, 0x61, 0x6e, 0x64,
	0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6d, 0x64, 0x62, 0x2e, 0x73, 0x71, 0x6c,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x2e, 0x53, 0x51, 0x4c, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x32, 0x30, 0x31, 0x36, 0x73, 0x70, 0x32, 0x73, 0x74, 0x64, 0x52, 0x0f, 0x65, 0x66, 0x66, 0x65,
	0x63, 0x74, 0x69, 0x76, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x60, 0x0a, 0x0b, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x3f, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x6d, 0x64, 0x62, 0x2e, 0x73, 0x71, 0x6c, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x53, 0x51, 0x4c, 0x53, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x32, 0x30, 0x31, 0x36, 0x73, 0x70, 0x32, 0x73, 0x74,
	0x64, 0x52, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x66, 0x0a,
	0x0e, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3f, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6d, 0x64, 0x62, 0x2e, 0x73, 0x71, 0x6c, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x53, 0x51, 0x4c,
	0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x32, 0x30, 0x31, 0x36,
	0x73, 0x70, 0x32, 0x73, 0x74, 0x64, 0x52, 0x0d, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0xe9, 0x03, 0x0a, 0x19, 0x53, 0x51, 0x4c, 0x53, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x32, 0x30, 0x31, 0x36, 0x73, 0x70, 0x32,
	0x65, 0x6e, 0x74, 0x12, 0x60, 0x0a, 0x19, 0x6d, 0x61, 0x78, 0x5f, 0x64, 0x65, 0x67, 0x72, 0x65,
	0x65, 0x5f, 0x6f, 0x66, 0x5f, 0x70, 0x61, 0x72, 0x61, 0x6c, 0x6c, 0x65, 0x6c, 0x69, 0x73, 0x6d,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x42, 0x08, 0xfa, 0xc7, 0x31, 0x04, 0x31, 0x2d, 0x39, 0x39, 0x52, 0x16, 0x6d,
	0x61, 0x78, 0x44, 0x65, 0x67, 0x72, 0x65, 0x65, 0x4f, 0x66, 0x50, 0x61, 0x72, 0x61, 0x6c, 0x6c,
	0x65, 0x6c, 0x69, 0x73, 0x6d, 0x12, 0x6d, 0x0a, 0x1e, 0x63, 0x6f, 0x73, 0x74, 0x5f, 0x74, 0x68,
	0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x5f, 0x66, 0x6f, 0x72, 0x5f, 0x70, 0x61, 0x72, 0x61,
	0x6c, 0x6c, 0x65, 0x6c, 0x69, 0x73, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x49, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x0b, 0xfa, 0xc7, 0x31, 0x07,
	0x35, 0x2d, 0x33, 0x32, 0x37, 0x36, 0x37, 0x52, 0x1b, 0x63, 0x6f, 0x73, 0x74, 0x54, 0x68, 0x72,
	0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x46, 0x6f, 0x72, 0x50, 0x61, 0x72, 0x61, 0x6c, 0x6c, 0x65,
	0x6c, 0x69, 0x73, 0x6d, 0x12, 0x45, 0x0a, 0x0b, 0x61, 0x75, 0x64, 0x69, 0x74, 0x5f, 0x6c, 0x65,
	0x76, 0x65, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74, 0x36,
	0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x07, 0xfa, 0xc7, 0x31, 0x03, 0x30, 0x2d, 0x33, 0x52,
	0x0a, 0x61, 0x75, 0x64, 0x69, 0x74, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x56, 0x0a, 0x13, 0x66,
	0x69, 0x6c, 0x6c, 0x5f, 0x66, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x5f, 0x70, 0x65, 0x72, 0x63, 0x65,
	0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74, 0x36, 0x34,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x09, 0xfa, 0xc7, 0x31, 0x05, 0x30, 0x2d, 0x31, 0x30, 0x30,
	0x52, 0x11, 0x66, 0x69, 0x6c, 0x6c, 0x46, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x50, 0x65, 0x72, 0x63,
	0x65, 0x6e, 0x74, 0x12, 0x5c, 0x0a, 0x1d, 0x6f, 0x70, 0x74, 0x69, 0x6d, 0x69, 0x7a, 0x65, 0x5f,
	0x66, 0x6f, 0x72, 0x5f, 0x61, 0x64, 0x5f, 0x68, 0x6f, 0x63, 0x5f, 0x77, 0x6f, 0x72, 0x6b, 0x6c,
	0x6f, 0x61, 0x64, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42, 0x6f, 0x6f,
	0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x19, 0x6f, 0x70, 0x74, 0x69, 0x6d, 0x69, 0x7a, 0x65,
	0x46, 0x6f, 0x72, 0x41, 0x64, 0x48, 0x6f, 0x63, 0x57, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64,
	0x73, 0x22, 0xd4, 0x02, 0x0a, 0x1c, 0x53, 0x51, 0x4c, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x53, 0x65, 0x74, 0x32, 0x30, 0x31, 0x36, 0x73, 0x70, 0x32, 0x65,
	0x6e, 0x74, 0x12, 0x6a, 0x0a, 0x10, 0x65, 0x66, 0x66, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x5f,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3f, 0x2e, 0x79,
	0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6d, 0x64, 0x62, 0x2e,
	0x73, 0x71, 0x6c, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x2e, 0x53, 0x51, 0x4c, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x32, 0x30, 0x31, 0x36, 0x73, 0x70, 0x32, 0x65, 0x6e, 0x74, 0x52, 0x0f, 0x65,
	0x66, 0x66, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x60,
	0x0a, 0x0b, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x3f, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x6d, 0x64, 0x62, 0x2e, 0x73, 0x71, 0x6c, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x53, 0x51, 0x4c, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x32, 0x30, 0x31, 0x36, 0x73, 0x70,
	0x32, 0x65, 0x6e, 0x74, 0x52, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x12, 0x66, 0x0a, 0x0e, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3f, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65,
	0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6d, 0x64, 0x62, 0x2e, 0x73, 0x71, 0x6c, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e,
	0x53, 0x51, 0x4c, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x32,
	0x30, 0x31, 0x36, 0x73, 0x70, 0x32, 0x65, 0x6e, 0x74, 0x52, 0x0d, 0x64, 0x65, 0x66, 0x61, 0x75,
	0x6c, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x42, 0x7e, 0x0a, 0x28, 0x79, 0x61, 0x6e, 0x64,
	0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x64, 0x62,
	0x2e, 0x73, 0x71, 0x6c, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x5a, 0x52, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x67, 0x6f,
	0x2d, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78,
	0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x6d, 0x64, 0x62, 0x2f, 0x73, 0x71, 0x6c, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x3b, 0x73,
	0x71, 0x6c, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_yandex_cloud_mdb_sqlserver_v1_config_sqlserver2016sp2_proto_rawDescOnce sync.Once
	file_yandex_cloud_mdb_sqlserver_v1_config_sqlserver2016sp2_proto_rawDescData = file_yandex_cloud_mdb_sqlserver_v1_config_sqlserver2016sp2_proto_rawDesc
)

func file_yandex_cloud_mdb_sqlserver_v1_config_sqlserver2016sp2_proto_rawDescGZIP() []byte {
	file_yandex_cloud_mdb_sqlserver_v1_config_sqlserver2016sp2_proto_rawDescOnce.Do(func() {
		file_yandex_cloud_mdb_sqlserver_v1_config_sqlserver2016sp2_proto_rawDescData = protoimpl.X.CompressGZIP(file_yandex_cloud_mdb_sqlserver_v1_config_sqlserver2016sp2_proto_rawDescData)
	})
	return file_yandex_cloud_mdb_sqlserver_v1_config_sqlserver2016sp2_proto_rawDescData
}

var file_yandex_cloud_mdb_sqlserver_v1_config_sqlserver2016sp2_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_yandex_cloud_mdb_sqlserver_v1_config_sqlserver2016sp2_proto_goTypes = []interface{}{
	(*SQLServerConfig2016Sp2Std)(nil),    // 0: yandex.cloud.mdb.sqlserver.v1.config.SQLServerConfig2016sp2std
	(*SQLServerConfigSet2016Sp2Std)(nil), // 1: yandex.cloud.mdb.sqlserver.v1.config.SQLServerConfigSet2016sp2std
	(*SQLServerConfig2016Sp2Ent)(nil),    // 2: yandex.cloud.mdb.sqlserver.v1.config.SQLServerConfig2016sp2ent
	(*SQLServerConfigSet2016Sp2Ent)(nil), // 3: yandex.cloud.mdb.sqlserver.v1.config.SQLServerConfigSet2016sp2ent
	(*wrappers.Int64Value)(nil),          // 4: google.protobuf.Int64Value
	(*wrappers.BoolValue)(nil),           // 5: google.protobuf.BoolValue
}
var file_yandex_cloud_mdb_sqlserver_v1_config_sqlserver2016sp2_proto_depIdxs = []int32{
	4,  // 0: yandex.cloud.mdb.sqlserver.v1.config.SQLServerConfig2016sp2std.max_degree_of_parallelism:type_name -> google.protobuf.Int64Value
	4,  // 1: yandex.cloud.mdb.sqlserver.v1.config.SQLServerConfig2016sp2std.cost_threshold_for_parallelism:type_name -> google.protobuf.Int64Value
	4,  // 2: yandex.cloud.mdb.sqlserver.v1.config.SQLServerConfig2016sp2std.audit_level:type_name -> google.protobuf.Int64Value
	4,  // 3: yandex.cloud.mdb.sqlserver.v1.config.SQLServerConfig2016sp2std.fill_factor_percent:type_name -> google.protobuf.Int64Value
	5,  // 4: yandex.cloud.mdb.sqlserver.v1.config.SQLServerConfig2016sp2std.optimize_for_ad_hoc_workloads:type_name -> google.protobuf.BoolValue
	0,  // 5: yandex.cloud.mdb.sqlserver.v1.config.SQLServerConfigSet2016sp2std.effective_config:type_name -> yandex.cloud.mdb.sqlserver.v1.config.SQLServerConfig2016sp2std
	0,  // 6: yandex.cloud.mdb.sqlserver.v1.config.SQLServerConfigSet2016sp2std.user_config:type_name -> yandex.cloud.mdb.sqlserver.v1.config.SQLServerConfig2016sp2std
	0,  // 7: yandex.cloud.mdb.sqlserver.v1.config.SQLServerConfigSet2016sp2std.default_config:type_name -> yandex.cloud.mdb.sqlserver.v1.config.SQLServerConfig2016sp2std
	4,  // 8: yandex.cloud.mdb.sqlserver.v1.config.SQLServerConfig2016sp2ent.max_degree_of_parallelism:type_name -> google.protobuf.Int64Value
	4,  // 9: yandex.cloud.mdb.sqlserver.v1.config.SQLServerConfig2016sp2ent.cost_threshold_for_parallelism:type_name -> google.protobuf.Int64Value
	4,  // 10: yandex.cloud.mdb.sqlserver.v1.config.SQLServerConfig2016sp2ent.audit_level:type_name -> google.protobuf.Int64Value
	4,  // 11: yandex.cloud.mdb.sqlserver.v1.config.SQLServerConfig2016sp2ent.fill_factor_percent:type_name -> google.protobuf.Int64Value
	5,  // 12: yandex.cloud.mdb.sqlserver.v1.config.SQLServerConfig2016sp2ent.optimize_for_ad_hoc_workloads:type_name -> google.protobuf.BoolValue
	2,  // 13: yandex.cloud.mdb.sqlserver.v1.config.SQLServerConfigSet2016sp2ent.effective_config:type_name -> yandex.cloud.mdb.sqlserver.v1.config.SQLServerConfig2016sp2ent
	2,  // 14: yandex.cloud.mdb.sqlserver.v1.config.SQLServerConfigSet2016sp2ent.user_config:type_name -> yandex.cloud.mdb.sqlserver.v1.config.SQLServerConfig2016sp2ent
	2,  // 15: yandex.cloud.mdb.sqlserver.v1.config.SQLServerConfigSet2016sp2ent.default_config:type_name -> yandex.cloud.mdb.sqlserver.v1.config.SQLServerConfig2016sp2ent
	16, // [16:16] is the sub-list for method output_type
	16, // [16:16] is the sub-list for method input_type
	16, // [16:16] is the sub-list for extension type_name
	16, // [16:16] is the sub-list for extension extendee
	0,  // [0:16] is the sub-list for field type_name
}

func init() { file_yandex_cloud_mdb_sqlserver_v1_config_sqlserver2016sp2_proto_init() }
func file_yandex_cloud_mdb_sqlserver_v1_config_sqlserver2016sp2_proto_init() {
	if File_yandex_cloud_mdb_sqlserver_v1_config_sqlserver2016sp2_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_yandex_cloud_mdb_sqlserver_v1_config_sqlserver2016sp2_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SQLServerConfig2016Sp2Std); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_yandex_cloud_mdb_sqlserver_v1_config_sqlserver2016sp2_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SQLServerConfigSet2016Sp2Std); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_yandex_cloud_mdb_sqlserver_v1_config_sqlserver2016sp2_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SQLServerConfig2016Sp2Ent); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_yandex_cloud_mdb_sqlserver_v1_config_sqlserver2016sp2_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SQLServerConfigSet2016Sp2Ent); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_yandex_cloud_mdb_sqlserver_v1_config_sqlserver2016sp2_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_yandex_cloud_mdb_sqlserver_v1_config_sqlserver2016sp2_proto_goTypes,
		DependencyIndexes: file_yandex_cloud_mdb_sqlserver_v1_config_sqlserver2016sp2_proto_depIdxs,
		MessageInfos:      file_yandex_cloud_mdb_sqlserver_v1_config_sqlserver2016sp2_proto_msgTypes,
	}.Build()
	File_yandex_cloud_mdb_sqlserver_v1_config_sqlserver2016sp2_proto = out.File
	file_yandex_cloud_mdb_sqlserver_v1_config_sqlserver2016sp2_proto_rawDesc = nil
	file_yandex_cloud_mdb_sqlserver_v1_config_sqlserver2016sp2_proto_goTypes = nil
	file_yandex_cloud_mdb_sqlserver_v1_config_sqlserver2016sp2_proto_depIdxs = nil
}