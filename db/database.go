/*
    Manages persistent state and keeps track of previous and current builds.
 */
package db

import(
    _ "github.com/mattn/go-sqlite3"
    "database/sql"
    "fmt"
    "os"
    util "../util"
)

var     dataLoc         string      =   os.Getenv("HOME")+"/.config/whiteblock/.gdata"
const   ServerTable     string      =   "servers"
const   TestTable       string      =   "testnets"
const   NodesTable      string      =   "nodes"
const   BuildsTable     string      =   "builds"

var conf = util.GetConfig()

var db *sql.DB

func init(){
    db = getDB()
    db.SetMaxOpenConns(50)
    CheckAndUpdate()
}
func getDB() *sql.DB {
    if _, err := os.Stat(dataLoc); os.IsNotExist(err) {
        dbInit()
    }
    d, err := sql.Open("sqlite3", dataLoc)
    if err != nil {
        panic(err)
    }
    return d
}

func dbInit() {
    _, err := os.Create(dataLoc)
    if err != nil {
        panic(err)
    }
    db = getDB()

    serverSchema := fmt.Sprintf("CREATE TABLE %s (%s,%s,%s, %s,%s,%s);",
        ServerTable,
        
        "id INTEGER PRIMARY KEY AUTOINCREMENT",
        "server_id INTEGER",
        "addr TEXT NOT NULL",
        "nodes INTEGER DEFAULT 0",
        "max INTEGER",
        "name TEXT")

    testSchema := fmt.Sprintf("CREATE TABLE %s (%s,%s,%s,%s,%s);",
        TestTable,
        "id TEXT",
        "blockchain TEXT NOT NULL",
        "nodes INTERGER",
        "image TEXT NOT NULL",
        "ts INTEGER")

    nodesSchema := fmt.Sprintf("CREATE TABLE %s (%s,%s,%s, %s,%s,%s);",
        NodesTable,
        "id TEXT",
        "test_net TEXT",
        "server INTEGER",
        "local_id INTEGER",
        "ip TEXT NOT NULL",
        "label TEXT")

    buildSchema := fmt.Sprintf("CREATE TABLE %s (%s,%s,%s, %s,%s,%s, %s,%s,%s);",
        BuildsTable,
        "id INTEGER PRIMARY KEY AUTOINCREMENT",
        "testnet TEXT",
        "servers TEXT",
        "blockchain TEXT",
        "nodes INTEGER",
        "image TEXT",
        "params TEXT",
        "resources TEXT",
        "environment TEXT")

    versionSchema := fmt.Sprintf("CREATE TABLE meta (%s,%s);",
        "key TEXT",
        "value TEXT",
        )

    _,err = db.Exec(serverSchema)
    if err != nil {
        panic(err)
    }
    _,err = db.Exec(testSchema)
    if err != nil {
        panic(err)
    }
    _,err = db.Exec(nodesSchema)
    if err != nil {
        panic(err)
    }
    _,err = db.Exec(buildSchema)
    if err != nil {
        panic(err)
    }
    _,err = db.Exec(versionSchema)
    if err != nil {
        panic(err)
    }
    InsertLocalServers(db);
    SetVersion(Version)
}

/*
    InsertLocalServers adds the default server(s) to the servers database, allowing immediate use of the application
    without having to register a server
 */
func InsertLocalServers(db *sql.DB) {
    InsertServer("cloud",
        Server{
            Addr:"127.0.0.1",
            Nodes:0,
            Max:conf.MaxNodes,
            SubnetID:1,
            Id:-1,
            Ips:[]string{}})
}
