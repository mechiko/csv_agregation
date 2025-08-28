package main

import (
	"agregat/process"
	"log"
	"os"
	"regexp"

	"github.com/mechiko/utility"
)

const inDir = `IN`
const outDir = `OUT`

func checkErr(e error) {
	if e != nil {
		log.Printf("%v", e)
		os.Exit(-1)
	}
}

func main() {
	re, err := regexp.Compile(`.*\.csv$`)
	checkErr(err)
	files, err := utility.FilteredSearchOfDirectoryTree(re, inDir)
	checkErr(err)
	checkErr(os.MkdirAll(outDir, 0o755))
	for _, file := range files {
		p, err := process.New(file)
		// заполняем короба палеты и КМ
		checkErr(err)
		err = p.Scan()
		checkErr(err)
		err = p.Save(outDir)
		checkErr(err)
	}
}
