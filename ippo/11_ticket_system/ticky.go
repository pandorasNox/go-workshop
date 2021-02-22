package main

import (
	"encoding/json"
	"io/ioutil"
	"os/user"
	"strconv"

	cli "github.com/urfave/cli/v2"
	// "ticky/pkg/ohmy"
	"fmt"
	"log"
	"os"
)

type Metadata struct {
	IdCounter       uint32
	AmountOfTickets int
	Tickets         []Ticket
}

type Ticket struct {
	ID          uint32 `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

func main() {
	//	jtest()
	//	return

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
						return nil
					}

					_, err = os.Create(usr.HomeDir + "/.ticky" + "/.metadata.json")
					if err != nil {
						log.Fatal(err)
					}
					initMetadata := Metadata{
						IdCounter:       0,
						AmountOfTickets: 0,
						Tickets:         []Ticket{},
					}
					metadataFile, err := json.MarshalIndent(initMetadata, "", " ")
					if err != nil {
						log.Fatalf("failed marshaling metadata %s", err)
					}
					err = ioutil.WriteFile(usr.HomeDir+"/.ticky"+"/.metadata.json", metadataFile, 0644)
					if err != nil {
						log.Fatalf("failed writing metadata file %s", err)
					}
					return nil
				},
			},
			{
				Name:  "create",
				Usage: "create a ticket with a title and a description, e.g. `ticky create \"my title\" \"my description\"`",
				Action: func(c *cli.Context) error {
					fmt.Println("createing ticket...")
					usr, err := user.Current()
					if err != nil {
						// log.Fatal( err )
						return err
					}
					// check metadata exist + message if not
					jsonFile, err := os.Open(usr.HomeDir + "/.ticky" + "/.metadata.json")
					if err != nil {
						return fmt.Errorf("please try to run ticky init: %s", err)
					}
					defer jsonFile.Close()

					// handle arguments
					//fmt.Printf("Hello %v \n", c.NArg())
					//https://godoc.org/gopkg.in/urfave/cli.v2#Context.NArg
					if c.NArg() != 2 {
						return fmt.Errorf("ticky create command accepts 2 args")
					}
					// fmt.Println("arg0", c.Args().Get(0), "arg1", c.Args().Get(1))
					// expect two args, if not throw message pls add 2 args or add...main
					// take the two args and map them to title and description var

					// read metadata.json into metadata golang struct
					readedMetadata := Metadata{}
					rawFileContentAsByteSlice, _ := ioutil.ReadAll(jsonFile)
					err = json.Unmarshal(rawFileContentAsByteSlice, &readedMetadata)
					if err != nil {
						return fmt.Errorf("Cannot parse metadata file: %s", err)
					}
					// fmt.Println("readed file: ", readedMetadata)

					// print them back to stdout

					//// error handling parse
					// increment IdCounter _> get the newest tickt id for the ticket we want to create in a second
					//// keep the IdCOunter in memory => in a varibale
					incrementedIDCount := readedMetadata.IdCounter + 1
					readedMetadata.IdCounter = incrementedIDCount
					readedMetadata.AmountOfTickets = readedMetadata.AmountOfTickets + 1

					// fmt.Println("incrementedIdCount", incrementedIDCount)
					// create the ticket variable instance with incremented IdCOunter variable with title and description
					newTicket := Ticket{
						ID:          incrementedIDCount,
						Title:       c.Args().Get(0),
						Description: c.Args().Get(1),
					}
					readedMetadata.Tickets = append(readedMetadata.Tickets, newTicket)

					fmt.Println("readedMetadata", readedMetadata)

					// write new ticket {write metadata file and new ticket because is the same file}
					metadataFile, err := json.MarshalIndent(readedMetadata, "", " ")
					if err != nil {
						return fmt.Errorf("Cannot parse metadata to json: %s", err)
					}
					err = ioutil.WriteFile(usr.HomeDir+"/.ticky"+"/.metadata.json", metadataFile, 0644)
					if err != nil {
						return fmt.Errorf("Cannot write to metadata file: %s", err)
					}
					return nil
				},
			},
			{
				Name:  "list",
				Usage: "list all tickets with title, e.g. `ticky list`",
				Action: func(c *cli.Context) error {
					//fmt.Println("listing tickets...")
					usr, err := user.Current()
					if err != nil {
						// log.Fatal( err )
						return err
					}

					jsonFile, err := os.Open(usr.HomeDir + "/.ticky" + "/.metadata.json")
					if err != nil {
						return fmt.Errorf("please try to run ticky init: %s", err)
					}
					defer jsonFile.Close()

					if c.NArg() != 0 {
						return fmt.Errorf("ticky list command does not support any args")
					}

					readedMetadata := Metadata{}
					rawFileContentAsByteSlice, _ := ioutil.ReadAll(jsonFile)
					err = json.Unmarshal(rawFileContentAsByteSlice, &readedMetadata)
					if err != nil {
						return fmt.Errorf("Cannot parse metadata file: %s", err)
					}

					fmt.Println("id title")

					for _, actualCurrentElementOfThisIteration := range readedMetadata.Tickets {
						//	fmt.Println(actualCurrentElementOfThisIteration)
						fmt.Println(actualCurrentElementOfThisIteration.ID, actualCurrentElementOfThisIteration.Title)
					}

					return nil
				},

				// read .metada.json file, extract id and title info, stdout
			},
			{
				Name:  "show",
				Usage: "shows all details of a specific ticket given by id, e.g. `ticky show <id>`",
				Action: func(c *cli.Context) error {
					//fmt.Println("listing tickets...")
					usr, err := user.Current()
					if err != nil {
						// log.Fatal( err )
						return err
					}

					if c.NArg() != 1 {
						return fmt.Errorf("ticky show command supports exact one args")
					}
					//convert maybeSearchID to uint32
					givenCLIArg := c.Args().Get(0)
					maybeSearchID64, _ := strconv.ParseUint(givenCLIArg, 10, 64)
					maybeSearchID := uint32(maybeSearchID64)
					//fmt.Printf("%T\n", maybeSearchID)

					jsonFile, err := os.Open(usr.HomeDir + "/.ticky" + "/.metadata.json")
					if err != nil {
						return fmt.Errorf("please try to run ticky init: %s", err)
					}
					defer jsonFile.Close()

					readedMetadata := Metadata{}
					rawFileContentAsByteSlice, _ := ioutil.ReadAll(jsonFile)
					err = json.Unmarshal(rawFileContentAsByteSlice, &readedMetadata)
					if err != nil {
						return fmt.Errorf("Cannot parse metadata file: %s", err)
					}
					found := false
					for _, currentTicket := range readedMetadata.Tickets {
						if currentTicket.ID == maybeSearchID {
							found = true
							fmt.Println(currentTicket.Title, currentTicket.Description)
						}
					}
					if found == false {
						fmt.Println("This ticket does not exist")
					}
					return nil
				},
			},
			{
				Name:  "delete",
				Usage: "delete the ticket given by id, e.g. `ticky delete <id>`",
				Action: func(c *cli.Context) error {
					usr, err := user.Current()
					if err != nil {
						// log.Fatal( err )
						return err
					}
					if c.NArg() != 1 {
						return fmt.Errorf("ticky delete command supports exact one args")
					}
					//convert maybeSearchID to uint32
					givenCLIArg := c.Args().Get(0)
					maybeSearchID64, _ := strconv.ParseUint(givenCLIArg, 10, 64)
					maybeSearchID := uint32(maybeSearchID64)

					jsonFile, err := os.Open(usr.HomeDir + "/.ticky" + "/.metadata.json")
					if err != nil {
						return fmt.Errorf("please try to run ticky init: %s", err)
					}
					defer jsonFile.Close()

					readedMetadata := Metadata{}
					rawFileContentAsByteSlice, _ := ioutil.ReadAll(jsonFile)
					err = json.Unmarshal(rawFileContentAsByteSlice, &readedMetadata)
					if err != nil {
						return fmt.Errorf("Cannot parse metadata file: %s", err)
					}
					// "1st WAY""
					//found := false
					//for _, currentTicket := range readedMetadata.Tickets {
					//	indexIDToBeRemoved := findIndexForTicketID(readedMetadata.Tickets, currentTicket.ID)
					//	if findIndexForTicketID(readedMetadata.Tickets, currentTicket.ID) == findIndexForTicketID(readedMetadata.Tickets, maybeSearchID) {
					//		found = true
					//		readedMetadata.Tickets = remove(readedMetadata.Tickets, indexIDToBeRemoved)
					//		//fmt.Println(readedMetadata.Tickets)
					//	}
					//}
					//if found == false {
					//	fmt.Println("This ticket does not exist")
					//}
					// "2nd WAY"
					indexForSearchedTicketID := findIndexForTicketID(readedMetadata.Tickets, maybeSearchID)
					if indexForSearchedTicketID == -1 {
						return fmt.Errorf("This ticket does not exist %s", "")
					}
					newTicketList := remove(readedMetadata.Tickets, indexForSearchedTicketID)
					readedMetadata.Tickets = newTicketList
					// write new ticket {write metadata file and new ticket because is the same file}
					metadataFile, err := json.MarshalIndent(readedMetadata, "", " ")
					if err != nil {
						return fmt.Errorf("Cannot parse metadata to json: %s", err)
					}
					err = ioutil.WriteFile(usr.HomeDir+"/.ticky"+"/.metadata.json", metadataFile, 0644)
					if err != nil {
						return fmt.Errorf("Cannot write to metadata file: %s", err)
					}
					return nil
				},
			},
			{
				Name:  "update",
				Usage: "update the ticket given by id and new title and description, e.g. `ticky update <id> \"New Title\" \"New Description\"` \n Notice: Use tickey show <id> to prevent updating wrong ticket",
				Action: func(c *cli.Context) error {
					usr, err := user.Current()
					if err != nil {
						return err
					}

					if c.NArg() != 3 {
						return fmt.Errorf("ticky update command supports three args")
					}
					//convert maybeSearchID to uint32
					givenCLIArg := c.Args().Get(0)
					maybeSearchID64, _ := strconv.ParseUint(givenCLIArg, 10, 64)
					maybeSearchID := uint32(maybeSearchID64)

					newTitle := c.Args().Get(1)
					newDescription := c.Args().Get(2)

					jsonFile, err := os.Open(usr.HomeDir + "/.ticky" + "/.metadata.json")
					if err != nil {
						return fmt.Errorf("please try to run ticky init: %s", err)
					}
					defer jsonFile.Close()

					readedMetadata := Metadata{}
					rawFileContentAsByteSlice, _ := ioutil.ReadAll(jsonFile)
					err = json.Unmarshal(rawFileContentAsByteSlice, &readedMetadata)
					if err != nil {
						return fmt.Errorf("Cannot parse metadata file: %s", err)
					}

					indexForSearchedTicketID := findIndexForTicketID(readedMetadata.Tickets, maybeSearchID)
					if indexForSearchedTicketID == -1 {
						return fmt.Errorf("This ticket does not exist %s", "")
					}
					readedMetadata.Tickets[indexForSearchedTicketID].Title = newTitle
					readedMetadata.Tickets[indexForSearchedTicketID].Description = newDescription

					// write new ticket {write metadata file and new ticket because is the same file}
					metadataFile, err := json.MarshalIndent(readedMetadata, "", " ")
					if err != nil {
						return fmt.Errorf("Cannot parse metadata to json: %s", err)
					}
					err = ioutil.WriteFile(usr.HomeDir+"/.ticky"+"/.metadata.json", metadataFile, 0644)
					if err != nil {
						return fmt.Errorf("Cannot write to metadata file: %s", err)
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

func remove(tickets []Ticket, indexOfTicketIDToBeRemoved int) []Ticket { //change uint32 to int!!!
	return append(tickets[:indexOfTicketIDToBeRemoved], tickets[indexOfTicketIDToBeRemoved+1:]...) //... spread operator
}

// from the list of ticket we want to get the index of the ticket we are looking for. if we don't find it return -1
func findIndexForTicketID(tickets []Ticket, ticketID uint32) (index int) {
	for index, element := range tickets {
		if element.ID == ticketID {
			return int(index) //type casting
		}
	}
	return -1
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
	defer jsonFile.Close() // => puts it on defer stack, the defer stack runs after the current function finished
	fmt.Println("Successfully Opened test.json")
	readedTicket := Ticket{}
	byteValue, _ := ioutil.ReadAll(jsonFile)
	err = json.Unmarshal(byteValue, &readedTicket)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("readed file: ", readedTicket)
	// jsonFile.Close()

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

// .metadata.json:
// write the input of the cmd line
