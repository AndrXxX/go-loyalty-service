package entities

type Withdraw struct {
	Order     string      `json:"order"`
	Sum       *float64    `json:"sum"`
	CreatedAt RFC3339Time `json:"processed_at"`
}
