package main

import (
	"dkg-client-go/sparqlParser"
)

func main() {


q := `
PREFIX schema:   <http://schema.org/>
CONSTRUCT { ?s ?p ?o }
WHERE
{
  GRAPH ?g
    { ?s ?p ?o .
    ?s schema:hasVisibility ?v }
}
`

 	_, _, _, err := sparql.ParseQuery(q)
 	if err != nil {
 		panic(err)
 	}

}