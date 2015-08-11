package hsmproto


type State struct {
	StateType     int
	StateObject   HsmState
}

type HsmState interface {

	AcceptState () int

}


const ( 
 	HSM_SYSSTAT_NULL     = (0 + iota)		// NULL PSEUDO STATE
 	HSM_SYSSTAT_ACTIVATE				// NULL PSEUDO STATE
 	HSM_SYSSTAT_DEACTIVATE 			      	// NULL PSEUDO STATE
 	HSM_SYSSTAT_INITIALISE		        	// NULL PSEUDO STATE
  	HSM_SYSSTAT_TRANSITION		        	// NULL PSEUDO STATE
 	HSM_SYSSTAT_EXITCHAIN		         	// NULL PSEUDO STATE
  	HSM_SYSSTAT_ENTRYCHAIN			      	// NULL PSEUDO STATE
  	HSM_SYSSTAT_REQUESTNEWERA		      	// NULL PSEUDO STATE
  	HSM_SYSSTAT_STARTNEWERA		        	// NULL PSEUDO STATE
  	HSM_SYSSTAT_AWAITNEWERA			      	// NULL PSEUDO STATE
  	HSM_SYSSTAT_RESTOREERA			      	// NULL PSEUDO STATE
  	HSM_SYSSTAT_REQUESTSYNCPOINT		      	// NULL PSEUDO STATE
  	HSM_SYSSTAT_SYNCPOINT		      		// NULL PSEUDO STATE
  	HSM_SYSSTAT_COMPLETESYNCPOINT		      	// NULL PSEUDO STATE
 	HSM_SYSSTAT_TOP		      			// NULL PSEUDO STATE
 	HSM_SYSSTAT_USER		      		// NULL PSEUDO STATE
)


type HsmEngine struct {

	LastState 	HsmState 
	CurrentState 	HsmState 
	
}




type HSM interface {

	Activate()  

}


type MyHSM struct {

	HsmEngine
	
	STATE_01 	State_01 
	STATE_01_01 	State_01_01 
	STATE_02 	State_02 
	
}


func ( myhsm MyHSM ) Activate()  {
    
} 