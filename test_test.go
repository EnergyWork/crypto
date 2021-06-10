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

	decoded = Decrypt(strings.Split("2a", ""))
	sdecoded = strings.Join(decoded, "")
	decoded_real = "aa"
	if sdecoded != decoded_real {
		t.Errorf("%s not equal %s", sdecoded, decoded_real)
	}

	decoded = Decrypt(strings.Split("a2(er)2(vf)2r", ""))
	sdecoded = strings.Join(decoded, "")
	decoded_real = "aerervfvfrr"
	if sdecoded != decoded_real {
		t.Errorf("%s not equal %s", sdecoded, decoded_real)
	}

	decoded = Decrypt(strings.Split("2(ar)2(e2(ho))", ""))
	sdecoded = strings.Join(decoded, "")
	decoded_real = "ararehohoehoho"
	if sdecoded != decoded_real {
		t.Errorf("%s not equal %s", sdecoded, decoded_real)
	}

	decoded = Decrypt(strings.Split("(abc2d)", ""))
	sdecoded = strings.Join(decoded, "")
	decoded_real = "aabcdd"
	if sdecoded != decoded_real {
		t.Errorf("%s not equal %s", sdecoded, decoded_real)
	}
}
