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

// Contains all the wrappers from the go-ethereum root package.

package geth

import (
	"errors"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
)

// Subscription represents an event subscription where events are
// delivered on a data channel.
type Subscription struct {
	sub ethereum.Subscription
}

// Unsubscribe cancels the sending of events to the data channel
// and closes the error channel.
func (s *Subscription) Unsubscribe() {
	log.DebugLog()
	s.sub.Unsubscribe()
}

// CallMsg contains parameters for contract calls.
type CallMsg struct {
	msg ethereum.CallMsg
}

// NewCallMsg creates an empty contract call parameter list.
func NewCallMsg() *CallMsg {
	log.DebugLog()
	return new(CallMsg)
}

func (msg *CallMsg) GetFrom() *Address {
	log.DebugLog()
	return &Address{msg.msg.From}
}
func (msg *CallMsg) GetGas() int64 {
	log.DebugLog()
	return int64(msg.msg.Gas)
}
func (msg *CallMsg) GetGasPrice() *BigInt {
	log.DebugLog()
	return &BigInt{msg.msg.GasPrice}
}
func (msg *CallMsg) GetValue() *BigInt {
	log.DebugLog()
	return &BigInt{msg.msg.Value}
}
func (msg *CallMsg) GetData() []byte {
	log.DebugLog()
	return msg.msg.Data
}
func (msg *CallMsg) GetTo() *Address {
	log.DebugLog()
	if to := msg.msg.To; to != nil {
		return &Address{*msg.msg.To}
	}
	return nil
}

func (msg *CallMsg) SetFrom(address *Address) {
	log.DebugLog()
	msg.msg.From = address.address
}
func (msg *CallMsg) SetGas(gas int64) {
	log.DebugLog()
	msg.msg.Gas = uint64(gas)
}
func (msg *CallMsg) SetGasPrice(price *BigInt) {
	log.DebugLog()
	msg.msg.GasPrice = price.bigint
}
func (msg *CallMsg) SetValue(value *BigInt) {
	log.DebugLog()
	msg.msg.Value = value.bigint
}
func (msg *CallMsg) SetData(data []byte) {
	log.DebugLog()
	msg.msg.Data = common.CopyBytes(data)
}
func (msg *CallMsg) SetTo(address *Address) {
	log.DebugLog()
	if address == nil {
		msg.msg.To = nil
	}
	msg.msg.To = &address.address
}

// SyncProgress gives progress indications when the node is synchronising with
// the Ethereum network.
type SyncProgress struct {
	progress ethereum.SyncProgress
}

func (p *SyncProgress) GetStartingBlock() int64 {
	log.DebugLog()
	return int64(p.progress.StartingBlock)
}
func (p *SyncProgress) GetCurrentBlock() int64 {
	log.DebugLog()
	return int64(p.progress.CurrentBlock)
}
func (p *SyncProgress) GetHighestBlock() int64 {
	log.DebugLog()
	return int64(p.progress.HighestBlock)
}
func (p *SyncProgress) GetPulledStates() int64 {
	log.DebugLog()
	return int64(p.progress.PulledStates)
}
func (p *SyncProgress) GetKnownStates() int64 {
	log.DebugLog()
	return int64(p.progress.KnownStates)
}

// Topics is a set of topic lists to filter events with.
type Topics struct{ topics [][]common.Hash }

// NewTopics creates a slice of uninitialized Topics.
func NewTopics(size int) *Topics {
	log.DebugLog()
	return &Topics{
		topics: make([][]common.Hash, size),
	}
}

// NewTopicsEmpty creates an empty slice of Topics values.
func NewTopicsEmpty() *Topics {
	log.DebugLog()
	return NewTopics(0)
}

// Size returns the number of topic lists inside the set
func (t *Topics) Size() int {
	log.DebugLog()
	return len(t.topics)
}

// Get returns the topic list at the given index from the slice.
func (t *Topics) Get(index int) (hashes *Hashes, _ error) {
	log.DebugLog()
	if index < 0 || index >= len(t.topics) {
		return nil, errors.New("index out of bounds")
	}
	return &Hashes{t.topics[index]}, nil
}

// Set sets the topic list at the given index in the slice.
func (t *Topics) Set(index int, topics *Hashes) error {
	log.DebugLog()
	if index < 0 || index >= len(t.topics) {
		return errors.New("index out of bounds")
	}
	t.topics[index] = topics.hashes
	return nil
}

// Append adds a new topic list to the end of the slice.
func (t *Topics) Append(topics *Hashes) {
	log.DebugLog()
	t.topics = append(t.topics, topics.hashes)
}

// FilterQuery contains options for contact log filtering.
type FilterQuery struct {
	query ethereum.FilterQuery
}

// NewFilterQuery creates an empty filter query for contact log filtering.
func NewFilterQuery() *FilterQuery {
	log.DebugLog()
	return new(FilterQuery)
}

func (fq *FilterQuery) GetFromBlock() *BigInt {
	log.DebugLog()
	return &BigInt{fq.query.FromBlock}
}
func (fq *FilterQuery) GetToBlock() *BigInt {
	log.DebugLog()
	return &BigInt{fq.query.ToBlock}
}
func (fq *FilterQuery) GetAddresses() *Addresses {
	log.DebugLog()
	return &Addresses{fq.query.Addresses}
}
func (fq *FilterQuery) GetTopics() *Topics {
	log.DebugLog()
	return &Topics{fq.query.Topics}
}

func (fq *FilterQuery) SetFromBlock(fromBlock *BigInt) {
	log.DebugLog()
	fq.query.FromBlock = fromBlock.bigint
}
func (fq *FilterQuery) SetToBlock(toBlock *BigInt) {
	log.DebugLog()
	fq.query.ToBlock = toBlock.bigint
}
func (fq *FilterQuery) SetAddresses(addresses *Addresses) {
	log.DebugLog()
	fq.query.Addresses = addresses.addresses
}
func (fq *FilterQuery) SetTopics(topics *Topics) {
	log.DebugLog()
	fq.query.Topics = topics.topics
}
