package engine

import (
	"fmt"

	"github.com/G1P0/hopless/internal/domain"
)

func MissionText() string {
	return `Mission:
1) Allow client -> server
2) Deny  client -> router

World:
client ---- router ---- server

Notes:
- Rules are evaluated by "last match wins".
- Default is DENY.`
}

func MissionComplete(w domain.World) (bool, []string) {
	var reasons []string

	ok1, why1 := CanReach(w, domain.Client, domain.Server)
	if !ok1 {
		reasons = append(reasons, fmt.Sprintf("FAIL: client -> server must be ALLOW (%s)", why1))
	} else {
		reasons = append(reasons, fmt.Sprintf("OK:   client -> server is ALLOW (%s)", why1))
	}

	ok2, why2 := CanReach(w, domain.Client, domain.Router)
	if ok2 {
		reasons = append(reasons, fmt.Sprintf("FAIL: client -> router must be DENY (%s)", why2))
	} else {
		reasons = append(reasons, fmt.Sprintf("OK:   client -> router is DENY (%s)", why2))
	}

	return ok1 && !ok2, reasons
}
