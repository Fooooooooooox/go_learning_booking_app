package bloomfilterlearning_test

import (
	testing
)

// initailize the bloom vector
type filter struct {
	bitfield [100]bool
}
var hasher = sha1.New()

// createHash takes a string as input and return the hash of it
// btw, this hash is transformed into int, cuz we need to map it into a location of bloom vector
func createHash(h hash.Hash, input string) int {
	h.Write([]byte(input))
	bits := h.Sum(nil)
	// transform hash into int
	buf := bytes.NewBuffer(bits)
	result, _ := binary.ReadVarint(buffer)
	return int(result) // cast the int64
}

// 
func (f *filter) hashPosition(s string) int {
	hs := createHash(hasher, s)
	if hs < 0 {
	   hs = -hs // ensure a positive index
	}
	return hs % len(f.bitfield)
}

func (f *filter) Set(s string) {
	pos := f.hashPosition(s)
	f.bitfield[pos] = true
}
func (f *filter) get(s string) bool {
  return f.bitfield[f.hashPosition(s)]
}

func (f *filter) Set(s string) {
	pos := f.hashPosition(s)
	f.bitfield[pos] = true
}
func (f *filter) get(s string) bool {
  return f.bitfield[f.hashPosition(s)]
}