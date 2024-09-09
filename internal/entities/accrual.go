package entities

type Accrual struct {
	Order   string   `json:"order"`
	Accrual *float64 `json:"accrual"`
	Status  string   `json:"status"`
}
