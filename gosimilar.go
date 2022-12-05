package gosimilar

import (
	"errors"
	"github.com/gounits/gosimilar/similar"
)

func PhraseSimilar(c similar.ISimilar) (score float64, err error) {
	word1, word2 := c.Compare64()

	var (
		w1Len    = len(word1)
		w2Len    = len(word2)
		maxLen   = w1Len
		validLen = 64
	)

	if w1Len == 0 || w2Len == 0 {
		return
	}

	if w1Len > validLen || w2Len > validLen {
		err = errors.New("长度不能超过64")
		return
	}

	if w1Len < w2Len {
		maxLen = w2Len
		word1, word2 = word2, word1
	}

	// 最大值是1<<16个位置
	var peq [0xff]uint64

	// 记录位置索引
	for i, c := range word1 {
		peq[c] |= 1 << i
	}

	// 初始化指标信息
	var (
		pv uint64 = 1<<validLen - 1
		mv uint64 = 0
		ls uint64 = 1 << (maxLen - 1)
		sc        = maxLen
	)

	for _, c := range word2 {
		eq := peq[c]
		xv := eq | mv
		eq |= ((eq & pv) + pv) ^ pv
		mv |= ^(eq | pv)
		pv &= eq
		if (mv & ls) != 0 {
			sc++
		}
		if (pv & ls) != 0 {
			sc--
		}
		mv = (mv << 1) | 1
		pv = (pv << 1) | ^(xv | mv)
		mv &= xv
	}
	score = float64(maxLen-sc) / float64(maxLen)
	return
}

var (
	EnglishSimilar     = similar.NewEnglishCompare
	EnglishDictSimilar = similar.NewEnglishDictCompare
)
