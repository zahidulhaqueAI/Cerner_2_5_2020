// cerner_2^5_2020

import (
	"golang.org/x/tour/wc"
	"strings"
	"fmt"
)

func WordCount(s string) map[string]int {
	// splits the string
	x := strings.Fields(s)
	// create the map
	mymap := make(map[string]int)
	for index,value := range x{
		mymap[value]++
		}
	// return the populated map
	return mymap
}
