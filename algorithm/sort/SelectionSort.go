package sort

// SelectionSort 选择排序
/**
  思想：每一次选出最小者直接交换到左侧，省出了多余的元素交换

  时间复杂度： 平均：O(n2) 最坏：O(n2)
  空间复杂度： O(1)
  稳定性： 不稳定。当数列包含多个值相等的元素时，选择排序有可能打乱它们的原有顺序。
*/
func SelectionSort(arr []int) []int {

	var min int
	for i := 0; i < len(arr); i++ {
		// 选择
		min = i
		for j := i + 1; j < len(arr); j++ {
			if arr[min] > arr[j] {
				min = j
			}
		}
		// 交换
		if min != i {
			arr[i], arr[min] = arr[min], arr[i]
		}
	}
	return arr
}
