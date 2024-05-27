package models

type Category struct {
	ID      int     `gorm:"column:cat_id" json:"cat_id"`
	Name    string  `gorm:"column:cat_name" json:"cat_name"`
	Deleted *bool   `gorm:"column:cat_deleted" json:"cat_deleted"`
	Icon    *string `gorm:"column:cat_icon" json:"cat_icon"`
}

// TableName 指定 Category 对应的数据库表名为 "category"
func (Category) TableName() string {
	return "category"
}
