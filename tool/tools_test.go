package tool

import (
	"fmt"
	"testing"
)

func TestGetAddress(t *testing.T) {
	//GetAddressFromPubKey("03870f2aebd7ac762bf26de14bf4624781cd4e4ed3ca4ada16c883f1d7a492ec0a")

	SetAesKey("03870f2aebd7ac762bf26de14bf4624781cd4e4ed3ca4ada16c883f1d7a492ec0a"[0:32])
	se, err := AesEncrypt("aes-20170416-30-1000")
	fmt.Println(se, err)
	sd, err := AesDecrypt(se)
	fmt.Println(sd, err)
}
