package main  
  
import (  
	"testing"  
	"github.com/kamihama-railway/uwasa"  
)  

  
func Benchmark_uwasa_basic(b *testing.B) {  
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
  
func Benchmark_uwasa_string_ops(b *testing.B) {  
	// Note: uwasa doesn't have startsWith, so we'll use substring comparison  
	params := map[string]interface{}{  
		"name":  "/groups/foo/bar",  
		"group": "foo",  
	}  
  
	engine, err := uwasa.NewEngine(`name == "/groups/" + group + "/bar"`)  
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
  
func Benchmark_uwasa_assignment(b *testing.B) {  
	// Note: uwasa doesn't support custom functions, but supports assignments  
	engine, err := uwasa.NewEngine(`result = "hello" + ", world"`)  
	if err != nil {  
		b.Fatal(err)  
	}  
  
	var out interface{}  
  
	b.ResetTimer()  
	for n := 0; n < b.N; n++ {  
		vars := map[string]interface{}{}  
		out, err = engine.Execute(vars)  
	}  
	b.StopTimer()  
  
	if err != nil {  
		b.Fatal(err)  
	}  
	if out.(string) != "hello, world" {  
		b.Fail()  
	}  
}  
  
func Benchmark_uwasa_arithmetic(b *testing.B) {  
	// Note: uwasa doesn't have map function, so we'll use basic arithmetic  
	env := map[string]interface{}{  
		"value": 1,  
	}  
	engine, err := uwasa.NewEngine(`value * 2`)  
	if err != nil {  
		b.Fatal(err)  
	}  
  
	var out interface{}  
  
	b.ResetTimer()  
	for n := 0; n < b.N; n++ {  
		out, err = engine.Execute(env)  
	}  
	b.StopTimer()  
  
	if err != nil {  
		b.Fatal(err)  
	}  
	if out.(float64) != 2.0 {  
		b.Fail()  
	}  
}
