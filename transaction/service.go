package transaction

import (
	"crypto/rand"
	"test_aac/item"
	"time"
)

type service struct {
	repository     Repository
	ItemRepository item.Reprository
}

type Service interface {
	GetAllTransaction() ([]Transaction, error)
	CreateTrans(input InputCreateTrans) (Transaction, error)
	GetTransactionbyNumber(number string) (Transaction, error)
	GetTransactionbyID(ID int) (Transaction, error)
	DeleteTrans(id InputGetTransactionID) (Transaction, error)
	GetTransactionbyDate(date string) ([]Transaction, error)
}

func NewService(repository Repository, ItemRepository item.Reprository) *service {
	return &service{repository, ItemRepository}
}

func (s *service) GetAllTransaction() ([]Transaction, error) {
	transaction, err := s.repository.GetAllTransaction()
	if err != nil {
		return transaction, err
	}
	return transaction, nil
}

func (s *service) GetTransactionbyNumber(number string) (Transaction, error) {
	transaction, err := s.repository.GetTransactionbyNumber(number)
	if err != nil {
		return transaction, err
	}
	return transaction, nil
}

func (s *service) GetTransactionbyDate(date string) ([]Transaction, error) {
	transaction, err := s.repository.GetTransactionbyDate(date)
	if err != nil {
		return transaction, err
	}
	return transaction, nil
}

func (s *service) GetTransactionbyID(id int) (Transaction, error) {
	transaction, err := s.repository.GetTransactionbyID(id)
	if err != nil {
		return transaction, err
	}
	return transaction, nil
}

func (s *service) CreateTrans(input InputCreateTrans) (Transaction, error) {
	priceTotal := 0
	costTotal := 0

	data := input.DetailData
	for _, req := range data {
		dt, err := s.repository.GetItem(req.ItemId)
		if err != nil {
			return Transaction{}, err
		}
		quantity := int(req.ItemQuantity)
		price := dt.Price
		priceTotal = (quantity * price) + priceTotal

		cost := dt.Cost
		costTotal = (quantity * cost) + costTotal
	}

	date := time.Now()

	Transactions := Transaction{}
	number, _ := rand.Prime(rand.Reader, 64)
	Transactions.Date = time.Now()
	Transactions.Number = "" + date.Format("200601") + "-" + number.String()
	Transactions.PriceTotal = priceTotal
	Transactions.CostTotal = costTotal

	trans, err := s.repository.SaveTrans(Transactions)
	if err != nil {
		return trans, err
	}

	detailTrans := DetailTransaction{}

	detail_data := input.DetailData
	for _, req := range detail_data {
		detailTrans.TransactionId = trans.ID
		detailTrans.ItemId = req.ItemId
		detailTrans.ItemQuantity = req.ItemQuantity

		dt, err := s.repository.GetItem(req.ItemId)
		if err != nil {
			return Transaction{}, err
		}
		detailTrans.ItemCost = dt.Cost
		detailTrans.ItemPrice = dt.Price
		detailTrans.ItemQuantity = req.ItemQuantity

		_, err_response := s.repository.SaveDetailTrans(detailTrans)
		if err != nil {
			return Transaction{}, err_response
		}

	}

	return trans, nil
}

func (s *service) DeleteTrans(id InputGetTransactionID) (Transaction, error) {
	trans, err := s.repository.DeleteTrans(id.ID)
	if err != nil {
		return trans, err
	}
	_, err = s.repository.DetailTransDelete(id.ID)
	if err != nil {
		return trans, err
	}
	return trans, nil
}
