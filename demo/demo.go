package main

import (
	"fmt"
	"github.com/huuzkee-go-hsm/hsmproto"
)

//var rootDir string = "./testData"
//type FileHash struct {
//	fullpath       string
//	hash           string
//	lastModifyTime time.Time
//}
//var fileHashes []FileHash

func main() {

	//take CLI input
	/*	dirPtr := flag.String("path", ".", " string")
		svrPtr := flag.String("server", "www.softlayer.com", " string")
		frqPtr := flag.Int("bkp-interval", 60, "backup interval in hours")
		comprsPtr := flag.Bool("compression", false, "a bool")
		encrptPtr := flag.Bool("encryption", false, "a bool")
		flag.Parse()
		dir_path := *dirPtr
		server := *svrPtr
		frequency := *frqPtr
		compress := *comprsPtr
		encrypt := *encrptPtr
	*/ //I am not including bakcup run time since

	var con hsmproto.Connector

	con = hsmproto.InitConnector()

	fmt.Printf("\r\nHello World !  MessageCls:  %v MessageType:  %v MessageCount: %v \r\n\r\n", con.MsgCls, con.MsgType, con.MsgCount)

	hsa := hsmproto.Create_MyActor()

	hsa.GoLive(&hsa)

	fmt.Printf("\r\nTHE END ! \r\n")

}
