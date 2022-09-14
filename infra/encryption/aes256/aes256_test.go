package aes256

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAes256(t *testing.T) {
	key := "a1b2c3d4e5f6g7h8"
	content := `{"use_info":"crmao"}`

	// encryption
	//content = "hello world"
	encrypted := Encrypt(content, key)
	// decryption
	decrypted := Decrypt(encrypted, key)
	assert.Equal(t, decrypted, content)
}
