// automatically generated, do not modify

package test

import (
	flatbuffers "github.com/google/flatbuffers/go"
)
type DvMonster struct {
	_tab flatbuffers.Table
}

func GetRootAsDvMonster(buf []byte, offset flatbuffers.UOffsetT) *DvMonster {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &DvMonster{}
	x.Init(buf, n + offset)
	return x
}

func (rcv *DvMonster) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *DvMonster) Id() uint64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetUint64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *DvMonster) Name() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *DvMonster) Flag() byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetByte(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *DvMonster) List(j int) uint64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetUint64(a + flatbuffers.UOffsetT(j * 8))
	}
	return 0
}

func (rcv *DvMonster) ListLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *DvMonster) Kv(obj *KV) *KV {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		x := o + rcv._tab.Pos
		if obj == nil {
			obj = new(KV)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func DvMonsterStart(builder *flatbuffers.Builder) { builder.StartObject(5) }
func DvMonsterAddId(builder *flatbuffers.Builder, id uint64) { builder.PrependUint64Slot(0, id, 0) }
func DvMonsterAddName(builder *flatbuffers.Builder, name flatbuffers.UOffsetT) { builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(name), 0) }
func DvMonsterAddFlag(builder *flatbuffers.Builder, flag byte) { builder.PrependByteSlot(2, flag, 0) }
func DvMonsterAddList(builder *flatbuffers.Builder, list flatbuffers.UOffsetT) { builder.PrependUOffsetTSlot(3, flatbuffers.UOffsetT(list), 0) }
func DvMonsterStartListVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT { return builder.StartVector(8, numElems, 8)
}
func DvMonsterAddKv(builder *flatbuffers.Builder, kv flatbuffers.UOffsetT) { builder.PrependStructSlot(4, flatbuffers.UOffsetT(kv), 0) }
func DvMonsterEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT { return builder.EndObject() }
