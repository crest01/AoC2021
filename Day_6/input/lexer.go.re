//go:generate re2go $GOFILE.re -o $GOFILE --no-debug-info -W
package input

import (
	"strconv"
	"strings"
	"fmt"
)

func asUInt(str *string, start int, end int) uint64 {
	substr := (*str)[start : end]
	val, err := strconv.Atoi(substr)
	if err != nil {
		panic(err)
	}
	return uint64(val)
}

func printError(str *string, start int, end int) {
	substr := (*str)[start : end]
	fmt.Printf("Error: unexpected input '%s' at position %d", substr, start) 
	panic("Unexpected Input") 
}


// Pad string with YYMAXFILL zeroes at the end.
func pad(str string) string {
	/*!max:re2c*/
	return str + strings.Repeat("\000", YYMAXFILL)
}

func lex(str string) []uint64 { // Go code
	var cursor int
	limit := len(str)
	fishes := make([]uint64, 9)
	/*!max:re2c*/

loop:
	/*!re2c
	re2c:flags:tags = 0;
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
	eol { return fishes }
	comma = ",";
	integer = [0-8];
	integer {
		idx := asUInt(&str, cursor -1, cursor)
		fishes[idx]++
		goto loop
	}
	comma { goto loop }
	[\x00] { 
		if limit - cursor == YYMAXFILL - 1 {
			return fishes
		} else {
			panic("parse error!")
		}
	}
	*	{ printError(&str, cursor-1, cursor) }
	*/
}
