package orchestrator

import "testing"

func TestDraftPlanHasApprovalGate(t *testing.T) {
	p := DraftPlan("analyze my project")
	if len(p.Steps) < 3 {
		t.Fatal("expected multi-step plan")
	}
	found := false
	for _, step := range p.Steps {
		found = found || step.RequiresApproval
	}
	if !found {
		t.Fatal("plan must include an approval gate")
	}
}
