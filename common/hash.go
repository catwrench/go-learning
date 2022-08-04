package common

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"hash/crc32"
)

// HashInt 计算hash 返回 int 值
func HashInt[T any](key T) int {
	// 将key序列化为byte数组，方便使用一种hash算法
	buf := bytes.Buffer{}
	encoder := gob.NewEncoder(&buf)
	err := encoder.Encode(key)
	if err != nil {
		fmt.Println("hash失败:", err)
		return 0
	}
	return HashCRC32(buf.Bytes())
}

// HashCRC32 计算hash
func HashCRC32(data []byte) int {
	v := int(crc32.ChecksumIEEE(data))
	if v >= 0 {
	} else if -v >= 0 {
		v = -v
	} else {
		// v == MinInt, 此时v等于int最小值
		v = 0
	}
	return v
}
