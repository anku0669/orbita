# Contributing

1. Open an issue describing the problem or proposal.
2. Add a focused test for behavior changes.
3. Run `gofmt -w ./cmd ./internal`, `go test ./...`, and `go vet ./...`.
4. Keep tool adapters least-privilege and route state changes through policy review.
5. Never commit credentials, real customer data, or unredacted traces.

Security reports should not include exploit payloads against systems you do not own. Use the repository's private vulnerability reporting feature when available.
