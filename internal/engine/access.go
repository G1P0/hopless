package engine

import (
	"fmt"

	"github.com/G1P0/hopless/internal/domain"
)

// CanReach: last match wins, default deny.
func CanReach(w domain.World, from, to domain.Node) (bool, string) {
	for i := len(w.Rules) - 1; i >= 0; i-- {
		r := w.Rules[i]
		if r.Src == from && r.Dst == to {
			if r.Allow {
				return true, fmt.Sprintf("ALLOW by rule #%d (%s -> %s)", i, r.Src, r.Dst)
			}
			return false, fmt.Sprintf("DENY by rule #%d (%s -> %s)", i, r.Src, r.Dst)
		}
	}
	return false, "DENY by default (no rule)"
}
