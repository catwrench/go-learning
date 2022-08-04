package st

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

type TypeST int

const (
	TypeSequentialSearchST = iota
	TypeBinarySearchST
	TypeBST
	TypeRedBlackBST
)

func Input1() []int {
	return []int{1, 2, 4, 9, 3, 8, 5, 7, 6, 0}
}

// FrequencyCounter 频率计数器
type FrequencyCounter struct {
	n    int    // 随机生成的数组长度
	t    int    // 执行次数
	base TypeST // 作为比较的基准数据结构
}

func NewFrequencyCounter() *FrequencyCounter {
	return &FrequencyCounter{}
}

func (f *FrequencyCounter) SetN(N int) *FrequencyCounter {
	f.n = N
	return f
}

func (f *FrequencyCounter) SetT(t int) *FrequencyCounter {
	f.t = t
	return f
}

func (f *FrequencyCounter) SetBase(base TypeST) *FrequencyCounter {
	f.base = base
	return f
}

func (f *FrequencyCounter) Time(t TypeST, arr []int) int64 {
	arrStr := make([]string, len(arr))
	for i := range arr {
		arrStr = append(arrStr, strconv.Itoa(arr[i]))
	}

	start := time.Now()
	switch t {
	case TypeSequentialSearchST:
		st := NewSequentialSearchST[string, int]()
		for i := range arr {
			if st.Contains(arrStr[i]) {
				st.Put(arrStr[i], st.Get(arrStr[i])+1)
			} else {
				st.Put(arrStr[i], arr[1])
			}
		}

		k, v, max := "", 0, " "
		st.Put(max, 0)
		it := st.NewIterator()
		for it.HasNext() {
			k, v = it.Next()
			if v > st.Get(max) {
				max = k
			}
		}
	case TypeBinarySearchST:
		st := NewBinarySearchST[int, int](len(arr))
		for i := range arr {
			if st.Contains(arr[i]) {
				st.Put(arr[i], st.Get(arr[i])+1)
			} else {
				st.Put(arr[i], 1)
			}
		}

		var k, v, max int
		st.NewIterator()
		for st.HasNext() {
			k, v = st.Next()
			if v > max {
				max = k
			}
		}
	case TypeBST:
		st := NewBST[int, int]()
		for i := range arr {
			if st.Contains(arr[i]) {
				st.Put(arr[i], st.Get(arr[i])+1)
			} else {
				st.Put(arr[i], 1)
			}
		}

		var k, max int
		it := st.NewIterator()
		for it.HasNext() {
			k = it.Next()
			if st.Get(k) > max {
				max = k
			}
		}
	case TypeRedBlackBST:
		st := NewRedBlackBST[int, int]()
		for i := range arr {
			if st.Contains(arr[i]) {
				st.Put(arr[i], st.Get(arr[i])+1)
			} else {
				st.Put(arr[i], 1)
			}
		}

		var k, max int
		it := st.NewIterator()
		for it.HasNext() {
			k = it.Next()
			if st.Get(k) > max {
				max = k
			}
		}
	}
	return time.Now().Sub(start).Nanoseconds()
}

func (f *FrequencyCounter) TimeRandomInput(t TypeST) int64 {
	var total int64
	arr := make([]int, f.n, f.n)
	for i := 0; i < f.t; i++ {
		for j := 0; j < f.n; j++ {
			arr[j] = rand.Intn(100)
		}
		total += f.Time(t, arr)
	}
	return total
}

func (f *FrequencyCounter) Compare(types ...TypeST) {
	fmt.Printf("每个数据结构执行 %d 次，每次处理 %d 个元素 \n\n", f.t, f.n)

	var baseTotal int64
	m := make(map[TypeST]int64)
	for _, typeSTt := range types {
		m[typeSTt] = f.TimeRandomInput(typeSTt)
		if f.base == typeSTt {
			baseTotal = m[typeSTt]
		}
	}
	for i, typeSTt := range types {
		total := m[typeSTt]
		fmt.Printf("数据结构%v执行耗时： %v ns | 单次耗时： %v ns | 达到基准：%.2f %% 的性能 \n",
			i+1,
			total,
			total/int64(f.t),
			float64(100*(baseTotal))/float64(total),
		)
	}
}
