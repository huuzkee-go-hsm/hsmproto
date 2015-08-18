package hsmproto


import (		
		
		"fmt"   

)


// SYSTEM STATE TOKENS ///////////////////////////////////////////////////////////

const ( 
 	HSM_SYSSTAT_NULL     = (0 + iota)		// NULL PSEUDO STATE
 	HSM_SYSSTAT_ACTIVATE				// NULL PSEUDO STATE
 	HSM_SYSSTAT_LIVE				// NULL PSEUDO STATE
 	HSM_SYSSTAT_PREHIBERNATE			// NULL PSEUDO STATE
 	HSM_SYSSTAT_HIBERNATE				// NULL PSEUDO STATE
 	HSM_SYSSTAT_POSTHIBERNATE			// NULL PSEUDO STATE
 	HSM_SYSSTAT_DEACTIVATE 			      	// NULL PSEUDO STATE
 	HSM_SYSSTAT_DEAD 			      	// NULL PSEUDO STATE
 	HSM_SYSSTAT_USER		      		// NULL PSEUDO STATE
)


/////


type HsmStateFrame interface {

	getLastState () int
	setLastState ( state int ) 
	getCurrentState () int
	setCurrentState ( state int ) 

}

type HsmActor interface {

	HsmStateFrame 

}

//--- HsmActorBase ---------------------------------------------------------//

type HsmActorBase struct {

	LastState 	int
	CurrentState 	int
		
}

func ( hsa HsmActorBase) getLastState () int  {  return hsa.LastState }

func ( hsa HsmActorBase) setLastState ( state int ) { hsa.LastState = state }

func ( hsa HsmActorBase) getCurrentState () int {  return hsa.CurrentState }

func ( hsa HsmActorBase) setCurrentState ( state int )  { hsa.CurrentState = state }

func ( hsa HsmActorBase) Live() ( exitstate int, err error) {

	if hsa.CurrentState == HSM_SYSSTAT_HIBERNATE { 
	    hsa.CurrentState = HSM_SYSSTAT_POSTHIBERNATE 
	}

	for hsa.CurrentState != HSM_SYSSTAT_DEAD {
	
		fmt.Printf("\tThe Current State of The World is:  %v \r\n", hsa.CurrentState )
		
		hsa.CurrentState += 1
		
	}

	return hsa.CurrentState, nil 
    
} 

//--- CreateActorBase ---------------------------------------------------------//	

func  CreateActorBase() (HsmActorBase) {

	hsa := 		HsmActorBase { HSM_SYSSTAT_NULL, HSM_SYSSTAT_ACTIVATE }	
	return hsa
} 


////////////////////////////////////////////////////////////////////////////////////////////////

///

type HsmState interface {

	AcceptState () int

}



//////


type HSM interface {

	Activate()  

}


type MyHSM struct {

	HsmActor
	
	STATE_01 	State_01 
	STATE_01_01 	State_01_01 
	STATE_02 	State_02 
	
}

////////////////////////////////////////////////////////////////////////////
// USER SPECIFIC IMPLEMENTATION ////////////////////////////////////////////

func ( myhsm MyHSM ) Activate()  {
    
} 


////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////




type State struct {
	StateType     int
	StateObject   HsmState
}


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