package util

import (
	"fmt"
	"math/rand"
)

//removed inti() becouse from go 1.20 rand generator initialise automaticly

func GenerateID() string {
	hi := rand.Uint64()
	lo := rand.Uint32()
	return fmt.Sprintf("%x%x", hi, lo)
}
