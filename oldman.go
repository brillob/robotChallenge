package oldman

// "bufio"
// challenge "challenge/robotwarehouse/Challenge"
// "fmt"
// "os"
// "strings"

// "flag"

// "rsc.io/quote"

func oldman() {
	// fmt.Println(challenge.BestRobot())
	// fmt.Println(quote.Go())

	// // Define flags
	// greeting := flag.String("greeting", "Hello", "Specify a greeting message")
	// name := flag.String("name", "User", "Specify a name")

	// flag.Usage = func() {
	// 	//fmt.Fprintf(os.Stderr, "Usage: %s [flags]\n", os.Args[0])
	// 	flag.PrintDefaults()
	// }
	// // Parse flags
	// flag.Parse()
	// // Accessing flag values
	// fmt.Printf("%s, %s!\n", *greeting, *name)

	// reader := bufio.NewReader(os.Stdin)
	// // Prompt the user for input
	// fmt.Print("Enter your favorite programming language: ")
	// language, _ := reader.ReadString('\n')
	// // Trim the newline character
	// language = strings.TrimSpace(language)
	// fmt.Printf("You entered: %s\n", language)\

	// app := cli.NewApp()
	// app.Name = "ColorApp"
	// app.Usage = "A CLI application with colors"
	// app.Version = "1.0.0"
	// // Define commands
	// app.Commands = []cli.Command{
	// 	{
	// 		Name:    "message",
	// 		Aliases: []string{"m"},
	// 		Usage:   "Display a colorful message",
	// 		Action: func(c *cli.Context) error {
	// 			green := color.New(color.FgGreen).SprintFunc()
	// 			red := color.New(color.FgRed).SprintFunc()
	// 			message := c.Args().First()
	// 			if message == "success" {
	// 				fmt.Printf("Success: %s\n", green(message))
	// 			} else if message == "error" {
	// 				fmt.Printf("Error: %s\n", red(message))
	// 			} else {
	// 				fmt.Println("Unknown message type.")
	// 			}
	// 			return nil
	// 		},
	// 	},
	// 	// Add more commands here...
	// }
	// err := app.Run(os.Args)
	// if err != nil {
	// 	fmt.Println(err)
	// }

}
