package sort

// QuickSort 快速排序
/**
  思想：在快速排序中，元素的比较和交换是从两端向中间进行的，较大的元素一轮就能交换到后面的位置，
	而较小的元素一轮就能交换到前面的位置，元素每次移动的距离较远，所以比较次数和移动次数较少，速度较快。
	1、在待排序的元素任取一个元素作为基准（通常选第一个元素，但最好的方法是从待排序元素中随机选取一个为基准），称为基准元素（pivot）
	2、将待排序的元素进行分区，比基准元素大的元素放在它的右边，比基准元素小的放在它的左边
	3、对左右两个分区重复以上步骤，直到所有的元素都是有序的

  时间复杂度：平均：O(nlogn) 最坏：O(n2)
  空间复杂度：平均：O(logn) 最坏：O(n)
  稳定性： 不稳定，因为基准元素的比较和交换是跳跃进行的
*/
func QuickSort(arr []int) {
	qSort(arr, 0, len(arr)-1)
	return
}

func qSort(arr []int, low, high int) {
	if low < high {
		pivot := partition(arr, low, high)
		// pivot := partition3(arr, low, high)  	// 获取基准值,三取中方式获取基准更平衡
		qSort(arr, low, pivot-1)
		qSort(arr, pivot+1, high)
	}
}

// partition 切分数组，返回支点索引
func partition(arr []int, low, high int) int {
	if len(arr) <= 1 {
		return low
	}

	// 获取基准值
	startIndex := low
	pivot := arr[startIndex]

	for low < high {
		// 因为选择的是低位做支点，所以需要从高位开始移动，否则在交换的过程中，无法保证顺序
		// 高位指针左移
		for low < high && arr[high] >= pivot {
			high--
		}
		// 低位指针右移
		for low < high && arr[low] <= pivot {
			low++
		}
		// 交换
		if low < high {
			arr[low], arr[high] = arr[high], arr[low]
		}
	}
	// 交换基准值和中间相等的值
	arr[low], arr[startIndex] = arr[startIndex], arr[low]
	return low
}
