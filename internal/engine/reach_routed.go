package engine

import (
	"fmt"

	"github.com/G1P0/hopless/internal/domain"
)

func CanReachRouted(w domain.World, q Query) (bool, string) {
	path, ok := FindPath(w, q.From, q.To)
	if !ok {
		return false, "NO ROUTE TO HOST"
	}

	for i := 0; i < len(path)-1; i++ {
		hopFrom := path[i]
		hopTo := path[i+1]

		ok, why := CanReach(w, Query{
			From: hopFrom,
			To:   hopTo,
			Port: q.Port,
		})
		if !ok {
			return false, fmt.Sprintf(
				"DROP at hop %s -> %s (%s)",
				hopFrom, hopTo, why,
			)
		}
	}

	return true, fmt.Sprintf("OK via path %v", path)
}
