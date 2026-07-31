// Package fixture is a minimal module used to smoke-test the reusable workflows.
package fixture

import "github.com/google/uuid"

// Add returns the sum of a and b.
func Add(a, b int) int {
	return a + b
}

// ID returns a new random identifier. It exists so the module has a real
// dependency, which gives it a go.sum.
func ID() string {
	return uuid.NewString()
}
