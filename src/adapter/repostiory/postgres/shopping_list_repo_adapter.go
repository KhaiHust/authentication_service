package postgres

import (
	"github.com/KhaiHust/authen_service/core/port"
	"gorm.io/gorm"
)

type ShoppingListRepoAdapter struct {
	base
}

func NewShoppingListRepoAdapter(db *gorm.DB) port.IShoppingListPort {
	return &ShoppingListRepoAdapter{
		base: base{db: db},
	}
}
