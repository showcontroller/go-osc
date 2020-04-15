package main

import (
	"github.com/eiannone/keyboard"
	"log"

	"time"
)
import "fmt"

//import "github.com/Lobaro/slip"
import "github.com/showcontroller/go-osc/osc"

func main() {
	sd := osc.NewStandardDispatcher()
	err := sd.AddMsgHandler("*", func(msg *osc.Message) {
		log.Println("received a message ", msg.String())
		//osc.PrintMessage(msg)
	})
	if err != nil {
		log.Println("error adding message handler", err)
	}

	// make a couple clients, so we can have two lights
	tc := osc.NewTCPClient("10.0.0.184:8765", sd)
	tc.Connect()
	go tc.Listen()
	tc2 := osc.NewTCPClient("10.0.0.221:8765", sd)
	tc2.Connect()
	go tc2.Listen()

	m1 := osc.NewMessage("/led/1/high")
	m2 := osc.NewMessage("/led/1/low")
	m3 := osc.NewMessage("/led/2/high")
	m4 := osc.NewMessage("/led/2/low")
	m5 := osc.NewMessage("/led/3/high")
	m6 := osc.NewMessage("/led/3/low")

	t := 200 * time.Millisecond // wait time in between messages

	//tc.ReconnectWait = 1*time.Second

	err = keyboard.Open()
	if err != nil {
		panic(err)
	}

	fmt.Println("Press ESC to quit")
	for {
		char, key, err := keyboard.GetKey()
		fmt.Printf("You pressed: %q\r\n", char)
		if err != nil {
			panic(err)
		} else if key == keyboard.KeyEsc {
			break
		} else if char == '1' {
			tc.Send(m1)
		} else if char == '4' {
			tc.Send(m2)
		} else if char == '2' {
			tc.Send(m3)
		} else if char == '5' {
			tc.Send(m4)
		} else if char == '3' {
			tc.Send(m5)
		} else if char == '6' {
			tc.Send(m6)
		} else if char == '7' {
			tc2.Send(m1)
		} else if char == '8' {
			tc2.Send(m3)
		} else if char == '9' {
			tc2.Send(m5)
		} else if char == '0' {
			tc2.Send(m2)
			tc2.Send(m4)
			tc2.Send(m6)
		}
	}

	keyboard.Close()

	for {
		tc.Send(m1)
		tc2.Send(m1)
		time.Sleep(t)
		tc.Send(m2)
		tc2.Send(m2)
		time.Sleep(t)
		tc.Send(m3)
		tc2.Send(m3)
		time.Sleep(t)
		tc.Send(m4)
		tc2.Send(m4)
		time.Sleep(t)
		tc.Send(m5)
		tc2.Send(m5)
		time.Sleep(t)
		tc.Send(m6)
		tc2.Send(m6)
		time.Sleep(t)

	}
}
