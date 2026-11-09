# ORBITA Architecture

## Design principles

1. **Provider neutrality:** model vendors are behind a small interface.
2. **Policy before effect:** every tool invocation must pass authorization and policy evaluation before execution.
3. **Human control:** state-changing operations require explicit approval.
4. **Bounded context:** filesystem tools operate under a configured workspace root.
5. **Inspectable execution:** plans and decisions are structured data, not hidden prompt state.
6. **Portable delivery:** the core runtime targets Windows, macOS, and Linux.

## Request lifecycle

1. The CLI parses a goal and loads validated configuration.
2. The planner emits a deterministic plan with explicit approval markers.
3. The tool registry resolves each requested operation. Unknown tools fail closed.
4. The policy engine checks authorization, target presence, operation class, and read-only mode.
5. A review decision is presented to an operator. An approval token must be scoped to the exact operation before execution is implemented.
6. The tool runs in a bounded execution adapter and returns structured output.
7. The runtime validates output and records an audit event. Persistent event storage is planned, not yet implemented.

## Trust boundaries

- Operator input is untrusted until parsed and validated.
- Model output is untrusted data and must never be treated as executable instructions.
- Tool adapters are privileged components and require independent review.
- Credentials belong in OS secret stores or short-lived environment injection, never source control.
- The workspace root is a security boundary only when filesystem operations use canonical paths and defend against symlinks and races. The prototype still needs this hardening.

## Planned service topology

A later deployment profile may separate the local agent runner, authenticated API, web UI, policy service, and audit store. The first release remains a local CLI to minimize exposed attack surface.
