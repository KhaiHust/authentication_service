package postgres

import (
	"context"
	"github.com/KhaiHust/authen_service/adapter/repostiory/postgres/mapper"
	"github.com/KhaiHust/authen_service/adapter/repostiory/postgres/model"
	"github.com/KhaiHust/authen_service/core/entity"
	"github.com/KhaiHust/authen_service/core/port"
	"gorm.io/gorm"
)

type ShoppingTaskRepoAdapter struct {
	base
}

func (s *ShoppingTaskRepoAdapter) GetTaskByID(ctx context.Context, taskID int64) (*entity.ShoppingTaskEntity, error) {
	var shoppingTaskModel model.ShoppingTaskModel
	if err := s.db.WithContext(ctx).Where("id = ?", taskID).First(&shoppingTaskModel).Error; err != nil {
		return nil, err
	}
	return mapper.ToShoppingTaskEntity(&shoppingTaskModel), nil
}

func (s *ShoppingTaskRepoAdapter) UpdateTaskByID(ctx context.Context, tx *gorm.DB, task *entity.ShoppingTaskEntity) (*entity.ShoppingTaskEntity, error) {
	shoppingTaskModel := mapper.ToShoppingTaskModel(task)
	if err := tx.WithContext(ctx).
		Model(&model.ShoppingTaskModel{}).
		Where("id = ?", shoppingTaskModel.ID).
		Updates(&shoppingTaskModel).
		Scan(&shoppingTaskModel).Error; err != nil {
		return nil, err
	}
	return mapper.ToShoppingTaskEntity(shoppingTaskModel), nil
}

func (s *ShoppingTaskRepoAdapter) DeleteTaskByID(ctx context.Context, tx *gorm.DB, taskID int64) error {
	if err := tx.WithContext(ctx).Where("id = ?", taskID).Delete(&model.ShoppingTaskModel{}).Error; err != nil {
		return err
	}
	return nil
}

func (s *ShoppingTaskRepoAdapter) CreateNewShoppingTasks(ctx context.Context, tx *gorm.DB, shoppingTasks []*entity.ShoppingTaskEntity) ([]*entity.ShoppingTaskEntity, error) {
	shoppingTaskModels := mapper.ToListShoppingTaskModel(shoppingTasks)
	if err := tx.WithContext(ctx).Create(&shoppingTaskModels).Error; err != nil {
		return nil, err
	}
	return mapper.ToListShoppingTaskEntity(shoppingTaskModels), nil
}
func (s *ShoppingTaskRepoAdapter) GetShoppingTasksByShoppingListID(ctx context.Context, shoppingListID int64) ([]*entity.ShoppingTaskEntity, error) {
	var shoppingTaskModels []*model.ShoppingTaskModel
	if err := s.db.WithContext(ctx).Where("shopping_list_id = ?", shoppingListID).Find(&shoppingTaskModels).Error; err != nil {
		return nil, err
	}
	return mapper.ToListShoppingTaskEntity(shoppingTaskModels), nil
}
func NewShoppingTaskRepoAdapter(db *gorm.DB) port.IShoppingTaskPort {
	return &ShoppingTaskRepoAdapter{base{db}}
}
