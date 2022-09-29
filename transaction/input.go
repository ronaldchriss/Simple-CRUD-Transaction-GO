package transaction

type InputGetTransactionNumber struct {
	Number string `json:"number" binding:"required"`
}

type InputGetTransactionDate struct {
	Date string `json:"date" binding:"required"`
}

type InputGetTransactionID struct {
	ID int `uri:"id" binding:"required"`
}

type InputCreateTrans struct {
	DetailData []InputCreateTransDetail `json:"detail_item"`
}

type InputCreateTransDetail struct {
	ItemId       string `json:"item_id"`
	ItemQuantity int    `json:"item_quantity"`
}
