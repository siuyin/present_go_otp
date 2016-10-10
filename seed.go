package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// S_S OMIT
	rand.Seed(time.Now().UnixNano()) // seed random generator with current time
	fmt.Println(rand.Intn(100))      // a random number in [0,100)
	// S_E OMIT
}
