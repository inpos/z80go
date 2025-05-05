package lexer

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
	ASSIGN         // =
	ASSIGNNEW      // :=
	PLUS           // +
	MINUS          // -
	EQUAL          // ==
	NOTEQUAL       // !=
	LESS           // <
	GREATER        // >
	LESSOREQUAL    // <=
	GREATEROREQUAL // >=
	NOT            // !
	LOGAND         // &
	LOGOR          // |
	AND            // &&
	OR             // ||
	// symbols
	SEMICOLON // ;
	COLON     // :
	DOT       // .
	COMMA     // ,
	//
	EOF
)

var Symbols = map[string]int8{
	`;`: SEMICOLON,
	`:`: COLON,
	`.`: DOT,
	`,`: COMMA,
}

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

var Blocks = map[string]int8{
	`{`: LBRACE,
	`}`: RBRACE,
	`(`: LPAREN,
	`)`: RPAREN,
	`[`: LBRACKET,
	`]`: RBRACKET,
	`'`: QUOTE,
	`"`: DQUOTE,
}

var Operators = map[string]int8{
	`=`:  ASSIGN,
	`:=`: ASSIGNNEW,
	`+`:  PLUS,
	`-`:  MINUS,
	`==`: EQUAL,
	`!=`: NOTEQUAL,
	`<`:  LESS,
	`>`:  GREATER,
	`<=`: LESSOREQUAL,
	`>=`: GREATEROREQUAL,
	`!`:  NOT,
	`&`:  LOGAND,
	`|`:  LOGOR,
	`&&`: AND,
	`||`: OR,
}

type Lexer struct {
	char  string
	kind  int8
	value string
}
