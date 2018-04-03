package models

import (
	"github.com/go-xorm/xorm"
	"time"
)

type Goods struct {
	GoodsId        int       `json:"goods_id" xorm:"not null pk autoincr INT(11)"`
	GoodsName      string    `json:"goods_name" xorm:"not null default '' VARCHAR(120)"`
	GoodsBrief     string    `json:"goods_brief" xorm:"not null default '' VARCHAR(60)"`
	CatePid        int       `json:"cate_pid" xorm:"not null default 0 SMALLINT(6)"`
	CateId         int       `json:"cate_id" xorm:"not null default 0 SMALLINT(6)"`
	BarCode        string    `json:"bar_code" xorm:"not null default '' VARCHAR(60)"`
	BrandId        int       `json:"brand_id" xorm:"not null default 0 INT(11)"`
	BrandName      string    `json:"brand_name" xorm:"not null default '' VARCHAR(60)"`
	ShopId         int       `json:"shop_id" xorm:"not null default 0 INT(11)"`
	ShopName       string    `json:"shop_name" xorm:"not null default '' VARCHAR(60)"`
	GoodsStock     int       `json:"goods_stock" xorm:"not null default 0 SMALLINT(6)"`
	GoodsUnit      string    `json:"goods_unit" xorm:"not null default '' VARCHAR(15)"`
	GoodsSpec      string    `json:"goods_spec" xorm:"not null default '' VARCHAR(30)"`
	Price          int       `json:"price" xorm:"not null default 0 INT(10)"`
	VipPrice       int       `json:"vip_price" xorm:"not null default 0 INT(10)"`
	GoodsImg       string    `json:"goods_img" xorm:"not null default '' VARCHAR(255)"`
	ImgPathUrl     string    `json:"img_path_url" xorm:"not null default '' VARCHAR(255)"`
	GoodsParameter string    `json:"goods_parameter" xorm:"not null default '' VARCHAR(255)"`
	State          int       `json:"state" xorm:"not null default 1 TINYINT(1)"`
	SaleNum        int       `json:"sale_num" xorm:"not null default 0 SMALLINT(6)"`
	ViewNum        int       `json:"view_num" xorm:"not null default 0 INT(10)"`
	UpdatedAt      int       `json:"updated_at" xorm:"not null default 0 INT(11)"`
	CreatedAt      int       `json:"created_at" xorm:"not null default 0 INT(11)"`
	DeletedAt      time.Time `json:"deleted_at" xorm:"not null DATETIME"`
}

func (m Goods) TableName() string {
	return "ty_goods"
}

func (m *Goods) AfterFind() {
}

func AddGoods(model *Goods) error {
	return AddModel(model)
}

func DelGoodsById(id int64, safe ...bool) error {
	model := new(Goods)
	return DelModel(model, Engine.ID(id), safe...)
}

func DelGoods(session *xorm.Session, safe ...bool) error {
	model := new(Goods)
	return DelModel(model, session, safe...)
}

func UpdateGoodsById(id int64, model *Goods) error {
	return UpdateModel(model, Engine.ID(id))
}

func UpdateGoods(model *Goods, session *xorm.Session) error {
	return UpdateModel(model, session)
}

func GetGoodsByID(id int64) (*Goods, error) {
	model := new(Goods)

	return model, GetModel(model, Engine.ID(id))
}

func GetGoods(session *xorm.Session) (*Goods, error) {
	model := new(Goods)

	return model, GetModel(model, session)
}

func GetGoodsList(session *xorm.Session, limit ...int) (models []*Goods, err error) {
	if len(limit) > 0 {
		models = make([]*Goods, 0, limit[0])

		err = session.Limit(limit[0]).Find(&models)

	} else {
		models = make([]*Goods, 0)

		err = session.Find(&models)

	}

	if err != nil {
		return nil, err
	}

	for i := range models {
		models[i].AfterFind()
	}

	return models, nil
}

func GetGoodsListForPage(session *xorm.Session, page int, pageSize int) ([]*Goods, error) {
	models := make([]*Goods, 0, pageSize)

	err := session.Limit(pageSize, (page-1)*pageSize).Find(&models)

	if err != nil {
		return nil, err
	}

	for i := range models {
		models[i].AfterFind()
	}

	return models, nil
}
