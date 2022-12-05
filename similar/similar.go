package similar

/*
ISimilar 相似算法的实现接口

	每一个uint8代表的是该信息的位置索引
*/
type ISimilar interface {
	// Compare64 获取要比较的词向量
	// 并且长度不能超过64个
	Compare64() (c1 []uint8, c2 []uint8)
}
