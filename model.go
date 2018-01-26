package sqlutil

import (
	"errors"

	"github.com/zr-hebo/state_machine/gfsm"
)

type sensiableRange struct {
	start      []rune
	end        []rune
	careEscape bool
}

const (
	// PlainState 语句可拆分状态
	plainStateCode int = iota
	// SensiableState 敏感状态，不能进行拆分，例如在引号或者注释内
	sensiableStateCode
	// EndState 结束状态
	endStateCode
)

const (
	escapseRune = rune('\\')
)

var stateNames = map[int]string{
	plainStateCode:     "plain_state",
	sensiableStateCode: "sensiable_state",
	endStateCode:       "end_state",
}
var (
	sensiableRanges = []sensiableRange{
		{start: []rune("'"), end: []rune("'"), careEscape: true},
		{start: []rune("\""), end: []rune("\""), careEscape: true},
		{start: []rune("`"), end: []rune("`"), careEscape: true},
		{start: []rune("-- "), end: []rune("\n"), careEscape: false},
		{start: []rune("#"), end: []rune("\n"), careEscape: false},
		{start: []rune("/*"), end: []rune("*/"), careEscape: false},
	}
)

type commonComponents struct {
	stateMachine *gfsm.StateMachine
	checkRunes   []rune
}

// matchState 引号或者注释匹配状态
type matchState struct {
	*commonComponents
	stateCode int
}

func (ms matchState) String() string {
	return stateNames[ms.stateCode]
}

func (ms *matchState) IsEnd() bool {
	return false
}

func (ms *matchState) GetVal() interface{} {
	return ms.stateCode
}

func (ms *matchState) Walk(val interface{}) (
	nextState gfsm.Stater, option interface{}, err error) {
	return
}

// SetMachine
func (ms matchState) SetMachine(sm *gfsm.StateMachine) {
	ms.stateMachine = sm
}

// PlainState PlainState
type PlainState struct {
	matchState
}

// NewStartState NewStartState
func NewStartState(checkRunes []rune) (ps *PlainState) {
	ps = &PlainState{}
	cc := &commonComponents{checkRunes: checkRunes}
	ps.commonComponents = cc
	ps.stateCode = plainStateCode
	return
}

// NewPlainState NewPlainState
func NewPlainState(cc *commonComponents) (ps *PlainState) {
	ps = &PlainState{}
	ps.stateCode = plainStateCode
	ps.commonComponents = cc
	return
}

// Walk Walk
func (ps *PlainState) Walk(val interface{}) (
	nextState gfsm.Stater, option interface{}, err error) {
	checkRuneLen := len(ps.checkRunes)
	startPos := val.(int)
	nextState = ps

	if startPos >= checkRuneLen {
		nextState = NewEndState(ps.commonComponents)
		return
	}

	option = 1
	for _, sr := range sensiableRanges {
		cmpLen := len(sr.start)
		endPos := startPos + cmpLen
		if endPos > checkRuneLen {
			endPos = checkRuneLen
		}

		if isRunesSame(ps.checkRunes[startPos:endPos], sr.start) {
			option = len(sr.start)
			nextState = NewSensiableState(ps.commonComponents)
			return
		}
	}

	return
}

// SensiableState SensiableState
type SensiableState struct {
	matchState
	sr *sensiableRange
}

// NewSensiableState NewPlainState
func NewSensiableState(cc *commonComponents) (ss *SensiableState) {
	ss = &SensiableState{}
	ss.stateCode = sensiableStateCode
	ss.commonComponents = cc

	return
}

// Walk Walk
func (ss *SensiableState) Walk(val interface{}) (
	nextState gfsm.Stater, option interface{}, err error) {
	checkRuneLen := len(ss.checkRunes)
	startPos := val.(int)
	nextState = ss

	option = 1
	for _, sr := range sensiableRanges {
		cmpLen := len(sr.end)
		endPos := startPos + cmpLen
		if endPos > checkRuneLen {
			endPos = checkRuneLen
		}

		if isRunesSame(ss.checkRunes[startPos:endPos], sr.end) &&
			checkEscape(sr.careEscape, ss.checkRunes, startPos) {
			option = len(sr.end)
			nextState = NewPlainState(ss.commonComponents)
			return
		}
	}

	return
}

func checkEscape(careEscape bool, checkRunes []rune, pos int) (ok bool) {
	if !careEscape {
		ok = true
		return
	}

	counter := 0
	for pos--; pos >= 0; pos-- {
		if checkRunes[pos] == escapseRune {
			counter++

		} else {
			break
		}
	}

	ok = counter%2 == 0
	return
}

// EndState EndState
type EndState struct {
	matchState
}

// NewEndState NewEndState
func NewEndState(cc *commonComponents) (es *EndState) {
	es = &EndState{}
	es.commonComponents = cc
	es.stateCode = endStateCode
	return
}

// Walk Walk
func (es *EndState) Walk(val interface{}) (
	nextState gfsm.Stater, option interface{}, err error) {
	err = errors.New(
		"state machine was already in finish state, cannot walk")
	return
}

func isRunesSame(src, dest []rune) (same bool) {
	srcLen := len(src)
	destLen := len(dest)
	if srcLen != destLen {
		same = false
		return
	}

	for idx, ch := range src {
		if ch != dest[idx] {
			same = false
			return
		}
	}

	same = true
	return
}

// IsEnd IsEnd
func (es *EndState) IsEnd() bool {
	return true
}
