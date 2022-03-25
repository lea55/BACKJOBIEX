package coreconstants

const (
	PaymentStsPending   = "PD001"
	PaymentStsApproved  = "DPA002"
	PaymentStsConfirmed = "APP003"
)

const (
	PaymentMthUndefined   = "SD001"
	PaymentMthMercadoPago = "MP002"
)

type ProjectStatus struct {
	Opened     string
	InProgress string
	Closed     string
}

func NewProjectStatus() *ProjectStatus {
	return &ProjectStatus{
		Opened:     "OPND",
		InProgress: "INPG",
		Closed:     "CLSD",
	}
}
