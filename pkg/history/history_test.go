package history

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	name = "Nik"
)

var (
	unreliablePersons = []string{"John", "Nik", "Fill"}
)

func TestReputation_CheckHistory(t *testing.T) {
	reputation := NewCreditReputation(unreliablePersons)
	status := reputation.CheckHistory(name)
	assert.Equal(t, false, status)
	status = reputation.CheckHistory("Bob")
	assert.Equal(t, true, status)
	status = reputation.CheckHistory("John")
	assert.Equal(t, true, status)
	status = reputation.CheckHistory("")
	assert.Equal(t, true, status)

	reputation = NewCreditReputation([]string{"Bob"})
	status = reputation.CheckHistory(name)
	assert.Equal(t, true, status)
	status = reputation.CheckHistory("Bob")
	assert.Equal(t, false, status)
}
