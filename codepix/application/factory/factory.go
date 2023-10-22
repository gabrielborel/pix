package factory

import (
	"github.com/gabrielborel/pix/codepix/application/usecase"
	"github.com/gabrielborel/pix/codepix/infra/repository"
	"github.com/jinzhu/gorm"
)

func TransactionUseCaseFactory(db *gorm.DB) *usecase.TransactionUseCase {
	pixRepository := repository.NewPixKeyRepositoryDb(db)
	transactionRepository := repository.NewTransactionRepositoryDb(db)

	transactionUseCase := usecase.NewTransactionUseCase(transactionRepository, pixRepository)
	return transactionUseCase
}
