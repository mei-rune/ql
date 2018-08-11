// CAUTION: Generated file - DO NOT EDIT.

// Copyright (c) 2015 The golex Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This is an example program using golex run time library. It is generated by
//
//	$ golex -o example_test.go example.l
//
// The complete input file, example.l, is at [3], the scan function excerpt is:
//
//	func (l *lexer) scan() lex.Char {
//		c := l.Enter()
//	%}
//
//	digit		[0-9]|{unicodeDigit}
//	identifier	{letter}({letter}|{digit})*
//	int		[0-9]+
//	letter		[_a-zA-Z]|{unicodeLetter}
//	unicodeDigit	\x81
//	unicodeLetter	\x80
//
//	%%
//
//		c = l.Rule0()
//
//	[ \t\r\n]+
//
//	func		return l.char(FUNC)
//	{identifier}	return l.char(IDENT)
//	{int}		return l.char(INT)
//
//
//	%%
//		if c, ok := l.Abort(); ok {
//			return l.char(c)
//		}
//
//		goto yyAction
//	}
package lex_test // import "github.com/cznic/ql/golex/lex"

import (
	"bytes"
	"fmt"
	"go/token"
	"unicode"

	"github.com/cznic/ql/golex/lex"
)

// Allocate Character classes anywhere in [0x80, 0xFF].
const (
	classUnicodeLeter = iota + 0x80
	classUnicodeDigit
	classOther
)

// Parser token values.
const (
	FUNC = iota + 0xE002
	INT
	IDENT
)

// For pretty printing.
func str(r rune) string {
	switch r {
	case FUNC:
		return "FUNC"
	case INT:
		return "INT"
	case IDENT:
		return "IDENT"
	case lex.RuneEOF:
		return "EOF"
	}

	return fmt.Sprintf("%q", r)
}

type lexer struct {
	*lex.Lexer
}

func (l *lexer) char(r int) lex.Char {
	return lex.NewChar(l.First.Pos(), rune(r))
}

func rune2Class(r rune) int {
	if r >= 0 && r < 0x80 { // Keep ASCII as it is.
		return int(r)
	}

	if unicode.IsLetter(r) {
		return classUnicodeLeter
	}

	if unicode.IsDigit(r) {
		return classUnicodeDigit
	}

	return classOther
}

const src = `

func Xφ42() int { return 314 }

`

func Example_completeGeneratedProgram() { // main
	fset := token.NewFileSet()
	file := fset.AddFile("example.go", -1, len(src))
	src := bytes.NewBufferString(src)
	lx, err := lex.New(file, src, lex.RuneClass(rune2Class))
	if err != nil {
		panic(err)
	}

	l := &lexer{lx}
	for {
		c := l.scan()
		fmt.Printf("%v: %v %q\n", file.Position(c.Pos()), str(c.Rune), l.TokenBytes(nil))
		if c.Rune == lex.RuneEOF {
			return
		}
	}
	// Output:
	// example.go:3:1: FUNC "func"
	// example.go:3:6: IDENT "Xφ42"
	// example.go:3:11: '(' "("
	// example.go:3:12: ')' ")"
	// example.go:3:14: IDENT "int"
	// example.go:3:18: '{' "{"
	// example.go:3:20: IDENT "return"
	// example.go:3:27: INT "314"
	// example.go:3:31: '}' "}"
	// example.go:4:2: EOF "\xff"
}

func (l *lexer) scan() lex.Char {
	c := l.Enter()

yystate0:
	yyrule := -1
	_ = yyrule
	c = l.Rule0()

	goto yystart1

	goto yystate0 // silence unused label error
	goto yyAction // silence unused label error
yyAction:
	switch yyrule {
	case 1:
		goto yyrule1
	case 2:
		goto yyrule2
	case 3:
		goto yyrule3
	case 4:
		goto yyrule4
	}
	goto yystate1 // silence unused label error
yystate1:
	c = l.Next()
yystart1:
	switch {
	default:
		goto yyabort
	case c == '\t' || c == '\n' || c == '\r' || c == ' ':
		goto yystate2
	case c == 'f':
		goto yystate5
	case c >= '0' && c <= '9':
		goto yystate3
	case c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'e' || c >= 'g' && c <= 'z' || c == '\u0080':
		goto yystate4
	}

yystate2:
	c = l.Next()
	yyrule = 1
	l.Mark()
	switch {
	default:
		goto yyrule1
	case c == '\t' || c == '\n' || c == '\r' || c == ' ':
		goto yystate2
	}

yystate3:
	c = l.Next()
	yyrule = 4
	l.Mark()
	switch {
	default:
		goto yyrule4
	case c >= '0' && c <= '9':
		goto yystate3
	}

yystate4:
	c = l.Next()
	yyrule = 3
	l.Mark()
	switch {
	default:
		goto yyrule3
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z' || c == '\u0080' || c == '\u0081':
		goto yystate4
	}

yystate5:
	c = l.Next()
	yyrule = 3
	l.Mark()
	switch {
	default:
		goto yyrule3
	case c == 'u':
		goto yystate6
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 't' || c >= 'v' && c <= 'z' || c == '\u0080' || c == '\u0081':
		goto yystate4
	}

yystate6:
	c = l.Next()
	yyrule = 3
	l.Mark()
	switch {
	default:
		goto yyrule3
	case c == 'n':
		goto yystate7
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'm' || c >= 'o' && c <= 'z' || c == '\u0080' || c == '\u0081':
		goto yystate4
	}

yystate7:
	c = l.Next()
	yyrule = 3
	l.Mark()
	switch {
	default:
		goto yyrule3
	case c == 'c':
		goto yystate8
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c == 'a' || c == 'b' || c >= 'd' && c <= 'z' || c == '\u0080' || c == '\u0081':
		goto yystate4
	}

yystate8:
	c = l.Next()
	yyrule = 2
	l.Mark()
	switch {
	default:
		goto yyrule2
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z' || c == '\u0080' || c == '\u0081':
		goto yystate4
	}

yyrule1: // [ \t\r\n]+

	goto yystate0
yyrule2: // func
	{
		return l.char(FUNC)
	}
yyrule3: // {identifier}
	{
		return l.char(IDENT)
	}
yyrule4: // {int}
	{
		return l.char(INT)
	}
	panic("unreachable")

	goto yyabort // silence unused label error

yyabort: // no lexem recognized
	if c, ok := l.Abort(); ok {
		return l.char(c)
	}

	goto yyAction
}
