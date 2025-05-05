package parser

import (
	"strings"
	"z80go/lexer"
)

func Parse(src string) (*lexer.Node, error) {
	r := strings.NewReader(src)

}
