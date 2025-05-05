package lexer

import (
	"errors"
	"io"
	"strings"
)

const (
	ID int8 = iota + 1
	NUM
	ALPHA
	// keywords
	IMPORT
	PACKAGE
	TYPE
	VAR
	CONST
	FUNC
	RETURN
	FOR
	CONTINUE
	BREAK
	RANGE // NOTE: keyword?
	IF
	ELSE
	ELSEIF
	SWITCH
	CASE
	DEFAULT
	SELECT
	IOTA
	// blocks
	LBRACE   // {
	RBRACE   // }
	LPAREN   // (
	RPAREN   // )
	LBRACKET // [
	RBRACKET // ]
	QUOTE    // '
	DQUOTE   // "
	// operators
	EQUAL    // =
	NOTEQUAL // !
	PLUS     // +
	MINUS    // -
	LESS     // <
	GREATER  // >
	NOT      // !
	LOGAND   // &
	LOGOR    // |
	// symbols
	SEMICOLON // ;
	COLON     // :
	DOT       // .
	COMMA     // ,
	//
	NEWLINE // \n
	EOF
)

//var Symbols = map[string]int8{
//	`;`: SEMICOLON,
//	`:`: COLON,
//	`.`: DOT,
//	`,`: COMMA,
//}

var Keywords = map[string]int8{
	`package`:  PACKAGE,
	`import`:   IMPORT,
	`type`:     TYPE,
	`var`:      VAR,
	`const`:    CONST,
	`func`:     FUNC,
	`return`:   RETURN,
	`for`:      FOR,
	`continue`: CONTINUE,
	`break`:    BREAK,
	`range`:    RANGE,
	`if`:       IF,
	`else`:     ELSE,
	`else if`:  ELSEIF,
	`switch`:   SWITCH,
	`case`:     CASE,
	`default`:  DEFAULT,
	`select`:   SELECT,
	`iota`:     IOTA,
}

const (
	RuneLBRACE    = '{'
	RuneRBRACE    = '}'
	RuneLPAREN    = '('
	RuneRPAREN    = ')'
	RuneLBRACKET  = '['
	RuneRBRACKET  = ']'
	RuneQUOTE     = '\''
	RuneDQUOTE    = '"'
	RuneSEMICOLON = ';'
	RuneCOLON     = ':'
	RuneDOT       = '.'
	RuneCOMMA     = ','
	RuneEQUAL     = '='
	RuneNOTEQUAL  = '!'
	RuneLESS      = '<'
	RuneGREATER   = '>'
	RunePLUS      = '+'
	RuneMINUS     = '-'
)

var SymbolRunes = []rune{
	RuneLBRACE,
	RuneRBRACE,
	RuneLPAREN,
	RuneRPAREN,
	RuneLBRACKET,
	RuneRBRACKET,
	RuneQUOTE,
	RuneDQUOTE,
	RuneSEMICOLON,
	RuneCOLON,
	RuneDOT,
	RuneCOMMA,
	RuneEQUAL,
	RuneNOTEQUAL,
	RuneLESS,
	RuneGREATER,
	RunePLUS,
	RuneMINUS,
}

var Symbols = map[rune]int8{
	RuneLBRACE:    LBRACE,
	RuneRBRACE:    RBRACE,
	RuneLPAREN:    LPAREN,
	RuneRPAREN:    RPAREN,
	RuneLBRACKET:  LBRACKET,
	RuneRBRACKET:  RBRACKET,
	RuneQUOTE:     QUOTE,
	RuneDQUOTE:    DQUOTE,
	RuneSEMICOLON: SEMICOLON,
	RuneCOLON:     COLON,
	RuneDOT:       DOT,
	RuneCOMMA:     COMMA,
	RuneEQUAL:     EQUAL,
	RuneNOTEQUAL:  NOTEQUAL,
	RuneLESS:      LESS,
	RuneGREATER:   GREATER,
	RunePLUS:      PLUS,
	RuneMINUS:     MINUS,
}

type Token struct {
	Symbol int8
	Value  string
}

type Lexer struct {
	Symbol int8
	Value  string

	ch rune
}

func (l *Lexer) GetCh(r *strings.Reader) {
	var err error

	l.ch, _, err = r.ReadRune()
	if errors.Is(err, io.EOF) {
		l.ch = 0
	}
}

func (l *Lexer) NextToken(r *strings.Reader) *Token {
	l.Symbol = 0
	l.Value = ``

	for l.Symbol == 0 {
		switch {
		case l.ch == 0:
			l.Symbol = EOF
		case l.ch == '\n':
			l.Symbol = NEWLINE
		case runeIn(l.ch, ' ', '\t', '\r'):
			l.GetCh(r)
		case runeIn(l.ch, SymbolRunes...):
			l.Symbol = Symbols[l.ch]
		case l.ch >= '0' && l.ch <= '9':
			for l.ch >= '0' && l.ch <= '9' {

			}
		}
	}
}

func runeIn(ch rune, runes ...rune) bool {
	for _, r := range runes {
		if ch == r {
			return true
		}
	}

	return false
}
