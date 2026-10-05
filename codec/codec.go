// Package codec 把玩家 id 编成邀请码并校验。
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

// Alphabet 是邀请码字符集：大写字母去掉 I、O，数字只用 2 到 9，共 32 个。
const Alphabet = "ABCDEFGHJKLMNPQRSTUVWXYZ23456789"

// Length 是邀请码长度。
const Length = 6

const base = len(Alphabet)

// slots 是前 5 位（数据位）能表示的空间，32^5。
const slots = base * base * base * base * base

var charIndex = func() [256]int {
	var table [256]int
	for i := range table {
		table[i] = -1
	}
	for i, ch := range Alphabet {
		table[byte(ch)] = i
	}
	return table
}()

// checksum 按 README 口径算校验位下标：
// 前 5 位字符序号（从 1 起）之和乘 31 再对 32 取模。
func checksum(data []byte) int {
	sum := 0
	for _, ch := range data {
		sum += charIndex[ch] + 1
	}
	return (sum * 31) % base
}

// Encode 把玩家 id 确定性地编成邀请码：前 5 位由 id 哈希填充，末位是校验位。
func Encode(player string) string {
	value := 0
	for i := 0; i < len(player); i++ {
		value = (value*31 + int(player[i])) % slots
	}
	out := make([]byte, Length)
	for i := Length - 2; i >= 0; i-- {
		out[i] = Alphabet[value%base]
		value /= base
	}
	out[Length-1] = Alphabet[checksum(out[:Length-1])]
	return string(out)
}

// Verify 校验邀请码：先统一大写，再依次判长度、字符集与校验位。
func Verify(code string) error {
	code = strings.ToUpper(code)
	if len(code) != Length {
		return ErrLength
	}
	for i := 0; i < Length; i++ {
		if charIndex[code[i]] < 0 {
			return ErrChar
		}
	}
	if charIndex[code[Length-1]] != checksum([]byte(code[:Length-1])) {
		return ErrChecksum
	}
	return nil
}
