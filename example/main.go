package main

import (
	"github.com/k0kubun/pp"
	"github.com/tzmfreedom/go-salesforce-formula-parser"
)

func main() {
	// ast := parser.ParseFile(os.Args[1])
	ast := parser.ParseString("IF(TRUE, \"foo\", IF(FALSE, \"bar\", \"baz\"))")
	pp.Println(ast)
	visitor := &Visitor{}
	_, err := visitor.Visit(ast)
	if err != nil {
		panic(err)
	}
}

type Visitor struct{}

func (v *Visitor) Visit(n interface{}) (interface{}, error) {
	if fn, ok := n.(*parser.FunctionCallNode); ok {
		return v.VisitFunctionCallNode(fn)
	}
	if bn, ok := n.(*parser.BinaryNode); ok {
		return v.VisitBinaryNode(bn)
	}
	if ln, ok := n.(*parser.StringNode); ok {
		return v.VisitLiteralNode(ln)
	}
	if ln, ok := n.(*parser.IntegerNode); ok {
		return v.VisitLiteralNode(ln)
	}
	if ln, ok := n.(*parser.FloatingPointNode); ok {
		return v.VisitLiteralNode(ln)
	}
	if ln, ok := n.(*parser.BooleanNode); ok {
		return v.VisitLiteralNode(ln)
	}
	return nil, nil
}

func (v *Visitor) VisitFunctionCallNode(n *parser.FunctionCallNode) (interface{}, error) {
	pp.Println(n.Name)
	return nil, nil
}

func (v *Visitor) VisitBinaryNode(n *parser.BinaryNode) (interface{}, error) {
	v.Visit(n.Right)
	v.Visit(n.Left)
	pp.Println(n.Op)
	return nil, nil
}

func (v *Visitor) VisitLiteralNode(n interface{}) (interface{}, error) {
	pp.Println(n)
	return nil, nil
}
