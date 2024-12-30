package usecase

import (
	"context"
	"github.com/KhaiHust/authen_service/core/entity"
	"github.com/KhaiHust/authen_service/core/exception"
	"github.com/KhaiHust/authen_service/core/port"
	"github.com/golibs-starter/golib/log"
)

type ICreateCategoryUsecase interface {
	CreateCategory(ctx context.Context, category *entity.CategoryEntity) (*entity.CategoryEntity, error)
}
type CreateCategoryUsecase struct {
	categoryPort               port.ICategoryPort
	databaseTransactionUsecase IDatabaseTransactionUsecase
}

func (c CreateCategoryUsecase) CreateCategory(ctx context.Context, category *entity.CategoryEntity) (*entity.CategoryEntity, error) {
	tx := c.databaseTransactionUsecase.StartTransaction()
	var err error
	defer func() {
		if r := recover(); r != nil {
			err = exception.InternalServerErrorException
		}
		if errRollback := tx.Rollback(); errRollback != nil {
			log.Error(ctx, "Rollback transaction failed: ", errRollback)
		} else {
			log.Info(ctx, "Rollback transaction success")
		}
	}()
	category, err = c.categoryPort.CreateNewCategory(ctx, tx, category)
	if err != nil {
		log.Error(ctx, "Create category failed: ", err)
		return nil, err
	}
	if err = tx.Commit().Error; err != nil {
		log.Error(ctx, "Commit transaction failed: ", err)
		return nil, err
	}
	return category, nil
}

func NewCreateCategoryUsecase(categoryPort port.ICategoryPort, databaseTransactionUsecase IDatabaseTransactionUsecase) ICreateCategoryUsecase {
	return &CreateCategoryUsecase{
		categoryPort:               categoryPort,
		databaseTransactionUsecase: databaseTransactionUsecase,
	}
}
