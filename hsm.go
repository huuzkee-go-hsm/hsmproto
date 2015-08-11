package hsmproto

// SYSTEM STATE TOKENS ///////////////////////////////////////////////////////////

const ( 
 	HSM_SYSSTAT_NULL     = (0 + iota)		// NULL PSEUDO STATE
 	HSM_SYSSTAT_ACTIVATE				// NULL PSEUDO STATE
 	HSM_SYSSTAT_LIVE				// NULL PSEUDO STATE
 	HSM_SYSSTAT_DEACTIVATE 			      	// NULL PSEUDO STATE
 	HSM_SYSSTAT_DEAD 			      	// NULL PSEUDO STATE
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
 	HSM_SYSSTAT_TOP		      			// NULL PSEUDO STATE
 	HSM_SYSSTAT_USER		      		// NULL PSEUDO STATE
)


/////

type HsmState interface {

	AcceptState () int

}
type HsmEngine struct {

	LastState 	HsmState 
	CurrentState 	HsmState 
	
}


//////


type HSM interface {

	Activate()  

}


type MyHSM struct {

	HsmEngine
	
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
