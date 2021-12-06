//go:generate re2go ./$GOFILE.re -o ./$GOFILE --no-debug-info
package input

import (
	"strconv"
	"strings"
)

func asInt(str *string, start int, end int) int {
	substr := (*str)[start : end]
	val, err := strconv.Atoi(substr)
	if err != nil {
		panic(err)
	}
	return val
}



// Pad string with YYMAXFILL zeroes at the end.
func pad(str string) string {
	/*!max:re2c*/
	return str + strings.Repeat("\000", YYMAXFILL)
}

func lex(str string) []int { // Go code
	var cursor int
	limit := len(str)
	var values []int
	var s, e int
	var /*!stags:re2c format = "@@{tag}"; separator=", ";*/ int
	/*!max:re2c*/

loop:
	/*!re2c
	re2c:flags:tags = 1;
	re2c:flags:type-header = 1;
	re2c:define:YYCTYPE    = byte;
	re2c:define:YYPEEK     = "str[cursor]";
	re2c:define:YYSKIP     = "cursor += 1";
	re2c:define:YYLESSTHAN = "limit - cursor < @@{len}";
	re2c:define:YYFILL     = "panic(\"yfill\")";
	re2c:define:YYBACKUP    = "marker = cursor";
	re2c:define:YYRESTORE   = "cursor = marker";
	re2c:define:YYSTAGP     = "@@{tag} = cursor";
	re2c:define:YYSTAGN     = "@@{tag} = -1";
	re2c:define:YYSHIFTSTAG = "@@{tag} += @@{shift}";

	eol  = "\n";
	eol { goto loop }
	space = " ";
	comma = ",";
	arrow = "->";
	ws = (space*|comma)+;
	integer = @s[1-9][0-9]*@e;
	integer {
		values = append(values, asInt(&str, s, e))
		goto loop
	}
	ws { goto loop }
	arrow {goto loop }
	[\x00] { 
		if limit - cursor == YYMAXFILL - 1 {
			return values
		} else {
			panic("parse error!")
		}
	}
	*	{ panic("default rule!") }       // no other rule matches
	*/
}
