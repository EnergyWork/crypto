package tests

import (
	"api-crypto/crypto"
	"strings"
	"testing"
)

func TestDecrypt(t *testing.T) {
	t.Log("1")
	decoded, _ := crypto.Decrypt(strings.Split("a2(c2(ab)f)d2c", ""))
	sdecoded := strings.Join(decoded, "")
	decoded_real := "acababfcababfdcc"
	if sdecoded != decoded_real {
		t.Errorf("%s not equal %s", sdecoded, decoded_real)
	}
	t.Log("2")
	decoded, _ = crypto.Decrypt(strings.Split("2a(c2e)", ""))
	sdecoded = strings.Join(decoded, "")
	decoded_real = "aacee"
	if sdecoded != decoded_real {
		t.Errorf("%s not equal %s", sdecoded, decoded_real)
	}
	t.Log("3")
	decoded, _ = crypto.Decrypt(strings.Split("2a", ""))
	sdecoded = strings.Join(decoded, "")
	decoded_real = "aa"
	if sdecoded != decoded_real {
		t.Errorf("%s not equal %s", sdecoded, decoded_real)
	}
	t.Log("4")
	decoded, _ = crypto.Decrypt(strings.Split("a2(er)2(vf)2r", ""))
	sdecoded = strings.Join(decoded, "")
	decoded_real = "aerervfvfrr"
	if sdecoded != decoded_real {
		t.Errorf("%s not equal %s", sdecoded, decoded_real)
	}
	t.Log("5")
	decoded, _ = crypto.Decrypt(strings.Split("2(ar)2(e2(ho))", ""))
	sdecoded = strings.Join(decoded, "")
	decoded_real = "ararehohoehoho"
	if sdecoded != decoded_real {
		t.Errorf("%s not equal %s", sdecoded, decoded_real)
	}
	t.Log("6")
	decoded, _ = crypto.Decrypt(strings.Split("(2a)", ""))
	sdecoded = strings.Join(decoded, "")
	decoded_real = "aa"
	if sdecoded != decoded_real {
		t.Errorf("%s not equal %s", sdecoded, decoded_real)
	}
}
