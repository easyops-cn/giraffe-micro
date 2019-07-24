// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: query.proto

package mock

import (
	encoding_binary "encoding/binary"
	fmt "fmt"
	io "io"
	math "math"

	proto "github.com/gogo/protobuf/proto"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type S struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *S) Reset()         { *m = S{} }
func (m *S) String() string { return proto.CompactTextString(m) }
func (*S) ProtoMessage()    {}
func (*S) Descriptor() ([]byte, []int) {
	return fileDescriptor_5c6ac9b241082464, []int{0}
}
func (m *S) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *S) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_S.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *S) XXX_Merge(src proto.Message) {
	xxx_messageInfo_S.Merge(m, src)
}
func (m *S) XXX_Size() int {
	return m.Size()
}
func (m *S) XXX_DiscardUnknown() {
	xxx_messageInfo_S.DiscardUnknown(m)
}

var xxx_messageInfo_S proto.InternalMessageInfo

func (m *S) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type Query struct {
	StringValue          string   `protobuf:"bytes,1,opt,name=string_value,json=stringValue,proto3" json:"string_value,omitempty"`
	BoolValue            bool     `protobuf:"varint,2,opt,name=bool_value,json=boolValue,proto3" json:"bool_value,omitempty"`
	Int32Value           int32    `protobuf:"varint,3,opt,name=int32_value,json=int32Value,proto3" json:"int32_value,omitempty"`
	Int64Value           int64    `protobuf:"varint,4,opt,name=int64_value,json=int64Value,proto3" json:"int64_value,omitempty"`
	Uint32Value          uint32   `protobuf:"varint,5,opt,name=uint32_value,json=uint32Value,proto3" json:"uint32_value,omitempty"`
	Uint64Value          uint64   `protobuf:"varint,6,opt,name=uint64_value,json=uint64Value,proto3" json:"uint64_value,omitempty"`
	FloatValue           float32  `protobuf:"fixed32,7,opt,name=float_value,json=floatValue,proto3" json:"float_value,omitempty"`
	DoubleValue          float64  `protobuf:"fixed64,8,opt,name=double_value,json=doubleValue,proto3" json:"double_value,omitempty"`
	StructValue          *S       `protobuf:"bytes,9,opt,name=struct_value,json=structValue,proto3" json:"struct_value,omitempty"`
	SliceValue           []string `protobuf:"bytes,10,rep,name=slice_value,json=sliceValue,proto3" json:"slice_value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Query) Reset()         { *m = Query{} }
func (m *Query) String() string { return proto.CompactTextString(m) }
func (*Query) ProtoMessage()    {}
func (*Query) Descriptor() ([]byte, []int) {
	return fileDescriptor_5c6ac9b241082464, []int{1}
}
func (m *Query) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Query) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Query.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Query) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Query.Merge(m, src)
}
func (m *Query) XXX_Size() int {
	return m.Size()
}
func (m *Query) XXX_DiscardUnknown() {
	xxx_messageInfo_Query.DiscardUnknown(m)
}

var xxx_messageInfo_Query proto.InternalMessageInfo

func (m *Query) GetStringValue() string {
	if m != nil {
		return m.StringValue
	}
	return ""
}

func (m *Query) GetBoolValue() bool {
	if m != nil {
		return m.BoolValue
	}
	return false
}

func (m *Query) GetInt32Value() int32 {
	if m != nil {
		return m.Int32Value
	}
	return 0
}

func (m *Query) GetInt64Value() int64 {
	if m != nil {
		return m.Int64Value
	}
	return 0
}

func (m *Query) GetUint32Value() uint32 {
	if m != nil {
		return m.Uint32Value
	}
	return 0
}

func (m *Query) GetUint64Value() uint64 {
	if m != nil {
		return m.Uint64Value
	}
	return 0
}

func (m *Query) GetFloatValue() float32 {
	if m != nil {
		return m.FloatValue
	}
	return 0
}

func (m *Query) GetDoubleValue() float64 {
	if m != nil {
		return m.DoubleValue
	}
	return 0
}

func (m *Query) GetStructValue() *S {
	if m != nil {
		return m.StructValue
	}
	return nil
}

func (m *Query) GetSliceValue() []string {
	if m != nil {
		return m.SliceValue
	}
	return nil
}

func init() {
	proto.RegisterType((*S)(nil), "mock.S")
	proto.RegisterType((*Query)(nil), "mock.Query")
}

func init() { proto.RegisterFile("query.proto", fileDescriptor_5c6ac9b241082464) }

var fileDescriptor_5c6ac9b241082464 = []byte{
	// 269 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x91, 0x4f, 0x4e, 0x83, 0x40,
	0x14, 0x87, 0xf3, 0xf8, 0xd3, 0x96, 0x37, 0xb8, 0x61, 0xa3, 0x1b, 0xe9, 0xb4, 0xab, 0x17, 0x17,
	0x2c, 0xda, 0xa6, 0x07, 0xf0, 0x06, 0x4e, 0x13, 0xb7, 0x06, 0x10, 0x0d, 0x91, 0x32, 0x4a, 0x67,
	0x4c, 0xbc, 0xa1, 0x4b, 0x8f, 0x60, 0x38, 0x49, 0xc3, 0xcc, 0x40, 0xd8, 0x91, 0xef, 0x7d, 0xf9,
	0x48, 0x7e, 0x83, 0xec, 0x4b, 0x57, 0xdd, 0x4f, 0xf6, 0xd9, 0x49, 0x25, 0x93, 0xe0, 0x2c, 0xcb,
	0x8f, 0xed, 0x2d, 0xc2, 0x29, 0x49, 0x30, 0x68, 0xf3, 0x73, 0x75, 0x07, 0x1c, 0x28, 0x12, 0xe6,
	0x7b, 0xdb, 0x7b, 0x18, 0x3e, 0x0d, 0x7a, 0xb2, 0xc1, 0xf8, 0xa2, 0xba, 0xba, 0x7d, 0x7f, 0xf9,
	0xce, 0x1b, 0x3d, 0x5a, 0xcc, 0xb2, 0xe7, 0x01, 0x25, 0xf7, 0x88, 0x85, 0x94, 0x8d, 0x13, 0x3c,
	0x0e, 0xb4, 0x12, 0xd1, 0x40, 0xec, 0x79, 0x8d, 0xac, 0x6e, 0xd5, 0x7e, 0xe7, 0xee, 0x3e, 0x07,
	0x0a, 0x05, 0x1a, 0x34, 0x17, 0x8e, 0x07, 0x27, 0x04, 0x1c, 0xc8, 0x37, 0xc2, 0xf1, 0x60, 0x85,
	0x0d, 0xc6, 0x7a, 0x9e, 0x08, 0x39, 0xd0, 0x8d, 0x60, 0x7a, 0xd6, 0x70, 0xca, 0x14, 0x59, 0x70,
	0xa0, 0xc0, 0x2a, 0x63, 0x65, 0x8d, 0xec, 0xad, 0x91, 0xb9, 0x72, 0xc6, 0x92, 0x03, 0x79, 0x02,
	0x0d, 0x9a, 0x1a, 0xaf, 0x52, 0x17, 0x4d, 0xe5, 0x8c, 0x15, 0x07, 0x02, 0xc1, 0x2c, 0xb3, 0xca,
	0x83, 0x59, 0x43, 0x97, 0x63, 0x24, 0xe2, 0x40, 0x6c, 0xb7, 0xcc, 0x86, 0x35, 0xb3, 0x93, 0x99,
	0x45, 0x97, 0x6a, 0xfa, 0xdf, 0xa5, 0xa9, 0xcb, 0xb1, 0x86, 0xdc, 0xa7, 0x48, 0xa0, 0x41, 0x46,
	0x78, 0x8c, 0x7f, 0xfb, 0x14, 0xfe, 0xfa, 0x14, 0xfe, 0xfb, 0x14, 0x8a, 0x85, 0x79, 0x98, 0xfd,
	0x35, 0x00, 0x00, 0xff, 0xff, 0xa4, 0x27, 0x97, 0x92, 0xa7, 0x01, 0x00, 0x00,
}

func (m *S) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *S) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Name) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *Query) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Query) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.StringValue) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintQuery(dAtA, i, uint64(len(m.StringValue)))
		i += copy(dAtA[i:], m.StringValue)
	}
	if m.BoolValue {
		dAtA[i] = 0x10
		i++
		if m.BoolValue {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.Int32Value != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintQuery(dAtA, i, uint64(m.Int32Value))
	}
	if m.Int64Value != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintQuery(dAtA, i, uint64(m.Int64Value))
	}
	if m.Uint32Value != 0 {
		dAtA[i] = 0x28
		i++
		i = encodeVarintQuery(dAtA, i, uint64(m.Uint32Value))
	}
	if m.Uint64Value != 0 {
		dAtA[i] = 0x30
		i++
		i = encodeVarintQuery(dAtA, i, uint64(m.Uint64Value))
	}
	if m.FloatValue != 0 {
		dAtA[i] = 0x3d
		i++
		encoding_binary.LittleEndian.PutUint32(dAtA[i:], uint32(math.Float32bits(float32(m.FloatValue))))
		i += 4
	}
	if m.DoubleValue != 0 {
		dAtA[i] = 0x41
		i++
		encoding_binary.LittleEndian.PutUint64(dAtA[i:], uint64(math.Float64bits(float64(m.DoubleValue))))
		i += 8
	}
	if m.StructValue != nil {
		dAtA[i] = 0x4a
		i++
		i = encodeVarintQuery(dAtA, i, uint64(m.StructValue.Size()))
		n1, err := m.StructValue.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if len(m.SliceValue) > 0 {
		for _, s := range m.SliceValue {
			dAtA[i] = 0x52
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *S) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *Query) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.StringValue)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	if m.BoolValue {
		n += 2
	}
	if m.Int32Value != 0 {
		n += 1 + sovQuery(uint64(m.Int32Value))
	}
	if m.Int64Value != 0 {
		n += 1 + sovQuery(uint64(m.Int64Value))
	}
	if m.Uint32Value != 0 {
		n += 1 + sovQuery(uint64(m.Uint32Value))
	}
	if m.Uint64Value != 0 {
		n += 1 + sovQuery(uint64(m.Uint64Value))
	}
	if m.FloatValue != 0 {
		n += 5
	}
	if m.DoubleValue != 0 {
		n += 9
	}
	if m.StructValue != nil {
		l = m.StructValue.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	if len(m.SliceValue) > 0 {
		for _, s := range m.SliceValue {
			l = len(s)
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovQuery(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *S) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: S: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: S: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Query) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Query: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Query: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StringValue", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.StringValue = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BoolValue", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.BoolValue = bool(v != 0)
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Int32Value", wireType)
			}
			m.Int32Value = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Int32Value |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Int64Value", wireType)
			}
			m.Int64Value = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Int64Value |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Uint32Value", wireType)
			}
			m.Uint32Value = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Uint32Value |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Uint64Value", wireType)
			}
			m.Uint64Value = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Uint64Value |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 5 {
				return fmt.Errorf("proto: wrong wireType = %d for field FloatValue", wireType)
			}
			var v uint32
			if (iNdEx + 4) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint32(encoding_binary.LittleEndian.Uint32(dAtA[iNdEx:]))
			iNdEx += 4
			m.FloatValue = float32(math.Float32frombits(v))
		case 8:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field DoubleValue", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint64(encoding_binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
			m.DoubleValue = float64(math.Float64frombits(v))
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StructValue", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.StructValue == nil {
				m.StructValue = &S{}
			}
			if err := m.StructValue.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SliceValue", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SliceValue = append(m.SliceValue, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthQuery
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowQuery
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipQuery(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthQuery
				}
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthQuery = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery   = fmt.Errorf("proto: integer overflow")
)