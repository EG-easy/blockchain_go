package main

import "bytes"

type TXInput struct {
	Txid     []byte
	Vout      int
	Signature[]byte
	PubKey   []byte
}

func (in *TXInput) UsesKey(publicKeyHash []byte) bool {
	lockingHash := HashPubKey(in.PubKey)
	return bytes.Compare(lockingHash, publicKeyHash) == 0
}
