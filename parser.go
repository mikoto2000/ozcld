
//line parser.go.y:2
package ozcld
import __yyfmt__ "fmt"
//line parser.go.y:2
		
//line parser.go.y:5
type yySymType struct{
	yys int
    classDiagram *ClassDiagram
    class *Class
    classes []*Class
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

const WORD = 57346
const EOM = 57347
const EQUAL = 57348
const LABEL_CLD = 57349
const LABEL_CLASS = 57350
const LABEL_NAMESPACE = 57351
const START_BLOCK = 57352
const END_BLOCK = 57353
const HYPHEN = 57354
const GT = 57355
const PIPE = 57356
const DOT = 57357
const COLON = 57358

var yyToknames = []string{
	"WORD",
	"EOM",
	"EQUAL",
	"LABEL_CLD",
	"LABEL_CLASS",
	"LABEL_NAMESPACE",
	"START_BLOCK",
	"END_BLOCK",
	"HYPHEN",
	"GT",
	"PIPE",
	"DOT",
	"COLON",
}
var yyStatenames = []string{}

const yyEofCode = 1
const yyErrCode = 2
const yyMaxDepth = 200

//line parser.go.y:266


//line yacctab:1
var yyExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyNprod = 39
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 319

var yyAct = []int{

	5, 16, 48, 18, 13, 15, 25, 20, 14, 75,
	41, 11, 9, 10, 39, 40, 7, 19, 47, 6,
	8, 46, 21, 13, 17, 52, 24, 21, 17, 2,
	34, 13, 13, 59, 23, 29, 28, 27, 38, 69,
	19, 56, 44, 42, 13, 1, 13, 0, 30, 0,
	13, 3, 24, 0, 0, 13, 13, 0, 13, 64,
	0, 13, 0, 0, 13, 0, 22, 13, 0, 0,
	31, 13, 0, 13, 0, 31, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 43, 0, 45, 0,
	0, 49, 0, 0, 0, 54, 55, 57, 14, 74,
	60, 11, 9, 10, 0, 63, 7, 66, 0, 6,
	8, 0, 4, 70, 0, 11, 9, 10, 72, 73,
	7, 0, 0, 6, 8, 14, 71, 0, 11, 9,
	10, 0, 0, 7, 0, 0, 6, 8, 14, 68,
	0, 11, 9, 10, 0, 0, 7, 0, 0, 6,
	8, 14, 67, 0, 11, 9, 10, 0, 0, 7,
	0, 0, 6, 8, 14, 65, 0, 11, 9, 10,
	0, 0, 7, 0, 0, 6, 8, 4, 0, 50,
	11, 9, 10, 0, 0, 7, 0, 0, 6, 8,
	14, 62, 0, 11, 9, 10, 0, 0, 7, 0,
	0, 6, 8, 14, 61, 0, 11, 9, 10, 0,
	0, 7, 0, 0, 6, 8, 14, 53, 0, 11,
	9, 10, 0, 0, 7, 0, 14, 6, 8, 11,
	9, 10, 37, 0, 7, 0, 4, 6, 8, 11,
	9, 10, 0, 33, 7, 0, 14, 6, 8, 11,
	9, 10, 32, 0, 7, 0, 14, 6, 8, 11,
	9, 10, 12, 0, 7, 0, 4, 6, 8, 11,
	9, 10, 0, 0, 7, 0, 14, 6, 8, 11,
	9, 10, 0, 0, 7, 0, 14, 6, 58, 11,
	9, 10, 0, 0, 7, 0, 14, 6, 51, 11,
	9, 10, 0, 0, 35, 0, 4, 36, 8, 11,
	26, 10, 0, 0, 7, 0, 0, 6, 8,
}
var yyPact = []int{

	22, -1000, 262, 252, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, 15, -1000, -1000, 19, -1000, 262, 302, -1000,
	-1000, 262, 242, 232, -1000, -1000, 262, -1000, -1000, -1000,
	292, 222, 15, -1000, -1000, 1, -4, 262, 19, 262,
	8, 5, 173, 282, 14, 212, 262, 262, 262, 272,
	27, 262, -1000, -1000, 199, 186, 173, 160, 262, -1000,
	147, -1000, -1000, 134, 262, -1000, 121, -1000, -1000, 108,
	94, -1000, 4, -1000, -1000, -1000,
}
var yyPgo = []int{

	0, 48, 0, 45, 7, 3, 43, 41, 39, 2,
	1, 5, 6, 37, 36, 35, 34,
}
var yyR1 = []int{

	0, 3, 5, 5, 5, 7, 7, 7, 8, 8,
	8, 4, 9, 6, 6, 6, 11, 11, 11, 10,
	16, 16, 16, 12, 12, 12, 13, 14, 15, 1,
	1, 1, 1, 2, 2, 2, 2, 2, 2,
}
var yyR2 = []int{

	0, 7, 0, 1, 2, 0, 2, 3, 0, 2,
	3, 9, 2, 0, 4, 5, 0, 1, 2, 6,
	0, 1, 2, 1, 1, 1, 5, 6, 6, 1,
	1, 2, 2, 1, 1, 1, 1, 1, 1,
}
var yyChk = []int{

	-1000, -3, 7, -1, 4, -2, 15, 12, 16, 8,
	9, 7, 10, -2, 4, -11, -10, 9, -5, -10,
	-4, 8, -1, -16, -4, -12, 8, -13, -14, -15,
	-1, -1, 10, 11, -12, 12, 15, 10, -11, 13,
	14, 14, -6, -1, -5, -1, 13, 13, -9, -1,
	6, 16, 11, 5, -1, -1, -7, -1, 16, 6,
	-1, 5, 5, -1, -9, 5, -1, 5, 5, -8,
	-1, 5, -1, 11, 5, 5,
}
var yyDef = []int{

	0, -2, 0, 0, 29, 30, 33, 34, 35, 36,
	37, 38, 16, 31, 32, 2, 17, 0, 20, 18,
	3, 0, 0, 0, 4, 21, 0, 23, 24, 25,
	0, 0, 16, 1, 22, 34, 33, 13, 2, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 5, 0,
	0, 0, 19, 26, 0, 0, 0, 0, 0, 12,
	0, 27, 28, 0, 8, 6, 0, 14, 7, 0,
	0, 15, 0, 11, 9, 10,
}
var yyTok1 = []int{

	1,
}
var yyTok2 = []int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16,
}
var yyTok3 = []int{
	0,
}

//line yaccpar:1

/*	parser for yacc output	*/

var yyDebug = 0

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

const yyFlag = -1000

func yyTokname(c int) string {
	// 4 is TOKSTART above
	if c >= 4 && c-4 < len(yyToknames) {
		if yyToknames[c-4] != "" {
			return yyToknames[c-4]
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

func yylex1(lex yyLexer, lval *yySymType) int {
	c := 0
	char := lex.Lex(lval)
	if char <= 0 {
		c = yyTok1[0]
		goto out
	}
	if char < len(yyTok1) {
		c = yyTok1[char]
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			c = yyTok2[char-yyPrivate]
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		c = yyTok3[i+0]
		if c == char {
			c = yyTok3[i+1]
			goto out
		}
	}

out:
	if c == 0 {
		c = yyTok2[1] /* unknown char */
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", yyTokname(c), uint(char))
	}
	return c
}

func yyParse(yylex yyLexer) int {
	var yyn int
	var yylval yySymType
	var yyVAL yySymType
	yyS := make([]yySymType, yyMaxDepth)

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yychar := -1
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	if yyDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", yyTokname(yychar), yyStatname(yystate))
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
	if yychar < 0 {
		yychar = yylex1(yylex, &yylval)
	}
	yyn += yychar
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = yyAct[yyn]
	if yyChk[yyn] == yychar { /* valid shift */
		yychar = -1
		yyVAL = yylval
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
		if yychar < 0 {
			yychar = yylex1(yylex, &yylval)
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
			if yyn < 0 || yyn == yychar {
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
			yylex.Error("syntax error")
			Nerrs++
			if yyDebug >= 1 {
				__yyfmt__.Printf("%s", yyStatname(yystate))
				__yyfmt__.Printf(" saw %s\n", yyTokname(yychar))
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
				__yyfmt__.Printf("error recovery discards %s\n", yyTokname(yychar))
			}
			if yychar == yyEofCode {
				goto ret1
			}
			yychar = -1
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
		//line parser.go.y:66
		{
	        //pp.Println("classDiagram")

	        namespaces := yyS[yypt-3].namespaces
	        classes := yyS[yypt-2].classes
	
	        yyVAL.classDiagram = CreateClassDiagram(wordsToString(yyS[yypt-5].words), namespaces, classes, yyS[yypt-1].relations)
	
	        yylex.(*Lexer).Result = yyVAL.classDiagram
	    }
	case 2:
		//line parser.go.y:79
		{
	        yyVAL.classes = []*Class{}
	    }
	case 3:
		//line parser.go.y:83
		{
	        yyVAL.classes = []*Class{yyS[yypt-0].class}
	    }
	case 4:
		//line parser.go.y:87
		{
	        yyVAL.classes = append(yyS[yypt-1].classes, yyS[yypt-0].class)
	    }
	case 5:
		//line parser.go.y:93
		{
	        yyVAL.fields = Fields{}
	    }
	case 6:
		//line parser.go.y:97
		{
	        fields := Fields{CreateFieldFromString(wordsToString(yyS[yypt-1].words))}
	        yyVAL.fields = fields
	    }
	case 7:
		//line parser.go.y:102
		{
	        field := CreateFieldFromString(wordsToString(yyS[yypt-1].words))
	        yyS[yypt-2].fields = append(yyS[yypt-2].fields, field)
	        yyVAL.fields = yyS[yypt-2].fields
	    }
	case 8:
		//line parser.go.y:110
		{
	        yyVAL.methods = Methods{}
	    }
	case 9:
		//line parser.go.y:114
		{
	        methods := Methods{CreateMethodFromString(wordsToString(yyS[yypt-1].words))}
	        yyVAL.methods = methods
	    }
	case 10:
		//line parser.go.y:119
		{
	        method := CreateMethodFromString(wordsToString(yyS[yypt-1].words))
	        yyS[yypt-2].methods = append(yyS[yypt-2].methods, method)
	        yyVAL.methods = yyS[yypt-2].methods
	    }
	case 11:
		//line parser.go.y:127
		{
	        //pp.Println("class")
        stereotype := yyS[yypt-5].classMetaDatas["stereotype"]
	
	        yyVAL.class = CreateClass(stereotype, wordsToString(yyS[yypt-7].words), yyS[yypt-3].fields, yyS[yypt-1].methods)
	
	        id := yyS[yypt-5].classMetaDatas["id"]
	        yyVAL.class.SetIdent(id)
	
	        yylex.(*Lexer).Result = yyVAL.class
	    }
	case 12:
		yyVAL.word = yyS[yypt-0].word
	case 13:
		//line parser.go.y:148
		{
	        yyVAL.classMetaDatas = map[string]string{}
	    }
	case 14:
		//line parser.go.y:152
		{
	        yyVAL.classMetaDatas = map[string]string{wordsToString(yyS[yypt-3].words):wordsToString(yyS[yypt-1].words)}
	    }
	case 15:
		//line parser.go.y:156
		{
	        yyVAL.classMetaDatas[wordsToString(yyS[yypt-3].words)] = wordsToString(yyS[yypt-1].words)
	    }
	case 16:
		//line parser.go.y:162
		{
	        yyVAL.namespaces = []*Namespace{}
	    }
	case 17:
		//line parser.go.y:166
		{
	        yyVAL.namespaces = []*Namespace{yyS[yypt-0].namespace}
	    }
	case 18:
		//line parser.go.y:170
		{
	        yyVAL.namespaces = append(yyS[yypt-1].namespaces, yyS[yypt-0].namespace)
	    }
	case 19:
		//line parser.go.y:176
		{
	        //pp.Println("namespace")

	        yyVAL.namespace = CreateNamespace(wordsToString(yyS[yypt-4].words), yyS[yypt-1].classes, yyS[yypt-2].namespaces)
	
	        yylex.(*Lexer).Result = yyVAL.namespace
	    }
	case 20:
		//line parser.go.y:186
		{
	        yyVAL.relations = []*Relation{}
	    }
	case 21:
		//line parser.go.y:190
		{
	        yyVAL.relations = []*Relation{yyS[yypt-0].relation}
	    }
	case 22:
		//line parser.go.y:194
		{
	        relations := append(yyS[yypt-1].relations, yyS[yypt-0].relation)
	
	        yyVAL.relations = relations
	
	        yylex.(*Lexer).Result = yyVAL.relations
	    }
	case 23:
		yyVAL.relation = yyS[yypt-0].relation
	case 24:
		yyVAL.relation = yyS[yypt-0].relation
	case 25:
		yyVAL.relation = yyS[yypt-0].relation
	case 26:
		//line parser.go.y:209
		{
	        yyVAL.relation = CreateRelation("", RELATION_NORMAL, wordsToString(yyS[yypt-4].words), wordsToString(yyS[yypt-1].words), "", "")
	    }
	case 27:
		//line parser.go.y:215
		{
	        yyVAL.relation = CreateRelation("", RELATION_INHERIT, wordsToString(yyS[yypt-5].words), wordsToString(yyS[yypt-1].words), "", "")
	    }
	case 28:
		//line parser.go.y:221
		{
	        yyVAL.relation = CreateRelation("", RELATION_IMPLEMENT, wordsToString(yyS[yypt-5].words), wordsToString(yyS[yypt-1].words), "", "")
	    }
	case 29:
		//line parser.go.y:228
		{
	        tokens := []*Token{yyS[yypt-0].word}
	
	        yyVAL.words = tokens
	
	        yylex.(*Lexer).Result = yyVAL.words
	    }
	case 30:
		//line parser.go.y:236
		{
	        tokens := []*Token{yyS[yypt-0].word}
	
	        yyVAL.words = tokens
	
	        yylex.(*Lexer).Result = yyVAL.words
	    }
	case 31:
		//line parser.go.y:244
		{
	        tokens := append(yyS[yypt-1].words, yyS[yypt-0].word)
	
	        yyVAL.words = tokens
	
	        yylex.(*Lexer).Result = yyVAL.words
	    }
	case 32:
		//line parser.go.y:251
		{
	        tokens := append(yyS[yypt-1].words, yyS[yypt-0].word)
	
	        yyVAL.words = tokens
	
	        yylex.(*Lexer).Result = yyVAL.words
	    }
	case 33:
		yyVAL.word = yyS[yypt-0].word
	case 34:
		yyVAL.word = yyS[yypt-0].word
	case 35:
		yyVAL.word = yyS[yypt-0].word
	case 36:
		yyVAL.word = yyS[yypt-0].word
	case 37:
		yyVAL.word = yyS[yypt-0].word
	case 38:
		yyVAL.word = yyS[yypt-0].word
	}
	goto yystack /* stack new state and value */
}
