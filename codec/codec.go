// Package codec 把玩家 id 编成邀请码并校验。
// 缺陷：字符集里留着易混淆字符、校验位跟着随机计数走、校验时不查校验位也不归一大小写。
package codec

import (
	"errors"
	"strings"
)

// ErrLength 长度不对。
var ErrLength = errors.New("邀请码长度不对")

// ErrChar 出现字符集之外的字符。
var ErrChar = errors.New("邀请码含非法字符")

// ErrChecksum 校验位不符。
var ErrChecksum = errors.New("邀请码校验位不符")

// Alphabet 是邀请码字符集（缺陷：里面还有 O 0 I 1）。
const Alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ012345"

// Length 是邀请码长度。
const Length = 6

var counter int

// Encode 把玩家 id 编成邀请码。缺陷：掺进了每次都变的计数器，同一个玩家会拿到不同码。
func Encode(player string) string {
	sum := 0
	for _, ch := range player {
		sum += int(ch)
	}
	counter++
	sum += counter
	out := make([]byte, Length)
	for index := 0; index < Length; index++ {
		out[index] = Alphabet[(sum+index*7+index*index)%len(Alphabet)]
	}
	return string(out)
}

// Verify 校验邀请码格式。缺陷：不查校验位、也不归一大小写。
func Verify(code string) error {
	if len(code) != Length {
		return ErrLength
	}
	for _, ch := range code {
		if !strings.ContainsRune(Alphabet, ch) {
			return ErrChar
		}
	}
	return nil
}
