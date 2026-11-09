# Packaging

The current automated release creates portable CLI binaries for supported target tuples. They are not OS-native installers.

Before creating `.dmg`, `.msi`, `.deb`, or `.rpm` packages, add platform-specific packaging definitions, signing identities, notarization, SBOM generation, and installation/uninstallation tests. Never embed API keys in distributables.
