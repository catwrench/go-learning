package sort

// QuickSort3Way 三向切分的快速排序
// 使用场景：适合存在大量重复元素的排序
func QuickSort3Way(arr []int) {
	q3WaySort(arr, 0, len(arr)-1)
	return
}

func q3WaySort(arr []int, low, high int) {
	if low < high {

		// [lo ... lt-1] 小于 pivot
		// [gt+1 ... hi] 大于 pivot
		// [lt ... i-1] 等于 pivot
		// [i ... gt] 之间的元素大小还未确定，整个循环就是在遍历 i...gt之前的元素，将其分配到上面三个数组
		pivot := arr[getPivotIdx(arr, low, high)]
		// 这里是使用的3取中，所以是从 arr[low] 开始计算的，i=low 开始计算
		// 通常算法默认以 arr[low] 为基准，i=low+1 开始计算
		lt, i, gt := low, low, high
		for i <= gt {
			if arr[i] < pivot {
				arr[i], arr[lt] = arr[lt], arr[i]
				i++
				lt++
			} else if arr[i] > pivot {
				arr[i], arr[gt] = arr[gt], arr[i]
				gt--
			} else {
				i++
			}
		}

		q3WaySort(arr, low, lt-1)
		q3WaySort(arr, gt+1, high)
	}
}

// partition3 切分数组，返回支点索引
func partition3(arr []int, low, high int) int {
	if len(arr) <= 1 {
		return low
	}

	// 获取基准值,三取中
	startIndex := getPivotIdx(arr, low, high)
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
	// 将中间值替换为 pivot就行
	// arr[low], arr[startIndex] = arr[startIndex], arr[low]
	arr[low] = pivot
	return low
}

// 获取基准值，这里采用三取中
func getPivotIdx(arr []int, low, high int) int {
	if high-low < 2 {
		return low
	}
	pivotIdx := (low + high) / 2
	// 低高交换，保证低小
	if arr[low] > arr[high] {
		arr[low], arr[high] = arr[high], arr[low]
	}
	// 中高交换，保证中小
	if arr[pivotIdx] > arr[high] {
		arr[pivotIdx], arr[high] = arr[high], arr[pivotIdx]
	}
	// 低中交换，保证低小
	if arr[low] > arr[pivotIdx] {
		arr[low], arr[pivotIdx] = arr[pivotIdx], arr[low]
	}
	return pivotIdx
}
