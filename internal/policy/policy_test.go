package policy

import "testing"

func TestEvaluateFailsClosed(t *testing.T) {
	cases := []struct {
		name string
		req  Request
		want Decision
	}{
		{"missing authorization", Request{Tool: "workspace.read", Target: "."}, Deny},
		{"unknown tool", Request{Tool: "shell.exec", Target: "host", Authorized: true}, Deny},
		{"safe read", Request{Tool: "workspace.read", Target: ".", Authorized: true}, Allow},
		{"write requires review", Request{Tool: "workspace.write", Target: "notes.txt", Authorized: true}, Review},
		{"missing target", Request{Tool: "workspace.read", Authorized: true}, Deny},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := Evaluate(tc.req).Decision; got != tc.want {
				t.Fatalf("decision = %q, want %q", got, tc.want)
			}
		})
	}
}
