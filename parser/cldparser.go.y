%{
package main

import (
    cld "../../ozcld"
    "fmt"
//    "github.com/k0kubun/pp"
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
    classes []*cld.Class
    fields *cld.Fields
    methods *cld.Methods
    namespace *cld.Namespace
    namespaces []*cld.Namespace
    relation *cld.Relation
    relations []*cld.Relation
    word *Token
    words []*Token
}

// 単語の繰り返し
%type<words> words

// 単語
%token<word> WORD


%type<classDiagram> classDiagram
%type<class> class
%type<classes> classes

%type<fields> fields
%type<methods> methods
// end of member
%token<word> EOM

%token<word> EQUAL
%type<word> divisionMarker
%type<namespace> namespace
%type<namespaces> namespaces

%type<relation> relation
%type<relations> relations

%token<word> LABEL_CLD
%token<word> LABEL_CLASS
%token<word> LABEL_NAMESPACE
%token<word> START_BLOCK
%token<word> END_BLOCK

// 関連
%token<word> HYPHEN
%token<word> GT

%%

classDiagram
    : LABEL_CLD WORD START_BLOCK namespaces classes relations END_BLOCK
    {
        //pp.Println("classDiagram")

        namespaces := $4
        classes := $5

        $$ = cld.CreateClassDiagram($2.Literal, namespaces, classes, $6)

        yylex.(*Lexer).Result = $$
    }

classes
    :
    {
        $$ = []*cld.Class{}
    }
    | class
    {
        $$ = []*cld.Class{$1}
    }
    | classes class
    {
        $$ = append($1, $2)
    }

fields
    :
    {
        $$ = &cld.Fields{}
    }
    | words EOM
    {
        fields := &cld.Fields{}
        field := cld.CreateFieldFromString(wordsToString($1))
        fields.Add(field)
        $$ = fields
    }
    | fields words EOM
    {
        field := cld.CreateFieldFromString(wordsToString($2))
        $1.Add(field)
        $$ = $1
    }

methods
    :
    {
        $$ = &cld.Methods{}
    }
    | words EOM
    {
        methods := &cld.Methods{}
        method := cld.CreateMethodFromString(wordsToString($1))
        methods.Add(method)
        $$ = methods
    }
    | methods words EOM
    {
        method := cld.CreateMethodFromString(wordsToString($2))
        $1.Add(method)
        $$ = $1
    }

class
    : LABEL_CLASS WORD START_BLOCK words divisionMarker fields divisionMarker methods END_BLOCK
    {
        //pp.Println("class")

        $$ = cld.CreateClass(wordsToString($4), $2.Literal, $6, $8)

        yylex.(*Lexer).Result = $$
    }
    | LABEL_CLASS WORD START_BLOCK divisionMarker fields divisionMarker methods END_BLOCK
    {
        //pp.Println("class")

        $$ = cld.CreateClass("", $2.Literal, $5, $7)

        yylex.(*Lexer).Result = $$
    }

divisionMarker
    : EQUAL EQUAL

namespaces
    :
    {
        $$ = []*cld.Namespace{}
    }
    | namespace
    {
        $$ = []*cld.Namespace{$1}
    }
    | namespaces namespace
    {
        $$ = append($1, $2)
    }

namespace
    : LABEL_NAMESPACE WORD START_BLOCK namespaces classes END_BLOCK
    {
        //pp.Println("namespace")

        $$ = cld.CreateNamespace($2.Literal, $5, $4)

        yylex.(*Lexer).Result = $$
    }

relations
    :
    {
        $$ = []*cld.Relation{}
    }
    | relation
    {
        $$ = []*cld.Relation{$1}
    }
    | relations relation
    {
        relations := append($1, $2)

        $$ = relations

        yylex.(*Lexer).Result = $$
    }

relation
    : words HYPHEN GT words
    {
        $$ = cld.CreateRelation("", cld.RELATION_NORMAL, wordsToString($1), wordsToString($4), "", "")
    }

// 単語(WORD)の繰り返しルール
words
    : WORD
    {
        tokens := []*Token{$1}

        $$ = tokens

        yylex.(*Lexer).Result = $$
    }
    | HYPHEN /* word は、 '-' も許容したいので、こんな感じでルールを追加。 */
    {
        tokens := []*Token{$1}

        $$ = tokens

        yylex.(*Lexer).Result = $$
    }
    | words HYPHEN /* word は、 '-' も許容したいので、こんな感じでルールを追加。 */
    {
        tokens := append($1, $2)

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
    } else if l.TokenText() == ";" {
        token = EOM
    } else if l.TokenText() == "-" {
        token = HYPHEN
    } else if l.TokenText() == ">" {
        token = GT
    } else if token != scanner.EOF {
        // WORD は、yacc 宣言部で書いた token<xxx> のどれか
        token = WORD
    }

    // 解析結果を lval に詰める。
    // WORD を Token 型に変更したので、そのように修正。
    lval.word = &Token{Type: token, Literal: l.TokenText()}

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
    //pp.Println(l.Result)

    fmt.Println(l.Result.(cld.Dot).ToDot())
}

func wordsToString(words []*Token) string {
    str := ""

    for _, v := range words {
        str = str + v.Literal
    }

    return str;
}
