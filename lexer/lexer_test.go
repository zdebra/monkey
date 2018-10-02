package lexer

import (
	"testing"

	"github.com/zdebra/monkey/token"
)

func TestNextToken(t *testing.T) {
	t.Run("basic", func(t *testing.T) {
		input := `=+(){},;`

		tests := []struct {
			expectedType    token.Type
			expectedLiteral string
		}{
			{token.Assign, "="},
			{token.Plus, "+"},
			{token.LParen, "("},
			{token.RParen, ")"},
			{token.LBrace, "{"},
			{token.RBrace, "}"},
			{token.Comma, ","},
			{token.Semicolon, ";"},
			{token.EOF, ""},
		}

		l := New(input)
		for i, tt := range tests {
			tok := l.NextToken()
			if tok.Type != tt.expectedType {
				t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q", i, tt.expectedType, tok.Type)
			}
			if tok.Literal != tt.expectedLiteral {
				t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q", i, tt.expectedLiteral, tok.Literal)
			}
		}
	})

	t.Run("support for more tokens", func(t *testing.T) {
		input := `let five = 5; 
		let ten = 10; 
		let add = fn(x, y) { 
			x + y; 
		}; 
		
		let result = add(five, ten); 
		!-/*5;
		5 < 10 > 5; 
		if (5 < 10) {
			return true;
		} else {
			return false;
		}`

		tests := []struct {
			expectedType    token.Type
			expectedLiteral string
		}{
			{token.Let, "let"},
			{token.Ident, "five"},
			{token.Assign, "="},
			{token.Int, "5"},
			{token.Semicolon, ";"},
			{token.Let, "let"},
			{token.Ident, "ten"},
			{token.Assign, "="},
			{token.Int, "10"},
			{token.Semicolon, ";"},
			{token.Let, "let"},
			{token.Ident, "add"},
			{token.Assign, "="},
			{token.Function, "fn"},
			{token.LParen, "("},
			{token.Ident, "x"},
			{token.Comma, ","},
			{token.Ident, "y"},
			{token.RParen, ")"},
			{token.LBrace, "{"},
			{token.Ident, "x"},
			{token.Plus, "+"},
			{token.Ident, "y"},
			{token.Semicolon, ";"},
			{token.RBrace, "}"},
			{token.Semicolon, ";"},
			{token.Let, "let"},
			{token.Ident, "result"},
			{token.Assign, "="},
			{token.Ident, "add"},
			{token.LParen, "("},
			{token.Ident, "five"},
			{token.Comma, ","},
			{token.Ident, "ten"},
			{token.RParen, ")"},
			{token.Semicolon, ";"},
			{token.Bang, "!"},
			{token.Minus, "-"},
			{token.Slash, "/"},
			{token.Asterisk, "*"},
			{token.Int, "5"},
			{token.Semicolon, ";"},
			{token.Int, "5"},
			{token.LT, "<"},
			{token.Int, "10"},
			{token.GT, ">"},
			{token.Int, "5"},
			{token.Semicolon, ";"},
			{token.If, "if"},
			{token.LParen, "("},
			{token.Int, "5"},
			{token.LT, "<"},
			{token.Int, "10"},
			{token.RParen, ")"},
			{token.LBrace, "{"},
			{token.Return, "return"},
			{token.True, "true"},
			{token.Semicolon, ";"},
			{token.RBrace, "}"},
			{token.Else, "else"},
			{token.LBrace, "{"},
			{token.Return, "return"},
			{token.False, "false"},
			{token.Semicolon, ";"},
			{token.RBrace, "}"},
			{token.EOF, ""},
		}

		l := New(input)
		for i, tt := range tests {
			tok := l.NextToken()
			if tok.Type != tt.expectedType {
				t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q", i, tt.expectedType, tok.Type)
			}
			if tok.Literal != tt.expectedLiteral {
				t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q", i, tt.expectedLiteral, tok.Literal)
			}
		}
	})
}
