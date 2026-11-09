// Package orchestrator contains deterministic task-planning primitives.
package orchestrator

import "strings"

type Step struct {
	ID               int    `json:"id"`
	Description      string `json:"description"`
	RequiresApproval bool   `json:"requires_approval"`
}

type Plan struct {
	Goal  string `json:"goal"`
	Steps []Step `json:"steps"`
}

// DraftPlan creates a transparent starter plan. It does not execute tools.
func DraftPlan(goal string) Plan {
	goal = strings.TrimSpace(goal)
	if goal == "" {
		return Plan{Steps: []Step{}}
	}
	return Plan{Goal: goal, Steps: []Step{
		{ID: 1, Description: "Clarify scope, success criteria, and authorization"},
		{ID: 2, Description: "Collect permitted context and identify constraints"},
		{ID: 3, Description: "Propose actions; route impactful changes for human approval", RequiresApproval: true},
		{ID: 4, Description: "Validate outcomes and produce an auditable report"},
	}}
}
