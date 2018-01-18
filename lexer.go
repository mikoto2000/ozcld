package ozcld

import (
	"errors"
	"fmt"
	//    "github.com/k0kubun/pp"
	"os"
	"strings"
	"text/scanner"
)

type Result interface{}

// トークンの情報を保持するための構造体を追加
type Token struct {
	// トークンの種類を表す数値
	Type int

	// 字句解析で分かち書きされた文字列(この表現でよいのかしら？)
	Literal string
}

// yyLexer インタフェースを実装したオブジェクトを作成する。
// こいつの Lex メソッドが、字句解析を担当する。
type Lexer struct {
	scanner.Scanner

	ErrorList []error

	// パースの結果を格納するオブジェクトを追加
	Result Result
}

func (this *Lexer) Parse() {
	yyParse(this)
}

// Lex : トークンを一つ読み進める。
// この実装では、スキャンした結果を何でもかんでも WORD として解釈する。
func (l *Lexer) Lex(lval *yySymType) int {
	token := int(l.Scan())

	//pp.Println(l.TokenText())

	// EOF 以外はすべて WORD
	if l.TokenText() == "classdiagram" {
		token = LABEL_CLD
	} else if l.TokenText() == "class" {
		token = LABEL_CLASS
	} else if l.TokenText() == "=" {
		token = EQUAL
	} else if l.TokenText() == "namespace" {
		token = LABEL_NAMESPACE
	} else if l.TokenText() == "{" {
		token = START_BLOCK
	} else if l.TokenText() == "}" {
		token = END_BLOCK
	} else if l.TokenText() == ":" {
		token = COLON
	} else if l.TokenText() == ";" {
		token = EOM
	} else if l.TokenText() == "." {
		token = DOT
	} else if l.TokenText() == "-" {
		token = HYPHEN
	} else if l.TokenText() == ">" {
		token = GT
	} else if l.TokenText() == "|" {
		token = PIPE
	} else if token != scanner.EOF {
		// WORD は、yacc 宣言部で書いた token<xxx> のどれか
		token = WORD
	}

	// 解析結果を lval に詰める。
	// WORD を Token 型に変更したので、そのように修正。
	lval.word = &Token{Type: token, Literal: strings.Trim(l.TokenText(), "'\"")}

	// 解釈したトークンの種類を返却する
	return token
}

/* エラー情報を追加 */
func (l *Lexer) Error(e string) {
	errorMessage := fmt.Sprintf("'%v' <Line %v, Column %v> in %v", l.TokenText(), l.Pos().Line, l.Pos().Column, l.Filename)

	l.ErrorList = append(l.ErrorList, errors.New(errorMessage))
}

// トークンのリストを結合し、ひとつの文字列にする
func wordsToString(words []*Token) string {
	str := ""

	for _, v := range words {
		str = str + v.Literal
	}

	return str
}

func Parse(file *os.File) (string, []error) {
	// 自作 Lexer 作成
	l := new(Lexer)
	l.Filename = file.Name()
	l.ErrorList = make([]error, 0)

	// 第一引数で渡されたファイルを読み込んでパースするように設定
	l.Init(file)

	// パース実行
	l.Parse()

	// パース結果表示
	// こういう使い方すると、もう l って Lexer の範疇超えてる気がするけれどどうなのだろうか？
	//pp.Println(l.Result)

	if l.Result == nil || len(l.ErrorList) != 0 {
		return "", l.ErrorList
	}

	return l.Result.(Dot).ToDot(), nil
}
