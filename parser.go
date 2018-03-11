//line .\parser.go.y:2
package ozcld

import __yyfmt__ "fmt"

//line .\parser.go.y:2
//line .\parser.go.y:5
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

//line .\parser.go.y:304

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyNprod = 46
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 387

var yyAct = [...]int{

	5, 21, 55, 26, 14, 17, 24, 32, 19, 52,
	53, 16, 50, 51, 15, 89, 62, 12, 9, 10,
	11, 25, 20, 7, 14, 61, 6, 8, 31, 14,
	27, 68, 22, 70, 18, 22, 27, 2, 43, 14,
	14, 41, 18, 38, 30, 37, 3, 20, 36, 14,
	49, 25, 35, 34, 83, 14, 59, 14, 65, 47,
	1, 14, 23, 31, 14, 0, 28, 14, 76, 0,
	14, 39, 0, 14, 14, 0, 14, 39, 0, 14,
	0, 0, 0, 0, 48, 14, 0, 14, 0, 0,
	54, 56, 0, 0, 60, 0, 0, 63, 0, 66,
	0, 0, 69, 0, 0, 72, 73, 0, 0, 75,
	0, 78, 15, 88, 0, 12, 9, 10, 11, 0,
	84, 7, 0, 0, 6, 8, 4, 86, 0, 12,
	9, 10, 11, 0, 87, 7, 0, 0, 6, 8,
	15, 85, 0, 12, 9, 10, 11, 0, 0, 7,
	0, 0, 6, 8, 15, 82, 0, 12, 9, 10,
	11, 0, 0, 7, 0, 0, 6, 8, 15, 81,
	0, 12, 9, 10, 11, 0, 0, 7, 0, 0,
	6, 8, 15, 80, 0, 12, 9, 10, 11, 0,
	0, 7, 0, 0, 6, 8, 15, 79, 0, 12,
	9, 10, 11, 0, 0, 7, 0, 0, 6, 8,
	15, 77, 0, 12, 9, 10, 11, 0, 0, 7,
	0, 0, 6, 8, 4, 0, 57, 12, 9, 10,
	11, 0, 0, 7, 0, 0, 6, 8, 15, 74,
	0, 12, 9, 10, 11, 0, 0, 7, 0, 0,
	6, 8, 15, 71, 0, 12, 9, 10, 11, 0,
	0, 7, 0, 15, 6, 8, 12, 9, 10, 11,
	0, 64, 7, 0, 15, 6, 8, 12, 9, 10,
	11, 46, 0, 7, 0, 4, 6, 8, 12, 9,
	10, 11, 0, 42, 7, 0, 15, 6, 8, 12,
	9, 10, 11, 40, 0, 7, 0, 15, 6, 8,
	12, 9, 10, 11, 29, 0, 7, 0, 15, 6,
	8, 12, 9, 10, 11, 13, 0, 7, 0, 4,
	6, 8, 12, 9, 10, 11, 0, 0, 7, 0,
	15, 6, 8, 12, 9, 10, 11, 0, 0, 7,
	0, 15, 6, 67, 12, 9, 10, 11, 0, 0,
	7, 0, 15, 6, 58, 12, 9, 10, 11, 0,
	0, 44, 0, 4, 45, 8, 12, 9, 33, 11,
	0, 0, 7, 0, 0, 6, 8,
}
var yyPact = [...]int{

	30, -1000, 325, 314, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, 32, -1000, -1000, 24, -1000, 325, 27,
	-1000, -1000, 325, 303, 369, -1000, -1000, 325, 292, 32,
	281, -1000, -1000, 325, -1000, -1000, -1000, -1000, 358, 270,
	325, 24, -1000, -1000, -2, -6, 325, 220, 347, 27,
	325, 11, 2, 325, 259, 325, 336, 25, 325, 21,
	248, 325, 325, 234, -1000, 220, 206, 325, -1000, 192,
	-1000, -1000, 178, 164, -1000, 150, 325, -1000, 136, -1000,
	-1000, -1000, -1000, 122, 108, -1000, 10, -1000, -1000, -1000,
}
var yyPgo = [...]int{

	0, 43, 0, 60, 1, 3, 8, 6, 59, 58,
	54, 2, 5, 11, 7, 53, 52, 48, 45, 44,
}
var yyR1 = [...]int{

	0, 3, 6, 6, 6, 7, 7, 7, 9, 9,
	9, 10, 10, 10, 4, 5, 11, 8, 8, 8,
	13, 13, 13, 12, 19, 19, 19, 14, 14, 14,
	14, 15, 16, 17, 18, 1, 1, 1, 1, 2,
	2, 2, 2, 2, 2, 2,
}
var yyR2 = [...]int{

	0, 8, 0, 1, 2, 0, 1, 2, 0, 2,
	3, 0, 2, 3, 9, 5, 2, 0, 4, 5,
	0, 1, 2, 7, 0, 1, 2, 1, 1, 1,
	1, 5, 6, 6, 5, 1, 1, 2, 2, 1,
	1, 1, 1, 1, 1, 1,
}
var yyChk = [...]int{

	-1000, -3, 7, -1, 4, -2, 16, 13, 17, 8,
	9, 10, 7, 11, -2, 4, -13, -12, 10, -6,
	-12, -4, 8, -1, -7, -4, -5, 9, -1, 11,
	-19, -5, -14, 9, -15, -16, -17, -18, -1, -1,
	11, -13, 12, -14, 13, 16, 11, -8, -1, -6,
	14, 15, 15, 16, -1, -11, -1, 6, 17, -7,
	-1, 14, 14, -1, 12, -9, -1, 17, 6, -1,
	12, 5, -1, -1, 5, -1, -11, 5, -1, 5,
	5, 5, 5, -10, -1, 5, -1, 12, 5, 5,
}
var yyDef = [...]int{

	0, -2, 0, 0, 35, 36, 39, 40, 41, 42,
	43, 44, 45, 20, 37, 38, 2, 21, 0, 5,
	22, 3, 0, 0, 24, 4, 6, 0, 0, 20,
	0, 7, 25, 0, 27, 28, 29, 30, 0, 0,
	17, 2, 1, 26, 40, 39, 0, 0, 0, 5,
	0, 0, 0, 0, 0, 8, 0, 0, 0, 0,
	0, 0, 0, 0, 15, 0, 0, 0, 16, 0,
	23, 31, 0, 0, 34, 0, 11, 9, 0, 18,
	32, 33, 10, 0, 0, 19, 0, 14, 12, 13,
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
		//line .\parser.go.y:72
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
		//line .\parser.go.y:86
		{
			yyVAL.classes = []*Class{}
		}
	case 3:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line .\parser.go.y:90
		{
			yyVAL.classes = []*Class{yyDollar[1].class}
		}
	case 4:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line .\parser.go.y:94
		{
			yyVAL.classes = append(yyDollar[1].classes, yyDollar[2].class)
		}
	case 5:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line .\parser.go.y:100
		{
			yyVAL.notes = []*Note{}
		}
	case 6:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line .\parser.go.y:104
		{
			yyVAL.notes = []*Note{yyDollar[1].note}
		}
	case 7:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line .\parser.go.y:108
		{
			yyVAL.notes = append(yyDollar[1].notes, yyDollar[2].note)
		}
	case 8:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line .\parser.go.y:114
		{
			yyVAL.fields = Fields{}
		}
	case 9:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line .\parser.go.y:118
		{
			fields := Fields{CreateFieldFromString(wordsToString(yyDollar[1].words))}
			yyVAL.fields = fields
		}
	case 10:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line .\parser.go.y:123
		{
			field := CreateFieldFromString(wordsToString(yyDollar[2].words))
			yyDollar[1].fields = append(yyDollar[1].fields, field)
			yyVAL.fields = yyDollar[1].fields
		}
	case 11:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line .\parser.go.y:131
		{
			yyVAL.methods = Methods{}
		}
	case 12:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line .\parser.go.y:135
		{
			methods := Methods{CreateMethodFromString(wordsToString(yyDollar[1].words))}
			yyVAL.methods = methods
		}
	case 13:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line .\parser.go.y:140
		{
			method := CreateMethodFromString(wordsToString(yyDollar[2].words))
			yyDollar[1].methods = append(yyDollar[1].methods, method)
			yyVAL.methods = yyDollar[1].methods
		}
	case 14:
		yyDollar = yyS[yypt-9 : yypt+1]
		//line .\parser.go.y:148
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
		//line .\parser.go.y:162
		{
			//pp.Println("note")
			yyVAL.note = CreateNote(wordsToString(yyDollar[2].words), wordsToString(yyDollar[4].words))

			yylex.(*Lexer).Result = yyVAL.note
		}
	case 17:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line .\parser.go.y:178
		{
			yyVAL.classMetaDatas = map[string]string{}
		}
	case 18:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line .\parser.go.y:182
		{
			yyVAL.classMetaDatas = map[string]string{wordsToString(yyDollar[1].words): wordsToString(yyDollar[3].words)}
		}
	case 19:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line .\parser.go.y:186
		{
			yyVAL.classMetaDatas[wordsToString(yyDollar[2].words)] = wordsToString(yyDollar[4].words)
		}
	case 20:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line .\parser.go.y:192
		{
			yyVAL.namespaces = []*Namespace{}
		}
	case 21:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line .\parser.go.y:196
		{
			yyVAL.namespaces = []*Namespace{yyDollar[1].namespace}
		}
	case 22:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line .\parser.go.y:200
		{
			yyVAL.namespaces = append(yyDollar[1].namespaces, yyDollar[2].namespace)
		}
	case 23:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line .\parser.go.y:206
		{
			//pp.Println("namespace")

			yyVAL.namespace = CreateNamespace(wordsToString(yyDollar[2].words), yyDollar[5].classes, yyDollar[6].notes, yyDollar[4].namespaces)

			yylex.(*Lexer).Result = yyVAL.namespace
		}
	case 24:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line .\parser.go.y:216
		{
			yyVAL.relations = []*Relation{}
		}
	case 25:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line .\parser.go.y:220
		{
			yyVAL.relations = []*Relation{yyDollar[1].relation}
		}
	case 26:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line .\parser.go.y:224
		{
			relations := append(yyDollar[1].relations, yyDollar[2].relation)

			yyVAL.relations = relations

			yylex.(*Lexer).Result = yyVAL.relations
		}
	case 31:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line .\parser.go.y:240
		{
			yyVAL.relation = CreateRelation("", RELATION_NORMAL, wordsToString(yyDollar[1].words), wordsToString(yyDollar[4].words), "", "")
		}
	case 32:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line .\parser.go.y:246
		{
			yyVAL.relation = CreateRelation("", RELATION_INHERIT, wordsToString(yyDollar[1].words), wordsToString(yyDollar[5].words), "", "")
		}
	case 33:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line .\parser.go.y:252
		{
			yyVAL.relation = CreateRelation("", RELATION_IMPLEMENT, wordsToString(yyDollar[1].words), wordsToString(yyDollar[5].words), "", "")
		}
	case 34:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line .\parser.go.y:258
		{
			yyVAL.relation = CreateRelation("", RELATION_NOTE, wordsToString(yyDollar[1].words), wordsToString(yyDollar[4].words), "", "")
		}
	case 35:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line .\parser.go.y:265
		{
			tokens := []*Token{yyDollar[1].word}

			yyVAL.words = tokens

			yylex.(*Lexer).Result = yyVAL.words
		}
	case 36:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line .\parser.go.y:273
		{
			tokens := []*Token{yyDollar[1].word}

			yyVAL.words = tokens

			yylex.(*Lexer).Result = yyVAL.words
		}
	case 37:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line .\parser.go.y:281
		{
			tokens := append(yyDollar[1].words, yyDollar[2].word)

			yyVAL.words = tokens

			yylex.(*Lexer).Result = yyVAL.words
		}
	case 38:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line .\parser.go.y:288
		{
			tokens := append(yyDollar[1].words, yyDollar[2].word)

			yyVAL.words = tokens

			yylex.(*Lexer).Result = yyVAL.words
		}
	}
	goto yystack /* stack new state and value */
}
