package utils

import (
	"IMChat/utils/encrypt"
	"fmt"
	"time"
)

func GenerateToken() string {
	str := fmt.Sprintf("%d", time.Now().Unix())
	return encrypt.Md5Encode(str)
}
