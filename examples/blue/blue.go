package main

import (
	"fmt"
	"github.com/jaedle/golang-tplink-hs100/pkg/configuration"
	"github.com/jaedle/golang-tplink-hs100/pkg/hs100"
	"github.com/showcontroller/go-osc/osc"
	"log"
	"time"
)

func main() {
	devices, err := hs100.Discover("10.0.0.0/24",
		configuration.Default().WithTimeout(5*time.Second),
	)

	if err != nil {
		panic(err)
	}

	log.Printf("Found devices: %d", len(devices))
	for _, d := range devices {
		name, _ := d.GetName()
		ip := d.Address
		log.Printf("Device name: %s", name)
		log.Printf("Device ip: %s", ip)
	}
	addr := "127.0.0.1:6969"

	h := hs100.NewHs100("10.0.0.104", configuration.Default().WithTimeout(3*time.Second))

	fmt.Println(h.GetName())
	d := osc.NewStandardDispatcher()
	err = d.AddMsgHandler("/blue/on", func(msg *osc.Message) {
		log.Println(msg.String())
		log.Println(devices[0].TurnOn())

	})
	if err != nil {
		log.Println("error adding message handler ", err)
	}
	err = d.AddMsgHandler("/blue/off", func(msg *osc.Message) {
		log.Println(msg.String())
		log.Println(devices[0].TurnOff())
	})
	if err != nil {
		log.Println("error adding message handler ", err)
	}
	server := &osc.Server{
		Addr:       addr,
		Dispatcher: d,
	}

	//log.Println(devices[0].IsOn())
	log.Println("started listening")
	log.Println(server.ListenAndServe())
}
