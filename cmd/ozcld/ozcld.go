package main

import (
	cld "github.com/mikoto2000/ozcld"
	"flag"
	"fmt"
	"os"
)

func printUsage() {
	fmt.Println("Usage: ozcld [OPTIONS] INFILE")
	flag.PrintDefaults()
}

func printStdout(message string) {
	fmt.Println(message)
}

func main() {
	flag.Usage = printUsage

	// 引数パース
	outFilePath := flag.String("o", "", "output file(default: stdout)")
	isShowHelp := flag.Bool("h", false, "show help(default: false)")
	flag.Parse()

	// コマンドライン引数不正の場合、 Usage を出力して終了
	if flag.NArg() != 1 {
		flag.Usage()
		os.Exit(1)
	}

	// ヘルプ表示
	if *isShowHelp {
		flag.Usage()
		os.Exit(0)
	}

	// 入力ファイルを読み込んでパース
	inFilePath := flag.Arg(0)
	inFile, err := os.Open(inFilePath)
	if err != nil {
		panic(err)
	}
	defer inFile.Close()

	// パースに成功したら、-o オプションで渡されたファイルに出力
	// -o が無ければ、標準出力に出力する。
	parseResult, parseErrors := cld.Parse(inFile)

	if parseErrors != nil {
		for _, e := range parseErrors {
			errorMessage := fmt.Sprintf("ERROR: %v\n", e)
			fmt.Fprintf(os.Stderr, errorMessage)
		}
		os.Exit(1)
	}

	if *outFilePath == "" {
		printStdout(parseResult)
	} else {
		outFile, err := os.Create(*outFilePath)
		if err != nil {
			panic(err)
		}
		defer outFile.Close()
		outFile.Write(([]byte)(parseResult))
	}
}
