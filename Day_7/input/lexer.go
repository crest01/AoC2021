// Code generated by re2c 2.2 on Tue Dec  7 20:12:10 2021, DO NOT EDIT.
//go:generate re2go $GOFILE.re -o $GOFILE --no-debug-info -W
package input


func lex(str *[]byte) []int { // Go code
	var YYMAXFILL int = 1

	str = pad(str, YYMAXFILL)
	limit := len(*str)
	var cursor, s int
	var yyt1 int

	crabs := make([]int, 0, limit)

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
	case ',':
		goto yy8
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
		yyt1 = cursor
		goto yy10
	default:
		goto yy4
	}
yy2:
	cursor += 1
	{ 
		if limit - cursor == YYMAXFILL - 1 {
			return crabs
		} else {
			panic("parse error!")
		}
	}
yy4:
	cursor += 1
	{ printError(str, cursor-1, cursor) }
yy6:
	cursor += 1
	{ return crabs }
yy8:
	cursor += 1
	{ goto loop }
yy10:
	cursor += 1
	if (limit - cursor < 1) {
		panic("yfill")
	}
	yych = (*str)[cursor]
	switch (yych) {
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
		goto yy10
	default:
		goto yy12
	}
yy12:
	s = yyt1
	{
		crabs = append(crabs, asInt(str, s, cursor))
		goto loop
	}
}

}
