package sort

// HSort = 2 : 1, 3, 7, 15, 31, 63
// HSort = 3 : 1, 4, 13, 40, 121, 364
const HSort int = 3 // h有序，数组任意间隔h个元素的元素都是有序的，h大小对性能影响不大

// ShellSort 希尔排序
/**
  思想：设置希尔增量每次折半，逐步分组进行粗调，最后进行插入排序。

  时间复杂度：平均：O(nlogn) 最坏：O(n2) + 检查的时间
  空间复杂度： O(1)
  稳定性： 不稳定
  ps: 实际上 希尔增量h 的大小对性能影响不大，不用刻意选择
*/
func ShellSort(arr []int) {
	N := len(arr)
	for h := N / HSort; h >= 1; h /= HSort {
		// 每次遍历h到末尾，h减小一次，直到h=1,i遍历完整个数组，排序也就完成了
		for i := h; i < N; i++ {
			// 不断和前面的相差h的元素进行比较，跳跃交换
			for j := i; j >= h && arr[j] < arr[j-h]; j -= h {
				arr[j], arr[j-h] = arr[j-h], arr[j]
			}
		}
	}
}
