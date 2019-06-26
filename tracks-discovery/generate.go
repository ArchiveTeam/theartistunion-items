package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println("Generating items files for theartistunion tracks-discovery...")
	var maxID = int(0x00FFFFFF)
	var maxItemsPerFile = 10000
	var maxIDperRange = 100
	var nFiles = 1
	var i, iNext, items int

	f, err := os.OpenFile("theartiststation-items-tracks-discovery-"+strconv.Itoa(maxItemsPerFile), os.O_CREATE, 0664)
	if err != nil {
		fmt.Println("Unable to open the file:", err)
		os.Exit(1)
	}

	for _ = i; i < maxID; i = i + (maxIDperRange - 1) {
		iNext = i + (maxIDperRange - 1)
		if iNext > maxID {
			iNext = maxID
		}

		items++

		if (items % maxItemsPerFile) == 0 {
			f.Close()

			nFiles++
			f, err = os.OpenFile("theartiststation-items-tracks-discovery-"+strconv.Itoa(maxItemsPerFile*nFiles), os.O_CREATE, 0664)
			if err != nil {
				fmt.Println("Unable to open the file: ", err)
				os.Exit(1)
			}

		}

		_, err = f.WriteString("tracks-discovery:" + fmt.Sprintf("%06x", i) + "-" + fmt.Sprintf("%06x", iNext) + "\n")
		if err != nil {
			fmt.Println("Unable to write the file:", err)
			os.Exit(1)
		}
	}
	f.Close()
	fmt.Println("Done generating items files for theartistunion tracks-discovery")
	return
}
