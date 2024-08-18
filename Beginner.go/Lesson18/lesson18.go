package main

import "fmt"

func main() {
	arr := [...]int{2, 4, 6, 8, 10}
	sum := 0

	for i := 0; i <= 4; i++ { // 配列のインデックス番号は0から始まるため、0~4番まで繰り返し処理
		sum += arr[i]
	}
	fmt.Println(sum)
}
