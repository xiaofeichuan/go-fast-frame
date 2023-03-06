package utils

import (
	"math/rand"
	"time"
)

type TextUtil struct{}

var Text = new(TextUtil)

// GetRoundNumber 获取随机数
func (TextUtil) GetRoundNumber(size int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyz"
	bytes := []byte(str)
	var result []byte
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < size; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}
