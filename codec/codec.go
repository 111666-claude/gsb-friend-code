// Package codec 把玩家 id 编成邀请码并校验。
package codec

import (
	"errors"
	"hash/fnv"
	"strings"
)

// ErrLength 长度不对。
var ErrLength = errors.New("邀请码长度不对")

// ErrChar 出现字符集之外的字符。
var ErrChar = errors.New("邀请码含非法字符")

// ErrChecksum 校验位不符。
var ErrChecksum = errors.New("邀请码校验位不符")

// Alphabet 是邀请码字符集：大写字母去掉 I、O，数字只用 2 到 9，共 32 个。
const Alphabet = "ABCDEFGHJKLMNPQRSTUVWXYZ23456789"

// Length 是邀请码长度。
const Length = 6

// Encode 把玩家 id 编成邀请码：同一个玩家永远得到同一个码。
func Encode(player string) string {
	hasher := fnv.New64a()
	hasher.Write([]byte(player))
	hash := hasher.Sum64()
	out := make([]byte, Length)
	for index := 0; index < Length-1; index++ {
		out[index] = Alphabet[hash&uint64(len(Alphabet)-1)]
		hash >>= 5
	}
	out[Length-1] = Alphabet[checksum(out[:Length-1])]
	return string(out)
}

// checksum 取前 5 位字符序号（从 1 数起）和乘 31 再对 32 取模得到的下标。
func checksum(body []byte) int {
	sum := 0
	for _, ch := range body {
		sum += strings.IndexByte(Alphabet, ch) + 1
	}
	return sum * 31 % len(Alphabet)
}

// Verify 校验邀请码：先统一成大写，再判长度、字符与校验位。
func Verify(code string) error {
	upper := strings.ToUpper(code)
	if len(upper) != Length {
		return ErrLength
	}
	for _, ch := range upper {
		if !strings.ContainsRune(Alphabet, ch) {
			return ErrChar
		}
	}
	if Alphabet[checksum([]byte(upper[:Length-1]))] != upper[Length-1] {
		return ErrChecksum
	}
	return nil
}
