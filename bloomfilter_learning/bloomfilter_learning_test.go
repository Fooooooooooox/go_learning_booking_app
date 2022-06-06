package bloomfilterlearning_test

import (
	"bytes"
	"crypto/sha1"
	"encoding/binary"
	"hash"
	"testing"

	"github.com/patrickmn/go-bloom"
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
	buffer := bytes.NewBuffer(bits)
	result, _ := binary.ReadVarint(buffer)
	return int(result) // cast the int64
}

// hashPosition takes a string and maps the hash of it into a location of bloomfilter vector
func (f *filter) hashPosition(s string) int {
	hs := createHash(hasher, s)
	if hs < 0 {
		hs = -hs // ensure a positive index
	}
	// return hs%len(vector)
	return hs % len(f.bitfield)
}

// set takes a string and set the corresponding location of bloom filter vector to 1
func (f *filter) Set(s string) {
	pos := f.hashPosition(s)
	f.bitfield[pos] = true
}

// get takes a string and return the bit in corresonding location
func (f *filter) get(s string) bool {
	return f.bitfield[f.hashPosition(s)]
}

// actually, go has a package for this
// go get github.com/patrickmn/go-bloom
// example:

func TestBloom(t *testing.T) {
	t.Log("hhh")
	// Create a bloom filter which will contain an expected 100,000 items, and which
	f := bloom.New(100000, 0.01)
	// t.Log(f)
	// Add an item to the filter
	t.Log("...adding foo")
	f.Add([]byte("foo"))
	// t.Log(f)
	// Check if an item has been added to the filter (if true, subject to the
	// false positive chance; if false, then the item definitely has not been
	// added to the filter.)
	t.Log(f.Test([]byte("foo")))
	f.Test([]byte("foo"))
}
