package main

import (
	"dkg-client-go/sparkqlParser"
)

func main() {

	q := `
	SELECT ?dcid,
				typeOf ?parent_node Place,
				typeOf ?node Place,
				subType ?node City,
				countryAlpha2Code ?node "country-code",
				containedInPlace ?node ?parent_node,
				dcid ?parent_node "dc/x333",
				dcid ?node ?dcid`

 	_, _, _, err := sparql.ParseQuery(q)
 	if err != nil {
 		panic(err)
 	}

}