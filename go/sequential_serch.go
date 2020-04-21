package algorithm

// 逐次探索（線形探索）
// 先頭から順番に値を比較して目的のものを探す
// 見つかった場合はその場所のインデックスを返す
// 目的のものがなかったら-1を返している

func intSequentialSerch(s []int, key int) int {
	for idx, value := range s {
		if value == key {
			return idx
		}
	}
	return -1
}