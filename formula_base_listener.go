// Code generated from formula.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // formula

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseformulaListener is a complete listener for a parse tree produced by formulaParser.
type BaseformulaListener struct{}

var _ formulaListener = &BaseformulaListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseformulaListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseformulaListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseformulaListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseformulaListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterCompilationUnit is called when production compilationUnit is entered.
func (s *BaseformulaListener) EnterCompilationUnit(ctx *CompilationUnitContext) {}

// ExitCompilationUnit is called when production compilationUnit is exited.
func (s *BaseformulaListener) ExitCompilationUnit(ctx *CompilationUnitContext) {}

// EnterPrimaryExpression is called when production primaryExpression is entered.
func (s *BaseformulaListener) EnterPrimaryExpression(ctx *PrimaryExpressionContext) {}

// ExitPrimaryExpression is called when production primaryExpression is exited.
func (s *BaseformulaListener) ExitPrimaryExpression(ctx *PrimaryExpressionContext) {}

// EnterBinaryExpression is called when production binaryExpression is entered.
func (s *BaseformulaListener) EnterBinaryExpression(ctx *BinaryExpressionContext) {}

// ExitBinaryExpression is called when production binaryExpression is exited.
func (s *BaseformulaListener) ExitBinaryExpression(ctx *BinaryExpressionContext) {}

// EnterFunctionCall is called when production functionCall is entered.
func (s *BaseformulaListener) EnterFunctionCall(ctx *FunctionCallContext) {}

// ExitFunctionCall is called when production functionCall is exited.
func (s *BaseformulaListener) ExitFunctionCall(ctx *FunctionCallContext) {}

// EnterFieldReference is called when production fieldReference is entered.
func (s *BaseformulaListener) EnterFieldReference(ctx *FieldReferenceContext) {}

// ExitFieldReference is called when production fieldReference is exited.
func (s *BaseformulaListener) ExitFieldReference(ctx *FieldReferenceContext) {}

// EnterPrimary is called when production primary is entered.
func (s *BaseformulaListener) EnterPrimary(ctx *PrimaryContext) {}

// ExitPrimary is called when production primary is exited.
func (s *BaseformulaListener) ExitPrimary(ctx *PrimaryContext) {}

// EnterArguments is called when production arguments is entered.
func (s *BaseformulaListener) EnterArguments(ctx *ArgumentsContext) {}

// ExitArguments is called when production arguments is exited.
func (s *BaseformulaListener) ExitArguments(ctx *ArgumentsContext) {}

// EnterLiteral is called when production literal is entered.
func (s *BaseformulaListener) EnterLiteral(ctx *LiteralContext) {}

// ExitLiteral is called when production literal is exited.
func (s *BaseformulaListener) ExitLiteral(ctx *LiteralContext) {}
