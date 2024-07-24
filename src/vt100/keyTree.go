package vt100

func GetCommandTree() *SymbolNode {
	escKeyTree := map[byte]*SymbolNode{
		KeyLeftBracket: {
			Symbol: KeyLeftBracket,
			Children: map[byte]*SymbolNode{
				KeyACapLetter: {Symbol: KeyACapLetter, Type: SymCommand, Command: "KeyArrowUp"},
				KeyBCapLetter: {Symbol: KeyBCapLetter, Type: SymCommand, Command: "KeyArrowDown"},
				KeyCCapLetter: {Symbol: KeyCCapLetter, Type: SymCommand, Command: "KeyArrowRight"},
				KeyDCapLetter: {Symbol: KeyDCapLetter, Type: SymCommand, Command: "KeyArrowLeft"},
			},
		},

		KeyOCapLetter: {
			Symbol: KeyOCapLetter,
			Children: map[byte]*SymbolNode{
				KeyPCapLetter: {Symbol: KeyPCapLetter, Type: SymCommand, Command: "KeyF1"},
				KeyQCapLetter: {Symbol: KeyQCapLetter, Type: SymCommand, Command: "KeyF2"},
				KeyRCapLetter: {Symbol: KeyRCapLetter, Type: SymCommand, Command: "KeyF3"},
				KeySCapLetter: {Symbol: KeySCapLetter, Type: SymCommand, Command: "KeyF4"},
			},
		},
	}

	tree := &SymbolNode{
		Children: map[byte]*SymbolNode{
			KeyEscape:      {Symbol: KeyEscape, Children: escKeyTree},
			KeyCarriageRet: {Symbol: KeyCarriageRet, Type: SymCommand, Command: "KeyEnter"},

			KeyNumber0: {Symbol: KeyNumber0, Type: SymDigit},
			KeyNumber1: {Symbol: KeyNumber1, Type: SymDigit},
			KeyNumber2: {Symbol: KeyNumber2, Type: SymDigit},
			KeyNumber3: {Symbol: KeyNumber3, Type: SymDigit},
			KeyNumber4: {Symbol: KeyNumber4, Type: SymDigit},
			KeyNumber5: {Symbol: KeyNumber5, Type: SymDigit},
			KeyNumber6: {Symbol: KeyNumber6, Type: SymDigit},
			KeyNumber7: {Symbol: KeyNumber7, Type: SymDigit},
			KeyNumber8: {Symbol: KeyNumber8, Type: SymDigit},
			KeyNumber9: {Symbol: KeyNumber9, Type: SymDigit},
		},
	}

	return tree
}
