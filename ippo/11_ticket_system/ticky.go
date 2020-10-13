package main

import (
	"encoding/json"
	"io/ioutil"
	"os/user"

	cli "github.com/urfave/cli/v2"
	// "ticky/pkg/ohmy"
	"fmt"
	"log"
	"os"
)

type Metadata struct {
	AmountOfTickets int
	Tickets         []Ticket
}
type Ticket struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

// .metadata.json:
// write the input of the cmd line

// $ ticky create "Monday tasks" "Go supermarket"
// .metadata.json:
// [
//     {
//         "id": 0,
//         "title": "init title",
//         "description": "init description"
//     },
//     {
//         "id": 1,
//         "title": "Monday tasks",
//         "description": "Go supermarket"
//     }
// ]

// {
//  amountOfTickets: 2
// 	tickets: [
// 		{
// 			"id": 0,
// 			"title": "init title",
// 			"description": "init description"
// 		},
// 		{
// 			"id": 1,
// 			"title": "Monday tasks",
// 			"description": "Go supermarket"
// 		}
// 	]
// }

func main() {
	jtest()
	return

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
				Action: func(c *cli.Context) error {
					usr, err := user.Current()
					if err != nil {
						// log.Fatal( err )
						return err
					}

					_, err = os.Stat(usr.HomeDir + "/.ticky")
					if os.IsNotExist(err) {
						log.Print("Folder `.ticky` does not exist, initialize filesystem")
						os.MkdirAll(usr.HomeDir+"/.ticky", os.FileMode(0755))
					} else {
						log.Println("Folder exists")
					}

					_, err = os.Create(usr.HomeDir + "/.ticky" + "/.metadata.json")
					if err != nil {
						log.Fatal(err)
					}
					return nil
				},
			},
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func jtest() {
	type Ticket struct {
		ID          int    `json:"id"`
		Title       string `json:"title"`
		Description string `json:"description"`
	}

	ticket0 := Ticket{
		ID:          1,
		Title:       "a title",
		Description: "a description",
	}

	// ticket0.write("file/path/test.json")
	// ticket0.read("file/path/test.json")

	// fmt.Println(ticket0)
	file, _ := json.MarshalIndent(ticket0, "", " ")
	_ = ioutil.WriteFile("test.json", file, 0644)

	//check open json file
	jsonFile, err := os.Open("test.json")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened test.json")
	readedTicket := Ticket{}
	byteValue, _ := ioutil.ReadAll(jsonFile)
	err = json.Unmarshal(byteValue, &readedTicket)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("readed file: ", readedTicket)
	defer jsonFile.Close()

	//-------------------------------------------------

	var jsonText = []byte(`[
	    {"id": 1, "title": "a title", "description": "description"}
	]`)
	var tickets []Ticket

	err = json.Unmarshal([]byte(jsonText), &tickets)
	if err != nil {
		fmt.Println(err)
	}

	tickets = append(tickets, Ticket{ID: 2, Title: "a new Title", Description: "a new description"})

	fmt.Println("tickets: ", tickets)

	file2, _ := json.MarshalIndent(tickets, "", " ")
	_ = ioutil.WriteFile("test2.json", file2, 0644)

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
