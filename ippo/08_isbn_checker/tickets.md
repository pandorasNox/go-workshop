
___

## 1 Task
As a (skilled) user I want to be able to pass an ISBN number via command line to a programm,
which returns the same ISBN and if it's a valid ISBN-10.

Acceptance criteria
- can run it via cli
- can pass an ISBN string
- returns to std.out the isbn and if it's valid accoring to ISBN-10

example:
```
$ go run main.go "foobarbaz"

the given input 'foobarbaz' is invalid



$ go run main.go "3-598-21508-8"

the given input '3-598-21508-8' is valid
```

### questions
- What is the first thing to do?

### 1.0 subtask  (√)
- √

### 1.1 subtask  (√)
- write a test

### 1.2 subtask  (√)
- write a function which accepts as input a string
- returns true for valid ISBN-10
- your test is passing

### 1.3 subtask
- add ability to pass a string from cli to the programm
- programm returns the passed string and if this string is a valid ISBN
```
$ go run . "123456789"

The input '123456789' is valid ISBN: true
```

### 1.4 subtask
- allow cli prog to accept csv file as input
- outputs a csv file with mapping isbn to if its valid
e.g.

inputFile:
```
123456789
123

3-598-21508-8
```

cmd:
```
$ go run . -f "isbn.csv"

123456789, true
123, false
, false
3-598-21508-8, true
```

### 1.5 subtask
- build it as a binary
- execute it like `isbnchecker ""`








### hints
- https://web.archive.org/web/20130522043458/http://www.isbn.org/standards/home/isbn/international/html/usm4.htm
- see https://gobyexample.com/command-line-arguments

___



















note:
task 1, work with test
task 2, works on cli, with cli input like `$ validisb "1-24325435-3435-X"`
		+ nice output "ISBN '' is (NOT) valid"
task 3 read csv file, print it out
task 4 read csv file, expect certain format of the file, return true/false
		for file validity
task 5 read valid csv file, valdate all isbns, print csv file
		with valid column in output