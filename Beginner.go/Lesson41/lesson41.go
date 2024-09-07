package main

import "fmt"

func main() {
	// スライスの宣言と初期化
	var sl1 []int              // 空のスライスを宣言
	sl1 = append(sl1, 1, 2, 3) // スライスに空の要素を追加
	fmt.Println("スライス:", sl1)  // 出力: スライス: [1 2 3]

	sl2 := []int{1, 2, 3}
	sl2 = append(sl2, 4, 5)   // // 新しい要素4, 5を追加
	fmt.Println("スライス:", sl2) // 出力: スライス: [1 2 3 4 5]

	arr := [5]int{1, 2, 3, 4, 5} // 配列
	sl3 := arr[1:4]              // 配列の部分をスライスとして切り取る
	fmt.Println("スライス:", sl3)    // 出力: スライス: [2 3 4]

	sl4 := make([]int, 3)     // 長さ3のスライスを作成（デフォルト値は0）
	fmt.Println("スライス:", sl4) // 出力: スライス: [0 0 0]

	s := []int{1, 2, 3, 4, 5}
	s = append(s[:2], s[3:]...) // インデックス2の要素を削除（s[:2]要素数1,2を取得、s[:3]要素数4,5を取得）
	fmt.Println("スライス:", s)     // 出力: スライス: [1 2 4 5]
}
