package algorithm

import (
	"fmt"
	"strings"
)

/************************************************************************
// 二分探索 ソート済み配列にのみ使用可能
// まず配列内中央の値を見る。下の配列で言うと5
//
// key=6
//              ↓
// [1, 2, 3, 4, 5, 6, 7, 8, 9]
//
// keyが6（中央の値より大きい）なら中央より右側に探索範囲を絞って中央の値を見る
//
//                    ↓            　　　↓探索範囲を絞ったときのイメージ
// [1, 2, 3, 4, 5, 6, 7, 8, 9]    　　　[6, 7, 8, 9]
//
//keyの方が小さい（6 < 7）ので今度は探索範囲を左側に絞る
//
// [6, 7, 8, 9] -> [6]
// 目的の数字が見つかる！
*************************************************************************/

func intBinarySerch(s []int, key int) int {
	// pl=探索範囲の先頭index pr=探索範囲の末尾index
	pl := 0
	pr := len(s) - 1

	for {
		// pc=探索範囲の中央index
		pc := (pl + pr) / 2

		// pcの要素とkeyが一致したらpc（目的の値のindex）を返す
		if s[pc] == key {
			return pc
		// pcの要素がkeyより小さかったら探索範囲の先頭indexをpc+1にする
		} else if s[pc] < key {
			pl = pc + 1
		// 探索範囲の末尾indexをpc-1にする
		} else {
			pr = pc - 1
		}

		// 先頭indexが末尾indexより大きくなったら目的の値がないのでブレーク
		if pl > pr {
			break
		}
	}

	return -1
}

func intBinarySerchVisualization(s []int, key int) int {
	// pl=探索範囲の先頭index pr=探索範囲の末尾index
	pl := 0
	pr := len(s) - 1

	fmt.Print("idx|")
	for ii:=0; ii < len(s); ii++ {
		fmt.Printf("%4v", ii)
	}
	fmt.Printf("\n---+")
	fmt.Printf(strings.Repeat("-", (4 * len(s) + 2)) + "\n")

	for {
		// pc=探索範囲の中央index
		pc := (pl + pr) / 2

		fmt.Print("   |")
		if pl != pc {
			fmt.Print(strings.Repeat(" ", (pl * 4 + 1)) + "←")
			fmt.Print(strings.Repeat(" ", ((pc - pl) * 4)) + "+")
		} else {
			fmt.Print(strings.Repeat(" ", (pc * 4 + 1)) + "←")
		}
		if pc != pr {
			fmt.Printf(strings.Repeat(" ", ((pr - pc) * 4 - 2)) + "→")
		} else {
			fmt.Printf("→")
		}
		fmt.Printf("\n%3v|", pc)
		for ii:=0; ii < len(s); ii++ {
			fmt.Printf("%4v", s[ii])
		}
		fmt.Print("\n   |\n")

		// pcの要素とkeyが一致したらpc（目的の値のindex）を返す
		if s[pc] == key {
			fmt.Printf("その値はi[%v]にあります。", pc)
			return pc
		// pcの要素がkeyより小さかったら探索範囲の先頭indexをpc+1にする
		} else if s[pc] < key {
			pl = pc + 1
		// 探索範囲の末尾indexをpc-1にする
		} else {
			pr = pc - 1
		}

		// 先頭indexが末尾indexより大きくなったら目的の値がないのでブレーク
		if pl > pr {
			break
		}
	}

	return -1
}