package main

import (
	"fmt"
	"github.com/kpeu3i/gods4"
	"github.com/showcontroller/go-osc/osc"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	fmt.Println("Hello CLI")

	// Find all controllers connected to your machine via USB or Bluetooth
	controllers := gods4.Find()
	if len(controllers) == 0 {
		panic("No connected DS4 controllers found")
	}

	// Select first controller from the list
	controller := controllers[0]

	// Connect to the controller
	err := controller.Connect()
	if err != nil {
		panic(err)
	}

	log.Printf("* Controller #1 | %-10s | name: %s, connection: %s\n", "Connect", controller, controller.ConnectionType())

	// Disconnect controller when a program is terminated
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-signals
		err := controller.Disconnect()
		if err != nil {
			panic(err)
		}
		log.Printf("* Controller #1 | %-10s | bye!\n", "Disconnect")
		os.Exit(3)
	}()

	disp := osc.NewStandardDispatcher()
	client := osc.NewTCPClient("127.0.0.1:6979", disp)
	client.Connect()

	// Register callback for "BatteryUpdate" event
	controller.On(gods4.EventBatteryUpdate, func(data interface{}) error {
		battery := data.(gods4.Battery)
		log.Printf("* Controller #1 | %-10s | capacity: %v%%, charging: %v, cable: %v\n",
			"Battery",
			battery.Capacity,
			battery.IsCharging,
			battery.IsCableConnected,
		)

		return nil
	})

	controller.On(gods4.EventDPadDownPress, func(data interface{}) error {
		log.Printf("* Controller #1 | %-10s | state: press\n", "DPadDown")
		log.Println("sent /dpad_down/press")
		msg := osc.NewMessage("/dpad_down/press")
		client.Send(msg)

		return nil
	})

	controller.On(gods4.EventDPadDownRelease, func(data interface{}) error {
		log.Printf("* Controller #1 | %-10s | state: release\n", "DPadDown")
		log.Println("sent /dpad_down/release")
		msg := osc.NewMessage("/dpad_down/release")
		client.Send(msg)

		return nil
	})

	controller.On(gods4.EventDPadLeftPress, func(data interface{}) error {
		log.Printf("* Controller #1 | %-10s | state: press\n", "DPadLeft")
		log.Println("sent /dpad_left/press")
		msg := osc.NewMessage("/dpad_left/press")
		client.Send(msg)

		return nil
	})

	controller.On(gods4.EventDPadLeftRelease, func(data interface{}) error {
		log.Printf("* Controller #1 | %-10s | state: release\n", "DPadLeft")
		log.Println("sent /dpad_left/release")
		msg := osc.NewMessage("/dpad_left/release")
		client.Send(msg)

		return nil
	})

	controller.On(gods4.EventDPadRightPress, func(data interface{}) error {
		log.Printf("* Controller #1 | %-10s | state: press\n", "DPadLeft")
		log.Println("sent /dpad_right/press")
		msg := osc.NewMessage("/dpad_right/press")
		client.Send(msg)

		return nil
	})

	controller.On(gods4.EventDPadRightRelease, func(data interface{}) error {
		log.Printf("* Controller #1 | %-10s | state: release\n", "DPadRight")
		log.Println("sent /dpad_right/release")
		msg := osc.NewMessage("/dpad_right/release")
		client.Send(msg)

		return nil
	})

	controller.On(gods4.EventDPadUpPress, func(data interface{}) error {
		log.Printf("* Controller #1 | %-10s | state: press\n", "DPadUp")
		log.Println("sent /dpad_up/press")
		msg := osc.NewMessage("/dpad_up/press")
		client.Send(msg)

		return nil
	})

	controller.On(gods4.EventDPadUpRelease, func(data interface{}) error {
		log.Printf("* Controller #1 | %-10s | state: release\n", "DPadUp")
		log.Println("sent /dpad_up/release")
		msg := osc.NewMessage("/dpad_up/release")
		client.Send(msg)

		return nil
	})

	controller.On(gods4.EventCrossPress, func(data interface{}) error {
		log.Printf("* Controller #1 | %-10s | state: press\n", "Cross")
		log.Println("sent /cross/press")
		msg := osc.NewMessage("/cross/press")
		client.Send(msg)

		return nil
	})

	controller.On(gods4.EventCrossRelease, func(data interface{}) error {
		log.Printf("* Controller #1 | %-10s | state: release\n", "Cross")
		log.Println("sent /cross/release")
		msg := osc.NewMessage("/cross/release")
		client.Send(msg)

		return nil
	})

	controller.On(gods4.EventCirclePress, func(data interface{}) error {
		log.Printf("* Controller #1 | %-10s | state: press\n", "Circle")
		log.Println("sent /circle/press")
		msg := osc.NewMessage("/circle/press")
		client.Send(msg)

		return nil
	})

	controller.On(gods4.EventCircleRelease, func(data interface{}) error {
		log.Printf("* Controller #1 | %-10s | state: release\n", "Circle")
		log.Println("sent /circle/release")
		msg := osc.NewMessage("/circle/release")
		client.Send(msg)

		return nil
	})

	controller.On(gods4.EventSquarePress, func(data interface{}) error {
		log.Printf("* Controller #1 | %-10s | state: press\n", "Square")
		log.Println("sent /square/press")
		msg := osc.NewMessage("/square/press")
		client.Send(msg)

		return nil
	})

	controller.On(gods4.EventSquareRelease, func(data interface{}) error {
		log.Printf("* Controller #1 | %-10s | state: release\n", "Square")
		log.Println("sent /square/release")
		msg := osc.NewMessage("/square/release")
		client.Send(msg)

		return nil
	})

	controller.On(gods4.EventTrianglePress, func(data interface{}) error {
		log.Printf("* Controller #1 | %-10s | state: press\n", "Triangle")
		log.Println("sent /triangle/press")
		msg := osc.NewMessage("/triangle/press")
		client.Send(msg)

		return nil
	})

	controller.On(gods4.EventTriangleRelease, func(data interface{}) error {
		log.Printf("* Controller #1 | %-10s | state: release\n", "Triangle")
		log.Println("sent /triangle/release")
		msg := osc.NewMessage("/triangle/release")
		client.Send(msg)

		return nil
	})

	controller.On(gods4.EventL1Press, func(data interface{}) error {
		log.Printf("* Controller #1 | %-10s | state: press\n", "L1")
		log.Println("sent /l1/press")
		msg := osc.NewMessage("/l1/press")
		client.Send(msg)

		return nil
	})

	controller.On(gods4.EventL1Release, func(data interface{}) error {
		log.Printf("* Controller #1 | %-10s | state: release\n", "L1")
		log.Println("sent /l1/release")
		msg := osc.NewMessage("/l1/release")
		client.Send(msg)

		return nil
	})

	controller.On(gods4.EventL2Press, func(data interface{}) error {
		log.Printf("* Controller #1 | %-10s | state: press\n", "L2")
		log.Println("sent /l2/press")
		msg := osc.NewMessage("/l2/press")
		client.Send(msg)

		return nil
	})

	controller.On(gods4.EventL2Release, func(data interface{}) error {
		log.Printf("* Controller #1 | %-10s | state: release\n", "L2")
		log.Println("sent /l2/release")
		msg := osc.NewMessage("/l2/release")
		client.Send(msg)

		return nil
	})
	controller.On(gods4.EventL3Press, func(data interface{}) error {
		log.Printf("* Controller #1 | %-10s | state: press\n", "L3")
		log.Println("sent /l3/press")
		msg := osc.NewMessage("/l3/press")
		client.Send(msg)

		return nil
	})

	controller.On(gods4.EventL3Release, func(data interface{}) error {
		log.Printf("* Controller #1 | %-10s | state: release\n", "L3")
		log.Println("sent /l3/release")
		msg := osc.NewMessage("/l3/release")
		client.Send(msg)

		return nil
	})

	controller.On(gods4.EventR1Press, func(data interface{}) error {
		log.Printf("* Controller #1 | %-10s | state: press\n", "R1")
		log.Println("sent /r1/press")
		msg := osc.NewMessage("/r1/press")
		client.Send(msg)

		return nil
	})

	controller.On(gods4.EventR1Release, func(data interface{}) error {
		log.Printf("* Controller #1 | %-10s | state: release\n", "R1")
		log.Println("sent /r1/release")
		msg := osc.NewMessage("/r1/release")
		client.Send(msg)

		return nil
	})

	controller.On(gods4.EventR2Press, func(data interface{}) error {
		log.Printf("* Controller #1 | %-10s | state: press\n", "R2")
		log.Println("sent /r2/press")
		msg := osc.NewMessage("/r2/press")
		client.Send(msg)

		return nil
	})

	controller.On(gods4.EventR2Release, func(data interface{}) error {
		log.Printf("* Controller #1 | %-10s | state: release\n", "R2")
		log.Println("sent /r2/release")
		msg := osc.NewMessage("/r2/release")
		client.Send(msg)

		return nil
	})
	controller.On(gods4.EventR3Press, func(data interface{}) error {
		log.Printf("* Controller #1 | %-10s | state: press\n", "R3")
		log.Println("sent /r3/press")
		msg := osc.NewMessage("/r3/press")
		client.Send(msg)

		return nil
	})

	controller.On(gods4.EventR3Release, func(data interface{}) error {
		log.Printf("* Controller #1 | %-10s | state: release\n", "R3")
		log.Println("sent /r3/release")
		msg := osc.NewMessage("/r3/release")
		client.Send(msg)

		return nil
	})

	// Register callback for "RightStickMove" event
	//controller.On(gods4.EventRightStickMove, func(data interface{}) error {
	//	stick := data.(gods4.Stick)
	//	log.Printf("* Controller #1 | %-10s | x: %v, y: %v\n", "RightStick", stick.X, stick.Y)
	//	msg := osc.NewMessage("/rightstick/")
	//	msg.Append(int32(stick.X))
	//	msg.Append(int32(stick.Y))
	//	client.Send(msg)
	//	return nil
	//})

	// Enable left and right rumble motors
	//err = controller.Rumble(rumble.Both())
	//if err != nil {
	//	panic(err)
	//}

	// Enable LED (yellow) with flash
	//err = controller.Led(led.Yellow().Flash(50, 50))
	//if err != nil {
	//	panic(err)
	//}

	err = controller.Listen()
	if err != nil {
		panic(err)
	}
}
