// Package crosswritecanary is a researcher-owned target for validating
// authorization boundaries between independent package records.
package crosswritecanary

import "os/exec"

// Baseline identifies the canonical contents published by the owner.
const Baseline = "robert-owned cross-package baseline v0.0.8"

// Command returns an inert command value. It is never run. The direct call to
// os/exec is intentional so Capslock produces a stable, non-empty capability
// record that a later owned-package test can attempt to suppress.
func Command() *exec.Cmd {
	return exec.Command("true")
}
