package parser

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

func ParseFile(f string) interface{} {
	input, _ := antlr.NewFileStream(f)
	return parse(input)
}

func ParseString(data string) interface{} {
	input := antlr.NewInputStream(data)
	return parse(input)
}

func parse(input antlr.CharStream) interface{} {
	lexer := NewformulaLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := NewformulaParser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	p.BuildParseTrees = true
	tree := p.CompilationUnit()
	return tree.Accept(&BaseformulaVisitor{})
}
