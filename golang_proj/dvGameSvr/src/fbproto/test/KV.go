// automatically generated, do not modify

package test

import (
	flatbuffers "github.com/google/flatbuffers/go"
)
type KV struct {
	_tab flatbuffers.Struct
}

func (rcv *KV) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *KV) Key() uint64 { return rcv._tab.GetUint64(rcv._tab.Pos + flatbuffers.UOffsetT(0)) }
func (rcv *KV) Value() float64 { return rcv._tab.GetFloat64(rcv._tab.Pos + flatbuffers.UOffsetT(8)) }

func CreateKV(builder *flatbuffers.Builder, key uint64, value float64) flatbuffers.UOffsetT {
    builder.Prep(8, 16)
    builder.PrependFloat64(value)
    builder.PrependUint64(key)
    return builder.Offset()
}
