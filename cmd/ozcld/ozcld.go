package main

import (
	cld "github.com/mikoto2000/ozcld"
	"fmt"
	"os"
)

func main() {
	// 第一引数で渡されたファイルを読み込んでパース
	file, _ := os.Open(os.Args[1])
	fmt.Println(cld.Parse(file))
}
