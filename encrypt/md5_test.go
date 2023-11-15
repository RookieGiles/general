package encrypt

import (
	"fmt"
	"testing"
)

func TestMd5(t *testing.T) {
	str := "123456"

	fmt.Println(Md5(str))
}
