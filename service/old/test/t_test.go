package test

import (
	"math/rand"
	"testing"
	"time"
)

func TestRandom(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10000; i++ {
		a := rand.Intn(10)
		if a >= 10 {
			t.Errorf("fuck ")
		}
	}
}
