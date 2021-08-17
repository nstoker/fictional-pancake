package version

import (
	"fmt"
	"testing"
)

func TestVersion(t *testing.T) {
	got := Version()
	expected := fmt.Sprintf("v%s", version)
	if got != expected {
		t.Errorf("Got %s, wanted %s", got, expected)
	}
}
