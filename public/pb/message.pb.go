// Code generated by protoc-gen-go. DO NOT EDIT.
// source: message.proto

/*
Package pb is a generated protocol buffer package.

It is generated from these files:
	message.proto

It has these top-level messages:
	Message
	MessageItem
*/
package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// 消息投递
type Message struct {
	Messages []*MessageItem `protobuf:"bytes,1,rep,name=Messages" json:"Messages,omitempty"`
}

func (m *Message) Reset()                    { *m = Message{} }
func (m *Message) String() string            { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()               {}
func (*Message) Descriptor() ([]byte, []int) { return messageFileDescriptor, []int{0} }

func (m *Message) GetMessages() []*MessageItem {
	if m != nil {
		return m.Messages
	}
	return nil
}

// 单条消息投递
type MessageItem struct {
	MessageId      int64  `protobuf:"varint,1,opt,name=message_id,json=messageId" json:"message_id,omitempty"`
	SenderType     int32  `protobuf:"varint,2,opt,name=sender_type,json=senderType" json:"sender_type,omitempty"`
	SenderId       int64  `protobuf:"varint,3,opt,name=sender_id,json=senderId" json:"sender_id,omitempty"`
	SenderDeviceId int64  `protobuf:"varint,4,opt,name=sender_device_id,json=senderDeviceId" json:"sender_device_id,omitempty"`
	ReceiverType   int32  `protobuf:"varint,5,opt,name=receiver_type,json=receiverType" json:"receiver_type,omitempty"`
	ReceiverId     int64  `protobuf:"varint,6,opt,name=receiver_id,json=receiverId" json:"receiver_id,omitempty"`
	Type           int32  `protobuf:"varint,7,opt,name=type" json:"type,omitempty"`
	Content        string `protobuf:"bytes,8,opt,name=content" json:"content,omitempty"`
	SyncSequence   int64  `protobuf:"varint,9,opt,name=sync_sequence,json=syncSequence" json:"sync_sequence,omitempty"`
	SendTime       int64  `protobuf:"varint,10,opt,name=send_time,json=sendTime" json:"send_time,omitempty"`
}

func (m *MessageItem) Reset()                    { *m = MessageItem{} }
func (m *MessageItem) String() string            { return proto.CompactTextString(m) }
func (*MessageItem) ProtoMessage()               {}
func (*MessageItem) Descriptor() ([]byte, []int) { return messageFileDescriptor, []int{1} }

func (m *MessageItem) GetMessageId() int64 {
	if m != nil {
		return m.MessageId
	}
	return 0
}

func (m *MessageItem) GetSenderType() int32 {
	if m != nil {
		return m.SenderType
	}
	return 0
}

func (m *MessageItem) GetSenderId() int64 {
	if m != nil {
		return m.SenderId
	}
	return 0
}

func (m *MessageItem) GetSenderDeviceId() int64 {
	if m != nil {
		return m.SenderDeviceId
	}
	return 0
}

func (m *MessageItem) GetReceiverType() int32 {
	if m != nil {
		return m.ReceiverType
	}
	return 0
}

func (m *MessageItem) GetReceiverId() int64 {
	if m != nil {
		return m.ReceiverId
	}
	return 0
}

func (m *MessageItem) GetType() int32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *MessageItem) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *MessageItem) GetSyncSequence() int64 {
	if m != nil {
		return m.SyncSequence
	}
	return 0
}

func (m *MessageItem) GetSendTime() int64 {
	if m != nil {
		return m.SendTime
	}
	return 0
}

func init() {
	proto.RegisterType((*Message)(nil), "pb.Message")
	proto.RegisterType((*MessageItem)(nil), "pb.MessageItem")
}

func init() { proto.RegisterFile("message.proto", messageFileDescriptor) }

var messageFileDescriptor = []byte{
	// 267 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x91, 0xcd, 0x6a, 0x83, 0x40,
	0x10, 0xc7, 0x51, 0x93, 0xa8, 0x63, 0xd2, 0x96, 0x39, 0x2d, 0x94, 0x12, 0x49, 0x2f, 0x42, 0xc1,
	0x43, 0x0b, 0x7d, 0x82, 0x5e, 0x3c, 0xf4, 0x62, 0x73, 0x97, 0xea, 0x0e, 0x65, 0x0f, 0xea, 0xd6,
	0xdd, 0x06, 0x7c, 0xc5, 0x3e, 0x55, 0xc9, 0xb8, 0x9b, 0xe6, 0x36, 0xfb, 0x9b, 0xff, 0x07, 0xc3,
	0xc2, 0xae, 0x27, 0x63, 0x3e, 0xbf, 0xa8, 0xd4, 0xd3, 0x68, 0x47, 0x0c, 0x75, 0x7b, 0x78, 0x85,
	0xf8, 0x7d, 0x81, 0xf8, 0x04, 0x89, 0x1b, 0x8d, 0x08, 0xf2, 0xa8, 0xc8, 0x9e, 0x6f, 0x4b, 0xdd,
	0x96, 0x8e, 0x55, 0x96, 0xfa, 0xfa, 0x22, 0x38, 0xfc, 0x86, 0x90, 0x5d, 0x6d, 0xf0, 0x01, 0xc0,
	0x85, 0x37, 0x4a, 0x8a, 0x20, 0x0f, 0x8a, 0xa8, 0x4e, 0x1d, 0xa9, 0x24, 0xee, 0x21, 0x33, 0x34,
	0x48, 0x9a, 0x1a, 0x3b, 0x6b, 0x12, 0x61, 0x1e, 0x14, 0xeb, 0x1a, 0x16, 0x74, 0x9c, 0x35, 0xe1,
	0x3d, 0xa4, 0x4e, 0xa0, 0xa4, 0x88, 0xd8, 0x9e, 0x2c, 0xa0, 0x92, 0x58, 0xc0, 0x9d, 0x5b, 0x4a,
	0x3a, 0xa9, 0x8e, 0x2b, 0x56, 0xac, 0xb9, 0x59, 0xf8, 0x1b, 0xe3, 0x4a, 0xe2, 0x23, 0xec, 0x26,
	0xea, 0x48, 0x9d, 0x7c, 0xd3, 0x9a, 0x9b, 0xb6, 0x1e, 0x72, 0xd7, 0x1e, 0xb2, 0x8b, 0x48, 0x49,
	0xb1, 0xe1, 0x24, 0xf0, 0xa8, 0x92, 0x88, 0xb0, 0x62, 0x73, 0xcc, 0x66, 0x9e, 0x51, 0x40, 0xdc,
	0x8d, 0x83, 0xa5, 0xc1, 0x8a, 0x24, 0x0f, 0x8a, 0xb4, 0xf6, 0xcf, 0x73, 0xa7, 0x99, 0x87, 0xae,
	0x31, 0xf4, 0xfd, 0x43, 0x43, 0x47, 0x22, 0xe5, 0xc0, 0xed, 0x19, 0x7e, 0x38, 0xe6, 0xef, 0x6b,
	0xac, 0xea, 0x49, 0xc0, 0xff, 0x7d, 0x47, 0xd5, 0x53, 0xbb, 0xe1, 0xff, 0x78, 0xf9, 0x0b, 0x00,
	0x00, 0xff, 0xff, 0xa7, 0xaf, 0x84, 0x25, 0xa0, 0x01, 0x00, 0x00,
}