//go:build !linux && !windows && !freebsd

package osl

// GC triggers garbage collection of namespace path right away
// and waits for it.
func GC() {
}

// GetSandboxForExternalKey returns sandbox object for the supplied path
func GetSandboxForExternalKey(path string, key string) (Sandbox, error) {
	return nil, nil
}
