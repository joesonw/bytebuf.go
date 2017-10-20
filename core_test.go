package bytebuf

import (
	. "gopkg.in/go-playground/assert.v1"
	"testing"
)

func TestCore1(t *testing.T) {
	ChunkSize = 1
	b, err := NewByteBuf([]byte{0x00, 0x01}, []byte{0x02, 0x03, 0x04})
	Equal(t, err, nil)
	NotEqual(t, b, nil)

	Equal(t, b.Capacity(), 5)
	Equal(t, b.Size(), 5)

	Equal(t, b.AvailableBytes(), []byte{0x00, 0x01, 0x02, 0x03, 0x04})
	Equal(t, b.Bytes(), []byte{0x00, 0x01, 0x02, 0x03, 0x04})

	b.Flush()
	Equal(t, b.Capacity(), 1)
	Equal(t, b.Size(), 0)
}

func TestCore2(t *testing.T) {
	ChunkSize = 10
	b, err := NewByteBuf([]byte{0x00, 0x01}, []byte{0x02, 0x03, 0x04})
	Equal(t, err, nil)
	NotEqual(t, b, nil)
	Equal(t, b.ReadableBytes(), 5)
	Equal(t, b.Capacity(), 10)
	Equal(t, b.Size(), 5)
	Equal(t, b.AvailableBytes(), []byte{0x00, 0x01, 0x02, 0x03, 0x04})
	Equal(t, b.Bytes(), []byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x00, 0x00, 0x00, 0x00, 0x00})
	Equal(t, b.UnsafeBytes(), []byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x00, 0x00, 0x00, 0x00, 0x00})

	Equal(t, b.readerMarker, -1)
	b.MarkReaderIndex()
	b.readerIndex = 1
	Equal(t, b.readerIndex, 1)
	Equal(t, b.readerMarker, 0)
	b.ResetReaderIndex()
	Equal(t, b.readerIndex, 0)
	Equal(t, b.readerMarker, -1)

	Equal(t, b.writerMarker, -1)
	b.writerIndex = 0
	b.MarkWriterIndex()
	b.writerIndex = 1
	Equal(t, b.writerIndex, 1)
	Equal(t, b.writerMarker, 0)
	b.ResetWriterIndex()
	Equal(t, b.writerIndex, 0)
	Equal(t, b.writerMarker, -1)

	Equal(t, b.IsReadable(), false)
	b.writerIndex = 1
	Equal(t, b.IsReadable(), true)

}
