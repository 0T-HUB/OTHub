// Code generated from Sparql.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Sparql

import "github.com/antlr/antlr4/runtime/Go/antlr"

// SparqlListener is a complete listener for a parse tree produced by SparqlParser.
type SparqlListener interface {
	antlr.ParseTreeListener

	// EnterQuery is called when entering the query production.
	EnterQuery(c *QueryContext)

	// EnterPrologue is called when entering the prologue production.
	EnterPrologue(c *PrologueContext)

	// EnterBaseDecl is called when entering the baseDecl production.
	EnterBaseDecl(c *BaseDeclContext)

	// EnterPrefixDecl is called when entering the prefixDecl production.
	EnterPrefixDecl(c *PrefixDeclContext)

	// EnterSelectQuery is called when entering the selectQuery production.
	EnterSelectQuery(c *SelectQueryContext)

	// EnterConstructQuery is called when entering the constructQuery production.
	EnterConstructQuery(c *ConstructQueryContext)

	// EnterDescribeQuery is called when entering the describeQuery production.
	EnterDescribeQuery(c *DescribeQueryContext)

	// EnterAskQuery is called when entering the askQuery production.
	EnterAskQuery(c *AskQueryContext)

	// EnterDatasetClause is called when entering the datasetClause production.
	EnterDatasetClause(c *DatasetClauseContext)

	// EnterDefaultGraphClause is called when entering the defaultGraphClause production.
	EnterDefaultGraphClause(c *DefaultGraphClauseContext)

	// EnterNamedGraphClause is called when entering the namedGraphClause production.
	EnterNamedGraphClause(c *NamedGraphClauseContext)

	// EnterSourceSelector is called when entering the sourceSelector production.
	EnterSourceSelector(c *SourceSelectorContext)

	// EnterWhereClause is called when entering the whereClause production.
	EnterWhereClause(c *WhereClauseContext)

	// EnterSolutionModifier is called when entering the solutionModifier production.
	EnterSolutionModifier(c *SolutionModifierContext)

	// EnterLimitOffsetClauses is called when entering the limitOffsetClauses production.
	EnterLimitOffsetClauses(c *LimitOffsetClausesContext)

	// EnterOrderClause is called when entering the orderClause production.
	EnterOrderClause(c *OrderClauseContext)

	// EnterOrderCondition is called when entering the orderCondition production.
	EnterOrderCondition(c *OrderConditionContext)

	// EnterLimitClause is called when entering the limitClause production.
	EnterLimitClause(c *LimitClauseContext)

	// EnterOffsetClause is called when entering the offsetClause production.
	EnterOffsetClause(c *OffsetClauseContext)

	// EnterGroupGraphPattern is called when entering the groupGraphPattern production.
	EnterGroupGraphPattern(c *GroupGraphPatternContext)

	// EnterTriplesBlock is called when entering the triplesBlock production.
	EnterTriplesBlock(c *TriplesBlockContext)

	// EnterGraphPatternNotTriples is called when entering the graphPatternNotTriples production.
	EnterGraphPatternNotTriples(c *GraphPatternNotTriplesContext)

	// EnterOptionalGraphPattern is called when entering the optionalGraphPattern production.
	EnterOptionalGraphPattern(c *OptionalGraphPatternContext)

	// EnterGraphGraphPattern is called when entering the graphGraphPattern production.
	EnterGraphGraphPattern(c *GraphGraphPatternContext)

	// EnterGroupOrUnionGraphPattern is called when entering the groupOrUnionGraphPattern production.
	EnterGroupOrUnionGraphPattern(c *GroupOrUnionGraphPatternContext)

	// EnterFilter_ is called when entering the filter_ production.
	EnterFilter_(c *Filter_Context)

	// EnterConstraint is called when entering the constraint production.
	EnterConstraint(c *ConstraintContext)

	// EnterFunctionCall is called when entering the functionCall production.
	EnterFunctionCall(c *FunctionCallContext)

	// EnterArgList is called when entering the argList production.
	EnterArgList(c *ArgListContext)

	// EnterConstructTemplate is called when entering the constructTemplate production.
	EnterConstructTemplate(c *ConstructTemplateContext)

	// EnterConstructTriples is called when entering the constructTriples production.
	EnterConstructTriples(c *ConstructTriplesContext)

	// EnterTriplesSameSubject is called when entering the triplesSameSubject production.
	EnterTriplesSameSubject(c *TriplesSameSubjectContext)

	// EnterPropertyListNotEmpty is called when entering the propertyListNotEmpty production.
	EnterPropertyListNotEmpty(c *PropertyListNotEmptyContext)

	// EnterPropertyList is called when entering the propertyList production.
	EnterPropertyList(c *PropertyListContext)

	// EnterObjectList is called when entering the objectList production.
	EnterObjectList(c *ObjectListContext)

	// EnterObject_ is called when entering the object_ production.
	EnterObject_(c *Object_Context)

	// EnterVerb is called when entering the verb production.
	EnterVerb(c *VerbContext)

	// EnterTriplesNode is called when entering the triplesNode production.
	EnterTriplesNode(c *TriplesNodeContext)

	// EnterBlankNodePropertyList is called when entering the blankNodePropertyList production.
	EnterBlankNodePropertyList(c *BlankNodePropertyListContext)

	// EnterCollection is called when entering the collection production.
	EnterCollection(c *CollectionContext)

	// EnterGraphNode is called when entering the graphNode production.
	EnterGraphNode(c *GraphNodeContext)

	// EnterVarOrTerm is called when entering the varOrTerm production.
	EnterVarOrTerm(c *VarOrTermContext)

	// EnterVarOrIRIref is called when entering the varOrIRIref production.
	EnterVarOrIRIref(c *VarOrIRIrefContext)

	// EnterVar_ is called when entering the var_ production.
	EnterVar_(c *Var_Context)

	// EnterGraphTerm is called when entering the graphTerm production.
	EnterGraphTerm(c *GraphTermContext)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterConditionalOrExpression is called when entering the conditionalOrExpression production.
	EnterConditionalOrExpression(c *ConditionalOrExpressionContext)

	// EnterConditionalAndExpression is called when entering the conditionalAndExpression production.
	EnterConditionalAndExpression(c *ConditionalAndExpressionContext)

	// EnterValueLogical is called when entering the valueLogical production.
	EnterValueLogical(c *ValueLogicalContext)

	// EnterRelationalExpression is called when entering the relationalExpression production.
	EnterRelationalExpression(c *RelationalExpressionContext)

	// EnterNumericExpression is called when entering the numericExpression production.
	EnterNumericExpression(c *NumericExpressionContext)

	// EnterAdditiveExpression is called when entering the additiveExpression production.
	EnterAdditiveExpression(c *AdditiveExpressionContext)

	// EnterMultiplicativeExpression is called when entering the multiplicativeExpression production.
	EnterMultiplicativeExpression(c *MultiplicativeExpressionContext)

	// EnterUnaryExpression is called when entering the unaryExpression production.
	EnterUnaryExpression(c *UnaryExpressionContext)

	// EnterPrimaryExpression is called when entering the primaryExpression production.
	EnterPrimaryExpression(c *PrimaryExpressionContext)

	// EnterBrackettedExpression is called when entering the brackettedExpression production.
	EnterBrackettedExpression(c *BrackettedExpressionContext)

	// EnterBuiltInCall is called when entering the builtInCall production.
	EnterBuiltInCall(c *BuiltInCallContext)

	// EnterRegexExpression is called when entering the regexExpression production.
	EnterRegexExpression(c *RegexExpressionContext)

	// EnterIriRefOrFunction is called when entering the iriRefOrFunction production.
	EnterIriRefOrFunction(c *IriRefOrFunctionContext)

	// EnterRdfLiteral is called when entering the rdfLiteral production.
	EnterRdfLiteral(c *RdfLiteralContext)

	// EnterNumericLiteral is called when entering the numericLiteral production.
	EnterNumericLiteral(c *NumericLiteralContext)

	// EnterNumericLiteralUnsigned is called when entering the numericLiteralUnsigned production.
	EnterNumericLiteralUnsigned(c *NumericLiteralUnsignedContext)

	// EnterNumericLiteralPositive is called when entering the numericLiteralPositive production.
	EnterNumericLiteralPositive(c *NumericLiteralPositiveContext)

	// EnterNumericLiteralNegative is called when entering the numericLiteralNegative production.
	EnterNumericLiteralNegative(c *NumericLiteralNegativeContext)

	// EnterBooleanLiteral is called when entering the booleanLiteral production.
	EnterBooleanLiteral(c *BooleanLiteralContext)

	// EnterString_ is called when entering the string_ production.
	EnterString_(c *String_Context)

	// EnterIriRef is called when entering the iriRef production.
	EnterIriRef(c *IriRefContext)

	// EnterPrefixedName is called when entering the prefixedName production.
	EnterPrefixedName(c *PrefixedNameContext)

	// EnterBlankNode is called when entering the blankNode production.
	EnterBlankNode(c *BlankNodeContext)

	// ExitQuery is called when exiting the query production.
	ExitQuery(c *QueryContext)

	// ExitPrologue is called when exiting the prologue production.
	ExitPrologue(c *PrologueContext)

	// ExitBaseDecl is called when exiting the baseDecl production.
	ExitBaseDecl(c *BaseDeclContext)

	// ExitPrefixDecl is called when exiting the prefixDecl production.
	ExitPrefixDecl(c *PrefixDeclContext)

	// ExitSelectQuery is called when exiting the selectQuery production.
	ExitSelectQuery(c *SelectQueryContext)

	// ExitConstructQuery is called when exiting the constructQuery production.
	ExitConstructQuery(c *ConstructQueryContext)

	// ExitDescribeQuery is called when exiting the describeQuery production.
	ExitDescribeQuery(c *DescribeQueryContext)

	// ExitAskQuery is called when exiting the askQuery production.
	ExitAskQuery(c *AskQueryContext)

	// ExitDatasetClause is called when exiting the datasetClause production.
	ExitDatasetClause(c *DatasetClauseContext)

	// ExitDefaultGraphClause is called when exiting the defaultGraphClause production.
	ExitDefaultGraphClause(c *DefaultGraphClauseContext)

	// ExitNamedGraphClause is called when exiting the namedGraphClause production.
	ExitNamedGraphClause(c *NamedGraphClauseContext)

	// ExitSourceSelector is called when exiting the sourceSelector production.
	ExitSourceSelector(c *SourceSelectorContext)

	// ExitWhereClause is called when exiting the whereClause production.
	ExitWhereClause(c *WhereClauseContext)

	// ExitSolutionModifier is called when exiting the solutionModifier production.
	ExitSolutionModifier(c *SolutionModifierContext)

	// ExitLimitOffsetClauses is called when exiting the limitOffsetClauses production.
	ExitLimitOffsetClauses(c *LimitOffsetClausesContext)

	// ExitOrderClause is called when exiting the orderClause production.
	ExitOrderClause(c *OrderClauseContext)

	// ExitOrderCondition is called when exiting the orderCondition production.
	ExitOrderCondition(c *OrderConditionContext)

	// ExitLimitClause is called when exiting the limitClause production.
	ExitLimitClause(c *LimitClauseContext)

	// ExitOffsetClause is called when exiting the offsetClause production.
	ExitOffsetClause(c *OffsetClauseContext)

	// ExitGroupGraphPattern is called when exiting the groupGraphPattern production.
	ExitGroupGraphPattern(c *GroupGraphPatternContext)

	// ExitTriplesBlock is called when exiting the triplesBlock production.
	ExitTriplesBlock(c *TriplesBlockContext)

	// ExitGraphPatternNotTriples is called when exiting the graphPatternNotTriples production.
	ExitGraphPatternNotTriples(c *GraphPatternNotTriplesContext)

	// ExitOptionalGraphPattern is called when exiting the optionalGraphPattern production.
	ExitOptionalGraphPattern(c *OptionalGraphPatternContext)

	// ExitGraphGraphPattern is called when exiting the graphGraphPattern production.
	ExitGraphGraphPattern(c *GraphGraphPatternContext)

	// ExitGroupOrUnionGraphPattern is called when exiting the groupOrUnionGraphPattern production.
	ExitGroupOrUnionGraphPattern(c *GroupOrUnionGraphPatternContext)

	// ExitFilter_ is called when exiting the filter_ production.
	ExitFilter_(c *Filter_Context)

	// ExitConstraint is called when exiting the constraint production.
	ExitConstraint(c *ConstraintContext)

	// ExitFunctionCall is called when exiting the functionCall production.
	ExitFunctionCall(c *FunctionCallContext)

	// ExitArgList is called when exiting the argList production.
	ExitArgList(c *ArgListContext)

	// ExitConstructTemplate is called when exiting the constructTemplate production.
	ExitConstructTemplate(c *ConstructTemplateContext)

	// ExitConstructTriples is called when exiting the constructTriples production.
	ExitConstructTriples(c *ConstructTriplesContext)

	// ExitTriplesSameSubject is called when exiting the triplesSameSubject production.
	ExitTriplesSameSubject(c *TriplesSameSubjectContext)

	// ExitPropertyListNotEmpty is called when exiting the propertyListNotEmpty production.
	ExitPropertyListNotEmpty(c *PropertyListNotEmptyContext)

	// ExitPropertyList is called when exiting the propertyList production.
	ExitPropertyList(c *PropertyListContext)

	// ExitObjectList is called when exiting the objectList production.
	ExitObjectList(c *ObjectListContext)

	// ExitObject_ is called when exiting the object_ production.
	ExitObject_(c *Object_Context)

	// ExitVerb is called when exiting the verb production.
	ExitVerb(c *VerbContext)

	// ExitTriplesNode is called when exiting the triplesNode production.
	ExitTriplesNode(c *TriplesNodeContext)

	// ExitBlankNodePropertyList is called when exiting the blankNodePropertyList production.
	ExitBlankNodePropertyList(c *BlankNodePropertyListContext)

	// ExitCollection is called when exiting the collection production.
	ExitCollection(c *CollectionContext)

	// ExitGraphNode is called when exiting the graphNode production.
	ExitGraphNode(c *GraphNodeContext)

	// ExitVarOrTerm is called when exiting the varOrTerm production.
	ExitVarOrTerm(c *VarOrTermContext)

	// ExitVarOrIRIref is called when exiting the varOrIRIref production.
	ExitVarOrIRIref(c *VarOrIRIrefContext)

	// ExitVar_ is called when exiting the var_ production.
	ExitVar_(c *Var_Context)

	// ExitGraphTerm is called when exiting the graphTerm production.
	ExitGraphTerm(c *GraphTermContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitConditionalOrExpression is called when exiting the conditionalOrExpression production.
	ExitConditionalOrExpression(c *ConditionalOrExpressionContext)

	// ExitConditionalAndExpression is called when exiting the conditionalAndExpression production.
	ExitConditionalAndExpression(c *ConditionalAndExpressionContext)

	// ExitValueLogical is called when exiting the valueLogical production.
	ExitValueLogical(c *ValueLogicalContext)

	// ExitRelationalExpression is called when exiting the relationalExpression production.
	ExitRelationalExpression(c *RelationalExpressionContext)

	// ExitNumericExpression is called when exiting the numericExpression production.
	ExitNumericExpression(c *NumericExpressionContext)

	// ExitAdditiveExpression is called when exiting the additiveExpression production.
	ExitAdditiveExpression(c *AdditiveExpressionContext)

	// ExitMultiplicativeExpression is called when exiting the multiplicativeExpression production.
	ExitMultiplicativeExpression(c *MultiplicativeExpressionContext)

	// ExitUnaryExpression is called when exiting the unaryExpression production.
	ExitUnaryExpression(c *UnaryExpressionContext)

	// ExitPrimaryExpression is called when exiting the primaryExpression production.
	ExitPrimaryExpression(c *PrimaryExpressionContext)

	// ExitBrackettedExpression is called when exiting the brackettedExpression production.
	ExitBrackettedExpression(c *BrackettedExpressionContext)

	// ExitBuiltInCall is called when exiting the builtInCall production.
	ExitBuiltInCall(c *BuiltInCallContext)

	// ExitRegexExpression is called when exiting the regexExpression production.
	ExitRegexExpression(c *RegexExpressionContext)

	// ExitIriRefOrFunction is called when exiting the iriRefOrFunction production.
	ExitIriRefOrFunction(c *IriRefOrFunctionContext)

	// ExitRdfLiteral is called when exiting the rdfLiteral production.
	ExitRdfLiteral(c *RdfLiteralContext)

	// ExitNumericLiteral is called when exiting the numericLiteral production.
	ExitNumericLiteral(c *NumericLiteralContext)

	// ExitNumericLiteralUnsigned is called when exiting the numericLiteralUnsigned production.
	ExitNumericLiteralUnsigned(c *NumericLiteralUnsignedContext)

	// ExitNumericLiteralPositive is called when exiting the numericLiteralPositive production.
	ExitNumericLiteralPositive(c *NumericLiteralPositiveContext)

	// ExitNumericLiteralNegative is called when exiting the numericLiteralNegative production.
	ExitNumericLiteralNegative(c *NumericLiteralNegativeContext)

	// ExitBooleanLiteral is called when exiting the booleanLiteral production.
	ExitBooleanLiteral(c *BooleanLiteralContext)

	// ExitString_ is called when exiting the string_ production.
	ExitString_(c *String_Context)

	// ExitIriRef is called when exiting the iriRef production.
	ExitIriRef(c *IriRefContext)

	// ExitPrefixedName is called when exiting the prefixedName production.
	ExitPrefixedName(c *PrefixedNameContext)

	// ExitBlankNode is called when exiting the blankNode production.
	ExitBlankNode(c *BlankNodeContext)
}
