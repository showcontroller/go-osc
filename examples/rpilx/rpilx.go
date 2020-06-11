package main

import (
	"github.com/showcontroller/go-osc/osc"
	"github.com/stianeikeland/go-rpio/v4"
	"log"
	"os"
)

func main() {

	// setup gpio pins
	pRed := rpio.Pin(26)
	pYellow := rpio.Pin(19)
	pGreen := rpio.Pin(13)
	if err := rpio.Open(); err != nil {
		log.Println("error opening gpio pins: ", err)
		os.Exit(1)
	}
	// unmap gpio when done

	defer log.Println(rpio.Close())
	// set pins to output mode
	pRed.Output()
	pYellow.Output()
	pGreen.Output()

	addr := ":8765"

	sd := osc.NewStandardDispatcher()
	err := sd.AddMsgHandler("*", func(msg *osc.Message) {
		log.Println("received a message ", msg.String())
	})
	if err != nil {
		log.Println("error adding default message handler", err)
	}
	err = sd.AddMsgHandler("/led/1/*", func(msg *osc.Message) {
		log.Println("received a led 1 message ", msg.String())
	})
	if err != nil {
		log.Println("error adding message handler", err)
	}
	err = sd.AddMsgHandler("/led/1/high", func(msg *osc.Message) {
		log.Println("received a led 1 high message ", msg.String())
		pRed.High()
	})
	if err != nil {
		log.Println("error adding led 1 high message handler", err)
	}
	err = sd.AddMsgHandler("/led/1/low", func(msg *osc.Message) {
		log.Println("received a led 1 low message ", msg.String())
		pRed.Low()
	})
	if err != nil {
		log.Println("error adding led 1 low message handler", err)
	}
	err = sd.AddMsgHandler("/led/2/high", func(msg *osc.Message) {
		log.Println("received a led 2 high message ", msg.String())
		pYellow.High()
	})
	if err != nil {
		log.Println("error adding led 2 high message handler", err)
	}
	err = sd.AddMsgHandler("/led/2/low", func(msg *osc.Message) {
		log.Println("received a led 2 low message ", msg.String())
		pYellow.Low()
	})
	if err != nil {
		log.Println("error adding led 2 low message handler", err)
	}
	err = sd.AddMsgHandler("/led/3/high", func(msg *osc.Message) {
		log.Println("received a led 3 high message ", msg.String())
		pGreen.High()
	})
	if err != nil {
		log.Println("error adding led 3 high message handler", err)
	}
	err = sd.AddMsgHandler("/led/3/low", func(msg *osc.Message) {
		log.Println("received a led 3 low message ", msg.String())
		pGreen.Low()
	})
	if err != nil {
		log.Println("error adding led 3 low message handler", err)
	}

	s := osc.TCPServer{Addr: addr, Dispatch: sd}
	log.Println(s.ListenServe())

}
