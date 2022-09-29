package transaction

import "time"

type TransactionFormatter struct {
	Number     string    `json: "number"`
	Date       time.Time `json: "date"`
	PriceTotal int       `json: "price_total"`
	CostTotal  int       `json: "cost_total"`
}

type DetailTransactionFormatter struct {
	Number            string    `json: "number"`
	Date              time.Time `json: "date"`
	PriceTotal        int       `json: "price_total"`
	CostTotal         int       `json: "cost_total"`
	TransactionDetail []TransactionDetailFormatter
}

func FormatTransaction(transaction Transaction) TransactionFormatter {
	TransactionFormatter := TransactionFormatter{}
	TransactionFormatter.Number = transaction.Number
	TransactionFormatter.Date = transaction.Date
	TransactionFormatter.PriceTotal = transaction.PriceTotal
	TransactionFormatter.CostTotal = transaction.CostTotal

	return TransactionFormatter
}

func FormatListTrans(transaction []Transaction) []TransactionFormatter {
	if len(transaction) == 0 {
		return []TransactionFormatter{}
	}

	var transactionFormatter []TransactionFormatter

	for _, transactions := range transaction {
		formatter := FormatTransaction(transactions)
		transactionFormatter = append(transactionFormatter, formatter)
	}

	return transactionFormatter
}

type TransactionDetailFormatter struct {
	ItemId       string `json: "item_id"`
	ItemQuantity int    `json: "item_quantity"`
	ItemPrice    int    `json: "item_price"`
	ItemCost     int    `json: "item_cost"`
}

func FormatDetailTransaction(transaction Transaction) DetailTransactionFormatter {
	TransactionFormatter := DetailTransactionFormatter{}
	TransactionFormatter.Date = transaction.Date
	TransactionFormatter.Number = transaction.Number
	TransactionFormatter.PriceTotal = transaction.PriceTotal
	TransactionFormatter.CostTotal = transaction.CostTotal
	details := []TransactionDetailFormatter{}

	for _, detail := range transaction.DetailTransaction {
		TransactionDetailFormatter := TransactionDetailFormatter{}
		TransactionDetailFormatter.ItemId = detail.ItemId
		TransactionDetailFormatter.ItemQuantity = detail.ItemQuantity
		TransactionDetailFormatter.ItemPrice = detail.ItemPrice
		TransactionDetailFormatter.ItemCost = detail.ItemCost

		details = append(details, TransactionDetailFormatter)
	}

	TransactionFormatter.TransactionDetail = details

	return TransactionFormatter
}
