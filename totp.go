package main

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base32"
	"encoding/binary"
	"fmt"
	"time"
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
	key, _ := base32.StdEncoding.DecodeString("CDIWOBLY")
	mac := hmac.New(sha1.New, key)
	timeStep := time.Now().Unix() / 30
	fmt.Printf("timestep is %d\n", timeStep)
	buf := &bytes.Buffer{}
	binary.Write(buf, binary.BigEndian, timeStep)
	mac.Write(buf.Bytes())
	expectedMAC := mac.Sum(nil)
	fmt.Printf("TOTP is %06d\n", Truncate(expectedMAC))
	// H_E OMIT
}
