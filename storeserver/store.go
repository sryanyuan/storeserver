package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var (
	storeSession *sql.DB
)

type StoreKVPair struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func storeInit(source string) error {
	var err error
	storeSession, err = sql.Open("sqlite3", source)
	if nil != err {
		return err
	}
	if err = storeSession.Ping(); nil != err {
		storeUninit()
		return err
	}
	// initialize table
	expr := `CREATE TABLE IF NOT EXISTS kvstore (id INTEGER PRIMARY KEY AUTOINCREMENT, key VARCHAR(128) UNIQUE, value VARCHAR(256))`
	if _, err = storeSession.Exec(expr); nil != err {
		storeUninit()
		return err
	}
	return nil
}

func storeUninit() {
	if nil != storeSession {
		storeSession.Close()
		storeSession = nil
	}
}

func storeGet(key string) (string, error) {
	if len(key) == 0 {
		return "", nil
	}
	row := storeSession.QueryRow("SELECT value FROM kvstore WHERE key = ?", key)
	var value string
	if err := row.Scan(
		&value,
	); nil != err {
		return "", err
	}

	return value, nil
}

func storeSet(key, value string) error {
	if len(key) == 0 {
		return nil
	}
	_, err := storeSession.Exec("REPLACE INTO kvstore (key, value) VALUES (?, ?)", key, value)
	if nil != err {
		return err
	}
	return nil
}

func storeDelete(key string) error {
	if len(key) == 0 {
		return nil
	}
	_, err := storeSession.Exec("DELETE FROM kvstore WHERE key = ?", key)
	if nil != err {
		return err
	}
	return nil
}

func storeGetAll() ([]StoreKVPair, error) {
	rows, err := storeSession.Query("SELECT key, value FROM kvstore")
	if nil != err {
		return nil, err
	}
	defer rows.Close()

	results := make([]StoreKVPair, 0, 128)
	for rows.Next() {
		var pair StoreKVPair
		if err = rows.Scan(
			&pair.Key,
			&pair.Value,
		); nil != err {
			return nil, err
		}
		results = append(results, pair)
	}

	return results, nil
}
