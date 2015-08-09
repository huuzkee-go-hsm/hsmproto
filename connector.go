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
	SIG_init 	= (99 + iota)
	SIG_enter 	 
	SIG_exit 	 
)

func InitConnector() Connector {

	var ch_init        chan int
	var ch_enter       chan int
	var ch_exit        chan int
  
	ch_init 	= make (chan int)
	ch_enter 	= make (chan int)
	ch_exit 	= make (chan int)

	return Connector{ ch_init, ch_enter, SIG_init, SIG_enter, SIG_exit }

}
