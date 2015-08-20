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
	HSM_SYSSTAT_FINALISED                   // NULL PSEUDO STATE
	HSM_SYSSTAT_INITHSM                     // NULL PSEUDO STATE
	HSM_SYSSTAT_ENTERHIERARCHY              // NULL PSEUDO STATE
	HSM_SYSSTAT_EXITHIERARCHY               // NULL PSEUDO STATE
	HSM_SYSSTAT_BLOCKING                    // NULL PSEUDO STATE
	HSM_SYSSTAT_EXPORTING                   // NULL PSEUDO STATE
	HSM_SYSSTAT_IMPORTING                   // NULL PSEUDO STATE
	HSM_SYSSTAT_REGENERATING                // NULL PSEUDO STATE
	HSM_SYSSTAT_DEBUGSTEP                   // NULL PSEUDO STATE
	HSM_SYSSTAT_DEBUGTRACE                  // NULL PSEUDO STATE
	HSM_SYSSTAT_DEAD                        // NULL PSEUDO STATE
	HSM_SYSSTAT_USER                        // NULL PSEUDO STATE
)

/////////////////////////////////////////////////////////////////////////////////
// INTERFACES ///////////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////////////////////////////

type HsmState interface {
	acceptState() int
	getState() int
}

//--- HsmStateFrame -----------------------------------------------------------//

type HsmStateFrame interface {
	getLastState() int
	setLastState(state int)
	getCurrentState() int
	setCurrentState(state int)

	initFrame(minState int, maxState int)
}

//--- HsmActor ----------------------------------------------------------------//

type HsmActor interface {
	HsmStateFrame
}

/////////////////////////////////////////////////////////////////////////////////
// TYPES ////////////////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////////////////////////////

//--- HsmActorBase ------------------------------------------------------------//

type HsmActorBase struct {
	LastState    int
	CurrentState int
	States       []HsmState
}

func (hsa HsmActorBase) getLastState() int { return hsa.LastState }

func (hsa HsmActorBase) setLastState(state int) { hsa.LastState = state }

func (hsa HsmActorBase) getCurrentState() int { return hsa.CurrentState }

func (hsa HsmActorBase) setCurrentState(state int) { hsa.CurrentState = state }

func (hsa HsmActorBase) Live() (exitstate int, err error) {

	if hsa.CurrentState == HSM_SYSSTAT_HIBERNATE {
		hsa.CurrentState = HSM_SYSSTAT_POSTHIBERNATE
	}

	for hsa.CurrentState != HSM_SYSSTAT_DEAD {

		fmt.Printf("\tThe Current State of The World is:  %v \r\n", hsa.CurrentState)

		hsa.CurrentState = hsa.States[hsa.CurrentState].acceptState()

	}

	return hsa.CurrentState, nil

}

/////////////////////////////////////////////////////////////////////////////////
// STATE DEFINITIONS ////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////////////////////////////

//+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++//
//--- hsm_systat_NULL ---------------------------------------------------------//

type hsm_systat_NULL struct{}

func (hst hsm_systat_NULL) acceptState() int      { return HSM_SYSSTAT_ACTIVATE }
func (hst hsm_systat_NULL) getState() int         { return HSM_SYSSTAT_NULL }
func (hst hsm_systat_NULL) describeState() string { return "HSM_SYSSTAT_NULL" }

//+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++//
//--- hsm_systat_ACTIVATE -----------------------------------------------------//

type hsm_systat_ACTIVATE struct{}

func (hst hsm_systat_ACTIVATE) acceptState() int { return HSM_SYSSTAT_LIVE }
func (hst hsm_systat_ACTIVATE) getState() int    { return HSM_SYSSTAT_ACTIVATE }

//+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++//
//--- hsm_systat_LIVE ---------------------------------------------------------//

type hsm_systat_LIVE struct{}

func (hst hsm_systat_LIVE) acceptState() int { return HSM_SYSSTAT_DEBUG }
func (hst hsm_systat_LIVE) getState() int    { return HSM_SYSSTAT_LIVE }

//+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++//
//--- hsm_systat_DEBUG ---------------------------------------------------------//

type hsm_systat_DEBUG struct{}

func (hst hsm_systat_DEBUG) acceptState() int { return HSM_SYSSTAT_PREHIBERNATE }
func (hst hsm_systat_DEBUG) getState() int    { return HSM_SYSSTAT_DEBUG }

//+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++//
//--- hsm_systat_PREHIBERNATE -------------------------------------------------//

type hsm_systat_PREHIBERNATE struct{}

func (hst hsm_systat_PREHIBERNATE) acceptState() int { return HSM_SYSSTAT_HIBERNATE }
func (hst hsm_systat_PREHIBERNATE) getState() int    { return HSM_SYSSTAT_PREHIBERNATE }

//+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++//
//--- hsm_systat_HIBERNATE ----------------------------------------------------//

type hsm_systat_HIBERNATE struct{}

func (hst hsm_systat_HIBERNATE) acceptState() int { return HSM_SYSSTAT_POSTHIBERNATE }
func (hst hsm_systat_HIBERNATE) getState() int    { return HSM_SYSSTAT_HIBERNATE }

//+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++//
//--- hsm_systat_POSTHIBERNATE ------------------------------------------------//

type hsm_systat_POSTHIBERNATE struct{}

func (hst hsm_systat_POSTHIBERNATE) acceptState() int { return HSM_SYSSTAT_DEACTIVATE }
func (hst hsm_systat_POSTHIBERNATE) getState() int    { return HSM_SYSSTAT_POSTHIBERNATE }

//+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++//
//--- hsm_systat_DEACTIVATE ---------------------------------------------------//

type hsm_systat_DEACTIVATE struct{}

func (hst hsm_systat_DEACTIVATE) acceptState() int { return HSM_SYSSTAT_FINALISED }
func (hst hsm_systat_DEACTIVATE) getState() int    { return HSM_SYSSTAT_DEACTIVATE }

//+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++//
//--- hsm_systat_FINALISED  ---------------------------------------------------//

type hsm_systat_FINALISED struct{}

func (hst hsm_systat_FINALISED) acceptState() int { return HSM_SYSSTAT_INITHSM }
func (hst hsm_systat_FINALISED) getState() int    { return HSM_SYSSTAT_FINALISED }

//+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++//
//--- hsm_systat_INITHSM  -----------------------------------------------------//

type hsm_systat_INITHSM struct{}

func (hst hsm_systat_INITHSM) acceptState() int { return HSM_SYSSTAT_ENTERHIERARCHY }
func (hst hsm_systat_INITHSM) getState() int    { return HSM_SYSSTAT_INITHSM }

//+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++//
//--- hsm_systat_ENTERHIERARCHY   ---------------------------------------------//

type hsm_systat_ENTERHIERARCHY struct{}

func (hst hsm_systat_ENTERHIERARCHY) acceptState() int { return HSM_SYSSTAT_EXITHIERARCHY }
func (hst hsm_systat_ENTERHIERARCHY) getState() int    { return HSM_SYSSTAT_ENTERHIERARCHY }

//+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++//
//--- hsm_systat_EXITHIERARCHY   ----------------------------------------------//

type hsm_systat_EXITHIERARCHY struct{}

func (hst hsm_systat_EXITHIERARCHY) acceptState() int { return HSM_SYSSTAT_BLOCKING }
func (hst hsm_systat_EXITHIERARCHY) getState() int    { return HSM_SYSSTAT_EXITHIERARCHY }

//+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++//
//--- hsm_systat_BLOCKING  ---------------------------------------------------------//

type hsm_systat_BLOCKING struct{}

func (hst hsm_systat_BLOCKING) acceptState() int { return HSM_SYSSTAT_EXPORTING }
func (hst hsm_systat_BLOCKING) getState() int    { return HSM_SYSSTAT_BLOCKING }

//+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++//
//--- hsm_systat_EXPORTING  ---------------------------------------------------------//

type hsm_systat_EXPORTING struct{}

func (hst hsm_systat_EXPORTING) acceptState() int { return HSM_SYSSTAT_IMPORTING }
func (hst hsm_systat_EXPORTING) getState() int    { return HSM_SYSSTAT_EXPORTING }

//+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++//
//--- hsm_systat_IMPORTING  ---------------------------------------------------------//

type hsm_systat_IMPORTING struct{}

func (hst hsm_systat_IMPORTING) acceptState() int { return HSM_SYSSTAT_REGENERATING }
func (hst hsm_systat_IMPORTING) getState() int    { return HSM_SYSSTAT_IMPORTING }

//+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++//
//--- hsm_systat_REGENERATING  ------------------------------------------------//

type hsm_systat_REGENERATING struct{}

func (hst hsm_systat_REGENERATING) acceptState() int { return HSM_SYSSTAT_DEBUGSTEP }
func (hst hsm_systat_REGENERATING) getState() int    { return HSM_SYSSTAT_REGENERATING }

//+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++//
//--- hsm_systat_DEBUGSTEP  ---------------------------------------------------//

type hsm_systat_DEBUGSTEP struct{}

func (hst hsm_systat_DEBUGSTEP) acceptState() int { return HSM_SYSSTAT_DEBUGTRACE }
func (hst hsm_systat_DEBUGSTEP) getState() int    { return HSM_SYSSTAT_DEBUGSTEP }

//+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++//
//--- hsm_systat_DEBUGTRACE  --------------------------------------------------//

type hsm_systat_DEBUGTRACE struct{}

func (hst hsm_systat_DEBUGTRACE) acceptState() int { return HSM_SYSSTAT_DEAD }
func (hst hsm_systat_DEBUGTRACE) getState() int    { return HSM_SYSSTAT_DEBUGTRACE }

//+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++//
//--- hsm_systat_DEAD  ---------------------------------------------------------//

type hsm_systat_DEAD struct{}

func (hst hsm_systat_DEAD) acceptState() int { return HSM_SYSSTAT_DEAD }
func (hst hsm_systat_DEAD) getState() int    { return HSM_SYSSTAT_DEAD }

//--- CreateActorBase ---------------------------------------------------------//

func CreateActorBase() HsmActorBase {

	var states []HsmState = make([]HsmState, HSM_SYSSTAT_DEAD+1)

	states[HSM_SYSSTAT_NULL] = hsm_systat_NULL{}
	states[HSM_SYSSTAT_ACTIVATE] = hsm_systat_ACTIVATE{}
	states[HSM_SYSSTAT_LIVE] = hsm_systat_LIVE{}
	states[HSM_SYSSTAT_LIVE] = hsm_systat_LIVE{}
	states[HSM_SYSSTAT_PREHIBERNATE] = hsm_systat_PREHIBERNATE{}
	states[HSM_SYSSTAT_HIBERNATE] = hsm_systat_HIBERNATE{}
	states[HSM_SYSSTAT_POSTHIBERNATE] = hsm_systat_POSTHIBERNATE{}
	states[HSM_SYSSTAT_DEACTIVATE] = hsm_systat_DEACTIVATE{}
	states[HSM_SYSSTAT_FINALISED] = hsm_systat_FINALISED{}
	states[HSM_SYSSTAT_INITHSM] = hsm_systat_INITHSM{}
	states[HSM_SYSSTAT_ENTERHIERARCHY] = hsm_systat_ENTERHIERARCHY{}
	states[HSM_SYSSTAT_EXITHIERARCHY] = hsm_systat_EXITHIERARCHY{}
	states[HSM_SYSSTAT_BLOCKING] = hsm_systat_BLOCKING{}
	states[HSM_SYSSTAT_EXPORTING] = hsm_systat_EXPORTING{}
	states[HSM_SYSSTAT_IMPORTING] = hsm_systat_IMPORTING{}
	states[HSM_SYSSTAT_REGENERATING] = hsm_systat_REGENERATING{}
	states[HSM_SYSSTAT_DEBUGSTEP] = hsm_systat_DEBUGSTEP{}
	states[HSM_SYSSTAT_DEBUGTRACE] = hsm_systat_DEBUGTRACE{}
	states[HSM_SYSSTAT_DEAD] = hsm_systat_DEAD{}

	hsa := HsmActorBase{HSM_SYSSTAT_NULL, HSM_SYSSTAT_ACTIVATE, states}
	return hsa
}

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
