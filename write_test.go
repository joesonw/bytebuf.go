package bytebuf

import (
	"bytes"
	. "gopkg.in/go-playground/assert.v1"
	"testing"
)

func TestWrite(t *testing.T) {
	ChunkSize = 5

	b, err := NewByteBuf()
	Equal(t, err, nil)
	NotEqual(t, b, nil)
	Equal(t, b.Capacity(), 5)

	p := []byte{0x01, 0x02, 0x03}
	n, err := b.Write(p)
	Equal(t, err, nil)
	Equal(t, n, 3)
	Equal(t, b.buf, []byte{0x01, 0x02, 0x03, 0x00, 0x00})

}

func TestWriteBool(t *testing.T) {
	ChunkSize = 1

	b, err := NewByteBuf()
	Equal(t, err, nil)
	NotEqual(t, b, nil)

	err = b.WriteBool(true)
	Equal(t, b.buf, []byte{0x01})
	Equal(t, err, nil)
	Equal(t, b.Capacity(), 1)

	b.Flush()
	err = b.WriteBool(false)
	Equal(t, b.buf, []byte{0x00})
	Equal(t, err, nil)
	Equal(t, b.Capacity(), 1)
}

func TestWriteByte(t *testing.T) {
	ChunkSize = 1

	b, err := NewByteBuf()
	Equal(t, err, nil)
	NotEqual(t, b, nil)

	err = b.WriteByte(0x08)
	Equal(t, b.buf, []byte{0x08})
	Equal(t, err, nil)
	Equal(t, b.Capacity(), 1)

}

func TestWriteBytes(t *testing.T) {
	ChunkSize = 1

	b, err := NewByteBuf()
	Equal(t, err, nil)
	NotEqual(t, b, nil)

	err = b.WriteBytes([]byte{0x00, 0x01, 0x02, 0x03})
	Equal(t, err, nil)
	Equal(t, b.buf, []byte{0x00, 0x01, 0x02, 0x03})
	Equal(t, b.Capacity(), 4)

	err = b.WriteBytes([]byte{0x00, 0x01, 0x02})
	Equal(t, err, nil)
	Equal(t, b.buf, []byte{0x00, 0x01, 0x02, 0x03, 0x00, 0x01, 0x02})
	Equal(t, b.Capacity(), 7)
}

func TestWriteUInt8(t *testing.T) {
	ChunkSize = 1

	b, err := NewByteBuf()
	Equal(t, err, nil)
	NotEqual(t, b, nil)
	Equal(t, b.Capacity(), 1)

	err = b.WriteUInt8(8)
	Equal(t, err, nil)
	Equal(t, b.buf, []byte{0x08})
	Equal(t, b.Capacity(), 1)
}

func TestWriteInt8(t *testing.T) {
	ChunkSize = 1

	b, err := NewByteBuf()
	Equal(t, err, nil)
	NotEqual(t, b, nil)
	Equal(t, b.Capacity(), 1)

	err = b.WriteInt8(-2)
	Equal(t, err, nil)
	Equal(t, b.buf, []byte{0xfe})
	Equal(t, b.Capacity(), 1)
}

func TestWriteUInt16LE(t *testing.T) {
	ChunkSize = 1

	b, err := NewByteBuf()
	Equal(t, err, nil)
	NotEqual(t, b, nil)

	err = b.WriteUInt16LE(264)
	Equal(t, err, nil)
	Equal(t, b.buf, []byte{0x08, 0x01})
	Equal(t, b.Capacity(), 2)
}

func TestWriteUInt16BE(t *testing.T) {
	ChunkSize = 1

	b, err := NewByteBuf()
	Equal(t, err, nil)
	NotEqual(t, b, nil)

	err = b.WriteUInt16BE(264)
	Equal(t, err, nil)
	Equal(t, b.buf, []byte{0x01, 0x08})
	Equal(t, b.Capacity(), 2)
}

func TestWriteInt16LE(t *testing.T) {
	ChunkSize = 1

	b, err := NewByteBuf()
	Equal(t, err, nil)
	NotEqual(t, b, nil)

	err = b.WriteInt16LE(-2)
	Equal(t, err, nil)
	Equal(t, b.buf, []byte{0xfe, 0xff})
	Equal(t, b.Capacity(), 2)
}

func TestWriteInt16BE(t *testing.T) {
	ChunkSize = 1

	b, err := NewByteBuf()
	Equal(t, err, nil)
	NotEqual(t, b, nil)

	err = b.WriteInt16BE(-2)
	Equal(t, err, nil)
	Equal(t, b.buf, []byte{0xff, 0xfe})
	Equal(t, b.Capacity(), 2)
}

func TestWriteUInt32LE(t *testing.T) {
	ChunkSize = 1

	b, err := NewByteBuf()
	Equal(t, err, nil)
	NotEqual(t, b, nil)

	err = b.WriteUInt32LE(67305985)
	Equal(t, err, nil)
	Equal(t, b.buf, []byte{0x01, 0x02, 0x03, 0x04})
	Equal(t, b.Capacity(), 4)
}

func TestWriteUInt32BE(t *testing.T) {
	ChunkSize = 1

	b, err := NewByteBuf()
	Equal(t, err, nil)
	NotEqual(t, b, nil)

	err = b.WriteUInt32BE(67305985)
	Equal(t, err, nil)
	Equal(t, b.buf, []byte{0x04, 0x03, 0x02, 0x01})
	Equal(t, b.Capacity(), 4)
}

func TestWriteInt32LE(t *testing.T) {
	ChunkSize = 1

	b, err := NewByteBuf()
	Equal(t, err, nil)
	NotEqual(t, b, nil)

	err = b.WriteInt32LE(-2)
	Equal(t, err, nil)
	Equal(t, b.buf, []byte{0xfe, 0xff, 0xff, 0xff})
	Equal(t, b.Capacity(), 4)
}

func TestWriteInt32BE(t *testing.T) {
	ChunkSize = 1

	b, err := NewByteBuf()
	Equal(t, err, nil)
	NotEqual(t, b, nil)

	err = b.WriteInt32BE(-2)
	Equal(t, err, nil)
	Equal(t, b.buf, []byte{0xff, 0xff, 0xff, 0xfe})
	Equal(t, b.Capacity(), 4)
}

func TestWriteUInt64LE(t *testing.T) {
	ChunkSize = 1

	b, err := NewByteBuf()
	Equal(t, err, nil)
	NotEqual(t, b, nil)

	err = b.WriteUInt64LE(67305985)
	Equal(t, err, nil)
	Equal(t, b.buf, []byte{0x01, 0x02, 0x03, 0x04, 0x00, 0x00, 0x00, 0x00})
	Equal(t, b.Capacity(), 8)
}

func TestWriteUInt64BE(t *testing.T) {
	ChunkSize = 1

	b, err := NewByteBuf()
	Equal(t, err, nil)
	NotEqual(t, b, nil)

	err = b.WriteUInt64BE(67305985)
	Equal(t, err, nil)
	Equal(t, b.buf, []byte{0x00, 0x00, 0x00, 0x00, 0x04, 0x03, 0x02, 0x01})
	Equal(t, b.Capacity(), 8)
}

func TestWriteInt64LE(t *testing.T) {
	ChunkSize = 1

	b, err := NewByteBuf()
	Equal(t, err, nil)
	NotEqual(t, b, nil)

	err = b.WriteInt64LE(-2)
	Equal(t, err, nil)
	Equal(t, b.buf, []byte{0xfe, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff})
	Equal(t, b.Capacity(), 8)
}

func TestWriteInt64BE(t *testing.T) {
	ChunkSize = 1

	b, err := NewByteBuf()
	Equal(t, err, nil)
	NotEqual(t, b, nil)

	err = b.WriteInt64BE(-2)
	Equal(t, err, nil)
	Equal(t, b.buf, []byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xfe})
	Equal(t, b.Capacity(), 8)
}

func TestWriteIntFloat32LE(t *testing.T) {
	ChunkSize = 1

	b, err := NewByteBuf()
	Equal(t, err, nil)
	NotEqual(t, b, nil)

	err = b.WriteFloat32LE(3.141592653589793)
	Equal(t, err, nil)
	Equal(t, b.buf, []byte{0xdb, 0x0f, 0x49, 0x40})
	Equal(t, b.Capacity(), 4)
}

func TestWriteIntFloat32BE(t *testing.T) {
	ChunkSize = 1

	b, err := NewByteBuf()
	Equal(t, err, nil)
	NotEqual(t, b, nil)

	err = b.WriteFloat32BE(3.141592653589793)
	Equal(t, err, nil)
	Equal(t, b.buf, []byte{0x40, 0x49, 0x0f, 0xdb})
	Equal(t, b.Capacity(), 4)
}

func TestWriteIntFloat64LE(t *testing.T) {
	ChunkSize = 1

	b, err := NewByteBuf()
	Equal(t, err, nil)
	NotEqual(t, b, nil)

	err = b.WriteFloat64LE(3.141592653589793)
	Equal(t, err, nil)
	Equal(t, b.buf, []byte{24, 45, 68, 84, 251, 33, 9, 64})
	Equal(t, b.Capacity(), 8)
}

func TestWriteIntFloat64BE(t *testing.T) {
	ChunkSize = 1

	b, err := NewByteBuf()
	Equal(t, err, nil)
	NotEqual(t, b, nil)

	err = b.WriteFloat64BE(3.141592653589793)
	Equal(t, err, nil)
	Equal(t, b.buf, []byte{64, 9, 33, 251, 84, 68, 45, 24})
	Equal(t, b.Capacity(), 8)
}

func TestWriteTo(t *testing.T) {
	ChunkSize = 1
	var buf bytes.Buffer

	b, err := NewByteBuf([]byte{0x00, 0x01, 0x02})
	Equal(t, err, nil)
	NotEqual(t, b, nil)

	b.WriteTo(&buf)
	Equal(t, buf.Bytes(), []byte{0x00, 0x01, 0x02})
}
