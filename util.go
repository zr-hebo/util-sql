package sqlutil

import (
	"fmt"

	"github.com/zr-hebo/state_machine/gfsm"
)

func getStateCode(s gfsm.Stater) (stateCode int, err error) {
	ps, ok := s.(*PlainState)
	if ok {
		stateCode = ps.stateCode
		return
	}

	ss, ok := s.(*SensiableState)
	if ok {
		stateCode = ss.stateCode
		return
	}

	es, ok := s.(*EndState)
	if ok {
		stateCode = es.stateCode
		return
	}

	err = fmt.Errorf("invalid state: %s", s)
	return
}
