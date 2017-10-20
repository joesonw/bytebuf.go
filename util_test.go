package bytebuf

import (
	"bytes"
	. "gopkg.in/go-playground/assert.v1"
	"testing"
)

func TestUtil(t *testing.T) {
	b, err := NewByteBuf()
	Equal(t, err, nil)
	NotEqual(t, b, nil)

	Equal(t, b.Capacity(), ChunkSize)
	b.EnsureCapacity(ChunkSize * 2)
	Equal(t, b.Capacity(), ChunkSize*2)
	b.EnsureCapacity(ChunkSize)
	Equal(t, b.Capacity(), ChunkSize*2)

	b, err = NewByteBuf([]byte{0x00, 0x01, 0x02})
	Equal(t, err, nil)
	NotEqual(t, b, nil)

	counter := byte(0)
	n := b.ForEachByte(func(p byte) bool {
		Equal(t, p, counter)
		counter++
		return true
	})
	Equal(t, n, 3)
	Equal(t, b.Index([]byte{0x01}), 1)
	Equal(t, b.Index([]byte{0x01}, 1), 0)
	Equal(t, b.Index([]byte{0x01}, 2), -1)
	Equal(t, b.Index([]byte{0x01}, 2, 3), -1)

	Equal(t, b.Equal(b.Copy()), true)
}

func TestUtil2(t *testing.T) {
	ChunkSize = 1
	reader := bytes.NewReader([]byte{0x00, 0x01, 0x02})

	b, err := NewReader(reader)
	Equal(t, err, nil)
	NotEqual(t, b, nil)

	Equal(t, b.buf, []byte{0x00, 0x01, 0x02})
}

func TestString(t *testing.T) {
	ChunkSize = 10

	b, err := NewByteBuf([]byte("test"))
	Equal(t, err, nil)
	NotEqual(t, b, nil)

	Equal(t, b.Capacity(), 10)
	Equal(t, b.String(), "test")

}
