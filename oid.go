package oid

import (
	"encoding/binary"
	"encoding/hex"
	"math/rand"
	"sync/atomic"
	"time"
)

func random() uint32 {
	value := make([]byte, 4)
	rand.Seed(time.Now().UnixNano())
	rand.Read(value)
	return binary.BigEndian.Uint32(value)
}

var initial = random()

// New - describe
func New() string {
	// Generate the timestamp portion
	ts := make([]byte, 4)
	binary.BigEndian.PutUint32(
		ts[:], uint32(time.Now().Unix()),
	)

	// Generate the random portion
	rnd := make([]byte, 5)
	rand.Seed(time.Now().UnixNano())
	rand.Read(rnd)

	// Generate the increment portion
	inc := atomic.AddUint32(&initial, 1)

	// Build the object id
	var b []byte
	b = append(ts, rnd...)
	b = append(b, []byte{
		byte(inc >> 16),
		byte(inc >> 8),
		byte(inc),
	}...)

	return hex.EncodeToString(b[:])
}
