//go:generate re2go $GOFILE.re -o $GOFILE --no-debug-info -W
package input


func lex(str *[]byte) [][]int8 { // Go code
	/*!max:re2c*/
	str = pad(str, YYMAXFILL)
	limit := len(*str)
	var cursor int
	canvas := make([][]int8, 0)
	line := make([]int8, 0)

loop:
	/*!re2c
	re2c:define:YYCTYPE    = byte;
	re2c:define:YYPEEK     = "(*str)[cursor]";
	re2c:define:YYSKIP     = "cursor += 1";
	re2c:define:YYLESSTHAN = "limit - cursor < @@{len}";
	re2c:define:YYFILL     = "panic(\"yfill\")";
	re2c:define:YYBACKUP    = "marker = cursor";
	re2c:define:YYRESTORE   = "cursor = marker";
	re2c:define:YYSTAGP     = "@@{tag} = cursor";
	re2c:define:YYSTAGN     = "@@{tag} = -1";
	re2c:define:YYSHIFTSTAG = "@@{tag} += @@{shift}";

	eol  = "\n";
	eol 
	{ 
		canvas = append(canvas, line)
		line = make([]int8, 0)
		goto loop
	}
	number = [0-9];
	number 
	{ 
		line = append(line, asInt(str, cursor-1))
		goto loop
	}
	[\x00] { 
		if limit - cursor == YYMAXFILL - 1 {
			return canvas
		} else {
			panic("parse error!")
		}
	}
	*	{ printError(str, cursor-1, cursor) }
	*/
	return canvas
}