package workspace

import "testing"

func TestResolveRejectsTraversal(t *testing.T) {
	s, err := New(t.TempDir())
	if err != nil {
		t.Fatal(err)
	}
	for _, input := range []string{"../escape", "../../etc/passwd"} {
		if _, err := s.Resolve(input); err == nil {
			t.Fatalf("expected traversal rejection for %q", input)
		}
	}
	if _, err := s.Resolve("reports/today.md"); err != nil {
		t.Fatalf("safe path rejected: %v", err)
	}
}
