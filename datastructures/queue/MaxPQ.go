package queue

// MaxPQ 最大优先队列（基于堆实现）
/**
  思想：优先队列基于一个堆的完全二叉树表示，
	1、添加时，将新元素添加到队尾，使用swim进行上浮。
	2、删除元素时，将首尾交换，移除尾部，此时队首为最小元素，然后使用sink进行下沉
*/
type MaxPQ struct {
	queue []int // 数组实现的堆,0号位不使用
	N     int   // 元素个数
}

func NewMaxPQ() *MaxPQ {
	res := &MaxPQ{
		queue: make([]int, 0),
	}
	res.queue = append(res.queue, 0)
	return res
}

func (m *MaxPQ) Insert(key int) {
	// 插入到队尾然后执行上浮
	m.queue = append(m.queue, key)
	m.N++
	m.swim(m.N)
}

func (m *MaxPQ) Max() int {
	if m.Size() > 0 {
		return m.queue[1]
	}
	return 0
}

func (m *MaxPQ) DelMax() {
	// 删除时，先交换末尾和头部元素，然后删除尾部元素，对头部元素执行下沉
	switch m.N {
	case 0:
	case 1:
		m.queue = m.queue[:1]
		m.N--
	default:
		m.Exchange(1, m.N)
		m.queue = m.queue[:m.N]
		m.N--
		m.sink(1)
	}
	return
}

func (m *MaxPQ) IsEmpty() bool {
	return len(m.queue) == 1 // 默认0号不使用
}

func (m *MaxPQ) Size() int {
	return m.N
}

func (m *MaxPQ) Less(firstIdx, secondIdx int) bool {
	return m.queue[firstIdx] < m.queue[secondIdx]
}

func (m *MaxPQ) Exchange(firstIdx, secondIdx int) {
	m.queue[firstIdx], m.queue[secondIdx] = m.queue[secondIdx], m.queue[firstIdx]
}

// Swim 元素上浮
func (m *MaxPQ) swim(index int) {
	for index > 1 && m.Less(index/2, index) {
		m.Exchange(index/2, index)
		index /= 2
	}
}

// Sink 元素下沉
func (m *MaxPQ) sink(index int) {
	// 下沉需要两次比较，一次用来找出较大的子节点，一次用来确定节点是否需要下沉

	for index*2 <= m.N {
		doubleIdx := index * 2
		// 找出较大的子节点，在堆底部只有一个子节点时不用比较
		if doubleIdx < m.N && m.Less(doubleIdx, doubleIdx+1) {
			doubleIdx += 1
		}

		// 比较，确定节点是否需要下沉
		if m.Less(index, doubleIdx) {
			m.Exchange(index, doubleIdx)
			index = doubleIdx
		} else {
			break
		}
	}
}
