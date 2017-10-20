package bytebuf

import (
	. "gopkg.in/go-playground/assert.v1"
	"testing"
)

func TestPutBool(t *testing.T) {
	ChunkSize = 1

	b, err := NewByteBuf(make([]byte, 1))
	Equal(t, err, nil)
	NotEqual(t, b, nil)

	err = b.PutBool(0, true)
	Equal(t, b.buf, []byte{0x01})
	Equal(t, err, nil)

	err = b.PutBool(1, true)
	Equal(t, err, ErrOutOfBound)

	err = b.PutBool(0, false)
	Equal(t, b.buf, []byte{0x00})
	Equal(t, err, nil)
}

func TestPutByte(t *testing.T) {
	ChunkSize = 1

	b, err := NewByteBuf([]byte{0x08})
	Equal(t, err, nil)
	NotEqual(t, b, nil)

	err = b.PutByte(0, 0x08)
	Equal(t, b.buf, []byte{0x08})
	Equal(t, err, nil)

	err = b.PutByte(1, 0x08)
	Equal(t, err, ErrOutOfBound)
}

func TestPutBytes(t *testing.T) {
	ChunkSize = 1

	b, err := NewByteBuf(make([]byte, 4))
	Equal(t, err, nil)
	NotEqual(t, b, nil)

	err = b.PutBytes(0, []byte{0x00, 0x01, 0x02, 0x03})
	Equal(t, err, nil)
	Equal(t, b.buf, []byte{0x00, 0x01, 0x02, 0x03})

	err = b.PutBytes(1, []byte{0x00, 0x01, 0x02})
	Equal(t, err, nil)
	Equal(t, b.buf, []byte{0x00, 0x00, 0x01, 0x02})

	err = b.PutBytes(1, []byte{0x00, 0x01, 0x02, 0x03})
	Equal(t, err, ErrOutOfBound)
}

func TestPutUInt8(t *testing.T) {
	ChunkSize = 1

	b, err := NewByteBuf(make([]byte, 1))
	Equal(t, err, nil)
	NotEqual(t, b, nil)

	err = b.PutUInt8(0, 8)
	Equal(t, err, nil)
	Equal(t, b.buf, []byte{0x08})

	err = b.PutUInt8(1, 8)
	Equal(t, err, ErrOutOfBound)
}

func TestPutInt8(t *testing.T) {
	ChunkSize = 1

	b, err := NewByteBuf(make([]byte, 1))
	Equal(t, err, nil)
	NotEqual(t, b, nil)

	err = b.PutInt8(0, -2)
	Equal(t, err, nil)
	Equal(t, b.buf, []byte{0xfe})

	err = b.PutInt8(1, -2)
	Equal(t, err, ErrOutOfBound)
}

func TestPutUInt16LE(t *testing.T) {
	ChunkSize = 1

	b, err := NewByteBuf(make([]byte, 2))
	Equal(t, err, nil)
	NotEqual(t, b, nil)

	err = b.PutUInt16LE(0, 264)
	Equal(t, err, nil)
	Equal(t, b.buf, []byte{0x08, 0x01})

	err = b.PutUInt16LE(1, 264)
	Equal(t, err, ErrOutOfBound)
}

func TestPutUInt16BE(t *testing.T) {
	ChunkSize = 1

	b, err := NewByteBuf(make([]byte, 2))
	Equal(t, err, nil)
	NotEqual(t, b, nil)

	err = b.PutUInt16BE(0, 264)
	Equal(t, err, nil)
	Equal(t, b.buf, []byte{0x01, 0x08})

	err = b.PutUInt16BE(1, 264)
	Equal(t, err, ErrOutOfBound)
}

func TestPutInt16LE(t *testing.T) {
	ChunkSize = 1

	b, err := NewByteBuf(make([]byte, 2))
	Equal(t, err, nil)
	NotEqual(t, b, nil)

	err = b.PutInt16LE(0, -2)
	Equal(t, err, nil)
	Equal(t, b.buf, []byte{0xfe, 0xff})

	err = b.PutInt16LE(1, -2)
	Equal(t, err, ErrOutOfBound)
}

func TestPutInt16BE(t *testing.T) {
	ChunkSize = 1

	b, err := NewByteBuf(make([]byte, 2))
	Equal(t, err, nil)
	NotEqual(t, b, nil)

	err = b.PutInt16BE(0, -2)
	Equal(t, err, nil)
	Equal(t, b.buf, []byte{0xff, 0xfe})

	err = b.PutInt16BE(1, -2)
	Equal(t, err, ErrOutOfBound)
}

func TestPutUInt32LE(t *testing.T) {
	ChunkSize = 1

	b, err := NewByteBuf(make([]byte, 4))
	Equal(t, err, nil)
	NotEqual(t, b, nil)

	err = b.PutUInt32LE(0, 67305985)
	Equal(t, err, nil)
	Equal(t, b.buf, []byte{0x01, 0x02, 0x03, 0x04})

	err = b.PutUInt32LE(1, 0)
	Equal(t, err, ErrOutOfBound)
}

func TestPutUInt32BE(t *testing.T) {
	ChunkSize = 1

	b, err := NewByteBuf(make([]byte, 4))
	Equal(t, err, nil)
	NotEqual(t, b, nil)

	err = b.PutUInt32BE(0, 67305985)
	Equal(t, err, nil)
	Equal(t, b.buf, []byte{0x04, 0x03, 0x02, 0x01})

	err = b.PutUInt32BE(1, 0)
	Equal(t, err, ErrOutOfBound)
}

func TestPutInt32LE(t *testing.T) {
	ChunkSize = 1

	b, err := NewByteBuf(make([]byte, 4))
	Equal(t, err, nil)
	NotEqual(t, b, nil)

	err = b.PutInt32LE(0, -2)
	Equal(t, err, nil)
	Equal(t, b.buf, []byte{0xfe, 0xff, 0xff, 0xff})

	err = b.PutInt32LE(1, 0)
	Equal(t, err, ErrOutOfBound)
}

func TestPutInt32BE(t *testing.T) {
	ChunkSize = 1

	b, err := NewByteBuf(make([]byte, 4))
	Equal(t, err, nil)
	NotEqual(t, b, nil)

	err = b.PutInt32BE(0, -2)
	Equal(t, err, nil)
	Equal(t, b.buf, []byte{0xff, 0xff, 0xff, 0xfe})

	err = b.PutInt32BE(1, 0)
	Equal(t, err, ErrOutOfBound)
}

func TestPutUInt64LE(t *testing.T) {
	ChunkSize = 1

	b, err := NewByteBuf(make([]byte, 8))
	Equal(t, err, nil)
	NotEqual(t, b, nil)

	err = b.PutUInt64LE(0, 67305985)
	Equal(t, err, nil)
	Equal(t, b.buf, []byte{0x01, 0x02, 0x03, 0x04, 0x00, 0x00, 0x00, 0x00})

	err = b.PutUInt64LE(1, 0)
	Equal(t, err, ErrOutOfBound)
}

func TestPutUInt64BE(t *testing.T) {
	ChunkSize = 1

	b, err := NewByteBuf(make([]byte, 8))
	Equal(t, err, nil)
	NotEqual(t, b, nil)

	err = b.PutUInt64BE(0, 67305985)
	Equal(t, err, nil)
	Equal(t, b.buf, []byte{0x00, 0x00, 0x00, 0x00, 0x04, 0x03, 0x02, 0x01})

	err = b.PutUInt64BE(1, 0)
	Equal(t, err, ErrOutOfBound)
}

func TestPutInt64LE(t *testing.T) {
	ChunkSize = 1

	b, err := NewByteBuf(make([]byte, 8))
	Equal(t, err, nil)
	NotEqual(t, b, nil)

	err = b.PutInt64LE(0, -2)
	Equal(t, err, nil)
	Equal(t, b.buf, []byte{0xfe, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff})

	err = b.PutInt64LE(1, 0)
	Equal(t, err, ErrOutOfBound)
}

func TestPutInt64BE(t *testing.T) {
	ChunkSize = 1

	b, err := NewByteBuf(make([]byte, 8))
	Equal(t, err, nil)
	NotEqual(t, b, nil)

	err = b.PutInt64BE(0, -2)
	Equal(t, err, nil)
	Equal(t, b.buf, []byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xfe})

	err = b.PutInt64BE(1, 0)
	Equal(t, err, ErrOutOfBound)
}

func TestPutIntFloat32LE(t *testing.T) {
	ChunkSize = 1

	b, err := NewByteBuf(make([]byte, 4))
	Equal(t, err, nil)
	NotEqual(t, b, nil)

	err = b.PutFloat32LE(0, 3.141592653589793)
	Equal(t, err, nil)
	Equal(t, b.buf, []byte{0xdb, 0x0f, 0x49, 0x40})

	err = b.PutFloat32LE(1, 3.141592653589793)
	Equal(t, err, ErrOutOfBound)
}

func TestPutIntFloat32BE(t *testing.T) {
	ChunkSize = 1

	b, err := NewByteBuf(make([]byte, 4))
	Equal(t, err, nil)
	NotEqual(t, b, nil)

	err = b.PutFloat32BE(0, 3.141592653589793)
	Equal(t, err, nil)
	Equal(t, b.buf, []byte{0x40, 0x49, 0x0f, 0xdb})

	err = b.PutFloat32BE(1, 3.141592653589793)
	Equal(t, err, ErrOutOfBound)
}

func TestPutIntFloat64LE(t *testing.T) {
	ChunkSize = 1

	b, err := NewByteBuf(make([]byte, 8))
	Equal(t, err, nil)
	NotEqual(t, b, nil)

	err = b.PutFloat64LE(0, 3.141592653589793)
	Equal(t, err, nil)
	Equal(t, b.buf, []byte{24, 45, 68, 84, 251, 33, 9, 64})

	err = b.PutFloat64LE(1, 3.141592653589793)
	Equal(t, err, ErrOutOfBound)
}

func TestPutIntFloat64BE(t *testing.T) {
	ChunkSize = 1

	b, err := NewByteBuf(make([]byte, 8))
	Equal(t, err, nil)
	NotEqual(t, b, nil)

	err = b.PutFloat64BE(0, 3.141592653589793)
	Equal(t, err, nil)
	Equal(t, b.buf, []byte{64, 9, 33, 251, 84, 68, 45, 24})

	err = b.PutFloat64BE(1, 3.141592653589793)
	Equal(t, err, ErrOutOfBound)
}
