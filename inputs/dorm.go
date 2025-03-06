package inputs

type DormInput struct {
	DormName           string `json:"dorm_name"`
	DormPriceAmount    string `json:"dorm_price_amount"`
	DormQuantityAmount string `json:"dorm_quantity_amount"`
	DormAddress        string `json:"dorm_address"`
	DormDescription    string `json:"dorm_description"`
	DormPhoneNumber    string `json:"dorm_phone_number"`
}
