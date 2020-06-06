package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

const driverName = "sqlite3"

// var dataSourceName = os.TempDir() + strconv.Itoa(int(time.Now().Unix()))
const dataSourceName = ":memory:"
const schemaQuery = `
CREATE TABLE goods (id INTEGER, name TEXT); 
CREATE TABLE tags (id INTEGER, name TEXT);
CREATE TABLE tags_goods (tag_id INTEGER, goods_id INTEGER, UNIQUE (tag_id, goods_id))`

const dataQuery = `
INSERT INTO
	goods
VALUES
	(1, "a"),
	(2, "aa"),
	(3, "aaa"),
	(4, "aaaa"),
	(5, "b"),
	(6, "bb"),
	(7, "bbb"),
	(8, "bbbb"),
	(9, "c"),
	(10, "cc"),
	(11, "ccc"),
	(12, "testName");

INSERT INTO
	tags
VALUES
	(1, "aCategory"),
	(2, "bCategory"),
	(3, "dCategory");
	
INSERT INTO
	tags_goods
VALUES
	(1, 1),
	(1, 2),
	(1, 3),
	(1, 4),
	(2, 5),
	(2, 6),
	(2, 7),
	(2, 8),
	(1, 12),
	(2, 12),
	(3, 12)
`

func NewDb() (*sql.DB, error) {
	db, err := sql.Open(driverName, dataSourceName)

	if err != nil {
		return nil, fmt.Errorf("sql.Open(\"%v\", %v): %w", driverName, dataSourceName, err)
	}

	if _, err := db.Exec(schemaQuery); err != nil {
		return nil, fmt.Errorf("db.Exec(\"%v\"): %w", schemaQuery, err)
	}

	if _, err := db.Exec(dataQuery); err != nil {
		return nil, fmt.Errorf("db.Exec(\"%v\"): %w", dataQuery, err)
	}

	return db, nil
}
