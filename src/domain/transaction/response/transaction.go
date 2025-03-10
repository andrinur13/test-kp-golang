package response

type TransactionResponse struct {
	ID                int    `json:"id"`
	AmountOtr         int    `json:"amount_otr"`
	AmountFee         int    `json:"amount_fee"`
	AmountInstallment int    `json:"amount_installment"`
	AmountInterest    int    `json:"amount_interest"`
	CreatedAt         string `json:"created_at"`
}
