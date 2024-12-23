// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.3
// source: datastreams/datastreams.proto

package datastreams

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type MessageType int32

const (
	MessageType_BROADCAST   MessageType = 0
	MessageType_MULTICAST   MessageType = 1
	MessageType_UNICAST     MessageType = 2
	MessageType_SUBSCRIBE   MessageType = 3
	MessageType_UNSUBSCRIBE MessageType = 4
)

// Enum value maps for MessageType.
var (
	MessageType_name = map[int32]string{
		0: "BROADCAST",
		1: "MULTICAST",
		2: "UNICAST",
		3: "SUBSCRIBE",
		4: "UNSUBSCRIBE",
	}
	MessageType_value = map[string]int32{
		"BROADCAST":   0,
		"MULTICAST":   1,
		"UNICAST":     2,
		"SUBSCRIBE":   3,
		"UNSUBSCRIBE": 4,
	}
)

func (x MessageType) Enum() *MessageType {
	p := new(MessageType)
	*p = x
	return p
}

func (x MessageType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MessageType) Descriptor() protoreflect.EnumDescriptor {
	return file_datastreams_datastreams_proto_enumTypes[0].Descriptor()
}

func (MessageType) Type() protoreflect.EnumType {
	return &file_datastreams_datastreams_proto_enumTypes[0]
}

func (x MessageType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MessageType.Descriptor instead.
func (MessageType) EnumDescriptor() ([]byte, []int) {
	return file_datastreams_datastreams_proto_rawDescGZIP(), []int{0}
}

type EventState int32

const (
	EventState_SUBSCRIBED   EventState = 0
	EventState_UNSUBSCRIBED EventState = 1
)

// Enum value maps for EventState.
var (
	EventState_name = map[int32]string{
		0: "SUBSCRIBED",
		1: "UNSUBSCRIBED",
	}
	EventState_value = map[string]int32{
		"SUBSCRIBED":   0,
		"UNSUBSCRIBED": 1,
	}
)

func (x EventState) Enum() *EventState {
	p := new(EventState)
	*p = x
	return p
}

func (x EventState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EventState) Descriptor() protoreflect.EnumDescriptor {
	return file_datastreams_datastreams_proto_enumTypes[1].Descriptor()
}

func (EventState) Type() protoreflect.EnumType {
	return &file_datastreams_datastreams_proto_enumTypes[1]
}

func (x EventState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EventState.Descriptor instead.
func (EventState) EnumDescriptor() ([]byte, []int) {
	return file_datastreams_datastreams_proto_rawDescGZIP(), []int{1}
}

type ApplicationState int32

const (
	ApplicationState_AVAILABLE   ApplicationState = 0
	ApplicationState_UNAVAILABLE ApplicationState = 1
	ApplicationState_DELETED     ApplicationState = 2
)

// Enum value maps for ApplicationState.
var (
	ApplicationState_name = map[int32]string{
		0: "AVAILABLE",
		1: "UNAVAILABLE",
		2: "DELETED",
	}
	ApplicationState_value = map[string]int32{
		"AVAILABLE":   0,
		"UNAVAILABLE": 1,
		"DELETED":     2,
	}
)

func (x ApplicationState) Enum() *ApplicationState {
	p := new(ApplicationState)
	*p = x
	return p
}

func (x ApplicationState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ApplicationState) Descriptor() protoreflect.EnumDescriptor {
	return file_datastreams_datastreams_proto_enumTypes[2].Descriptor()
}

func (ApplicationState) Type() protoreflect.EnumType {
	return &file_datastreams_datastreams_proto_enumTypes[2]
}

func (x ApplicationState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ApplicationState.Descriptor instead.
func (ApplicationState) EnumDescriptor() ([]byte, []int) {
	return file_datastreams_datastreams_proto_rawDescGZIP(), []int{2}
}

type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sender *UUID       `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty"`
	Stream string      `protobuf:"bytes,2,opt,name=stream,proto3" json:"stream,omitempty"`
	Type   MessageType `protobuf:"varint,3,opt,name=type,proto3,enum=datastreams.MessageType" json:"type,omitempty"`
	Data   *anypb.Any  `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`
	Event  *Event      `protobuf:"bytes,5,opt,name=event,proto3" json:"event,omitempty"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_datastreams_datastreams_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_datastreams_datastreams_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_datastreams_datastreams_proto_rawDescGZIP(), []int{0}
}

func (x *Message) GetSender() *UUID {
	if x != nil {
		return x.Sender
	}
	return nil
}

func (x *Message) GetStream() string {
	if x != nil {
		return x.Stream
	}
	return ""
}

func (x *Message) GetType() MessageType {
	if x != nil {
		return x.Type
	}
	return MessageType_BROADCAST
}

func (x *Message) GetData() *anypb.Any {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *Message) GetEvent() *Event {
	if x != nil {
		return x.Event
	}
	return nil
}

type DataStream struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name   string            `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Events map[string]*Event `protobuf:"bytes,2,rep,name=events,proto3" json:"events,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *DataStream) Reset() {
	*x = DataStream{}
	if protoimpl.UnsafeEnabled {
		mi := &file_datastreams_datastreams_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataStream) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataStream) ProtoMessage() {}

func (x *DataStream) ProtoReflect() protoreflect.Message {
	mi := &file_datastreams_datastreams_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataStream.ProtoReflect.Descriptor instead.
func (*DataStream) Descriptor() ([]byte, []int) {
	return file_datastreams_datastreams_proto_rawDescGZIP(), []int{1}
}

func (x *DataStream) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DataStream) GetEvents() map[string]*Event {
	if x != nil {
		return x.Events
	}
	return nil
}

type UUID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *UUID) Reset() {
	*x = UUID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_datastreams_datastreams_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UUID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UUID) ProtoMessage() {}

func (x *UUID) ProtoReflect() protoreflect.Message {
	mi := &file_datastreams_datastreams_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UUID.ProtoReflect.Descriptor instead.
func (*UUID) Descriptor() ([]byte, []int) {
	return file_datastreams_datastreams_proto_rawDescGZIP(), []int{2}
}

func (x *UUID) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type Application struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       *UUID                  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name     string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Hostname string                 `protobuf:"bytes,3,opt,name=hostname,proto3" json:"hostname,omitempty"`
	State    ApplicationState       `protobuf:"varint,4,opt,name=state,proto3,enum=datastreams.ApplicationState" json:"state,omitempty"`
	Streams  map[string]*DataStream `protobuf:"bytes,5,rep,name=streams,proto3" json:"streams,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Application) Reset() {
	*x = Application{}
	if protoimpl.UnsafeEnabled {
		mi := &file_datastreams_datastreams_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Application) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Application) ProtoMessage() {}

func (x *Application) ProtoReflect() protoreflect.Message {
	mi := &file_datastreams_datastreams_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Application.ProtoReflect.Descriptor instead.
func (*Application) Descriptor() ([]byte, []int) {
	return file_datastreams_datastreams_proto_rawDescGZIP(), []int{3}
}

func (x *Application) GetId() *UUID {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *Application) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Application) GetHostname() string {
	if x != nil {
		return x.Hostname
	}
	return ""
}

func (x *Application) GetState() ApplicationState {
	if x != nil {
		return x.State
	}
	return ApplicationState_AVAILABLE
}

func (x *Application) GetStreams() map[string]*DataStream {
	if x != nil {
		return x.Streams
	}
	return nil
}

type Event struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string     `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	State EventState `protobuf:"varint,2,opt,name=state,proto3,enum=datastreams.EventState" json:"state,omitempty"`
}

func (x *Event) Reset() {
	*x = Event{}
	if protoimpl.UnsafeEnabled {
		mi := &file_datastreams_datastreams_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Event) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Event) ProtoMessage() {}

func (x *Event) ProtoReflect() protoreflect.Message {
	mi := &file_datastreams_datastreams_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Event.ProtoReflect.Descriptor instead.
func (*Event) Descriptor() ([]byte, []int) {
	return file_datastreams_datastreams_proto_rawDescGZIP(), []int{4}
}

func (x *Event) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Event) GetState() EventState {
	if x != nil {
		return x.State
	}
	return EventState_SUBSCRIBED
}

var File_datastreams_datastreams_proto protoreflect.FileDescriptor

var file_datastreams_datastreams_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x64, 0x61, 0x74, 0x61, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x2f, 0x64, 0x61,
	0x74, 0x61, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0b, 0x64, 0x61, 0x74, 0x61, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x1a, 0x1b, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d,
	0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xce, 0x01, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x29, 0x0a, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x11, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x2e, 0x55,
	0x55, 0x49, 0x44, 0x52, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x12, 0x2c, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x18, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x2e,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x12, 0x28, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x28, 0x0a, 0x05, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x64, 0x61, 0x74,
	0x61, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x05,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x22, 0xac, 0x01, 0x0a, 0x0a, 0x44, 0x61, 0x74, 0x61, 0x53, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x3b, 0x0a, 0x06, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x73,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x53, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x73, 0x1a, 0x4d, 0x0a, 0x0b, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x28, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x73, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x73, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x3a, 0x02, 0x38, 0x01, 0x22, 0x1c, 0x0a, 0x04, 0x55, 0x55, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x22, 0xab, 0x02, 0x0a, 0x0b, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x21, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11,
	0x2e, 0x64, 0x61, 0x74, 0x61, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x2e, 0x55, 0x55, 0x49,
	0x44, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x68, 0x6f, 0x73,
	0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x68, 0x6f, 0x73,
	0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x33, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x1d, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x73, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x73, 0x2e, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x3f, 0x0a, 0x07, 0x73, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x64, 0x61,
	0x74, 0x61, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x2e, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x07, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x1a, 0x53, 0x0a, 0x0c, 0x53,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2d, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x64,
	0x61, 0x74, 0x61, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x53,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x22, 0x4a, 0x0a, 0x05, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2d, 0x0a,
	0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x17, 0x2e, 0x64,
	0x61, 0x74, 0x61, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2a, 0x58, 0x0a, 0x0b,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0d, 0x0a, 0x09, 0x42,
	0x52, 0x4f, 0x41, 0x44, 0x43, 0x41, 0x53, 0x54, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x4d, 0x55,
	0x4c, 0x54, 0x49, 0x43, 0x41, 0x53, 0x54, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x49,
	0x43, 0x41, 0x53, 0x54, 0x10, 0x02, 0x12, 0x0d, 0x0a, 0x09, 0x53, 0x55, 0x42, 0x53, 0x43, 0x52,
	0x49, 0x42, 0x45, 0x10, 0x03, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x53, 0x55, 0x42, 0x53, 0x43,
	0x52, 0x49, 0x42, 0x45, 0x10, 0x04, 0x2a, 0x2e, 0x0a, 0x0a, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x12, 0x0e, 0x0a, 0x0a, 0x53, 0x55, 0x42, 0x53, 0x43, 0x52, 0x49, 0x42,
	0x45, 0x44, 0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c, 0x55, 0x4e, 0x53, 0x55, 0x42, 0x53, 0x43, 0x52,
	0x49, 0x42, 0x45, 0x44, 0x10, 0x01, 0x2a, 0x3f, 0x0a, 0x10, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x0d, 0x0a, 0x09, 0x41, 0x56,
	0x41, 0x49, 0x4c, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x41,
	0x56, 0x41, 0x49, 0x4c, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x44, 0x45,
	0x4c, 0x45, 0x54, 0x45, 0x44, 0x10, 0x02, 0x32, 0x91, 0x01, 0x0a, 0x11, 0x44, 0x61, 0x74, 0x61,
	0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3a, 0x0a,
	0x08, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x14, 0x2e, 0x64, 0x61, 0x74, 0x61,
	0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x40, 0x0a, 0x08, 0x52, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x65, 0x72, 0x12, 0x18, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x73, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x73, 0x2e, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a,
	0x18, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x2e, 0x41, 0x70,
	0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x00, 0x42, 0x32, 0x5a, 0x30, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x61, 0x6d, 0x75, 0x65, 0x6c,
	0x68, 0x65, 0x6d, 0x2f, 0x67, 0x6f, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x73, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x73, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_datastreams_datastreams_proto_rawDescOnce sync.Once
	file_datastreams_datastreams_proto_rawDescData = file_datastreams_datastreams_proto_rawDesc
)

func file_datastreams_datastreams_proto_rawDescGZIP() []byte {
	file_datastreams_datastreams_proto_rawDescOnce.Do(func() {
		file_datastreams_datastreams_proto_rawDescData = protoimpl.X.CompressGZIP(file_datastreams_datastreams_proto_rawDescData)
	})
	return file_datastreams_datastreams_proto_rawDescData
}

var file_datastreams_datastreams_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_datastreams_datastreams_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_datastreams_datastreams_proto_goTypes = []interface{}{
	(MessageType)(0),      // 0: datastreams.MessageType
	(EventState)(0),       // 1: datastreams.EventState
	(ApplicationState)(0), // 2: datastreams.ApplicationState
	(*Message)(nil),       // 3: datastreams.Message
	(*DataStream)(nil),    // 4: datastreams.DataStream
	(*UUID)(nil),          // 5: datastreams.UUID
	(*Application)(nil),   // 6: datastreams.Application
	(*Event)(nil),         // 7: datastreams.Event
	nil,                   // 8: datastreams.DataStream.EventsEntry
	nil,                   // 9: datastreams.Application.StreamsEntry
	(*anypb.Any)(nil),     // 10: google.protobuf.Any
	(*emptypb.Empty)(nil), // 11: google.protobuf.Empty
}
var file_datastreams_datastreams_proto_depIdxs = []int32{
	5,  // 0: datastreams.Message.sender:type_name -> datastreams.UUID
	0,  // 1: datastreams.Message.type:type_name -> datastreams.MessageType
	10, // 2: datastreams.Message.data:type_name -> google.protobuf.Any
	7,  // 3: datastreams.Message.event:type_name -> datastreams.Event
	8,  // 4: datastreams.DataStream.events:type_name -> datastreams.DataStream.EventsEntry
	5,  // 5: datastreams.Application.id:type_name -> datastreams.UUID
	2,  // 6: datastreams.Application.state:type_name -> datastreams.ApplicationState
	9,  // 7: datastreams.Application.streams:type_name -> datastreams.Application.StreamsEntry
	1,  // 8: datastreams.Event.state:type_name -> datastreams.EventState
	7,  // 9: datastreams.DataStream.EventsEntry.value:type_name -> datastreams.Event
	4,  // 10: datastreams.Application.StreamsEntry.value:type_name -> datastreams.DataStream
	3,  // 11: datastreams.DataStreamService.Exchange:input_type -> datastreams.Message
	6,  // 12: datastreams.DataStreamService.Register:input_type -> datastreams.Application
	11, // 13: datastreams.DataStreamService.Exchange:output_type -> google.protobuf.Empty
	6,  // 14: datastreams.DataStreamService.Register:output_type -> datastreams.Application
	13, // [13:15] is the sub-list for method output_type
	11, // [11:13] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_datastreams_datastreams_proto_init() }
func file_datastreams_datastreams_proto_init() {
	if File_datastreams_datastreams_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_datastreams_datastreams_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message); i {
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
		file_datastreams_datastreams_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataStream); i {
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
		file_datastreams_datastreams_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UUID); i {
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
		file_datastreams_datastreams_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Application); i {
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
		file_datastreams_datastreams_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Event); i {
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
			RawDescriptor: file_datastreams_datastreams_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_datastreams_datastreams_proto_goTypes,
		DependencyIndexes: file_datastreams_datastreams_proto_depIdxs,
		EnumInfos:         file_datastreams_datastreams_proto_enumTypes,
		MessageInfos:      file_datastreams_datastreams_proto_msgTypes,
	}.Build()
	File_datastreams_datastreams_proto = out.File
	file_datastreams_datastreams_proto_rawDesc = nil
	file_datastreams_datastreams_proto_goTypes = nil
	file_datastreams_datastreams_proto_depIdxs = nil
}
