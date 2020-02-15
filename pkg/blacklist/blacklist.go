package blacklist

type Blocked interface {
	Check(name string) bool
}

type blocked struct {
	persons []string
}

func (b *blocked) Check(name string) bool {
	for i := 0; i < len(b.persons); i++ {
		if b.persons[i] == name {
			return false
		}
	}
	return true
}

// NewBlocked initializes the Blocked.
func NewBlocked(persons []string) Blocked {
	return &blocked{
		persons: persons,
	}
}
