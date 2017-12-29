package main

import (
	cld "github.com/mikoto2000/ozcld"
	"os"
)

func main() {
	// 第一引数で渡されたファイルを読み込んでパース
	inFile, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	defer inFile.Close()

	// パースに成功したら、第二引数で渡されたファイルに出力
	parseResult := cld.Parse(inFile)

	outFile, err := os.Create(os.Args[2])
	if err != nil {
		panic(err)
	}
	defer outFile.Close()
	outFile.Write(([]byte)(parseResult))
}
