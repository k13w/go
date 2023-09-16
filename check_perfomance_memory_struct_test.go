package main

import (
	"testing"
)

func passStruckByReference(s *Souls) {
	s.Age++
}

func passStruckByValue(s Souls) {
	s.Age++
}

func BenchmarkStructPassByReference(b *testing.B) {
	soul := Souls{
		Name:        "Player",
		Age:         30,
		Experience:  100,
		CurrentBoss: "Boss 1",
		MainWeapon:  "Sword",
	}
	for i := 0; i < b.N; i++ {
		passStruckByReference(&soul)
	}
}

func BenchmarkStructPassByValue(b *testing.B) {
	soul := Souls{
		Name:        "Player",
		Age:         30,
		Experience:  100,
		CurrentBoss: "Boss 1",
		MainWeapon:  "Sword",
	}
	for i := 0; i < b.N; i++ {
		passStruckByValue(soul)
	}
}

func checkPerformanceMemoryStruct() {
	testing.Benchmark(BenchmarkStructPassByValue)
	testing.Benchmark(BenchmarkStructPassByReference)
}

type Souls struct {
	Name        string
	Age         int
	Experience  int
	CurrentBoss string
	MainWeapon  string
}
