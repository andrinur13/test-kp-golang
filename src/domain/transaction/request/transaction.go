package request

type TransactionRequest struct {
	ProducId int `json:"product_id"`
	Tenor    int `json:"tenor"`
}
