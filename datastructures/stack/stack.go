package stack

// LIFOStack 后进先出栈
type LIFOStack[T any] struct {
	data    []T
	num     int // 元素数量
	max     int
	idxIter int // 当前索引，用于支持迭代
}

func NewLIFOStack[T any]() *LIFOStack[T] {
	return &LIFOStack[T]{
		data: make([]T, 0),
		num:  0,
		max:  1,
	}
}

func (l *LIFOStack[T]) IsEmpty() bool {
	return l.num == 0
}

func (l *LIFOStack[T]) Push(str T) {
	// go切片到达容量上限自动扩容,如果用数组的话需要手动扩容
	// l.Resize()
	l.data = append(l.data, str)
	l.num++
}

func (l LIFOStack[T]) Resize() {
	if len(l.data) == l.max {
		l.max *= 2
		temp := make([]T, 0, l.max)    // 扩容
		temp = append(temp, l.data...) // 迁移数据
	}
}

func (l *LIFOStack[T]) Pop() (res T) {
	if l.num > 0 {
		res = l.data[l.num-1]
		l.data = l.data[:l.num-1]
		l.num--
	}
	return
}

func (l *LIFOStack[T]) NewIterator() *LIFOStack[T] {
	l.idxIter = l.num
	return l
}

func (l *LIFOStack[T]) HasNext() bool {
	// 因为是栈，所以是从最后一位向前遍历的
	return l.idxIter > 0
}

func (l *LIFOStack[T]) Next() (res T) {
	if l.idxIter > 0 {
		l.idxIter--
		res = l.data[l.idxIter]
	}
	return
}
