package models

type GoodsDetail struct {
	GoodsID        int     `gorm:"column:goods_id" json:"goods_id"`
	GoodsName      string  `gorm:"column:goods_name" json:"goods_name"`
	GoodsPrice     float64 `gorm:"column:goods_price" json:"goods_price"`
	GoodsNumber    *int    `gorm:"column:goods_price" json:"goods_number"`
	GoodsWeight    *int    `gorm:"column:goods_weight" json:"goods_weight"`
	GoodsIntroduce *string `gorm:"column:goods_introduce" json:"goods_introduce"`
	GoodsDeleted   *bool   `gorm:"column:goods_deleted" json:"goods_deleted"`
	GoodsState     *int    `gorm:"column:goods_state" json:"goods_state"`
	GoodsCat       *string `gorm:"column:goods_cat" json:"goods_cat"`
	PicUrl         *string `gorm:"column:pic_url" json:"pic_url"`
}

func (GoodsDetail) TableName() string {
	return "goods_detail"
}
