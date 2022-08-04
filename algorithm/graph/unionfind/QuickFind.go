package unionfind

// QuickFind 算法
type QuickFind struct {
	storage []int // 存储结构
	count   int   // 连通分量数量
}

// NewQuickFind max 为最大值，共需要初始化 max+1 的容量
func NewQuickFind(max int) *QuickFind {
	res := &QuickFind{
		storage: make([]int, max+1, max+1),
	}
	// 按索引顺序初始化数组
	for k := range res.storage {
		res.storage[k] = k
	}

	res.count = max + 1
	return res
}

// Find 查找节点
func (q *QuickFind) Find(data int) int {
	return q.storage[data]
}

// Connected 判断是否联通
func (q *QuickFind) Connected(first, second int) bool {
	return q.Find(first) == q.Find(second)
}

// Union 算法核心，将值不同的两个节点染色 （值相同的视为在同一个连通分量中）
func (q *QuickFind) Union(first, second int) {
	// 两节点不同时，完成转变
	dF := q.Find(first)
	dS := q.Find(second)
	if dF == dS {
		return
	}

	// 将 first 转变为 second
	for i := 0; i < len(q.storage); i++ {
		if q.storage[i] == dF {
			q.storage[i] = dS
		}
	}
	q.count--
	return
}

// Count 获取连通分量数
func (q *QuickFind) Count() int {
	return q.count
}
