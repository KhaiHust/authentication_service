package usecase

import (
	"github.com/KhaiHust/authen_service/core/port"
	"gorm.io/gorm"
)

type IDatabaseTransactionUsecase interface {
	StartTransaction() *gorm.DB
	Commit(txDB *gorm.DB) error
	Rollback(txDB *gorm.DB) error
}
type DatabaseTransactionUsecase struct {
	databaseTransactionPort port.IDatabaseTransactionPort
}

func (d *DatabaseTransactionUsecase) StartTransaction() *gorm.DB {
	return d.databaseTransactionPort.StartTransaction()
}

func (d *DatabaseTransactionUsecase) Commit(txDB *gorm.DB) error {
	return d.databaseTransactionPort.Commit(txDB)
}

func (d *DatabaseTransactionUsecase) Rollback(txDB *gorm.DB) error {
	return d.databaseTransactionPort.Rollback(txDB)
}

func NewDatabaseTransactionUsecase(databaseTransactionPort port.IDatabaseTransactionPort) IDatabaseTransactionUsecase {
	return &DatabaseTransactionUsecase{
		databaseTransactionPort: databaseTransactionPort,
	}
}
