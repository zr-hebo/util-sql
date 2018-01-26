package sqlutil

// "fmt"

var (
	// commentArray []*CommentRule
	commentArray = []*CommentRule{
		&CommentRule{[]rune("#"), []rune("\n"), []rune{}, []rune("\n")},
		&CommentRule{[]rune("--"), []rune("\n"), []rune{}, []rune("\n")},
		&CommentRule{[]rune("/*"), []rune("*/"), []rune("/*!"), []rune(" ")},
	}
)

type commentMatcher struct {
	CurrentComment *CommentRule
	StartPos       int
	EndPos         int
}

func isCharInCommentStart(currChar rune) bool {
	for _, currComment := range commentArray {
		if currChar == currComment.Start[0] {
			return true
		}
	}

	return false
}

func isCharInCommentEnd(currChar rune) bool {
	for _, currComment := range commentArray {
		if len(currComment.End) > 0 && currChar == currComment.End[0] {
			return true
		}
	}

	return false
}

func (cm *commentMatcher) matchMeaningfullCommentStart(
	runeArray []rune, idx int) (bool, *CommentRule) {
	if cm.StartPos != -1 {
		return false, nil
	}

	for _, currComment := range commentArray {
		if matchSubArray(runeArray, currComment.Start, idx) &&
			!matchSubArray(runeArray, currComment.Exception, idx) {
			return true, currComment
		}
	}

	return false, nil
}

func (cm *commentMatcher) matchCommentEnd(
	runeArray []rune, idx int) (bool, *CommentRule) {
	if cm.StartPos == -1 {
		return false, nil
	}

	for _, currComment := range commentArray {
		if matchSubArray(runeArray, cm.CurrentComment.End, idx) {
			return true, currComment
		}
	}

	return false, nil
}
