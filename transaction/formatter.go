package transaction

import "time"

type TransactionFormatter struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Amount    int       `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
}

func FormatTransaction(transaction Transaction) TransactionFormatter {
	formatter := TransactionFormatter{
		ID:        transaction.ID,
		Name:      transaction.User.Name,
		Amount:    transaction.Amount,
		CreatedAt: transaction.CreatedAt,
	}
	return formatter
}

func FormatTransactions(transactions []Transaction) []TransactionFormatter {
	if len(transactions) == 0 {
		return []TransactionFormatter{}
	}

	var transactionsFormatter []TransactionFormatter

	for _, transaction := range transactions {
		formatter := FormatTransaction(transaction)
		transactionsFormatter = append(transactionsFormatter, formatter)
	}

	return transactionsFormatter
}
