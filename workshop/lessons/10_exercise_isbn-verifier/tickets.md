
___

## 1 Task
As a (skilled) user I want to be able to pass an ISBN number via command line to a programm,
which returns the same ISBN and if it's a valid ISBN-10.

Acceptance criteria
- can run it via cli
- can pass an ISBN string
- returns to std.in the isbn and if it's valid accoring to ISBN-10

### questions
- What is the first thing to do?

### 1.0 subtask
- ???

### 1.1 subtask
- write a test

### 1.2 subtask
- write a function which accepts as input a string
- returns true for valid ISBN-10
- your test is passing

### 1.2 subtask
- add ability to pass a string from cli to the programm
- programm returns the passed string and if this string is a valid ISBN











### hints
- //https://web.archive.org/web/20130522043458/http://www.isbn.org/standards/home/isbn/international/html/usm4.htm
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