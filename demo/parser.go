//line demo.y:2

//TODO Put your favorite license here

// yacc source generated by ebnf2y[1]
// at 2013-07-17 12:48:48.235216367 +0200 CEST.
//
// CAUTION: If this file is a Go source file (*.go), it was generated
// automatically by '$ go tool yacc' from a *.y file - DO NOT EDIT in that case!
//
//   [1]: http://github.com/cznic/ebnf2y

package main //TODO real package name
import __yyfmt__ "fmt"

//line demo.y:13
//TODO required only be the demo _dump function
import (
	"bytes"
	"fmt"
	"strings"

	"github.com/cznic/strutil"
)

//line demo.y:26
type yySymType struct {
	yys  int
	item interface{} //TODO insert real field(s)
}

const ANDNOT = 57346
const BOOLEAN = 57347
const FLOAT = 57348
const IDENTIFIER = 57349
const IMAGINARY = 57350
const INTEGER = 57351
const LSH = 57352
const RSH = 57353
const STR = 57354

var yyToknames = []string{
	"ANDNOT",
	"BOOLEAN",
	"FLOAT",
	"IDENTIFIER",
	"IMAGINARY",
	"INTEGER",
	"LSH",
	"RSH",
	"STR",
}
var yyStatenames = []string{}

const yyEofCode = 1
const yyErrCode = 2
const yyMaxDepth = 200

//line demo.y:278

//TODO remove demo stuff below

var _parserResult interface{}

type (
	Expression      interface{}
	Expression1     interface{}
	Expression11    interface{}
	ExpressionList  interface{}
	ExpressionList1 interface{}
	Factor          interface{}
	Factor1         interface{}
	Factor11        interface{}
	Literal         interface{}
	Operand         interface{}
	Operand1        interface{}
	QualifiedIdent  interface{}
	QualifiedIdent1 interface{}
	Start           interface{}
	Term            interface{}
	Term1           interface{}
	Term11          interface{}
)

func _dump() {
	s := fmt.Sprintf("%#v", _parserResult)
	s = strings.Replace(s, "%", "%%", -1)
	s = strings.Replace(s, "{", "{%i\n", -1)
	s = strings.Replace(s, "}", "%u\n}", -1)
	s = strings.Replace(s, ", ", ",\n", -1)
	var buf bytes.Buffer
	strutil.IndentFormatter(&buf, ". ").Format(s)
	buf.WriteString("\n")
	a := strings.Split(buf.String(), "\n")
	for _, v := range a {
		if strings.HasSuffix(v, "(nil)") || strings.HasSuffix(v, "(nil),") {
			continue
		}

		fmt.Println(v)
	}
}

// End of demo stuff

//line yacctab:1
var yyExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 36,
	20, 27,
	-2, 12,
}

const yyNprod = 43
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 51

var yyAct = []int{

	2, 29, 4, 3, 39, 47, 45, 31, 32, 17,
	18, 22, 19, 20, 36, 49, 21, 37, 46, 30,
	33, 34, 35, 16, 28, 12, 1, 40, 38, 15,
	7, 41, 9, 10, 42, 8, 13, 44, 24, 25,
	26, 27, 14, 6, 5, 48, 43, 23, 11, 0,
	50,
}
var yyPact = []int{

	17, -1000, -1000, -1000, -1000, 4, -1000, -1000, -1000, -1000,
	-1000, 25, -3, -1000, -1000, -5, 17, -1000, -1000, -1000,
	-1000, -1000, -17, 17, -1000, -1000, -1000, -1000, 17, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, 17, -14, -1000, 11,
	-1000, -1000, -15, -1000, -1000, -1000, -1000, -1000, -2, 17,
	-1000,
}
var yyPgo = []int{

	0, 0, 48, 47, 46, 45, 2, 44, 43, 42,
	36, 34, 29, 28, 26, 3, 25, 24,
}
var yyR1 = []int{

	0, 1, 2, 2, 3, 3, 3, 3, 4, 5,
	5, 6, 7, 7, 8, 8, 8, 8, 9, 9,
	9, 9, 9, 9, 10, 10, 10, 11, 11, 12,
	13, 13, 14, 15, 16, 16, 17, 17, 17, 17,
	17, 17, 17,
}
var yyR2 = []int{

	0, 2, 0, 3, 1, 1, 1, 1, 2, 0,
	3, 2, 0, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 4, 3, 0, 1, 2,
	0, 2, 1, 2, 0, 3, 1, 1, 1, 1,
	1, 1, 1,
}
var yyChk = []int{

	-1000, -14, -1, -15, -6, -7, -8, 13, 18, 15,
	16, -2, -16, -10, -9, -12, 19, 5, 6, 8,
	9, 12, 7, -3, 13, 14, 15, 16, -17, 4,
	22, 10, 11, 23, 24, 25, 19, -1, -13, 21,
	-15, -6, -11, -4, -1, 20, 7, 20, -5, 17,
	-1,
}
var yyDef = []int{

	12, -2, 32, 2, 34, 0, 13, 14, 15, 16,
	17, 1, 33, 11, 24, 20, 12, 18, 19, 21,
	22, 23, 30, 12, 4, 5, 6, 7, 12, 36,
	37, 38, 39, 40, 41, 42, -2, 0, 29, 0,
	3, 35, 0, 28, 9, 26, 31, 25, 8, 12,
	10,
}
var yyTok1 = []int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 18, 3, 3, 3, 23, 22, 3,
	19, 20, 25, 16, 17, 15, 21, 24, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 13, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 14,
}
var yyTok2 = []int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12,
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
		__yyfmt__.Printf("lex %U %s\n", uint(char), yyTokname(c))
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
				__yyfmt__.Printf("saw %s\n", yyTokname(yychar))
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
		//line demo.y:78
		{
			yyVAL.item = []Expression{yyS[yypt-1].item, yyS[yypt-0].item} //TODO 1
		}
	case 2:
		//line demo.y:84
		{
			yyVAL.item = []Expression1(nil) //TODO 2
		}
	case 3:
		//line demo.y:88
		{
			yyVAL.item = append(yyS[yypt-2].item.([]Expression1), yyS[yypt-1].item, yyS[yypt-0].item) //TODO 3
		}
	case 4:
		//line demo.y:94
		{
			yyVAL.item = "^" //TODO 4
		}
	case 5:
		//line demo.y:98
		{
			yyVAL.item = "|" //TODO 5
		}
	case 6:
		//line demo.y:102
		{
			yyVAL.item = "-" //TODO 6
		}
	case 7:
		//line demo.y:106
		{
			yyVAL.item = "+" //TODO 7
		}
	case 8:
		//line demo.y:112
		{
			yyVAL.item = []ExpressionList{yyS[yypt-1].item, yyS[yypt-0].item} //TODO 8
		}
	case 9:
		//line demo.y:118
		{
			yyVAL.item = []ExpressionList1(nil) //TODO 9
		}
	case 10:
		//line demo.y:122
		{
			yyVAL.item = append(yyS[yypt-2].item.([]ExpressionList1), ",", yyS[yypt-0].item) //TODO 10
		}
	case 11:
		//line demo.y:128
		{
			yyVAL.item = []Factor{yyS[yypt-1].item, yyS[yypt-0].item} //TODO 11
		}
	case 12:
		//line demo.y:134
		{
			yyVAL.item = nil //TODO 12
		}
	case 13:
		//line demo.y:138
		{
			yyVAL.item = yyS[yypt-0].item //TODO 13
		}
	case 14:
		//line demo.y:144
		{
			yyVAL.item = "^" //TODO 14
		}
	case 15:
		//line demo.y:148
		{
			yyVAL.item = "!" //TODO 15
		}
	case 16:
		//line demo.y:152
		{
			yyVAL.item = "-" //TODO 16
		}
	case 17:
		//line demo.y:156
		{
			yyVAL.item = "+" //TODO 17
		}
	case 18:
		//line demo.y:162
		{
			yyVAL.item = yyS[yypt-0].item //TODO 18
		}
	case 19:
		//line demo.y:166
		{
			yyVAL.item = yyS[yypt-0].item //TODO 19
		}
	case 20:
		//line demo.y:170
		{
			yyVAL.item = yyS[yypt-0].item //TODO 20
		}
	case 21:
		//line demo.y:174
		{
			yyVAL.item = yyS[yypt-0].item //TODO 21
		}
	case 22:
		//line demo.y:178
		{
			yyVAL.item = yyS[yypt-0].item //TODO 22
		}
	case 23:
		//line demo.y:182
		{
			yyVAL.item = yyS[yypt-0].item //TODO 23
		}
	case 24:
		//line demo.y:188
		{
			yyVAL.item = yyS[yypt-0].item //TODO 24
		}
	case 25:
		//line demo.y:192
		{
			yyVAL.item = []Operand{yyS[yypt-3].item, "(", yyS[yypt-1].item, ")"} //TODO 25
		}
	case 26:
		//line demo.y:196
		{
			yyVAL.item = []Operand{"(", yyS[yypt-1].item, ")"} //TODO 26
		}
	case 27:
		//line demo.y:202
		{
			yyVAL.item = nil //TODO 27
		}
	case 28:
		//line demo.y:206
		{
			yyVAL.item = yyS[yypt-0].item //TODO 28
		}
	case 29:
		//line demo.y:212
		{
			yyVAL.item = []QualifiedIdent{yyS[yypt-1].item, yyS[yypt-0].item} //TODO 29
		}
	case 30:
		//line demo.y:218
		{
			yyVAL.item = nil //TODO 30
		}
	case 31:
		//line demo.y:222
		{
			yyVAL.item = []QualifiedIdent1{".", yyS[yypt-0].item} //TODO 31
		}
	case 32:
		//line demo.y:228
		{
			_parserResult = yyS[yypt-0].item //TODO 32
		}
	case 33:
		//line demo.y:234
		{
			yyVAL.item = []Term{yyS[yypt-1].item, yyS[yypt-0].item} //TODO 33
		}
	case 34:
		//line demo.y:240
		{
			yyVAL.item = []Term1(nil) //TODO 34
		}
	case 35:
		//line demo.y:244
		{
			yyVAL.item = append(yyS[yypt-2].item.([]Term1), yyS[yypt-1].item, yyS[yypt-0].item) //TODO 35
		}
	case 36:
		//line demo.y:250
		{
			yyVAL.item = yyS[yypt-0].item //TODO 36
		}
	case 37:
		//line demo.y:254
		{
			yyVAL.item = "&" //TODO 37
		}
	case 38:
		//line demo.y:258
		{
			yyVAL.item = yyS[yypt-0].item //TODO 38
		}
	case 39:
		//line demo.y:262
		{
			yyVAL.item = yyS[yypt-0].item //TODO 39
		}
	case 40:
		//line demo.y:266
		{
			yyVAL.item = "%" //TODO 40
		}
	case 41:
		//line demo.y:270
		{
			yyVAL.item = "/" //TODO 41
		}
	case 42:
		//line demo.y:274
		{
			yyVAL.item = "*" //TODO 42
		}
	}
	goto yystack /* stack new state and value */
}
