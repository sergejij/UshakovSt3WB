package models

// Status ...
type Status = string

// Status states
const (
	CreditApproved 	Status = "The credit was approved."
	CreditReject 	Status = "The credit was reject."
	CurrentBalance	Status = "Current balance"
)
