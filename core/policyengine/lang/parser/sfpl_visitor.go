//
// Copyright (C) 2020 IBM Corporation.
//
// Authors:
// Frederico Araujo <frederico.araujo@ibm.com>
// Teryl Taylor <terylt@ibm.com>
//
// Code generated from Sfpl.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // Sfpl
import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by SfplParser.
type SfplVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by SfplParser#policy.
	VisitPolicy(ctx *PolicyContext) interface{}

	// Visit a parse tree produced by SfplParser#prule.
	VisitPrule(ctx *PruleContext) interface{}

	// Visit a parse tree produced by SfplParser#pfilter.
	VisitPfilter(ctx *PfilterContext) interface{}

	// Visit a parse tree produced by SfplParser#pmacro.
	VisitPmacro(ctx *PmacroContext) interface{}

	// Visit a parse tree produced by SfplParser#plist.
	VisitPlist(ctx *PlistContext) interface{}

	// Visit a parse tree produced by SfplParser#expression.
	VisitExpression(ctx *ExpressionContext) interface{}

	// Visit a parse tree produced by SfplParser#or_expression.
	VisitOr_expression(ctx *Or_expressionContext) interface{}

	// Visit a parse tree produced by SfplParser#and_expression.
	VisitAnd_expression(ctx *And_expressionContext) interface{}

	// Visit a parse tree produced by SfplParser#term.
	VisitTerm(ctx *TermContext) interface{}

	// Visit a parse tree produced by SfplParser#items.
	VisitItems(ctx *ItemsContext) interface{}

	// Visit a parse tree produced by SfplParser#variable.
	VisitVariable(ctx *VariableContext) interface{}

	// Visit a parse tree produced by SfplParser#atom.
	VisitAtom(ctx *AtomContext) interface{}

	// Visit a parse tree produced by SfplParser#text.
	VisitText(ctx *TextContext) interface{}

	// Visit a parse tree produced by SfplParser#binary_operator.
	VisitBinary_operator(ctx *Binary_operatorContext) interface{}

	// Visit a parse tree produced by SfplParser#unary_operator.
	VisitUnary_operator(ctx *Unary_operatorContext) interface{}
}