// Copyright (C) 2018 go-dappley authors
//
// This file is part of the go-dappley library.
//
// the go-dappley library is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// the go-dappley library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with the go-dappley library.  If not, see <http://www.gnu.org/licenses/>.
//
package core

import (
	"bytes"
	"container/heap"
	"crypto/sha256"
	"encoding/gob"
	"log"
	"time"
	"github.com/gogo/protobuf/proto"
	"github.com/dappley/go-dappley/core/pb"
)

type BlockHeader struct {
	hash      Hash
	prevHash  Hash
	nonce     int64
	timestamp int64
}

type Block struct {
	header       *BlockHeader
	transactions []*Transaction
}


func NewBlock(prevHash []byte) *Block {
	sortedTransactions := []*Transaction{}

	for GetTxnPoolInstance().Len() > 0 {
		if len(sortedTransactions) < TransactionPoolLimit {
			var transaction = heap.Pop(GetTxnPoolInstance()).(Transaction)
			sortedTransactions = append(sortedTransactions, &transaction)
		}
	}
	return &Block{
		header: &BlockHeader{
			hash:      []byte{},
			prevHash:  prevHash,
			nonce:     0,
			timestamp: time.Now().Unix(),
		},
		transactions: sortedTransactions,
	}
}

func (b *Block) HashTransactions() []byte {
	var txHashes [][]byte
	var txHash [32]byte

	for _, tx := range b.transactions {
		txHashes = append(txHashes, tx.Hash())
	}
	txHash = sha256.Sum256(bytes.Join(txHashes, []byte{}))

	return txHash[:]
}

func (b *Block) Serialize() []byte {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)

	bs := &BlockStream{
		Header: &BlockHeaderStream{
			Hash:      b.header.hash,
			PrevHash:  b.header.prevHash,
			Nonce:     b.header.nonce,
			Timestamp: b.header.timestamp,
		},
		Transactions: b.transactions,
	}

	err := encoder.Encode(bs)
	if err != nil {
		log.Panic(err)
	}
	return result.Bytes()
}

func Deserialize(d []byte) *Block {
	var bs BlockStream
	decoder := gob.NewDecoder(bytes.NewReader(d))
	err := decoder.Decode(&bs)
	if err != nil {
		log.Panic(err)
	}

	return &Block{
		header: &BlockHeader{
			hash:      bs.Header.Hash,
			prevHash:  bs.Header.PrevHash,
			nonce:     bs.Header.Nonce,
			timestamp: bs.Header.Timestamp,
		},
		transactions: bs.Transactions,
	}
}

func (b *Block) SetHash(hash Hash) {
	b.header.hash = hash
}

func (b *Block) GetHash() Hash {
	return b.header.hash
}

func (b *Block) GetPrevHash() Hash {
	return b.header.prevHash
}

func (b *Block) SetNonce(nonce int64) {
	b.header.nonce = nonce
}

func (b *Block) GetNonce() int64 {
	return b.header.nonce
}

func (b *Block) GetTimestamp() int64 {
	return b.header.timestamp
}

func (b *Block) GetTransactions() []*Transaction {
	return b.transactions
}

func (b *Block) ToProto() proto.Message{

	txArray := []*corepb.Transaction{}
	for _,tx := range b.transactions {
		txArray = append(txArray,tx.ToProto().(*corepb.Transaction))
	}

	return &corepb.Block{
		Header:			b.header.ToProto().(*corepb.BlockHeader),
		Transactions:	txArray,
	}
}

func (b *Block) FromProto(pb proto.Message){

	bh := BlockHeader{}
	bh.FromProto(pb.(*corepb.Block).Header)
	b.header = &bh

	txs := []*Transaction{}
	tx := &Transaction{}
	for _,txpb := range pb.(*corepb.Block).Transactions{
		tx.FromProto(txpb)
		txs = append(txs,tx)
	}
	b.transactions = txs
}

func (bh *BlockHeader) ToProto() proto.Message{
	return &corepb.BlockHeader{
		Hash:		bh.hash,
		Prevhash:	bh.prevHash,
		Nonce:		bh.nonce,
		Timestamp:	bh.timestamp,
	}
}

func (bh *BlockHeader) FromProto(pb proto.Message){
	bh.hash = pb.(*corepb.BlockHeader).Hash
	bh.prevHash = pb.(*corepb.BlockHeader).Prevhash
	bh.nonce = pb.(*corepb.BlockHeader).Nonce
	bh.timestamp = pb.(*corepb.BlockHeader).Timestamp
}