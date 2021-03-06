// Code generated by protoc-gen-go.
// source: player.proto
// DO NOT EDIT!

/*
Package Player is a generated protocol buffer package.

It is generated from these files:
	player.proto

It has these top-level messages:
	Player
*/
package Player

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Player struct {
	Name             *string            `protobuf:"bytes,1,req,name=name" json:"name,omitempty"`
	Id               *int32             `protobuf:"varint,2,req,name=id" json:"id,omitempty"`
	Pos              []*Player_Position `protobuf:"bytes,3,rep,name=pos" json:"pos,omitempty"`
	XXX_unrecognized []byte             `json:"-"`
}

func (m *Player) Reset()                    { *m = Player{} }
func (m *Player) String() string            { return proto.CompactTextString(m) }
func (*Player) ProtoMessage()               {}
func (*Player) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Player) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *Player) GetId() int32 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *Player) GetPos() []*Player_Position {
	if m != nil {
		return m.Pos
	}
	return nil
}

type Player_Position struct {
	X                *int32 `protobuf:"varint,1,req,name=X" json:"X,omitempty"`
	Y                *int32 `protobuf:"varint,2,req,name=Y" json:"Y,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *Player_Position) Reset()                    { *m = Player_Position{} }
func (m *Player_Position) String() string            { return proto.CompactTextString(m) }
func (*Player_Position) ProtoMessage()               {}
func (*Player_Position) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

func (m *Player_Position) GetX() int32 {
	if m != nil && m.X != nil {
		return *m.X
	}
	return 0
}

func (m *Player_Position) GetY() int32 {
	if m != nil && m.Y != nil {
		return *m.Y
	}
	return 0
}

func init() {
	proto.RegisterType((*Player)(nil), "Player.Player")
	proto.RegisterType((*Player_Position)(nil), "Player.Player.Position")
}

var fileDescriptor0 = []byte{
	// 118 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x29, 0xc8, 0x49, 0xac,
	0x4c, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x0b, 0x00, 0xf3, 0x94, 0xb2, 0xb8,
	0xa0, 0x2c, 0x21, 0x1e, 0x2e, 0x96, 0xbc, 0xc4, 0xdc, 0x54, 0x09, 0x46, 0x05, 0x26, 0x0d, 0x4e,
	0x21, 0x2e, 0x2e, 0xa6, 0xcc, 0x14, 0x09, 0x26, 0x20, 0x9b, 0x55, 0x48, 0x85, 0x8b, 0xb9, 0x20,
	0xbf, 0x58, 0x82, 0x59, 0x81, 0x59, 0x83, 0xdb, 0x48, 0x5c, 0x0f, 0xa2, 0x1e, 0x4e, 0xe5, 0x17,
	0x67, 0x96, 0x64, 0xe6, 0xe7, 0x49, 0x29, 0x70, 0x71, 0xc0, 0xd8, 0x42, 0x9c, 0x5c, 0x8c, 0x11,
	0x60, 0x83, 0x58, 0x41, 0xcc, 0x48, 0x88, 0x39, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0xb8, 0x31,
	0xca, 0xd2, 0x82, 0x00, 0x00, 0x00,
}
