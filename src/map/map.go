package main

import (
	"fmt"
)

func main(){
	var countryCapitalMap map[string]string

	countryCapitalMap = make(map[string]string)

	//insert key-value
	countryCapitalMap["France"] = "Paris"
	countryCapitalMap["Italy"] = "Rome"
	countryCapitalMap["Japan"] = "Tokyo"
	countryCapitalMap["India"] = "New Delhi"

	for c := range countryCapitalMap{
		fmt.Println("Capital of", c, "is", countryCapitalMap[c])
	}

	//check for a data if exist
	capital, ok := countryCapitalMap["United States"]
	if(ok){
		fmt.Println("Capital of United States is", countryCapitalMap[capital])
	} else{
		fmt.Println("Capital of United States is not present")
	}

	//Delete functions
	delete(countryCapitalMap, "France")
	fmt.Println("After delete element:")
	for c := range countryCapitalMap{
		fmt.Println("Capital of", c, "is", countryCapitalMap[c])
	}
}
