package main

import (
	cli "github.com/urfave/cli/v2"
	// "ticky/pkg/ohmy"
	"fmt"
	"log"
	"os"
	"os/user"
)

func main() {
	app := &cli.App{
		Name:  "ticky",
		Usage: "ticky is a ticket system...",
		Action: func(c *cli.Context) error {
			fmt.Println("Hello friend!")
			return nil
		},
		Commands: []*cli.Command{
			{
				Name:    "init",
				Aliases: []string{"i"},
				Usage:   "initialize filesystem (`.ticky`, e.g. `~/.ticky`) and `.metadata` config",
				Action:  func(c *cli.Context) error {
							if _, err := os.Stat("/Users/ippo/.ticky"); err != nil {
								if os.IsNotExist(err) {
									fmt.Println("init task: ")
								} else {
									fmt.Println("init task: 2 ")
								}
							}
					}
					//usr, err := user.Current()
					//if err != nil {
					//	// log.Fatal( err )
					//	return err
					//}  //homework, check if the .ticky dir exist and if not create it
					//
					// fmt.Println( "homedir:", usr.HomeDir )
					//
					//// fmt.Println("init task: ", c.Args().First())
					//return nil
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Println(ohmy.Add(4,6))

//we need smt like `ticky init` to create the basic setup (directory, metadata file etc)
// `$ ticky` will check if the besic setup exists
// if not print out "please run ticky init"
// if yes will show all the sub commands
//5 basics functions. "create", "show", "update", "list" and "delete"
//implement "create" function

// how to handle the id increment ?
// what happen when you have three tickets, with id 1,2,3
// you delete the first ticket with id one
// you create a new ticket, what id does the new ticket have ?  --> the next highest id + 1
// how do I remember the latest/higest id? (bec I could always delete the lates ticket, I still want the new ticket to have this id+1)
// => solution is: `.metadata` or `.idlock` file
// ticket file content format ? how does it look? (we want to solve parsing)
// solutions: `json` or `yaml`

// myvar := func() {
	// 	fmt.Println("dsfsdf")
	// }
	// myvar()