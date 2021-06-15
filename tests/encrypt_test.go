package tests

import (
	"api-crypto/crypto"
	"strings"
	"testing"
)

func TestEncrypt(t *testing.T) {
	decoded, _ := crypto.Encrypt(strings.Split("aaabbb", ""))
	sdecoded := strings.Join(decoded, "")
	decoded_real := "3a3b"
	if sdecoded != decoded_real {
		t.Errorf("%s not equal %s", sdecoded, decoded_real)
	}
	decoded, _ = crypto.Encrypt(strings.Split("aaabcceffcceff", ""))
	sdecoded = strings.Join(decoded, "")
	decoded_real = "3ab2(2ce2f)"
	if sdecoded != decoded_real {
		t.Errorf("%s not equal %s", sdecoded, decoded_real)
	}
	decoded, _ = crypto.Encrypt(strings.Split("aaababaaa", ""))
	sdecoded = strings.Join(decoded, "")
	decoded_real = "2a2(ab)3a"
 	if sdecoded != decoded_real {
		t.Errorf("%s not equal %s", sdecoded, decoded_real)
	}
	decoded, _ = crypto.Encrypt(strings.Split("aaaaaaaaaaaaaa", ""))
	sdecoded = strings.Join(decoded, "")
	decoded_real = "14a"
	if sdecoded != decoded_real {
		t.Errorf("%s not equal %s", sdecoded, decoded_real)
	}
}
