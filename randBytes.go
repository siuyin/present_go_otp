package main

import (
	"encoding/base32"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// R_S OMIT
	rand.Seed(time.Now().UnixNano()) // seed random generator with current time
	b := make([]byte, 5, 5)          // make a slice of 5 bytes with a capacity of 5
	rand.Read(b)
	fmt.Println(b)

	s := base32.StdEncoding.EncodeToString(b) // Note: no 0 (zeros) in StdEncoding
	fmt.Println(s)                            // eg. CDIWOBLY
	d, _ := base32.StdEncoding.DecodeString(s)
	fmt.Println(d)
	// R_E OMIT
}
