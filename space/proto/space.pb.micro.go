// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/space.proto

package space

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	client "github.com/micro/micro/v5/service/client"
	server "github.com/micro/micro/v5/service/server"
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

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Space service

type SpaceService interface {
	Create(ctx context.Context, in *CreateRequest, opts ...client.CallOption) (*CreateResponse, error)
	Update(ctx context.Context, in *UpdateRequest, opts ...client.CallOption) (*UpdateResponse, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...client.CallOption) (*DeleteResponse, error)
	List(ctx context.Context, in *ListRequest, opts ...client.CallOption) (*ListResponse, error)
	Head(ctx context.Context, in *HeadRequest, opts ...client.CallOption) (*HeadResponse, error)
	Read(ctx context.Context, in *ReadRequest, opts ...client.CallOption) (*ReadResponse, error)
	Download(ctx context.Context, in *DownloadRequest, opts ...client.CallOption) (*DownloadResponse, error)
	Upload(ctx context.Context, in *UploadRequest, opts ...client.CallOption) (*UploadResponse, error)
}

type spaceService struct {
	c    client.Client
	name string
}

func NewSpaceService(name string, c client.Client) SpaceService {
	return &spaceService{
		c:    c,
		name: name,
	}
}

func (c *spaceService) Create(ctx context.Context, in *CreateRequest, opts ...client.CallOption) (*CreateResponse, error) {
	req := c.c.NewRequest(c.name, "Space.Create", in)
	out := new(CreateResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *spaceService) Update(ctx context.Context, in *UpdateRequest, opts ...client.CallOption) (*UpdateResponse, error) {
	req := c.c.NewRequest(c.name, "Space.Update", in)
	out := new(UpdateResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *spaceService) Delete(ctx context.Context, in *DeleteRequest, opts ...client.CallOption) (*DeleteResponse, error) {
	req := c.c.NewRequest(c.name, "Space.Delete", in)
	out := new(DeleteResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *spaceService) List(ctx context.Context, in *ListRequest, opts ...client.CallOption) (*ListResponse, error) {
	req := c.c.NewRequest(c.name, "Space.List", in)
	out := new(ListResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *spaceService) Head(ctx context.Context, in *HeadRequest, opts ...client.CallOption) (*HeadResponse, error) {
	req := c.c.NewRequest(c.name, "Space.Head", in)
	out := new(HeadResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *spaceService) Read(ctx context.Context, in *ReadRequest, opts ...client.CallOption) (*ReadResponse, error) {
	req := c.c.NewRequest(c.name, "Space.Read", in)
	out := new(ReadResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *spaceService) Download(ctx context.Context, in *DownloadRequest, opts ...client.CallOption) (*DownloadResponse, error) {
	req := c.c.NewRequest(c.name, "Space.Download", in)
	out := new(DownloadResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *spaceService) Upload(ctx context.Context, in *UploadRequest, opts ...client.CallOption) (*UploadResponse, error) {
	req := c.c.NewRequest(c.name, "Space.Upload", in)
	out := new(UploadResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Space service

type SpaceHandler interface {
	Create(context.Context, *CreateRequest, *CreateResponse) error
	Update(context.Context, *UpdateRequest, *UpdateResponse) error
	Delete(context.Context, *DeleteRequest, *DeleteResponse) error
	List(context.Context, *ListRequest, *ListResponse) error
	Head(context.Context, *HeadRequest, *HeadResponse) error
	Read(context.Context, *ReadRequest, *ReadResponse) error
	Download(context.Context, *DownloadRequest, *DownloadResponse) error
	Upload(context.Context, *UploadRequest, *UploadResponse) error
}

func RegisterSpaceHandler(s server.Server, hdlr SpaceHandler, opts ...server.HandlerOption) error {
	type space interface {
		Create(ctx context.Context, in *CreateRequest, out *CreateResponse) error
		Update(ctx context.Context, in *UpdateRequest, out *UpdateResponse) error
		Delete(ctx context.Context, in *DeleteRequest, out *DeleteResponse) error
		List(ctx context.Context, in *ListRequest, out *ListResponse) error
		Head(ctx context.Context, in *HeadRequest, out *HeadResponse) error
		Read(ctx context.Context, in *ReadRequest, out *ReadResponse) error
		Download(ctx context.Context, in *DownloadRequest, out *DownloadResponse) error
		Upload(ctx context.Context, in *UploadRequest, out *UploadResponse) error
	}
	type Space struct {
		space
	}
	h := &spaceHandler{hdlr}
	return s.Handle(s.NewHandler(&Space{h}, opts...))
}

type spaceHandler struct {
	SpaceHandler
}

func (h *spaceHandler) Create(ctx context.Context, in *CreateRequest, out *CreateResponse) error {
	return h.SpaceHandler.Create(ctx, in, out)
}

func (h *spaceHandler) Update(ctx context.Context, in *UpdateRequest, out *UpdateResponse) error {
	return h.SpaceHandler.Update(ctx, in, out)
}

func (h *spaceHandler) Delete(ctx context.Context, in *DeleteRequest, out *DeleteResponse) error {
	return h.SpaceHandler.Delete(ctx, in, out)
}

func (h *spaceHandler) List(ctx context.Context, in *ListRequest, out *ListResponse) error {
	return h.SpaceHandler.List(ctx, in, out)
}

func (h *spaceHandler) Head(ctx context.Context, in *HeadRequest, out *HeadResponse) error {
	return h.SpaceHandler.Head(ctx, in, out)
}

func (h *spaceHandler) Read(ctx context.Context, in *ReadRequest, out *ReadResponse) error {
	return h.SpaceHandler.Read(ctx, in, out)
}

func (h *spaceHandler) Download(ctx context.Context, in *DownloadRequest, out *DownloadResponse) error {
	return h.SpaceHandler.Download(ctx, in, out)
}

func (h *spaceHandler) Upload(ctx context.Context, in *UploadRequest, out *UploadResponse) error {
	return h.SpaceHandler.Upload(ctx, in, out)
}
