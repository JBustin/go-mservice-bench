package account

type Account struct {
	Id       uint    `json:"id" gorm:"primary_key"`
	ClientId uint    `json:"clientId"`
	Amount   float64 `json:"amount"`
}

type CreateAccountInput struct {
	ClientId uint    `json:"clientId" binding:"required"`
	Amount   float64 `json:"amount" binding:"required"`
}

type UpdateAccountInput struct {
	Id       uint    `json:"id"`
	ClientId uint    `json:"clientId"`
	Amount   float64 `json:"amount"`
}
