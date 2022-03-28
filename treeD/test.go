package main

import (
	"encoding/json"
	"fmt"
	tree "github.com/azhengyongqin/golang-tree-menu"
)

func main() {
	//var systemMenus = []SystemMenu{
	//	{
	//		Id:     8,
	//		Name: "7",
	//		FatherId:  8,
	//	},
	//	{
	//		Id:     7,
	//		Name: "4",
	//		FatherId:  8,
	//	},
	//	{
	//		Id:     4,
	//		Name: "5",
	//		FatherId:  7,
	//	},
	//	{
	//		Id:     1,
	//		Name: "3",
	//		FatherId:  4,
	//	},
	//	{
	//		Id:     3,
	//		Name: "1",
	//		FatherId:  4,
	//	},
	//	{
	//		Id:     5,
	//		Name: "6",
	//		FatherId:  4,
	//	},
	//	{
	//		Id:     2,
	//		Name: "2",
	//		FatherId:  3,
	//	},
	//	{
	//		Id:     6,
	//		Name: "2",
	//		FatherId:  5,
	//	},
	//}
	//
	//fmt.Println(systemMenus)
	//var systemMenu SystemMenu{}

	allMenu := []SystemMenu{
		{Id: 1, FatherId: 0, Name: "系统总览", Route: "/systemOverview", Icon: "icon-system"},
		{Id: 2, FatherId: 0, Name: "系统配置", Route: "/systemConfig", Icon: "icon-config"},
		{Id: 3, FatherId: 1, Name: "资产", Route: "/asset", Icon: "icon-asset"},
		{Id: 4, FatherId: 1, Name: "动环", Route: "/pe", Icon: "icon-pe"},
		{Id: 5, FatherId: 2, Name: "菜单配置", Route: "/menuConfig", Icon: "icon-menu-config"},
		{Id: 6, FatherId: 3, Name: "设备", Route: "/device", Icon: "icon-device"},
		{Id: 7, FatherId: 3, Name: "机柜", Route: "/device", Icon: "icon-device"},
	}

	resp := tree.GenerateTree(SystemMenus.ConvertToINodeArray(allMenu), nil)
	bytes, _ := json.MarshalIndent(resp, "", "\t")
	fmt.Println(string(bytes))
	fmt.Println("=================")
	bys, _ := json.Marshal(resp)
	fmt.Println(string(bys))

}

// SystemMenu 定义我们自己的菜单对象
type SystemMenu struct {
	Id       int    `json:"id"`        //id
	FatherId int    `json:"father_id"` //上级菜单id
	Name     string `json:"name"`      //菜单名
	Route    string `json:"route"`     //页面路径
	Icon     string `json:"icon"`      //图标路径
}

func (s SystemMenu) GetTitle() string {
	return s.Name
}
func (s SystemMenu) GetId() int {
	return s.Id
}
func (s SystemMenu) GetFatherId() int {
	return s.FatherId
}
func (s SystemMenu) GetData() interface{} {
	return s
}
func (s SystemMenu) IsRoot() bool {
	// 这里通过FatherId等于0 或者 FatherId等于自身Id表示顶层根节点
	return s.FatherId == 0 || s.FatherId == s.Id
}

type SystemMenus []SystemMenu

// ConvertToINodeArray 将当前数组转换成父类 INode 接口 数组
func (s SystemMenus) ConvertToINodeArray() (nodes []tree.INode) {
	for _, v := range s {
		nodes = append(nodes, v)
	}
	return
}
