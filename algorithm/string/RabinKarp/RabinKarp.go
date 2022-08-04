package RabinKarp

import (
	"crypto/rand"
	"go_learning/common"
)

// RabinKarp 指纹字(散列)符串查找算法
type RabinKarp struct {
	pat     string // 模式串
	patHash int64  // 模式串hash
	M       int    // 模式串长度
	R       int    // 字符表 字符数
	Q       int64  // 一个较大的素数
	RM      int64  // R^(M-1)%Q
}

func NewRabinKarp(pat string) *RabinKarp {
	res := &RabinKarp{}
	R := common.R256
	M := len(pat)
	Q := res.longRandomPrime()
	RM := int64(1)
	for i := 1; i <= M-1; i++ {
		RM = (int64(R) * RM) % Q // 用于减去第一个数字时的计算
	}

	res.pat = pat
	res.M = M
	res.R = R
	res.Q = Q
	res.RM = RM
	res.patHash = res.hash(pat, M)
	return res
}

func (r *RabinKarp) Search(txt string) int {
	N := len(txt)
	textHash := r.hash(txt, r.M)

	// 首位命中，直接返回
	if (r.patHash == textHash) && r.check(txt, 0) {
		return 0
	}

	for i := r.M; i < N; i++ {
		// 计算hash,去掉首位，加上末尾
		textHash = (textHash + r.Q - r.RM*int64(txt[i-r.M])%r.Q) % r.Q
		textHash = (textHash*int64(r.R) + int64(txt[i])) % r.Q

		// 命中,计算命中时的索引是否和模式串匹配
		index := i - r.M + 1
		if (r.patHash == textHash) && r.check(txt, index) {
			return index
		}
	}
	return N // 未命中
}

// 拉斯维加斯算法检查 模式 是否与text[i...i-M+1]匹配
func (r *RabinKarp) check(txt string, i int) bool {
	for j := 0; j < r.M; j++ {
		if r.pat[j] != txt[i+j] {
			return false
		}
	}
	return true // 蒙特卡洛算法始终返回true
}

func (r *RabinKarp) hash(key string, M int) (res int64) {
	// 计算key[0...M-1]的散列值
	for i := 0; i < M; i++ {
		res = (int64(r.R)*res + int64(key[i])) % r.Q
	}
	return
}

func (r *RabinKarp) longRandomPrime() int64 {
	res, _ := rand.Prime(rand.Reader, 32)
	return res.Int64()
}
