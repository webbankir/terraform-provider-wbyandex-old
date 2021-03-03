// Code generated by protoc-gen-goext. DO NOT EDIT.

package kafka

import (
	field_mask "google.golang.org/genproto/protobuf/field_mask"
)

func (m *GetTopicRequest) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *GetTopicRequest) SetTopicName(v string) {
	m.TopicName = v
}

func (m *ListTopicsRequest) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *ListTopicsRequest) SetPageSize(v int64) {
	m.PageSize = v
}

func (m *ListTopicsRequest) SetPageToken(v string) {
	m.PageToken = v
}

func (m *ListTopicsResponse) SetTopics(v []*Topic) {
	m.Topics = v
}

func (m *ListTopicsResponse) SetNextPageToken(v string) {
	m.NextPageToken = v
}

func (m *CreateTopicRequest) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *CreateTopicRequest) SetTopicSpec(v *TopicSpec) {
	m.TopicSpec = v
}

func (m *CreateTopicMetadata) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *CreateTopicMetadata) SetTopicName(v string) {
	m.TopicName = v
}

func (m *UpdateTopicRequest) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *UpdateTopicRequest) SetTopicName(v string) {
	m.TopicName = v
}

func (m *UpdateTopicRequest) SetUpdateMask(v *field_mask.FieldMask) {
	m.UpdateMask = v
}

func (m *UpdateTopicRequest) SetTopicSpec(v *TopicSpec) {
	m.TopicSpec = v
}

func (m *UpdateTopicMetadata) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *UpdateTopicMetadata) SetTopicName(v string) {
	m.TopicName = v
}

func (m *DeleteTopicRequest) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *DeleteTopicRequest) SetTopicName(v string) {
	m.TopicName = v
}

func (m *DeleteTopicMetadata) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *DeleteTopicMetadata) SetTopicName(v string) {
	m.TopicName = v
}
