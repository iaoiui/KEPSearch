package neo4jwrapper

import (
	"fmt"

	"github.com/neo4j/neo4j-go-driver/neo4j"
)

// Neo4jQueryHandler はNeo4jへのクエリをハンドルする型です
// type Neo4jQueryHandler func(transaction neo4j.Transaction) (interface{}, error)

// CipherQuery はNeo4jで実行するCipherクエリです
type CipherQuery string

// ShowQueryResult はクエリ結果を表示します
func ShowQueryResult(s []string) {
	for i, r := range s {
		fmt.Println(i, ": \t", r)
	}
}

// GenerateCipherQuery は文字列を元にCipherクエリを生成します
func GenerateCipherQuery(query CipherQuery) func(transaction neo4j.Transaction) (interface{}, error) {
	return func(transaction neo4j.Transaction) (interface{}, error) {
		result, err := transaction.Run(
			string(query),
			map[string]interface{}{})
		if err != nil {
			return nil, err
		}

		s := []string{}
		hasNext := true
		for {
			hasNext = result.Next()
			if hasNext == false {
				return s, nil
			}
			if result.Record().GetByIndex(0) == nil {
				return s, nil
			}
			s = append(s, (result.Record().GetByIndex(0)).(string))
		}
	}

}

// Query はNeo4Jへのクエリラッパ
func Query(handler func(transaction neo4j.Transaction) (interface{}, error)) ([]string, error) {
	// 認証情報をハードコード
	return handleNeo4j("bolt://localhost:7687", "neo4j", "neo4j", false, handler)
}

// handleNeo4j はNeo4Jからクエリ結果を受け取ります
func handleNeo4j(uri, username, password string, encrypted bool, queryHandler func(transaction neo4j.Transaction) (interface{}, error)) ([]string, error) {
	driver, err := neo4j.NewDriver(uri, neo4j.BasicAuth(username, password, ""), func(c *neo4j.Config) {
		c.Encrypted = encrypted
	})
	if err != nil {
		return []string{}, err
	}
	defer driver.Close()

	session, err := driver.Session(neo4j.AccessModeWrite)
	if err != nil {
		return []string{}, err
	}
	defer session.Close()
	greeting, err := session.WriteTransaction(queryHandler)
	if err != nil {
		return []string{}, err
	}

	return greeting.([]string), nil
}

// test は、Neo4Jに適当なデータを作成します
// func test() error {
// 	configForNeo4j40 := func(conf *neo4j.Config) { conf.Encrypted = false }
// 	driver, err := neo4j.NewDriver("bolt://localhost:7687", neo4j.BasicAuth("neo4j", "neo4j", ""), configForNeo4j40)
// 	if err != nil {
// 		return err
// 	}
// 	defer driver.Close()
// 	sessionConfig := neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite}
// 	session, err := driver.NewSession(sessionConfig)
// 	if err != nil {
// 		return err
// 	}
// 	defer session.Close()
// 	result, err := session.Run("CREATE (n:Item { id: $id, name: $name })   RETURN n.id, n.name", map[string]interface{}{
// 		"id":   1,
// 		"name": "Item 1",
// 	})
// 	if err != nil {
// 		return err
// 	}
// 	for result.Next() {
// 		fmt.Printf("Created Item with Id = '%d' and Name = '%s'", result.Record().GetByIndex(0).(int64), result.Record().GetByIndex(1).(string))
// 		fmt.Println("")
// 	}
// 	return result.Err()
// }
