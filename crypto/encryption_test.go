package crypto

import (
	"testing"
	"fmt"
)
var key16Text="qIcbf8KhDglhwzTQ"

func TestEncrypt(t *testing.T) {
	data := "Hello"
	want := "-SBbgC8="
	result := Encrypt(data,key16Text)
	fmt.Printf("result is %s",result)
	if result != want {
		t.Errorf("result must be %s", want)
	}
}

func TestDecrypt(t *testing.T) {
	data := "-SBbgC8="
	want := "Hello"
	result := Decrypt(data,key16Text)
	fmt.Printf("result is %s",result)
	if result != want {

		t.Errorf("result must be %s", want)
	}
}

func TestAesEncrypt(t *testing.T) {
	data := "1"
	want := "FiPRVth1I0srMiwtSnfdpw=="
	result := AesEncrypt(data,key16Text)
	if result != want {
		fmt.Printf("result is %s",result)
		t.Errorf("result must be %s", want)
	}
}

func TestAesDecrypt(t *testing.T) {
	data := "9fHBKYvjm9VJuoQ_fEQkgA=="
	want := "Hello world"
	result := AesDecrypt(data,key16Text)
	if result != want {
		t.Errorf("result must be %s", want)
	}
}

