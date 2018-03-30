package elastics

import (
	"gopkg.in/olivere/elastic.v5"
	"time"
)

type Goods struct {
	Id             string      `json:"-"`
	GoodsId        int         `json:"goods_id"`
	GoodsName      string      `json:"goods_name"`
	GoodsBrief     string      `json:"goods_brief"`
	CatePid        int         `json:"cate_pid"`
	CateId         int         `json:"cate_id"`
	BarCode        string      `json:"bar_code"`
	BrandId        int         `json:"brand_id"`
	BrandName      string      `json:"brand_name"`
	ShopId         int         `json:"shop_id"`
	ShopName       string      `json:"shop_name"`
	GoodsStock     int         `json:"goods_stock"`
	GoodsUnit      string      `json:"goods_unit"`
	GoodsSpec      string      `json:"goods_spec"`
	Price          int         `json:"price"`
	VipPrice       int         `json:"vip_price"`
	GoodsImg       string      `json:"goods_img"`
	ImgPathUrl     string      `json:"img_path_url"`
	GoodsParameter string      `json:"goods_parameter"`
	State          int         `json:"state"`
	SaleNum        int         `json:"sale_num"`
	ViewNum        int         `json:"view_num"`
	UpdatedAt      int         `json:"updated_at"`
	CreatedAt      int         `json:"created_at"`
	DeletedAt      time.Time   `json:"deleted_at"`
	Snapshot       interface{} `json:"_"`
}

func (g *Goods) Index() string {
	return "goods_test"
}

func (g *Goods) Type() string {
	return "goods"
}

func (g *Goods) AfterFind() {
	g.Snapshot = *g
}

func (g *Goods) SetId(id string) {
	g.Id = id
}

func (g *Goods) GetSnapshot() interface{} {
	return g.Snapshot
}

func (g *Goods) Mapping() map[string]map[string]map[string]map[string]map[string]string {
	return map[string]map[string]map[string]map[string]map[string]string{
		"mappings": {
			g.Type(): {
				"properties": {
					"goods_id":        {"type": "long"},
					"goods_name":      {"type": "string"},
					"goods_brief":     {"type": "string"},
					"cate_pid":        {"type": "string"},
					"cate_id":         {"type": "string"},
					"bar_code":        {"type": "string"},
					"brand_id":        {"type": "long"},
					"brand_name":      {"type": "string"},
					"shop_id":         {"type": "long"},
					"shop_name":       {"type": "string"},
					"goods_stock":     {"type": "long"},
					"goods_unit":      {"type": "string"},
					"goods_spec":      {"type": "string"},
					"price":           {"type": "long"},
					"vip_price":       {"type": "long"},
					"goods_img":       {"type": "string"},
					"img_path_url":    {"type": "string"},
					"goods_parameter": {"type": "string"},
					"state":           {"type": "long"},
					"sale_num":        {"type": "long"},
					"view_num":        {"type": "long"},
					"created_at":      {"type": "date"},
					"updated_at":      {"type": "date"},
					"deleted_at":      {"type": "date"},
				},
			},
		},
	}
}

func CreateGoodsIndex() {
	CreateIndex(new(Goods))
}

func DelGoodsIndex() {
	DelIndex(new(Goods))
}

func AddGoods(model *Goods) error {
	return AddDocument(model)
}

func DelGoods(id string) error {
	return DelDocument(new(Goods), id)
}

func UpdateGoodsById(goods *Goods) error {
	if goods.Id == "" {

		return ErrNotExist

	}

	return UpdateDocumentById(goods, goods.Id)
}

func GetGoodsByID(id string) (*Goods, error) {
	goods := new(Goods)
	return goods, GetDocumentById(goods, id)
}

func GetGoodsList(query elastic.Query, sort map[string]bool, limit ...int) ([]*ElasticSearch, error) {
	return GetDocuments(new(Goods), query, sort, limit...)
}
