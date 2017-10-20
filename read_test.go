package bytebuf

import (
	"bytes"
	. "gopkg.in/go-playground/assert.v1"
	"io"
	"testing"
)

func TestRead(t *testing.T) {
	b, err := NewByteBuf([]byte{0x00, 0x01}, []byte{0x02, 0x03, 0x04})
	Equal(t, err, nil)
	NotEqual(t, b, nil)

	p := make([]byte, 5)
	n, err := b.Read(p)

	Equal(t, err, nil)
	Equal(t, n, 5)
	Equal(t, p, []byte{0x00, 0x01, 0x02, 0x03, 0x04})

	_, err = b.Read(p)
	Equal(t, err, io.EOF)
}

func TestReadBool(t *testing.T) {
	b, err := NewByteBuf([]byte{0x01})
	Equal(t, err, nil)
	NotEqual(t, b, nil)
	b.writerIndex = 1

	val, err := b.ReadBool()
	Equal(t, err, nil)
	Equal(t, val, true)
	Equal(t, b.ReaderIndex(), 1)

	_, err = b.ReadBool()
	Equal(t, err, ErrOutOfBound)
}

func TestReadByte(t *testing.T) {
	b, err := NewByteBuf([]byte{0x08})
	Equal(t, err, nil)
	NotEqual(t, b, nil)
	b.writerIndex = 1

	val, err := b.ReadByte()
	Equal(t, err, nil)
	Equal(t, val, byte(0x08))
	Equal(t, b.ReaderIndex(), 1)

	_, err = b.ReadByte()
	Equal(t, err, ErrOutOfBound)
}

func TestReadBytes(t *testing.T) {
	b, err := NewByteBuf([]byte{0x00, 0x01})
	Equal(t, err, nil)
	NotEqual(t, b, nil)
	b.writerIndex = 2

	p := make([]byte, 2)
	n, err := b.ReadBytes(p)
	Equal(t, err, nil)
	Equal(t, p, []byte{0x00, 0x01})
	Equal(t, n, 2)
	Equal(t, b.ReaderIndex(), 2)

	n, _ = b.ReadBytes(p)
	Equal(t, n, 0)
}

func TestReadUInt8(t *testing.T) {
	b, err := NewByteBuf([]byte{0x08})
	Equal(t, err, nil)
	NotEqual(t, b, nil)
	b.writerIndex = 1

	v, err := b.ReadUInt8()
	Equal(t, v, uint8(0x08))
	Equal(t, err, nil)
	Equal(t, b.ReaderIndex(), 1)

	v, err = b.ReadUInt8()
	Equal(t, err, ErrOutOfBound)
}

func TestReadInt8(t *testing.T) {
	b, err := NewByteBuf([]byte{0xfe})
	Equal(t, err, nil)
	NotEqual(t, b, nil)
	b.writerIndex = 1

	v, err := b.ReadInt8()
	Equal(t, v, int8(-2))
	Equal(t, err, nil)
	Equal(t, b.ReaderIndex(), 1)

	v, err = b.ReadInt8()
	Equal(t, err, ErrOutOfBound)
}

func TestReadUInt16LE(t *testing.T) {
	b, err := NewByteBuf([]byte{0x08, 0x01})
	Equal(t, err, nil)
	NotEqual(t, b, nil)
	b.writerIndex = 2

	v, err := b.ReadUInt16LE()
	Equal(t, err, nil)
	Equal(t, v, uint16(264))
	Equal(t, b.ReaderIndex(), 2)

	v, err = b.ReadUInt16LE()
	Equal(t, err, ErrOutOfBound)
}

func TestReadUInt16BE(t *testing.T) {
	b, err := NewByteBuf([]byte{0x01, 0x08})
	Equal(t, err, nil)
	NotEqual(t, b, nil)
	b.writerIndex = 2

	v, err := b.ReadUInt16BE()
	Equal(t, err, nil)
	Equal(t, v, uint16(264))
	Equal(t, b.ReaderIndex(), 2)

	v, err = b.ReadUInt16BE()
	Equal(t, err, ErrOutOfBound)
}

func TestReadInt16LE(t *testing.T) {
	b, err := NewByteBuf([]byte{0xfe, 0xff})
	Equal(t, err, nil)
	NotEqual(t, b, nil)
	b.writerIndex = 2

	v, err := b.ReadInt16LE()
	Equal(t, err, nil)
	Equal(t, v, int16(-2))
	Equal(t, b.ReaderIndex(), 2)

	v, err = b.ReadInt16LE()
	Equal(t, err, ErrOutOfBound)
}

func TestReadInt16BE(t *testing.T) {
	b, err := NewByteBuf([]byte{0xff, 0xfe})
	Equal(t, err, nil)
	NotEqual(t, b, nil)
	b.writerIndex = 2

	v, err := b.ReadInt16BE()
	Equal(t, err, nil)
	Equal(t, v, int16(-2))
	Equal(t, b.ReaderIndex(), 2)

	v, err = b.ReadInt16BE()
	Equal(t, err, ErrOutOfBound)
}

func TestReadUInt32LE(t *testing.T) {
	b, err := NewByteBuf([]byte{0x01, 0x02, 0x03, 0x04})
	Equal(t, err, nil)
	NotEqual(t, b, nil)
	b.writerIndex = 4

	v, err := b.ReadUInt32LE()
	Equal(t, v, uint32(67305985))
	Equal(t, err, nil)
	Equal(t, b.ReaderIndex(), 4)

	v, err = b.ReadUInt32LE()
	Equal(t, err, ErrOutOfBound)
}

func TestReadUInt32BE(t *testing.T) {
	b, err := NewByteBuf([]byte{0x04, 0x03, 0x02, 0x01})
	Equal(t, err, nil)
	NotEqual(t, b, nil)
	b.writerIndex = 4

	v, err := b.ReadUInt32BE()
	Equal(t, v, uint32(67305985))
	Equal(t, err, nil)
	Equal(t, b.ReaderIndex(), 4)

	v, err = b.ReadUInt32BE()
	Equal(t, err, ErrOutOfBound)
}

func TestReadInt32LE(t *testing.T) {
	b, err := NewByteBuf([]byte{0xfe, 0xff, 0xff, 0xff})
	Equal(t, err, nil)
	NotEqual(t, b, nil)
	b.writerIndex = 4

	v, err := b.ReadInt32LE()
	Equal(t, err, nil)
	Equal(t, v, int32(-2))
	Equal(t, b.ReaderIndex(), 4)

	v, err = b.ReadInt32LE()
	Equal(t, err, ErrOutOfBound)
}

func TestReadInt32BE(t *testing.T) {
	b, err := NewByteBuf([]byte{0xff, 0xff, 0xff, 0xfe})
	Equal(t, err, nil)
	NotEqual(t, b, nil)
	b.writerIndex = 4

	v, err := b.ReadInt32BE()
	Equal(t, v, int32(-2))
	Equal(t, err, nil)
	Equal(t, b.ReaderIndex(), 4)

	v, err = b.ReadInt32BE()
	Equal(t, err, ErrOutOfBound)
}

func TestReadUInt64LE(t *testing.T) {
	b, err := NewByteBuf([]byte{0x01, 0x02, 0x03, 0x04, 0x00, 0x00, 0x00, 0x00})
	Equal(t, err, nil)
	NotEqual(t, b, nil)
	b.writerIndex = 8

	v, err := b.ReadUInt64LE()
	Equal(t, err, nil)
	Equal(t, v, uint64(67305985))
	Equal(t, b.ReaderIndex(), 8)

	v, err = b.ReadUInt64LE()
	Equal(t, err, ErrOutOfBound)
}

func TestReadUInt64BE(t *testing.T) {
	b, err := NewByteBuf([]byte{0x00, 0x00, 0x00, 0x00, 0x04, 0x03, 0x02, 0x01})
	Equal(t, err, nil)
	NotEqual(t, b, nil)
	b.writerIndex = 8

	v, err := b.ReadUInt64BE()
	Equal(t, err, nil)
	Equal(t, v, uint64(67305985))
	Equal(t, b.ReaderIndex(), 8)

	v, err = b.ReadUInt64BE()
	Equal(t, err, ErrOutOfBound)
}

func TestReadInt64LE(t *testing.T) {
	b, err := NewByteBuf([]byte{0xfe, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff})
	Equal(t, err, nil)
	NotEqual(t, b, nil)
	b.writerIndex = 8

	v, err := b.ReadInt64LE()
	Equal(t, err, nil)
	Equal(t, v, int64(-2))
	Equal(t, b.ReaderIndex(), 8)

	v, err = b.ReadInt64LE()
	Equal(t, err, ErrOutOfBound)
}

func TestReadInt64BE(t *testing.T) {
	b, err := NewByteBuf([]byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xfe})
	Equal(t, err, nil)
	NotEqual(t, b, nil)
	b.writerIndex = 8

	v, err := b.ReadInt64BE()
	Equal(t, err, nil)
	Equal(t, v, int64(-2))
	Equal(t, b.ReaderIndex(), 8)

	v, err = b.ReadInt64BE()
	Equal(t, err, ErrOutOfBound)
}

func TestReadIntFloat32LE(t *testing.T) {
	b, err := NewByteBuf([]byte{0xdb, 0x0f, 0x49, 0x40})
	Equal(t, err, nil)
	NotEqual(t, b, nil)
	b.writerIndex = 4

	v, err := b.ReadFloat32LE()
	Equal(t, err, nil)
	Equal(t, v, float32(3.141592653589793))
	Equal(t, b.ReaderIndex(), 4)

	v, err = b.ReadFloat32LE()
	Equal(t, err, ErrOutOfBound)
}

func TestReadIntFloat32BE(t *testing.T) {
	b, err := NewByteBuf([]byte{0x40, 0x49, 0x0f, 0xdb})
	Equal(t, err, nil)
	NotEqual(t, b, nil)
	b.writerIndex = 4

	v, err := b.ReadFloat32BE()
	Equal(t, err, nil)
	Equal(t, v, float32(3.141592653589793))
	Equal(t, b.ReaderIndex(), 4)

	v, err = b.ReadFloat32BE()
	Equal(t, err, ErrOutOfBound)
}

func TestReadIntFloat64LE(t *testing.T) {
	b, err := NewByteBuf([]byte{24, 45, 68, 84, 251, 33, 9, 64})
	Equal(t, err, nil)
	NotEqual(t, b, nil)
	b.writerIndex = 8

	v, err := b.ReadFloat64LE()
	Equal(t, err, nil)
	Equal(t, v, float64(3.141592653589793))
	Equal(t, b.ReaderIndex(), 8)

	v, err = b.ReadFloat64LE()
	Equal(t, err, ErrOutOfBound)
}

func TestReadIntFloat64BE(t *testing.T) {
	b, err := NewByteBuf([]byte{64, 9, 33, 251, 84, 68, 45, 24})
	Equal(t, err, nil)
	NotEqual(t, b, nil)
	b.writerIndex = 8

	v, err := b.ReadFloat64BE()
	Equal(t, err, nil)
	Equal(t, v, float64(3.141592653589793))
	Equal(t, b.ReaderIndex(), 8)

	v, err = b.ReadFloat64BE()
	Equal(t, err, ErrOutOfBound)
}

func TestReadFrom(t *testing.T) {
	ChunkSize = 1
	reader := bytes.NewReader([]byte{0x00, 0x01, 0x02})

	b, err := NewByteBuf()
	Equal(t, err, nil)
	NotEqual(t, b, nil)

	b.ReadFrom(reader)
	Equal(t, b.buf, []byte{0x00, 0x01, 0x02})
}
