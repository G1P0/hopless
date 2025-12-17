package engine

import (
	"fmt"

	"github.com/G1P0/hopless/internal/domain"
)

func MissionText() string {
	return `Mission:
1) Allow client -> server on port 80
2) Deny  client -> server on port 22

World:
client ---- router ---- server

Notes:
- Rules are evaluated by "last match wins".
- Default is DENY.
- Port 0 means ANY.`
}

func MissionComplete(w domain.World) (bool, []string) {
	var reasons []string

	ok1, why1 := CanReachRouted(w, Query{
		From: domain.Client,
		To:   domain.Server,
		Port: 80,
	})
	if !ok1 {
		reasons = append(reasons, fmt.Sprintf("FAIL: client -> server:80 must be ALLOW (%s)", why1))
	} else {
		reasons = append(reasons, fmt.Sprintf("OK:   client -> server:80 is ALLOW (%s)", why1))
	}

	ok2, why2 := CanReachRouted(w, Query{
		From: domain.Client,
		To:   domain.Server,
		Port: 22,
	})
	if ok2 {
		reasons = append(reasons, fmt.Sprintf("FAIL: client -> server:22 must be DENY (%s)", why2))
	} else {
		reasons = append(reasons, fmt.Sprintf("OK:   client -> server:22 is DENY (%s)", why2))
	}

	return ok1 && !ok2, reasons
}
