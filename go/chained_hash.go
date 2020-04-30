package algorithm

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
)

/************************************************************************
//  同一はハッシュ値を持つデータを鎖状に線形リストでつなげる
//  node1 = node{
//    key   interface{}
//	  value int
//	  next  ここにnode2へのポインタを入れる
//  }
*************************************************************************/

const errUnknownType = "不明な型です"
const errKeyExists = "すでに存在するキーです"
const errKeyNotExist = "存在しないkeyです"
const errNoExist = "目的の値が見つかりませんでした"

type chainedHash struct {
	capacity int
	table    []*intNode
}

type intNode struct {
	key   interface{}
	value int
	next  *intNode
}

func newChainedHash(c int) chainedHash {
	return chainedHash{
		capacity: c,
		table:    make([]*intNode, c),
	}
}

func (h chainedHash) hashValue(key interface{}) (int, error) {
	switch v := key.(type) {
	case int:
		return v % h.capacity, nil

	case string:
		sha := sha256.Sum256([]byte(v))
		dst := make([]byte, hex.EncodedLen(len(sha)))
		hx := hex.Encode(dst, sha[:])
		return hx % h.capacity, nil
	}

	return 0, errors.New(errUnknownType)
}

func (h chainedHash) serch(key interface{}) (int, error) {
	var err error

	hash, err := h.hashValue(key)
	if err != nil {
		return 0, err
	}
	p := h.table[hash]

	for p != nil {
		if p.key == key {
			return p.value, err
		}
		p = p.next
	}

	return 0, errors.New(errNoExist)
}

func (h chainedHash) add(key interface{}, value int) error {
	var err error

	hash, err := h.hashValue(key)
	if err != nil {
		return err
	}
	p := h.table[hash]

	// すでに存在するキーじゃないかチェックする
	for p != nil {
		if p.key == key {
			return errors.New(errKeyExists)
		}
		p = p.next
	}

	h.table[hash] = &intNode{
		key:   key,
		value: value,
		next:  h.table[hash],
	}

	return err
}

func (h chainedHash) remove(key interface{}) error {
	var err error

	hash, err := h.hashValue(key)
	if err != nil {
		return err
	}
	p := h.table[hash]
	var prevp *intNode

	for p != nil {
		// 削除成功時
		// 削除したいキーとノードのキーが一致したら
		if p.key == key {
			// 前回のノードがnilなら今見ているノードが先頭なのでnextをh.table[hash]に入れる
			if prevp == nil {
				h.table[hash] = p.next
				// 前回のノードがあった場合。今見ているノードの次のノードを前回ノードのnextにいれる（今見ているノードを飛ばす=削除）
			} else {
				prevp.next = p.next
			}
			return err
		}
		prevp = p
		p = p.next
	}

	return errors.New(errKeyNotExist)
}

func (h chainedHash) dump() {
	for i := 0; i < h.capacity; i++ {
		p := h.table[i]
		fmt.Print(i)

		for p != nil {
			fmt.Printf("  -> %v (%v)", p.key, p.value)
			p = p.next
		}
		fmt.Println("")
	}
}
