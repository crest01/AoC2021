//go:generate re2go $GOFILE.re -o $GOFILE --no-debug-info -W
package input

import ("advent_of_code/day13/sparse_field")

func lex(str *[]byte) ([]sparse_field.Int2, []sparse_field.Int2) { // Go code
	/*!max:re2c*/
	str = pad(str, YYMAXFILL)
	limit := len(*str)
	var cursor, marker int
	var s1, s2 int
	var /*!stags:re2c format = "@@{tag}"; separator=", ";*/ int
	dots := make([]sparse_field.Int2, 0)
	folds := make([]sparse_field.Int2, 0)

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
	re2c:define:YYSTAGN     = "@@{tag} = -1go";
	re2c:define:YYSHIFTSTAG = "@@{tag} += @@{shift}";

	eol  = "\n";
	eol 
	{ 
		goto loop
	}
	number = [0-9]+;
	coord =  @s1 number ',' @s2 number;
	coord
	{
		dots = append(dots, sparse_field.MakeInt2(asInt(str, s1, s2-1), asInt(str, s2, cursor)))
		goto loop
	}
	fold = "fold along " @s1 [xy] "=" @s2 number;
	fold 
	{
		if (*str)[s1] == 'x' {
			folds = append(folds, sparse_field.MakeInt2(asInt(str, s2, cursor), 0))
		} else if (*str)[s1] == 'y' {
			folds = append(folds, sparse_field.MakeInt2(0, asInt(str, s2, cursor)))
		}
		goto loop
	}
	[\x00] { 
		if limit - cursor == YYMAXFILL - 1 {
			return dots, folds
		} else {
			panic("parse error!")
		}
	}
	*	{ printError(str, cursor-1, cursor) }
	*/
	return dots, folds
}
