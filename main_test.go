package main

import (
	"strconv"
	"testing"
	. "github.com/smartystreets/goconvey/convey"
)
//测试add方法
func TestAdd(t *testing.T) {
	//列出要测试的参数和返回结果
	cases := []struct {
		first    int
		second   int
		excepted int
	}{
		{1, 2, 3},
		{1, 3, 4},
		{2, 2, 4},
	}

	//遍历，进行测试
	for _, c := range cases {
		result := Add(c.first, c.second)
		if result != c.excepted {
			t.Fatalf("add function failed, first: %d, second:%d, execpted:%d, result:%d", c.first, c.second, c.excepted, result)
		}
	}
}
//测试add方法，convey
func TestAddConvey(t *testing.T) {
	//列出要测试的参数和返回结果
	cases := []struct {
		first    int
		second   int
		excepted int
	}{
		{1, 2, 3},
		{1, 3, 4},
		{2, 2, 4},
	}

	//遍历，进行测试
	for _, c := range cases {
		Convey("测试加和结果", t, func() {
			So(Add(c.first, c.second), ShouldEqual,c.excepted)
		})
	}
}

//测试add方法，convey
func TestAddConveyNest(t *testing.T) {
	//列出要测试的参数和返回结果
	cases := []struct {
		first    int
		second   int
		excepted int
	}{
		{1, 2, 3},
		{1, 3, 4},
		{2, 2, 4},
	}

	Convey("开始一组测试加和结果", t, func() {
		//遍历，进行测试
		for _, c := range cases {
			Convey("加和测试:"+strconv.Itoa(c.first)+"+"+strconv.Itoa(c.second), func() {
				So(Add(c.first, c.second), ShouldEqual,c.excepted)
			})
		}
	})



}


