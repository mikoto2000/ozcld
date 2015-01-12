%{
package main

import (
    cld "../../ozcld"
    "fmt"
    "github.com/k0kubun/pp"
    "os"
    "text/scanner"
)

type Result interface {}

// トークンの情報を保持するための構造体を追加
type Token struct {
    // トークンの種類を表す数値
    Type int

    // 字句解析で分かち書きされた文字列(この表現でよいのかしら？)
    Literal string
}

%}

%union{
    classDiagram *cld.ClassDiagram
    class *cld.Class
    namespace *cld.Namespace
    word Token
    words []Token
}

// 単語の繰り返し
%type<words> words

// 単語
%token<word> WORD


%type<classDiagram> classDiagram
%type<class> class
%type<namespace> namespace

%token<word> LABEL_CLD
%token<word> LABEL_CLASS
%token<word> LABEL_NAMESPACE
%token<word> START_BLOCK
%token<word> END_BLOCK

%%

classDiagram
    : LABEL_CLD WORD START_BLOCK namespace class END_BLOCK
    {
        pp.Println("classDiagram")

        namespaces := []*cld.Namespace{$4}
        classes := []*cld.Class{$5}

        $$ = cld.CreateClassDiagram($2.Literal, namespaces, classes, nil)

        yylex.(*Lexer).Result = $$
    }

classDiagramItem
    : namespace
    | class

class
    : LABEL_CLASS WORD START_BLOCK words END_BLOCK
    {
        pp.Println("class")

        $$ = cld.CreateClassFromDefs("", $2.Literal, nil, nil)

        yylex.(*Lexer).Result = $$
    }

namespace
    : LABEL_NAMESPACE WORD START_BLOCK words END_BLOCK
    {
        pp.Println("namespace")

        $$ = cld.CreateNamespace($2.Literal, nil, nil)

        yylex.(*Lexer).Result = $$
    }

// 単語(WORD)の繰り返しルール
words
    : WORD {
        tokens := []Token{$1}

        $$ = tokens

        yylex.(*Lexer).Result = $$
    }
    | words WORD {
        tokens := append($1, $2)

        $$ = tokens

        yylex.(*Lexer).Result = $$
    }
%%

// yyLexer インタフェースを実装したオブジェクトを作成する。
// こいつの Lex メソッドが、字句解析を担当する。
type Lexer struct {
    scanner.Scanner

    // パースの結果を格納するオブジェクトを追加
    Result Result
}


// Lex : トークンを一つ読み進める。
// この実装では、スキャンした結果を何でもかんでも WORD として解釈する。
func (l *Lexer) Lex(lval *yySymType) int {
    token := int(l.Scan())

    pp.Println(l.TokenText())

    // EOF 以外はすべて WORD
    if l.TokenText() == "classdiagram" {
        token = LABEL_CLD
    } else if l.TokenText() == "class" {
        token = LABEL_CLASS
    } else if l.TokenText() == "namespace" {
        token = LABEL_NAMESPACE
    } else if l.TokenText() == "{" {
        token = START_BLOCK
    } else if l.TokenText() == "}" {
        token = END_BLOCK
    } else if token != scanner.EOF {
        // WORD は、yacc 宣言部で書いた token<xxx> のどれか
        token = WORD
    }

    // 解析結果を lval に詰める。
    // WORD を Token 型に変更したので、そのように修正。
    lval.word = Token{Type: token, Literal: l.TokenText()}

    // 解釈したトークンの種類を返却する
    return token
}

/* エラーはとりあえずパニック */
func (l *Lexer) Error(e string) {
    panic(e)
}


func main() {
    // 自作 Lexer 作成
    l := new(Lexer)

    // 第一引数で渡されたファイルを読み込んでパースするように設定
    file, _ := os.Open(os.Args[1])
    l.Init(file)

    // パース実行
    yyParse(l)

    // パース結果表示
    // こういう使い方すると、もう l って Lexer の範疇超えてる気がするけれどどうなのだろうか？
    pp.Println(l.Result)

    fmt.Println(l.Result.(cld.Dot).ToDot())
}
