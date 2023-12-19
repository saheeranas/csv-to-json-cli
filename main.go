package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Item map[string]interface{}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {    
	if len(os.Args) < 2 {
		fmt.Println("Please run on a file")
		os.Exit(1)
	} 

	file, err := os.Open(os.Args[1])
	check(err) 
	defer file.Close()

	var res []Item
	count := 0

	reader := bufio.NewReader(file) 

	for  {
		line, _, err := reader.ReadLine()
		if err != nil {
			break 
		}
		
		slices := strings.Split(string(line), ",") 

		temp := Item{}
		for i, v := range slices {  
			temp[strconv.Itoa(i)] = v
		} 

		res = append(res, temp)  
		count++
	}

	
	jsonData, err := json.Marshal(res)
	check(err)

	fmt.Printf("%+v\n", string(jsonData))
}