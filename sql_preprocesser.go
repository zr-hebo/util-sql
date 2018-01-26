package sqlutil

import (
	"errors"
	"fmt"
	"strings"
)

// CommentRule 注释规则
type CommentRule struct {
	Start       []rune
	End         []rune
	Exception   []rune
	Replacement []rune
}

// RemoveComments 删除sql语句中的注释信息
func RemoveComments(stmtStr string) (string, error) {
	stmtRuneArray := []rune(stmtStr)
	cleanStmtRuneArray, err := removeCommentAvoidQuote(stmtRuneArray)
	if err != nil {
		return "", err
	}

	//cleanStmtRuneArray = removeSpecialCommentAvoidQuote(cleanStmtRuneArray,
	//	[]rune("--\n"))

	return string(cleanStmtRuneArray), nil
}

func removeSpecialCommentAvoidQuote(rawStmtRuneArray,
	commentSpecialRuneArray []rune) []rune {
	// remove comment like "--\n" or "--"
	stmtRuneArray := append(rawStmtRuneArray, '\n')
	// commentSpecialRuneArray := []rune(specialCommentStr)
	cleanStmtRuneArray := []rune{}
	specialCommentLength := len(commentSpecialRuneArray)
	stmtLength := len(stmtRuneArray)
	if specialCommentLength < 1 || stmtLength < 1 {
		return cleanStmtRuneArray
	}

	firstRuneOfSpecial := commentSpecialRuneArray[0]
	qm := quoteMatcher{}
	cursor := 0
	commentStartPos := -1
	commentEndPos := -1
	for idx := 0; idx < stmtLength; idx++ {
		currChar := stmtRuneArray[idx]

		if isMeaningfulQuoteInIndex(stmtRuneArray, idx) {
			qm.pushQuote(currChar)

		} else if currChar == firstRuneOfSpecial {
			if qm.isQuoteMatch() &&
				matchSubArray(stmtRuneArray, commentSpecialRuneArray, idx) {
				commentStartPos = idx
				commentEndPos = idx + specialCommentLength - 1
				if cursor < commentStartPos {
					cleanStmtRuneArray = append(append(cleanStmtRuneArray, '\n'),
						stmtRuneArray[cursor:commentStartPos]...)
				}

				cursor = commentEndPos + 1
				idx = commentEndPos
				commentStartPos = -1
				commentEndPos = -1
			}
		}
	}

	if cursor < stmtLength {
		if commentStartPos < 0 && qm.isQuoteMatch() {
			cleanStmtRuneArray = append(cleanStmtRuneArray,
				stmtRuneArray[cursor:]...)
		} else if cursor < commentStartPos {
			cleanStmtRuneArray = append(cleanStmtRuneArray,
				stmtRuneArray[cursor:commentStartPos]...)
		}
	}

	cleanStmtLength := len(cleanStmtRuneArray)
	if cleanStmtLength > 0 &&
		cleanStmtRuneArray[cleanStmtLength-1] == '\n' {
		cleanStmtRuneArray = cleanStmtRuneArray[:cleanStmtLength-1]
	}

	return cleanStmtRuneArray
}

func splitAvoidQuote(stmtStr string, splitor rune) ([]string, error) {
	runeArray := []rune(stmtStr)
	runeArrayLen := len(runeArray)
	if runeArray[runeArrayLen-1] != splitor {
		runeArray = append(runeArray, splitor)
	}

	subStrBeginPos := 0
	statementArray := []string{}
	qm := quoteMatcher{}

	for pos, currentChar := range runeArray {
		if currentChar != splitor {
			if isMeaningfulQuoteInIndex(runeArray, pos) {
				qm.pushQuote(currentChar)
			}
		} else {

			if qm.isQuoteMatch() {
				subStr := string(runeArray[subStrBeginPos:pos])
				if len(strings.TrimSpace(subStr)) > 0 {
					statementArray = append(statementArray, subStr)
				}
				subStrBeginPos = pos + 1
			}
		}
	}

	if subStrBeginPos < len(runeArray) {
		return nil, errors.New("invalid SQL statement in statements")
	}
	return statementArray, nil
}

func isMeaningfulQuoteInIndex(runeArray []rune, idx int) bool {
	currentChar := runeArray[idx]
	return isCharInArray(currentChar, quoteArray) &&
		(!haveUnmatchedEscapeBeforeIndex(runeArray, idx))
}

func removeCommentAvoidQuote(rawStmtRuneArray []rune) ([]rune, error) {
	stmtRuneArray := append(rawStmtRuneArray, '\n')
	cleanStmtRuneArray := []rune{}
	stmtLength := len(stmtRuneArray)

	qm := quoteMatcher{}
	cm := commentMatcher{nil, -1, -1}
	cursor := 0
	for idx := 0; idx < stmtLength; idx++ {
		currChar := stmtRuneArray[idx]

		// fmt.Printf("currChar:%s", string(currChar))
		if isMeaningfulQuoteInIndex(stmtRuneArray, idx) {
			qm.pushQuote(currChar)

		} else if isCharInCommentStart(currChar) {
			isCommentStartMatch, matchedComment :=
				cm.matchMeaningfullCommentStart(stmtRuneArray, idx)
			if qm.isQuoteMatch() && isCommentStartMatch {
				cm.StartPos = idx
				cm.CurrentComment = matchedComment
				idx = idx + len(matchedComment.Start) - 1
			}

		} else if isCharInCommentEnd(currChar) {
			isCommentEndMatch, _ :=
				cm.matchCommentEnd(stmtRuneArray, idx)
			if qm.isQuoteMatch() && isCommentEndMatch {

				cm.EndPos = idx + len(
					cm.CurrentComment.End) - 1
				if cursor < cm.EndPos {
					cleanStmtRuneArray = append(cleanStmtRuneArray,
						stmtRuneArray[cursor:cm.StartPos]...)
					cleanStmtRuneArray = append(cleanStmtRuneArray,
						cm.CurrentComment.Replacement...)
				}

				cursor = cm.EndPos + 1
				idx = cm.EndPos
				cm.StartPos = -1
				cm.EndPos = -1
			}
		}
	}

	if cursor < stmtLength {
		if qm.isQuoteMatch() &&
			cm.StartPos == -1 {
			cleanStmtRuneArray = append(cleanStmtRuneArray,
				stmtRuneArray[cursor:]...)
		} else {
			return nil, fmt.Errorf("have invalid comments in raw statements")
		}
	}

	cleanStmtLength := len(cleanStmtRuneArray)
	if cleanStmtLength > 0 && cleanStmtRuneArray[cleanStmtLength-1] == '\n' {
		cleanStmtRuneArray = cleanStmtRuneArray[:cleanStmtLength-1]
	}

	return cleanStmtRuneArray, nil
}

func matchSubArray(mainArray, subArray []rune, idx int) bool {
	subArrayLength := len(subArray)
	mainArrayLength := len(mainArray)
	if subArrayLength < 1 || mainArrayLength < 1 {
		return false
	}

	if idx+subArrayLength <= mainArrayLength {
		return isSameArray(mainArray[idx:idx+subArrayLength], subArray)
	}

	return false
}

func isSameArray(srcArray, destArray []rune) bool {
	if len(srcArray) != len(destArray) {
		return false
	}

	for idx := 0; idx < len(srcArray); idx++ {
		if srcArray[idx] != destArray[idx] {
			return false
		}

	}
	return true
}

func haveUnmatchedEscapeBeforeIndex(runeArray []rune, beginIdx int) bool {
	haveUnmatchedEscape := false
	lookBackIndex := beginIdx - 1
	for lookBackIndex >= 0 {
		preChar := runeArray[lookBackIndex]
		if preChar == '\\' {
			haveUnmatchedEscape = !haveUnmatchedEscape
			lookBackIndex = lookBackIndex - 1
		} else {
			break
		}
	}

	return haveUnmatchedEscape
}
