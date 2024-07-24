package vt100

import (
	"io"
)

var (
	keyTreeRoot = GetCommandTree()
)

func KeyEval(r io.Reader) (SymbolNode, error) {
	var (
		inputBuffer [1]byte
	)

	keyCurrentNode := keyTreeRoot

	for {
		_, err := r.Read(inputBuffer[:])
		if err != nil {
			return SymbolNode{}, err
		}

		keyCurrentNode, err = keyCurrentNode.FindChild(inputBuffer[0])
		if err != nil {
			return SymbolNode{}, err
		}

		if keyCurrentNode.Type != SymUndefined {
			return *keyCurrentNode, nil
		}
	}
}
