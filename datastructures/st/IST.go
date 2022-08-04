package st

// IST 符号表 接口
type IST[K comparable, V any] interface {
	Put(K, V)
	Get(K) V
	Del(K)
	Contains(K) bool
	IsEmpty() bool
	Size() int
}

// IOrderST 有序符号表 接口
type IOrderST[K comparable, V any] interface {
	Put(K, V)
	Get(K) V
	Del(K)
	Contains(K) bool
	IsEmpty() bool
	Size() int
	Rank(K) int   // 小于k的键的数量
	Select(int) K // 返回排名为rank的键
	Min() K
	Max() K
	DelMin()
	DelMax()
	Keys() []K
	Floor(K) K   // 返回小于等于K的最大键
	Ceiling(K) K // 返回大于等于K的最小键
}
