package sort

import (
	"fmt"
	"testing"
)

func TestSortCompare(t *testing.T) {
	c := NewSortCompare()
	c.SetBase(TypeSelection) // 比较用的基准算法

	sorts := []SortType{
		TypeSelection, // 1 选择排序：每趟选择一个最小的元素，无法为下次循环减少比较
		TypeInsertion, // 2 插入排序：适合数量较小的排序, 不同算法 7-15 个元素 的排序速度优于其他算法。（因为没有递归的开销）
		TypeQuick,     // 3 快速排序：将元素插入到一个位置，使得左边的元素都比他小，右边的比他大。（基准元素选择很重要）
		TypeQuick3Way, // 4 快速排序（三向切分）：适合重复元素较多的排序
		TypeShell,     // 5 希尔排序：切分为多个数组，间隔n个元素进行比较，一次交换的位置更多。（相当于插入排序的优化版本）
		TypeMergeUB,   // 6 归并排序（自顶向下）：排序很快，缺点需要一倍辅助内存
		TypeMergeBU,   // 7 归并排序（自底向上）：自底向上少了分拆数组的过程，更快一点
		TypeHeapSort,  // 8 堆排序（大顶堆）
	}

	fmt.Println("\n--------------- 短数组 ------------------")
	c.SetN(15)    // 随机数组长度
	c.SetT(10000) // 执行次数
	c.Compare(sorts...)
	fmt.Println("\n--------------- 中数组 ------------------")
	c.SetN(2000) // 随机数组长度
	c.SetT(100)  // 执行次数
	c.Compare(sorts...)
	fmt.Println("\n--------------- 长数组 ------------------")
	c.SetN(50000) // 随机数组长度s
	c.SetT(3)     // 执行次数
	c.Compare(sorts...)
}
