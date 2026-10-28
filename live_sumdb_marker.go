package crosswritecanary

// LiveSumDBMarker is an inert owner-published marker used to verify that a
// fresh consumer records and consumes this new release through sum.golang.org.
func LiveSumDBMarker() string {
	return "owned-live-sumdb-v0.0.6"
}
