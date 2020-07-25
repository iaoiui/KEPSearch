# KEPSearch

This repo is just sample for accessing Neo4J through Go.

KEP(https://github.com/kubernetes/enhancements#enhancements-tracking-spreadsheet) can be retrieved via this tool.

# Usage

Launch Neo4j locally and then they mount KEP.csv as docker volume.
It is better to retrieve KEP from Google Spreadsheet directry.

```
cd docker_neo4j
docker-compose up -d
```

# Verification

```
❯ docker-compose logs 
Attaching to testneo4j
testneo4j | Directories in use:
testneo4j |   home:         /var/lib/neo4j
testneo4j |   config:       /var/lib/neo4j/conf
testneo4j |   logs:         /logs
testneo4j |   plugins:      /plugins
testneo4j |   import:       NOT SET
testneo4j |   data:         /var/lib/neo4j/data
testneo4j |   certificates: /var/lib/neo4j/certificates
testneo4j |   run:          /var/lib/neo4j/run
testneo4j | Starting Neo4j.
testneo4j | 2020-07-25 08:51:51.231+0000 INFO  Starting...
testneo4j | 2020-07-25 08:51:54.882+0000 INFO  ======== Neo4j 4.1.1 ========
testneo4j | 2020-07-25 08:51:57.508+0000 INFO  Performing postInitialization step for component 'security-users' with version 2 and status CURRENT
testneo4j | 2020-07-25 08:51:57.511+0000 INFO  Updating the initial password in component 'security-users'  
testneo4j | 2020-07-25 08:51:58.270+0000 INFO  Bolt enabled on 0.0.0.0:7687.
testneo4j | 2020-07-25 08:52:00.031+0000 INFO  Remote interface available at http://localhost:7474/
testneo4j | 2020-07-25 08:52:00.034+0000 INFO  Started.
```

# Configuration

The username and password of Neo4j are hard-coded in cmd/neo4jwrapper/neo4jwrapper.go.
- username: neo4j
- password: neo4j
