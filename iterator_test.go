package bytebuf

import (
	. "gopkg.in/go-playground/assert.v1"
	"testing"
)

func TestFindNullIterator(t *testing.T) {
	b, err := NewByteBuf([]byte{0x01, 0x01, 0x00})
	Equal(t, err, nil)
	NotEqual(t, b, nil)

	Equal(t, b.ForEachByte(FindNullIterator), 2)
}

func TestFindNotNullIterator(t *testing.T) {
	b, err := NewByteBuf([]byte{0x00, 0x00, 0x01})
	Equal(t, err, nil)
	NotEqual(t, b, nil)

	Equal(t, b.ForEachByte(FindNotNullIterator), 2)
}

func TestFindCRIterator(t *testing.T) {
	b, err := NewByteBuf([]byte{0x01, 0x01, '\r'})
	Equal(t, err, nil)
	NotEqual(t, b, nil)

	Equal(t, b.ForEachByte(FindCRIterator), 2)
}

func TestFindNotCRIterator(t *testing.T) {
	b, err := NewByteBuf([]byte{'\r', '\r', 0x01})
	Equal(t, err, nil)
	NotEqual(t, b, nil)

	Equal(t, b.ForEachByte(FindNotCRIterator), 2)
}

func TestFindLFIterator(t *testing.T) {
	b, err := NewByteBuf([]byte{0x01, 0x01, '\n'})
	Equal(t, err, nil)
	NotEqual(t, b, nil)

	Equal(t, b.ForEachByte(FindLFIterator), 2)
}

func TestFindNotLFIterator(t *testing.T) {
	b, err := NewByteBuf([]byte{'\n', '\n', 0x01})
	Equal(t, err, nil)
	NotEqual(t, b, nil)

	Equal(t, b.ForEachByte(FindNotLFIterator), 2)
}

func TestFindCRLFIterator(t *testing.T) {
	b, err := NewByteBuf([]byte{0x01, 0x01, '\n', 0x01, 0x01, '\r'})
	Equal(t, err, nil)
	NotEqual(t, b, nil)

	Equal(t, b.ForEachByte(FindCRLFIterator), 2)
	Equal(t, b.ForEachByte(FindCRLFIterator, 3), 2)
}

func TestFindNotCRLFIterator(t *testing.T) {
	b, err := NewByteBuf([]byte{'\n', '\n', 0x01, '\r', '\r', 0x01})
	Equal(t, err, nil)
	NotEqual(t, b, nil)

	Equal(t, b.ForEachByte(FindNotCRLFIterator), 2)
	Equal(t, b.ForEachByte(FindNotCRLFIterator, 3), 2)
}

func TestFindWhiteSpaceIterator(t *testing.T) {
	b, err := NewByteBuf([]byte{0x01, 0x01, '\n', 0x01, 0x01, '\t', 0x01, 0x01, ' '})
	Equal(t, err, nil)
	NotEqual(t, b, nil)

	Equal(t, b.ForEachByte(FindWhiteSpaceIterator), 2)
	Equal(t, b.ForEachByte(FindWhiteSpaceIterator, 3), 2)
	Equal(t, b.ForEachByte(FindWhiteSpaceIterator, 6), 2)
}

func TestFindNotWhiteSpaceIterator(t *testing.T) {
	b, err := NewByteBuf([]byte{'\n', '\t', ' ', 0x01})
	Equal(t, err, nil)
	NotEqual(t, b, nil)

	Equal(t, b.ForEachByte(FindNotWhiteSpaceIterator), 3)
}
