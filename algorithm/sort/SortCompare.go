package sort

import (
	"fmt"
	"math/rand"
	"time"
)

type SortType int

const (
	TypeSelection SortType = iota
	TypeInsertion
	TypeQuick
	TypeQuick3Way
	TypeShell
	TypeMergeUB
	TypeMergeBU
	TypeHeapSort
)

func Input() []int {
	return []int{1, 2, 3, 4, 5, 9, 8, 6, 0, 7}
}

func Input2() []int {
	return []int{7, 2, 2, 3, 4, 0, 9, 8, 6, 2}
}

func Input3() []int {
	return []int{2, 2, 2, 4, 4, 4, 3, 3, 3, 6, 2}
}

// IsSorted 判断数组是否有序
func IsSorted(arr []int) bool {
	if len(arr) <= 1 {
		return true
	}
	for i := 1; i < len(arr); i++ {
		if arr[i] < arr[i-1] {
			return false
		}
	}
	return true
}

// SortCompare 数组比较
type SortCompare struct {
	n    int      // 随机生成的数组长度
	t    int      // 执行次数
	base SortType // 作为比较的基准算法
}

func NewSortCompare() *SortCompare {
	return &SortCompare{}
}

func (s *SortCompare) SetN(N int) *SortCompare {
	s.n = N
	return s
}

func (s *SortCompare) SetT(T int) *SortCompare {
	s.t = T
	return s
}
func (s *SortCompare) SetBase(base SortType) *SortCompare {
	s.base = base
	return s
}

// Time 计算执行一次排序的耗时
func (s *SortCompare) Time(t SortType, arr []int) int64 {
	start := time.Now()
	switch t {
	case TypeSelection:
		SelectionSort(arr)
	case TypeInsertion:
		InsertSort(arr)
	case TypeQuick:
		QuickSort(arr)
	case TypeQuick3Way:
		QuickSort3Way(arr)
	case TypeShell:
		ShellSort(arr)
	case TypeMergeUB:
		NewMergeSort().MergeUBSort(arr)
	case TypeMergeBU:
		NewMergeSort().MergeBUSort(arr)
	case TypeHeapSort:
		HeapSort(arr)
	}

	return time.Now().Sub(start).Nanoseconds()
}

// TimeRandomInput 随机生成N个数组，每个数组长度为T，完成排序的耗时
func (s *SortCompare) TimeRandomInput(sortType SortType) int64 {
	var total int64
	arr := make([]int, s.n, s.n)
	for i := 0; i < s.t; i++ {
		for j := 0; j < s.n; j++ {
			arr[j] = rand.Int()
		}
		total += s.Time(sortType, arr)
	}
	return total
}

// Compare 随机生成N个数组，每个数组长度为T，完成排序的耗时
func (s *SortCompare) Compare(sortTypes ...SortType) {
	fmt.Printf("每个算法执行 %d 次排序，每次排序 %d 个元素 \n\n", s.t, s.n)

	var baseTotal int64
	m := make(map[SortType]int64)

	for _, sortType := range sortTypes {
		m[sortType] = s.TimeRandomInput(sortType)
		if s.base == sortType {
			baseTotal = m[sortType]
		}
	}
	for i, sortType := range sortTypes {
		total := m[sortType]
		fmt.Printf("算法%v耗时： %v ns | 单次耗时： %v ns | 达到基准算法：%.2f %% 的性能 \n",
			i+1,
			total,
			total/int64(s.t),
			float64(100*(baseTotal))/float64(total),
		)
	}
}
