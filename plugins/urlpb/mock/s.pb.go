package mock

import "github.com/gogo/protobuf/proto"

type SS struct {
	Name                 string   `json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SS) Reset()         { *m = SS{} }
func (m *SS) String() string { return proto.CompactTextString(m) }
func (*SS) ProtoMessage()    {}
