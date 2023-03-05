package other

import (
	"math/rand"
	"testing"
)

func Test_Rand(t *testing.T) {
	// rand.Seed(1) GO 1.20版本弃用了
	for i := 0; i < 3; i++ {
		t.Log(rand.Intn(100))
	}
}
