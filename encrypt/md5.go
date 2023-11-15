package encrypt

import (
	"crypto/md5"
	"encoding/hex"
)

var (
	defaultSalt string = "12321"
	buildSalt   string
)

// Md5 对字符串进行md5 加密操作 默认盐为 salt
func Md5(str string, salt ...string) string {
	if len(salt) == 0 || salt[0] == "" {
		buildSalt = defaultSalt
	} else {
		buildSalt = salt[0]
	}

	buildStr := str + buildSalt

	data := []byte(buildStr)
	sum := md5.Sum(data)
	// hex转字符串
	md5Str := hex.EncodeToString(sum[:])

	return md5Str
}
