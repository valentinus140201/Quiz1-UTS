package Controllers

type Wallet struct {
	WalletId    int         `form:"walletId" json:"walletId"`
	Currency    string      `form:"currency" json:"currency"`
	Username    string      `form:"username" json:"username"`
	Password    string      `form:"password" json:"password"`
	DisableUser int         `form:"disableUser" json:"disableUser"`
	Transaction Transaction `form:"transaction" json:"transaction"`
}

type WalletResponse struct {
	Status  int      `form:"status" json:"status"`
	Message string   `form:"message" json:"message"`
	Data    []Wallet `form:"data" json:"data"`
}

type Transaction struct {
	TransactionId int    `form:"transactionId" json:"transactionId"`
	IdWallet      int    `form:"idWallet" json:"idWallet"`
	DateTime      string `form:"dateTime" json:"dateTime"`
	Amount        int    `form:"amount" json:"amount"`
	Description   string `form:"description" json:"description"`
}

type TransactionResponse struct {
	Status  int           `form:"status" json:"status"`
	Message string        `form:"message" json:"message"`
	Data    []Transaction `form:"data" json:"data"`
}
