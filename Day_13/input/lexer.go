// Code generated by re2c 2.2 on Mon Dec 13 20:20:47 2021, DO NOT EDIT.
//go:generate re2go $GOFILE.re -o $GOFILE --no-debug-info -W
package input

import ("advent_of_code/day13/sparse_field")

func lex(str *[]byte) ([]sparse_field.Int2, []sparse_field.Int2) { // Go code
	var YYMAXFILL int = 14

	str = pad(str, YYMAXFILL)
	limit := len(*str)
	var cursor, marker int
	var s1, s2 int
	var yyt1, yyt2 int
	dots := make([]sparse_field.Int2, 0)
	folds := make([]sparse_field.Int2, 0)

loop:
	
{
	var yych byte
	if (limit - cursor < 14) {
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
		yyt1 = cursor
		goto yy8
	case 'f':
		goto yy9
	default:
		goto yy4
	}
yy2:
	cursor += 1
	{ 
		if limit - cursor == YYMAXFILL - 1 {
			return dots, folds
		} else {
			panic("parse error!")
		}
	}
yy4:
	cursor += 1
yy5:
	{ printError(str, cursor-1, cursor) }
yy6:
	cursor += 1
	{ 
		goto loop
	}
yy8:
	cursor += 1
	marker = cursor
	yych = (*str)[cursor]
	switch (yych) {
	case ',':
		goto yy10
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
		goto yy12
	default:
		goto yy5
	}
yy9:
	cursor += 1
	marker = cursor
	yych = (*str)[cursor]
	switch (yych) {
	case 'o':
		goto yy14
	default:
		goto yy5
	}
yy10:
	cursor += 1
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
		yyt2 = cursor
		goto yy15
	default:
		goto yy11
	}
yy11:
	cursor = marker
	goto yy5
yy12:
	cursor += 1
	if (limit - cursor < 2) {
		panic("yfill")
	}
	yych = (*str)[cursor]
	switch (yych) {
	case ',':
		goto yy10
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
		goto yy12
	default:
		goto yy11
	}
yy14:
	cursor += 1
	yych = (*str)[cursor]
	switch (yych) {
	case 'l':
		goto yy18
	default:
		goto yy11
	}
yy15:
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
		goto yy15
	default:
		goto yy17
	}
yy17:
	s1 = yyt1
	s2 = yyt2
	{
		dots = append(dots, sparse_field.MakeInt2(asInt(str, s1, s2-1), asInt(str, s2, cursor)))
		goto loop
	}
yy18:
	cursor += 1
	yych = (*str)[cursor]
	switch (yych) {
	case 'd':
		goto yy19
	default:
		goto yy11
	}
yy19:
	cursor += 1
	yych = (*str)[cursor]
	switch (yych) {
	case ' ':
		goto yy20
	default:
		goto yy11
	}
yy20:
	cursor += 1
	yych = (*str)[cursor]
	switch (yych) {
	case 'a':
		goto yy21
	default:
		goto yy11
	}
yy21:
	cursor += 1
	yych = (*str)[cursor]
	switch (yych) {
	case 'l':
		goto yy22
	default:
		goto yy11
	}
yy22:
	cursor += 1
	yych = (*str)[cursor]
	switch (yych) {
	case 'o':
		goto yy23
	default:
		goto yy11
	}
yy23:
	cursor += 1
	yych = (*str)[cursor]
	switch (yych) {
	case 'n':
		goto yy24
	default:
		goto yy11
	}
yy24:
	cursor += 1
	yych = (*str)[cursor]
	switch (yych) {
	case 'g':
		goto yy25
	default:
		goto yy11
	}
yy25:
	cursor += 1
	yych = (*str)[cursor]
	switch (yych) {
	case ' ':
		goto yy26
	default:
		goto yy11
	}
yy26:
	cursor += 1
	yych = (*str)[cursor]
	switch (yych) {
	case 'x':
		fallthrough
	case 'y':
		goto yy27
	default:
		goto yy11
	}
yy27:
	cursor += 1
	yych = (*str)[cursor]
	switch (yych) {
	case '=':
		goto yy28
	default:
		goto yy11
	}
yy28:
	cursor += 1
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
		yyt1 = cursor
		goto yy29
	default:
		goto yy11
	}
yy29:
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
		goto yy29
	default:
		goto yy31
	}
yy31:
	s2 = yyt1
	s1 = yyt1
	s1 += -2
	{
		if (*str)[s1] == 'x' {
			folds = append(folds, sparse_field.MakeInt2(asInt(str, s2, cursor), 0))
		} else if (*str)[s1] == 'y' {
			folds = append(folds, sparse_field.MakeInt2(0, asInt(str, s2, cursor)))
		}
		goto loop
	}
}

	return dots, folds
}