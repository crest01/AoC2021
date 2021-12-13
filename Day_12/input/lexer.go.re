//go:generate re2go $GOFILE.re -o $GOFILE --no-debug-info -W
package input


func lex(str *[]byte) [][]string { // Go code
	/*!max:re2c*/
	str = pad(str, YYMAXFILL)
	limit := len(*str)
	var cursor, marker int
	var s, s1 int
	var /*!stags:re2c format = "@@{tag}"; separator=", ";*/ int
	edge := make([]string, 2)
	edges := make([][]string, 0)

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
		edges = append(edges, edge)
		edge = make([]string, 2)
		goto loop
	}
	edge = [a-zA-Z]+;
	sep = '-';
	node = @s edge sep @s1 edge;
	node {
		edge[0] = asString(str, s, s1-1)
		edge[1] = asString(str, s1, cursor)
		goto loop
	}
	[\x00] { 
		if limit - cursor == YYMAXFILL - 1 {
			return edges
		} else {
			panic("parse error!")
		}
	}
	*	{ printError(str, cursor-1, cursor) }
	*/
	return edges
}
