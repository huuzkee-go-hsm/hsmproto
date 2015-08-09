package hsmproto


type Connector struct {
 	Init       chan int
	Enter      chan int
	Exit       chan int
	MsgCls     int
	MsgType    int
	MsgCount   int
}

const ( 
	SIG_Init 	= (99 + iota)
	SIG_Enter 	 
	SIG_Exit 	 
)

func InitConnector() Connector {

	var ch_init        chan int
	var ch_enter       chan int
	var ch_exit        chan int
  
	ch_init 	= make (chan int)
	ch_enter 	= make (chan int)
	ch_exit 	= make (chan int)

	return Connector{ ch_init, ch_enter, ch_exit, SIG_Init, SIG_Enter, SIG_Exit }

}
