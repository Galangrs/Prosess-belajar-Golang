package main


import (
	"fmt"
)

var (
	data = []map[string]string{
		{ "name" : "Hank", 			"Age": "50", 	"Job" : "Polisi" },
		{ "name" : "Heisenberg", 	"Age" : "52", 	"Job" : "Ilmuwan" },
		{ "name" : "Skyler", 		"Age" : "48", 	"Job" : "Akuntan" },
	}
	row1 int8 = 5
	row3 int8 = 5
	i int8
	j int8
)

func main(){
	// Nomer 1
	for _ , value := range data {
		fmt.Printf("Hi Perkenalkan, Nama saya %s, umur saya %s, dan saya bekerja sebagai %s\n",value["name"],value["Age"],value["Job"])
	}
	// Nomer 2
	i = 0
	start_row1:
		if row1 == i {
			goto end_row1
		}
		i++
		fmt.Println("*")
		goto start_row1
	end_row1:
		i =	0
	// Nomer 3
	j = 0
	start_row3:
		if row3 == i {
			goto end_row3
		}
		i++
		start_row2:
			if j == i {
				goto end_row2
			}
			j++
			fmt.Printf("*")
			goto start_row2
		end_row2:
			j = 0
			fmt.Printf("\n")
		goto start_row3
	end_row3:
		i =	0
		j = 0
}