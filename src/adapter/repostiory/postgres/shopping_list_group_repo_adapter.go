package postgres

import (
	"github.com/KhaiHust/authen_service/core/port"
	"gorm.io/gorm"
)

type ShoppingListGroupRepoAdapter struct {
	base
}

func NewShoppingListGroupRepoAdapter(db *gorm.DB) port.IShoppingListGroupPort {
	return &ShoppingListGroupRepoAdapter{
		base: base{db: db},
	}
}
