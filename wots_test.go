package xmssmt

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"testing"
)

func testWotsGenChain(params *Params, expect string, t *testing.T) {
	var pubSeed []byte = make([]byte, params.N)
	var in []byte = make([]byte, params.N)
	var addr [8]uint32
	for i := 0; i < int(params.N); i++ {
		pubSeed[i] = byte(2 * i)
		in[i] = byte(i)
	}
	for i := 0; i < 8; i++ {
		addr[i] = 500000000 * uint32(i)
	}
	val := hex.EncodeToString(params.wotsGenChain(
		in, 4, 5, pubSeed, address(addr)))
	if val != expect {
		t.Errorf("%s wotsGenChain returned %s instead of %s", params.Name(), val, expect)
	}
}

func TestWotsGenChain(t *testing.T) {
	testWotsGenChain(ParamsFromOid(false, 1), "2dd7fcc039afb02d35c4b370172a7714b909d74a6ef2463538e87b05ab573d18", t)
	testWotsGenChain(ParamsFromOid(false, 4), "9b4cda48d43e57bf4b5eb57c7bd86126d523517f9f27dbe287c8501d3c00f4f1e37fab649ac4bec337bc92623acc837af3ac5be17ed1624a335eb02d0771a68c", t)
	testWotsGenChain(ParamsFromOid(false, 7), "14f78e435e3758a862fedea60af053374390d9cc3b140a2221e03281b2d84cf0", t)
	testWotsGenChain(ParamsFromOid(false, 10), "252e91e199a755ef156c9671f1e35d1853653f2956a167bc548ae3def7fc7f0842f2825ed674c212cb156c0c2908c8d3835d22c5aaf1140bcc0cffdc8b96b89f", t)
}

func testWotsPkGen(params *Params, expect string, t *testing.T) {
	var pubSeed []byte = make([]byte, params.N)
	var seed []byte = make([]byte, params.N)
	var addr [8]uint32
	for i := 0; i < int(params.N); i++ {
		pubSeed[i] = byte(2 * i)
		seed[i] = byte(i)
	}
	for i := 0; i < 8; i++ {
		addr[i] = 500000000 * uint32(i)
	}
	valHash := sha256.Sum256(
		params.wotsPkGen(seed, pubSeed, address(addr)))
	valHashPref := hex.EncodeToString(valHash[:8])
	if valHashPref != expect {
		t.Errorf("%s hash of wotsPkGen return value starts with %s instead of %s", params.Name(), valHashPref, expect)
	}
}

func TestWotsPkGen(t *testing.T) {
	testWotsPkGen(ParamsFromOid(false, 1), "4bad377b36d488f0", t)
	testWotsPkGen(ParamsFromOid(false, 4), "2da374aa0c3c48cf", t)
	testWotsPkGen(ParamsFromOid(false, 7), "a63529f0d6c4c965", t)
	testWotsPkGen(ParamsFromOid(false, 10), "65ab3d40673846d7", t)
}

func testWotsSign(params *Params, expect string, t *testing.T) {
	var pubSeed []byte = make([]byte, params.N)
	var seed []byte = make([]byte, params.N)
	var msg []byte = make([]byte, params.N)
	var addr [8]uint32
	for i := 0; i < int(params.N); i++ {
		pubSeed[i] = byte(2 * i)
		seed[i] = byte(i)
		msg[i] = byte(3 * i)
	}
	for i := 0; i < 8; i++ {
		addr[i] = 500000000 * uint32(i)
	}
	valHash := sha256.Sum256(
		params.wotsSign(msg, seed, pubSeed, address(addr)))
	valHashPref := hex.EncodeToString(valHash[:8])
	if valHashPref != expect {
		t.Errorf("%s hash of wotsSign return value starts with %s instead of %s", params.Name(), valHashPref, expect)
	}
}

func TestWotsSign(t *testing.T) {
	testWotsSign(ParamsFromOid(false, 1), "ddef75e06556e4a0", t)
	testWotsSign(ParamsFromOid(false, 4), "eaca616e882a8afc", t)
	testWotsSign(ParamsFromOid(false, 7), "03c64b093f123bb9", t)
	testWotsSign(ParamsFromOid(false, 10), "3b526d0b89d463c7", t)
}

func testWotSignThenVerify(params *Params, t *testing.T) {
	var pubSeed []byte = make([]byte, params.N)
	var seed []byte = make([]byte, params.N)
	var msg []byte = make([]byte, params.N)
	var addr [8]uint32
	for i := 0; i < int(params.N); i++ {
		pubSeed[i] = byte(2 * i)
		seed[i] = byte(i)
		msg[i] = byte(3 * i)
	}
	for i := 0; i < 8; i++ {
		addr[i] = 500000000 * uint32(i)
	}
	sig := params.wotsSign(msg, seed, pubSeed, address(addr))
	pk1 := params.wotsPkFromSig(sig, msg, pubSeed, address(addr))
	pk2 := params.wotsPkGen(seed, pubSeed, address(addr))
	if !bytes.Equal(pk1, pk2) {
		t.Errorf("%s verification of signature failed", params.Name())
	}
}

func TestWotsSignThenVerify(t *testing.T) {
	testWotSignThenVerify(ParamsFromOid(false, 1), t)
	testWotSignThenVerify(ParamsFromOid(false, 4), t)
	testWotSignThenVerify(ParamsFromOid(false, 7), t)
	testWotSignThenVerify(ParamsFromOid(false, 10), t)
}
