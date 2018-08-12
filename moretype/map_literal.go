package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex{
	"Bell Labs": Vertex{
		40.11111, -74.22222,
	},
	"Google": Vertex{
		37.33333, -122.44444,
	},
}

func main() {
	fmt.Println(m)
}
