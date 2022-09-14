package hash

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMd5(t *testing.T) {
	res := Md5("MAOZHONGYU")
	//对应PHP方法md5() ,长度32
	assert.Equal(t, "2d0a0f514a039d97e3f273c7ae7d2158", res)
}

func TestSha256(t *testing.T) {
	res := Sha256("MAOZHONGYU")
	// 对应php 方法  hash("sha256","MAOZHONGYU");   长度64
	assert.Equal(t, "32e4e2d7be8996c7d70ac44575e967ee02eadf3d2f5b6a9354a2a84fc834375a", res)
}
