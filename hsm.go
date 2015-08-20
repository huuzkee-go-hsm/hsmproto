package hsmproto

import (
	"fmt"
)

/////////////////////////////////////////////////////////////////////////////////
// SYSTEM STATE TOKENS //////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////////////////////////////

const (
	HSM_SYSSTAT_NULL           = (0 + iota) // NULL PSEUDO STATE
	HSM_SYSSTAT_ACTIVATE                    // NULL PSEUDO STATE
	HSM_SYSSTAT_LIVE                        // NULL PSEUDO STATE
	HSM_SYSSTAT_DEBUG                       // NULL PSEUDO STATE
	HSM_SYSSTAT_PREHIBERNATE                // NULL PSEUDO STATE
	HSM_SYSSTAT_HIBERNATE                   // NULL PSEUDO STATE
	HSM_SYSSTAT_POSTHIBERNATE               // NULL PSEUDO STATE
	HSM_SYSSTAT_DEACTIVATE                  // NULL PSEUDO STATE
	HSM_SYSSTAT_INITHSM                     // NULL PSEUDO STATE
	HSM_SYSSTAT_ENTERHIERARCHY              // NULL PSEUDO STATE
	HSM_SYSSTAT_EXITHIERARCHY               // NULL PSEUDO STATE
	HSM_SYSSTAT_BLOCKING                    // NULL PSEUDO STATE
	HSM_SYSSTAT_DEMATERIALISE               // NULL PSEUDO STATE
	HSM_SYSSTAT_REMATERIALISE               // NULL PSEUDO STATE
	HSM_SYSSTAT_REGENERATING                // NULL PSEUDO STATE
	HSM_SYSSTAT_DEBUGSTEP                   // NULL PSEUDO STATE
	HSM_SYSSTAT_DEBUGTRACE                  // NULL PSEUDO STATE
	HSM_SYSSTAT_FINALISED                   // NULL PSEUDO STATE
	HSM_SYSSTAT_DEAD                        // NULL PSEUDO STATE
	HSM_SYSSTAT_USER                        // NULL PSEUDO STATE
)

/////////////////////////////////////////////////////////////////////////////////
// INTERFACES ///////////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////////////////////////////

type HsmState interface {
	acceptState(layer *HsmActorLayer) int
	getState() int
	describeState() string
}

//--- HsmActorLayerStateFrame -----------------------------------------------------------//

type HsmActorLayerStateFrame interface {
	getLastState() int
	setLastState(state int)
	getCurrentState() int
	setCurrentState(state int)
	getStates() []HsmState

	initFrame(minState int, maxState int) error
}

//--- HsmActor ----------------------------------------------------------------//

type HsmActor interface {
	//GoLive() (exitstate int, err error)
	getUserLayerSegment(seg int) HsmActorUserLayer
}

//--- HsmActorLayer-----------------------------------------------------------//

type HsmActorLayer interface {
	HsmActorLayerStateFrame
	Live(layer *HsmActorLayer) (exitstate int, err error)
	getActor() *HsmActor
	//setActor(actor *HsmActor)
}

/////////////////////////////////////////////////////////////////////////////////
// TYPES ////////////////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////////////////////////////

//--- HsmActorBaseLayer ------------------------------------------------------------//

type HsmActorBaseLayer struct {
	LastState    int
	CurrentState int
	States       []HsmState
	actor        *HsmActor
}

func (hsa HsmActorBaseLayer) getLastState() int { return hsa.LastState }

func (hsa HsmActorBaseLayer) setLastState(state int) { hsa.LastState = state }

func (hsa HsmActorBaseLayer) getCurrentState() int { return hsa.CurrentState }

func (hsa HsmActorBaseLayer) setCurrentState(state int) { hsa.CurrentState = state }

func (hsa HsmActorBaseLayer) getActor() *HsmActor {
	//fmt.Printf("\tx2:  \r\n")

	//if nil == &hsa { fmt.Printf("\tx2!:  \r\n") }

	//act := hsa.actor

	//if nil == act { fmt.Printf("\tx2!act:  \r\n") }

	return hsa.actor
}

func (hsa *HsmActorBaseLayer) setActor(pactor *HsmActor) {

	//fmt.Printf("\tx1:  \r\n")
	//if nil == *pactor { fmt.Printf("\tx1!:  \r\n") }

	hsa.actor = pactor

	act := hsa.actor

	if nil == act {
		fmt.Printf("\tx1!!!!:  \r\n")
	}

}

func (hsa HsmActorBaseLayer) initFrame(minState int, MaxState int) error { return nil }

func (hsa HsmActorBaseLayer) getStates() []HsmState { return hsa.States }

func (hsa HsmActorBaseLayer) Live(layer *HsmActorLayer) (exitstate int, err error) {

	m1 := *layer // actorp := m1.getActor() ; actor := *actorp ; //ul := actor.getUserLayerSegment(HSM_SYSSTAT_LIVE) ;
	st := m1.getStates()
	cs := st[m1.getCurrentState()]

	if m1.getCurrentState() == HSM_SYSSTAT_HIBERNATE {
		m1.setCurrentState(HSM_SYSSTAT_POSTHIBERNATE)
	}

	hsal := HsmActorLayer(hsa)

	hsal.getActor()

	for m1.getCurrentState() != HSM_SYSSTAT_DEAD {

		fmt.Printf("\tThe Current State of The World is:  # %v  %v \r\n", m1.getCurrentState(), cs.describeState())

		m1.setCurrentState(cs.acceptState(&hsal))

	}

	return hsa.CurrentState, nil

}

func Create_HsmActorBaseLayer() HsmActorBaseLayer {

	var states []HsmState = make([]HsmState, HSM_SYSSTAT_DEAD+1)

	states[HSM_SYSSTAT_NULL] = hsm_systat_NULL{}
	states[HSM_SYSSTAT_ACTIVATE] = hsm_systat_ACTIVATE{}
	states[HSM_SYSSTAT_LIVE] = hsm_systat_LIVE{}
	states[HSM_SYSSTAT_DEBUG] = hsm_systat_DEBUG{}
	states[HSM_SYSSTAT_PREHIBERNATE] = hsm_systat_PREHIBERNATE{}
	states[HSM_SYSSTAT_HIBERNATE] = hsm_systat_HIBERNATE{}
	states[HSM_SYSSTAT_POSTHIBERNATE] = hsm_systat_POSTHIBERNATE{}
	states[HSM_SYSSTAT_DEACTIVATE] = hsm_systat_DEACTIVATE{}
	states[HSM_SYSSTAT_INITHSM] = hsm_systat_INITHSM{}
	states[HSM_SYSSTAT_ENTERHIERARCHY] = hsm_systat_ENTERHIERARCHY{}
	states[HSM_SYSSTAT_EXITHIERARCHY] = hsm_systat_EXITHIERARCHY{}
	states[HSM_SYSSTAT_BLOCKING] = hsm_systat_BLOCKING{}
	states[HSM_SYSSTAT_DEMATERIALISE] = hsm_systat_DEMATERIALISE{}
	states[HSM_SYSSTAT_REMATERIALISE] = hsm_systat_REMATERIALISE{}
	states[HSM_SYSSTAT_REGENERATING] = hsm_systat_REGENERATING{}
	states[HSM_SYSSTAT_DEBUGSTEP] = hsm_systat_DEBUGSTEP{}
	states[HSM_SYSSTAT_DEBUGTRACE] = hsm_systat_DEBUGTRACE{}
	states[HSM_SYSSTAT_FINALISED] = hsm_systat_FINALISED{}
	states[HSM_SYSSTAT_DEAD] = hsm_systat_DEAD{}

	hsa := HsmActorBaseLayer{HSM_SYSSTAT_NULL, HSM_SYSSTAT_ACTIVATE, states, nil}
	return hsa
}

/////////////////////////////////////////////////////////////////////////////////
// STATE DEFINITIONS ////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////////////////////////////

//+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++//
//--- hsm_systat_NULL ---------------------------------------------------------//

type hsm_systat_NULL struct{}

func (hst hsm_systat_NULL) acceptState(layer *HsmActorLayer) int { return HSM_SYSSTAT_ACTIVATE }
func (hst hsm_systat_NULL) getState() int                        { return HSM_SYSSTAT_NULL }
func (hst hsm_systat_NULL) describeState() string                { return "HSM_SYSSTAT_NULL" }

//+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++//
//--- hsm_systat_ACTIVATE -----------------------------------------------------//

type hsm_systat_ACTIVATE struct{}

func (hst hsm_systat_ACTIVATE) acceptState(layer *HsmActorLayer) int { return HSM_SYSSTAT_LIVE }
func (hst hsm_systat_ACTIVATE) getState() int                        { return HSM_SYSSTAT_ACTIVATE }
func (hst hsm_systat_ACTIVATE) describeState() string                { return "HSM_SYSSTAT_ACTIVATE" }

//+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++//
//--- hsm_systat_LIVE ---------------------------------------------------------//

type hsm_systat_LIVE struct{}

func (hst hsm_systat_LIVE) acceptState(layer *HsmActorLayer) int {

	m1 := *layer
	actorp := m1.getActor()
	actor := *actorp
	ul := actor.getUserLayerSegment(HSM_SYSSTAT_LIVE)
	st := m1.getStates()
	cs := st[m1.getCurrentState()]

	fmt.Printf("\tThe Current State of The World is:  # %v  %v \r\n", m1.getCurrentState(), cs.describeState())

	if ul.CurrentState < HSM_USRSTAT_ENTER {
		return HSM_SYSSTAT_DEBUG
	}

	var retval int
	//var ucheck HsmActorLayer{} = m1 ;

	for ul.CurrentState < HSM_USRSTAT_EXIT {

		if ul.LastState != ul.CurrentState {
			if m1.getLastState() != HSM_SYSSTAT_EXITHIERARCHY {
				m1.setLastState(m1.getCurrentState())
				m1.setCurrentState(HSM_SYSSTAT_EXITHIERARCHY)
				return HSM_SYSSTAT_EXITHIERARCHY
			} else {
				m1.setLastState(m1.getCurrentState())
				m1.setCurrentState(HSM_SYSSTAT_ENTERHIERARCHY)
				return HSM_SYSSTAT_ENTERHIERARCHY
			}

		}

		st := m1.getStates()
		cs := st[m1.getCurrentState()]

		fmt.Printf("\tThe Current State of The World is:  # %v  %v \r\n", m1.getCurrentState(), cs.describeState())

		uref := HsmActorLayer(ul)

		retval = cs.acceptState(&uref)

		if retval != HSM_SYSSTAT_NULL {

			m1.setLastState(m1.getCurrentState())
			m1.setCurrentState(retval)
			return m1.getCurrentState()
		}

	}

	return HSM_SYSSTAT_DEBUG
}

func (hst hsm_systat_LIVE) getState() int         { return HSM_SYSSTAT_LIVE }
func (hst hsm_systat_LIVE) describeState() string { return "HSM_SYSSTAT_LIVE" }

//+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++//
//--- hsm_systat_DEBUG ---------------------------------------------------------//

type hsm_systat_DEBUG struct{}

func (hst hsm_systat_DEBUG) acceptState(layer *HsmActorLayer) int { return HSM_SYSSTAT_PREHIBERNATE }
func (hst hsm_systat_DEBUG) getState() int                        { return HSM_SYSSTAT_DEBUG }
func (hst hsm_systat_DEBUG) describeState() string                { return "HSM_SYSSTAT_DEBUG" }

//+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++//
//--- hsm_systat_PREHIBERNATE -------------------------------------------------//

type hsm_systat_PREHIBERNATE struct{}

func (hst hsm_systat_PREHIBERNATE) acceptState(layer *HsmActorLayer) int {
	return HSM_SYSSTAT_HIBERNATE
}
func (hst hsm_systat_PREHIBERNATE) getState() int         { return HSM_SYSSTAT_PREHIBERNATE }
func (hst hsm_systat_PREHIBERNATE) describeState() string { return "HSM_SYSSTAT_PREHIBERNATE" }

//+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++//
//--- hsm_systat_HIBERNATE ----------------------------------------------------//

type hsm_systat_HIBERNATE struct{}

func (hst hsm_systat_HIBERNATE) acceptState(layer *HsmActorLayer) int {
	return HSM_SYSSTAT_POSTHIBERNATE
}
func (hst hsm_systat_HIBERNATE) getState() int         { return HSM_SYSSTAT_HIBERNATE }
func (hst hsm_systat_HIBERNATE) describeState() string { return "HSM_SYSSTAT_HIBERNATE" }

//+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++//
//--- hsm_systat_POSTHIBERNATE ------------------------------------------------//

type hsm_systat_POSTHIBERNATE struct{}

func (hst hsm_systat_POSTHIBERNATE) acceptState(layer *HsmActorLayer) int {
	return HSM_SYSSTAT_DEACTIVATE
}
func (hst hsm_systat_POSTHIBERNATE) getState() int         { return HSM_SYSSTAT_POSTHIBERNATE }
func (hst hsm_systat_POSTHIBERNATE) describeState() string { return "HSM_SYSSTAT_POSTHIBERNATE" }

//+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++//
//--- hsm_systat_DEACTIVATE ---------------------------------------------------//

type hsm_systat_DEACTIVATE struct{}

func (hst hsm_systat_DEACTIVATE) acceptState(layer *HsmActorLayer) int { return HSM_SYSSTAT_INITHSM }
func (hst hsm_systat_DEACTIVATE) getState() int                        { return HSM_SYSSTAT_DEACTIVATE }
func (hst hsm_systat_DEACTIVATE) describeState() string                { return "HSM_SYSSTAT_DEACTIVATE" }

//+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++//
//--- hsm_systat_INITHSM  -----------------------------------------------------//

type hsm_systat_INITHSM struct{}

func (hst hsm_systat_INITHSM) acceptState(layer *HsmActorLayer) int {
	return HSM_SYSSTAT_ENTERHIERARCHY
}
func (hst hsm_systat_INITHSM) getState() int         { return HSM_SYSSTAT_INITHSM }
func (hst hsm_systat_INITHSM) describeState() string { return "HSM_SYSSTAT_INITHSM" }

//+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++//
//--- hsm_systat_ENTERHIERARCHY   ---------------------------------------------//

type hsm_systat_ENTERHIERARCHY struct{}

func (hst hsm_systat_ENTERHIERARCHY) acceptState(layer *HsmActorLayer) int {
	return HSM_SYSSTAT_EXITHIERARCHY
}
func (hst hsm_systat_ENTERHIERARCHY) getState() int         { return HSM_SYSSTAT_ENTERHIERARCHY }
func (hst hsm_systat_ENTERHIERARCHY) describeState() string { return "HSM_SYSSTAT_ENTERHIERARCHY" }

//+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++//
//--- hsm_systat_EXITHIERARCHY   ----------------------------------------------//

type hsm_systat_EXITHIERARCHY struct{}

func (hst hsm_systat_EXITHIERARCHY) acceptState(layer *HsmActorLayer) int {
	return HSM_SYSSTAT_BLOCKING
}
func (hst hsm_systat_EXITHIERARCHY) getState() int         { return HSM_SYSSTAT_EXITHIERARCHY }
func (hst hsm_systat_EXITHIERARCHY) describeState() string { return "HSM_SYSSTAT_EXITHIERARCHY" }

//+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++//
//--- hsm_systat_BLOCKING  ---------------------------------------------------------//

type hsm_systat_BLOCKING struct{}

func (hst hsm_systat_BLOCKING) acceptState(layer *HsmActorLayer) int {
	return HSM_SYSSTAT_DEMATERIALISE
}
func (hst hsm_systat_BLOCKING) getState() int         { return HSM_SYSSTAT_BLOCKING }
func (hst hsm_systat_BLOCKING) describeState() string { return "HSM_SYSSTAT_BLOCKING" }

//+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++//
//--- hsm_systat_DEMATERIALISE  ---------------------------------------------------------//

type hsm_systat_DEMATERIALISE struct{}

func (hst hsm_systat_DEMATERIALISE) acceptState(layer *HsmActorLayer) int {
	return HSM_SYSSTAT_REMATERIALISE
}
func (hst hsm_systat_DEMATERIALISE) getState() int         { return HSM_SYSSTAT_DEMATERIALISE }
func (hst hsm_systat_DEMATERIALISE) describeState() string { return "HSM_SYSSTAT_DEMATERIALISE" }

//+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++//
//--- hsm_systat_REMATERIALISE  ---------------------------------------------------------//

type hsm_systat_REMATERIALISE struct{}

func (hst hsm_systat_REMATERIALISE) acceptState(layer *HsmActorLayer) int {
	return HSM_SYSSTAT_REGENERATING
}
func (hst hsm_systat_REMATERIALISE) getState() int         { return HSM_SYSSTAT_REMATERIALISE }
func (hst hsm_systat_REMATERIALISE) describeState() string { return "HSM_SYSSTAT_REMATERIALISE" }

//+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++//
//--- hsm_systat_REGENERATING  ------------------------------------------------//

type hsm_systat_REGENERATING struct{}

func (hst hsm_systat_REGENERATING) acceptState(layer *HsmActorLayer) int {
	return HSM_SYSSTAT_DEBUGSTEP
}
func (hst hsm_systat_REGENERATING) getState() int         { return HSM_SYSSTAT_REGENERATING }
func (hst hsm_systat_REGENERATING) describeState() string { return "HSM_SYSSTAT_REGENERATING" }

//+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++//
//--- hsm_systat_DEBUGSTEP  ---------------------------------------------------//

type hsm_systat_DEBUGSTEP struct{}

func (hst hsm_systat_DEBUGSTEP) acceptState(layer *HsmActorLayer) int { return HSM_SYSSTAT_DEBUGTRACE }
func (hst hsm_systat_DEBUGSTEP) getState() int                        { return HSM_SYSSTAT_DEBUGSTEP }
func (hst hsm_systat_DEBUGSTEP) describeState() string                { return "HSM_SYSSTAT_DEBUGSTEP" }

//+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++//
//--- hsm_systat_DEBUGTRACE  --------------------------------------------------//

type hsm_systat_DEBUGTRACE struct{}

func (hst hsm_systat_DEBUGTRACE) acceptState(layer *HsmActorLayer) int { return HSM_SYSSTAT_FINALISED }
func (hst hsm_systat_DEBUGTRACE) getState() int                        { return HSM_SYSSTAT_DEBUGTRACE }
func (hst hsm_systat_DEBUGTRACE) describeState() string                { return "HSM_SYSSTAT_DEBUGTRACE" }

//+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++//
//--- hsm_systat_FINALISED  ---------------------------------------------------//

type hsm_systat_FINALISED struct{}

func (hst hsm_systat_FINALISED) acceptState(layer *HsmActorLayer) int { return HSM_SYSSTAT_DEAD }
func (hst hsm_systat_FINALISED) getState() int                        { return HSM_SYSSTAT_FINALISED }
func (hst hsm_systat_FINALISED) describeState() string                { return "HSM_SYSSTAT_FINALISED" }

//+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++//
//--- hsm_systat_DEAD  ---------------------------------------------------------//

type hsm_systat_DEAD struct{}

func (hst hsm_systat_DEAD) acceptState(layer *HsmActorLayer) int { return HSM_SYSSTAT_DEAD }
func (hst hsm_systat_DEAD) getState() int                        { return HSM_SYSSTAT_DEAD }
func (hst hsm_systat_DEAD) describeState() string                { return "HSM_SYSSTAT_DEAD" }

////////////////////////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////////////////////

/////////////////////////////////////////////////////////////////////////////////
// SYSTEM STATE TOKENS //////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////////////////////////////

const (
	HSM_USRSTAT_ENTER = (HSM_SYSSTAT_USER + iota) // NULL PSEUDO STATE
	HSM_USRSTAT_01                                // NULL PSEUDO STATE
	HSM_USRSTAT_01_01                             // NULL PSEUDO STATE
	HSM_USRSTAT_01_02                             // NULL PSEUDO STATE
	HSM_USRSTAT_02                                // NULL PSEUDO STATE
	HSM_USRSTAT_02_01                             // NULL PSEUDO STATE
	HSM_USRSTAT_02_02                             // NULL PSEUDO STATE
	HSM_USRSTAT_EXIT                              // NULL PSEUDO STATE
)

////////////////////////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////////////////////

//+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++//
//--- hsm_systat_ENTER ---------------------------------------------------------//

type hsm_usrstat_ENTER struct{}

func (hst hsm_usrstat_ENTER) acceptState(layer *HsmActorLayer) int { return HSM_USRSTAT_01 }
func (hst hsm_usrstat_ENTER) getState() int                        { return HSM_USRSTAT_ENTER }
func (hst hsm_usrstat_ENTER) describeState() string                { return "HSM_USRSTAT_ENTER" }

//+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++//
//--- hsm_systat_01 ---------------------------------------------------------//

type hsm_usrstat_01 struct{}

func (hst hsm_usrstat_01) acceptState(layer *HsmActorLayer) int { return HSM_USRSTAT_01_01 }
func (hst hsm_usrstat_01) getState() int                        { return HSM_USRSTAT_01 }
func (hst hsm_usrstat_01) describeState() string                { return "HSM_USRSTAT_01" }

//+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++//
//--- hsm_systat_01_01 ---------------------------------------------------------//

type hsm_usrstat_01_01 struct{}

func (hst hsm_usrstat_01_01) acceptState(layer *HsmActorLayer) int { return HSM_USRSTAT_01_02 }
func (hst hsm_usrstat_01_01) getState() int                        { return HSM_USRSTAT_01_01 }
func (hst hsm_usrstat_01_01) describeState() string                { return "HSM_USRSTAT_01_01" }

//+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++//
//--- hsm_systat_01_02 ---------------------------------------------------------//

type hsm_usrstat_01_02 struct{}

func (hst hsm_usrstat_01_02) acceptState(layer *HsmActorLayer) int { return HSM_USRSTAT_02 }
func (hst hsm_usrstat_01_02) getState() int                        { return HSM_USRSTAT_01_02 }
func (hst hsm_usrstat_01_02) describeState() string                { return "HSM_USRSTAT_01_02" }

//+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++//
//--- hsm_systat_02 ---------------------------------------------------------//

type hsm_usrstat_02 struct{}

func (hst hsm_usrstat_02) acceptState(layer *HsmActorLayer) int { return HSM_USRSTAT_02_01 }
func (hst hsm_usrstat_02) getState() int                        { return HSM_USRSTAT_02 }
func (hst hsm_usrstat_02) describeState() string                { return "HSM_USRSTAT_02" }

//+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++//
//--- hsm_systat_02_01 ---------------------------------------------------------//

type hsm_usrstat_02_01 struct{}

func (hst hsm_usrstat_02_01) acceptState(layer *HsmActorLayer) int { return HSM_USRSTAT_02_02 }
func (hst hsm_usrstat_02_01) getState() int                        { return HSM_USRSTAT_02_01 }
func (hst hsm_usrstat_02_01) describeState() string                { return "HSM_USRSTAT_02_01" }

//+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++//
//--- hsm_systat_02_02 ---------------------------------------------------------//

type hsm_usrstat_02_02 struct{}

func (hst hsm_usrstat_02_02) acceptState(layer *HsmActorLayer) int { return HSM_USRSTAT_EXIT }
func (hst hsm_usrstat_02_02) getState() int                        { return HSM_USRSTAT_02_02 }
func (hst hsm_usrstat_02_02) describeState() string                { return "HSM_USRSTAT_02_02" }

//+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++//
//--- hsm_systat_EXIT ---------------------------------------------------------//

type hsm_usrstat_EXIT struct{}

func (hst hsm_usrstat_EXIT) acceptState(layer *HsmActorLayer) int { return HSM_USRSTAT_EXIT }
func (hst hsm_usrstat_EXIT) getState() int                        { return HSM_USRSTAT_EXIT }
func (hst hsm_usrstat_EXIT) describeState() string                { return "HSM_USRSTAT_EXIT " }

////////////////////////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////////////////////

//--- CreateUserLaer ---------------------------------------------------------//
//--- HsmActorUserLayer ------------------------------------------------------------//

type HsmActorUserLayer struct {
	LastState    int
	CurrentState int
	States       []HsmState
	actor        *HsmActor
}

func (hsa HsmActorUserLayer) getLastState() int { return hsa.LastState }

func (hsa HsmActorUserLayer) setLastState(state int) { hsa.LastState = state }

func (hsa HsmActorUserLayer) getCurrentState() int { return hsa.CurrentState }

func (hsa HsmActorUserLayer) setCurrentState(state int) { hsa.CurrentState = state }

func (hsa HsmActorUserLayer) getActor() *HsmActor {
	//fmt.Printf("\ty2:  \r\n")
	return hsa.actor
}

func (hsa *HsmActorUserLayer) setActor(actor *HsmActor) {

	//fmt.Printf("\ty1:  \r\n")
	hsa.actor = actor
}

func (hsa HsmActorUserLayer) initFrame(minState int, MaxState int) error { return nil }

func (hsa HsmActorUserLayer) getStates() []HsmState { return hsa.States }

func (hsa HsmActorUserLayer) Live(layer *HsmActorLayer) (exitstate int, err error) {

	if hsa.CurrentState < HSM_USRSTAT_ENTER {
		return hsa.CurrentState, nil
	}

	for hsa.CurrentState < HSM_USRSTAT_EXIT {

		fmt.Printf("\tThe Current State of The World is:  # %v  %v \r\n", hsa.CurrentState, hsa.States[hsa.CurrentState].describeState())

		hsal := HsmActorLayer(hsa)

		hsa.CurrentState = hsa.States[hsa.CurrentState].acceptState(&hsal)

	}

	return hsa.CurrentState, nil

}

func Create_HsmActorUserLayer() HsmActorUserLayer {

	var states []HsmState = make([]HsmState, HSM_USRSTAT_EXIT+1)

	states[HSM_USRSTAT_ENTER] = hsm_usrstat_ENTER{}
	states[HSM_USRSTAT_01] = hsm_usrstat_01{}
	states[HSM_USRSTAT_02] = hsm_usrstat_02{}
	states[HSM_USRSTAT_01_01] = hsm_usrstat_01_01{}
	states[HSM_USRSTAT_01_02] = hsm_usrstat_01_02{}
	states[HSM_USRSTAT_02_01] = hsm_usrstat_02_01{}
	states[HSM_USRSTAT_02_02] = hsm_usrstat_02_02{}
	states[HSM_USRSTAT_EXIT] = hsm_usrstat_EXIT{}

	hsa := HsmActorUserLayer{HSM_USRSTAT_ENTER, HSM_USRSTAT_01, states, nil}
	return hsa
}

////////////////////////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////////////////////

type MyActor struct {
	BaseLayer HsmActorBaseLayer
	UserLayer []HsmActorUserLayer
}

func (hsa *MyActor) GoLive() (exitstate int, err error) {

	//fmt.Printf("\r\ngolive 1 \r\n")

	hbl := hsa.BaseLayer

	lay := HsmActorLayer(hbl)

	myexit, myerr := hbl.Live(&lay)

	//fmt.Printf("\r\ngolive 2 \r\n")

	return myexit, myerr

}

func (hsa MyActor) getUserLayerSegment(seg int) HsmActorUserLayer {
	return hsa.UserLayer[seg]
}

func Create_MyActor() MyActor {

	//fmt.Printf("\r\nCREATE ACTOR PRE 1 \r\n")

	var usrlayer []HsmActorUserLayer = make([]HsmActorUserLayer, HSM_SYSSTAT_DEAD+1)

	//fmt.Printf("\r\nCREATE ACTOR PRE 2 \r\n")

	uslayer := Create_HsmActorUserLayer()

	usrlayer[HSM_SYSSTAT_NULL] = uslayer
	usrlayer[HSM_SYSSTAT_ACTIVATE] = uslayer
	usrlayer[HSM_SYSSTAT_LIVE] = uslayer
	usrlayer[HSM_SYSSTAT_DEBUG] = uslayer
	usrlayer[HSM_SYSSTAT_PREHIBERNATE] = uslayer
	usrlayer[HSM_SYSSTAT_HIBERNATE] = uslayer
	usrlayer[HSM_SYSSTAT_POSTHIBERNATE] = uslayer
	usrlayer[HSM_SYSSTAT_DEACTIVATE] = uslayer
	usrlayer[HSM_SYSSTAT_INITHSM] = uslayer
	usrlayer[HSM_SYSSTAT_ENTERHIERARCHY] = uslayer
	usrlayer[HSM_SYSSTAT_EXITHIERARCHY] = uslayer
	usrlayer[HSM_SYSSTAT_BLOCKING] = uslayer
	usrlayer[HSM_SYSSTAT_DEMATERIALISE] = uslayer
	usrlayer[HSM_SYSSTAT_REMATERIALISE] = uslayer
	usrlayer[HSM_SYSSTAT_REGENERATING] = uslayer
	usrlayer[HSM_SYSSTAT_DEBUGSTEP] = uslayer
	usrlayer[HSM_SYSSTAT_DEBUGTRACE] = uslayer
	usrlayer[HSM_SYSSTAT_FINALISED] = uslayer

	baslayer := Create_HsmActorBaseLayer()

	hsa := MyActor{baslayer, usrlayer}
	hsal := HsmActor(hsa)

	//if nil == hsal  { fmt.Printf("\tx2no! actor:  \r\n") }

	hsa.BaseLayer.setActor(&hsal)

	hsa.BaseLayer.getActor()

	//if nil == * hsa.BaseLayer.getActor( ) { fmt.Printf("\tx2! actor:  \r\n") }

	//fmt.Printf("\r\nCREATE ACTOR PRE 3 \r\n")

	hsa.UserLayer[HSM_SYSSTAT_NULL].setActor(&hsal)

	//fmt.Printf("\r\nCREATE ACTOR PRE 4 \r\n")

	hsa.UserLayer[HSM_SYSSTAT_ACTIVATE].setActor(&hsal)
	hsa.UserLayer[HSM_SYSSTAT_LIVE].setActor(&hsal)
	hsa.UserLayer[HSM_SYSSTAT_DEBUG].setActor(&hsal)
	hsa.UserLayer[HSM_SYSSTAT_PREHIBERNATE].setActor(&hsal)
	hsa.UserLayer[HSM_SYSSTAT_HIBERNATE].setActor(&hsal)
	hsa.UserLayer[HSM_SYSSTAT_POSTHIBERNATE].setActor(&hsal)
	hsa.UserLayer[HSM_SYSSTAT_DEACTIVATE].setActor(&hsal)
	hsa.UserLayer[HSM_SYSSTAT_INITHSM].setActor(&hsal)
	hsa.UserLayer[HSM_SYSSTAT_ENTERHIERARCHY].setActor(&hsal)
	hsa.UserLayer[HSM_SYSSTAT_EXITHIERARCHY].setActor(&hsal)
	hsa.UserLayer[HSM_SYSSTAT_BLOCKING].setActor(&hsal)
	hsa.UserLayer[HSM_SYSSTAT_DEMATERIALISE].setActor(&hsal)
	hsa.UserLayer[HSM_SYSSTAT_REMATERIALISE].setActor(&hsal)
	hsa.UserLayer[HSM_SYSSTAT_REGENERATING].setActor(&hsal)
	hsa.UserLayer[HSM_SYSSTAT_DEBUGSTEP].setActor(&hsal)
	hsa.UserLayer[HSM_SYSSTAT_DEBUGTRACE].setActor(&hsal)
	hsa.UserLayer[HSM_SYSSTAT_FINALISED].setActor(&hsal)
	//hsa.UserLayer[HSM_SYSSTAT_DEAD].setActor( &hsal  )

	//fmt.Printf("\r\nCREATE ACTOR PRE 5 \r\n")

	//fmt.Printf("\r\nCREATE ACTOR POST ! \r\n")

	return hsa
}

////////////////////////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////////////////////
///

//////

type HSM interface {
	Activate()
}

type MyHSM struct {
	HsmActor

	STATE_01    State_01
	STATE_01_01 State_01_01
	STATE_02    State_02
}

////////////////////////////////////////////////////////////////////////////
// USER SPECIFIC IMPLEMENTATION ////////////////////////////////////////////

func (myhsm MyHSM) Activate() {

}

////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////

//////////////////////////////////////////////////////////////////////////////

/*

 	//------------------------------------------------------------------
 	HSM_SYSSTAT_INITIALISING	        	// NULL PSEUDO STATE
 	//------------------------------------------------------------------
  	HSM_SYSSTAT_STARTTRANSITION		        // NULL PSEUDO STATE
 	HSM_SYSSTAT_EXITCHAIN		         	// NULL PSEUDO STATE
  	HSM_SYSSTAT_ENTRYCHAIN			      	// NULL PSEUDO STATE
   	HSM_SYSSTAT_COMPLETETRANSITION		        // NULL PSEUDO STATE
 	//------------------------------------------------------------------
  	HSM_SYSSTAT_REQUESTNEWERA		      	// NULL PSEUDO STATE
  	HSM_SYSSTAT_PREPARENEWERA		        // NULL PSEUDO STATE
  	HSM_SYSSTAT_NEWERA		     		// NULL PSEUDO STATE
  	HSM_SYSSTAT_COMMITNEWERA		     	// NULL PSEUDO STATE
 	//------------------------------------------------------------------
  	HSM_SYSSTAT_RESTORING			      	// NULL PSEUDO STATE
 	//------------------------------------------------------------------
  	HSM_SYSSTAT_REQUESTSYNCPOINT		      	// NULL PSEUDO STATE
  	HSM_SYSSTAT_PREPARESYNCPOINT		      	// NULL PSEUDO STATE
  	HSM_SYSSTAT_SYNCPOINT		      		// NULL PSEUDO STATE
  	HSM_SYSSTAT_COMMITSYNCPOINT		      	// NULL PSEUDO STATE
 	//------------------------------------------------------------------
   	HSM_SYSSTAT_REQUESTCONFIG		      	// NULL PSEUDO STATE
   	HSM_SYSSTAT_PREPARECONFIGT		      	// NULL PSEUDO STATE
   	HSM_SYSSTAT_CONFIG		      		// NULL PSEUDO STATE
  	HSM_SYSSTAT_COMMITCONFIG		      	// NULL PSEUDO STATE
 	//------------------------------------------------------------------
 	HSM_SYSSTAT_CURRENTSTATE      			// NULL PSEUDO STATE
 	//------------------------------------------------------------------
 	HSM_SYSSTAT_TOP		      			// NULL PSEUDO STATE

*/
