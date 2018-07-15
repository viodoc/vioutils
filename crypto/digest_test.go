package crypto

import (
	"fmt"
	"testing"
)

func TestSha256Hex(t *testing.T) {
	data := "Hello"
	want := "185f8db32271fe25f561a6fc938b2e264306ec304eda518007d1764826381969"
	result := Sha256Hex(data)
	if result != want {
		t.Errorf("result must be %s", want)
	}
}

func TestMd5Hex(t *testing.T) {
	data := "Hello"
	want := "8b1a9953c4611296a827abf8c47804d7"
	result := Md5Hex(data)
	fmt.Println(len(want))
	fmt.Println(result)
	if result != want {
		t.Errorf("result must be %s", want)
	}
}
