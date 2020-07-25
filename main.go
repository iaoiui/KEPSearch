package main

import (
	"fmt"

	"github.com/iaoiui/Neo4j/cmd/neo4jwrapper"
)

// QueryGetKEP はKEP取得用のCipherQueryです
var QueryGetKEP neo4jwrapper.CipherQuery = "match (k:KEP) return k.name"

// QueryGetSIG はSIGを取得するためのCipherQueryです
var QueryGetSIG neo4jwrapper.CipherQuery = "match (k:KEP) return k.SIG"

func main() {
	result, err := neo4jwrapper.Query(neo4jwrapper.GenerateCipherQuery(QueryGetSIG))
	if err != nil {
		fmt.Printf("%vというエラーが発生\n", err)
		return
	}
	neo4jwrapper.ShowQueryResult(result)

}
