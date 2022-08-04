package others

import (
	"go_learning/datastructures/st"
	"strings"
)

// LookUpIndex 索引（及反向索引）查找
type LookUpIndex struct {
	keysSt  *st.SeparateChainingHashST[string, string] // 正向索引 k -> v
	valueSt *st.SeparateChainingHashST[string, string] // 反向索引 v -> k
}

func NewLookUpIndex(cap int) *LookUpIndex {
	return &LookUpIndex{
		keysSt:  st.NewSeparateChainingHashST[string, string](cap),
		valueSt: st.NewSeparateChainingHashST[string, string](cap),
	}
}

func (l *LookUpIndex) Add(data []string, sep string) {
	for i := 0; i < len(data); i++ {
		arr := strings.Split(data[i], sep)
		if !l.keysSt.Contains(arr[0]) {
			l.keysSt.Put(arr[0], arr[1])
		}
		if !l.valueSt.Contains(arr[1]) {
			l.valueSt.Put(arr[1], arr[0])
		}
	}
}

func (l *LookUpIndex) GetVal(key string) (val string) {
	return l.keysSt.Get(key)
}

func (l *LookUpIndex) GetKey(val string) (key string) {
	return l.valueSt.Get(val)
}
