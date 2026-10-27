# Go Byteman cross-write canary

This public repository is an inert, researcher-owned target used to verify
whether one Google package-analysis worker can write a cached download record
attributed to a different package owned by the same researcher.

It contains no workflow, callback, credential, dependency, generated binary,
or runtime side effect. The canonical `v0.0.3` tag is reserved for a first-fetch
integrity test and does not contain the private altered-archive marker.
