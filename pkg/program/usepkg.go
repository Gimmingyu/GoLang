package main

import (
	"GOLANG/pkg/custompkg"
	"fmt"

	"github.com/guptarohit/asciigraph"
	"github.com/tuckersGo/musthaveGo/ch16/expkg"
)

func main() {
	pl := fmt.Println
	custompkg.PrintCustom()
	custompkg.printcustom2()
	expkg.PrintSample()

	data := []float64{3, 4, 5, 6, 7, 8, 9, 10, 11, 2, 7, 2, 5, 6}
	graph := asciigraph.Plot(data)
	pl(graph)
}
