package sqlutil

import (
	"fmt"
	"strings"

	"github.com/zr-hebo/state_machine/gfsm"
)

var splitRune = rune(';')

const (
	defaultSliceLen = 64
)

// Split 分割字符串，保留引号内部的分割符，用于sql语句拆分之类的语句.
func Split(rawStr string) (splitSQLs []string, err error) {
	splitSQLs = make([]string, 0, defaultSliceLen)
	rawRunes := []rune(rawStr)
	rawRunes = append(rawRunes, rune('\n'))
	splitStateMachine := NewSplitChecker(rawStr)
	var stateCode int
	subStrStartPos := 0
	allRuneLen := len(rawRunes)
	// var option interface{}

	for idx := 0; idx < allRuneLen; idx++ {
		cr := rawRunes[idx]

		a := string(rawRunes[idx])
		_ = a

		currState := splitStateMachine.GetState()
		stateCode, err = getStateCode(currState)
		if err != nil {
			return
		}

		// 遇到分割符并且可以分割
		if cr == splitRune && stateCode != sensiableStateCode {
			subRunes := rawRunes[subStrStartPos : idx+1]
			splitSQLs = append(splitSQLs, string(subRunes))
			subStrStartPos = idx + 1
			continue
		}

		// 处理当前字符
		err = splitStateMachine.Walk(idx)
		if err != nil {
			return
		}
		/*
		option = splitStateMachine.GetState().GetVal()
		if option != nil {
			step := option.(int)
			idx = idx + step - 1
		}
		*/
	}
	/*
		err = splitStateMachine.Walk(allRuneLen)
		if err != nil {
			return
		}
	*/

	if !splitStateMachine.IsFinished() {
		err = fmt.Errorf(
			"invalid SQL begin from %dth charater : %s",
			subStrStartPos, rawStr[subStrStartPos:])
		return
	}

	if subStrStartPos < allRuneLen {
		subRunes := rawRunes[subStrStartPos:allRuneLen]
		if len(subRunes) <= 1 {
			return
		}

		subRunes = subRunes[:len(subRunes)-1]
		subStr := strings.TrimSpace(string(subRunes))
		if len(subStr) > 0 {
			splitSQLs = append(splitSQLs)
		}
	}

	return
}

// NewSplitChecker SplitNew
func NewSplitChecker(sql string) (sm *gfsm.StateMachine) {
	startState := NewStartState([]rune(sql))
	sm = gfsm.NewStateMachine(startState)
	return
}
