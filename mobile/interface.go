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

// Contains perverted wrappers to allow crossing over empty interfaces.

package geth

import (
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

// Interface represents a wrapped version of Go's interface{}, with the capacity
// to store arbitrary data types.
//
// Since it's impossible to get the arbitrary-ness converted between Go and mobile
// platforms, we're using explicit getters and setters for the conversions. There
// is of course no point in enumerating everything, just enough to support the
// contract bindins requiring client side generated code.
type Interface struct {
	object interface{}
}

// NewInterface creates a new empty interface that can be used to pass around
// generic types.
func NewInterface() *Interface { log.DebugLog()
	return new(Interface)
}

func (i *Interface) SetBool(b bool)                { log.DebugLog() i.object = &b }
func (i *Interface) SetBools(bs []bool)            { log.DebugLog() i.object = &bs }
func (i *Interface) SetString(str string)          { log.DebugLog() i.object = &str }
func (i *Interface) SetStrings(strs *Strings)      { log.DebugLog() i.object = &strs.strs }
func (i *Interface) SetBinary(binary []byte)       { log.DebugLog() b := common.CopyBytes(binary); i.object = &b }
func (i *Interface) SetBinaries(binaries [][]byte) { log.DebugLog() i.object = &binaries }
func (i *Interface) SetAddress(address *Address)   { log.DebugLog() i.object = &address.address }
func (i *Interface) SetAddresses(addrs *Addresses) { log.DebugLog() i.object = &addrs.addresses }
func (i *Interface) SetHash(hash *Hash)            { log.DebugLog() i.object = &hash.hash }
func (i *Interface) SetHashes(hashes *Hashes)      { log.DebugLog() i.object = &hashes.hashes }
func (i *Interface) SetInt8(n int8)                { log.DebugLog() i.object = &n }
func (i *Interface) SetInt16(n int16)              { log.DebugLog() i.object = &n }
func (i *Interface) SetInt32(n int32)              { log.DebugLog() i.object = &n }
func (i *Interface) SetInt64(n int64)              { log.DebugLog() i.object = &n }
func (i *Interface) SetUint8(bigint *BigInt)       { log.DebugLog() n := uint8(bigint.bigint.Uint64()); i.object = &n }
func (i *Interface) SetUint16(bigint *BigInt)      { log.DebugLog() n := uint16(bigint.bigint.Uint64()); i.object = &n }
func (i *Interface) SetUint32(bigint *BigInt)      { log.DebugLog() n := uint32(bigint.bigint.Uint64()); i.object = &n }
func (i *Interface) SetUint64(bigint *BigInt)      { log.DebugLog() n := bigint.bigint.Uint64(); i.object = &n }
func (i *Interface) SetBigInt(bigint *BigInt)      { log.DebugLog() i.object = &bigint.bigint }
func (i *Interface) SetBigInts(bigints *BigInts)   { log.DebugLog() i.object = &bigints.bigints }

func (i *Interface) SetDefaultBool()      { log.DebugLog() i.object = new(bool) }
func (i *Interface) SetDefaultBools()     { log.DebugLog() i.object = new([]bool) }
func (i *Interface) SetDefaultString()    { log.DebugLog() i.object = new(string) }
func (i *Interface) SetDefaultStrings()   { log.DebugLog() i.object = new([]string) }
func (i *Interface) SetDefaultBinary()    { log.DebugLog() i.object = new([]byte) }
func (i *Interface) SetDefaultBinaries()  { log.DebugLog() i.object = new([][]byte) }
func (i *Interface) SetDefaultAddress()   { log.DebugLog() i.object = new(common.Address) }
func (i *Interface) SetDefaultAddresses() { log.DebugLog() i.object = new([]common.Address) }
func (i *Interface) SetDefaultHash()      { log.DebugLog() i.object = new(common.Hash) }
func (i *Interface) SetDefaultHashes()    { log.DebugLog() i.object = new([]common.Hash) }
func (i *Interface) SetDefaultInt8()      { log.DebugLog() i.object = new(int8) }
func (i *Interface) SetDefaultInt16()     { log.DebugLog() i.object = new(int16) }
func (i *Interface) SetDefaultInt32()     { log.DebugLog() i.object = new(int32) }
func (i *Interface) SetDefaultInt64()     { log.DebugLog() i.object = new(int64) }
func (i *Interface) SetDefaultUint8()     { log.DebugLog() i.object = new(uint8) }
func (i *Interface) SetDefaultUint16()    { log.DebugLog() i.object = new(uint16) }
func (i *Interface) SetDefaultUint32()    { log.DebugLog() i.object = new(uint32) }
func (i *Interface) SetDefaultUint64()    { log.DebugLog() i.object = new(uint64) }
func (i *Interface) SetDefaultBigInt()    { log.DebugLog() i.object = new(*big.Int) }
func (i *Interface) SetDefaultBigInts()   { log.DebugLog() i.object = new([]*big.Int) }

func (i *Interface) GetBool() bool            { log.DebugLog() return *i.object.(*bool) }
func (i *Interface) GetBools() []bool         { log.DebugLog() return *i.object.(*[]bool) }
func (i *Interface) GetString() string        { log.DebugLog() return *i.object.(*string) }
func (i *Interface) GetStrings() *Strings     { log.DebugLog() return &Strings{*i.object.(*[]string)} }
func (i *Interface) GetBinary() []byte        { log.DebugLog() return *i.object.(*[]byte) }
func (i *Interface) GetBinaries() [][]byte    { log.DebugLog() return *i.object.(*[][]byte) }
func (i *Interface) GetAddress() *Address     { log.DebugLog() return &Address{*i.object.(*common.Address)} }
func (i *Interface) GetAddresses() *Addresses { log.DebugLog() return &Addresses{*i.object.(*[]common.Address)} }
func (i *Interface) GetHash() *Hash           { log.DebugLog() return &Hash{*i.object.(*common.Hash)} }
func (i *Interface) GetHashes() *Hashes       { log.DebugLog() return &Hashes{*i.object.(*[]common.Hash)} }
func (i *Interface) GetInt8() int8            { log.DebugLog() return *i.object.(*int8) }
func (i *Interface) GetInt16() int16          { log.DebugLog() return *i.object.(*int16) }
func (i *Interface) GetInt32() int32          { log.DebugLog() return *i.object.(*int32) }
func (i *Interface) GetInt64() int64          { log.DebugLog() return *i.object.(*int64) }
func (i *Interface) GetUint8() *BigInt { log.DebugLog()
	return &BigInt{new(big.Int).SetUint64(uint64(*i.object.(*uint8)))}
}
func (i *Interface) GetUint16() *BigInt { log.DebugLog()
	return &BigInt{new(big.Int).SetUint64(uint64(*i.object.(*uint16)))}
}
func (i *Interface) GetUint32() *BigInt { log.DebugLog()
	return &BigInt{new(big.Int).SetUint64(uint64(*i.object.(*uint32)))}
}
func (i *Interface) GetUint64() *BigInt { log.DebugLog()
	return &BigInt{new(big.Int).SetUint64(*i.object.(*uint64))}
}
func (i *Interface) GetBigInt() *BigInt   { log.DebugLog() return &BigInt{*i.object.(**big.Int)} }
func (i *Interface) GetBigInts() *BigInts { log.DebugLog() return &BigInts{*i.object.(*[]*big.Int)} }

// Interfaces is a slices of wrapped generic objects.
type Interfaces struct {
	objects []interface{}
}

// NewInterfaces creates a slice of uninitialized interfaces.
func NewInterfaces(size int) *Interfaces { log.DebugLog()
	return &Interfaces{
		objects: make([]interface{}, size),
	}
}

// Size returns the number of interfaces in the slice.
func (i *Interfaces) Size() int { log.DebugLog()
	return len(i.objects)
}

// Get returns the bigint at the given index from the slice.
func (i *Interfaces) Get(index int) (iface *Interface, _ error) { log.DebugLog()
	if index < 0 || index >= len(i.objects) {
		return nil, errors.New("index out of bounds")
	}
	return &Interface{i.objects[index]}, nil
}

// Set sets the big int at the given index in the slice.
func (i *Interfaces) Set(index int, object *Interface) error { log.DebugLog()
	if index < 0 || index >= len(i.objects) {
		return errors.New("index out of bounds")
	}
	i.objects[index] = object.object
	return nil
}
