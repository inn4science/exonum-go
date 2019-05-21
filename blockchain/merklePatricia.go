package blockchain

//exonumCrypto "github.com/inn4science/exonum-go/crypto"

type Proof struct {
	Path *ProofPath
	Hash []byte
}

type MapProof struct {
	entries []string
	proof   []Proof
}

type ProofJson struct {
	Path string `json:"path"`
	Hash string `json:"hash"`
}

type Json struct {
	Entries []interface{} `json:"entries"`
	Proof   []ProofJson   `json:"proof"`
}

type Entries struct {
	Key   string
	Value string
	Path  string
	Hash  []byte
}

//func (MapProof) New(json Json) (MapProof, error) {
//	var proofPaths []Proof
//	for _, value := range json.Proof {
//		path, err := ProofPath{}.New(value.Path, BitLength)
//		if err != nil {
//			return MapProof{}, err
//		}
//		hash, err := exonumCrypto.Hash{}.FromString(value.Hash)
//		if err != nil {
//			return MapProof{}, err
//		}
//		proofPaths = append(proofPaths, Proof{
//			&path,
//			hash.GetData(),
//		})
//	}
//
//	var entries Entries
//	for _, value := range json.Entries {
//		if value.missing != nil {
//			return MapProof{}, nil
//		}
//
//		entries = Entries{
//			Key:   value.key,
//			Value: value.value
//			Path:  createPath(value.key),
//			Hash:  hash,
//		}
//	}
//
//	fmt.Println(entries)
//	fmt.Println(proofPaths)
//
//	err := precheckProof(proofPaths)
//	if err != nil {
//		fmt.Println("err", err)
//	}
//
//	completeProof := append(proofPaths, entries)
//
//	for i := 1; i < len(completeProof); i++ {
//		pathA := completeProof[i-1]
//		pathB := completeProof[i]
//
//		if pathA.Path.Compare(pathB.Path) == 0 {
//			return MapProof{}, errors.New("duplicatePath")
//		}
//	}
//
//	merkleRoot := collect(completeProof)
//
//	missingKeys := entries
//	entries = entries
//
//	return MapProof{
//		proofPath,
//	}
//}
//
//func precheckProof(proofs []Proof) error {
//	for i := 1; i < len(proofs); i++ {
//		prevPath := proofs[i-1].Path
//		path := proofs[i].Path
//
//		switch prevPath.Compare(path) {
//		case -1:
//			return errors.New("error")
//		case 0:
//			return errors.New("error")
//		case 1:
//			return errors.New("error")
//		}
//	}
//}
//
//func collect(entries []Entries) (string, error) {
//	switch len(entries) {
//	case 0:
//		return "0000000000000000000000000000000000000000000000000000000000000000", nil
//	case 1:
//		if !entries[0].Path.isTerminal {
//			return "", errors.New("nonTerminalNode")
//		}
//		return hashIsolatedNode(entries[0]), nil
//	default:
//		var contour []string
//
//		lastPrefix := entries[0].Path.CommonPrefix(entries[1].Path)
//		contour = append(entries[0], entries[1])
//
//		for i := 2; i < len(entries); i++ {
//			entry := entries[i]
//			newPrefix := entry.Path.CommonPrefix(contour[len(contour)-1].Path)
//
//			for len(contour) > 1 && newPrefix.bitLength() < lastPrefix.bitLength() {
//				foldedPrefix := fold(contour, lastPrefix)
//				if foldedPrefix != nil {
//					lastPrefix = foldedPrefix
//				}
//			}
//
//			contour = append(contour, entry)
//			lastPrefix = newPrefix
//		}
//
//		for len(contour) > 1 {
//			lastPrefix = fold(contour, lastPrefix)
//		}
//		return contour[0].Hash
//
//	}
//}
//
//func hashIsolatedNode(path, valueHash []byte) exonumCrypto.Hash {
//	buffer := serializeIsolatedNode(path, valueHash)
//	return exonumCrypto.Hash{}.FromData(buffer)
//}
//
//func serializeIsolatedNode(path, hash []byte) []byte {
//	buf := make([]byte, 0)
//	serializeProofPathType(path, buf)
//	buf = append(buf, hash...)
//	return buf
//}
//
//func fold(contour, lastPrefix string) {
//	var lastEntry string
//	var penultimateEntry string
//	contour, lastEntry = contour[len(contour)-1], contour[:len(contour)-1]
//	contour, penultimateEntry = contour[len(contour)-1], contour[:len(contour)-1]
//
//	contour = append(contour, Proof{
//		Path: lastPrefix,
//		Hash: hashBranch(penultimateEntry, lastEntry),
//	})
//
//	if len(contour) > 1 {
//		return lastPrefix.CommonPrefix(contour[contour.length-2].Path)
//	} else {
//		return nil
//	}
//}
//
//func hashBranch(left, right string) exonumCrypto.Hash {
//	buffer := serializeBranchNode(left.Hash, right.Hash, left.Path, right.Path)
//	return exonumCrypto.Hash{}.FromData(buffer)
//}
//
//func serializeBranchNode(leftHash, rightHash, leftPath, rightPath string) []byte {
//	var buffer []byte
//	buffer = append(buffer, leftHash...)
//	buffer = append(buffer, rightHash...)
//	buffer = append(buffer, serializeProofPathType(leftPath, buffer)...)
//	buffer = append(buffer, serializeProofPathType(rightPath, buffer)...)
//	return buffer
//}
//
//func serializeProofPathType(typeHash ProofPath, buffer []byte) []byte {
//	if typeHash.IsTerminal {
//		buffer = append(buffer, byte(1))
//	} else {
//		buffer = append(buffer, byte(0))
//	}
//	buffer = append(buffer, typeHash.Key...)
//	buffer = append(buffer, byte(typeHash.LengthByte))
//	return buffer
//}
