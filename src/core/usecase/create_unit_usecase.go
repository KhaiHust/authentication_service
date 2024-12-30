package usecase

import (
	"context"
	"github.com/KhaiHust/authen_service/core/entity"
	"github.com/KhaiHust/authen_service/core/exception"
	"github.com/KhaiHust/authen_service/core/port"
	"github.com/golibs-starter/golib/log"
)

type ICreateUnitUseCase interface {
	CreateNewUnit(ctx context.Context, unit *entity.UnitEntity) (*entity.UnitEntity, error)
}
type CreateUnitUseCase struct {
	unitPort                   port.IUnitPort
	databaseTransactionUsecase IDatabaseTransactionUsecase
}

func (c CreateUnitUseCase) CreateNewUnit(ctx context.Context, unit *entity.UnitEntity) (*entity.UnitEntity, error) {
	tx := c.databaseTransactionUsecase.StartTransaction()
	var err error
	defer func() {
		if r := recover(); r != nil {
			err = exception.InternalServerErrorException
		}
		if errRollback := tx.Rollback(); errRollback != nil {
			log.Error(ctx, "Rollback transaction failed: ", errRollback)
		} else {
			log.Info(ctx, "Rollback transaction successfully")
		}
	}()
	unit, err = c.unitPort.SaveUnit(ctx, tx, unit)
	if err != nil {
		log.Error(ctx, "Save new unit failed: ", err)
		return nil, err
	}
	if errCommit := tx.Commit().Error; errCommit != nil {
		log.Error(ctx, "Commit transaction failed: ", errCommit)
		return nil, errCommit
	}
	return unit, nil
}

func NewCreateUnitUseCase(unitPort port.IUnitPort, databaseTransactionUsecase IDatabaseTransactionUsecase) ICreateUnitUseCase {
	return &CreateUnitUseCase{
		unitPort:                   unitPort,
		databaseTransactionUsecase: databaseTransactionUsecase,
	}
}
