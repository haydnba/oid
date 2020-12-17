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
	tstmp := make([]byte, 4)
	binary.BigEndian.PutUint32(
		tstmp[:], uint32(time.Now().Unix()),
	)

	// Generate the random portion
	token := make([]byte, 5)
	rand.Seed(time.Now().UnixNano())
	rand.Read(token)

	// Generate the increment portion
	count := atomic.AddUint32(&counter, 1)

	// Build the object id
	b = append(tstmp, token...)
	b = append(b, make([]byte, 3)...)
	b[9] = byte(count >> 16)
	b[10] = byte(count >> 8)
	b[11] = byte(count)

	return hex.EncodeToString(b[:])
}
