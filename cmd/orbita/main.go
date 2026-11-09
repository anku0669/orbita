package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/anku0669/orbita/internal/config"
	"github.com/anku0669/orbita/internal/orchestrator"
	"github.com/anku0669/orbita/internal/policy"
	"github.com/anku0669/orbita/internal/workspace"
)

const version = "0.1.0-dev"

func main() {
	if len(os.Args) < 2 {
		usage()
		return
	}
	switch os.Args[1] {
	case "version":
		fmt.Printf("ORBITA %s\n", version)
	case "status":
		cfg, err := config.Load()
		fatalIf(err)
		store, err := workspace.New(cfg.Workspace)
		fatalIf(err)
		fmt.Printf("ORBITA %s\nworkspace: %s\nlog level: %s\n", version, store.Root(), cfg.LogLevel)
	case "plan":
		fs := flag.NewFlagSet("plan", flag.ExitOnError)
		goal := fs.String("goal", "", "goal to plan")
		_ = fs.Parse(os.Args[2:])
		writeJSON(orchestrator.DraftPlan(*goal))
	case "policy-check":
		fs := flag.NewFlagSet("policy-check", flag.ExitOnError)
		tool := fs.String("tool", "", "registered tool")
		target := fs.String("target", "", "operation target")
		authorized := fs.Bool("authorized", false, "assert explicit authorization")
		readOnly := fs.Bool("read-only", false, "disallow writes")
		_ = fs.Parse(os.Args[2:])
		writeJSON(policy.Evaluate(policy.Request{Tool: *tool, Target: *target, Authorized: *authorized, ReadOnly: *readOnly}))
	default:
		usage()
		os.Exit(2)
	}
}

func usage() {
	fmt.Print(`ORBITA: Open Runtime for Building Intelligent Task Agents

Usage:
  orbita version
  orbita status
  orbita plan --goal "describe a goal"
  orbita policy-check --tool workspace.read --target ./ --authorized
`)
}
func writeJSON(v any) {
	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	fatalIf(enc.Encode(v))
}
func fatalIf(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}
