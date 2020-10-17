# Bradfield Takehome - Exercise 2

## System setup 
This solution is written in Golang. 

To prepare your system to run this solution: 

1. Verify you have Golang installed: `go --version`
2. If not installed: `brew install golang`  
3. Set your GOPATH: 
	`export GOPATH=<insert-path-to-directory-where-you-will-save-go-projects>`
4. This project should be saved in the GOPATH directory

## Running the program
There are two ways to benchmark the functions. 
1. IDE via *_test.go files 
2. Command line
    1. build the project `go build`
    2. run main `go run main.go`

## Notes
* For benchmarking, I arbitrarily select the maximum value for values in the generate array to be 1000 

## Analysis 
Detailed analysis of benchmarking in [Google Doc](https://docs.google.com/document/d/1F-eGgUg6_c0UZuk0HssdgZLyPe9k1zeEeCdCs-tGaWA/edit?usp=sharing)

## Future work / shortcomings 
* Create TargetFinder interface
* Generalize for different data types using reflection 
* Run benchmarking in the main function
* Modify benchTime for main function benchmarking calls 
* Further Benchmark research to determine if: 
    * Input & benchmark names can be changed interatively
    * A single slice can be shared by two Benchmarking functions 


