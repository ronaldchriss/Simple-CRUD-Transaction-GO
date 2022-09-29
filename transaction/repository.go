package transaction

import (
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

type Repository interface {
	SaveDetailTrans(detail DetailTransaction) (DetailTransaction, error)
	SaveTrans(transaction Transaction) (Transaction, error)
	GetItem(id string) (Item, error)
	GetTransactionbyID(ID int) (Transaction, error)
	GetAllTransaction() ([]Transaction, error)
	GetTransactionbyNumber(number string) (Transaction, error)
	GetTransactionbyDate(date string) ([]Transaction, error)
	DeleteTrans(id int) (Transaction, error)
	DetailTransDelete(id int) (DetailTransaction, error)
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetAllTransaction() ([]Transaction, error) {
	var transaction []Transaction

	err := r.db.Order("id desc").Find(&transaction).Error
	if err != nil {
		return transaction, err
	}

	return transaction, nil
}

func (r *repository) GetTransactionbyID(ID int) (Transaction, error) {
	var transaction Transaction

	err := r.db.Preload("DetailTransaction").Where("id", ID).Find(&transaction).Error
	if err != nil {
		return transaction, err
	}

	return transaction, nil
}

func (r *repository) GetTransactionbyDate(date string) ([]Transaction, error) {
	var transaction []Transaction

	err := r.db.Preload("DetailTransaction").Where("date", date).Find(&transaction).Error
	if err != nil {
		return transaction, err
	}

	return transaction, nil
}

func (r *repository) GetTransactionbyNumber(number string) (Transaction, error) {
	var transaction Transaction

	err := r.db.Preload("DetailTransaction").Where("number", number).Find(&transaction).Error
	if err != nil {
		return transaction, err
	}

	return transaction, err
}

func (r *repository) GetItem(id string) (Item, error) {
	var item Item

	err := r.db.Where("id", id).Find(&item).Error
	if err != nil {
		return item, err
	}

	return item, err
}

func (r *repository) SaveDetailTrans(detail DetailTransaction) (DetailTransaction, error) {
	err := r.db.Create(&detail).Error
	if err != nil {
		return detail, err
	}

	return detail, nil
}

func (r *repository) SaveTrans(transaction Transaction) (Transaction, error) {
	err := r.db.Create(&transaction).Error
	if err != nil {
		return transaction, err
	}

	return transaction, nil
}

func (r *repository) DeleteTrans(id int) (Transaction, error) {
	var trans Transaction

	err := r.db.Where("id=?", id).Delete(&trans).Error

	if err != nil {
		return trans, err
	}

	return trans, nil
}

func (r *repository) DetailTransDelete(id int) (DetailTransaction, error) {
	var detail DetailTransaction

	err := r.db.Where("transaction_id=?", id).Delete(&detail).Error

	if err != nil {
		return detail, err
	}

	return detail, nil
}
