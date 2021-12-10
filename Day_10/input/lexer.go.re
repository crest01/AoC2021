//go:generate re2go $GOFILE.re -o $GOFILE --no-debug-info -W
package input


func lex(str *[]byte) ([][]byte, []byte, []int) { // Go code
	/*!max:re2c*/
	str = pad(str, YYMAXFILL)
	limit := len(*str)
	var cursor int
	stack := make([]byte, 0)
	line := make([]byte, 0)
	lines := make([][]byte, 0)
	corrupt_chars := make([]byte, 0)
	corrupt_lines := make([]int, 0)
	line_idx := 0

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
		lines = append(lines, line)
		stack = make([]byte, 0)
		line = make([]byte, 0)
		line_idx ++
		goto loop
	}
	open = [(<\{\[];
	open 
	{
		stack = append(stack, (*str)[cursor-1])
		line = append(line, (*str)[cursor-1])
		goto loop
	}
	close = [\])>\}];
	close 
	{
		top_of_stack := stack[len(stack)-1]
		closing := getClosingBracket(top_of_stack)
		current_bracket := (*str)[cursor-1]

		line = append(line, current_bracket)

		if current_bracket != closing {
			corrupt_chars = append(corrupt_chars, current_bracket)
			corrupt_lines = append(corrupt_lines, line_idx)
		}
	}
	[\x00] { 
		if limit - cursor == YYMAXFILL - 1 {
			return lines, corrupt_chars, corrupt_lines
		} else {
			panic("parse error!")
		}
	}
	*	{ printError(str, cursor-1, cursor) }
	*/
	return canvas
}
