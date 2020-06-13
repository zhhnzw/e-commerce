package models

import (
	"fmt"
	"goods/pb"
	"time"
)

// 下面的模型是老版本迁移留下的，目前没用了，实际用的pb模块的模型
type Commodity struct {
	Id              int       `json:"-"`
	CommodityUuid   string    `json:"commodityUuid" form:"commodityUuid"`
	CommodityFrom   string    `json:"commodityFrom" form:"commodityFrom"`
	CommodityTypeId int       `json:"commodityTypeId" gorm:"commodity_type_id"`
	PrimaryType     string    `json:"primaryType" form:"primaryType"`
	SecondaryType   string    `json:"secondaryType" form:"secondaryType"`
	Img             string    `json:"img"`
	Imgs            string    `json:"imgs"`
	IsValid         bool      `json:"isValid" gorm:"default:true,column:is_valid" form:"isValid"`
	Title           string    `json:"title"`
	SubTitle        string    `json:"subTitle"`
	Price           int       `json:"price"`
	PublishDate     JSONTime  `json:"publishDate" form:"-" gorm:"-"`
	CreatedTime     time.Time `json:"-" form:"-" gorm:"-"`
	UpdatedTime     JSONTime  `json:"-" form:"-" gorm:"-"`
	DeletedTime     time.Time `json:"-" form:"-" gorm:"-"`
	PageSize        int       `gorm:"-" json:"-" form:"pageSize"`
	PageIndex       int       `gorm:"-" json:"-" form:"pageIndex"`
}

func (*Commodity) TableName() string {
	return "tb_commodity"
}

type CommodityType struct {
	Id            int       `json:"-"`
	PrimaryType   string    `json:"primaryType"`
	SecondaryType string    `json:"secondaryType"`
	CreatedTime   time.Time `json:"-" form:"-" gorm:"-"`
	UpdatedTime   JSONTime  `json:"updateTime" form:"-" gorm:"-"`
	DeletedTime   time.Time `json:"-" form:"-" gorm:"-"`
	PageSize      int       `gorm:"-" json:"-" form:"pageSize"`
	PageIndex     int       `gorm:"-" json:"-" form:"pageIndex"`
}

func (*CommodityType) TableName() string {
	return "tb_commodity_type"
}

/*
通过商品类型筛选商品
查询优化: 延迟关联,覆盖索引.内连接查询已经从索引上取得了需要的主键，只会对pageSize个主键关联原表查找，减少了mysql扫描那些需要丢弃的行
SELECT t1.commodity_uuid,t1.commodity_from,t1.commodity_type_id,t1.primary_type,t1.secondary_type,t1.publish_date,t1.price,t1.title,t1.sub_title
FROM tb_commodity AS t1 INNER JOIN (
	SELECT id FROM tb_commodity WHERE primary_type='pants' AND secondary_type='casual_pants' ORDER BY id DESC LIMIT 10000,20
) AS t2
ON t1.id=t2.id
*/
func GetCommoditiesByType(request *pb.GoodsRequest) (*pb.GoodsReply, error) {
	results := make([]*pb.GoodsReplyItem, 0, request.PageSize)
	sql := `SELECT t1.commodity_uuid,t1.commodity_from,t1.img,t1.title,t1.sub_title,t1.price,t1.publish_date,t1.primary_type,t1.secondary_type,t1.is_valid,t1.commodity_type_id FROM tb_commodity AS t1 INNER JOIN (
	SELECT id FROM tb_commodity WHERE primary_type='%s' AND secondary_type='%s' ORDER BY id DESC LIMIT %d OFFSET %d) AS t2 ON t1.id=t2.id`
	sql = fmt.Sprintf(sql, request.PrimaryType, request.SecondaryType, request.PageSize, (request.PageIndex-1)*request.PageSize)
	db := DB.Raw(sql).Find(&results)
	reply := &pb.GoodsReply{Data: results}
	return reply, db.Error
}

// 根据uuid查找商品详情记录
// 给commodity_uuid添加了唯一约束, 根据commodity_uuid来查找，mysql可直接定位到这条记录，无需再优化了
func GetCommodity(request *pb.GoodsRequest) (*pb.GoodsReplyItem, error) {
	reply := &pb.GoodsReplyItem{}
	queryFields := []string{"commodity_uuid", "commodity_from", "img", "title", "sub_title", "price", "publish_date", "primary_type", "secondary_type", "commodity_type_id", "is_valid", "imgs"}
	db := DB.Table("tb_commodity").Select(queryFields).Where("commodity_uuid=?", request.GoodsUuid).First(reply)
	return reply, db.Error
}