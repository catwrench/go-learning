package sort

// HeapSort 堆排序
/**
  思想：本质是完全二叉树，根节点为最大值或者最小值

  时间复杂度： O(nlogn)
  空间复杂度： O(n)
*/
func HeapSort(arr []int) {
	N := len(arr)
	// 构造堆:构造大顶堆，按顺序出队后即从小到大排序
	// 1、只需要遍历一半的数组就行了，因为我们不用比较大小为1的堆
	// 2、(N-2)/2 也是同样的原因，最后一位无需下沉操作
	for i := N/2 - 1; i >= 0; i-- {
		// 较小元素下沉到底部
		sink(arr, i, N)
	}
	// 使用下沉排序，等同于优先队列的 delMax, 队尾此时有序
	for N > 0 {
		arr[0], arr[N-1] = arr[N-1], arr[0]
		N--
		sink(arr, 0, N)
	}
}

// N为数组长度，最大索引等于N-1
func sink(arr []int, idx int, N int) {
	doubleIdx := idx*2 + 1
	// temp 用于存储当前比较用的元素，每次比较其实都是这个元素和子级元素比较，用于减少数组访问次数
	// ps:目前来看数据基本有序后，比较次数不多，内存分配代价反而较高，所以性能变化不明显
	temp := arr[idx]
	for doubleIdx < N {
		// 选取较大的子节点，如果存在右子节点的话
		if doubleIdx+1 < N && arr[doubleIdx] < arr[doubleIdx+1] {
			doubleIdx++
		}
		// 下沉
		if temp < arr[doubleIdx] {
			arr[idx], arr[doubleIdx] = arr[doubleIdx], arr[idx]
		} else {
			break
		}
		idx = doubleIdx
		doubleIdx = 2*idx + 1
	}
}
