// Copyright 2024 Flant JSC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.2
// source: check.proto

package dhctl

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CheckRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Message:
	//
	//	*CheckRequest_Start
	//	*CheckRequest_Stop
	Message isCheckRequest_Message `protobuf_oneof:"message"`
}

func (x *CheckRequest) Reset() {
	*x = CheckRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_check_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckRequest) ProtoMessage() {}

func (x *CheckRequest) ProtoReflect() protoreflect.Message {
	mi := &file_check_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckRequest.ProtoReflect.Descriptor instead.
func (*CheckRequest) Descriptor() ([]byte, []int) {
	return file_check_proto_rawDescGZIP(), []int{0}
}

func (m *CheckRequest) GetMessage() isCheckRequest_Message {
	if m != nil {
		return m.Message
	}
	return nil
}

func (x *CheckRequest) GetStart() *CheckStart {
	if x, ok := x.GetMessage().(*CheckRequest_Start); ok {
		return x.Start
	}
	return nil
}

func (x *CheckRequest) GetStop() *CheckStop {
	if x, ok := x.GetMessage().(*CheckRequest_Stop); ok {
		return x.Stop
	}
	return nil
}

type isCheckRequest_Message interface {
	isCheckRequest_Message()
}

type CheckRequest_Start struct {
	Start *CheckStart `protobuf:"bytes,1,opt,name=start,proto3,oneof"`
}

type CheckRequest_Stop struct {
	Stop *CheckStop `protobuf:"bytes,2,opt,name=stop,proto3,oneof"`
}

func (*CheckRequest_Start) isCheckRequest_Message() {}

func (*CheckRequest_Stop) isCheckRequest_Message() {}

type CheckResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Message:
	//
	//	*CheckResponse_Result
	//	*CheckResponse_Logs
	Message isCheckResponse_Message `protobuf_oneof:"message"`
}

func (x *CheckResponse) Reset() {
	*x = CheckResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_check_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckResponse) ProtoMessage() {}

func (x *CheckResponse) ProtoReflect() protoreflect.Message {
	mi := &file_check_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckResponse.ProtoReflect.Descriptor instead.
func (*CheckResponse) Descriptor() ([]byte, []int) {
	return file_check_proto_rawDescGZIP(), []int{1}
}

func (m *CheckResponse) GetMessage() isCheckResponse_Message {
	if m != nil {
		return m.Message
	}
	return nil
}

func (x *CheckResponse) GetResult() *CheckResult {
	if x, ok := x.GetMessage().(*CheckResponse_Result); ok {
		return x.Result
	}
	return nil
}

func (x *CheckResponse) GetLogs() *Logs {
	if x, ok := x.GetMessage().(*CheckResponse_Logs); ok {
		return x.Logs
	}
	return nil
}

type isCheckResponse_Message interface {
	isCheckResponse_Message()
}

type CheckResponse_Result struct {
	Result *CheckResult `protobuf:"bytes,1,opt,name=result,proto3,oneof"`
}

type CheckResponse_Logs struct {
	Logs *Logs `protobuf:"bytes,2,opt,name=logs,proto3,oneof"`
}

func (*CheckResponse_Result) isCheckResponse_Message() {}

func (*CheckResponse_Logs) isCheckResponse_Message() {}

type CheckStart struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ConnectionConfig              string             `protobuf:"bytes,1,opt,name=connection_config,json=connectionConfig,proto3" json:"connection_config,omitempty"`
	ClusterConfig                 string             `protobuf:"bytes,2,opt,name=cluster_config,json=clusterConfig,proto3" json:"cluster_config,omitempty"`
	ProviderSpecificClusterConfig string             `protobuf:"bytes,3,opt,name=provider_specific_cluster_config,json=providerSpecificClusterConfig,proto3" json:"provider_specific_cluster_config,omitempty"`
	State                         string             `protobuf:"bytes,4,opt,name=state,proto3" json:"state,omitempty"`
	Options                       *CheckStartOptions `protobuf:"bytes,5,opt,name=options,proto3" json:"options,omitempty"`
}

func (x *CheckStart) Reset() {
	*x = CheckStart{}
	if protoimpl.UnsafeEnabled {
		mi := &file_check_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckStart) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckStart) ProtoMessage() {}

func (x *CheckStart) ProtoReflect() protoreflect.Message {
	mi := &file_check_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckStart.ProtoReflect.Descriptor instead.
func (*CheckStart) Descriptor() ([]byte, []int) {
	return file_check_proto_rawDescGZIP(), []int{2}
}

func (x *CheckStart) GetConnectionConfig() string {
	if x != nil {
		return x.ConnectionConfig
	}
	return ""
}

func (x *CheckStart) GetClusterConfig() string {
	if x != nil {
		return x.ClusterConfig
	}
	return ""
}

func (x *CheckStart) GetProviderSpecificClusterConfig() string {
	if x != nil {
		return x.ProviderSpecificClusterConfig
	}
	return ""
}

func (x *CheckStart) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

func (x *CheckStart) GetOptions() *CheckStartOptions {
	if x != nil {
		return x.Options
	}
	return nil
}

type CheckStop struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CheckStop) Reset() {
	*x = CheckStop{}
	if protoimpl.UnsafeEnabled {
		mi := &file_check_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckStop) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckStop) ProtoMessage() {}

func (x *CheckStop) ProtoReflect() protoreflect.Message {
	mi := &file_check_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckStop.ProtoReflect.Descriptor instead.
func (*CheckStop) Descriptor() ([]byte, []int) {
	return file_check_proto_rawDescGZIP(), []int{3}
}

type CheckStartOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CommanderMode    bool                 `protobuf:"varint,1,opt,name=commander_mode,json=commanderMode,proto3" json:"commander_mode,omitempty"`
	SanityCheck      bool                 `protobuf:"varint,2,opt,name=sanity_check,json=sanityCheck,proto3" json:"sanity_check,omitempty"`
	LogWidth         int32                `protobuf:"varint,3,opt,name=log_width,json=logWidth,proto3" json:"log_width,omitempty"`
	ResourcesTimeout *durationpb.Duration `protobuf:"bytes,4,opt,name=resources_timeout,json=resourcesTimeout,proto3" json:"resources_timeout,omitempty"`
	DeckhouseTimeout *durationpb.Duration `protobuf:"bytes,5,opt,name=deckhouse_timeout,json=deckhouseTimeout,proto3" json:"deckhouse_timeout,omitempty"`
}

func (x *CheckStartOptions) Reset() {
	*x = CheckStartOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_check_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckStartOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckStartOptions) ProtoMessage() {}

func (x *CheckStartOptions) ProtoReflect() protoreflect.Message {
	mi := &file_check_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckStartOptions.ProtoReflect.Descriptor instead.
func (*CheckStartOptions) Descriptor() ([]byte, []int) {
	return file_check_proto_rawDescGZIP(), []int{4}
}

func (x *CheckStartOptions) GetCommanderMode() bool {
	if x != nil {
		return x.CommanderMode
	}
	return false
}

func (x *CheckStartOptions) GetSanityCheck() bool {
	if x != nil {
		return x.SanityCheck
	}
	return false
}

func (x *CheckStartOptions) GetLogWidth() int32 {
	if x != nil {
		return x.LogWidth
	}
	return 0
}

func (x *CheckStartOptions) GetResourcesTimeout() *durationpb.Duration {
	if x != nil {
		return x.ResourcesTimeout
	}
	return nil
}

func (x *CheckStartOptions) GetDeckhouseTimeout() *durationpb.Duration {
	if x != nil {
		return x.DeckhouseTimeout
	}
	return nil
}

type CheckResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result string `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *CheckResult) Reset() {
	*x = CheckResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_check_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckResult) ProtoMessage() {}

func (x *CheckResult) ProtoReflect() protoreflect.Message {
	mi := &file_check_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckResult.ProtoReflect.Descriptor instead.
func (*CheckResult) Descriptor() ([]byte, []int) {
	return file_check_proto_rawDescGZIP(), []int{5}
}

func (x *CheckResult) GetResult() string {
	if x != nil {
		return x.Result
	}
	return ""
}

var File_check_proto protoreflect.FileDescriptor

var file_check_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x64,
	0x68, 0x63, 0x74, 0x6c, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x6c, 0x0a, 0x0c, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x29, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x11, 0x2e, 0x64, 0x68, 0x63, 0x74, 0x6c, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x53,
	0x74, 0x61, 0x72, 0x74, 0x48, 0x00, 0x52, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x12, 0x26, 0x0a,
	0x04, 0x73, 0x74, 0x6f, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x64, 0x68,
	0x63, 0x74, 0x6c, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x53, 0x74, 0x6f, 0x70, 0x48, 0x00, 0x52,
	0x04, 0x73, 0x74, 0x6f, 0x70, 0x42, 0x09, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x22, 0x6b, 0x0a, 0x0d, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x2c, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x12, 0x2e, 0x64, 0x68, 0x63, 0x74, 0x6c, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x48, 0x00, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12,
	0x21, 0x0a, 0x04, 0x6c, 0x6f, 0x67, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e,
	0x64, 0x68, 0x63, 0x74, 0x6c, 0x2e, 0x4c, 0x6f, 0x67, 0x73, 0x48, 0x00, 0x52, 0x04, 0x6c, 0x6f,
	0x67, 0x73, 0x42, 0x09, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0xf3, 0x01,
	0x0a, 0x0a, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x53, 0x74, 0x61, 0x72, 0x74, 0x12, 0x2b, 0x0a, 0x11,
	0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x6c, 0x75,
	0x73, 0x74, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0d, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x12, 0x47, 0x0a, 0x20, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x5f, 0x73, 0x70, 0x65,
	0x63, 0x69, 0x66, 0x69, 0x63, 0x5f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x1d, 0x70, 0x72, 0x6f, 0x76,
	0x69, 0x64, 0x65, 0x72, 0x53, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x43, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61,
	0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12,
	0x32, 0x0a, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x18, 0x2e, 0x64, 0x68, 0x63, 0x74, 0x6c, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x53, 0x74,
	0x61, 0x72, 0x74, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x07, 0x6f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x22, 0x0b, 0x0a, 0x09, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x53, 0x74, 0x6f, 0x70,
	0x22, 0x8a, 0x02, 0x0a, 0x11, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x53, 0x74, 0x61, 0x72, 0x74, 0x4f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e,
	0x64, 0x65, 0x72, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d,
	0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x65, 0x72, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x21, 0x0a,
	0x0c, 0x73, 0x61, 0x6e, 0x69, 0x74, 0x79, 0x5f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x0b, 0x73, 0x61, 0x6e, 0x69, 0x74, 0x79, 0x43, 0x68, 0x65, 0x63, 0x6b,
	0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x6f, 0x67, 0x5f, 0x77, 0x69, 0x64, 0x74, 0x68, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x08, 0x6c, 0x6f, 0x67, 0x57, 0x69, 0x64, 0x74, 0x68, 0x12, 0x46, 0x0a,
	0x11, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x6f,
	0x75, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x10, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x54, 0x69,
	0x6d, 0x65, 0x6f, 0x75, 0x74, 0x12, 0x46, 0x0a, 0x11, 0x64, 0x65, 0x63, 0x6b, 0x68, 0x6f, 0x75,
	0x73, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x10, 0x64, 0x65, 0x63,
	0x6b, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x22, 0x25, 0x0a,
	0x0b, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x16, 0x0a, 0x06,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x42, 0x0a, 0x5a, 0x08, 0x70, 0x62, 0x2f, 0x64, 0x68, 0x63, 0x74, 0x6c,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_check_proto_rawDescOnce sync.Once
	file_check_proto_rawDescData = file_check_proto_rawDesc
)

func file_check_proto_rawDescGZIP() []byte {
	file_check_proto_rawDescOnce.Do(func() {
		file_check_proto_rawDescData = protoimpl.X.CompressGZIP(file_check_proto_rawDescData)
	})
	return file_check_proto_rawDescData
}

var file_check_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_check_proto_goTypes = []interface{}{
	(*CheckRequest)(nil),        // 0: dhctl.CheckRequest
	(*CheckResponse)(nil),       // 1: dhctl.CheckResponse
	(*CheckStart)(nil),          // 2: dhctl.CheckStart
	(*CheckStop)(nil),           // 3: dhctl.CheckStop
	(*CheckStartOptions)(nil),   // 4: dhctl.CheckStartOptions
	(*CheckResult)(nil),         // 5: dhctl.CheckResult
	(*Logs)(nil),                // 6: dhctl.Logs
	(*durationpb.Duration)(nil), // 7: google.protobuf.Duration
}
var file_check_proto_depIdxs = []int32{
	2, // 0: dhctl.CheckRequest.start:type_name -> dhctl.CheckStart
	3, // 1: dhctl.CheckRequest.stop:type_name -> dhctl.CheckStop
	5, // 2: dhctl.CheckResponse.result:type_name -> dhctl.CheckResult
	6, // 3: dhctl.CheckResponse.logs:type_name -> dhctl.Logs
	4, // 4: dhctl.CheckStart.options:type_name -> dhctl.CheckStartOptions
	7, // 5: dhctl.CheckStartOptions.resources_timeout:type_name -> google.protobuf.Duration
	7, // 6: dhctl.CheckStartOptions.deckhouse_timeout:type_name -> google.protobuf.Duration
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_check_proto_init() }
func file_check_proto_init() {
	if File_check_proto != nil {
		return
	}
	file_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_check_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckRequest); i {
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
		file_check_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckResponse); i {
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
		file_check_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckStart); i {
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
		file_check_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckStop); i {
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
		file_check_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckStartOptions); i {
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
		file_check_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckResult); i {
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
	file_check_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*CheckRequest_Start)(nil),
		(*CheckRequest_Stop)(nil),
	}
	file_check_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*CheckResponse_Result)(nil),
		(*CheckResponse_Logs)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_check_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_check_proto_goTypes,
		DependencyIndexes: file_check_proto_depIdxs,
		MessageInfos:      file_check_proto_msgTypes,
	}.Build()
	File_check_proto = out.File
	file_check_proto_rawDesc = nil
	file_check_proto_goTypes = nil
	file_check_proto_depIdxs = nil
}