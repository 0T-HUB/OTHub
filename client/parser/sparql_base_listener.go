// Code generated from Sparql.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Sparql

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseSparqlListener is a complete listener for a parse tree produced by SparqlParser.
type BaseSparqlListener struct{}

var _ SparqlListener = &BaseSparqlListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseSparqlListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseSparqlListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseSparqlListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseSparqlListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterQuery is called when production query is entered.
func (s *BaseSparqlListener) EnterQuery(ctx *QueryContext) {}

// ExitQuery is called when production query is exited.
func (s *BaseSparqlListener) ExitQuery(ctx *QueryContext) {}

// EnterPrologue is called when production prologue is entered.
func (s *BaseSparqlListener) EnterPrologue(ctx *PrologueContext) {}

// ExitPrologue is called when production prologue is exited.
func (s *BaseSparqlListener) ExitPrologue(ctx *PrologueContext) {}

// EnterBaseDecl is called when production baseDecl is entered.
func (s *BaseSparqlListener) EnterBaseDecl(ctx *BaseDeclContext) {}

// ExitBaseDecl is called when production baseDecl is exited.
func (s *BaseSparqlListener) ExitBaseDecl(ctx *BaseDeclContext) {}

// EnterPrefixDecl is called when production prefixDecl is entered.
func (s *BaseSparqlListener) EnterPrefixDecl(ctx *PrefixDeclContext) {}

// ExitPrefixDecl is called when production prefixDecl is exited.
func (s *BaseSparqlListener) ExitPrefixDecl(ctx *PrefixDeclContext) {}

// EnterSelectQuery is called when production selectQuery is entered.
func (s *BaseSparqlListener) EnterSelectQuery(ctx *SelectQueryContext) {}

// ExitSelectQuery is called when production selectQuery is exited.
func (s *BaseSparqlListener) ExitSelectQuery(ctx *SelectQueryContext) {}

// EnterConstructQuery is called when production constructQuery is entered.
func (s *BaseSparqlListener) EnterConstructQuery(ctx *ConstructQueryContext) {}

// ExitConstructQuery is called when production constructQuery is exited.
func (s *BaseSparqlListener) ExitConstructQuery(ctx *ConstructQueryContext) {}

// EnterDescribeQuery is called when production describeQuery is entered.
func (s *BaseSparqlListener) EnterDescribeQuery(ctx *DescribeQueryContext) {}

// ExitDescribeQuery is called when production describeQuery is exited.
func (s *BaseSparqlListener) ExitDescribeQuery(ctx *DescribeQueryContext) {}

// EnterAskQuery is called when production askQuery is entered.
func (s *BaseSparqlListener) EnterAskQuery(ctx *AskQueryContext) {}

// ExitAskQuery is called when production askQuery is exited.
func (s *BaseSparqlListener) ExitAskQuery(ctx *AskQueryContext) {}

// EnterDatasetClause is called when production datasetClause is entered.
func (s *BaseSparqlListener) EnterDatasetClause(ctx *DatasetClauseContext) {}

// ExitDatasetClause is called when production datasetClause is exited.
func (s *BaseSparqlListener) ExitDatasetClause(ctx *DatasetClauseContext) {}

// EnterDefaultGraphClause is called when production defaultGraphClause is entered.
func (s *BaseSparqlListener) EnterDefaultGraphClause(ctx *DefaultGraphClauseContext) {}

// ExitDefaultGraphClause is called when production defaultGraphClause is exited.
func (s *BaseSparqlListener) ExitDefaultGraphClause(ctx *DefaultGraphClauseContext) {}

// EnterNamedGraphClause is called when production namedGraphClause is entered.
func (s *BaseSparqlListener) EnterNamedGraphClause(ctx *NamedGraphClauseContext) {}

// ExitNamedGraphClause is called when production namedGraphClause is exited.
func (s *BaseSparqlListener) ExitNamedGraphClause(ctx *NamedGraphClauseContext) {}

// EnterSourceSelector is called when production sourceSelector is entered.
func (s *BaseSparqlListener) EnterSourceSelector(ctx *SourceSelectorContext) {}

// ExitSourceSelector is called when production sourceSelector is exited.
func (s *BaseSparqlListener) ExitSourceSelector(ctx *SourceSelectorContext) {}

// EnterWhereClause is called when production whereClause is entered.
func (s *BaseSparqlListener) EnterWhereClause(ctx *WhereClauseContext) {}

// ExitWhereClause is called when production whereClause is exited.
func (s *BaseSparqlListener) ExitWhereClause(ctx *WhereClauseContext) {}

// EnterSolutionModifier is called when production solutionModifier is entered.
func (s *BaseSparqlListener) EnterSolutionModifier(ctx *SolutionModifierContext) {}

// ExitSolutionModifier is called when production solutionModifier is exited.
func (s *BaseSparqlListener) ExitSolutionModifier(ctx *SolutionModifierContext) {}

// EnterLimitOffsetClauses is called when production limitOffsetClauses is entered.
func (s *BaseSparqlListener) EnterLimitOffsetClauses(ctx *LimitOffsetClausesContext) {}

// ExitLimitOffsetClauses is called when production limitOffsetClauses is exited.
func (s *BaseSparqlListener) ExitLimitOffsetClauses(ctx *LimitOffsetClausesContext) {}

// EnterOrderClause is called when production orderClause is entered.
func (s *BaseSparqlListener) EnterOrderClause(ctx *OrderClauseContext) {}

// ExitOrderClause is called when production orderClause is exited.
func (s *BaseSparqlListener) ExitOrderClause(ctx *OrderClauseContext) {}

// EnterOrderCondition is called when production orderCondition is entered.
func (s *BaseSparqlListener) EnterOrderCondition(ctx *OrderConditionContext) {}

// ExitOrderCondition is called when production orderCondition is exited.
func (s *BaseSparqlListener) ExitOrderCondition(ctx *OrderConditionContext) {}

// EnterLimitClause is called when production limitClause is entered.
func (s *BaseSparqlListener) EnterLimitClause(ctx *LimitClauseContext) {}

// ExitLimitClause is called when production limitClause is exited.
func (s *BaseSparqlListener) ExitLimitClause(ctx *LimitClauseContext) {}

// EnterOffsetClause is called when production offsetClause is entered.
func (s *BaseSparqlListener) EnterOffsetClause(ctx *OffsetClauseContext) {}

// ExitOffsetClause is called when production offsetClause is exited.
func (s *BaseSparqlListener) ExitOffsetClause(ctx *OffsetClauseContext) {}

// EnterGroupGraphPattern is called when production groupGraphPattern is entered.
func (s *BaseSparqlListener) EnterGroupGraphPattern(ctx *GroupGraphPatternContext) {}

// ExitGroupGraphPattern is called when production groupGraphPattern is exited.
func (s *BaseSparqlListener) ExitGroupGraphPattern(ctx *GroupGraphPatternContext) {}

// EnterTriplesBlock is called when production triplesBlock is entered.
func (s *BaseSparqlListener) EnterTriplesBlock(ctx *TriplesBlockContext) {}

// ExitTriplesBlock is called when production triplesBlock is exited.
func (s *BaseSparqlListener) ExitTriplesBlock(ctx *TriplesBlockContext) {}

// EnterGraphPatternNotTriples is called when production graphPatternNotTriples is entered.
func (s *BaseSparqlListener) EnterGraphPatternNotTriples(ctx *GraphPatternNotTriplesContext) {}

// ExitGraphPatternNotTriples is called when production graphPatternNotTriples is exited.
func (s *BaseSparqlListener) ExitGraphPatternNotTriples(ctx *GraphPatternNotTriplesContext) {}

// EnterOptionalGraphPattern is called when production optionalGraphPattern is entered.
func (s *BaseSparqlListener) EnterOptionalGraphPattern(ctx *OptionalGraphPatternContext) {}

// ExitOptionalGraphPattern is called when production optionalGraphPattern is exited.
func (s *BaseSparqlListener) ExitOptionalGraphPattern(ctx *OptionalGraphPatternContext) {}

// EnterGraphGraphPattern is called when production graphGraphPattern is entered.
func (s *BaseSparqlListener) EnterGraphGraphPattern(ctx *GraphGraphPatternContext) {}

// ExitGraphGraphPattern is called when production graphGraphPattern is exited.
func (s *BaseSparqlListener) ExitGraphGraphPattern(ctx *GraphGraphPatternContext) {}

// EnterGroupOrUnionGraphPattern is called when production groupOrUnionGraphPattern is entered.
func (s *BaseSparqlListener) EnterGroupOrUnionGraphPattern(ctx *GroupOrUnionGraphPatternContext) {}

// ExitGroupOrUnionGraphPattern is called when production groupOrUnionGraphPattern is exited.
func (s *BaseSparqlListener) ExitGroupOrUnionGraphPattern(ctx *GroupOrUnionGraphPatternContext) {}

// EnterFilter_ is called when production filter_ is entered.
func (s *BaseSparqlListener) EnterFilter_(ctx *Filter_Context) {}

// ExitFilter_ is called when production filter_ is exited.
func (s *BaseSparqlListener) ExitFilter_(ctx *Filter_Context) {}

// EnterConstraint is called when production constraint is entered.
func (s *BaseSparqlListener) EnterConstraint(ctx *ConstraintContext) {}

// ExitConstraint is called when production constraint is exited.
func (s *BaseSparqlListener) ExitConstraint(ctx *ConstraintContext) {}

// EnterFunctionCall is called when production functionCall is entered.
func (s *BaseSparqlListener) EnterFunctionCall(ctx *FunctionCallContext) {}

// ExitFunctionCall is called when production functionCall is exited.
func (s *BaseSparqlListener) ExitFunctionCall(ctx *FunctionCallContext) {}

// EnterArgList is called when production argList is entered.
func (s *BaseSparqlListener) EnterArgList(ctx *ArgListContext) {}

// ExitArgList is called when production argList is exited.
func (s *BaseSparqlListener) ExitArgList(ctx *ArgListContext) {}

// EnterConstructTemplate is called when production constructTemplate is entered.
func (s *BaseSparqlListener) EnterConstructTemplate(ctx *ConstructTemplateContext) {}

// ExitConstructTemplate is called when production constructTemplate is exited.
func (s *BaseSparqlListener) ExitConstructTemplate(ctx *ConstructTemplateContext) {}

// EnterConstructTriples is called when production constructTriples is entered.
func (s *BaseSparqlListener) EnterConstructTriples(ctx *ConstructTriplesContext) {}

// ExitConstructTriples is called when production constructTriples is exited.
func (s *BaseSparqlListener) ExitConstructTriples(ctx *ConstructTriplesContext) {}

// EnterTriplesSameSubject is called when production triplesSameSubject is entered.
func (s *BaseSparqlListener) EnterTriplesSameSubject(ctx *TriplesSameSubjectContext) {}

// ExitTriplesSameSubject is called when production triplesSameSubject is exited.
func (s *BaseSparqlListener) ExitTriplesSameSubject(ctx *TriplesSameSubjectContext) {}

// EnterPropertyListNotEmpty is called when production propertyListNotEmpty is entered.
func (s *BaseSparqlListener) EnterPropertyListNotEmpty(ctx *PropertyListNotEmptyContext) {}

// ExitPropertyListNotEmpty is called when production propertyListNotEmpty is exited.
func (s *BaseSparqlListener) ExitPropertyListNotEmpty(ctx *PropertyListNotEmptyContext) {}

// EnterPropertyList is called when production propertyList is entered.
func (s *BaseSparqlListener) EnterPropertyList(ctx *PropertyListContext) {}

// ExitPropertyList is called when production propertyList is exited.
func (s *BaseSparqlListener) ExitPropertyList(ctx *PropertyListContext) {}

// EnterObjectList is called when production objectList is entered.
func (s *BaseSparqlListener) EnterObjectList(ctx *ObjectListContext) {}

// ExitObjectList is called when production objectList is exited.
func (s *BaseSparqlListener) ExitObjectList(ctx *ObjectListContext) {}

// EnterObject_ is called when production object_ is entered.
func (s *BaseSparqlListener) EnterObject_(ctx *Object_Context) {}

// ExitObject_ is called when production object_ is exited.
func (s *BaseSparqlListener) ExitObject_(ctx *Object_Context) {}

// EnterVerb is called when production verb is entered.
func (s *BaseSparqlListener) EnterVerb(ctx *VerbContext) {}

// ExitVerb is called when production verb is exited.
func (s *BaseSparqlListener) ExitVerb(ctx *VerbContext) {}

// EnterTriplesNode is called when production triplesNode is entered.
func (s *BaseSparqlListener) EnterTriplesNode(ctx *TriplesNodeContext) {}

// ExitTriplesNode is called when production triplesNode is exited.
func (s *BaseSparqlListener) ExitTriplesNode(ctx *TriplesNodeContext) {}

// EnterBlankNodePropertyList is called when production blankNodePropertyList is entered.
func (s *BaseSparqlListener) EnterBlankNodePropertyList(ctx *BlankNodePropertyListContext) {}

// ExitBlankNodePropertyList is called when production blankNodePropertyList is exited.
func (s *BaseSparqlListener) ExitBlankNodePropertyList(ctx *BlankNodePropertyListContext) {}

// EnterCollection is called when production collection is entered.
func (s *BaseSparqlListener) EnterCollection(ctx *CollectionContext) {}

// ExitCollection is called when production collection is exited.
func (s *BaseSparqlListener) ExitCollection(ctx *CollectionContext) {}

// EnterGraphNode is called when production graphNode is entered.
func (s *BaseSparqlListener) EnterGraphNode(ctx *GraphNodeContext) {}

// ExitGraphNode is called when production graphNode is exited.
func (s *BaseSparqlListener) ExitGraphNode(ctx *GraphNodeContext) {}

// EnterVarOrTerm is called when production varOrTerm is entered.
func (s *BaseSparqlListener) EnterVarOrTerm(ctx *VarOrTermContext) {}

// ExitVarOrTerm is called when production varOrTerm is exited.
func (s *BaseSparqlListener) ExitVarOrTerm(ctx *VarOrTermContext) {}

// EnterVarOrIRIref is called when production varOrIRIref is entered.
func (s *BaseSparqlListener) EnterVarOrIRIref(ctx *VarOrIRIrefContext) {}

// ExitVarOrIRIref is called when production varOrIRIref is exited.
func (s *BaseSparqlListener) ExitVarOrIRIref(ctx *VarOrIRIrefContext) {}

// EnterVar_ is called when production var_ is entered.
func (s *BaseSparqlListener) EnterVar_(ctx *Var_Context) {}

// ExitVar_ is called when production var_ is exited.
func (s *BaseSparqlListener) ExitVar_(ctx *Var_Context) {}

// EnterGraphTerm is called when production graphTerm is entered.
func (s *BaseSparqlListener) EnterGraphTerm(ctx *GraphTermContext) {}

// ExitGraphTerm is called when production graphTerm is exited.
func (s *BaseSparqlListener) ExitGraphTerm(ctx *GraphTermContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseSparqlListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseSparqlListener) ExitExpression(ctx *ExpressionContext) {}

// EnterConditionalOrExpression is called when production conditionalOrExpression is entered.
func (s *BaseSparqlListener) EnterConditionalOrExpression(ctx *ConditionalOrExpressionContext) {}

// ExitConditionalOrExpression is called when production conditionalOrExpression is exited.
func (s *BaseSparqlListener) ExitConditionalOrExpression(ctx *ConditionalOrExpressionContext) {}

// EnterConditionalAndExpression is called when production conditionalAndExpression is entered.
func (s *BaseSparqlListener) EnterConditionalAndExpression(ctx *ConditionalAndExpressionContext) {}

// ExitConditionalAndExpression is called when production conditionalAndExpression is exited.
func (s *BaseSparqlListener) ExitConditionalAndExpression(ctx *ConditionalAndExpressionContext) {}

// EnterValueLogical is called when production valueLogical is entered.
func (s *BaseSparqlListener) EnterValueLogical(ctx *ValueLogicalContext) {}

// ExitValueLogical is called when production valueLogical is exited.
func (s *BaseSparqlListener) ExitValueLogical(ctx *ValueLogicalContext) {}

// EnterRelationalExpression is called when production relationalExpression is entered.
func (s *BaseSparqlListener) EnterRelationalExpression(ctx *RelationalExpressionContext) {}

// ExitRelationalExpression is called when production relationalExpression is exited.
func (s *BaseSparqlListener) ExitRelationalExpression(ctx *RelationalExpressionContext) {}

// EnterNumericExpression is called when production numericExpression is entered.
func (s *BaseSparqlListener) EnterNumericExpression(ctx *NumericExpressionContext) {}

// ExitNumericExpression is called when production numericExpression is exited.
func (s *BaseSparqlListener) ExitNumericExpression(ctx *NumericExpressionContext) {}

// EnterAdditiveExpression is called when production additiveExpression is entered.
func (s *BaseSparqlListener) EnterAdditiveExpression(ctx *AdditiveExpressionContext) {}

// ExitAdditiveExpression is called when production additiveExpression is exited.
func (s *BaseSparqlListener) ExitAdditiveExpression(ctx *AdditiveExpressionContext) {}

// EnterMultiplicativeExpression is called when production multiplicativeExpression is entered.
func (s *BaseSparqlListener) EnterMultiplicativeExpression(ctx *MultiplicativeExpressionContext) {}

// ExitMultiplicativeExpression is called when production multiplicativeExpression is exited.
func (s *BaseSparqlListener) ExitMultiplicativeExpression(ctx *MultiplicativeExpressionContext) {}

// EnterUnaryExpression is called when production unaryExpression is entered.
func (s *BaseSparqlListener) EnterUnaryExpression(ctx *UnaryExpressionContext) {}

// ExitUnaryExpression is called when production unaryExpression is exited.
func (s *BaseSparqlListener) ExitUnaryExpression(ctx *UnaryExpressionContext) {}

// EnterPrimaryExpression is called when production primaryExpression is entered.
func (s *BaseSparqlListener) EnterPrimaryExpression(ctx *PrimaryExpressionContext) {}

// ExitPrimaryExpression is called when production primaryExpression is exited.
func (s *BaseSparqlListener) ExitPrimaryExpression(ctx *PrimaryExpressionContext) {}

// EnterBrackettedExpression is called when production brackettedExpression is entered.
func (s *BaseSparqlListener) EnterBrackettedExpression(ctx *BrackettedExpressionContext) {}

// ExitBrackettedExpression is called when production brackettedExpression is exited.
func (s *BaseSparqlListener) ExitBrackettedExpression(ctx *BrackettedExpressionContext) {}

// EnterBuiltInCall is called when production builtInCall is entered.
func (s *BaseSparqlListener) EnterBuiltInCall(ctx *BuiltInCallContext) {}

// ExitBuiltInCall is called when production builtInCall is exited.
func (s *BaseSparqlListener) ExitBuiltInCall(ctx *BuiltInCallContext) {}

// EnterRegexExpression is called when production regexExpression is entered.
func (s *BaseSparqlListener) EnterRegexExpression(ctx *RegexExpressionContext) {}

// ExitRegexExpression is called when production regexExpression is exited.
func (s *BaseSparqlListener) ExitRegexExpression(ctx *RegexExpressionContext) {}

// EnterIriRefOrFunction is called when production iriRefOrFunction is entered.
func (s *BaseSparqlListener) EnterIriRefOrFunction(ctx *IriRefOrFunctionContext) {}

// ExitIriRefOrFunction is called when production iriRefOrFunction is exited.
func (s *BaseSparqlListener) ExitIriRefOrFunction(ctx *IriRefOrFunctionContext) {}

// EnterRdfLiteral is called when production rdfLiteral is entered.
func (s *BaseSparqlListener) EnterRdfLiteral(ctx *RdfLiteralContext) {}

// ExitRdfLiteral is called when production rdfLiteral is exited.
func (s *BaseSparqlListener) ExitRdfLiteral(ctx *RdfLiteralContext) {}

// EnterNumericLiteral is called when production numericLiteral is entered.
func (s *BaseSparqlListener) EnterNumericLiteral(ctx *NumericLiteralContext) {}

// ExitNumericLiteral is called when production numericLiteral is exited.
func (s *BaseSparqlListener) ExitNumericLiteral(ctx *NumericLiteralContext) {}

// EnterNumericLiteralUnsigned is called when production numericLiteralUnsigned is entered.
func (s *BaseSparqlListener) EnterNumericLiteralUnsigned(ctx *NumericLiteralUnsignedContext) {}

// ExitNumericLiteralUnsigned is called when production numericLiteralUnsigned is exited.
func (s *BaseSparqlListener) ExitNumericLiteralUnsigned(ctx *NumericLiteralUnsignedContext) {}

// EnterNumericLiteralPositive is called when production numericLiteralPositive is entered.
func (s *BaseSparqlListener) EnterNumericLiteralPositive(ctx *NumericLiteralPositiveContext) {}

// ExitNumericLiteralPositive is called when production numericLiteralPositive is exited.
func (s *BaseSparqlListener) ExitNumericLiteralPositive(ctx *NumericLiteralPositiveContext) {}

// EnterNumericLiteralNegative is called when production numericLiteralNegative is entered.
func (s *BaseSparqlListener) EnterNumericLiteralNegative(ctx *NumericLiteralNegativeContext) {}

// ExitNumericLiteralNegative is called when production numericLiteralNegative is exited.
func (s *BaseSparqlListener) ExitNumericLiteralNegative(ctx *NumericLiteralNegativeContext) {}

// EnterBooleanLiteral is called when production booleanLiteral is entered.
func (s *BaseSparqlListener) EnterBooleanLiteral(ctx *BooleanLiteralContext) {}

// ExitBooleanLiteral is called when production booleanLiteral is exited.
func (s *BaseSparqlListener) ExitBooleanLiteral(ctx *BooleanLiteralContext) {}

// EnterString_ is called when production string_ is entered.
func (s *BaseSparqlListener) EnterString_(ctx *String_Context) {}

// ExitString_ is called when production string_ is exited.
func (s *BaseSparqlListener) ExitString_(ctx *String_Context) {}

// EnterIriRef is called when production iriRef is entered.
func (s *BaseSparqlListener) EnterIriRef(ctx *IriRefContext) {}

// ExitIriRef is called when production iriRef is exited.
func (s *BaseSparqlListener) ExitIriRef(ctx *IriRefContext) {}

// EnterPrefixedName is called when production prefixedName is entered.
func (s *BaseSparqlListener) EnterPrefixedName(ctx *PrefixedNameContext) {}

// ExitPrefixedName is called when production prefixedName is exited.
func (s *BaseSparqlListener) ExitPrefixedName(ctx *PrefixedNameContext) {}

// EnterBlankNode is called when production blankNode is entered.
func (s *BaseSparqlListener) EnterBlankNode(ctx *BlankNodeContext) {}

// ExitBlankNode is called when production blankNode is exited.
func (s *BaseSparqlListener) ExitBlankNode(ctx *BlankNodeContext) {}
