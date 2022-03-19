package root

import (
	"database/sql"
	"fmt"
	"sync"
)

var (
	once sync.Once
	db   *sql.DB
)

type DB struct {
	ip       string
	port     int
	dbname   string
	user     string
	password string
}

func InstanceDB() *DB {
	var ip = "localhost"
	var port = 3306
	var dbname = "test"
	var user = "root"
	var password = "123456"
	once.Do(func() {
		dns := fmt.Sprintf("%s:%s@/%s", user, password, dbname)
		db, _ = sql.Open("mysql", dns)
		if err := db.Ping(); err != nil {
			panic("Ping error")
		}
	})
	return &DB{
		ip:       ip,
		port:     port,
		dbname:   dbname,
		user:     user,
		password: password,
	}
}
func (dbs *DB) GetTables() []string {
	query := fmt.Sprintf(`SELECT table_name FROM information_schema.TABLES WHERE table_schema = '%s';`, dbs.dbname)
	rows, err := db.Query(query)
	if err != nil {
		panic(err)
	}
	var tables []string
	for rows.Next() {
		var table string
		err = rows.Scan(&table)
		if err != nil {
			panic(err)
		}
		tables = append(tables, table)
	}
	return tables
}
func (dbs *DB) GetTableSchema(table_name string) []M2M {
	query := fmt.Sprintf(`SELECT ordinal_position, column_name, data_type, IFNULL(character_maximum_length,0) AS length, is_nullable, IFNULL(column_comment,"") AS comment, IFNULL(column_key,"") as column_key, IFNULL(column_default,"") as column_default, IFNULL(extra,"") as extra FROM information_schema.COLUMNS WHERE table_name = '%s' AND table_schema = '%s';`, table_name, dbs.dbname)
	rows, err := db.Query(query)
	if err != nil {
		panic(err)
	}
	var m2ms []M2M
	for rows.Next() {
		var m2m M2M
		err = rows.Scan(
			&m2m.OrdinalPosition,
			&m2m.ColumnName,
			&m2m.DataType,
			&m2m.Length,
			&m2m.IsNullable,
			&m2m.Comment,
			&m2m.ColumnKey,
			&m2m.ColumnDefault,
			&m2m.Extra,
		)
		if err != nil {
			panic(err)
		}
		m2ms = append(m2ms, m2m)
	}
	return m2ms
}
