package unionfind

type IUnionFind interface {
	Find(data int) int                // 查找节点
	Connected(first, second int) bool // 判断是否联通
	Union(first, second int)          // 在两个节点之间添加连接
	Count() int                       // 连通分量数量
}
