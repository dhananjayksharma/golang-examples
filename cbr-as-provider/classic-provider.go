package main

import (
	"fmt"
	"golang-examples/cbr-as-provider/utils"
	"os"
)

type Data struct {
	Key   string
	Count int
}

func CountFreq(n string) {
	var dataUnique []Data
	makeUnique(&dataUnique, n)
	fmt.Printf("dataUnique for %v\n", dataUnique)

}

func makeUnique(dataUnique *[]Data, s string) {
	for _, row := range s {
		var skip bool = false
		searchStr := string(row)
		cnt := 1
		if len(*dataUnique) > 0 {
			for _, du := range *dataUnique {
				if du.Key == searchStr {
					cnt++
					skip = true
				}
			}
		}
		if !skip {
			*dataUnique = append(*dataUnique, Data{Key: searchStr, Count: cnt})
		} else {
			updateCount(dataUnique, searchStr, cnt)
		}
	}
}

func updateCount(dataUnique *[]Data, searchStr string, cnt int) {
	for i, d := range *dataUnique {
		if d.Key == searchStr {
			(*dataUnique)[i].Count = (*dataUnique)[i].Count + 1
		}
	}
}

// string of "main"
func main() {
	name := os.Args[1]
	currTime := utils.GetCurrentTime()
	cbrServiceZoneApiKey := os.Getenv("IBM_CLOUD_CBR_SERVICE_ZONE_API_KEY")
	fmt.Println("Running now:", currTime)
	fmt.Println("Running now ver 2:", currTime)
	fmt.Println("IBM_CLOUD_CBR_SERVICE_ZONE_API_KEY:", cbrServiceZoneApiKey[5:])
	fmt.Println("Running as classic provider params:", name)
	CountFreq(name)
}
