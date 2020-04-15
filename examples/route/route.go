package main

import (
	"github.com/showcontroller/go-osc/osc"
	"log"
)

func main() {
	// listen in on udp for qlab
	// send it out to two different tcp clients

	qDispatcher := osc.NewStandardDispatcher()
	udpQLabServer := osc.Server{Addr: ":6969", Dispatcher: qDispatcher}
	uDispatcher := osc.NewStandardDispatcher()
	tcpDS4Server := osc.TCPServer{Addr: ":6979", Dispatch: uDispatcher}

	lxDispatcher := osc.NewStandardDispatcher()

	blueLx := osc.NewTCPClient("10.0.0.221:8765", lxDispatcher)
	blackLx := osc.NewTCPClient("10.0.0.184:8765", lxDispatcher)

	blueLx.Connect()
	blackLx.Connect()

	// make a couple clients, so we can have two lights
	tc := osc.NewTCPClient("127.0.0.1:53000", uDispatcher)
	tc.Connect()
	//go tc.Listen()
	//tc2 := osc.NewTCPClient("10.0.0.221:8765", uDispatcher)
	//tc2.Connect()
	//go tc2.Listen()

	err := uDispatcher.AddMsgHandler("*", func(msg *osc.Message) {
		log.Println("received a message from ds4 ", msg.String())
		//log.Println(msg.Address, msg.Match("*"))
		tc.Send(msg)
		//tc2.Send(msg)

	})
	if err != nil {
		log.Println("error adding message handler", err)
	}
	err = qDispatcher.AddMsgHandler("/led/*", func(msg *osc.Message) {
		log.Println("received a message from qlab ", msg.String())
		//log.Println(msg.Address, msg.Match("*"))
		blueLx.Send(msg)
		blackLx.Send(msg)
		//tc2.Send(msg)

	})
	if err != nil {
		log.Println("error adding message handler", err)
	}
	//err = uDispatcher.AddMsgHandler("/led/*", func(msg *osc.Message) {
	//	log.Println("received a message led ", msg.String())
	//
	//})
	//if err != nil {
	//	log.Println("error adding led message handler", err)
	//}

	//udpQLabServer.Dispatcher = uDispatcher
	go udpQLabServer.ListenAndServe()
	err = tcpDS4Server.ListenServe()
	log.Println(err)
}
