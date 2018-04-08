// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: github.com/micro/message-srv/proto/message/message.proto

/*
Package message is a generated protocol buffer package.

It is generated from these files:
	github.com/micro/message-srv/proto/message/message.proto

It has these top-level messages:
	Event
	CreateRequest
	CreateResponse
	UpdateRequest
	UpdateResponse
	DeleteRequest
	DeleteResponse
	ReadRequest
	ReadResponse
	SearchRequest
	SearchResponse
	StreamRequest
	StreamResponse
*/
package message

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "context"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Message service

type MessageService interface {
	Create(ctx context.Context, in *CreateRequest, opts ...client.CallOption) (*CreateResponse, error)
	Update(ctx context.Context, in *UpdateRequest, opts ...client.CallOption) (*UpdateResponse, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...client.CallOption) (*DeleteResponse, error)
	Search(ctx context.Context, in *SearchRequest, opts ...client.CallOption) (*SearchResponse, error)
	Stream(ctx context.Context, in *StreamRequest, opts ...client.CallOption) (Message_StreamService, error)
	Read(ctx context.Context, in *ReadRequest, opts ...client.CallOption) (*ReadResponse, error)
}

type messageService struct {
	c           client.Client
	serviceName string
}

func MessageServiceClient(serviceName string, c client.Client) MessageService {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "message"
	}
	return &messageService{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *messageService) Create(ctx context.Context, in *CreateRequest, opts ...client.CallOption) (*CreateResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Message.Create", in)
	out := new(CreateResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageService) Update(ctx context.Context, in *UpdateRequest, opts ...client.CallOption) (*UpdateResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Message.Update", in)
	out := new(UpdateResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageService) Delete(ctx context.Context, in *DeleteRequest, opts ...client.CallOption) (*DeleteResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Message.Delete", in)
	out := new(DeleteResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageService) Search(ctx context.Context, in *SearchRequest, opts ...client.CallOption) (*SearchResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Message.Search", in)
	out := new(SearchResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageService) Stream(ctx context.Context, in *StreamRequest, opts ...client.CallOption) (Message_StreamService, error) {
	req := c.c.NewRequest(c.serviceName, "Message.Stream", &StreamRequest{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	if err := stream.Send(in); err != nil {
		return nil, err
	}
	return &messageStreamService{stream}, nil
}

type Message_StreamService interface {
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Recv() (*StreamResponse, error)
}

type messageStreamService struct {
	stream client.Streamer
}

func (x *messageStreamService) Close() error {
	return x.stream.Close()
}

func (x *messageStreamService) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *messageStreamService) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *messageStreamService) Recv() (*StreamResponse, error) {
	m := new(StreamResponse)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (c *messageService) Read(ctx context.Context, in *ReadRequest, opts ...client.CallOption) (*ReadResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Message.Read", in)
	out := new(ReadResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Message service

type MessageHandler interface {
	Create(context.Context, *CreateRequest, *CreateResponse) error
	Update(context.Context, *UpdateRequest, *UpdateResponse) error
	Delete(context.Context, *DeleteRequest, *DeleteResponse) error
	Search(context.Context, *SearchRequest, *SearchResponse) error
	Stream(context.Context, *StreamRequest, Message_StreamStream) error
	Read(context.Context, *ReadRequest, *ReadResponse) error
}

func RegisterMessageHandler(s server.Server, hdlr MessageHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&Message{hdlr}, opts...))
}

type Message struct {
	MessageHandler
}

func (h *Message) Create(ctx context.Context, in *CreateRequest, out *CreateResponse) error {
	return h.MessageHandler.Create(ctx, in, out)
}

func (h *Message) Update(ctx context.Context, in *UpdateRequest, out *UpdateResponse) error {
	return h.MessageHandler.Update(ctx, in, out)
}

func (h *Message) Delete(ctx context.Context, in *DeleteRequest, out *DeleteResponse) error {
	return h.MessageHandler.Delete(ctx, in, out)
}

func (h *Message) Search(ctx context.Context, in *SearchRequest, out *SearchResponse) error {
	return h.MessageHandler.Search(ctx, in, out)
}

func (h *Message) Stream(ctx context.Context, stream server.Streamer) error {
	m := new(StreamRequest)
	if err := stream.Recv(m); err != nil {
		return err
	}
	return h.MessageHandler.Stream(ctx, m, &messageStreamStream{stream})
}

type Message_StreamStream interface {
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*StreamResponse) error
}

type messageStreamStream struct {
	stream server.Streamer
}

func (x *messageStreamStream) Close() error {
	return x.stream.Close()
}

func (x *messageStreamStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *messageStreamStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *messageStreamStream) Send(m *StreamResponse) error {
	return x.stream.Send(m)
}

func (h *Message) Read(ctx context.Context, in *ReadRequest, out *ReadResponse) error {
	return h.MessageHandler.Read(ctx, in, out)
}