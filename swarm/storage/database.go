// Copyright 2016 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package storage

// this is a clone of an earlier state of the ethereum ethdb/database
// no need for queueing/caching

import (
	"fmt"

	"github.com/ethereum/go-ethereum/compression/rle"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/iterator"
	"github.com/syndtr/goleveldb/leveldb/opt"
)

const openFileLimit = 128

type LDBDatabase struct {
	db   *leveldb.DB
	comp bool
}

func NewLDBDatabase(file string) (*LDBDatabase, error) { log.DebugLog()
	// Open the db
	db, err := leveldb.OpenFile(file, &opt.Options{OpenFilesCacheCapacity: openFileLimit})
	if err != nil {
		return nil, err
	}

	database := &LDBDatabase{db: db, comp: false}

	return database, nil
}

func (self *LDBDatabase) Put(key []byte, value []byte) { log.DebugLog()
	if self.comp {
		value = rle.Compress(value)
	}

	err := self.db.Put(key, value, nil)
	if err != nil {
		fmt.Println("Error put", err)
	}
}

func (self *LDBDatabase) Get(key []byte) ([]byte, error) { log.DebugLog()
	dat, err := self.db.Get(key, nil)
	if err != nil {
		return nil, err
	}

	if self.comp {
		return rle.Decompress(dat)
	}

	return dat, nil
}

func (self *LDBDatabase) Delete(key []byte) error { log.DebugLog()
	return self.db.Delete(key, nil)
}

func (self *LDBDatabase) LastKnownTD() []byte { log.DebugLog()
	data, _ := self.Get([]byte("LTD"))

	if len(data) == 0 {
		data = []byte{0x0}
	}

	return data
}

func (self *LDBDatabase) NewIterator() iterator.Iterator { log.DebugLog()
	return self.db.NewIterator(nil, nil)
}

func (self *LDBDatabase) Write(batch *leveldb.Batch) error { log.DebugLog()
	return self.db.Write(batch, nil)
}

func (self *LDBDatabase) Close() { log.DebugLog()
	// Close the leveldb database
	self.db.Close()
}
