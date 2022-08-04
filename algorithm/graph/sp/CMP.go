package sp

import "go_learning/datastructures/graph"

// CMP 关键路径 优先级限制下的并行任务调度问题
// 在原始任务构成的带权有向图基础上，需要添加3组额外的边：权重为0的边用于分隔
// 1、添加任务开始指向任务结束的边，带权重
// 2、添加起点指向任务开始的边，权重为0
// 3、添加任务结束指向终点的边，权重为0
// 再基于 AcyclicLP 无环有向图 单点最长路径,计算得到关键路径，也就是任务最佳的执行顺序
// ps: 无环+最长路径，这条路径必然会连通所有的任务顶点
type CMP struct {
	N          int // 任务数
	start, end int // 任务起点终点
	lp         *AcyclicLP[int]
}

type Jobs struct {
	Weight float64 // 权重：任务耗时
	Next   []int   // 方向：该任务完成后才可继续执行的任务
}

func NewCMP(jobs []Jobs) *CMP {
	// 起点终点，因为还有等同于任务数的结束顶点
	N := len(jobs)
	start, end := 2*N, 2*N+1

	// 构建带权有向图
	G := graph.NewEdgeWeightedDiGraph[int]()
	for i, job := range jobs {
		// 添加任务开始指向任务结束的边
		G.AddEdge(*graph.NewDirectedEdge[int](i, i+N, job.Weight))
		// 添加起点指向任务开始的边
		G.AddEdge(*graph.NewDirectedEdge[int](start, i, 0))
		// 添加任务结束指向终点的边
		G.AddEdge(*graph.NewDirectedEdge[int](i+N, end, 0))

		// 添加任务结束指向下个任务开始的边
		for _, next := range job.Next {
			G.AddEdge(*graph.NewDirectedEdge[int](i+N, next, 0))
		}
	}

	// 计算关键路径
	lp := NewAcyclicLP[int](*G, start)

	return &CMP{
		N:     N,
		start: start,
		end:   end,
		lp:    lp,
	}
}

// StartTimes 打印关键路径上，每个任务应当开始执行的时间
func (c *CMP) StartTimes() (res []float64) {
	res = make([]float64, 0)
	for i := 0; i < c.N; i++ {
		res = append(res, c.lp.DistTo(i))
	}
	return res
}

// FinishTime 全部任务执行完总耗时
func (c *CMP) FinishTime() float64 {
	return c.lp.DistTo(c.end)
}
