# Bradfield Takehome - Exercise 2 - Range count

## About
This project implements two different approaches (linear scan, double binary
search) to finding the number of occurrences of a given value in a sorted
array.

## System setup 
This project is written in Golang, version 1.15. 

To prepare your system to run this project: 

1. Verify you have Golang installed: `go --version`
2. If not, install: `brew install golang`
3. Verify your GOPATH environment variable is set: `echo $GOPATH`
4. If not, set: `export GOPATH=<insert-path-to-directory-where-you-will-save-go-projects>`  
5. This project should be saved in the GOPATH

## Running the program
There are three ways to benchmark the linear scan and double binary search
functions. Each approach assumes you have already completed the **System Setup**
above. 
1. IDE (e.g. GoLand) via `*_test.go` files 
2. Command line via `_*test.go` files:
    1. `cd $GOPATH/src/bradfield_takehome_exercise2/targetfinder`
    2. `go test -bench=. -benchtime=10s`
2. Command line via main() function:
    1. `cd $GOPATH/src/bradfield_takehome_exercise2`
    2. build the project `go build`
    3. run main `go run main.go`

## Notes
For benchmarking, I arbitrarily set the maximum value of the randomly generated
integers in test arrays to be 1000.  

## Analysis 
Detailed analysis of benchmarking in [Google sheet](https://docs.google.com/document/d/1F-eGgUg6_c0UZuk0HssdgZLyPe9k1zeEeCdCs-tGaWA/edit?usp=sharing).

## Future work / shortcomings 
* Generalize functions for different data types using 
[reflection](https://golang.org/pkg/reflect/)
* Create TargetFinder interface, implemented by both linear scan and double
binary search  
* Modify benchTime for main function benchmarking (currently the default is 1s
when run via main vs. 10s when run via tests) 
* Continue reading Golang benchmarking docs to determine if: 
    * Input & benchmark function names can be dynamically generated 
    * A single input value can be shared by separate benchmarking functions 

## Helpful resources 
* [Go official docs - Package testing](https://golang.org/pkg/testing/)
