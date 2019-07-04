// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: status.proto

package status

import (
	fmt "fmt"
	io "io"
	math "math"

	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
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

type Status struct {
	// 错误码, 由 giraffe.status.Code 统一定义
	Code Code `protobuf:"varint,1,opt,name=code,proto3,enum=status.Code" json:"code,omitempty"`
	// 错误描述, 由错误码翻译+错误解释组成, 面向研发者, 要求全英文编写
	Error string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	// 错误解释, 面向用户, 用业务语言描述
	Message string `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	// 错误的详细信息, map格式
	Data *types.Struct `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`
	// 错误的详细信息, any格式
	Details              []*types.Any `protobuf:"bytes,5,rep,name=details,proto3" json:"details,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Status) Reset()         { *m = Status{} }
func (m *Status) String() string { return proto.CompactTextString(m) }
func (*Status) ProtoMessage()    {}
func (*Status) Descriptor() ([]byte, []int) {
	return fileDescriptor_dfe4fce6682daf5b, []int{0}
}
func (m *Status) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Status) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Status.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Status) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Status.Merge(m, src)
}
func (m *Status) XXX_Size() int {
	return m.Size()
}
func (m *Status) XXX_DiscardUnknown() {
	xxx_messageInfo_Status.DiscardUnknown(m)
}

var xxx_messageInfo_Status proto.InternalMessageInfo

func (m *Status) GetCode() Code {
	if m != nil {
		return m.Code
	}
	return Code_OK
}

func (m *Status) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *Status) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *Status) GetData() *types.Struct {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *Status) GetDetails() []*types.Any {
	if m != nil {
		return m.Details
	}
	return nil
}

func init() {
	proto.RegisterType((*Status)(nil), "status.Status")
}

func init() { proto.RegisterFile("status.proto", fileDescriptor_dfe4fce6682daf5b) }

var fileDescriptor_dfe4fce6682daf5b = []byte{
	// 223 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x8e, 0x41, 0x4a, 0xc4, 0x30,
	0x14, 0x86, 0x79, 0x4e, 0xa7, 0x83, 0x6f, 0x8a, 0x8b, 0x30, 0x60, 0x1c, 0xa4, 0x04, 0x57, 0x01,
	0x21, 0x85, 0x7a, 0x02, 0xf5, 0x06, 0xe9, 0x09, 0xd2, 0x26, 0x16, 0xa1, 0x36, 0x92, 0xa4, 0x8b,
	0x9e, 0xcc, 0x2b, 0xb8, 0xf4, 0x08, 0xd2, 0x93, 0x88, 0x49, 0xbb, 0xd1, 0xe5, 0xff, 0xbe, 0x8f,
	0xff, 0xfd, 0x58, 0xf8, 0xa0, 0xc2, 0xe4, 0xc5, 0xbb, 0xb3, 0xc1, 0x92, 0x3c, 0xa5, 0x33, 0x76,
	0x56, 0x9b, 0x74, 0x3b, 0xdf, 0xf6, 0xd6, 0xf6, 0x83, 0xa9, 0x62, 0x6a, 0xa7, 0x97, 0xca, 0x07,
	0x37, 0x75, 0x61, 0xa5, 0x37, 0x7f, 0xa9, 0x1a, 0xe7, 0x84, 0xee, 0x3e, 0x00, 0xf3, 0x26, 0xf6,
	0x11, 0x86, 0xd9, 0x6f, 0x23, 0x05, 0x06, 0xfc, 0xaa, 0x2e, 0xc4, 0xfa, 0xf4, 0xd9, 0x6a, 0x23,
	0x23, 0x21, 0x27, 0xdc, 0x1b, 0xe7, 0xac, 0xa3, 0x17, 0x0c, 0xf8, 0xa5, 0x4c, 0x81, 0x50, 0x3c,
	0xbc, 0x19, 0xef, 0x55, 0x6f, 0xe8, 0x2e, 0xde, 0xb7, 0x48, 0xee, 0x31, 0xd3, 0x2a, 0x28, 0x9a,
	0x31, 0xe0, 0xc7, 0xfa, 0x5a, 0xa4, 0x19, 0x62, 0x9b, 0x21, 0x9a, 0x38, 0x52, 0x46, 0x89, 0x08,
	0x3c, 0x68, 0x13, 0xd4, 0xeb, 0xe0, 0xe9, 0x9e, 0xed, 0xf8, 0xb1, 0x3e, 0xfd, 0xf3, 0x1f, 0xc7,
	0x59, 0x6e, 0xd2, 0x53, 0xf1, 0xb9, 0x94, 0xf0, 0xb5, 0x94, 0xf0, 0xbd, 0x94, 0xd0, 0xe6, 0x51,
	0x7a, 0xf8, 0x09, 0x00, 0x00, 0xff, 0xff, 0x37, 0x90, 0x4e, 0x5f, 0x2b, 0x01, 0x00, 0x00,
}

func (m *Status) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Status) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Code != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintStatus(dAtA, i, uint64(m.Code))
	}
	if len(m.Error) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintStatus(dAtA, i, uint64(len(m.Error)))
		i += copy(dAtA[i:], m.Error)
	}
	if len(m.Message) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintStatus(dAtA, i, uint64(len(m.Message)))
		i += copy(dAtA[i:], m.Message)
	}
	if m.Data != nil {
		dAtA[i] = 0x22
		i++
		i = encodeVarintStatus(dAtA, i, uint64(m.Data.Size()))
		n1, err := m.Data.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if len(m.Details) > 0 {
		for _, msg := range m.Details {
			dAtA[i] = 0x2a
			i++
			i = encodeVarintStatus(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintStatus(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Status) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Code != 0 {
		n += 1 + sovStatus(uint64(m.Code))
	}
	l = len(m.Error)
	if l > 0 {
		n += 1 + l + sovStatus(uint64(l))
	}
	l = len(m.Message)
	if l > 0 {
		n += 1 + l + sovStatus(uint64(l))
	}
	if m.Data != nil {
		l = m.Data.Size()
		n += 1 + l + sovStatus(uint64(l))
	}
	if len(m.Details) > 0 {
		for _, e := range m.Details {
			l = e.Size()
			n += 1 + l + sovStatus(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovStatus(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozStatus(x uint64) (n int) {
	return sovStatus(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Status) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStatus
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
			return fmt.Errorf("proto: Status: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Status: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Code", wireType)
			}
			m.Code = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStatus
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Code |= Code(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Error", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStatus
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
				return ErrInvalidLengthStatus
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStatus
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Error = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Message", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStatus
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
				return ErrInvalidLengthStatus
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStatus
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Message = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStatus
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
				return ErrInvalidLengthStatus
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthStatus
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Data == nil {
				m.Data = &types.Struct{}
			}
			if err := m.Data.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Details", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStatus
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
				return ErrInvalidLengthStatus
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthStatus
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Details = append(m.Details, &types.Any{})
			if err := m.Details[len(m.Details)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipStatus(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthStatus
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthStatus
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
func skipStatus(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowStatus
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
					return 0, ErrIntOverflowStatus
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
					return 0, ErrIntOverflowStatus
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
				return 0, ErrInvalidLengthStatus
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthStatus
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowStatus
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
				next, err := skipStatus(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthStatus
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
	ErrInvalidLengthStatus = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowStatus   = fmt.Errorf("proto: integer overflow")
)