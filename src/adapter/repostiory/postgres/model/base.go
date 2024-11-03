package model

type BaseModel struct {
	ID        int64 `gorm:"column:id;primary_key"`
	CreatedAt int64 `gorm:"column:created_at" sql:"default:current_timestamp"`
	UpdatedAt int64 `gorm:"column:updated_at" sql:"default:current_timestamp"`
}
