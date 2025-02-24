package csvtojson
import (
	"bufio"
	"fmt"
	"os" 
	"strings"
	"encoding/json" 
)

type Item map[string]interface{}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func Convert() {    
	if len(os.Args) < 2 {
		fmt.Println("Syntax:\n\t csvtojsoncli <file.csv>\n\t or \n\t go run main.go <file.csv>")
		os.Exit(1)
	} 

	file, err := os.Open(os.Args[1])
	check(err) 
	defer file.Close()

	var result []Item 

	reader := bufio.NewReader(file) 

	header_fields_line, _, err := reader.ReadLine()
	if err != nil {
		fmt.Println(err)
	}

	header_fields_slices := strings.Split(string(header_fields_line), ",")  

	for  {
		line, _, err := reader.ReadLine()
		if err != nil {
			break 
		}
		
		slices := strings.Split(string(line), ",") 

		
		temp := Item{}
		for i, v := range slices {   
			if len(header_fields_slices) > i {
				fmt.Printf("%+v\n", header_fields_slices[i])
				temp[header_fields_slices[i]] = v
			}
		}  
			
		result = append(result, temp)   
	}  
	
	jsonData, err := json.Marshal(result)
	check(err)

	fmt.Printf("%+v\n", string(jsonData))
}