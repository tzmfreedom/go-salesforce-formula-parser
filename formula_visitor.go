// Code generated from formula.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // formula

import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by formulaParser.
type formulaVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by formulaParser#compilationUnit.
	VisitCompilationUnit(ctx *CompilationUnitContext) interface{}

	// Visit a parse tree produced by formulaParser#primaryExpression.
	VisitPrimaryExpression(ctx *PrimaryExpressionContext) interface{}

	// Visit a parse tree produced by formulaParser#binaryExpression.
	VisitBinaryExpression(ctx *BinaryExpressionContext) interface{}

	// Visit a parse tree produced by formulaParser#functionCall.
	VisitFunctionCall(ctx *FunctionCallContext) interface{}

	// Visit a parse tree produced by formulaParser#fieldReference.
	VisitFieldReference(ctx *FieldReferenceContext) interface{}

	// Visit a parse tree produced by formulaParser#primary.
	VisitPrimary(ctx *PrimaryContext) interface{}

	// Visit a parse tree produced by formulaParser#arguments.
	VisitArguments(ctx *ArgumentsContext) interface{}

	// Visit a parse tree produced by formulaParser#literal.
	VisitLiteral(ctx *LiteralContext) interface{}
}
