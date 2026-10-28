package crosswritecanary

import "testing"

func TestLiveSumDBMarker(t *testing.T) {
	if got := LiveSumDBMarker(); got != "owned-live-sumdb-v0.0.6" {
		t.Fatalf("marker = %q", got)
	}
}
