package oid

import (
	"encoding/binary"
	"encoding/hex"
	"math/rand"
	"sync/atomic"
	"time"
)

func randCount() uint32 {
	counter := make([]byte, 4)
	rand.Seed(time.Now().UnixNano())
	rand.Read(counter)
	return binary.BigEndian.Uint32(counter)
}

var counter = randCount()

// New - describe
func New() string {
	// Hold the result
	var b []byte

	// Generate the timestamp portion
	timestamp := make([]byte, 4)
	binary.BigEndian.PutUint32(
		timestamp[:], uint32(time.Now().Unix()),
	)

	// Generate the random portion
	random := make([]byte, 5)
	rand.Seed(time.Now().UnixNano())
	rand.Read(random)

	// Generate the increment portion
	increment := atomic.AddUint32(&counter, 1)

	// Build the object id
	b = append(timestamp, random...)
	b = append(b, []byte{
		byte(increment >> 16),
		byte(increment >> 8),
		byte(increment),
	}...)

	return hex.EncodeToString(b[:])
}
