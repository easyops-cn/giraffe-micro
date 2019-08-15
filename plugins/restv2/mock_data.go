package restv2

import (
	"github.com/gogo/protobuf/proto"
	"github.com/gogo/protobuf/types"
)

type getDetailRequest struct {
	ObjectId             string   `protobuf:"bytes,1,opt,name=objectId,proto3" json:"objectId" form:"objectId"`
	InstanceId           string   `protobuf:"bytes,2,opt,name=instanceId,proto3" json:"instanceId" form:"instanceId"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *getDetailRequest) Reset()         { *m = getDetailRequest{} }
func (m *getDetailRequest) String() string { return proto.CompactTextString(m) }
func (*getDetailRequest) ProtoMessage()    {}

type deleteInstanceRequest struct {
	ObjectId             string   `protobuf:"bytes,1,opt,name=objectId,proto3" json:"objectId" form:"objectId"`
	InstanceId           string   `protobuf:"bytes,2,opt,name=instanceId,proto3" json:"instanceId" form:"instanceId"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *deleteInstanceRequest) Reset()         { *m = deleteInstanceRequest{} }
func (m *deleteInstanceRequest) String() string { return proto.CompactTextString(m) }
func (*deleteInstanceRequest) ProtoMessage()    {}

type createInstanceRequest struct {
	ObjectId             string        `protobuf:"bytes,1,opt,name=objectId,proto3" json:"objectId" form:"objectId"`
	Instance             *types.Struct `protobuf:"bytes,2,opt,name=instance,proto3" json:"instance" form:"instance"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *createInstanceRequest) Reset()         { *m = createInstanceRequest{} }
func (m *createInstanceRequest) String() string { return proto.CompactTextString(m) }
func (*createInstanceRequest) ProtoMessage()    {}

type updateInstanceRequest struct {
	ObjectId             string        `protobuf:"bytes,1,opt,name=objectId,proto3" json:"objectId" form:"objectId"`
	InstanceId           string        `protobuf:"bytes,2,opt,name=instanceId,proto3" json:"instanceId" form:"instanceId"`
	Instance             *types.Struct `protobuf:"bytes,3,opt,name=instance,proto3" json:"instance" form:"instance"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *updateInstanceRequest) Reset()         { *m = updateInstanceRequest{} }
func (m *updateInstanceRequest) String() string { return proto.CompactTextString(m) }
func (*updateInstanceRequest) ProtoMessage()    {}

type getDetailRequestWrapper struct {
	Data                 []byte
	ObjectId             string           `protobuf:"bytes,1,opt,name=objectId,proto3" json:"objectId" form:"objectId"`
	InstanceId           string           `protobuf:"bytes,2,opt,name=instanceId,proto3" json:"instanceId" form:"instanceId"`
	Wrapper              getDetailRequest `protobuf:"bytes,2,opt,name=wrapper,proto3" json:"wrapper"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *getDetailRequestWrapper) Reset()         { *m = getDetailRequestWrapper{} }
func (m *getDetailRequestWrapper) String() string { return proto.CompactTextString(m) }
func (m *getDetailRequestWrapper) ProtoMessage()  {}
