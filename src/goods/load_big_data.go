// 为大数据量查询测试做准备, 插入1000万条记录
package main

import (
	"bytes"
	"fmt"
	"github.com/jinzhu/gorm"
	"goods/conf"
	"goods/models"
	"goods/pb"
	"goods/utils"
	"math/rand"
	"strconv"
	"time"
)

// 批量插入数据
func BatchSave(db *gorm.DB, data []*pb.GoodsRequest) error {
	var buffer bytes.Buffer
	sql := "insert into `tb_goods` (`goods_uuid`,`goods_from`,`goods_type_id`,`primary_type`,`secondary_type`,`price`, `title`, `subtitle`, `img`, `imgs`) values"
	if _, err := buffer.WriteString(sql); err != nil {
		return err
	}
	for i, e := range data {
		if i == len(data)-1 {
			buffer.WriteString(fmt.Sprintf("('%s','%s','%d','%s','%s','%d','%s','%s','%s','%s');",
				e.GoodsUuid, e.GoodsFrom, e.GoodsTypeId, e.PrimaryType, e.SecondaryType, e.Price, e.Title, e.Subtitle, e.Img, e.Imgs))
		} else {
			buffer.WriteString(fmt.Sprintf("('%s','%s','%d','%s','%s','%d','%s','%s','%s','%s'),",
				e.GoodsUuid, e.GoodsFrom, e.GoodsTypeId, e.PrimaryType, e.SecondaryType, e.Price, e.Title, e.Subtitle, e.Img, e.Imgs))
		}
	}
	return db.Exec(buffer.String()).Error
}

func main() {
	conf.InitConfig()
	models.InitGorm()
	goodsFromSrc := []string{"taobao", "jd"}
	comoditySrc := [][5]string{
		{"clothes", "shirt", "1", "CYGNENOIR 20AW秋季字母印花宽松迷彩衬衫日系复古街头衬衣男潮", "衬衣 subtitle"},
		{"clothes", "jacket", "2", "SSUR PLUS X PCMY 联名刷漆教练夹克 PC225136", "夹克subtitle"},
		{"pants", "casual_pants", "3", "休闲裤 title", "休闲裤 subtitle"},
		{"pants", "sports_pants", "4", "运动裤 subtitle", "运动裤 subtitle"},
		{"shoes", "basketball_shoes", "5", "adidas Yeezy 350 Boost V2", "夏日的街头潮鞋之王"},
		{"shoes", "casual_shoes", "6", "Converse 1970s", "简约黑白配色延续经典"},
	}
	// 分2000次批量插入
	for k := 0; k < 2000; k++ {
		data := make([]*pb.GoodsRequest, 0, 5000)
		for i := 0; i < 5000; i++ {
			rand.Seed(time.Now().UnixNano()) // 设置不同的seed，后续的随机操作才是真随机
			randSrc := comoditySrc[rand.Intn(len(comoditySrc))]
			goodsTypeId, err := strconv.ParseInt(randSrc[2], 10, 64)
			utils.CheckErr(err, "")
			imgs := "http://shihuo.hupucdn.com/def/20200430/390bb16aeb93924ccd4c67f5f178cf9a1588213097.jpg," +
				"http://shihuo.hupucdn.com/def/20200430/390bb16aeb93924ccd4c67f5f178cf9a1588213097.jpg," +
				"http://shihuo.hupucdn.com/def/20200409/c182b32c1ef7fcd2fa3b0775cd44f7201586412353.jpg," +
				"http://shihuo.hupucdn.com/def/20200409/2835820a582cc50edcce4c5efe11d8fd1586412413.jpg," +
				"http://shihuo.hupucdn.com/def/20200514/4524408b332abcbf4687d475840325ba1589435366.jpg"
			model := pb.GoodsRequest{
				GoodsUuid:     utils.GetUUID(),
				GoodsFrom:     goodsFromSrc[rand.Intn(len(goodsFromSrc))],
				PrimaryType:   randSrc[0],
				SecondaryType: randSrc[1],
				GoodsTypeId:   goodsTypeId,
				Title:         randSrc[3],
				Subtitle:      randSrc[4],
				Img:           "http://shihuo.hupucdn.com/def/20200430/390bb16aeb93924ccd4c67f5f178cf9a1588213097.jpg",
				Imgs:          imgs,
				Price:         rand.Int63n(100000),
			}
			data = append(data, &model)
		}
		err := BatchSave(models.DB, data)
		utils.CheckErr(err, "")
	}
}
