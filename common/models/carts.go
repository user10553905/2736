package models

import "time"

// 购物车表
type Carts struct {
	Id        uint32    `gorm:"column:id;type:int UNSIGNED;comment:购物车ID;primaryKey;not null;" json:"id"`           // 购物车ID
	CreatedAt time.Time `gorm:"column:created_at;type:datetime(3);comment:创建时间;default:NULL;" json:"created_at"`    // 创建时间
	UpdatedAt time.Time `gorm:"column:updated_at;type:datetime(3);comment:更新/修改时间;default:NULL;" json:"updated_at"` // 更新/修改时间
	DeletedAt time.Time `gorm:"column:deleted_at;type:datetime(3);comment:删除时间;default:NULL;" json:"deleted_at"`    // 删除时间
	IsUserId  uint32    `gorm:"column:is_user_id;type:int UNSIGNED;comment:用户ID;not null;" json:"is_user_id"`       // 用户ID
	ProductId int32     `gorm:"column:product_id;type:int;comment:商品ID;not null;" json:"product_id"`                // 商品ID
	Quantity  int32     `gorm:"column:quantity;type:int;comment:商品数量;not null;" json:"quantity"`                    // 商品数量
}
