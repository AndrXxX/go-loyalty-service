package entities

type Order struct {
	Number    string      `json:"number"`
	Status    string      `json:"status"`
	Accrual   *float64    `json:"accrual"`
	CreatedAt RFC3339Time `json:"uploaded_at"`
}
