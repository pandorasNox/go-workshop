given two strings with the same length calculate the hamming distance, which describes the amount of differences (amount of different charackters) they have to each other

e.g.
in1 in2 outInt,outErr
“ABC” “ABC” 0
“ABC” “AAA” 2
“AB” “ABC” 0 (default zero value of int), error
“ABC” “AB” 0 (default zero value of int), error
“A” “ABÜ” 0 (default zero value of int), error
“ABC” “ABB” 1
“ABC” “ABD” 1
“ABC” “ABÜ” 1, nil
“” “” 0

## testing
go test .
go test main.go main_test.go

## info
- https://en.wikipedia.org/wiki/Hamming_distance
