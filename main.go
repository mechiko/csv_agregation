package main

import (
	"agregat/process"
	"agregat/processall"
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
	pAll, err := processall.New()
	checkErr(err)
	for _, file := range files {
		p, err := process.New(file)
		// заполняем короба палеты и КМ
		checkErr(err)
		err = p.Scan()
		checkErr(err)
		err = p.Save(outDir)
		checkErr(err)
		err = pAll.Add(file)
		checkErr(err)
	}
	err = pAll.ScanAll()
	checkErr(err)
	err = pAll.Save(outDir)
}
