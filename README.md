# Go Byteman cross-write canary

This public repository is a researcher-owned target used to verify whether one
Google package-analysis worker can alter an analysis record attributed to a
different package owned by the same researcher.

It contains no workflow, callback, credential, generated binary, or runtime
side effect. Its exported helper constructs an `os/exec.Cmd` but never runs it.
That call deliberately gives the canonical `v0.0.8` tag a stable, non-empty
Capslock result for an attributable, owner-only suppression and restoration
test.
