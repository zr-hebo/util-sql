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
func Split(rawStr string) (sts []string, err error) {
	sts = make([]string, 0, defaultSliceLen)
	rawRunes := []rune(rawStr)
	rawRunes = append(rawRunes, rune('\n'))
	sm := NewSplitChecker(rawStr)
	var stateCode int
	subStrStartPos := 0
	allRuneLen := len(rawRunes)
	var option interface{}

	for idx := 0; idx < allRuneLen; idx++ {
		cr := rawRunes[idx]

		currState := sm.GetState()
		stateCode, err = getStateCode(currState)
		if err != nil {
			return
		}

		if cr == splitRune && stateCode != sensiableStateCode {
			subRunes := rawRunes[subStrStartPos : idx+1]
			subStr := fmt.Sprintf("%s", string(subRunes))
			sts = append(sts, subStr)
			subStrStartPos = idx + 1
			continue
		}

		err = sm.Walk(idx)
		if err != nil {
			return
		}

		option = sm.GetState().GetVal()
		if option != nil {
			step := option.(int)
			idx = idx + step - 1
		}
	}
	sm.Walk(allRuneLen)

	if !sm.IsFinished() {
		err = fmt.Errorf(
			"invalid begin from %dth charater in sql: %s",
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
			sts = append(sts)
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
