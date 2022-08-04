package unionfind

// QuickUnion 加权算法
type QuickUnion struct {
	storage []int // 存储结构
	deep    []int // 各个根节点对应的树 的深度
	count   int   // 连通分量数量
}

func NewQuickUnion(max int) *QuickUnion {
	res := &QuickUnion{
		storage: make([]int, max+1, max+1),
		deep:    make([]int, max+1, max+1),
	}
	for i := 0; i < max+1; i++ {
		res.storage[i] = i
		res.deep[i] = 1
	}
	res.count = max + 1
	return res
}

// Find 查找节点
func (q *QuickUnion) Find(data int) int {
	for data != q.storage[data] {
		data = q.storage[data]
	}
	return data
}

// Connected 判断是否联通
func (q *QuickUnion) Connected(first, second int) bool {
	return q.Find(first) == q.Find(second)
}

// Union 算法核心，将不同的树进行连接，根据树深度判断，小树连接到大树
func (q *QuickUnion) Union(first, second int) {
	dF := q.Find(first)
	dS := q.Find(second)
	if dF == dS {
		return
	}
	// 小树 连接到 大树
	if q.deep[first] < q.deep[second] {
		q.storage[first] = dS
		q.deep[second] += q.deep[first]
	} else {
		q.storage[second] = dF
		q.deep[first] += q.deep[second]
	}
	q.count--
}

// Count 获取连通分量数
func (q *QuickUnion) Count() int {
	return q.count
}
