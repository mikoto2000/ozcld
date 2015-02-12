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
    classMetaDatas map[string]string
    fields cld.Fields
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
%type<word> validSigne


%type<classDiagram> classDiagram
%type<class> class
%type<classes> classes

%type<classMetaDatas> classMetaDatas
%type<fields> fields
%type<methods> methods
// end of member
%token<word> EOM

%token<word> EQUAL
%type<word> divisionMarker
%type<namespace> namespace
%type<namespaces> namespaces

%type<relation> relation
%type<relation> relation_normal
%type<relation> relation_inherit
%type<relation> relation_implement
%type<relations> relations

%token<word> LABEL_CLD
%token<word> LABEL_CLASS
%token<word> LABEL_NAMESPACE
%token<word> START_BLOCK
%token<word> END_BLOCK

// 関連
%token<word> HYPHEN
%token<word> GT
%token<word> PIPE
%token<word> DOT
%token<word> COLON

%%

classDiagram
    : LABEL_CLD words START_BLOCK namespaces classes relations END_BLOCK
    {
        //pp.Println("classDiagram")

        namespaces := $4
        classes := $5

        $$ = cld.CreateClassDiagram(wordsToString($2), namespaces, classes, $6)

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
        $$ = make(cld.Fields, 0)
    }
    | words EOM
    {
        fields := make(cld.Fields, 0)
        field := cld.CreateFieldFromString(wordsToString($1))
        fields = append(fields, field)
        $$ = fields
    }
    | fields words EOM
    {
        field := cld.CreateFieldFromString(wordsToString($2))
        $1 = append($1, field)
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
    : LABEL_CLASS words START_BLOCK classMetaDatas divisionMarker fields divisionMarker methods END_BLOCK
    {
        //pp.Println("class")
        stereotype := $4["stereotype"]

        $$ = cld.CreateClass(stereotype, wordsToString($2), $6, $8)

        id := $4["id"]
        $$.SetIdent(id)

        yylex.(*Lexer).Result = $$
    }

divisionMarker
    : EQUAL EQUAL

// メタデータを 「プロパティ : 値」 の形でパースするためのルール
// id         : 関連を記述するための識別子を別途指定したい場合に記述。
//              デフォルトは、ラベル名と同一。
// stereotype : ステレオタイプが必要な場合に記述。
classMetaDatas
    :
    {
        $$ = map[string]string{}
    }
    | words COLON words EOM
    {
        $$ = map[string]string{wordsToString($1):wordsToString($3)}
    }
    | classMetaDatas words COLON words EOM
    {
        $$[wordsToString($2)] = wordsToString($4)
    }

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
    : LABEL_NAMESPACE words START_BLOCK namespaces classes END_BLOCK
    {
        //pp.Println("namespace")

        $$ = cld.CreateNamespace(wordsToString($2), $5, $4)

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
    : relation_normal
    | relation_inherit
    | relation_implement

relation_normal
    : words HYPHEN GT words EOM
    {
        $$ = cld.CreateRelation("", cld.RELATION_NORMAL, wordsToString($1), wordsToString($4), "", "")
    }

relation_inherit
    : words HYPHEN PIPE GT words EOM
    {
        $$ = cld.CreateRelation("", cld.RELATION_INHERIT, wordsToString($1), wordsToString($5), "", "")
    }

relation_implement
    : words DOT PIPE GT words EOM
    {
        $$ = cld.CreateRelation("", cld.RELATION_IMPLEMENT, wordsToString($1), wordsToString($5), "", "")
    }

// 単語(WORD)の繰り返しルール
words
    : WORD
    {
        tokens := []*Token{$1}

        $$ = tokens

        yylex.(*Lexer).Result = $$
    }
    | validSigne /* word は、 '-' も許容したいので、こんな感じでルールを追加。 */
    {
        tokens := []*Token{$1}

        $$ = tokens

        yylex.(*Lexer).Result = $$
    }
    | words validSigne /* word は、 '-' も許容したいので、こんな感じでルールを追加。 */
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

validSigne
    : DOT
    | HYPHEN
    | COLON
    | LABEL_CLASS
    | LABEL_NAMESPACE
    | LABEL_CLD
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
