package aes

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAesCBC(t *testing.T) {
	key := []byte("ir489u58ir489u54") //key必须是长度为16的byte数组
	plain := `sksdkweoweowoe`
	cipher, err := AesEncrypt(plain, key)
	//fmt.Println(len(cipher))
	//fmt.Println(len(cipher))
	if err != nil {
		t.Logf("加密错误 %s", err)
	}
	//t.Log("密文:", cipher)
	plain, err = AesDecrypt(cipher, key)
	if err != nil {
		t.Logf("解密错误 %s", err)
	}
	//fmt.Printf("明文：%s\n", plain)

}

//
func TestAesCtr(t *testing.T) {
	plain := `速度快速度快肯定是`

	text := []byte(plain)
	//fmt.Println(len(text))
	key := []byte("1a1a1a1a1a1a1a1a")
	res := AesCtrEncrypt(text, key)
	//fmt.Println(string(res))
	//fmt.Println(len(res))
	//fmt.Println(len(res))
	res1 := AesCtrDecrypt(res, key)
	assert.Equal(t, string(res1), plain)
}
