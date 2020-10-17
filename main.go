package main

import (
	"fmt"
	"target_finder/slicegenerator"
	"target_finder/targetfinder"
	"testing"
)

func benchmarkFindValueInSliceBinary(value int, slice []int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		targetfinder.FindValueInSliceBinary(value, slice)
	}
}

func benchmarkFindValueInSliceLinear(value int, slice []int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		targetfinder.FindValueInSliceLinear(value, slice)
	}
}

func BenchmarkFindValueInSliceBinary10(b *testing.B) {
	value, slice := slicegenerator.GenerateRandomOrderedSlice(10, 1000)
	benchmarkFindValueInSliceBinary(value, slice, b)
}

func BenchmarkFindValueInSliceBinary100(b *testing.B) {
	value, slice := slicegenerator.GenerateRandomOrderedSlice(100, 1000)
	benchmarkFindValueInSliceBinary(value, slice, b)
}

func BenchmarkFindValueInSliceBinary1000(b *testing.B) {
	value, slice := slicegenerator.GenerateRandomOrderedSlice(1000, 1000)
	benchmarkFindValueInSliceBinary(value, slice, b)
}

func BenchmarkFindValueInSliceBinary10000(b *testing.B) {
	value, slice := slicegenerator.GenerateRandomOrderedSlice(10000, 1000)
	benchmarkFindValueInSliceBinary(value, slice, b)
}

func BenchmarkFindValueInSliceBinary100000(b *testing.B) {
	value, slice := slicegenerator.GenerateRandomOrderedSlice(100000, 1000)
	benchmarkFindValueInSliceBinary(value, slice, b)
}

func BenchmarkFindValueInSliceBinary1000000(b *testing.B) {
	value, slice := slicegenerator.GenerateRandomOrderedSlice(1000000, 1000)
	benchmarkFindValueInSliceBinary(value, slice, b)
}

func BenchmarkFindValueInSliceBinary10000000(b *testing.B) {
	value, slice := slicegenerator.GenerateRandomOrderedSlice(10000000, 1000)
	benchmarkFindValueInSliceBinary(value, slice, b)
}

func BenchmarkFindValueInSliceLinear10(b *testing.B) {
	value, slice := slicegenerator.GenerateRandomOrderedSlice(10, 1000)
	benchmarkFindValueInSliceLinear(value, slice, b)
}

func BenchmarkFindValueInSliceLinear100(b *testing.B) {
	value, slice := slicegenerator.GenerateRandomOrderedSlice(100, 1000)
	benchmarkFindValueInSliceLinear(value, slice, b)
}

func BenchmarkFindValueInSliceLinear1000(b *testing.B) {
	value, slice := slicegenerator.GenerateRandomOrderedSlice(1000, 1000)
	benchmarkFindValueInSliceLinear(value, slice, b)
}

func BenchmarkFindValueInSliceLinear10000(b *testing.B) {
	value, slice := slicegenerator.GenerateRandomOrderedSlice(10000, 1000)
	benchmarkFindValueInSliceLinear(value, slice, b)
}

func BenchmarkFindValueInSliceLinear100000(b *testing.B) {
	value, slice := slicegenerator.GenerateRandomOrderedSlice(100000, 1000)
	benchmarkFindValueInSliceLinear(value, slice, b)
}

func BenchmarkFindValueInSliceLinear1000000(b *testing.B) {
	value, slice := slicegenerator.GenerateRandomOrderedSlice(1000000, 1000)
	benchmarkFindValueInSliceLinear(value, slice, b)
}

func BenchmarkFindValueInSliceLinear10000000(b *testing.B) {
	value, slice := slicegenerator.GenerateRandomOrderedSlice(10000000, 1000)
	benchmarkFindValueInSliceLinear(value, slice, b)
}

func main() {
	fmt.Println("BenchmarkFindValueInSliceBinary10: ", testing.Benchmark(BenchmarkFindValueInSliceBinary10))
	fmt.Println("BenchmarkFindValueInSliceBinary100: ", testing.Benchmark(BenchmarkFindValueInSliceBinary100))
	fmt.Println("BenchmarkFindValueInSliceBinary1000: ", testing.Benchmark(BenchmarkFindValueInSliceBinary1000))
	fmt.Println("BenchmarkFindValueInSliceBinary10000: ", testing.Benchmark(BenchmarkFindValueInSliceBinary10000))
	fmt.Println("BenchmarkFindValueInSliceBinary100000: ", testing.Benchmark(BenchmarkFindValueInSliceBinary100000))
	fmt.Println("BenchmarkFindValueInSliceBinary1000000: ", testing.Benchmark(BenchmarkFindValueInSliceBinary1000000))
	fmt.Println("BenchmarkFindValueInSliceBinary10000000: ", testing.Benchmark(BenchmarkFindValueInSliceBinary10000000))
	fmt.Println("BenchmarkFindValueInSliceLinear10: ", testing.Benchmark(BenchmarkFindValueInSliceLinear10))
	fmt.Println("BenchmarkFindValueInSliceLinear100: ", testing.Benchmark(BenchmarkFindValueInSliceLinear100))
	fmt.Println("BenchmarkFindValueInSliceLinear1000: ", testing.Benchmark(BenchmarkFindValueInSliceLinear1000))
	fmt.Println("BenchmarkFindValueInSliceLinear10000: ", testing.Benchmark(BenchmarkFindValueInSliceLinear10000))
	fmt.Println("BenchmarkFindValueInSliceLinear100000: ", testing.Benchmark(BenchmarkFindValueInSliceLinear100000))
	fmt.Println("BenchmarkFindValueInSliceLinear1000000: ", testing.Benchmark(BenchmarkFindValueInSliceLinear1000000))
	fmt.Println("BenchmarkFindValueInSliceLinear10000000: ", testing.Benchmark(BenchmarkFindValueInSliceLinear10000000))
}
