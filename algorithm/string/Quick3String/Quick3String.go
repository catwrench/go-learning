package Quick3String

// Quick3String 三向切分的字符串排序
// 按待排序字符串数组进行切分：
// 一个数组包含所有首字母小于切分字符的字符串
// 一个数组包含所有首字母等于切分字符的字符串
// 一个数组包含所有首字母大于切分字符的字符串
type Quick3String struct {
}

func Quick3StringSort(data []string) {
	q3s := &Quick3String{}
	q3s.sort(data, 0, len(data)-1, 0)
}

func (q *Quick3String) sort(data []string, lo, hi, d int) {
	if hi <= lo {
		return
	}

	lt, gt := lo, hi
	current := q.charAt(data[lo], d)

	i := lo + 1
	for i <= gt {
		next := q.charAt(data[i], d)
		if next < current {
			data[lt], data[i] = data[i], data[lt]
			lt++
			i++
		} else if next > current {
			data[gt], data[i] = data[i], data[gt]
			gt--
		} else {
			i++
		}
	}

	q.sort(data, lo, lt-1, d)
	if current > 0 { // 只有首字母相同的子数组，需要递归比较下一个字符
		q.sort(data, lt, gt, d+1)
	}
	q.sort(data, gt+1, hi, d)
}

func (q *Quick3String) charAt(str string, d int) int {
	if d < len(str) {
		return int(str[d])
	} else {
		return -1
	}
}
