package common

type KVNode[K comparable, V any] struct {
	key   K
	value V
	next  *KVNode[K, V]
}

func NewKVNode[K comparable, V any](k K, v V) *KVNode[K, V] {
	return &KVNode[K, V]{
		key:   k,
		value: v,
	}
}

func (n *KVNode[K, V]) GetValue() V {
	return n.value
}

func (n *KVNode[K, V]) GetKey() K {
	return n.key
}

func (n *KVNode[K, V]) SetKey(k K) *KVNode[K, V] {
	n.key = k
	return n
}

func (n *KVNode[K, V]) SetValue(v V) *KVNode[K, V] {
	n.value = v
	return n
}

func (n *KVNode[K, V]) AddNext(node *KVNode[K, V]) *KVNode[K, V] {
	n.next = node
	return n
}

func (n *KVNode[K, V]) HasNext() bool {
	return n.next != nil
}

func (n *KVNode[K, V]) Next() *KVNode[K, V] {
	return n.next
}
