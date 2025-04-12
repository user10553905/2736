package models

import "time"

// 购物车结算表
type CartCheckouts struct {
	Id           uint32    `gorm:"column:id;type:int UNSIGNED;comment:结算记录ID;primaryKey;not null;" json:"id"`          // 结算记录ID
	CreatedAt    time.Time `gorm:"column:created_at;type:datetime(3);comment:创建时间;default:NULL;" json:"created_at"`    // 创建时间
	UpdatedAt    time.Time `gorm:"column:updated_at;type:datetime(3);comment:更新/修改时间;default:NULL;" json:"updated_at"` // 更新/修改时间
	DeletedAt    time.Time `gorm:"column:deleted_at;type:datetime(3);comment:删除时间;default:NULL;" json:"deleted_at"`    // 删除时间
	CartId       int32     `gorm:"column:cart_id;type:int;comment:购物车ID;not null;" json:"cart_id"`                     // 购物车ID
	TotalAmount  int32     `gorm:"column:total_amount;type:int;comment:结算总金额;not null;" json:"total_amount"`           // 结算总金额
	CheckoutTime string    `gorm:"column:checkout_time;type:varchar(10);comment:结算时间;not null;" json:"checkout_time"`  // 结算时间
}
