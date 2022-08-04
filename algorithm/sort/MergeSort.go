package sort

// MergeSort 归并排序
/**
  思想：最小分组比较，依次合并
	1、创建一个大集合，长度为两个小集合之和
	2、从左到右逐一比较两个小集合中的元素，把较小的优先放入大集合

  时间复杂度： O(nlogn)
  空间复杂度： O(n)
*/
type MergeSort struct {
	N       int   // 数组长度
	storage []int // 归并时使用的辅助数组
}

func NewMergeSort() *MergeSort {
	return &MergeSort{}
}

func (m *MergeSort) initArr(arr []int) *MergeSort {
	m.N = len(arr)
	m.storage = make([]int, m.N, m.N)
	return m
}

// MergeBUSort 归并排序，自底向上
func (m *MergeSort) MergeBUSort(arr []int) {
	m.initArr(arr)
	for subSize := 1; subSize < m.N; subSize *= 2 { // subSize:子数组大小，每次归并翻倍
		// 归并，最后一次合并的时候，可能高位数组不足一个subSize的长度，需要处理数组索引
		for i := 0; i < m.N-subSize; i += subSize * 2 {
			if i+subSize*2 > m.N {
				m.merge(arr, i, i+subSize-1, m.N-1)
			} else {
				m.merge(arr, i, i+subSize-1, i+2*subSize-1)
			}
		}
	}
}

// MergeUBSort 归并排序，自顶向下
func (m *MergeSort) MergeUBSort(arr []int) {
	m.initArr(arr)
	m.mergeUBSort(arr, 0, m.N-1)
}

// mergeUBSort 自顶向下，递归，对左右子数组进行排序
func (m *MergeSort) mergeUBSort(arr []int, low, high int) {
	if low >= high {
		return
	}
	mid := (low + high) / 2
	m.mergeUBSort(arr, low, mid)
	m.mergeUBSort(arr, mid+1, high)
	m.merge(arr, low, mid, high)
}

// merge 合并，使用辅助数组完成排序
// low: 左边数组开始位置
// mid: 左边数组结束位置
// high: 右边数组结束位置（右边数组可能不足一个子数组长度）
func (m *MergeSort) merge(arr []int, low, mid, high int) {
	// 复制元素到辅助数组
	for i := low; i < high; i++ {
		m.storage[i] = arr[i]
	}

	// 两个子数组的起始索引，分别从 low 和 mid+1 开始
	subIndexL, subIndexR := low, mid+1

	// 比较排序辅助数组元素大小，放到原始数组对应位置
	for i := low; i <= high; i++ {
		if subIndexR > high { // 索引到达右边数组最后一位，就只取边数组剩余元素
			arr[i] = m.storage[subIndexL]
			subIndexL++
		} else if subIndexL > mid { // 索引到达左边数组最后一位，就只取右边数组剩余元素
			arr[i] = m.storage[subIndexR]
			subIndexR++
		} else if m.storage[subIndexL] > m.storage[subIndexR] { // 未到两个数组边界时，对比两个数组对应索引位置的元素大小，优先取小的
			arr[i] = m.storage[subIndexR]
			subIndexR++
		} else {
			arr[i] = m.storage[subIndexL]
			subIndexL++
		}
	}
}
