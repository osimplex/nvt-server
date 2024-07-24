package vt100

import (
	"errors"
)

const (
	SymUndefined SymbolType = iota
	SymDigit
	SymCommand
)

var (
	ErrNoChildren    = errors.New("no children in this node")
	ErrChildNotFound = errors.New("requested child for key node not found")
)

type SymbolType int

type SymbolNode struct {
	Symbol   byte
	Type     SymbolType
	Command  string
	Children map[byte]*SymbolNode
}

func (node *SymbolNode) FindChild(b byte) (*SymbolNode, error) {
	if node.Children == nil {
		return &SymbolNode{}, ErrNoChildren
	}

	if child := node.Children[b]; child != nil {
		return child, nil
	}

	return &SymbolNode{}, ErrChildNotFound
}
