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

// Contains all the wrappers from the core/types package.

package geth

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/ethereum/go-ethereum/log"
)

// A Nonce is a 64-bit hash which proves (combined with the mix-hash) that
// a sufficient amount of computation has been carried out on a block.
type Nonce struct {
	nonce types.BlockNonce
}

// GetBytes retrieves the byte representation of the block nonce.
func (n *Nonce) GetBytes() []byte {
	log.DebugLog()
	return n.nonce[:]
}

// GetHex retrieves the hex string representation of the block nonce.
func (n *Nonce) GetHex() string {
	log.DebugLog()
	return fmt.Sprintf("0x%x", n.nonce[:])
}

// Bloom represents a 256 bit bloom filter.
type Bloom struct {
	bloom types.Bloom
}

// GetBytes retrieves the byte representation of the bloom filter.
func (b *Bloom) GetBytes() []byte {
	log.DebugLog()
	return b.bloom[:]
}

// GetHex retrieves the hex string representation of the bloom filter.
func (b *Bloom) GetHex() string {
	log.DebugLog()
	return fmt.Sprintf("0x%x", b.bloom[:])
}

// Header represents a block header in the Ethereum blockchain.
type Header struct {
	header *types.Header
}

// NewHeaderFromRLP parses a header from an RLP data dump.
func NewHeaderFromRLP(data []byte) (*Header, error) {
	log.DebugLog()
	h := &Header{
		header: new(types.Header),
	}
	if err := rlp.DecodeBytes(common.CopyBytes(data), h.header); err != nil {
		return nil, err
	}
	return h, nil
}

// EncodeRLP encodes a header into an RLP data dump.
func (h *Header) EncodeRLP() ([]byte, error) {
	log.DebugLog()
	return rlp.EncodeToBytes(h.header)
}

// NewHeaderFromJSON parses a header from an JSON data dump.
func NewHeaderFromJSON(data string) (*Header, error) {
	log.DebugLog()
	h := &Header{
		header: new(types.Header),
	}
	if err := json.Unmarshal([]byte(data), h.header); err != nil {
		return nil, err
	}
	return h, nil
}

// EncodeJSON encodes a header into an JSON data dump.
func (h *Header) EncodeJSON() (string, error) {
	log.DebugLog()
	data, err := json.Marshal(h.header)
	return string(data), err
}

// String implements the fmt.Stringer interface to print some semi-meaningful
// data dump of the header for debugging purposes.
func (h *Header) String() string {
	log.DebugLog()
	return h.header.String()
}

func (h *Header) GetParentHash() *Hash   { log.DebugLog()
											 return &Hash{h.header.ParentHash} }
func (h *Header) GetUncleHash() *Hash    { log.DebugLog()
											 return &Hash{h.header.UncleHash} }
func (h *Header) GetCoinbase() *Address  { log.DebugLog()
											 return &Address{h.header.Coinbase} }
func (h *Header) GetRoot() *Hash         { log.DebugLog()
											 return &Hash{h.header.Root} }
func (h *Header) GetTxHash() *Hash       { log.DebugLog()
											 return &Hash{h.header.TxHash} }
func (h *Header) GetReceiptHash() *Hash  { log.DebugLog()
											 return &Hash{h.header.ReceiptHash} }
func (h *Header) GetBloom() *Bloom       { log.DebugLog()
											 return &Bloom{h.header.Bloom} }
func (h *Header) GetDifficulty() *BigInt { log.DebugLog()
											 return &BigInt{h.header.Difficulty} }
func (h *Header) GetNumber() int64       { log.DebugLog()
											 return h.header.Number.Int64() }
func (h *Header) GetGasLimit() int64     { log.DebugLog()
											 return int64(h.header.GasLimit) }
func (h *Header) GetGasUsed() int64      { log.DebugLog()
											 return int64(h.header.GasUsed) }
func (h *Header) GetTime() int64         { log.DebugLog()
											 return h.header.Time.Int64() }
func (h *Header) GetExtra() []byte       { log.DebugLog()
											 return h.header.Extra }
func (h *Header) GetMixDigest() *Hash    { log.DebugLog()
											 return &Hash{h.header.MixDigest} }
func (h *Header) GetNonce() *Nonce       { log.DebugLog()
											 return &Nonce{h.header.Nonce} }
func (h *Header) GetHash() *Hash         { log.DebugLog()
											 return &Hash{h.header.Hash()} }

// Headers represents a slice of headers.
type Headers struct{ headers []*types.Header }

// Size returns the number of headers in the slice.
func (h *Headers) Size() int {
	log.DebugLog()
	return len(h.headers)
}

// Get returns the header at the given index from the slice.
func (h *Headers) Get(index int) (header *Header, _ error) {
	log.DebugLog()
	if index < 0 || index >= len(h.headers) {
		return nil, errors.New("index out of bounds")
	}
	return &Header{h.headers[index]}, nil
}

// Block represents an entire block in the Ethereum blockchain.
type Block struct {
	block *types.Block
}

// NewBlockFromRLP parses a block from an RLP data dump.
func NewBlockFromRLP(data []byte) (*Block, error) {
	log.DebugLog()
	b := &Block{
		block: new(types.Block),
	}
	if err := rlp.DecodeBytes(common.CopyBytes(data), b.block); err != nil {
		return nil, err
	}
	return b, nil
}

// EncodeRLP encodes a block into an RLP data dump.
func (b *Block) EncodeRLP() ([]byte, error) {
	log.DebugLog()
	return rlp.EncodeToBytes(b.block)
}

// NewBlockFromJSON parses a block from an JSON data dump.
func NewBlockFromJSON(data string) (*Block, error) {
	log.DebugLog()
	b := &Block{
		block: new(types.Block),
	}
	if err := json.Unmarshal([]byte(data), b.block); err != nil {
		return nil, err
	}
	return b, nil
}

// EncodeJSON encodes a block into an JSON data dump.
func (b *Block) EncodeJSON() (string, error) {
	log.DebugLog()
	data, err := json.Marshal(b.block)
	return string(data), err
}

// String implements the fmt.Stringer interface to print some semi-meaningful
// data dump of the block for debugging purposes.
func (b *Block) String() string {
	log.DebugLog()
	return b.block.String()
}

func (b *Block) GetParentHash() *Hash   { log.DebugLog()
											return &Hash{b.block.ParentHash()} }
func (b *Block) GetUncleHash() *Hash    { log.DebugLog()
											return &Hash{b.block.UncleHash()} }
func (b *Block) GetCoinbase() *Address  { log.DebugLog()
											return &Address{b.block.Coinbase()} }
func (b *Block) GetRoot() *Hash         { log.DebugLog()
											return &Hash{b.block.Root()} }
func (b *Block) GetTxHash() *Hash       { log.DebugLog()
											return &Hash{b.block.TxHash()} }
func (b *Block) GetReceiptHash() *Hash  { log.DebugLog()
											return &Hash{b.block.ReceiptHash()} }
func (b *Block) GetBloom() *Bloom       { log.DebugLog()
											return &Bloom{b.block.Bloom()} }
func (b *Block) GetDifficulty() *BigInt { log.DebugLog()
											return &BigInt{b.block.Difficulty()} }
func (b *Block) GetNumber() int64       { log.DebugLog()
											return b.block.Number().Int64() }
func (b *Block) GetGasLimit() int64     { log.DebugLog()
											return int64(b.block.GasLimit()) }
func (b *Block) GetGasUsed() int64      { log.DebugLog()
											return int64(b.block.GasUsed()) }
func (b *Block) GetTime() int64         { log.DebugLog()
											return b.block.Time().Int64() }
func (b *Block) GetExtra() []byte       { log.DebugLog()
											return b.block.Extra() }
func (b *Block) GetMixDigest() *Hash    { log.DebugLog()
											return &Hash{b.block.MixDigest()} }
func (b *Block) GetNonce() int64        { log.DebugLog()
											return int64(b.block.Nonce()) }

func (b *Block) GetHash() *Hash        { log.DebugLog()
										   return &Hash{b.block.Hash()} }
func (b *Block) GetHashNoNonce() *Hash { log.DebugLog()
										   return &Hash{b.block.HashNoNonce()} }

func (b *Block) GetHeader() *Header             { log.DebugLog()
													return &Header{b.block.Header()} }
func (b *Block) GetUncles() *Headers            { log.DebugLog()
													return &Headers{b.block.Uncles()} }
func (b *Block) GetTransactions() *Transactions { log.DebugLog()
													return &Transactions{b.block.Transactions()} }
func (b *Block) GetTransaction(hash *Hash) *Transaction {
	log.DebugLog()
	return &Transaction{b.block.Transaction(hash.hash)}
}

// Transaction represents a single Ethereum transaction.
type Transaction struct {
	tx *types.Transaction
}

// NewTransaction creates a new transaction with the given properties.
func NewTransaction(nonce int64, to *Address, amount *BigInt, gasLimit int64, gasPrice *BigInt, data []byte) *Transaction {
	log.DebugLog()
	return &Transaction{types.NewTransaction(uint64(nonce), to.address, amount.bigint, uint64(gasLimit), gasPrice.bigint, common.CopyBytes(data))}
}

// NewTransactionFromRLP parses a transaction from an RLP data dump.
func NewTransactionFromRLP(data []byte) (*Transaction, error) {
	log.DebugLog()
	tx := &Transaction{
		tx: new(types.Transaction),
	}
	if err := rlp.DecodeBytes(common.CopyBytes(data), tx.tx); err != nil {
		return nil, err
	}
	return tx, nil
}

// EncodeRLP encodes a transaction into an RLP data dump.
func (tx *Transaction) EncodeRLP() ([]byte, error) {
	log.DebugLog()
	return rlp.EncodeToBytes(tx.tx)
}

// NewTransactionFromJSON parses a transaction from an JSON data dump.
func NewTransactionFromJSON(data string) (*Transaction, error) {
	log.DebugLog()
	tx := &Transaction{
		tx: new(types.Transaction),
	}
	if err := json.Unmarshal([]byte(data), tx.tx); err != nil {
		return nil, err
	}
	return tx, nil
}

// EncodeJSON encodes a transaction into an JSON data dump.
func (tx *Transaction) EncodeJSON() (string, error) {
	log.DebugLog()
	data, err := json.Marshal(tx.tx)
	return string(data), err
}

// String implements the fmt.Stringer interface to print some semi-meaningful
// data dump of the transaction for debugging purposes.
func (tx *Transaction) String() string {
	log.DebugLog()
	return tx.tx.String()
}

func (tx *Transaction) GetData() []byte      { log.DebugLog()
												 return tx.tx.Data() }
func (tx *Transaction) GetGas() int64        { log.DebugLog()
												 return int64(tx.tx.Gas()) }
func (tx *Transaction) GetGasPrice() *BigInt { log.DebugLog()
												 return &BigInt{tx.tx.GasPrice()} }
func (tx *Transaction) GetValue() *BigInt    { log.DebugLog()
												 return &BigInt{tx.tx.Value()} }
func (tx *Transaction) GetNonce() int64      { log.DebugLog()
												 return int64(tx.tx.Nonce()) }

func (tx *Transaction) GetHash() *Hash   { log.DebugLog()
											 return &Hash{tx.tx.Hash()} }
func (tx *Transaction) GetCost() *BigInt { log.DebugLog()
											 return &BigInt{tx.tx.Cost()} }

// Deprecated: GetSigHash cannot know which signer to use.
func (tx *Transaction) GetSigHash() *Hash { log.DebugLog()
											  return &Hash{types.HomesteadSigner{}.Hash(tx.tx)} }

// Deprecated: use EthereumClient.TransactionSender
func (tx *Transaction) GetFrom(chainID *BigInt) (address *Address, _ error) {
	log.DebugLog()
	var signer types.Signer = types.HomesteadSigner{}
	if chainID != nil {
		signer = types.NewEIP155Signer(chainID.bigint)
	}
	from, err := types.Sender(signer, tx.tx)
	return &Address{from}, err
}

func (tx *Transaction) GetTo() *Address {
	log.DebugLog()
	if to := tx.tx.To(); to != nil {
		return &Address{*to}
	}
	return nil
}

func (tx *Transaction) WithSignature(sig []byte, chainID *BigInt) (signedTx *Transaction, _ error) {
	log.DebugLog()
	var signer types.Signer = types.HomesteadSigner{}
	if chainID != nil {
		signer = types.NewEIP155Signer(chainID.bigint)
	}
	rawTx, err := tx.tx.WithSignature(signer, common.CopyBytes(sig))
	return &Transaction{rawTx}, err
}

// Transactions represents a slice of transactions.
type Transactions struct{ txs types.Transactions }

// Size returns the number of transactions in the slice.
func (txs *Transactions) Size() int {
	log.DebugLog()
	return len(txs.txs)
}

// Get returns the transaction at the given index from the slice.
func (txs *Transactions) Get(index int) (tx *Transaction, _ error) {
	log.DebugLog()
	if index < 0 || index >= len(txs.txs) {
		return nil, errors.New("index out of bounds")
	}
	return &Transaction{txs.txs[index]}, nil
}

// Receipt represents the results of a transaction.
type Receipt struct {
	receipt *types.Receipt
}

// NewReceiptFromRLP parses a transaction receipt from an RLP data dump.
func NewReceiptFromRLP(data []byte) (*Receipt, error) {
	log.DebugLog()
	r := &Receipt{
		receipt: new(types.Receipt),
	}
	if err := rlp.DecodeBytes(common.CopyBytes(data), r.receipt); err != nil {
		return nil, err
	}
	return r, nil
}

// EncodeRLP encodes a transaction receipt into an RLP data dump.
func (r *Receipt) EncodeRLP() ([]byte, error) {
	log.DebugLog()
	return rlp.EncodeToBytes(r.receipt)
}

// NewReceiptFromJSON parses a transaction receipt from an JSON data dump.
func NewReceiptFromJSON(data string) (*Receipt, error) {
	log.DebugLog()
	r := &Receipt{
		receipt: new(types.Receipt),
	}
	if err := json.Unmarshal([]byte(data), r.receipt); err != nil {
		return nil, err
	}
	return r, nil
}

// EncodeJSON encodes a transaction receipt into an JSON data dump.
func (r *Receipt) EncodeJSON() (string, error) {
	log.DebugLog()
	data, err := rlp.EncodeToBytes(r.receipt)
	return string(data), err
}

// String implements the fmt.Stringer interface to print some semi-meaningful
// data dump of the transaction receipt for debugging purposes.
func (r *Receipt) String() string {
	log.DebugLog()
	return r.receipt.String()
}

func (r *Receipt) GetPostState() []byte         { log.DebugLog()
													return r.receipt.PostState }
func (r *Receipt) GetCumulativeGasUsed() int64  { log.DebugLog()
													return int64(r.receipt.CumulativeGasUsed) }
func (r *Receipt) GetBloom() *Bloom             { log.DebugLog()
													return &Bloom{r.receipt.Bloom} }
func (r *Receipt) GetLogs() *Logs               { log.DebugLog()
													return &Logs{r.receipt.Logs} }
func (r *Receipt) GetTxHash() *Hash             { log.DebugLog()
													return &Hash{r.receipt.TxHash} }
func (r *Receipt) GetContractAddress() *Address { log.DebugLog()
													return &Address{r.receipt.ContractAddress} }
func (r *Receipt) GetGasUsed() int64            { log.DebugLog()
													return int64(r.receipt.GasUsed) }
