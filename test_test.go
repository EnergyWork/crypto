package main

import (
	"strings"
	"testing"
)

func TestDecrypt(t *testing.T) {
	decoded := Decrypt(strings.Split("a2(c2(ab)f)d2c", ""))
	sdecoded := strings.Join(decoded, "")
	decoded_real := "acababfcababfdcc"
	if sdecoded != decoded_real {
		t.Errorf("%s not equal %s", sdecoded, decoded_real)
	}

	decoded = Decrypt(strings.Split("2a(c2e)", ""))
	sdecoded = strings.Join(decoded, "")
	decoded_real = "aacee"
	if sdecoded != decoded_real {
		t.Errorf("%s not equal %s", sdecoded, decoded_real)
	}

	decoded = Decrypt(strings.Split("22a", ""))
	sdecoded = strings.Join(decoded, "")
	decoded_real = "aaaa"
	if sdecoded != decoded_real {
		t.Errorf("%s not equal %s", sdecoded, decoded_real)
	}
}
