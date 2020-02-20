package history

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	name = "Nik"
)

var (
	unreliablePersons = []string{"John", "Nik", "Fill"}
)

// TestReputation_CheckHistory ...
func TestReputation_CheckHistory(t *testing.T) {
	reputation := NewReputation(unreliablePersons)
	status := reputation.CheckHistory(name)
	assert.Equal(t, false, status)
	status = reputation.CheckHistory("Bob")
	assert.Equal(t, true, status)
	status = reputation.CheckHistory("John")
	assert.Equal(t, false, status)
	status = reputation.CheckHistory("")
	assert.Equal(t, true, status)

	reputation = NewReputation([]string{"Bob"})
	status = reputation.CheckHistory(name)
	assert.Equal(t, true, status)
	status = reputation.CheckHistory("Bob")
	assert.Equal(t, false, status)
}
