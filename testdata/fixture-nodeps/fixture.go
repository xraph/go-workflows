// Package fixture is a minimal, dependency-free module used to smoke-test the
// reusable workflows against a module that has no go.sum.
package fixture

// Add returns the sum of a and b.
func Add(a, b int) int {
	return a + b
}
