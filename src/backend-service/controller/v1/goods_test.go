package v1_test

import (
	"backend-service/test"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestGetCommodities(t *testing.T) {
	Convey("TestGetCommodities", t, func() {
		URL := "/v1/goods?primaryType=clothes&secondaryType=shirt&pageSize=20&pageIndex=1&isValid=test"
		_, statusCode := test.SendRequest(URL, "", "GET")
		So(statusCode, ShouldEqual, 400)
	})
	Convey("TestGetCommodities", t, func() {
		URL := "/v1/goods?primaryType=clothes&secondaryType=shirt&pageSize=20&pageIndex=1"
		_, statusCode := test.SendRequest(URL, "", "GET")
		So(statusCode, ShouldEqual, 200)
	})
}

func TestGetgoods(t *testing.T) {
	Convey("TestGetgoods", t, func() {
		URL := "/v1/goods/b7b10c01-62b8-42c7-a8c4-8efe119cd326"
		_, statusCode := test.SendRequest(URL, "", "GET")
		So(statusCode, ShouldEqual, 200)
	})
}

func TestGetHot(t *testing.T) {
	Convey("TestGetHot", t, func() {
		URL := "/v1/hot?primaryType=clothes&pageSize=10&pageIndex=test"
		_, statusCode := test.SendRequest(URL, "", "GET")
		So(statusCode, ShouldEqual, 400)
	})
	Convey("TestGetHot", t, func() {
		URL := "/v1/hot?primaryType=clothes&pageSize=10&pageIndex=1"
		_, statusCode := test.SendRequest(URL, "", "GET")
		So(statusCode, ShouldEqual, 200)
	})
}

func TestHitCache(t *testing.T) {
	Convey("TestHitCache", t, func() {
		// 先访问商品榜单数据, 若内部没有缓存，触发数据库访问
		urlForHot := "/v1/hot?primaryType=shoes&pageSize=10&pageIndex=1"
		_, statusCode := test.SendRequest(urlForHot, "", "GET")

		// 访问2次商品详情页面 -> 详情数据载入缓存
		urlForDetail := "/v1/goods/b7b10c01-62b8-42c7-a8c4-8efe119cd326"
		_, statusCode = test.SendRequest(urlForDetail, "", "GET")
		So(statusCode, ShouldEqual, 200)
		_, statusCode = test.SendRequest(urlForDetail, "", "GET")
		So(statusCode, ShouldEqual, 200)

		// 第3次访问商品详情页，命中缓存
		_, statusCode = test.SendRequest(urlForDetail, "", "GET")
		So(statusCode, ShouldEqual, 200)

		// 再次访问商品榜单数据, 触发商品榜单数据接口内部的缓存访问
		_, statusCode = test.SendRequest(urlForHot, "", "GET")
		So(statusCode, ShouldEqual, 200)

		// 第3次访问商品榜单数据，命中缓存
		_, statusCode = test.SendRequest(urlForHot, "", "GET")
		So(statusCode, ShouldEqual, 200)
	})
}
