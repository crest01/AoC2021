//go:generate re2go $GOFILE.re -o $GOFILE --no-debug-info -W
package input
import "advent_of_code/day8/signal"


func lex(str *[]byte) []signal.Pattern { // Go code
	/*!max:re2c*/
	str = pad(str, YYMAXFILL)
	limit := len(*str)
	var cursor, s int
	flag := false
	var /*!stags:re2c format = "@@{tag}"; separator=", ";*/ int
	
	var pattern signal.Pattern
	result := make([]signal.Pattern, 0)

loop:
	/*!re2c
	re2c:flags:tags = 1;
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
		result = append(result, pattern)
		pattern = signal.Pattern{}
		flag = false
		goto loop
	}

	ws = " "|"\r";
	ws { goto loop }

	segment = @s [a-g]+;
	segment 
	{
		if !flag {
			pattern.S = append(pattern.S, asSignal(str, s, cursor))
		} else {
			pattern.O = append(pattern.O, asSignal(str, s, cursor))
		}
		goto loop
	}

	separator = "|";
	separator 
	{
		flag = true
		goto loop
	}
	[\x00] { 
		if limit - cursor == YYMAXFILL - 1 {
			return result
		} else {
			panic("parse error!")
		}
	}
	*	{ printError(str, cursor-1, cursor) }
	*/
}
