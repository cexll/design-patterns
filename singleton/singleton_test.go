package singleton

import (
	"fmt"
	"testing"
)

func TestSingle(t *testing.T) {
	for i := 0; i <= 10; i++ {
		go GetInstance()
	}

	fmt.Scanln()
}
