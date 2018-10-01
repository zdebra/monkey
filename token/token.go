package token

type Type string

type Token struct {
	Type Type
	Literal string
}

const (
	Illegal = "ILLEGAL"
	EOF = "EOF"

	// Identifiers and literals
	Ident = "IDENT" // add, foobar, x, y, ...
	Int = "INT"

	// Operators
	Assign = "="
	Plus = "+"

	// Delimiters
	Comma = ","
	Semicolon = ";"
	LParen = "("
	RParen = ")"
	LBrace = "{"
	RBrace = "}"

	// Keywords
	Function = "FUNCTION"
	Let = "LET"
)
