package parser

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/zdebra/monkey/ast"

	"github.com/stretchr/testify/assert"

	"github.com/zdebra/monkey/lexer"
)

func TestLetStatements(t *testing.T) {
	input := `
	let x = 5;
	let y = 10;
	let foobar = 838383;
	`

	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	assert.NotNil(t, program)
	assert.Len(t, program.Statements, 3)

	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}

	for i, tt := range tests {
		stmt := program.Statements[i]
		testLetStatement(t, stmt, tt.expectedIdentifier)
	}
}

func testLetStatement(t *testing.T, s ast.Statement, name string) {
	require.Equal(t, "let", s.TokenLiteral())
	letStmt, ok := s.(*ast.LetStatement)
	require.True(t, ok)
	require.Equal(t, name, letStmt.Name.Value)
	require.Equal(t, name, letStmt.Name.TokenLiteral())
}
