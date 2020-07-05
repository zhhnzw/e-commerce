package models

import (
	"fmt"
	"order/pb"
	"order/utils"
	"strings"
	"time"
)

type User struct {
	Id          int            `json:"id"`
	UserName    string         `json:"userName" form:"userName"`
	NickName    string         `json:"nickName" form:"nickName"`
	Mobile      string         `json:"mobile" form:"mobile"`
	Email       string         `json:"email" form:"email"`
	Avatar      string         `json:"avatar" form:"avatar"`
	CreatedTime time.Time      `json:"-" form:"-" gorm:"-"`
	UpdatedTime utils.JSONTime `json:"updateTime" form:"-" gorm:"-"`
}

func (*User) TableName() string {
	return "tb_user"
}

type Goods struct {
	Id            int
	GoodsUuid     string
	GoodsTypeId   int64
	PrimaryType   string
	SecondaryType string
	Img           string
	Title         string
	Subtitle      string
	Price         int64
	CreatedTime   time.Time      `json:"-" form:"-" gorm:"-"`
	UpdatedTime   utils.JSONTime `json:"-" form:"-" gorm:"-"`
	DeletedTime   time.Time      `json:"-" form:"-" gorm:"-"`
}

func (*Goods) TableName() string {
	return "tb_goods"
}

type Order struct {
	Id            int
	OrderId       string
	GoodsUuid     string
	GoodsTypeId   int64
	PrimaryType   string `gorm:"-"`
	SecondaryType string `gorm:"-"`
	Img           string `gorm:"-"`
	Title         string `gorm:"-"`
	Subtitle      string `gorm:"-"`
	Price         int64
	OrderStatus   string
	UserName      string
	NickName      string         `json:"nickName" form:"nickName" gorm:"-"`
	Mobile        string         `json:"mobile" form:"mobile" gorm:"-"`
	Email         string         `json:"email" form:"email" gorm:"-"`
	Avatar        string         `json:"avatar" form:"avatar" gorm:"-"`
	CreatedTime   time.Time      `json:"-" form:"-" gorm:"-"`
	UpdatedTime   utils.JSONTime `json:"-" form:"-" gorm:"-"`
	DeletedTime   time.Time      `json:"-" form:"-" gorm:"-"`
	PageSize      int            `gorm:"-" json:"-" form:"pageSize"`
	PageIndex     int            `gorm:"-" json:"-" form:"pageIndex"`
}

func (*Order) TableName() string {
	return "tb_order"
}

func CreateOrder(model *Order) (*pb.OrderCommonReply, error) {
	var reply pb.OrderCommonReply
	goodsModel := &Goods{
		GoodsUuid:     model.GoodsUuid,
		GoodsTypeId:   model.GoodsTypeId,
		PrimaryType:   model.PrimaryType,
		SecondaryType: model.SecondaryType,
		Img:           model.Img,
		Title:         model.Title,
		Subtitle:      model.Subtitle,
		Price:         model.Price,
	}
	dbGoods := DB.Create(goodsModel)
	if !strings.Contains(dbGoods.Error.Error(), "Duplicate") {
		msg := fmt.Sprintf("mysql tb_goods 插入失败, model:%+v", model)
		reply.Msg = msg
		return &reply, dbGoods.Error
	}
	userModel := &User{
		UserName: model.UserName,
		NickName: model.NickName,
		Mobile:   model.Mobile,
		Email:    model.Email,
		Avatar:   model.Avatar,
	}
	dbUser := DB.Create(userModel)
	if !strings.Contains(dbUser.Error.Error(), "Duplicate") {
		msg := fmt.Sprintf("mysql tb_user 插入失败, model:%+v", model)
		reply.Msg = msg
		return &reply, dbGoods.Error
	}
	model.OrderId = utils.GenerateOrderId()
	db := DB.Create(model)
	if db.Error != nil {
		msg := fmt.Sprintf("mysql tb_order 插入失败, model:%+v", model)
		reply.Msg = msg
		return &reply, db.Error
	}
	return &reply, nil
}

/*
综合查询
查询优化: 延迟关联,覆盖索引.内连接查询已经从索引上取得了需要的主键，只会对pageSize个主键关联原表查找，减少了mysql扫描那些需要丢弃的行
SELECT t1.goods_uuid,t1.goods_type_id,t1.primary_type,t1.secondary_type,t1.price,t1.title,t1.subtitle
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
		sql := `SELECT goods_uuid,price,goods_type_id,order_status,user_name FROM tb_order WHERE id>%d ORDER BY id ASC LIMIT %d`
		sql = fmt.Sprintf(sql, (request.PageIndex-1)*request.PageSize, request.PageSize)
		db := DB.Raw(sql).Find(&results)
		reply := &pb.OrderReply{Data: results}
		return reply, db.Error
	}
	// 传了商品类型
	results := make([]*pb.OrderReplyItem, 0, request.PageSize)
	sql := `SELECT t1.goods_uuid,t1.price,t1.user_name FROM tb_order AS t1 INNER JOIN (
	SELECT id FROM tb_order WHERE goods_type_id=%d ORDER BY id DESC LIMIT %d OFFSET %d) AS t2 ON t1.id=t2.id`
	sql = fmt.Sprintf(sql, request.GoodsTypeId, request.PageSize, (request.PageIndex-1)*request.PageSize)
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
