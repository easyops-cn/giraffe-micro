package mock_restv2

import (
	"github.com/gogo/protobuf/proto"
	"github.com/gogo/protobuf/types"
)

type GetDetailRequest struct {
	ObjectId             string   `protobuf:"bytes,1,opt,name=objectId,proto3" json:"objectId" form:"objectId"`
	InstanceId           string   `protobuf:"bytes,2,opt,name=instanceId,proto3" json:"instanceId" form:"instanceId"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetDetailRequest) Reset()         { *m = GetDetailRequest{} }
func (m *GetDetailRequest) String() string { return proto.CompactTextString(m) }
func (*GetDetailRequest) ProtoMessage()    {}

type DeleteInstanceRequest struct {
	ObjectId             string   `protobuf:"bytes,1,opt,name=objectId,proto3" json:"objectId" form:"objectId"`
	InstanceId           string   `protobuf:"bytes,2,opt,name=instanceId,proto3" json:"instanceId" form:"instanceId"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteInstanceRequest) Reset()         { *m = DeleteInstanceRequest{} }
func (m *DeleteInstanceRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteInstanceRequest) ProtoMessage()    {}

type CreateInstanceRequest struct {
	ObjectId             string        `protobuf:"bytes,1,opt,name=objectId,proto3" json:"objectId" form:"objectId"`
	Instance             *types.Struct `protobuf:"bytes,2,opt,name=instance,proto3" json:"instance" form:"instance"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *CreateInstanceRequest) Reset()         { *m = CreateInstanceRequest{} }
func (m *CreateInstanceRequest) String() string { return proto.CompactTextString(m) }
func (*CreateInstanceRequest) ProtoMessage()    {}

type UpdateInstanceRequest struct {
	ObjectId             string        `protobuf:"bytes,1,opt,name=objectId,proto3" json:"objectId" form:"objectId"`
	InstanceId           string        `protobuf:"bytes,2,opt,name=instanceId,proto3" json:"instanceId" form:"instanceId"`
	Instance             *types.Struct `protobuf:"bytes,3,opt,name=instance,proto3" json:"instance" form:"instance"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *UpdateInstanceRequest) Reset()         { *m = UpdateInstanceRequest{} }
func (m *UpdateInstanceRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateInstanceRequest) ProtoMessage()    {}

type GetDetailRequestWrapper struct {
	Data                 []byte
	ObjectId             string           `protobuf:"bytes,1,opt,name=objectId,proto3" json:"objectId" form:"objectId"`
	InstanceId           string           `protobuf:"bytes,2,opt,name=instanceId,proto3" json:"instanceId" form:"instanceId"`
	Wrapper              GetDetailRequest `protobuf:"bytes,2,opt,name=wrapper,proto3" json:"wrapper"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *GetDetailRequestWrapper) Reset()         { *m = GetDetailRequestWrapper{} }
func (m *GetDetailRequestWrapper) String() string { return proto.CompactTextString(m) }
func (m *GetDetailRequestWrapper) ProtoMessage()  {}
