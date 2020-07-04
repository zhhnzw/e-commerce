package models

import (
	"fmt"
	"goods/pb"
)

/*
商品的综合查询
查询优化: 延迟关联,覆盖索引.内连接查询已经从索引上取得了需要的主键，只会对pageSize个主键关联原表查找，减少了mysql扫描那些需要丢弃的行
SELECT t1.goods_uuid,t1.goods_from,t1.goods_type_id,t1.primary_type,t1.secondary_type,t1.publish_date,t1.price,t1.title,t1.subtitle
FROM tb_goods AS t1 INNER JOIN (
	SELECT id FROM tb_goods WHERE primary_type='pants' AND secondary_type='casual_pants' ORDER BY id DESC LIMIT 10000,20
) AS t2
ON t1.id=t2.id
*/
func QueryGoods(request *pb.GoodsRequest) (*pb.GoodsReply, error) {
	// 如果是查指定id
	if len(request.GoodsUuid) > 0 {
		item, err := GetGoods(request)
		m := pb.GoodsReply{Data: []*pb.GoodsReplyItem{item}}
		return &m, err
	}
	// 如果没传商品类型
	if len(request.PrimaryType) == 0 || len(request.SecondaryType) == 0 {
		results := make([]*pb.GoodsReplyItem, 0, request.PageSize)
		sql := `SELECT goods_uuid,goods_from,img,title,subtitle,price,publish_date,primary_type,secondary_type,is_valid,goods_type_id FROM tb_goods WHERE id>%d ORDER BY id ASC LIMIT %d`
		sql = fmt.Sprintf(sql, (request.PageIndex-1)*request.PageSize, request.PageSize)
		db := DB.Raw(sql).Find(&results)
		reply := &pb.GoodsReply{Data: results}
		return reply, db.Error
	}
	// 传了商品类型
	results := make([]*pb.GoodsReplyItem, 0, request.PageSize)
	sql := `SELECT t1.goods_uuid,t1.goods_from,t1.img,t1.title,t1.subtitle,t1.price,t1.publish_date,t1.primary_type,t1.secondary_type,t1.is_valid,t1.goods_type_id FROM tb_goods AS t1 INNER JOIN (
	SELECT id FROM tb_goods WHERE primary_type='%s' AND secondary_type='%s' ORDER BY id DESC LIMIT %d OFFSET %d) AS t2 ON t1.id=t2.id`
	sql = fmt.Sprintf(sql, request.PrimaryType, request.SecondaryType, request.PageSize, (request.PageIndex-1)*request.PageSize)
	db := DB.Raw(sql).Find(&results)
	reply := &pb.GoodsReply{Data: results}
	return reply, db.Error
}

// 根据uuid查找商品详情记录
// 给goods_uuid添加了唯一约束, 根据goods_uuid来查找，mysql可直接定位到这条记录，无需再优化了
func GetGoods(request *pb.GoodsRequest) (*pb.GoodsReplyItem, error) {
	reply := &pb.GoodsReplyItem{}
	queryFields := []string{"goods_uuid", "goods_from", "img", "title", "subtitle", "price", "publish_date", "primary_type", "secondary_type", "goods_type_id", "is_valid", "imgs"}
	db := DB.Table("tb_goods").Select(queryFields).Where("goods_uuid=?", request.GoodsUuid).First(reply)
	return reply, db.Error
}

// 查count，查出粗略值即可
//func GetGoodsStatistic() *pb.GoodsStatisticReply {
//	results := &pb.GoodsStatisticReply{}
//	type explainStr struct {
//		Rows int64
//	}
//	relate := map[int]string{1: "shirt", 2: "jacket", 3: "casual_pants", 4: "sports_pants", 5: "basketball_shoes", 6: "casual_shoes"}
//	for goodsTypeId, goodsTypeName := range relate {
//		result := pb.GoodsStatisticItem{}
//		explain := explainStr{}
//		sql := fmt.Sprintf("EXPLAIN SELECT count(*) FROM tb_goods WHERE goods_type_id=%d", goodsTypeId)
//		db := DB.Raw(sql).Find(&explain)
//		if db.Error != nil {
//			return results
//		}
//		result.GoodsType = goodsTypeName
//		result.Count = explain.Rows
//		results.Data = append(results.Data, &result)
//	}
//	return results
//}
// 粗略值误差还是太大了, 用准确值吧
func GetGoodsStatistic() *pb.GoodsStatisticReply {
	results := &pb.GoodsStatisticReply{}
	relate := map[int]string{1: "shirt", 2: "jacket", 3: "casual_pants", 4: "sports_pants", 5: "basketball_shoes", 6: "casual_shoes"}
	for goodsTypeId, goodsTypeName := range relate {
		result := pb.GoodsStatisticItem{}
		var count int64
		db := DB.Table("tb_goods").Where("goods_type_id=?", goodsTypeId).Count(&count)
		if db.Error != nil {
			return results
		}
		result.GoodsType = goodsTypeName
		result.Count = count
		results.Data = append(results.Data, &result)
	}
	return results
}
