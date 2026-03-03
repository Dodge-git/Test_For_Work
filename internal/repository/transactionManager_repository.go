package repository

import "gorm.io/gorm"

type TransactionManager interface{
	NewTransaction(fn func() error)error
}

type transactionManager struct {
	db *gorm.DB
}

func NewTransactionManager(db *gorm.DB) TransactionManager{
	return &transactionManager{db:db}
}

func (t *transactionManager) NewTransaction(fn func() error)error{
	return t.db.Transaction(func (tx *gorm.DB)error  {
	return fn()	
	})
}