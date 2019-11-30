// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: proto/beacon/p2p/v1/archive.proto

package ethereum_beacon_p2p_v1

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	v1alpha1 "github.com/prysmaticlabs/ethereumapis/eth/v1alpha1"
	io "io"
	math "math"
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

type ArchivedActiveSetChanges struct {
	Activated            []uint64                     `protobuf:"varint,1,rep,packed,name=activated,proto3" json:"activated,omitempty"`
	Exited               []uint64                     `protobuf:"varint,2,rep,packed,name=exited,proto3" json:"exited,omitempty"`
	Slashed              []uint64                     `protobuf:"varint,4,rep,packed,name=slashed,proto3" json:"slashed,omitempty"`
	VoluntaryExits       []*v1alpha1.VoluntaryExit    `protobuf:"bytes,6,rep,name=voluntary_exits,json=voluntaryExits,proto3" json:"voluntary_exits,omitempty"`
	ProposerSlashings    []*v1alpha1.ProposerSlashing `protobuf:"bytes,7,rep,name=proposer_slashings,json=proposerSlashings,proto3" json:"proposer_slashings,omitempty"`
	AttesterSlashings    []*v1alpha1.AttesterSlashing `protobuf:"bytes,8,rep,name=attester_slashings,json=attesterSlashings,proto3" json:"attester_slashings,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *ArchivedActiveSetChanges) Reset()         { *m = ArchivedActiveSetChanges{} }
func (m *ArchivedActiveSetChanges) String() string { return proto.CompactTextString(m) }
func (*ArchivedActiveSetChanges) ProtoMessage()    {}
func (*ArchivedActiveSetChanges) Descriptor() ([]byte, []int) {
	return fileDescriptor_289929478e9672a3, []int{0}
}
func (m *ArchivedActiveSetChanges) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ArchivedActiveSetChanges) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ArchivedActiveSetChanges.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ArchivedActiveSetChanges) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ArchivedActiveSetChanges.Merge(m, src)
}
func (m *ArchivedActiveSetChanges) XXX_Size() int {
	return m.Size()
}
func (m *ArchivedActiveSetChanges) XXX_DiscardUnknown() {
	xxx_messageInfo_ArchivedActiveSetChanges.DiscardUnknown(m)
}

var xxx_messageInfo_ArchivedActiveSetChanges proto.InternalMessageInfo

func (m *ArchivedActiveSetChanges) GetActivated() []uint64 {
	if m != nil {
		return m.Activated
	}
	return nil
}

func (m *ArchivedActiveSetChanges) GetExited() []uint64 {
	if m != nil {
		return m.Exited
	}
	return nil
}

func (m *ArchivedActiveSetChanges) GetSlashed() []uint64 {
	if m != nil {
		return m.Slashed
	}
	return nil
}

func (m *ArchivedActiveSetChanges) GetVoluntaryExits() []*v1alpha1.VoluntaryExit {
	if m != nil {
		return m.VoluntaryExits
	}
	return nil
}

func (m *ArchivedActiveSetChanges) GetProposerSlashings() []*v1alpha1.ProposerSlashing {
	if m != nil {
		return m.ProposerSlashings
	}
	return nil
}

func (m *ArchivedActiveSetChanges) GetAttesterSlashings() []*v1alpha1.AttesterSlashing {
	if m != nil {
		return m.AttesterSlashings
	}
	return nil
}

type ArchivedCommitteeInfo struct {
	ProposerSeed         []byte   `protobuf:"bytes,1,opt,name=proposer_seed,json=proposerSeed,proto3" json:"proposer_seed,omitempty" ssz-size:"32"`
	AttesterSeed         []byte   `protobuf:"bytes,2,opt,name=attester_seed,json=attesterSeed,proto3" json:"attester_seed,omitempty" ssz-size:"32"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ArchivedCommitteeInfo) Reset()         { *m = ArchivedCommitteeInfo{} }
func (m *ArchivedCommitteeInfo) String() string { return proto.CompactTextString(m) }
func (*ArchivedCommitteeInfo) ProtoMessage()    {}
func (*ArchivedCommitteeInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_289929478e9672a3, []int{1}
}
func (m *ArchivedCommitteeInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ArchivedCommitteeInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ArchivedCommitteeInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ArchivedCommitteeInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ArchivedCommitteeInfo.Merge(m, src)
}
func (m *ArchivedCommitteeInfo) XXX_Size() int {
	return m.Size()
}
func (m *ArchivedCommitteeInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ArchivedCommitteeInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ArchivedCommitteeInfo proto.InternalMessageInfo

func (m *ArchivedCommitteeInfo) GetProposerSeed() []byte {
	if m != nil {
		return m.ProposerSeed
	}
	return nil
}

func (m *ArchivedCommitteeInfo) GetAttesterSeed() []byte {
	if m != nil {
		return m.AttesterSeed
	}
	return nil
}

func init() {
	proto.RegisterType((*ArchivedActiveSetChanges)(nil), "ethereum.beacon.p2p.v1.ArchivedActiveSetChanges")
	proto.RegisterType((*ArchivedCommitteeInfo)(nil), "ethereum.beacon.p2p.v1.ArchivedCommitteeInfo")
}

func init() { proto.RegisterFile("proto/beacon/p2p/v1/archive.proto", fileDescriptor_289929478e9672a3) }

var fileDescriptor_289929478e9672a3 = []byte{
	// 396 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0x41, 0x8b, 0xd3, 0x40,
	0x14, 0xc7, 0xc9, 0xee, 0xd2, 0xd5, 0xb1, 0xab, 0xec, 0x80, 0x4b, 0x58, 0xa4, 0x5b, 0x83, 0x60,
	0x2f, 0x9d, 0x90, 0x14, 0x3c, 0x78, 0x6b, 0x8b, 0x07, 0x0f, 0x82, 0xa4, 0xd0, 0x6b, 0x99, 0x24,
	0xaf, 0x99, 0xc1, 0x24, 0x33, 0x64, 0x26, 0xa1, 0xf6, 0x0b, 0xf8, 0xd5, 0x3c, 0xfa, 0x09, 0x44,
	0x7a, 0xf3, 0xea, 0x27, 0x90, 0x4c, 0x92, 0x46, 0x0b, 0xdd, 0x5b, 0xde, 0x7b, 0xff, 0xff, 0xef,
	0xfd, 0x67, 0x32, 0xe8, 0xb5, 0x2c, 0x84, 0x16, 0x6e, 0x08, 0x34, 0x12, 0xb9, 0x2b, 0x7d, 0xe9,
	0x56, 0x9e, 0x4b, 0x8b, 0x88, 0xf1, 0x0a, 0x88, 0x99, 0xe1, 0x3b, 0xd0, 0x0c, 0x0a, 0x28, 0x33,
	0xd2, 0xa8, 0x88, 0xf4, 0x25, 0xa9, 0xbc, 0xfb, 0x69, 0xc2, 0x35, 0x2b, 0x43, 0x12, 0x89, 0xcc,
	0x4d, 0x44, 0x22, 0x5c, 0x23, 0x0f, 0xcb, 0xad, 0xa9, 0x1a, 0x6e, 0xfd, 0xd5, 0x60, 0xee, 0x1f,
	0x40, 0x33, 0xb7, 0xf2, 0x68, 0x2a, 0x19, 0xf5, 0xda, 0x85, 0x9b, 0x30, 0x15, 0xd1, 0x97, 0x46,
	0xe0, 0xfc, 0xbe, 0x40, 0xf6, 0xbc, 0xd9, 0x1c, 0xcf, 0x23, 0xcd, 0x2b, 0x58, 0x81, 0x5e, 0x32,
	0x9a, 0x27, 0xa0, 0xf0, 0x2b, 0xf4, 0x94, 0xd6, 0x3d, 0xaa, 0x21, 0xb6, 0xad, 0xf1, 0xe5, 0xe4,
	0x2a, 0xe8, 0x1b, 0xf8, 0x0e, 0x0d, 0x60, 0xc7, 0xeb, 0xd1, 0x85, 0x19, 0xb5, 0x15, 0xb6, 0xd1,
	0xb5, 0x4a, 0xa9, 0x62, 0x10, 0xdb, 0x57, 0x66, 0xd0, 0x95, 0xf8, 0x13, 0x7a, 0x51, 0x89, 0xb4,
	0xcc, 0x35, 0x2d, 0xbe, 0x6e, 0x6a, 0xb5, 0xb2, 0x07, 0xe3, 0xcb, 0xc9, 0x33, 0xff, 0x0d, 0x39,
	0x1e, 0x17, 0x34, 0x23, 0x5d, 0x60, 0xb2, 0xee, 0xd4, 0x1f, 0x76, 0x5c, 0x07, 0xcf, 0xab, 0x7f,
	0x4b, 0x85, 0xd7, 0x08, 0xcb, 0x42, 0x48, 0xa1, 0xa0, 0xd8, 0x98, 0x15, 0x3c, 0x4f, 0x94, 0x7d,
	0x6d, 0x88, 0x6f, 0xcf, 0x10, 0x3f, 0xb7, 0x86, 0x55, 0xab, 0x0f, 0x6e, 0xe5, 0x49, 0xc7, 0x70,
	0xa9, 0xd6, 0xa0, 0xf4, 0x7f, 0xdc, 0x27, 0x8f, 0x72, 0xe7, 0xad, 0xa1, 0xe7, 0xd2, 0x93, 0x8e,
	0x72, 0xbe, 0x59, 0xe8, 0x65, 0x77, 0xd7, 0x4b, 0x91, 0x65, 0x5c, 0x6b, 0x80, 0x8f, 0xf9, 0x56,
	0xe0, 0x77, 0xe8, 0xa6, 0x3f, 0x09, 0x98, 0xcb, 0xb6, 0x26, 0xc3, 0xc5, 0xed, 0x9f, 0x9f, 0x0f,
	0x37, 0x4a, 0xed, 0xa7, 0x8a, 0xef, 0xe1, 0xbd, 0x33, 0xf3, 0x9d, 0x60, 0x78, 0x8c, 0x0b, 0x10,
	0xd7, 0xbe, 0x3e, 0x29, 0x98, 0x3f, 0x71, 0xce, 0x77, 0x8c, 0x03, 0x10, 0x2f, 0x86, 0xdf, 0x0f,
	0x23, 0xeb, 0xc7, 0x61, 0x64, 0xfd, 0x3a, 0x8c, 0xac, 0x70, 0x60, 0x9e, 0xc2, 0xec, 0x6f, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x33, 0xde, 0x85, 0x08, 0x97, 0x02, 0x00, 0x00,
}

func (m *ArchivedActiveSetChanges) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ArchivedActiveSetChanges) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Activated) > 0 {
		dAtA2 := make([]byte, len(m.Activated)*10)
		var j1 int
		for _, num := range m.Activated {
			for num >= 1<<7 {
				dAtA2[j1] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j1++
			}
			dAtA2[j1] = uint8(num)
			j1++
		}
		dAtA[i] = 0xa
		i++
		i = encodeVarintArchive(dAtA, i, uint64(j1))
		i += copy(dAtA[i:], dAtA2[:j1])
	}
	if len(m.Exited) > 0 {
		dAtA4 := make([]byte, len(m.Exited)*10)
		var j3 int
		for _, num := range m.Exited {
			for num >= 1<<7 {
				dAtA4[j3] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j3++
			}
			dAtA4[j3] = uint8(num)
			j3++
		}
		dAtA[i] = 0x12
		i++
		i = encodeVarintArchive(dAtA, i, uint64(j3))
		i += copy(dAtA[i:], dAtA4[:j3])
	}
	if len(m.Slashed) > 0 {
		dAtA6 := make([]byte, len(m.Slashed)*10)
		var j5 int
		for _, num := range m.Slashed {
			for num >= 1<<7 {
				dAtA6[j5] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j5++
			}
			dAtA6[j5] = uint8(num)
			j5++
		}
		dAtA[i] = 0x22
		i++
		i = encodeVarintArchive(dAtA, i, uint64(j5))
		i += copy(dAtA[i:], dAtA6[:j5])
	}
	if len(m.VoluntaryExits) > 0 {
		for _, msg := range m.VoluntaryExits {
			dAtA[i] = 0x32
			i++
			i = encodeVarintArchive(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if len(m.ProposerSlashings) > 0 {
		for _, msg := range m.ProposerSlashings {
			dAtA[i] = 0x3a
			i++
			i = encodeVarintArchive(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if len(m.AttesterSlashings) > 0 {
		for _, msg := range m.AttesterSlashings {
			dAtA[i] = 0x42
			i++
			i = encodeVarintArchive(dAtA, i, uint64(msg.Size()))
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

func (m *ArchivedCommitteeInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ArchivedCommitteeInfo) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.ProposerSeed) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintArchive(dAtA, i, uint64(len(m.ProposerSeed)))
		i += copy(dAtA[i:], m.ProposerSeed)
	}
	if len(m.AttesterSeed) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintArchive(dAtA, i, uint64(len(m.AttesterSeed)))
		i += copy(dAtA[i:], m.AttesterSeed)
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintArchive(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *ArchivedActiveSetChanges) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Activated) > 0 {
		l = 0
		for _, e := range m.Activated {
			l += sovArchive(uint64(e))
		}
		n += 1 + sovArchive(uint64(l)) + l
	}
	if len(m.Exited) > 0 {
		l = 0
		for _, e := range m.Exited {
			l += sovArchive(uint64(e))
		}
		n += 1 + sovArchive(uint64(l)) + l
	}
	if len(m.Slashed) > 0 {
		l = 0
		for _, e := range m.Slashed {
			l += sovArchive(uint64(e))
		}
		n += 1 + sovArchive(uint64(l)) + l
	}
	if len(m.VoluntaryExits) > 0 {
		for _, e := range m.VoluntaryExits {
			l = e.Size()
			n += 1 + l + sovArchive(uint64(l))
		}
	}
	if len(m.ProposerSlashings) > 0 {
		for _, e := range m.ProposerSlashings {
			l = e.Size()
			n += 1 + l + sovArchive(uint64(l))
		}
	}
	if len(m.AttesterSlashings) > 0 {
		for _, e := range m.AttesterSlashings {
			l = e.Size()
			n += 1 + l + sovArchive(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *ArchivedCommitteeInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ProposerSeed)
	if l > 0 {
		n += 1 + l + sovArchive(uint64(l))
	}
	l = len(m.AttesterSeed)
	if l > 0 {
		n += 1 + l + sovArchive(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovArchive(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozArchive(x uint64) (n int) {
	return sovArchive(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ArchivedActiveSetChanges) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowArchive
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
			return fmt.Errorf("proto: ArchivedActiveSetChanges: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ArchivedActiveSetChanges: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType == 0 {
				var v uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowArchive
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.Activated = append(m.Activated, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowArchive
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthArchive
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthArchive
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				var count int
				for _, integer := range dAtA[iNdEx:postIndex] {
					if integer < 128 {
						count++
					}
				}
				elementCount = count
				if elementCount != 0 && len(m.Activated) == 0 {
					m.Activated = make([]uint64, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowArchive
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.Activated = append(m.Activated, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field Activated", wireType)
			}
		case 2:
			if wireType == 0 {
				var v uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowArchive
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.Exited = append(m.Exited, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowArchive
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthArchive
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthArchive
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				var count int
				for _, integer := range dAtA[iNdEx:postIndex] {
					if integer < 128 {
						count++
					}
				}
				elementCount = count
				if elementCount != 0 && len(m.Exited) == 0 {
					m.Exited = make([]uint64, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowArchive
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.Exited = append(m.Exited, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field Exited", wireType)
			}
		case 4:
			if wireType == 0 {
				var v uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowArchive
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.Slashed = append(m.Slashed, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowArchive
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthArchive
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthArchive
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				var count int
				for _, integer := range dAtA[iNdEx:postIndex] {
					if integer < 128 {
						count++
					}
				}
				elementCount = count
				if elementCount != 0 && len(m.Slashed) == 0 {
					m.Slashed = make([]uint64, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowArchive
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.Slashed = append(m.Slashed, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field Slashed", wireType)
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VoluntaryExits", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowArchive
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
				return ErrInvalidLengthArchive
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthArchive
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.VoluntaryExits = append(m.VoluntaryExits, &v1alpha1.VoluntaryExit{})
			if err := m.VoluntaryExits[len(m.VoluntaryExits)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProposerSlashings", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowArchive
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
				return ErrInvalidLengthArchive
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthArchive
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ProposerSlashings = append(m.ProposerSlashings, &v1alpha1.ProposerSlashing{})
			if err := m.ProposerSlashings[len(m.ProposerSlashings)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AttesterSlashings", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowArchive
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
				return ErrInvalidLengthArchive
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthArchive
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AttesterSlashings = append(m.AttesterSlashings, &v1alpha1.AttesterSlashing{})
			if err := m.AttesterSlashings[len(m.AttesterSlashings)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipArchive(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthArchive
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthArchive
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
func (m *ArchivedCommitteeInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowArchive
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
			return fmt.Errorf("proto: ArchivedCommitteeInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ArchivedCommitteeInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProposerSeed", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowArchive
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthArchive
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthArchive
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ProposerSeed = append(m.ProposerSeed[:0], dAtA[iNdEx:postIndex]...)
			if m.ProposerSeed == nil {
				m.ProposerSeed = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AttesterSeed", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowArchive
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthArchive
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthArchive
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AttesterSeed = append(m.AttesterSeed[:0], dAtA[iNdEx:postIndex]...)
			if m.AttesterSeed == nil {
				m.AttesterSeed = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipArchive(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthArchive
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthArchive
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
func skipArchive(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowArchive
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
					return 0, ErrIntOverflowArchive
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
					return 0, ErrIntOverflowArchive
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
				return 0, ErrInvalidLengthArchive
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthArchive
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowArchive
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
				next, err := skipArchive(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthArchive
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
	ErrInvalidLengthArchive = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowArchive   = fmt.Errorf("proto: integer overflow")
)
