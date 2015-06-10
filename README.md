#ei [![GoDoc](https://godoc.org/github.com/jaracil/ei?status.png)](https://godoc.org/github.com/jaracil/ei)
Painless empty interface type conversion.

Download:
```shell
go get github.com/jaracil/ei
```

##Description

ei package allows easy type conversion between empty interfaces and basic go types.

Supported types are:
 * int8, int16, int32, int64, int
 * uin8, byte, uint16, uint32, uint64, uint
 * float32, float64
 * string, []byte
 * time.Time
 * map[string]interface{} or shorthand type ei.M
 * []interface{} or shorthand type ei.S
 * interface{} pointing to one of above types

###Examples

```go
package main

import(
	"github.com/jaracil/ei"
)

func main(){
	var i interface{}
	
	// float to int (zero value on error)
	i = 5.2
	println(ei.N(i).IntZ()) // 5
	
	// string to int64 (zero value on error)
	i = "5"
	println(ei.N(i).Int64Z()) // 5
	
	// map[string]interface{} access. Using shorthand ei.M type
	i = ei.M{"sota":10, "caballo":"11"}
	println(ei.N(i).M("sota").IntZ(), ei.N(i).M("caballo").IntZ()) // 10 11
	l, _ := ei.N(i).Len()
	println(l) // 2 

	// []interface{} access. Using shorthand ei.S type
	i = ei.S{10, 11, 12}
	println(ei.N(i).S(0).IntZ(), ei.N(i).S(1).StringZ()) // 10 11
	l, _ = ei.N(i).Len()
	println(l) // 3
	
	// Nested access
	i = ei.M{"red":ei.S{255,0,0}, "green":ei.S{0,255,0}, "blue":ei.S{0,0,255}}
	println(ei.N(i).M("green").S(1).ByteZ()) // 255
	
	// Without shorthand types
	i = map[string]interface{}{"red":[]interface{}{255,0,0}, "green":[]interface{}{0,255,0}, "blue":[]interface{}{0,0,255}}
	println(ei.N(i).M("green").S(1).IntZ()) // 255
	
	// Errors are propagated 
	i = ei.M{"red":ei.S{255,0,0}, "green":ei.S{0,255,0}, "blue":ei.S{0,0,255}}
	v, err := ei.N(i).M("brown").S(1).Int() // Key error on M("brown") is propagated up to Int()
	println(v, err.Error()) // 0 key not found

}

```


