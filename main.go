package main

import "fmt"

//func main() {
//等待同步示例
//concurrency.WaitForGroup()

//同步示例
//con := concurrency.Concurrency{}
//con.Test()

//测试打乒乓
//concurrency.TestPingPong()
//fmt.Println("ok")
//}

type CascaderTree struct {
	ID       int            `json:"id" bson:"_id"`             // 节点id
	ParentId string         `json:"parentId" bson:"parent_id"` // 父节点id
	Name     string         `json:"name" bson:"name"`          // 名字
	Nodes    []CascaderTree `json:"nodes"`
}

func main() {
	a := CascaderTree{}
	var b CascaderTree
	var c map[int]string
	var d interface{} = c

	fmt.Println(a, b, c,d)
}
