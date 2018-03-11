//line ./parser.go.y:2
package ozcld

import __yyfmt__ "fmt"

//line ./parser.go.y:2
//line ./parser.go.y:5
type yySymType struct {
	yys            int
	classDiagram   *ClassDiagram
	class          *Class
	note           *Note
	classes        []*Class
	notes          []*Note
	classMetaDatas map[string]string
	fields         Fields
	methods        Methods
	namespace      *Namespace
	namespaces     []*Namespace
	relation       *Relation
	relations      []*Relation
	word           *Token
	words          []*Token
}

const WORD = 57346
const EOM = 57347
const EQUAL = 57348
const LABEL_CLD = 57349
const LABEL_CLASS = 57350
const LABEL_NOTE = 57351
const LABEL_NAMESPACE = 57352
const START_BLOCK = 57353
const END_BLOCK = 57354
const HYPHEN = 57355
const GT = 57356
const PIPE = 57357
const DOT = 57358
const COLON = 57359

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"WORD",
	"EOM",
	"EQUAL",
	"LABEL_CLD",
	"LABEL_CLASS",
	"LABEL_NOTE",
	"LABEL_NAMESPACE",
	"START_BLOCK",
	"END_BLOCK",
	"HYPHEN",
	"GT",
	"PIPE",
	"DOT",
	"COLON",
}
var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line ./parser.go.y:296

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyNprod = 44
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 360

var yyAct = [...]int{

	5, 21, 26, 53, 14, 24, 17, 16, 51, 32,
	19, 49, 50, 15, 85, 60, 12, 9, 10, 11,
	59, 25, 7, 20, 14, 6, 8, 31, 27, 14,
	22, 67, 18, 22, 27, 2, 18, 40, 14, 14,
	42, 65, 30, 36, 35, 34, 37, 20, 14, 3,
	25, 48, 79, 14, 57, 14, 62, 46, 1, 14,
	31, 0, 0, 0, 14, 23, 72, 14, 0, 28,
	14, 14, 14, 0, 38, 14, 0, 0, 0, 0,
	38, 14, 0, 14, 0, 4, 47, 55, 12, 9,
	10, 11, 52, 54, 7, 0, 58, 6, 8, 0,
	63, 0, 0, 66, 0, 0, 69, 70, 0, 71,
	0, 74, 15, 84, 0, 12, 9, 10, 11, 80,
	0, 7, 0, 0, 6, 8, 82, 4, 0, 0,
	12, 9, 10, 11, 0, 83, 7, 0, 0, 6,
	8, 15, 81, 0, 12, 9, 10, 11, 0, 0,
	7, 0, 0, 6, 8, 15, 78, 0, 12, 9,
	10, 11, 0, 0, 7, 0, 0, 6, 8, 15,
	77, 0, 12, 9, 10, 11, 0, 0, 7, 0,
	0, 6, 8, 15, 76, 0, 12, 9, 10, 11,
	0, 0, 7, 0, 0, 6, 8, 15, 75, 0,
	12, 9, 10, 11, 0, 0, 7, 0, 0, 6,
	8, 15, 73, 0, 12, 9, 10, 11, 0, 0,
	7, 0, 0, 6, 8, 15, 68, 0, 12, 9,
	10, 11, 0, 0, 7, 0, 15, 6, 8, 12,
	9, 10, 11, 0, 61, 7, 0, 15, 6, 8,
	12, 9, 10, 11, 45, 0, 7, 0, 4, 6,
	8, 12, 9, 10, 11, 0, 41, 7, 0, 15,
	6, 8, 12, 9, 10, 11, 39, 0, 7, 0,
	15, 6, 8, 12, 9, 10, 11, 29, 0, 7,
	0, 15, 6, 8, 12, 9, 10, 11, 13, 0,
	7, 0, 4, 6, 8, 12, 9, 10, 11, 0,
	0, 7, 0, 15, 6, 8, 12, 9, 10, 11,
	0, 0, 7, 0, 15, 6, 64, 12, 9, 10,
	11, 0, 0, 7, 0, 15, 6, 56, 12, 9,
	10, 11, 0, 0, 43, 0, 4, 44, 8, 12,
	9, 33, 11, 0, 0, 7, 0, 0, 6, 8,
}
var yyPact = [...]int{

	28, -1000, 298, 287, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, 26, -1000, -1000, 22, -1000, 298, 25,
	-1000, -1000, 298, 276, 342, -1000, -1000, 298, 265, 26,
	254, -1000, -1000, 298, -1000, -1000, -1000, 331, 243, 298,
	22, -1000, -1000, -3, -7, 298, 81, 320, 25, 298,
	6, 1, 232, 298, 309, 35, 298, 19, 221, 298,
	298, -1000, 81, 207, 298, -1000, 193, -1000, -1000, 179,
	165, 151, 298, -1000, 137, -1000, -1000, -1000, -1000, 123,
	108, -1000, 9, -1000, -1000, -1000,
}
var yyPgo = [...]int{

	0, 46, 0, 58, 1, 2, 10, 5, 57, 56,
	52, 3, 6, 7, 9, 45, 44, 43, 42,
}
var yyR1 = [...]int{

	0, 3, 6, 6, 6, 7, 7, 7, 9, 9,
	9, 10, 10, 10, 4, 5, 11, 8, 8, 8,
	13, 13, 13, 12, 18, 18, 18, 14, 14, 14,
	15, 16, 17, 1, 1, 1, 1, 2, 2, 2,
	2, 2, 2, 2,
}
var yyR2 = [...]int{

	0, 8, 0, 1, 2, 0, 1, 2, 0, 2,
	3, 0, 2, 3, 9, 5, 2, 0, 4, 5,
	0, 1, 2, 7, 0, 1, 2, 1, 1, 1,
	5, 6, 6, 1, 1, 2, 2, 1, 1, 1,
	1, 1, 1, 1,
}
var yyChk = [...]int{

	-1000, -3, 7, -1, 4, -2, 16, 13, 17, 8,
	9, 10, 7, 11, -2, 4, -13, -12, 10, -6,
	-12, -4, 8, -1, -7, -4, -5, 9, -1, 11,
	-18, -5, -14, 9, -15, -16, -17, -1, -1, 11,
	-13, 12, -14, 13, 16, 11, -8, -1, -6, 14,
	15, 15, -1, -11, -1, 6, 17, -7, -1, 14,
	14, 12, -9, -1, 17, 6, -1, 12, 5, -1,
	-1, -1, -11, 5, -1, 5, 5, 5, 5, -10,
	-1, 5, -1, 12, 5, 5,
}
var yyDef = [...]int{

	0, -2, 0, 0, 33, 34, 37, 38, 39, 40,
	41, 42, 43, 20, 35, 36, 2, 21, 0, 5,
	22, 3, 0, 0, 24, 4, 6, 0, 0, 20,
	0, 7, 25, 0, 27, 28, 29, 0, 0, 17,
	2, 1, 26, 38, 37, 0, 0, 0, 5, 0,
	0, 0, 0, 8, 0, 0, 0, 0, 0, 0,
	0, 15, 0, 0, 0, 16, 0, 23, 30, 0,
	0, 0, 11, 9, 0, 18, 31, 32, 10, 0,
	0, 19, 0, 14, 12, 13,
}
var yyTok1 = [...]int{

	1,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17,
}
var yyTok3 = [...]int{
	0,
}

var yyErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

//line yaccpar:1

/*	parser for yacc output	*/

var (
	yyDebug        = 0
	yyErrorVerbose = false
)

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

type yyParser interface {
	Parse(yyLexer) int
	Lookahead() int
}

type yyParserImpl struct {
	lval  yySymType
	stack [yyInitialStackSize]yySymType
	char  int
}

func (p *yyParserImpl) Lookahead() int {
	return p.char
}

func yyNewParser() yyParser {
	return &yyParserImpl{}
}

const yyFlag = -1000

func yyTokname(c int) string {
	if c >= 1 && c-1 < len(yyToknames) {
		if yyToknames[c-1] != "" {
			return yyToknames[c-1]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func yyStatname(s int) string {
	if s >= 0 && s < len(yyStatenames) {
		if yyStatenames[s] != "" {
			return yyStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func yyErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !yyErrorVerbose {
		return "syntax error"
	}

	for _, e := range yyErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + yyTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := yyPact[state]
	for tok := TOKSTART; tok-1 < len(yyToknames); tok++ {
		if n := base + tok; n >= 0 && n < yyLast && yyChk[yyAct[n]] == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if yyDef[state] == -2 {
		i := 0
		for yyExca[i] != -1 || yyExca[i+1] != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; yyExca[i] >= 0; i += 2 {
			tok := yyExca[i]
			if tok < TOKSTART || yyExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if yyExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += yyTokname(tok)
	}
	return res
}

func yylex1(lex yyLexer, lval *yySymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = yyTok1[0]
		goto out
	}
	if char < len(yyTok1) {
		token = yyTok1[char]
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			token = yyTok2[char-yyPrivate]
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		token = yyTok3[i+0]
		if token == char {
			token = yyTok3[i+1]
			goto out
		}
	}

out:
	if token == 0 {
		token = yyTok2[1] /* unknown char */
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", yyTokname(token), uint(char))
	}
	return char, token
}

func yyParse(yylex yyLexer) int {
	return yyNewParser().Parse(yylex)
}

func (yyrcvr *yyParserImpl) Parse(yylex yyLexer) int {
	var yyn int
	var yyVAL yySymType
	var yyDollar []yySymType
	_ = yyDollar // silence set and not used
	yyS := yyrcvr.stack[:]

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yyrcvr.char = -1
	yytoken := -1 // yyrcvr.char translated into internal numbering
	defer func() {
		// Make sure we report no lookahead when not parsing.
		yystate = -1
		yyrcvr.char = -1
		yytoken = -1
	}()
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	if yyDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", yyTokname(yytoken), yyStatname(yystate))
	}

	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	yyn = yyPact[yystate]
	if yyn <= yyFlag {
		goto yydefault /* simple state */
	}
	if yyrcvr.char < 0 {
		yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
	}
	yyn += yytoken
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = yyAct[yyn]
	if yyChk[yyn] == yytoken { /* valid shift */
		yyrcvr.char = -1
		yytoken = -1
		yyVAL = yyrcvr.lval
		yystate = yyn
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	}

yydefault:
	/* default state action */
	yyn = yyDef[yystate]
	if yyn == -2 {
		if yyrcvr.char < 0 {
			yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
		}

		/* look through exception table */
		xi := 0
		for {
			if yyExca[xi+0] == -1 && yyExca[xi+1] == yystate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			yyn = yyExca[xi+0]
			if yyn < 0 || yyn == yytoken {
				break
			}
		}
		yyn = yyExca[xi+1]
		if yyn < 0 {
			goto ret0
		}
	}
	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			yylex.Error(yyErrorMessage(yystate, yytoken))
			Nerrs++
			if yyDebug >= 1 {
				__yyfmt__.Printf("%s", yyStatname(yystate))
				__yyfmt__.Printf(" saw %s\n", yyTokname(yytoken))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				yyn = yyPact[yyS[yyp].yys] + yyErrCode
				if yyn >= 0 && yyn < yyLast {
					yystate = yyAct[yyn] /* simulate a shift of "error" */
					if yyChk[yystate] == yyErrCode {
						goto yystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if yyDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", yyS[yyp].yys)
				}
				yyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", yyTokname(yytoken))
			}
			if yytoken == yyEofCode {
				goto ret1
			}
			yyrcvr.char = -1
			yytoken = -1
			goto yynewstate /* try again in the same state */
		}
	}

	/* reduction by production yyn */
	if yyDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", yyn, yyStatname(yystate))
	}

	yynt := yyn
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= yyR2[yyn]
	// yyp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if yyp+1 >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	yyn = yyR1[yyn]
	yyg := yyPgo[yyn]
	yyj := yyg + yyS[yyp].yys + 1

	if yyj >= yyLast {
		yystate = yyAct[yyg]
	} else {
		yystate = yyAct[yyj]
		if yyChk[yystate] != -yyn {
			yystate = yyAct[yyg]
		}
	}
	// dummy call; replaced with literal code
	switch yynt {

	case 1:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line ./parser.go.y:71
		{
			//pp.Println("classDiagram")

			namespaces := yyDollar[4].namespaces
			classes := yyDollar[5].classes
			notes := yyDollar[6].notes

			yyVAL.classDiagram = CreateClassDiagram(wordsToString(yyDollar[2].words), namespaces, classes, notes, yyDollar[7].relations)

			yylex.(*Lexer).Result = yyVAL.classDiagram
		}
	case 2:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./parser.go.y:85
		{
			yyVAL.classes = []*Class{}
		}
	case 3:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser.go.y:89
		{
			yyVAL.classes = []*Class{yyDollar[1].class}
		}
	case 4:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./parser.go.y:93
		{
			yyVAL.classes = append(yyDollar[1].classes, yyDollar[2].class)
		}
	case 5:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./parser.go.y:99
		{
			yyVAL.notes = []*Note{}
		}
	case 6:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser.go.y:103
		{
			yyVAL.notes = []*Note{yyDollar[1].note}
		}
	case 7:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./parser.go.y:107
		{
			yyVAL.notes = append(yyDollar[1].notes, yyDollar[2].note)
		}
	case 8:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./parser.go.y:113
		{
			yyVAL.fields = Fields{}
		}
	case 9:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./parser.go.y:117
		{
			fields := Fields{CreateFieldFromString(wordsToString(yyDollar[1].words))}
			yyVAL.fields = fields
		}
	case 10:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser.go.y:122
		{
			field := CreateFieldFromString(wordsToString(yyDollar[2].words))
			yyDollar[1].fields = append(yyDollar[1].fields, field)
			yyVAL.fields = yyDollar[1].fields
		}
	case 11:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./parser.go.y:130
		{
			yyVAL.methods = Methods{}
		}
	case 12:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./parser.go.y:134
		{
			methods := Methods{CreateMethodFromString(wordsToString(yyDollar[1].words))}
			yyVAL.methods = methods
		}
	case 13:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line ./parser.go.y:139
		{
			method := CreateMethodFromString(wordsToString(yyDollar[2].words))
			yyDollar[1].methods = append(yyDollar[1].methods, method)
			yyVAL.methods = yyDollar[1].methods
		}
	case 14:
		yyDollar = yyS[yypt-9 : yypt+1]
		//line ./parser.go.y:147
		{
			//pp.Println("class")
			stereotype := yyDollar[4].classMetaDatas["stereotype"]

			yyVAL.class = CreateClass(stereotype, wordsToString(yyDollar[2].words), yyDollar[6].fields, yyDollar[8].methods)

			id := yyDollar[4].classMetaDatas["id"]
			yyVAL.class.SetIdent(id)

			yylex.(*Lexer).Result = yyVAL.class
		}
	case 15:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./parser.go.y:161
		{
			//pp.Println("note")
			yyVAL.note = CreateNote(wordsToString(yyDollar[2].words), wordsToString(yyDollar[4].words))

			yylex.(*Lexer).Result = yyVAL.note
		}
	case 17:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./parser.go.y:177
		{
			yyVAL.classMetaDatas = map[string]string{}
		}
	case 18:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line ./parser.go.y:181
		{
			yyVAL.classMetaDatas = map[string]string{wordsToString(yyDollar[1].words): wordsToString(yyDollar[3].words)}
		}
	case 19:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./parser.go.y:185
		{
			yyVAL.classMetaDatas[wordsToString(yyDollar[2].words)] = wordsToString(yyDollar[4].words)
		}
	case 20:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./parser.go.y:191
		{
			yyVAL.namespaces = []*Namespace{}
		}
	case 21:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser.go.y:195
		{
			yyVAL.namespaces = []*Namespace{yyDollar[1].namespace}
		}
	case 22:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./parser.go.y:199
		{
			yyVAL.namespaces = append(yyDollar[1].namespaces, yyDollar[2].namespace)
		}
	case 23:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line ./parser.go.y:205
		{
			//pp.Println("namespace")

			yyVAL.namespace = CreateNamespace(wordsToString(yyDollar[2].words), yyDollar[5].classes, yyDollar[6].notes, yyDollar[4].namespaces)

			yylex.(*Lexer).Result = yyVAL.namespace
		}
	case 24:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line ./parser.go.y:215
		{
			yyVAL.relations = []*Relation{}
		}
	case 25:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser.go.y:219
		{
			yyVAL.relations = []*Relation{yyDollar[1].relation}
		}
	case 26:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./parser.go.y:223
		{
			relations := append(yyDollar[1].relations, yyDollar[2].relation)

			yyVAL.relations = relations

			yylex.(*Lexer).Result = yyVAL.relations
		}
	case 30:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line ./parser.go.y:238
		{
			yyVAL.relation = CreateRelation("", RELATION_NORMAL, wordsToString(yyDollar[1].words), wordsToString(yyDollar[4].words), "", "")
		}
	case 31:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line ./parser.go.y:244
		{
			yyVAL.relation = CreateRelation("", RELATION_INHERIT, wordsToString(yyDollar[1].words), wordsToString(yyDollar[5].words), "", "")
		}
	case 32:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line ./parser.go.y:250
		{
			yyVAL.relation = CreateRelation("", RELATION_IMPLEMENT, wordsToString(yyDollar[1].words), wordsToString(yyDollar[5].words), "", "")
		}
	case 33:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser.go.y:257
		{
			tokens := []*Token{yyDollar[1].word}

			yyVAL.words = tokens

			yylex.(*Lexer).Result = yyVAL.words
		}
	case 34:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line ./parser.go.y:265
		{
			tokens := []*Token{yyDollar[1].word}

			yyVAL.words = tokens

			yylex.(*Lexer).Result = yyVAL.words
		}
	case 35:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./parser.go.y:273
		{
			tokens := append(yyDollar[1].words, yyDollar[2].word)

			yyVAL.words = tokens

			yylex.(*Lexer).Result = yyVAL.words
		}
	case 36:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line ./parser.go.y:280
		{
			tokens := append(yyDollar[1].words, yyDollar[2].word)

			yyVAL.words = tokens

			yylex.(*Lexer).Result = yyVAL.words
		}
	}
	goto yystack /* stack new state and value */
}
