package main

import (
	//"bytes"
	"flag"
	"fmt"
	"github.com/ebayboy/gohs/hyperscan"
	"io/ioutil"
	"os"
)

func eventHandler(id uint, from, to uint64, flags uint, context interface{}) error {
	inputData := context.([]byte)

	/*
		start := bytes.LastIndexByte(inputData[:from], '\n')
		end := int(to) + bytes.IndexByte(inputData[to:], '\n')

		if start == -1 {
			start = 0
		} else {
			start += 1
		}

		if end == -1 {
			end = len(inputData)
		}

		fmt.Printf("Matched:%s%s\n", inputData[start:from], inputData[to:end])
	*/
	fmt.Printf("Matched:%s id:%d\n", inputData[from:to], id)

	return nil
}

func main() {
	flag.Parse()

	if flag.NArg() != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <pattern> <input file>\n", os.Args[0])
		os.Exit(-1)
	}

	inputFN := flag.Arg(1)

	/* Next, we read the input data file into a buffer. */
	inputData, err := ioutil.ReadFile(inputFN)
	if err != nil {
		os.Exit(-1)
	}
	fmt.Println("inputData:", string(inputData))
	fmt.Println("pattern:", flag.Arg(0))

	pattern0 := hyperscan.NewPattern(".*bcde", hyperscan.DotAll|hyperscan.SomLeftMost)
	pattern0.Id = 10000
	pattern := hyperscan.NewPattern(flag.Arg(0), hyperscan.DotAll|hyperscan.SomLeftMost)
	pattern.Id = 10001
	database, err := hyperscan.NewBlockDatabase(pattern0, pattern)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: Unable to compile pattern \"%s\": %s\n", pattern.String(), err.Error())
		os.Exit(-1)
	}
	defer database.Close()

	scratch, err := hyperscan.NewScratch(database)
	if err != nil {
		fmt.Fprint(os.Stderr, "ERROR: Unable to allocate scratch space. Exiting.\n")
		os.Exit(-1)
	}

	defer scratch.Free()

	if err := database.Scan(inputData, scratch, eventHandler, inputData); err != nil {
		fmt.Fprint(os.Stderr, "ERROR: Unable to scan input buffer. Exiting.\n")
		os.Exit(-1)
	}

	return
}
