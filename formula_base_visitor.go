// Code generated from formula.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // formula

import (
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type BaseformulaVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseformulaVisitor) VisitCompilationUnit(ctx *CompilationUnitContext) interface{} {
	return ctx.Expression().Accept(v)
}

func (v *BaseformulaVisitor) VisitPrimaryExpression(ctx *PrimaryExpressionContext) interface{} {
	return ctx.Primary().Accept(v)
}

func (v *BaseformulaVisitor) VisitBinaryExpression(ctx *BinaryExpressionContext) interface{} {
	left := ctx.Expression(0).Accept(v)
	right := ctx.Expression(1).Accept(v)
	op := ctx.GetOp().GetText()
	return &BinaryNode{
		Left:  left,
		Right: right,
		Op:    op,
	}
}

func (v *BaseformulaVisitor) VisitFunctionCall(ctx *FunctionCallContext) interface{} {
	name := ctx.Identifier().GetText()
	args := ctx.Arguments().Accept(v).([]interface{})
	return &FunctionCallNode{
		Name: name,
		Args: args,
	}
}

func (v *BaseformulaVisitor) VisitFieldReference(ctx *FieldReferenceContext) interface{} {
	path := []string{}
	for _, ident := range ctx.AllIdentifier() {
		path = append(path, ident.GetText())
	}
	return &FieldReferenceNode{
		Path: path,
	}
}

func (v *BaseformulaVisitor) VisitPrimary(ctx *PrimaryContext) interface{} {
	if ctx.Literal() != nil {
		return ctx.Literal().Accept(v)
	}
	return ctx.Expression().Accept(v)
}

func (v *BaseformulaVisitor) VisitArguments(ctx *ArgumentsContext) interface{} {
	args := []interface{}{}
	for _, exp := range ctx.AllExpression() {
		args = append(args, exp.Accept(v))
	}
	return args
}

func (v *BaseformulaVisitor) VisitLiteral(ctx *LiteralContext) interface{} {
	if ctx.StringLiteral() != nil {
		txt := ctx.GetText()
		return &StringNode{txt[1 : len(txt)-1]}
	}
	if ctx.IntegerLiteral() != nil {
		val, err := strconv.Atoi(ctx.GetText())
		if err != nil {
			panic(err)
		}
		return &IntegerNode{val}
	}
	if ctx.FloatingPointLiteral() != nil {
		val, err := strconv.ParseFloat(ctx.GetText(), 32)
		if err != nil {
			panic(err)
		}
		return &FloatingPointNode{val}
	}
	if ctx.BooleanLiteral() != nil {
		val, err := strconv.ParseBool(ctx.GetText())
		if err != nil {
			panic(err)
		}
		return &BooleanNode{val}
	}
	return nil
}
