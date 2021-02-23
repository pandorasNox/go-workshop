# exercise cli ticket system

Write a command line tool which allows you to create, update, delete and list tickets.
A ticket itsel is just a textfile with description and id (id might be the name of the file). The id's are just intergers counting up, which means each new tickets gets a new higher id.

name of the tool: ticky

The cli tool expects in your home directory a `.ticky` directory, where it will store all the tickets as separate files

usage e.g. (the final result can look differently):
```
$ ticky
> how to use ticky ...
> ...
> ...

$ ticky -h
> how to use ticky ...
> ...
> ...

$ ticky init
> ...

$ ticky create "my title" "don't forget to drink water"
> ticket with id 1 created
> title: my title
> description:
> don't forget to drink water

$ ticky show 1
> showing ticket with id 1
> title: my title
> description:
> don't forget to drink water

$ ticky show 7239847
> ticket with id 7239847 not found: don't exist

$ ticky delete 1
> deleted ticket with id 1

$ ticky delete 738245237
> couldn't delete ticket, not exiting

$ ticky update 1 "new title" "don't drink to much alcohole"
> ticket with id 1 updated
> title: new title
> description:
> don't drink to much alcohole


$ ticky list
> id title
> 1 "my title"
> 2 "another title"
> ...

```

## Future implementaitions:
$ ticky destruct: this command will remove all traces of ticky
$ improve ticky create output, shows only new created ticket
$ improve ticky show output, more beautiful
$ ticky show command can show more than one ticket details.
$ ticky update command:
        a)shall first print out the original ticket title/description before modification (cancel option?)
        b)takes only id argument and listening to keyboard for input e.g.: Current Title: blabla
                                                                           New Title: ....
                                                                           Current description: bloubla
                                                                           New description: ....
