package main

import (
	"testing"

	"github.com/kamihama-railway/uwasa"
)

func Benchmark_uwasa_Basic(b *testing.B) {
	params := createParams()

	engine, err := uwasa.NewEngineWithOptions(example, uwasa.EngineOptions{OptimizationLevel: uwasa.OptBasic})
	if err != nil {
		b.Fatal(err)
	}

	var out interface{}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		out, err = engine.Execute(params)
	}
	b.StopTimer()

	if err != nil {
		b.Fatal(err)
	}
	if !out.(bool) {
		b.Fail()
	}
}

func Benchmark_uwasa_None(b *testing.B) {
	params := createParams()

	engine, err := uwasa.NewEngineWithOptions(example, uwasa.EngineOptions{OptimizationLevel: uwasa.OptNone})
	if err != nil {
		b.Fatal(err)
	}

	var out interface{}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		out, err = engine.Execute(params)
	}
	b.StopTimer()

	if err != nil {
		b.Fatal(err)
	}
	if !out.(bool) {
		b.Fail()
	}
}

func Benchmark_uwasa_Recompiler(b *testing.B) {
	params := createParams()

	engine, err := uwasa.NewEngineWithOptions(example, uwasa.EngineOptions{OptimizationLevel: uwasa.OptBasic, UseRecompiler: true})
	if err != nil {
		b.Fatal(err)
	}

	var out interface{}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		out, err = engine.Execute(params)
	}
	b.StopTimer()

	if err != nil {
		b.Fatal(err)
	}
	if !out.(bool) {
		b.Fail()
	}
}

func Benchmark_uwasa_VM(b *testing.B) {
	params := createParams()

	engine, err := uwasa.NewEngineVM(example)
	if err != nil {
		b.Fatal(err)
	}

	var out interface{}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		out, err = engine.Execute(params)
	}
	b.StopTimer()

	if err != nil {
		b.Fatal(err)
	}
	if !out.(bool) {
		b.Fail()
	}
}
