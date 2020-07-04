package models

import (
	"fmt"
	"log"
	"order/pb"
	"time"
)

type Order struct {
	Id            int
	GoodsUuid     string
	GoodsTypeId   int64
	PrimaryType   string
	SecondaryType string
	Img           string
	Title         string
	Subtitle      string
	Price         int64
	OrderStatus   string
	CreatedTime   time.Time `json:"-" form:"-" gorm:"-"`
	UpdatedTime   JSONTime  `json:"-" form:"-" gorm:"-"`
	DeletedTime   time.Time `json:"-" form:"-" gorm:"-"`
	PageSize      int       `gorm:"-" json:"-" form:"pageSize"`
	PageIndex     int       `gorm:"-" json:"-" form:"pageIndex"`
}

func CreateOrder(model *Order) (*pb.OrderCommonReply, error) {
	db := DB.Table("tb_order").Create(model)
	var m pb.OrderCommonReply
	if db.Error != nil {
		msg := fmt.Sprintf("mysql tb_order 插入失败, model:%+v", model)
		log.Printf(msg)
		m.Msg = msg
		return &m, db.Error
	}
	return &m, nil
}

/*
综合查询
查询优化: 延迟关联,覆盖索引.内连接查询已经从索引上取得了需要的主键，只会对pageSize个主键关联原表查找，减少了mysql扫描那些需要丢弃的行
SELECT t1.goods_uuid,t1.goods_type_id,t1.primary_type,t1.secondary_type,t1.price,t1.title,t1.subtitle,t1.stock
FROM tb_order AS t1 INNER JOIN (
	SELECT id FROM tb_order WHERE primary_type='pants' AND secondary_type='casual_pants' ORDER BY id DESC LIMIT 10000,20
) AS t2
ON t1.id=t2.id
*/
func QueryOrder(request *pb.OrderRequest) (*pb.OrderReply, error) {
	// 如果是查指定id
	if len(request.GoodsUuid) > 0 {
		item, err := GetOrder(request)
		m := pb.OrderReply{Data: []*pb.OrderReplyItem{item}}
		return &m, err
	}
	// 如果没传商品类型
	if len(request.PrimaryType) == 0 || len(request.SecondaryType) == 0 {
		results := make([]*pb.OrderReplyItem, 0, request.PageSize)
		sql := `SELECT goods_uuid,img,title,subtitle,price,primary_type,secondary_type,goods_type_id,order_status FROM tb_order WHERE id>%d ORDER BY id ASC LIMIT %d`
		sql = fmt.Sprintf(sql, (request.PageIndex-1)*request.PageSize, request.PageSize)
		db := DB.Raw(sql).Find(&results)
		reply := &pb.OrderReply{Data: results}
		return reply, db.Error
	}
	// 传了商品类型
	results := make([]*pb.OrderReplyItem, 0, request.PageSize)
	sql := `SELECT t1.goods_uuid,t1.img,t1.title,t1.subtitle,t1.price,t1.primary_type,t1.secondary_type,t1.stock FROM tb_order AS t1 INNER JOIN (
	SELECT id FROM tb_order WHERE primary_type='%s' AND secondary_type='%s' ORDER BY id DESC LIMIT %d OFFSET %d) AS t2 ON t1.id=t2.id`
	sql = fmt.Sprintf(sql, request.PrimaryType, request.SecondaryType, request.PageSize, (request.PageIndex-1)*request.PageSize)
	db := DB.Raw(sql).Find(&results)
	reply := &pb.OrderReply{Data: results}
	return reply, db.Error
}

// 根据uuid查找商品详情记录
// 给goods_uuid添加了唯一约束, 根据goods_uuid来查找，mysql可直接定位到这条记录，无需再优化了
func GetOrder(request *pb.OrderRequest) (*pb.OrderReplyItem, error) {
	reply := &pb.OrderReplyItem{}
	queryFields := []string{"goods_uuid", "img", "title", "subtitle", "price", "primary_type", "secondary_type", "goods_type_id"}
	db := DB.Table("tb_order").Select(queryFields).Where("goods_uuid=?", request.GoodsUuid).First(reply)
	return reply, db.Error
}

func GetOrderStatistic() *pb.OrderStatisticReply {
	results := &pb.OrderStatisticReply{}
	relate := map[int]string{1: "shirt", 2: "jacket", 3: "casual_pants", 4: "sports_pants", 5: "basketball_shoes", 6: "casual_shoes"}
	for goodsTypeId, goodsTypeName := range relate {
		result := pb.OrderStatisticItem{}
		var count int64
		db := DB.Table("tb_order").Where("goods_type_id=?", goodsTypeId).Count(&count)
		if db.Error != nil {
			return results
		}
		result.GoodsType = goodsTypeName
		result.Count = count
		results.Data = append(results.Data, &result)
	}
	return results
}
