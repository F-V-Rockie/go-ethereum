// Copyright 2017 The go-ethereum Authors
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

package enr

import (
	"crypto/ecdsa"
	"fmt"
	"io"
	"net"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/ethereum/go-ethereum/log"
)

// Entry is implemented by known node record entry types.
//
// To define a new entry that is to be included in a node record,
// create a Go type that satisfies this interface. The type should
// also implement rlp.Decoder if additional checks are needed on the value.
type Entry interface {
	ENRKey() string
}

type generic struct {
	key   string
	value interface{}
}

func (g generic) ENRKey() string { log.DebugLog()
									 return g.key }

func (g generic) EncodeRLP(w io.Writer) error {
	log.DebugLog()
	return rlp.Encode(w, g.value)
}

func (g *generic) DecodeRLP(s *rlp.Stream) error {
	log.DebugLog()
	return s.Decode(g.value)
}

// WithEntry wraps any value with a key name. It can be used to set and load arbitrary values
// in a record. The value v must be supported by rlp. To use WithEntry with Load, the value
// must be a pointer.
func WithEntry(k string, v interface{}) Entry {
	log.DebugLog()
	return &generic{key: k, value: v}
}

// DiscPort is the "discv5" key, which holds the UDP port for discovery v5.
type DiscPort uint16

func (v DiscPort) ENRKey() string { log.DebugLog()
									  return "discv5" }

// ID is the "id" key, which holds the name of the identity scheme.
type ID string

func (v ID) ENRKey() string { log.DebugLog()
								return "id" }

// IP4 is the "ip4" key, which holds a 4-byte IPv4 address.
type IP4 net.IP

func (v IP4) ENRKey() string { log.DebugLog()
								 return "ip4" }

// EncodeRLP implements rlp.Encoder.
func (v IP4) EncodeRLP(w io.Writer) error {
	log.DebugLog()
	ip4 := net.IP(v).To4()
	if ip4 == nil {
		return fmt.Errorf("invalid IPv4 address: %v", v)
	}
	return rlp.Encode(w, ip4)
}

// DecodeRLP implements rlp.Decoder.
func (v *IP4) DecodeRLP(s *rlp.Stream) error {
	log.DebugLog()
	if err := s.Decode((*net.IP)(v)); err != nil {
		return err
	}
	if len(*v) != 4 {
		return fmt.Errorf("invalid IPv4 address, want 4 bytes: %v", *v)
	}
	return nil
}

// IP6 is the "ip6" key, which holds a 16-byte IPv6 address.
type IP6 net.IP

func (v IP6) ENRKey() string { log.DebugLog()
								 return "ip6" }

// EncodeRLP implements rlp.Encoder.
func (v IP6) EncodeRLP(w io.Writer) error {
	log.DebugLog()
	ip6 := net.IP(v)
	return rlp.Encode(w, ip6)
}

// DecodeRLP implements rlp.Decoder.
func (v *IP6) DecodeRLP(s *rlp.Stream) error {
	log.DebugLog()
	if err := s.Decode((*net.IP)(v)); err != nil {
		return err
	}
	if len(*v) != 16 {
		return fmt.Errorf("invalid IPv6 address, want 16 bytes: %v", *v)
	}
	return nil
}

// Secp256k1 is the "secp256k1" key, which holds a public key.
type Secp256k1 ecdsa.PublicKey

func (v Secp256k1) ENRKey() string { log.DebugLog()
									   return "secp256k1" }

// EncodeRLP implements rlp.Encoder.
func (v Secp256k1) EncodeRLP(w io.Writer) error {
	log.DebugLog()
	return rlp.Encode(w, crypto.CompressPubkey((*ecdsa.PublicKey)(&v)))
}

// DecodeRLP implements rlp.Decoder.
func (v *Secp256k1) DecodeRLP(s *rlp.Stream) error {
	log.DebugLog()
	buf, err := s.Bytes()
	if err != nil {
		return err
	}
	pk, err := crypto.DecompressPubkey(buf)
	if err != nil {
		return err
	}
	*v = (Secp256k1)(*pk)
	return nil
}

// KeyError is an error related to a key.
type KeyError struct {
	Key string
	Err error
}

// Error implements error.
func (err *KeyError) Error() string {
	log.DebugLog()
	if err.Err == errNotFound {
		return fmt.Sprintf("missing ENR key %q", err.Key)
	}
	return fmt.Sprintf("ENR key %q: %v", err.Key, err.Err)
}

// IsNotFound reports whether the given error means that a key/value pair is
// missing from a record.
func IsNotFound(err error) bool {
	log.DebugLog()
	kerr, ok := err.(*KeyError)
	return ok && kerr.Err == errNotFound
}
