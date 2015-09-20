// Code generated by protoc-gen-go.
// source: proto/tx.proto
// DO NOT EDIT!

/*
Package proto is a generated protocol buffer package.

It is generated from these files:
	proto/tx.proto

It has these top-level messages:
	TxMsg
	InitialReport
	Report
*/
package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type TxMsg struct {
	Id  string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Sig []byte `protobuf:"bytes,2,opt,name=sig,proto3" json:"sig,omitempty"`
	Seq int32  `protobuf:"varint,3,opt,name=seq" json:"seq,omitempty"`
	// Types that are valid to be assigned to Inner:
	//	*TxMsg_InitialReport
	//	*TxMsg_Report
	Inner isTxMsg_Inner `protobuf_oneof:"inner"`
}

func (m *TxMsg) Reset()         { *m = TxMsg{} }
func (m *TxMsg) String() string { return proto1.CompactTextString(m) }
func (*TxMsg) ProtoMessage()    {}

type isTxMsg_Inner interface {
	isTxMsg_Inner()
}

type TxMsg_InitialReport struct {
	InitialReport *InitialReport `protobuf:"bytes,4,opt,name=initialReport,oneof"`
}
type TxMsg_Report struct {
	Report *Report `protobuf:"bytes,5,opt,name=report,oneof"`
}

func (*TxMsg_InitialReport) isTxMsg_Inner() {}
func (*TxMsg_Report) isTxMsg_Inner()        {}

func (m *TxMsg) GetInner() isTxMsg_Inner {
	if m != nil {
		return m.Inner
	}
	return nil
}

func (m *TxMsg) GetInitialReport() *InitialReport {
	if x, ok := m.GetInner().(*TxMsg_InitialReport); ok {
		return x.InitialReport
	}
	return nil
}

func (m *TxMsg) GetReport() *Report {
	if x, ok := m.GetInner().(*TxMsg_Report); ok {
		return x.Report
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*TxMsg) XXX_OneofFuncs() (func(msg proto1.Message, b *proto1.Buffer) error, func(msg proto1.Message, tag, wire int, b *proto1.Buffer) (bool, error), []interface{}) {
	return _TxMsg_OneofMarshaler, _TxMsg_OneofUnmarshaler, []interface{}{
		(*TxMsg_InitialReport)(nil),
		(*TxMsg_Report)(nil),
	}
}

func _TxMsg_OneofMarshaler(msg proto1.Message, b *proto1.Buffer) error {
	m := msg.(*TxMsg)
	// inner
	switch x := m.Inner.(type) {
	case *TxMsg_InitialReport:
		b.EncodeVarint(4<<3 | proto1.WireBytes)
		if err := b.EncodeMessage(x.InitialReport); err != nil {
			return err
		}
	case *TxMsg_Report:
		b.EncodeVarint(5<<3 | proto1.WireBytes)
		if err := b.EncodeMessage(x.Report); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("TxMsg.Inner has unexpected type %T", x)
	}
	return nil
}

func _TxMsg_OneofUnmarshaler(msg proto1.Message, tag, wire int, b *proto1.Buffer) (bool, error) {
	m := msg.(*TxMsg)
	switch tag {
	case 4: // inner.initialReport
		if wire != proto1.WireBytes {
			return true, proto1.ErrInternalBadWireType
		}
		msg := new(InitialReport)
		err := b.DecodeMessage(msg)
		m.Inner = &TxMsg_InitialReport{msg}
		return true, err
	case 5: // inner.report
		if wire != proto1.WireBytes {
			return true, proto1.ErrInternalBadWireType
		}
		msg := new(Report)
		err := b.DecodeMessage(msg)
		m.Inner = &TxMsg_Report{msg}
		return true, err
	default:
		return false, nil
	}
}

type InitialReport struct {
	Time int64  `protobuf:"varint,1,opt,name=time" json:"time,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
}

func (m *InitialReport) Reset()         { *m = InitialReport{} }
func (m *InitialReport) String() string { return proto1.CompactTextString(m) }
func (*InitialReport) ProtoMessage()    {}

type Report struct {
	Diff int64 `protobuf:"varint,1,opt,name=diff" json:"diff,omitempty"`
	Time int64 `protobuf:"varint,2,opt,name=time" json:"time,omitempty"`
}

func (m *Report) Reset()         { *m = Report{} }
func (m *Report) String() string { return proto1.CompactTextString(m) }
func (*Report) ProtoMessage()    {}