// Code generated by re2c 2.2 on Thu Dec  9 19:10:28 2021, DO NOT EDIT.
//go:generate re2go $GOFILE.re -o $GOFILE --no-debug-info -W
package input


func lex(str *[]byte) [][]int8 { // Go code
	var YYMAXFILL int = 1

	str = pad(str, YYMAXFILL)
	limit := len(*str)
	var cursor int
	canvas := make([][]int8, 0)
	line := make([]int8, 0)

loop:
	
{
	var yych byte
	if (limit - cursor < 1) {
		panic("yfill")
	}
	yych = (*str)[cursor]
	switch (yych) {
	case 0x00:
		goto yy2
	case '\n':
		goto yy6
	case '0':
		fallthrough
	case '1':
		fallthrough
	case '2':
		fallthrough
	case '3':
		fallthrough
	case '4':
		fallthrough
	case '5':
		fallthrough
	case '6':
		fallthrough
	case '7':
		fallthrough
	case '8':
		fallthrough
	case '9':
		goto yy8
	default:
		goto yy4
	}
yy2:
	cursor += 1
	{ 
		if limit - cursor == YYMAXFILL - 1 {
			return canvas
		} else {
			panic("parse error!")
		}
	}
yy4:
	cursor += 1
	{ printError(str, cursor-1, cursor) }
yy6:
	cursor += 1
	{ 
		canvas = append(canvas, line)
		line = make([]int8, 0)
		goto loop
	}
yy8:
	cursor += 1
	{ 
		line = append(line, asInt(str, cursor-1))
		goto loop
	}
}

	return canvas
}
