package sqlutil

var (
	quoteArray = []rune{
		'"', '\'', '`',
	}
	parenthesisArray = []rune{
		'(', ')', '[', ']', '{', '}',
	}
	heheArray = []rune{
		'(', ')', '[', ']', '{', '}',
	}

	parenthesisMap = map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}
)

type quote struct {
	Front rune
	Back  rune
}

type quoteMatcher struct {
	quoteStack rune
}

type parenthesisMatcher struct {
	parenthesisStack []rune
}

type plainChecker struct {
	quoteMatcher       quoteMatcher
	parenthesisMatcher parenthesisMatcher
}

//IsQuote 字符是否是引号
func IsQuote(quote rune) bool {
	return isCharInArray(quote, quoteArray)
}

func (qm *quoteMatcher) pushQuote(quote rune) {
	if !IsQuote(quote) {
		return
	}

	if qm.quoteStack == 0 {
		qm.quoteStack = quote
	} else if quote == qm.quoteStack {
		qm.quoteStack = 0
	}
}

func isCharInArray(char rune, charArray []rune) bool {
	for _, tempChar := range charArray {
		if char == tempChar {
			return true
		}
	}

	return false
}

func (qm *quoteMatcher) isQuoteMatch() bool {
	if qm.quoteStack == 0 {
		return true
	}

	return false
}

func (pc *plainChecker) isPlain() bool {
	if pc.parenthesisMatcher.isParenthesisMatch() &&
		pc.quoteMatcher.isQuoteMatch() {
		return true
	}

	return false
}

func (pc *plainChecker) pushChar(runeArray []rune, idx int) {
	if idx >= len(runeArray) {
		return
	}

	char := runeArray[idx]
	if isCharInArray(char, quoteArray) {
		if !haveUnmatchedEscapeBeforeIndex(runeArray, idx) {
			pc.quoteMatcher.pushQuote(char)
			return
		}
	} else if isCharInArray(char, parenthesisArray) {
		if pc.quoteMatcher.isQuoteMatch() {
			pc.parenthesisMatcher.pushParenthesis(char)
		}
	}
}

func (pm *parenthesisMatcher) pushParenthesis(parenthesis rune) {
	matchedParenthesis, found := parenthesisMap[parenthesis]
	if found {
		stackLen := len(pm.parenthesisStack)
		if stackLen > 0 && matchedParenthesis == pm.parenthesisStack[stackLen-1] {
			pm.parenthesisStack = pm.parenthesisStack[:stackLen-1]
			return
		}
	}
	pm.parenthesisStack = append(pm.parenthesisStack, parenthesis)
}

func (pm *parenthesisMatcher) isParenthesisMatch() bool {
	if len(pm.parenthesisStack) == 0 {
		return true
	}

	return false
}
