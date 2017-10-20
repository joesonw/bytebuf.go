package bytebuf

import (
	"io"
)

// Read implements io.Reader
func (b *ByteBuf) Read(p []byte) (int, error) {
	if b.readerIndex >= b.writerIndex {
		return 0, io.EOF
	}

	n := len(p)
	if len(p)+b.readerIndex > b.writerIndex {
		n = b.writerIndex - b.readerIndex
	}

	copy(p, b.buf[b.readerIndex:b.readerIndex+n])
	b.readerIndex += n

	return n, nil
}

// ReadBool read a bool
func (b *ByteBuf) ReadBool() (bool, error) {
	if b.readerIndex > b.writerIndex-1 {
		return false, ErrOutOfBound
	}

	val, _ := b.GetBool(b.readerIndex)
	b.readerIndex++
	return val, nil
}

// ReadByte read a byte
func (b *ByteBuf) ReadByte() (byte, error) {
	if b.readerIndex > b.writerIndex-1 {
		return 0, ErrOutOfBound
	}
	val, _ := b.GetByte(b.readerIndex)
	b.readerIndex++
	return val, nil
}

// ReadBytes read bytes
func (b *ByteBuf) ReadBytes(p []byte) (int, error) {
	n, _ := b.GetBytes(p, b.readerIndex, b.writerIndex)
	b.readerIndex += n
	return n, nil
}

// ReadByteBuf read bytes with given size into a new ByteBuf
func (b *ByteBuf) ReadByteBuf(size int) (*ByteBuf, error) {
	p := make([]byte, size)
	n, err := b.ReadBytes(p)
	if err != nil {
		return nil, err
	}
	p = p[:n]
	return NewByteBuf(p)
}

// ReadUInt8 read a uint8
func (b *ByteBuf) ReadUInt8() (uint8, error) {
	if b.readerIndex > b.writerIndex-1 {
		return 0, ErrOutOfBound
	}
	val, _ := b.GetUInt8(b.readerIndex)
	b.readerIndex++
	return val, nil
}

// ReadInt8 read an int8
func (b *ByteBuf) ReadInt8() (int8, error) {
	if b.readerIndex > b.writerIndex-1 {
		return 0, ErrOutOfBound
	}
	val, _ := b.GetInt8(b.readerIndex)
	b.readerIndex++
	return val, nil
}

// ReadUInt16BE read a uint16 in big endian
func (b *ByteBuf) ReadUInt16BE() (uint16, error) {
	if b.readerIndex > b.writerIndex-2 {
		return 0, ErrOutOfBound
	}
	val, _ := b.GetUInt16BE(b.readerIndex)
	b.readerIndex += 2
	return val, nil
}

// ReadUInt16LE read a uint16 in little endian
func (b *ByteBuf) ReadUInt16LE() (uint16, error) {
	if b.readerIndex > b.writerIndex-2 {
		return 0, ErrOutOfBound
	}
	val, _ := b.GetUInt16LE(b.readerIndex)
	b.readerIndex += 2
	return val, nil
}

// ReadInt16BE read an int16 in big endian
func (b *ByteBuf) ReadInt16BE() (int16, error) {
	if b.readerIndex > b.writerIndex-2 {
		return 0, ErrOutOfBound
	}
	val, _ := b.GetInt16BE(b.readerIndex)
	b.readerIndex += 2
	return val, nil
}

// ReadInt16LE read an int16 in little endian
func (b *ByteBuf) ReadInt16LE() (int16, error) {
	if b.readerIndex > b.writerIndex-2 {
		return 0, ErrOutOfBound
	}
	val, _ := b.GetInt16LE(b.readerIndex)
	b.readerIndex += 2
	return val, nil
}

// ReadUInt32BE read a uint32 in big endian
func (b *ByteBuf) ReadUInt32BE() (uint32, error) {
	if b.readerIndex > b.writerIndex-4 {
		return 0, ErrOutOfBound
	}
	val, _ := b.GetUInt32BE(b.readerIndex)
	b.readerIndex += 4
	return val, nil
}

// ReadUInt32LE read a uint32 in little endian
func (b *ByteBuf) ReadUInt32LE() (uint32, error) {
	if b.readerIndex > b.writerIndex-4 {
		return 0, ErrOutOfBound
	}
	val, _ := b.GetUInt32LE(b.readerIndex)
	b.readerIndex += 4
	return val, nil
}

// ReadInt32BE read an int32 in big endian
func (b *ByteBuf) ReadInt32BE() (int32, error) {
	if b.readerIndex > b.writerIndex-4 {
		return 0, ErrOutOfBound
	}
	val, _ := b.GetInt32BE(b.readerIndex)
	b.readerIndex += 4
	return val, nil
}

// ReadInt32LE read an int32 in little endian
func (b *ByteBuf) ReadInt32LE() (int32, error) {
	if b.readerIndex > b.writerIndex-4 {
		return 0, ErrOutOfBound
	}
	val, _ := b.GetInt32LE(b.readerIndex)
	b.readerIndex += 4
	return val, nil
}

// ReadUInt64BE read a uint64 in big endian
func (b *ByteBuf) ReadUInt64BE() (uint64, error) {
	if b.readerIndex > b.writerIndex-8 {
		return 0, ErrOutOfBound
	}
	val, _ := b.GetUInt64BE(b.readerIndex)
	b.readerIndex += 8
	return val, nil
}

// ReadUInt64LE read a uint64 in little endian
func (b *ByteBuf) ReadUInt64LE() (uint64, error) {
	if b.readerIndex > b.writerIndex-8 {
		return 0, ErrOutOfBound
	}
	val, _ := b.GetUInt64LE(b.readerIndex)
	b.readerIndex += 8
	return val, nil
}

// ReadInt64BE read an int64 in big endian
func (b *ByteBuf) ReadInt64BE() (int64, error) {
	if b.readerIndex > b.writerIndex-8 {
		return 0, ErrOutOfBound
	}
	val, _ := b.GetInt64BE(b.readerIndex)
	b.readerIndex += 8
	return val, nil
}

// ReadInt64LE read an int64 in little endian
func (b *ByteBuf) ReadInt64LE() (int64, error) {
	if b.readerIndex > b.writerIndex-8 {
		return 0, ErrOutOfBound
	}
	val, _ := b.GetInt64LE(b.readerIndex)
	b.readerIndex += 8
	return val, nil
}

// ReadFloat32BE read a float32 in big endian
func (b *ByteBuf) ReadFloat32BE() (float32, error) {
	if b.readerIndex > b.writerIndex-4 {
		return 0, ErrOutOfBound
	}
	val, _ := b.GetFloat32BE(b.readerIndex)
	b.readerIndex += 4
	return val, nil
}

// ReadFloat32LE read a float32 in little endian
func (b *ByteBuf) ReadFloat32LE() (float32, error) {
	if b.readerIndex > b.writerIndex-4 {
		return 0, ErrOutOfBound
	}
	val, _ := b.GetFloat32LE(b.readerIndex)
	b.readerIndex += 4
	return val, nil
}

// ReadFloat64BE read a float64 in big endian
func (b *ByteBuf) ReadFloat64BE() (float64, error) {
	if b.readerIndex > b.writerIndex-8 {
		return 0, ErrOutOfBound
	}
	val, _ := b.GetFloat64BE(b.readerIndex)
	b.readerIndex += 8
	return val, nil
}

// ReadFloat64LE read a float64 in little endian
func (b *ByteBuf) ReadFloat64LE() (float64, error) {
	if b.readerIndex > b.writerIndex-8 {
		return 0, ErrOutOfBound
	}
	val, _ := b.GetFloat64LE(b.readerIndex)
	b.readerIndex += 8
	return val, nil
}

// ReadFrom read from io.reader
func (b *ByteBuf) ReadFrom(reader io.Reader) error {
	for {
		chunk := make([]byte, ChunkSize)
		n, err := reader.Read(chunk)
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}
		_, err = b.Write(chunk[:n])
		if err != nil {
			return err
		}
	}
	return nil
}
