package bytebuf

import (
	"io"
)

// ByteBuf ByteBuf itself
type ByteBuf struct {
	buf []byte

	capacity int
	refCount int

	readerMarker int
	writerMarker int
	readerIndex  int
	writerIndex  int
}

// NewByteBuf create new ByteBuf, pass any []byte in as initial buffer, it will be auto capped
func NewByteBuf(bufs ...[]byte) (*ByteBuf, error) {
	b := &ByteBuf{
		buf: make([]byte, ChunkSize),

		capacity:     ChunkSize,
		readerMarker: -1,
		writerMarker: -1,
	}

	for _, buf := range bufs {
		n, err := b.Write(buf)
		if err != nil {
			return nil, err
		}
		if n != len(buf) {
			return nil, io.ErrShortWrite
		}
	}

	return b, nil
}

// ReaderIndex access current reader index
func (b *ByteBuf) ReaderIndex() int {
	return b.readerIndex
}

// WriterIndex access current writer index
func (b *ByteBuf) WriterIndex() int {
	return b.writerIndex
}

// AvailableBytes get available bytes([:size])
func (b *ByteBuf) AvailableBytes() []byte {
	p := make([]byte, b.writerIndex)
	copy(p, b.buf)
	return p
}

// Bytes get all the bytes in buffer ([:capacity])
func (b *ByteBuf) Bytes() []byte {
	p := make([]byte, b.capacity)
	copy(p, b.buf)
	return p
}

// UnsafeBytes DANGER! access underlying buffer directly
func (b *ByteBuf) UnsafeBytes() []byte {
	return b.buf
}

// Size get current written buffer size
func (b *ByteBuf) Size() int {
	return b.writerIndex
}

// Capacity get buffer capacity (how much it could read/write)
func (b *ByteBuf) Capacity() int {
	return b.capacity
}

// MarkReaderIndex mark current reader index for reset in the future
func (b *ByteBuf) MarkReaderIndex() {
	b.readerMarker = b.readerIndex
}

// ResetReaderIndex reset reader index to marked index
func (b *ByteBuf) ResetReaderIndex() {
	if b.readerMarker != -1 {
		b.readerIndex = b.readerMarker
		b.readerMarker = -1
	}
}

// MarkWriterIndex mark current writer index for reset in the future
func (b *ByteBuf) MarkWriterIndex() {
	b.writerMarker = b.writerIndex
}

// ResetWriterIndex reset writer index to marked index
func (b *ByteBuf) ResetWriterIndex() {
	if b.writerMarker != -1 {
		b.writerIndex = b.writerMarker
		b.writerMarker = -1
	}
}

// Flush flush all bytes and reset indexes
func (b *ByteBuf) Flush() {
	b.buf = nil
	b.buf = make([]byte, ChunkSize)

	b.readerIndex = 0
	b.readerMarker = -1
	b.writerIndex = 0
	b.writerMarker = 0

	b.capacity = ChunkSize
}

// IsReadable if buf is readable (b.writerIndex > b.readerIndex)
func (b *ByteBuf) IsReadable() bool {
	return b.writerIndex > b.readerIndex
}

// ReadableBytes how much bytes are there to be read
func (b *ByteBuf) ReadableBytes() int {
	return b.writerIndex - b.readerIndex
}

// Ref increment reference counter
func (b *ByteBuf) Ref() {
	b.refCount++
}

// Release decrement reference counter, underlying buf will be cleared once reference count is 0
func (b *ByteBuf) Release() {
	b.refCount--
	if b.refCount == 0 {
		b.ForceRelease()
	}
}

// ForceRelease force release bufs
func (b *ByteBuf) ForceRelease() {
	b.refCount = 0
	b.buf = nil
}
