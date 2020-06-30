package main

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/binary"
	"fmt"
)

// T_S OMIT

// Truncate truncates b, a 20-byte big-endian byte slice,
// and returns an unsigned 32-bit int
func Truncate(b []byte) uint32 {
	lsb := b[19]     // least significant byte
	lsn := lsb & 0xf // least significant nibble
	trunc := binary.BigEndian.Uint32(b[lsn : lsn+4])

	truncMasked := trunc & 0x7fffffff
	otp := truncMasked % 1000000 // 6-digit OTP
	return otp
}

// T_E OMIT
func main() {
	// H_S OMIT
	key := []byte("12345678901234567890") // 1  // HL
	mac := hmac.New(sha1.New, key)
	// C_S OMIT

	buf := make([]byte, 8, 8)
	i := uint64(1) // 2: counter is set to 1  // HL
	binary.BigEndian.PutUint64(buf, i)

	mac.Write(buf)
	// C_E OMIT
	expectedMAC := mac.Sum(nil) // 3 // HL
	fmt.Printf("%x\n", expectedMAC)
	fmt.Println(hmac.Equal(expectedMAC, mac.Sum(nil))) // 4 // HL
	fmt.Printf("HOTP is %d\n", Truncate(expectedMAC))
	// H_E OMIT
}
