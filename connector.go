package hsmproto


type connector struct {
 	init       chan int
	enter      chan int
	exit       chan int
	msgtyp     int
}

const ( 
	SIG_init 	= (1<iota)
	SIG_enter 	= (1<iota)
	SIG_exit 	= (1<iota)
)

func initConnector() connector {

	var ch_init        chan int
	var ch_enter       chan int
	var ch_exit        chan int
  
	ch_init 	= make (chan int)
	ch_enter 	= make (chan int)
	ch_exit 	= make (chan int)

	return connector{ ch_init, ch_enter, ch_exit, 0}

}