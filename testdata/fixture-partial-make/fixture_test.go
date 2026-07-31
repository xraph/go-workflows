package fixture

import "testing"

func TestAdd(t *testing.T) {
	if got := Add(2, 3); got != 5 {
		t.Fatalf("Add(2, 3) = %d, want 5", got)
	}
}

func TestID(t *testing.T) {
	if ID() == "" {
		t.Fatal("ID() returned an empty string")
	}
}
