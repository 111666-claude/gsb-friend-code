// Package registry 是邀请码台账：玩家与码双向唯一。
package registry

import (
	"errors"
	"sort"
	"strconv"
	"strings"

	"example.com/friendcode/codec"
)

// ErrTaken 目标码已被别的玩家占用。
var ErrTaken = errors.New("邀请码已被占用")

// Registry 是邀请码台账。
type Registry struct {
	byPlayer map[string]string
	byCode   map[string]string
	scanned  int
}

// New 建空台账。
func New() *Registry {
	return &Registry{byPlayer: map[string]string{}, byCode: map[string]string{}}
}

// Assign 给玩家分配邀请码：同一玩家幂等返回同一个码；码被别人占了就换盐重编。
func (r *Registry) Assign(player string) string {
	if code, ok := r.byPlayer[player]; ok {
		return code
	}
	code := codec.Encode(player)
	for salt := 1; ; salt++ {
		if _, taken := r.byCode[code]; !taken {
			break
		}
		code = codec.Encode(player + "\x00" + strconv.Itoa(salt))
	}
	r.byPlayer[player] = code
	r.byCode[code] = player
	return code
}

// Owner 按码反查归属玩家，O(1) 哈希查询，不全扫台账。
func (r *Registry) Owner(code string) string {
	r.scanned++
	return r.byCode[strings.ToUpper(code)]
}

// Players 返回已分配的玩家（升序）。
func (r *Registry) Players() []string {
	out := []string{}
	for player := range r.byPlayer {
		out = append(out, player)
	}
	sort.Strings(out)
	return out
}

// Scanned 是累计查询探测的台账条目数（规模观测）。
func (r *Registry) Scanned() int { return r.scanned }
