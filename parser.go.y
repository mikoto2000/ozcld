%{
package ozcld
%}

%union{
    classDiagram *ClassDiagram
    class *Class
    note *Note
    classes []*Class
    notes []*Note
    classMetaDatas map[string]string
    fields Fields
    methods Methods
    namespace *Namespace
    namespaces []*Namespace
    relation *Relation
    relations []*Relation
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
%type<note> note
%type<classes> classes
%type<notes> notes

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
%token<word> LABEL_NOTE
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
    : LABEL_CLD words START_BLOCK namespaces classes notes relations END_BLOCK
    {
        //pp.Println("classDiagram")

        namespaces := $4
        classes := $5
        notes := $6

        $$ = CreateClassDiagram(wordsToString($2), namespaces, classes, notes, $7)

        yylex.(*Lexer).Result = $$
    }

classes
    :
    {
        $$ = []*Class{}
    }
    | class
    {
        $$ = []*Class{$1}
    }
    | classes class
    {
        $$ = append($1, $2)
    }

notes
    :
    {
        $$ = []*Note{}
    }
    | note
    {
        $$ = []*Note{$1}
    }
    | notes note
    {
        $$ = append($1, $2)
    }

fields
    :
    {
        $$ = Fields{}
    }
    | words EOM
    {
        fields := Fields{CreateFieldFromString(wordsToString($1))}
        $$ = fields
    }
    | fields words EOM
    {
        field := CreateFieldFromString(wordsToString($2))
        $1 = append($1, field)
        $$ = $1
    }

methods
    :
    {
        $$ = Methods{}
    }
    | words EOM
    {
        methods := Methods{CreateMethodFromString(wordsToString($1))}
        $$ = methods
    }
    | methods words EOM
    {
        method := CreateMethodFromString(wordsToString($2))
        $1 = append($1, method)
        $$ = $1
    }

class
    : LABEL_CLASS words START_BLOCK classMetaDatas divisionMarker fields divisionMarker methods END_BLOCK
    {
        //pp.Println("class")
        stereotype := $4["stereotype"]

        $$ = CreateClass(stereotype, wordsToString($2), $6, $8)

        id := $4["id"]
        $$.SetIdent(id)

        yylex.(*Lexer).Result = $$
    }

note
    : LABEL_NOTE words START_BLOCK words END_BLOCK
    {
        //pp.Println("note")
        $$ = CreateNote(wordsToString($2), wordsToString($4))

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
        $$ = []*Namespace{}
    }
    | namespace
    {
        $$ = []*Namespace{$1}
    }
    | namespaces namespace
    {
        $$ = append($1, $2)
    }

namespace
    : LABEL_NAMESPACE words START_BLOCK namespaces classes notes END_BLOCK
    {
        //pp.Println("namespace")

        $$ = CreateNamespace(wordsToString($2), $5, $6, $4)

        yylex.(*Lexer).Result = $$
    }

relations
    :
    {
        $$ = []*Relation{}
    }
    | relation
    {
        $$ = []*Relation{$1}
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
        $$ = CreateRelation("", RELATION_NORMAL, wordsToString($1), wordsToString($4), "", "")
    }

relation_inherit
    : words HYPHEN PIPE GT words EOM
    {
        $$ = CreateRelation("", RELATION_INHERIT, wordsToString($1), wordsToString($5), "", "")
    }

relation_implement
    : words DOT PIPE GT words EOM
    {
        $$ = CreateRelation("", RELATION_IMPLEMENT, wordsToString($1), wordsToString($5), "", "")
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
    | LABEL_NOTE
    | LABEL_NAMESPACE
    | LABEL_CLD
%%
