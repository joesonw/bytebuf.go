package bytebuf

import (
	"bytes"
	"errors"
	"io"
)

// Global parameters, users can modify accordingly
var (
	// ChunkSize how much byte to allocate if there were not enough room
	ChunkSize = 1024
)

// Errors
var (
	ErrOutOfBound = errors.New("index is out of bound")
)

// ByteIterator type used for byte iteration
type ByteIterator func(byte) bool

// EnsureCapacity allocate more bytes to ensure the capacity
func (b *ByteBuf) EnsureCapacity(size int) {
	if size <= b.capacity {
		return
	}
	capacity := ((size-1)/ChunkSize + 1) * ChunkSize
	newBuf := make([]byte, capacity)
	copy(newBuf, b.buf)
	b.buf = nil
	b.buf = newBuf
	b.capacity = capacity
}

// Index index a byte slice inside buffer, Index(p), Index(p, start), Index(p, start, end)
func (b *ByteBuf) Index(p []byte, indexes ...int) int {
	if len(indexes) >= 2 {
		if indexes[1] > indexes[0] && indexes[1] < b.capacity && indexes[0] > 0 {
			return bytes.Index(b.buf[indexes[0]:indexes[1]], p)
		}
		return -1
	} else if len(indexes) >= 1 {
		if indexes[0] > 0 && indexes[0] < b.capacity {
			return bytes.Index(b.buf[indexes[0]:], p)
		}
		return -1
	}
	return bytes.Index(b.buf, p)
}

// Equal if current ByteBuf is equal to another ByteBuf, compared by underlying byte slice
func (b *ByteBuf) Equal(ob *ByteBuf) bool {
	return bytes.Equal(ob.buf, b.buf)
}

// DiscardReadBytes discard bytes that are read, adjust readerIndex/writerIndex accordingly
func (b *ByteBuf) DiscardReadBytes() {
	b.buf = b.buf[b.readerIndex:]
	b.writerIndex -= b.readerIndex
	b.readerIndex = 0
}

// Copy deep copy to create an brand new ByteBuf
func (b *ByteBuf) Copy() *ByteBuf {
	p := make([]byte, len(b.buf))
	copy(p, b.buf)
	return &ByteBuf{
		buf: p,

		capacity: b.capacity,

		readerIndex:  b.readerIndex,
		readerMarker: b.readerMarker,
		writerIndex:  b.writerIndex,
		writerMarker: b.writerMarker,
	}
}

// ForEachByte iterate through readable bytes, ForEachByte(iterator, start), ForEachByte(iterator, start, end)
func (b *ByteBuf) ForEachByte(iterator ByteIterator, indexes ...int) int {
	start := b.readerIndex
	end := b.writerIndex
	if len(indexes) >= 1 {
		start = indexes[0]
	}
	if len(indexes) >= 2 && indexes[1] < end {
		end = indexes[1]
	}

	if start > end {
		return 0
	}

	if end > b.capacity {
		end = b.capacity
	}

	count := 0
	for ; start < end; start++ {
		if !iterator(b.buf[start]) {
			break
		}
		count++
	}

	return count
}

// NewReader create ByteBuf from io.Reader
func NewReader(reader io.Reader) (*ByteBuf, error) {
	b, err := NewByteBuf()
	if err != nil {
		return nil, err
	}
	return b, b.ReadFrom(reader)
}

// String buf to string
func (b *ByteBuf) String() string {
	return string(b.AvailableBytes())
}
