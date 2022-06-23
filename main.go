package main

import (
  // "fmt"
	// "dkg-client-go/client"
  "dkg-client-go/parser"
  // "github.com/antlr/antlr4/runtime/Go/antlr"
)

var query = `PREFIX schema: <http://schema.org/>
            CONSTRUCT { ?s ?p ?o }
            WHERE {
                GRAPH ?g {
                ?s ?p ?o .
                ?s schema:hasVisibility ?v
            }}`

func main() {
  sc := parser.NewSparqlSyntaxCheck()
  err := sc.Check(query)
  if err != nil {
    panic(err)
  }
}