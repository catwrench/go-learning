package st

import (
	"fmt"
	"testing"
)

func TestFrequencyCounter(t *testing.T) {
	c := NewFrequencyCounter()
	c.SetBase(TypeSequentialSearchST) // 比较用的基准数据结构

	sorts := []TypeST{
		TypeSequentialSearchST, // 1、顺序查找符号表（基于链表）
		TypeBinarySearchST,     // 2、二分查找符号表（基于数组): 插入删除元素时，数组移动成本较高
		TypeBST,                // 3、符号表（基于二叉查找树）
		TypeRedBlackBST,        // 4、符号表（基于红黑树）
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
