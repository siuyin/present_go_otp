package main

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/binary"
	"fmt"
)

// T_S OMIT
func Truncate(b []byte) uint32 {
	lsb := b[19]
	lsn := lsb & 0xf // least significant nibble
	trunc := binary.BigEndian.Uint32(b[lsn : lsn+4])
	truncMasked := trunc & 0x7fffffff
	otp := truncMasked % 1000000 // 6-digit OTP
	return otp
}

// T_E OMIT
func main() {
	// H_S OMIT
	key := []byte("12345678901234567890")
	mac := hmac.New(sha1.New, key)
	// C_S OMIT
	mac.Write([]byte{0, 0, 0, 0, 0, 0, 0, 1}) // HOTP specifies 8-byte counter
	// C_E OMIT
	expectedMAC := mac.Sum(nil)
	fmt.Printf("%x\n", expectedMAC)
	fmt.Println(hmac.Equal(expectedMAC, mac.Sum(nil)))
	fmt.Printf("HOTP is %d\n", Truncate(expectedMAC)) // HL
	// H_E OMIT
}
