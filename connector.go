package hsmproto


public type connector struct {
 	init       chan
	enter      chan
	exit       chan
	msg-typ    int
}

const ( 
	SIG_init 	= (1<iota)
	SIG_enter 	= (1<iota)
	SIG_exit 	= (1<iota)
)

public func initConnector() connector {

	ch_init 	= make (chan int)
	ch_enter 	= make (chan int)
	ch_exit 	= make (chan int)

	return connector{ ch_init, ch_enter, ch_exit, 0}

}
