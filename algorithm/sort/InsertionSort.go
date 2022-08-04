package sort

// InsertSort 插入排序
/**
  思想：维护一个有序区，把元素一个一个插入到有序区的适当位置，直到所有元素有序为止

  时间复杂度： O(n2)
  空间复杂度： O(1)
  稳定性： 稳定
  ps: 插入排序适合『部分有序』和『小规模排序』
*/
func InsertSort(arr []int) {
	N := len(arr)
	for i := 1; i < N; i++ {
		// 将目标元素，插入到有序数组的合适位置
		for j := i; j > 0 && arr[j] < arr[j-1]; j-- {
			arr[j], arr[j-1] = arr[j-1], arr[j]
		}
	}
}
