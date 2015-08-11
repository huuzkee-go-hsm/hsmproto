package hsmproto


type State struct {
	StateType     int
	StateObject   HsmState
}

type HsmState interface {

	AcceptState () int

}



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