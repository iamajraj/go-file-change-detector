package main

import (
	"fmt"
	"os"
	"path"
	"time"
)

var directory string

func parseArgs(arguments []string) {
	for i, arg := range arguments {
		idx := i + 1
		if arg == "-d" {
			directory = arguments[idx]
		}
	}
}

func main() {
	if len(os.Args[1:]) > 1 {
		parseArgs(os.Args[1:])
	} else {
		fmt.Println("Please provide the directory e.g -d <path>")
		os.Exit(0)
	}

	var initialStat []time.Time
	var nextStat []time.Time
	getDirectoryModifiedTime(&initialStat, directory)

	fmt.Printf("Watching for changes in `%s`\n", directory)

	for {
		time.Sleep(1 * time.Second)
		nextStat = nil
		getDirectoryModifiedTime(&nextStat, directory)

		for i := 0; i < len(initialStat); i++ {
			if !initialStat[i].Equal(nextStat[i]) {
				fmt.Println("The File has changed")
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
