package queue

import "go_learning/common"

// IMaxPQ 最大优先队列
type IMaxPQ interface {
	Insert(key int)                    // 向优先队列插入一个元素
	Max() int                          // 返回优先队列最大的元素
	DelMax()                           // 删除优先队列最大的元素
	IsEmpty() bool                     // 判断优先队列是否为空
	Size() int                         // 获取优先队列大小
	Less(firstIdx, secondIdx int) bool // 判断元素1是否小于元素2
}

// IMinPQ 最小优先队列
// 和IMaxPQ 差不多，注意下比较方向就行了
type IMinPQ interface {
	Insert(key int)
	Min() int
	DelMin()
	IsEmpty() bool
	Size() int
	Less(firstIdx, secondIdx int) bool
}

// IIndexMinPQ 索引最小优先队列
type IIndexMinPQ[T common.IntAny] interface {
	Insert(idx int, item *T)           // 插入元素到idx索引位置
	Change(idx int, item *T)           // 修改idx索引对应的元素
	Contain(idx int) bool              // 检查是否存在索引idx的元素
	Delete(idx int)                    // 删除索引idx的元素
	Min() *T                           // 返回最小元素
	MinIdx() int                       // 返回最小元素的索引
	DelMin() int                       // 删除最小元素
	IsEmpty() bool                     // 判断队列是否为空
	Size() int                         // 获取队列元素个数
	Less(firstIdx, secondIdx int) bool // 判断元素1是否小于元素2
}
