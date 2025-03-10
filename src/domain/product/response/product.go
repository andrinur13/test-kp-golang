package response

type Product struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	AmountPrice int    `json:"amount_price"`
	AmountShip  int    `json:"amount_ship"`
}
