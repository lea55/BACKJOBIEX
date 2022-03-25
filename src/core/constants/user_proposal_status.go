package constants

type UserProposalStatusCode struct {
	Created     string
	Approved    string
	DevProgress string
}

func NewUserProposalStatusCode() *UserProposalStatusCode {
	return &UserProposalStatusCode{
		Created:     "CUP001",
		Approved:    "CUP002",
		DevProgress: "CUP003",
	}
}
