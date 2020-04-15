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

	// Register callback for "CrossPress" event
	controller.On(gods4.EventCrossPress, func(data interface{}) error {
		log.Printf("* Controller #1 | %-10s | state: press\n", "Cross")

		fmt.Println("sent osc message to qlab")
		msg := osc.NewMessage("/cue/l2/go")
		client.Send(msg)

		return nil
	})

	controller.On(gods4.EventR3Press, func(data interface{}) error {
		log.Printf("* Controller #1 | %-10s | state: press\n", "R3")

		fmt.Println("sent osc message to qlab to PANIC")
		msg := osc.NewMessage("/panic")
		client.Send(msg)

		return nil
	})

	controller.On(gods4.EventCirclePress, func(data interface{}) error {
		log.Printf("* Controller #1 | %-10s | state: press\n", "Circle")

		fmt.Println("sent osc message to qlab to pause")
		msg := osc.NewMessage("/cue/selected/stop")
		client.Send(msg)

		return nil
	})

	controller.On(gods4.EventTrianglePress, func(data interface{}) error {
		log.Printf("* Controller #1 | %-10s | state: press\n", "Triangle")

		fmt.Println("sent osc message to qlab to pause")
		msg := osc.NewMessage("/pause")
		client.Send(msg)

		return nil
	})

	controller.On(gods4.EventSquarePress, func(data interface{}) error {
		log.Printf("* Controller #1 | %-10s | state: press\n", "Square")

		fmt.Println("sent osc message to qlab to resume")
		msg := osc.NewMessage("/resume")
		client.Send(msg)

		return nil
	})

	controller.On(gods4.EventDPadDownPress, func(data interface{}) error {
		log.Printf("* Controller #1 | %-10s | state: press\n", "DPadDown")

		fmt.Println("sent osc message to qlab")
		msg := osc.NewMessage("/playhead/next")
		client.Send(msg)

		return nil
	})

	controller.On(gods4.EventDPadUpPress, func(data interface{}) error {
		log.Printf("* Controller #1 | %-10s | state: press\n", "DPadUp")

		fmt.Println("sent osc message to qlab")
		msg := osc.NewMessage("/playhead/previous")
		client.Send(msg)

		return nil
	})

	// Register callback for "CrossRelease" event
	controller.On(gods4.EventCrossRelease, func(data interface{}) error {
		log.Printf("* Controller #1 | %-10s | state: release\n", "Cross")

		return nil
	})

	// Register callback for "RightStickMove" event
	controller.On(gods4.EventRightStickMove, func(data interface{}) error {
		stick := data.(gods4.Stick)
		log.Printf("* Controller #1 | %-10s | x: %v, y: %v\n", "RightStick", stick.X, stick.Y)

		return nil
	})

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
