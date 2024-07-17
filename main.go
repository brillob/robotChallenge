package main

import (
	"fmt"
	"os"
	"strings"

	challenge "challenge/robotwarehouse/Challenge"

	"github.com/urfave/cli"
)

func main() {
	// Create a new warehouse
	warehouseInstance := challenge.WarehouseBuilder("R")
	robots := warehouseInstance.Robots()

	app := cli.NewApp()
	app.Name = "WarehouseApp"
	app.Usage = "A CLI application for moving robots in warehouse"
	app.Version = "1.0.0"
	// Define commands
	app.Commands = []cli.Command{
		{
			Name:    "move",
			Aliases: []string{"m"},
			Usage:   "Move the robot in the warehouse N for North, S for South, E for East, W for West. Eg: move NNN NN",
			Action: func(c *cli.Context) error {

				for _, args := range os.Args[2:] {
					message := args

					robot := robots[0]
					fmt.Println("Command:", message)
					taskID, _, err := robot.EnqueueTask(strings.TrimSpace(message))
					if err != nil {
						fmt.Println("Error enqueuing task:", err)
						return nil
					}
					fmt.Println("Task ID:", taskID)
					fmt.Println("X:", robots[0].CurrentState().X)
					fmt.Println("Y:", robots[0].CurrentState().Y)
					fmt.Println("HasCrate", robots[0].CurrentState().HasCrate)
				}
				return nil
			},
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err)
	}
}
