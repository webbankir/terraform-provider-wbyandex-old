// Code generated by protoc-gen-goext. DO NOT EDIT.

package kafka

import (
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
)

type Topic_TopicConfig = isTopic_TopicConfig

func (m *Topic) SetTopicConfig(v Topic_TopicConfig) {
	m.TopicConfig = v
}

func (m *Topic) SetName(v string) {
	m.Name = v
}

func (m *Topic) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *Topic) SetPartitions(v *wrappers.Int64Value) {
	m.Partitions = v
}

func (m *Topic) SetReplicationFactor(v *wrappers.Int64Value) {
	m.ReplicationFactor = v
}

func (m *Topic) SetTopicConfig_2_1(v *TopicConfig2_1) {
	m.TopicConfig = &Topic_TopicConfig_2_1{
		TopicConfig_2_1: v,
	}
}

func (m *Topic) SetTopicConfig_2_6(v *TopicConfig2_6) {
	m.TopicConfig = &Topic_TopicConfig_2_6{
		TopicConfig_2_6: v,
	}
}

type TopicSpec_TopicConfig = isTopicSpec_TopicConfig

func (m *TopicSpec) SetTopicConfig(v TopicSpec_TopicConfig) {
	m.TopicConfig = v
}

func (m *TopicSpec) SetName(v string) {
	m.Name = v
}

func (m *TopicSpec) SetPartitions(v *wrappers.Int64Value) {
	m.Partitions = v
}

func (m *TopicSpec) SetReplicationFactor(v *wrappers.Int64Value) {
	m.ReplicationFactor = v
}

func (m *TopicSpec) SetTopicConfig_2_1(v *TopicConfig2_1) {
	m.TopicConfig = &TopicSpec_TopicConfig_2_1{
		TopicConfig_2_1: v,
	}
}

func (m *TopicSpec) SetTopicConfig_2_6(v *TopicConfig2_6) {
	m.TopicConfig = &TopicSpec_TopicConfig_2_6{
		TopicConfig_2_6: v,
	}
}

func (m *TopicConfig2_1) SetCleanupPolicy(v TopicConfig2_1_CleanupPolicy) {
	m.CleanupPolicy = v
}

func (m *TopicConfig2_1) SetCompressionType(v CompressionType) {
	m.CompressionType = v
}

func (m *TopicConfig2_1) SetDeleteRetentionMs(v *wrappers.Int64Value) {
	m.DeleteRetentionMs = v
}

func (m *TopicConfig2_1) SetFileDeleteDelayMs(v *wrappers.Int64Value) {
	m.FileDeleteDelayMs = v
}

func (m *TopicConfig2_1) SetFlushMessages(v *wrappers.Int64Value) {
	m.FlushMessages = v
}

func (m *TopicConfig2_1) SetFlushMs(v *wrappers.Int64Value) {
	m.FlushMs = v
}

func (m *TopicConfig2_1) SetMinCompactionLagMs(v *wrappers.Int64Value) {
	m.MinCompactionLagMs = v
}

func (m *TopicConfig2_1) SetRetentionBytes(v *wrappers.Int64Value) {
	m.RetentionBytes = v
}

func (m *TopicConfig2_1) SetRetentionMs(v *wrappers.Int64Value) {
	m.RetentionMs = v
}

func (m *TopicConfig2_1) SetMaxMessageBytes(v *wrappers.Int64Value) {
	m.MaxMessageBytes = v
}

func (m *TopicConfig2_1) SetMinInsyncReplicas(v *wrappers.Int64Value) {
	m.MinInsyncReplicas = v
}

func (m *TopicConfig2_1) SetSegmentBytes(v *wrappers.Int64Value) {
	m.SegmentBytes = v
}

func (m *TopicConfig2_1) SetPreallocate(v *wrappers.BoolValue) {
	m.Preallocate = v
}

func (m *TopicConfig2_6) SetCleanupPolicy(v TopicConfig2_6_CleanupPolicy) {
	m.CleanupPolicy = v
}

func (m *TopicConfig2_6) SetCompressionType(v CompressionType) {
	m.CompressionType = v
}

func (m *TopicConfig2_6) SetDeleteRetentionMs(v *wrappers.Int64Value) {
	m.DeleteRetentionMs = v
}

func (m *TopicConfig2_6) SetFileDeleteDelayMs(v *wrappers.Int64Value) {
	m.FileDeleteDelayMs = v
}

func (m *TopicConfig2_6) SetFlushMessages(v *wrappers.Int64Value) {
	m.FlushMessages = v
}

func (m *TopicConfig2_6) SetFlushMs(v *wrappers.Int64Value) {
	m.FlushMs = v
}

func (m *TopicConfig2_6) SetMinCompactionLagMs(v *wrappers.Int64Value) {
	m.MinCompactionLagMs = v
}

func (m *TopicConfig2_6) SetRetentionBytes(v *wrappers.Int64Value) {
	m.RetentionBytes = v
}

func (m *TopicConfig2_6) SetRetentionMs(v *wrappers.Int64Value) {
	m.RetentionMs = v
}

func (m *TopicConfig2_6) SetMaxMessageBytes(v *wrappers.Int64Value) {
	m.MaxMessageBytes = v
}

func (m *TopicConfig2_6) SetMinInsyncReplicas(v *wrappers.Int64Value) {
	m.MinInsyncReplicas = v
}

func (m *TopicConfig2_6) SetSegmentBytes(v *wrappers.Int64Value) {
	m.SegmentBytes = v
}

func (m *TopicConfig2_6) SetPreallocate(v *wrappers.BoolValue) {
	m.Preallocate = v
}
