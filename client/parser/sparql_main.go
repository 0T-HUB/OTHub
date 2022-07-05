package parser

import (
	"errors"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type sparqlListener struct {
	*BaseSparqlListener
}

type CustomErrorListener struct {
	*antlr.DefaultErrorListener // Embed default which ensures we fit the interface
	haveError                   bool
	Message                     string
}

func (c *CustomErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	c.haveError = true
	c.Message = msg
}

type SparqlSyntaxCheck struct {
	lexerErrors  *CustomErrorListener
	parserErrors *CustomErrorListener
}

func NewSparqlSyntaxCheck() SparqlSyntaxCheck {
	l := &CustomErrorListener{}
	p := &CustomErrorListener{}

	return SparqlSyntaxCheck{l, p}
}

func (sc *SparqlSyntaxCheck) Check(query string) error {
	// Setup the input
	is := antlr.NewInputStream(query)

	// Create the Lexer
	lexer := NewSparqlLexer(is)
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(sc.lexerErrors)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the Parser
	p := NewSparqlParser(stream)
	p.RemoveErrorListeners()
	p.AddErrorListener(sc.parserErrors)
	p.BuildParseTrees = true

	tree := p.Query()

	// Finally parse the expression
	antlr.ParseTreeWalkerDefault.Walk(sparqlListener{}, tree)

	if sc.lexerErrors.haveError {
		return errors.New(sc.lexerErrors.Message)
	}

	if sc.parserErrors.haveError {
		return errors.New(sc.parserErrors.Message)
	}

	return nil
}
