package generic

import "github.com/segmentio/fasthash/fnv1a"

type Uint8 uint8
type Uint16 uint16
type Uint32 uint32
type Uint64 uint64

type Int8 int8
type Int16 int16
type Int32 int32
type Int64 int64

type Float32 float32
type Float64 float64

type Bool bool

type String string
type ByteSlice []byte

func (u Uint8) Less(other Uint8) bool {
	return u < other
}
func (u Uint8) Equals(other Uint8) bool {
	return u == other
}
func (u Uint8) Hash() uint64 {
	return hash(uint64(u))
}

func (u Uint16) Less(other Uint16) bool {
	return u < other
}
func (u Uint16) Equals(other Uint16) bool {
	return u == other
}
func (u Uint16) Hash() uint64 {
	return hash(uint64(u))
}

func (u Uint32) Less(other Uint32) bool {
	return u < other
}
func (u Uint32) Equals(other Uint32) bool {
	return u == other
}
func (u Uint32) Hash() uint64 {
	return hash(uint64(u))
}

func (u Uint64) Less(other Uint64) bool {
	return u < other
}
func (u Uint64) Equals(other Uint64) bool {
	return u == other
}
func (u Uint64) Hash() uint64 {
	return hash(uint64(u))
}

func (i Int8) Less(other Int8) bool {
	return i < other
}
func (i Int8) Equals(other Int8) bool {
	return i == other
}
func (i Int8) Hash() uint64 {
	return hash(uint64(i))
}

func (i Int16) Less(other Int16) bool {
	return i < other
}
func (i Int16) Equals(other Int16) bool {
	return i == other
}
func (i Int16) Hash() uint64 {
	return hash(uint64(i))
}

func (i Int32) Less(other Int32) bool {
	return i < other
}
func (i Int32) Equals(other Int32) bool {
	return i == other
}
func (i Int32) Hash() uint64 {
	return hash(uint64(i))
}

func (i Int64) Less(other Int64) bool {
	return i < other
}
func (i Int64) Equals(other Int64) bool {
	return i == other
}
func (i Int64) Hash() uint64 {
	return hash(uint64(i))
}

func (b Bool) Equals(other Bool) bool {
	return b == other
}
func (b Bool) Hash() uint64 {
	if b {
		return 1
	}
	return 0
}

func (s String) Less(other String) bool {
	return s < other
}
func (s String) Equals(other String) bool {
	return s == other
}
func (s String) Hash() uint64 {
	return fnv1a.HashString64(string(s))
}

func (b ByteSlice) Equals(other ByteSlice) bool {
	// TODO: update when stdlib slices package is available
	for i := range b {
		if b[i] != other[i] {
			return false
		}
	}
	return true
}
func (b ByteSlice) Hash() uint64 {
	return fnv1a.HashBytes64([]byte(b))
}

func hash(u uint64) uint64 {
	u ^= u >> 33
	u *= 0xff51afd7ed558ccd
	u ^= u >> 33
	u *= 0xc4ceb9fe1a85ec53
	u ^= u >> 33
	return u
}