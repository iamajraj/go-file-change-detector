package main

import (
	"fmt"
	"os"
	"path"
	"time"
)

// var directory string
// var commandExec string

// func parseArgs(arguments []string) {

// }

func main() {
	// parseArgs(os.Args)

	directory := "detectme"

	var initialStat []time.Time
	var nextStat []time.Time
	getDirectoryModifiedTime(&initialStat, directory)

	fmt.Printf("Watching for changes in `%s`\n", directory)

	for {
		time.Sleep(2 * time.Second)
		nextStat = nil
		getDirectoryModifiedTime(&nextStat, directory)

		for i := 0; i < len(initialStat); i++ {
			if !initialStat[i].Equal(nextStat[i]) {
				fmt.Println("File has changed")
				initialStat = nextStat
				break
			}
		}
	}
}

func getDirectoryModifiedTime(statArr *[]time.Time, dirName string) {
	dirEntry, err := os.ReadDir(dirName)
	if err != nil {
		panic(err)
	}

	for _, dir := range dirEntry {
		info, _ := dir.Info()
		if info.IsDir() {
			getDirectoryModifiedTime(statArr, path.Join(dirName, info.Name()))
		} else {
			*statArr = append(*statArr, info.ModTime())
		}
	}
}
