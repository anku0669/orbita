// Package policy provides fail-closed authorization for tool execution.
package policy

import (
	"errors"
	"strings"
)

type Decision string

const (
	Allow  Decision = "allow"
	Deny   Decision = "deny"
	Review Decision = "review"
)

type Request struct {
	Tool       string
	Target     string
	Authorized bool
	ReadOnly   bool
}

type Result struct {
	Decision Decision `json:"decision"`
	Reason   string   `json:"reason"`
}

// Evaluate is intentionally conservative: unknown tools, absent authorization,
// and empty targets are denied. This is a baseline guard, not a sandbox.
func Evaluate(req Request) Result {
	tool := strings.TrimSpace(req.Tool)
	target := strings.TrimSpace(req.Target)
	if tool == "" {
		return Result{Deny, "tool name is required"}
	}
	if target == "" {
		return Result{Deny, "target is required"}
	}
	if !req.Authorized {
		return Result{Deny, "explicit authorization is required"}
	}
	switch strings.ToLower(tool) {
	case "workspace.read", "workspace.list", "report.render":
		return Result{Allow, "known low-impact operation"}
	case "workspace.write":
		if req.ReadOnly {
			return Result{Deny, "request conflicts with read-only mode"}
		}
		return Result{Review, "write operation requires human approval"}
	default:
		return Result{Deny, "tool is not registered in the baseline policy"}
	}
}

var ErrApprovalRequired = errors.New("human approval required")
