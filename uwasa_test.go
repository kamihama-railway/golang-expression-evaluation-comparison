package main  
  
import (  
	"testing"  
	"github.com/kamihama-railway/uwasa"  
)  

  
func Benchmark_uwasa_exc(b *testing.B) {  
	params := createParams()  
  
	engine, err := uwasa.NewEngine(example)  
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

