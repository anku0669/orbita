# ORBITA: A Policy-Gated, Human-Supervised Runtime for Portable Agent Workflows

**Research manuscript draft — version 0.1 — 9 November 2026**  
**Author:** ORBITA contributors  
**Status:** Design and evaluation protocol; empirical results not yet collected.

## Abstract

Tool-using AI agents combine model-generated plans with operations that can affect files, services, and external systems. This creates a control problem: a natural-language request may be ambiguous, model output is probabilistic, and an otherwise useful tool can cause unintended effects when authorization and execution are coupled. This paper introduces ORBITA, a research prototype for policy-gated agent workflows. ORBITA separates planning from execution, represents tool requests as structured operations, applies a fail-closed policy decision, and routes state-changing operations through human review. Its core is provider-neutral and implemented in Go to support reproducible builds across Windows, macOS, and Linux. We present the architecture, threat model, prototype scope, and a reproducible evaluation protocol. We do not claim measured security or performance gains: those claims require the proposed experiments. The central research hypothesis is that explicit policy gates and auditable intermediate representations can reduce unauthorized tool execution while retaining acceptable task completion and operator effort.

**Keywords:** agentic systems, tool use, AI safety, authorization, human-in-the-loop, policy enforcement, reproducible systems, cross-platform software.

## 1. Introduction

Agentic applications increasingly connect language-model reasoning to tools. This can improve task completion, but it also creates a path from ambiguous natural-language intent to real effects. Prompt injection, confused-deputy behavior, over-broad credentials, unsafe tool composition, and insufficient observability are practical concerns. A design that relies only on model instructions places the authorization boundary inside the same probabilistic component that proposes actions.

ORBITA explores a different division of responsibility. A planner may propose a sequence of operations, but a separate policy layer decides whether each operation is allowed, denied, or requires human review. The model is treated as an untrusted proposer rather than an authority. The runtime is deliberately provider-neutral so the control boundary can be evaluated across model backends.

This paper makes three contributions:

1. A reference architecture separating goal interpretation, structured planning, policy evaluation, tool execution, and audit reporting.
2. A fail-closed baseline policy with explicit approval requirements for state-changing operations.
3. An experimental protocol for measuring policy effectiveness, task utility, portability, latency, and operator burden.

## 2. Problem statement and research questions

Let a workflow consist of a goal `g`, a proposed plan `P`, a set of registered tools `T`, an authorization context `A`, and an execution trace `E`. The system should maximize valid task completion while minimizing unauthorized side effects.

**RQ1:** Does an external policy gate reduce the rate of unauthorized tool invocations relative to a baseline where tool authorization is delegated to agent prompting?

**RQ2:** What is the task-completion and latency cost of routing state-changing operations through human approval?

**RQ3:** How do provider choice and prompt-injection attempts affect policy decisions when policy inputs are structured and independent of model prose?

**RQ4:** Can a single Go core produce functionally equivalent command-line behavior across Windows, macOS (Intel and Apple Silicon), and Linux architectures?

## 3. System design

### 3.1 Components

- **CLI and configuration:** parses commands and validates runtime configuration.
- **Planner:** emits inspectable steps. The current prototype uses deterministic plan templates; model-backed planning is future work.
- **Provider interface:** defines a backend-neutral message/response contract. Adapters are not yet implemented.
- **Tool registry:** maps stable operation identifiers to vetted handlers. The current prototype includes policy names and bounded workspace primitives, not arbitrary shell access.
- **Policy engine:** returns `allow`, `deny`, or `review`. Empty targets, missing authorization, and unknown tools are denied. Workspace writes require review.
- **Workspace:** provides local file operations scoped to a configured root. Full symlink- and race-safe containment remains future hardening.
- **Audit/report layer:** structured reporting is planned; persistent, tamper-evident storage is not yet implemented.

### 3.2 Policy semantics

For a request `r`, the baseline evaluator computes:

`D(r) ∈ {ALLOW, DENY, REVIEW}`.

The evaluator denies if the tool identifier or target is missing, or if authorization is absent. A small allowlist of low-impact operations may pass. State-changing operations are routed to review. Unknown operations are denied. This is intentionally conservative. It is not a proof of safety, nor does it replace OS-level sandboxing, credential scoping, or operator training.

### 3.3 Portability

The implementation language is Go. The release pipeline cross-compiles the CLI for Linux amd64/arm64, macOS amd64/arm64, and Windows amd64/arm64. Cross-compilation establishes build portability, not complete runtime compatibility. Each target must be exercised on native runners before release claims are made. Desktop installers and code signing are outside the current prototype.

## 4. Threat model

### 4.1 Assets

- Local files and workspace contents.
- Credentials and provider API tokens.
- Integrity of tool execution and generated reports.
- Availability and resource budgets.
- Operator authorization decisions.

### 4.2 Adversaries and failure modes

We consider malicious or compromised web content, prompt injection embedded in retrieved documents, malicious tool output, a model proposing an unsafe operation, accidental over-broad authorization, path traversal, symlink races, credential leakage, and dependency compromise. We also consider benign operator errors and model hallucinations.

### 4.3 Security assumptions and non-goals

The prototype assumes the host OS and Go runtime are not compromised. A boolean authorization field in the current command-line demonstration is not a production credential or cryptographic approval token. The prototype does not provide a hardened sandbox, arbitrary remote execution, distributed multi-tenant isolation, formal verification, or protection from a fully compromised host. The path resolver currently performs lexical traversal checks; it must be hardened with canonical path and symlink-aware checks before untrusted workloads are supported.

## 5. Implementation status

The current repository contains a Go CLI, configuration validation, deterministic plan skeleton, provider interface, baseline policy, workspace helper, and unit tests. It does not yet contain a production agent loop, live LLM integrations, authenticated web service, persistent audit database, or a GUI. These absences are material and should not be hidden behind broad feature claims.

## 6. Evaluation methodology

### 6.1 Baselines

- **B0 — Prompt-only:** tool eligibility is described in the agent prompt; no independent policy gate.
- **B1 — ORBITA policy:** every proposed operation is evaluated by the policy engine before tool dispatch.
- **B2 — Policy plus approval:** B1 plus explicit operator confirmation for every state-changing operation.

### 6.2 Workloads

Create a versioned suite of benign and adversarial scenarios: workspace read, workspace write, unknown tool request, missing authorization, path traversal attempts, symlink escape attempts, prompt-injection text in documents, conflicting instructions, and multi-step tasks mixing read-only and mutating operations. Each scenario should have a pre-registered expected decision and reproducible fixture.

### 6.3 Metrics

- Unauthorized execution rate (primary security outcome).
- False denial and false approval rates.
- Task success rate against a task-specific oracle.
- End-to-end latency and policy-decision latency (median, p95).
- Operator interventions per task and approval turnaround time.
- Resource consumption and binary size per target.
- Cross-platform test pass rate.

Report confidence intervals and raw per-trial outcomes. Do not collapse security and utility into one score without disclosing weights.

### 6.4 Experimental controls

Pin model versions, prompts, tool schemas, dependencies, OS images, and random seeds where supported. Run each scenario across multiple trials and report failures, not just successful cases. Separate development fixtures from held-out evaluation cases. Have an independent reviewer label expected authorization outcomes. Publish the scripts, environment manifests, anonymized traces, and artifact checksums.

### 6.5 Ablations

Remove the external policy gate, remove human review, disable explicit authorization, and replace structured tool identifiers with free-form names. Compare each ablation to the complete configuration to estimate the contribution of each control.

## 7. Results

**No empirical results are reported in this draft.** The evaluation has not been run. This section must be populated only after the protocol is executed. Include complete tables, confidence intervals, failure analysis, platform-specific build/test results, and links to raw data. Do not insert expected values as though measured.

## 8. Discussion

A policy gate can create a clearer control boundary, but its effectiveness depends on complete mediation: every effectful tool must pass through it. A single unregistered adapter or alternate execution path can invalidate the model. Human review improves oversight but may produce fatigue or approval habituation. The design therefore needs risk-based review, clear operation previews, short-lived scoped approvals, and auditability. Portability at compile time does not establish identical behavior across operating systems.

## 9. Limitations and future work

The current system is a small prototype and not a mature agent platform. Planned work includes robust filesystem containment, cryptographically scoped approvals, append-only audit logs, provider adapters, a typed tool registry, sandboxed execution, resource budgets, replayable traces, native cross-platform integration testing, and user studies of operator review. A formal policy specification and independent adversarial evaluation are also needed.

## 10. Conclusion

ORBITA presents a research direction for tool-using agents in which models propose plans but do not independently authorize effects. The prototype demonstrates a small fail-closed policy core and a portable Go foundation. Whether the architecture meaningfully improves security without unacceptable utility cost remains an empirical question. The proposed evaluation is designed to answer that question transparently and reproducibly.

## References

[1] OWASP Foundation, *OWASP Top 10 for Large Language Model Applications*, 2025 edition. https://genai.owasp.org/llm-top-10/

[2] NIST, *Artificial Intelligence Risk Management Framework (AI RMF 1.0)*, NIST AI 100-1, 2023. https://doi.org/10.6028/NIST.AI.100-1

[3] NIST, *Generative Artificial Intelligence Profile*, NIST AI 600-1, 2024. https://doi.org/10.6028/NIST.AI.600-1

[4] MITRE, *ATLAS: Adversarial Threat Landscape for Artificial-Intelligence Systems*. https://atlas.mitre.org/

[5] Go Authors, *Go Documentation: Compile and install applications*. https://go.dev/doc/

[6] Open Worldwide Application Security Project, *Application Security Verification Standard*. https://owasp.org/www-project-application-security-verification-standard/

---

## Reproducibility checklist

- [ ] Freeze the exact commit and dependency set.
- [ ] Publish scenario fixtures and expected decisions.
- [ ] Publish evaluation scripts and environment manifests.
- [ ] Run tests on native Windows, macOS, and Linux workers.
- [ ] Record model/provider identifiers and prompt versions.
- [ ] Report all trials and negative results.
- [ ] Archive raw data and release artifact checksums.
- [ ] Obtain independent review of the threat model and results.
