package engine

import (
	"fmt"

	"github.com/G1P0/hopless/internal/domain"
)

type Query struct {
	From  domain.Node
	To    domain.Node
	Proto string
	Port  int
}

// CanReach: last match wins, default deny.
func CanReach(w domain.World, q Query) (bool, string) {
	for i := len(w.Rules) - 1; i >= 0; i-- {
		r := w.Rules[i]
		if matches(r, q) {
			if r.Allow {
				return true, fmt.Sprintf("ALLOW by rule #%d (%s -> %s port=%d)", i, r.Src, r.Dst, r.Port)
			}
			return false, fmt.Sprintf("DENY by rule #%d (%s -> %s port=%d)", i, r.Src, r.Dst, r.Port)
		}
	}
	return false, "DENY by default (no matching rule)"
}

func matches(r domain.Rule, q Query) bool {
	if r.Src != q.From || r.Dst != q.To {
		return false
	}
	if r.Port != 0 && r.Port != q.Port { // 0 = ANY
		return false
	}
	return true
}
