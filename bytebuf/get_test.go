package bytebuf

import (
	. "gopkg.in/go-playground/assert.v1"
	"testing"
)

func TestGetBool(t *testing.T) {
	ChunkSize = 1

	b, err := NewByteBuf([]byte{0x01})
	Equal(t, err, nil)
	NotEqual(t, b, nil)

	val, err := b.GetBool(0)
	Equal(t, val, true)
	Equal(t, err, nil)
	val, err = b.GetBool(1)
	Equal(t, err, ErrOutOfBound)
}

func TestGetByte(t *testing.T) {
	ChunkSize = 1

	b, err := NewByteBuf([]byte{0x08})
	Equal(t, err, nil)
	NotEqual(t, b, nil)

	val, err := b.GetByte(0)
	Equal(t, val, byte(0x08))
	Equal(t, err, nil)

	val, err = b.GetByte(1)
	Equal(t, err, ErrOutOfBound)
}

func TestGetBytes(t *testing.T) {
	ChunkSize = 1

	b, err := NewByteBuf([]byte{0x00, 0x01, 0x02, 0x03})
	Equal(t, err, nil)
	NotEqual(t, b, nil)

	p := make([]byte, 4)
	n, err := b.GetBytes(p, 0)
	Equal(t, n, 4)
	Equal(t, err, nil)
	Equal(t, p, []byte{0x00, 0x01, 0x02, 0x03})

	p = make([]byte, 4)
	n, err = b.GetBytes(p, 1, 5)
	Equal(t, n, 3)
	Equal(t, err, nil)
	Equal(t, p, []byte{0x01, 0x02, 0x03, 0x00})

	p = make([]byte, 4)
	n, err = b.GetBytes(p, 1, 2)
	Equal(t, n, 1)
	Equal(t, err, nil)
	Equal(t, p, []byte{0x01, 0x00, 0x00, 0x00})

	p = make([]byte, 4)
	n, err = b.GetBytes(p, 3, 2)
	Equal(t, err, ErrOutOfBound)
}

func TestGetUInt8(t *testing.T) {
	ChunkSize = 1

	b, err := NewByteBuf([]byte{0x08})
	Equal(t, err, nil)
	NotEqual(t, b, nil)

	v, err := b.GetUInt8(0)
	Equal(t, v, uint8(0x08))
	Equal(t, err, nil)

	v, err = b.GetUInt8(1)
	Equal(t, err, ErrOutOfBound)
}

func TestGetInt8(t *testing.T) {
	ChunkSize = 1

	b, err := NewByteBuf([]byte{0xfe})
	Equal(t, err, nil)
	NotEqual(t, b, nil)

	v, err := b.GetInt8(0)
	Equal(t, v, int8(-2))
	Equal(t, err, nil)

	v, err = b.GetInt8(1)
	Equal(t, err, ErrOutOfBound)
}

func TestGetUInt16LE(t *testing.T) {
	ChunkSize = 1

	b, err := NewByteBuf([]byte{0x08, 0x01})
	Equal(t, err, nil)
	NotEqual(t, b, nil)

	v, err := b.GetUInt16LE(0)
	Equal(t, err, nil)
	Equal(t, v, uint16(264))

	v, err = b.GetUInt16LE(1)
	Equal(t, err, ErrOutOfBound)
}

func TestGetUInt16BE(t *testing.T) {
	ChunkSize = 1

	b, err := NewByteBuf([]byte{0x01, 0x08})
	Equal(t, err, nil)
	NotEqual(t, b, nil)

	v, err := b.GetUInt16BE(0)
	Equal(t, err, nil)
	Equal(t, v, uint16(264))

	v, err = b.GetUInt16BE(1)
	Equal(t, err, ErrOutOfBound)
}

func TestGetInt16LE(t *testing.T) {
	ChunkSize = 1

	b, err := NewByteBuf([]byte{0xfe, 0xff})
	Equal(t, err, nil)
	NotEqual(t, b, nil)

	v, err := b.GetInt16LE(0)
	Equal(t, err, nil)
	Equal(t, v, int16(-2))

	v, err = b.GetInt16LE(1)
	Equal(t, err, ErrOutOfBound)
}

func TestGetInt16BE(t *testing.T) {
	ChunkSize = 1

	b, err := NewByteBuf([]byte{0xff, 0xfe})
	Equal(t, err, nil)
	NotEqual(t, b, nil)

	v, err := b.GetInt16BE(0)
	Equal(t, err, nil)
	Equal(t, v, int16(-2))

	v, err = b.GetInt16BE(1)
	Equal(t, err, ErrOutOfBound)
}

func TestGetUInt32LE(t *testing.T) {
	ChunkSize = 1

	b, err := NewByteBuf([]byte{0x01, 0x02, 0x03, 0x04})
	Equal(t, err, nil)
	NotEqual(t, b, nil)

	v, err := b.GetUInt32LE(0)
	Equal(t, v, uint32(67305985))
	Equal(t, err, nil)

	v, err = b.GetUInt32LE(1)
	Equal(t, err, ErrOutOfBound)
}

func TestGetUInt32BE(t *testing.T) {
	ChunkSize = 1

	b, err := NewByteBuf([]byte{0x04, 0x03, 0x02, 0x01})
	Equal(t, err, nil)
	NotEqual(t, b, nil)

	v, err := b.GetUInt32BE(0)
	Equal(t, v, uint32(67305985))
	Equal(t, err, nil)

	v, err = b.GetUInt32BE(1)
	Equal(t, err, ErrOutOfBound)
}

func TestGetInt32LE(t *testing.T) {
	ChunkSize = 1

	b, err := NewByteBuf([]byte{0xfe, 0xff, 0xff, 0xff})
	Equal(t, err, nil)
	NotEqual(t, b, nil)

	v, err := b.GetInt32LE(0)
	Equal(t, err, nil)
	Equal(t, v, int32(-2))

	v, err = b.GetInt32LE(1)
	Equal(t, err, ErrOutOfBound)
}

func TestGetInt32BE(t *testing.T) {
	ChunkSize = 1

	b, err := NewByteBuf([]byte{0xff, 0xff, 0xff, 0xfe})
	Equal(t, err, nil)
	NotEqual(t, b, nil)

	v, err := b.GetInt32BE(0)
	Equal(t, v, int32(-2))
	Equal(t, err, nil)

	v, err = b.GetInt32BE(1)
	Equal(t, err, ErrOutOfBound)
}

func TestGetUInt64LE(t *testing.T) {
	ChunkSize = 1

	b, err := NewByteBuf([]byte{0x01, 0x02, 0x03, 0x04, 0x00, 0x00, 0x00, 0x00})
	Equal(t, err, nil)
	NotEqual(t, b, nil)

	v, err := b.GetUInt64LE(0)
	Equal(t, err, nil)
	Equal(t, v, uint64(67305985))

	v, err = b.GetUInt64LE(1)
	Equal(t, err, ErrOutOfBound)
}

func TestGetUInt64BE(t *testing.T) {
	ChunkSize = 1

	b, err := NewByteBuf([]byte{0x00, 0x00, 0x00, 0x00, 0x04, 0x03, 0x02, 0x01})
	Equal(t, err, nil)
	NotEqual(t, b, nil)

	v, err := b.GetUInt64BE(0)
	Equal(t, err, nil)
	Equal(t, v, uint64(67305985))

	v, err = b.GetUInt64BE(1)
	Equal(t, err, ErrOutOfBound)
}

func TestGetInt64LE(t *testing.T) {
	ChunkSize = 1

	b, err := NewByteBuf([]byte{0xfe, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff})
	Equal(t, err, nil)
	NotEqual(t, b, nil)

	v, err := b.GetInt64LE(0)
	Equal(t, err, nil)
	Equal(t, v, int64(-2))

	v, err = b.GetInt64LE(1)
	Equal(t, err, ErrOutOfBound)
}

func TestGetInt64BE(t *testing.T) {
	ChunkSize = 1

	b, err := NewByteBuf([]byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xfe})
	Equal(t, err, nil)
	NotEqual(t, b, nil)

	v, err := b.GetInt64BE(0)
	Equal(t, err, nil)
	Equal(t, v, int64(-2))

	v, err = b.GetInt64BE(1)
	Equal(t, err, ErrOutOfBound)
}

func TestGetIntFloat32LE(t *testing.T) {
	ChunkSize = 1

	b, err := NewByteBuf([]byte{0xdb, 0x0f, 0x49, 0x40})
	Equal(t, err, nil)
	NotEqual(t, b, nil)

	v, err := b.GetFloat32LE(0)
	Equal(t, err, nil)
	Equal(t, v, float32(3.141592653589793))

	v, err = b.GetFloat32LE(1)
	Equal(t, err, ErrOutOfBound)
}

func TestGetIntFloat32BE(t *testing.T) {
	ChunkSize = 1

	b, err := NewByteBuf([]byte{0x40, 0x49, 0x0f, 0xdb})
	Equal(t, err, nil)
	NotEqual(t, b, nil)

	v, err := b.GetFloat32BE(0)
	Equal(t, err, nil)
	Equal(t, v, float32(3.141592653589793))

	v, err = b.GetFloat32BE(1)
	Equal(t, err, ErrOutOfBound)
}

func TestGetIntFloat64LE(t *testing.T) {
	ChunkSize = 1

	b, err := NewByteBuf([]byte{24, 45, 68, 84, 251, 33, 9, 64})
	Equal(t, err, nil)
	NotEqual(t, b, nil)

	v, err := b.GetFloat64LE(0)
	Equal(t, err, nil)
	Equal(t, v, float64(3.141592653589793))

	v, err = b.GetFloat64LE(1)
	Equal(t, err, ErrOutOfBound)
}

func TestGetIntFloat64BE(t *testing.T) {
	ChunkSize = 1

	b, err := NewByteBuf([]byte{64, 9, 33, 251, 84, 68, 45, 24})
	Equal(t, err, nil)
	NotEqual(t, b, nil)

	v, err := b.GetFloat64BE(0)
	Equal(t, err, nil)
	Equal(t, v, float64(3.141592653589793))

	v, err = b.GetFloat64BE(1)
	Equal(t, err, ErrOutOfBound)
}
