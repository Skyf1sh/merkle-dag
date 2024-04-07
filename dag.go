package merkledag

import (
	"hash"
)

<<<<<<< HEAD
func Add(store KVStore, node Node, h hash.Hash) []byte {
	switch node.Type() {
	case FILE:
		fileNode := node.(File)
		data := fileNode.Bytes()
		hash := HashData(data, h)
		if err := store.Put(hash, data); err != nil {
			panic(err) // Handle error appropriately
		}
		return hash
	case DIR:
		dirNode := node.(Dir)
		it := dirNode.It()
		var childrenHashes [][]byte
		for it.Next() {
			childNode := it.Node()
			childHash := Add(store, childNode, h)
			childrenHashes = append(childrenHashes, childHash)
		}
		hash := HashChildren(childrenHashes, h)
		return hash
	default:
		panic("Unknown node type")
	}
}

func HashData(data []byte, h hash.Hash) []byte {
	h.Reset()
	h.Write(data)
	return h.Sum(nil)
}

func HashChildren(children [][]byte, h hash.Hash) []byte {
	h.Reset()
	for _, child := range children {
		h.Write(child)
	}
	return h.Sum(nil)
=======
type Link struct {
	Name string
	Hash []byte
	Size int
}

type Object struct {
	Links []Link
	Data  []byte
}

func Add(store KVStore, node Node, hp HashPool) []byte {
	// TODO 将分片写入到KVStore中，并返回Merkle Root
	return nil
>>>>>>> c1a60202beb4fd3abcbf7a2bc99c8ff22ea55320
}
