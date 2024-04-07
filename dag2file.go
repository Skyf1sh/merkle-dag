package merkledag

func Hash2File(store KVStore, hash []byte, path string, hp HashPool) []byte {

	data, err := store.Get(hash)
	if err != nil {
		return nil
	}

	switch path {
	case "tree":

		return data
	case "hash":

		return hash
	default:

		return nil
	}
}
