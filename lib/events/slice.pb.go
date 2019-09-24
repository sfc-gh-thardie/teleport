// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: slice.proto

package events

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import empty "github.com/golang/protobuf/ptypes/empty"

import context "golang.org/x/net/context"
import grpc "google.golang.org/grpc"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// SessionSlice is a slice of submitted chunks
type SessionSlice struct {
	// Namespace is a session namespace
	Namespace string `protobuf:"bytes,1,opt,name=Namespace,proto3" json:"Namespace,omitempty"`
	// SessionID is a session ID associated with this chunk
	SessionID string `protobuf:"bytes,2,opt,name=SessionID,proto3" json:"SessionID,omitempty"`
	// Chunks is a list of submitted session chunks
	Chunks []*SessionChunk `protobuf:"bytes,3,rep,name=Chunks" json:"Chunks,omitempty"`
	// Version specifies session slice version
	Version              int64    `protobuf:"varint,4,opt,name=Version,proto3" json:"Version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SessionSlice) Reset()         { *m = SessionSlice{} }
func (m *SessionSlice) String() string { return proto.CompactTextString(m) }
func (*SessionSlice) ProtoMessage()    {}
func (*SessionSlice) Descriptor() ([]byte, []int) {
	return fileDescriptor_slice_c27e90ee7997b6d0, []int{0}
}
func (m *SessionSlice) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SessionSlice) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SessionSlice.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *SessionSlice) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SessionSlice.Merge(dst, src)
}
func (m *SessionSlice) XXX_Size() int {
	return m.Size()
}
func (m *SessionSlice) XXX_DiscardUnknown() {
	xxx_messageInfo_SessionSlice.DiscardUnknown(m)
}

var xxx_messageInfo_SessionSlice proto.InternalMessageInfo

func (m *SessionSlice) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *SessionSlice) GetSessionID() string {
	if m != nil {
		return m.SessionID
	}
	return ""
}

func (m *SessionSlice) GetChunks() []*SessionChunk {
	if m != nil {
		return m.Chunks
	}
	return nil
}

func (m *SessionSlice) GetVersion() int64 {
	if m != nil {
		return m.Version
	}
	return 0
}

// SessionChunk is a chunk to be posted in the context of the session
type SessionChunk struct {
	// Time is the occurence of this event
	Time int64 `protobuf:"varint,2,opt,name=Time,proto3" json:"Time,omitempty"`
	// Data is captured data, contains event fields in case of event, session data otherwise
	Data []byte `protobuf:"bytes,3,opt,name=Data,proto3" json:"Data,omitempty"`
	// EventType is event type
	EventType string `protobuf:"bytes,4,opt,name=EventType,proto3" json:"EventType,omitempty"`
	// EventIndex is the event global index
	EventIndex int64 `protobuf:"varint,5,opt,name=EventIndex,proto3" json:"EventIndex,omitempty"`
	// Index is the autoincremented chunk index
	ChunkIndex int64 `protobuf:"varint,6,opt,name=ChunkIndex,proto3" json:"ChunkIndex,omitempty"`
	// Offset is an offset from the previous chunk in bytes
	Offset int64 `protobuf:"varint,7,opt,name=Offset,proto3" json:"Offset,omitempty"`
	// Delay is a delay from the previous event in milliseconds
	Delay                int64    `protobuf:"varint,8,opt,name=Delay,proto3" json:"Delay,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SessionChunk) Reset()         { *m = SessionChunk{} }
func (m *SessionChunk) String() string { return proto.CompactTextString(m) }
func (*SessionChunk) ProtoMessage()    {}
func (*SessionChunk) Descriptor() ([]byte, []int) {
	return fileDescriptor_slice_c27e90ee7997b6d0, []int{1}
}
func (m *SessionChunk) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SessionChunk) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SessionChunk.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *SessionChunk) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SessionChunk.Merge(dst, src)
}
func (m *SessionChunk) XXX_Size() int {
	return m.Size()
}
func (m *SessionChunk) XXX_DiscardUnknown() {
	xxx_messageInfo_SessionChunk.DiscardUnknown(m)
}

var xxx_messageInfo_SessionChunk proto.InternalMessageInfo

func (m *SessionChunk) GetTime() int64 {
	if m != nil {
		return m.Time
	}
	return 0
}

func (m *SessionChunk) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *SessionChunk) GetEventType() string {
	if m != nil {
		return m.EventType
	}
	return ""
}

func (m *SessionChunk) GetEventIndex() int64 {
	if m != nil {
		return m.EventIndex
	}
	return 0
}

func (m *SessionChunk) GetChunkIndex() int64 {
	if m != nil {
		return m.ChunkIndex
	}
	return 0
}

func (m *SessionChunk) GetOffset() int64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *SessionChunk) GetDelay() int64 {
	if m != nil {
		return m.Delay
	}
	return 0
}

func init() {
	proto.RegisterType((*SessionSlice)(nil), "events.SessionSlice")
	proto.RegisterType((*SessionChunk)(nil), "events.SessionChunk")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for AuditLog service

type AuditLogClient interface {
	SubmitSessionSlice(ctx context.Context, opts ...grpc.CallOption) (AuditLog_SubmitSessionSliceClient, error)
}

type auditLogClient struct {
	cc *grpc.ClientConn
}

func NewAuditLogClient(cc *grpc.ClientConn) AuditLogClient {
	return &auditLogClient{cc}
}

func (c *auditLogClient) SubmitSessionSlice(ctx context.Context, opts ...grpc.CallOption) (AuditLog_SubmitSessionSliceClient, error) {
	stream, err := c.cc.NewStream(ctx, &_AuditLog_serviceDesc.Streams[0], "/events.AuditLog/SubmitSessionSlice", opts...)
	if err != nil {
		return nil, err
	}
	x := &auditLogSubmitSessionSliceClient{stream}
	return x, nil
}

type AuditLog_SubmitSessionSliceClient interface {
	Send(*SessionSlice) error
	CloseAndRecv() (*empty.Empty, error)
	grpc.ClientStream
}

type auditLogSubmitSessionSliceClient struct {
	grpc.ClientStream
}

func (x *auditLogSubmitSessionSliceClient) Send(m *SessionSlice) error {
	return x.ClientStream.SendMsg(m)
}

func (x *auditLogSubmitSessionSliceClient) CloseAndRecv() (*empty.Empty, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(empty.Empty)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for AuditLog service

type AuditLogServer interface {
	SubmitSessionSlice(AuditLog_SubmitSessionSliceServer) error
}

func RegisterAuditLogServer(s *grpc.Server, srv AuditLogServer) {
	s.RegisterService(&_AuditLog_serviceDesc, srv)
}

func _AuditLog_SubmitSessionSlice_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(AuditLogServer).SubmitSessionSlice(&auditLogSubmitSessionSliceServer{stream})
}

type AuditLog_SubmitSessionSliceServer interface {
	SendAndClose(*empty.Empty) error
	Recv() (*SessionSlice, error)
	grpc.ServerStream
}

type auditLogSubmitSessionSliceServer struct {
	grpc.ServerStream
}

func (x *auditLogSubmitSessionSliceServer) SendAndClose(m *empty.Empty) error {
	return x.ServerStream.SendMsg(m)
}

func (x *auditLogSubmitSessionSliceServer) Recv() (*SessionSlice, error) {
	m := new(SessionSlice)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _AuditLog_serviceDesc = grpc.ServiceDesc{
	ServiceName: "events.AuditLog",
	HandlerType: (*AuditLogServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SubmitSessionSlice",
			Handler:       _AuditLog_SubmitSessionSlice_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "slice.proto",
}

func (m *SessionSlice) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SessionSlice) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Namespace) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintSlice(dAtA, i, uint64(len(m.Namespace)))
		i += copy(dAtA[i:], m.Namespace)
	}
	if len(m.SessionID) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintSlice(dAtA, i, uint64(len(m.SessionID)))
		i += copy(dAtA[i:], m.SessionID)
	}
	if len(m.Chunks) > 0 {
		for _, msg := range m.Chunks {
			dAtA[i] = 0x1a
			i++
			i = encodeVarintSlice(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if m.Version != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintSlice(dAtA, i, uint64(m.Version))
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *SessionChunk) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SessionChunk) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Time != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintSlice(dAtA, i, uint64(m.Time))
	}
	if len(m.Data) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintSlice(dAtA, i, uint64(len(m.Data)))
		i += copy(dAtA[i:], m.Data)
	}
	if len(m.EventType) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintSlice(dAtA, i, uint64(len(m.EventType)))
		i += copy(dAtA[i:], m.EventType)
	}
	if m.EventIndex != 0 {
		dAtA[i] = 0x28
		i++
		i = encodeVarintSlice(dAtA, i, uint64(m.EventIndex))
	}
	if m.ChunkIndex != 0 {
		dAtA[i] = 0x30
		i++
		i = encodeVarintSlice(dAtA, i, uint64(m.ChunkIndex))
	}
	if m.Offset != 0 {
		dAtA[i] = 0x38
		i++
		i = encodeVarintSlice(dAtA, i, uint64(m.Offset))
	}
	if m.Delay != 0 {
		dAtA[i] = 0x40
		i++
		i = encodeVarintSlice(dAtA, i, uint64(m.Delay))
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintSlice(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *SessionSlice) Size() (n int) {
	var l int
	_ = l
	l = len(m.Namespace)
	if l > 0 {
		n += 1 + l + sovSlice(uint64(l))
	}
	l = len(m.SessionID)
	if l > 0 {
		n += 1 + l + sovSlice(uint64(l))
	}
	if len(m.Chunks) > 0 {
		for _, e := range m.Chunks {
			l = e.Size()
			n += 1 + l + sovSlice(uint64(l))
		}
	}
	if m.Version != 0 {
		n += 1 + sovSlice(uint64(m.Version))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *SessionChunk) Size() (n int) {
	var l int
	_ = l
	if m.Time != 0 {
		n += 1 + sovSlice(uint64(m.Time))
	}
	l = len(m.Data)
	if l > 0 {
		n += 1 + l + sovSlice(uint64(l))
	}
	l = len(m.EventType)
	if l > 0 {
		n += 1 + l + sovSlice(uint64(l))
	}
	if m.EventIndex != 0 {
		n += 1 + sovSlice(uint64(m.EventIndex))
	}
	if m.ChunkIndex != 0 {
		n += 1 + sovSlice(uint64(m.ChunkIndex))
	}
	if m.Offset != 0 {
		n += 1 + sovSlice(uint64(m.Offset))
	}
	if m.Delay != 0 {
		n += 1 + sovSlice(uint64(m.Delay))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovSlice(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozSlice(x uint64) (n int) {
	return sovSlice(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *SessionSlice) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSlice
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: SessionSlice: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SessionSlice: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Namespace", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSlice
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSlice
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Namespace = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SessionID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSlice
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSlice
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SessionID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Chunks", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSlice
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthSlice
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Chunks = append(m.Chunks, &SessionChunk{})
			if err := m.Chunks[len(m.Chunks)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Version", wireType)
			}
			m.Version = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSlice
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Version |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipSlice(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSlice
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
func (m *SessionChunk) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSlice
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: SessionChunk: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SessionChunk: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Time", wireType)
			}
			m.Time = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSlice
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Time |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSlice
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthSlice
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = append(m.Data[:0], dAtA[iNdEx:postIndex]...)
			if m.Data == nil {
				m.Data = []byte{}
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EventType", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSlice
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSlice
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EventType = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EventIndex", wireType)
			}
			m.EventIndex = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSlice
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.EventIndex |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChunkIndex", wireType)
			}
			m.ChunkIndex = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSlice
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ChunkIndex |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Offset", wireType)
			}
			m.Offset = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSlice
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Offset |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Delay", wireType)
			}
			m.Delay = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSlice
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Delay |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipSlice(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSlice
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
func skipSlice(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSlice
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
					return 0, ErrIntOverflowSlice
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
					return 0, ErrIntOverflowSlice
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthSlice
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowSlice
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
				next, err := skipSlice(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
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
	ErrInvalidLengthSlice = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSlice   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("slice.proto", fileDescriptor_slice_c27e90ee7997b6d0) }

var fileDescriptor_slice_c27e90ee7997b6d0 = []byte{
	// 329 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x91, 0xc1, 0x4e, 0xc2, 0x40,
	0x10, 0x86, 0x59, 0x0b, 0x05, 0x16, 0x0e, 0x66, 0x43, 0xc8, 0x06, 0x4d, 0xd3, 0x70, 0xea, 0xc1,
	0x94, 0x04, 0x9f, 0x40, 0x05, 0x13, 0x12, 0xa3, 0xc9, 0x42, 0xbc, 0x17, 0x18, 0xb0, 0x91, 0x76,
	0x1b, 0x76, 0x6b, 0xe4, 0x35, 0x3c, 0xf9, 0x3c, 0x9e, 0x3c, 0xfa, 0x08, 0x06, 0x5f, 0xc4, 0xec,
	0x6c, 0xb1, 0xea, 0x6d, 0xff, 0xef, 0xff, 0x77, 0x66, 0x32, 0x43, 0x5b, 0x6a, 0x13, 0x2f, 0x20,
	0xcc, 0xb6, 0x52, 0x4b, 0xe6, 0xc2, 0x13, 0xa4, 0x5a, 0xf5, 0x4e, 0xd6, 0x52, 0xae, 0x37, 0x30,
	0x40, 0x3a, 0xcf, 0x57, 0x03, 0x48, 0x32, 0xbd, 0xb3, 0xa1, 0xfe, 0x0b, 0xa1, 0xed, 0x29, 0x28,
	0x15, 0xcb, 0x74, 0x6a, 0xfe, 0xb2, 0x53, 0xda, 0xbc, 0x8d, 0x12, 0x50, 0x59, 0xb4, 0x00, 0x4e,
	0x7c, 0x12, 0x34, 0x45, 0x09, 0x8c, 0x5b, 0xa4, 0x27, 0x23, 0x7e, 0x64, 0xdd, 0x1f, 0xc0, 0xce,
	0xa8, 0x7b, 0xf5, 0x90, 0xa7, 0x8f, 0x8a, 0x3b, 0xbe, 0x13, 0xb4, 0x86, 0x9d, 0xd0, 0x8e, 0x10,
	0x16, 0x11, 0x34, 0x45, 0x91, 0x61, 0x9c, 0xd6, 0xef, 0x61, 0x6b, 0x38, 0xaf, 0xfa, 0x24, 0x70,
	0xc4, 0x41, 0xf6, 0xdf, 0xca, 0xa1, 0x30, 0xcb, 0x18, 0xad, 0xce, 0xe2, 0x04, 0xb0, 0xa3, 0x23,
	0xf0, 0x6d, 0xd8, 0x28, 0xd2, 0x11, 0x77, 0x7c, 0x12, 0xb4, 0x05, 0xbe, 0xcd, 0x78, 0x63, 0xd3,
	0x71, 0xb6, 0xcb, 0x00, 0x8b, 0x36, 0x45, 0x09, 0x98, 0x47, 0x29, 0x8a, 0x49, 0xba, 0x84, 0x67,
	0x5e, 0xc3, 0x5a, 0xbf, 0x88, 0xf1, 0xb1, 0x9d, 0xf5, 0x5d, 0xeb, 0x97, 0x84, 0x75, 0xa9, 0x7b,
	0xb7, 0x5a, 0x29, 0xd0, 0xbc, 0x8e, 0x5e, 0xa1, 0x58, 0x87, 0xd6, 0x46, 0xb0, 0x89, 0x76, 0xbc,
	0x81, 0xd8, 0x8a, 0xa1, 0xa0, 0x8d, 0x8b, 0x7c, 0x19, 0xeb, 0x1b, 0xb9, 0x66, 0xd7, 0x94, 0x4d,
	0xf3, 0x79, 0x12, 0xeb, 0x3f, 0xab, 0xfe, 0xbf, 0x1e, 0xa4, 0xbd, 0x6e, 0x68, 0xef, 0x15, 0x1e,
	0xee, 0x15, 0x8e, 0xcd, 0xbd, 0xfa, 0x95, 0x80, 0x5c, 0x1e, 0xbf, 0xef, 0x3d, 0xf2, 0xb1, 0xf7,
	0xc8, 0xe7, 0xde, 0x23, 0xaf, 0x5f, 0x5e, 0x65, 0xee, 0x62, 0xea, 0xfc, 0x3b, 0x00, 0x00, 0xff,
	0xff, 0xa5, 0x2e, 0xde, 0xc9, 0xfa, 0x01, 0x00, 0x00,
}