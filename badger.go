package main

import (
	"fmt"

	"github.com/dgraph-io/badger"
)

// BadgerDB struct
type BadgerDB struct {
	db *badger.DB
}

// GBadgerDB global BadgerDB
var GBadgerDB *BadgerDB

// GetBadgerDB Get Persist
func GetBadgerDB() *BadgerDB {

	if GBadgerDB != nil {
		return GBadgerDB
	}

	GBadgerDB = new(BadgerDB)

	GBadgerDB.db = newTestBadger("data", nil)

	fmt.Println("Init BadgerDB!")

	return GBadgerDB
}

func (badgerDB *BadgerDB) Write(key string, value []byte) error {

	givenKey := []byte(key)
	// set
	err := badgerDB.db.Update(func(txn *badger.Txn) error {
		err := txn.Set(givenKey, value)
		return err
	})

	if err != nil {
		return err
	}
	return nil
}

func (badgerDB *BadgerDB) Read(key string) ([]byte, error) {
	givenKey := []byte(key)

	var value []byte

	err := badgerDB.db.View(func(txn *badger.Txn) error {
		item, err := txn.Get(givenKey)
		if err != nil {
			return err
		}
		value, err = item.Value()
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return nil, err
	}
	return value, nil
}

func newTestBadger(name string, applyOpts func(opts *badger.Options)) *badger.DB {

	opts := badger.DefaultOptions
	opts.Dir = "./badger-data"
	opts.ValueDir = "./badger-data"

	db, err := badger.Open(opts)
	if err != nil {
		return nil
	}

	return db
}
