// Package registry 是邀请码台账：玩家与码双向唯一。
// 缺陷：同一个玩家重复生成会拿到新码（不幂等）、查询码时全扫台账。
package registry

import (
	"errors"
	"sort"

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

// Assign 给玩家分配邀请码。缺陷：不查幂等，每次都重新编一个新码并覆盖旧记录。
func (r *Registry) Assign(player string) string {
	code := codec.Encode(player)
	r.byPlayer[player] = code
	r.byCode[code] = player
	return code
}

// Owner 返回码的归属玩家。缺陷：全扫台账。
func (r *Registry) Owner(code string) string {
	r.scanned += len(r.byPlayer)
	for player, value := range r.byPlayer {
		if value == code {
			return player
		}
	}
	return ""
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

// Scanned 是累计扫描的台账条目数（规模观测）。
func (r *Registry) Scanned() int { return r.scanned }
