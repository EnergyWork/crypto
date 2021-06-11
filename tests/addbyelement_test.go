package tests

import (
	"api-crypto/crypto"
	"fmt"
	"strings"
	"testing"
)

func TestAddByElement(t *testing.T) {
	var seq = []string{"a", "b"}
	var str = "cd"
	crypto.AddByElement(&seq, strings.Split(str, ""))
	fmt.Println(seq)
	if len(seq) != 4 {
		t.Errorf("Error")
	}
}
