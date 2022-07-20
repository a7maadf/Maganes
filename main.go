package main

import (
	"fmt"
	"os"
	"strings"
)

func init() {
	
	fmt.Println("Welcome to Maganes programming language, we fight feminists 24/7 here...\n")
	defer func() {
		if r := recover(); r != nil {
			if _, ok := r.(error); ok {
				fmt.Println("No file specified")
				os.Exit(0)
				return
			}
			panic(r)
		}
	}()
	x := sys.argv[1]
}

func lineArguments() {
	var f int
	func() {
		defer func() {
			if r := recover(); r != nil {
				if _, ok := r.(error); ok {
					fmt.Println("Error, Check if file exists")
					os.Exit(0)
					f := 2
					return
				}
				panic(r)
			}
		}()
		f := func() *os.File {
			f, err := os.OpenFile(sys.argv[1], os.O_RDONLY, 0o777)
			if err != nil {
				panic(err)
			}
			return f
		}()
	}()
	f
}

func importingTheJson() {
	f := func() *os.File {
		f, err := os.OpenFile("cmd/cmds.json", os.O_RDONLY, 0o777)
		if err != nil {
			panic(err)
		}
		return f
	}()
	jsondic := json.load(f)
	jsondic
}

func commandsCodes() {
	f := lineArguments()
	jsondic := importingTheJson()
	TheWholeCode := ""
	intolines := f.readlines()
	for _, i := range intolines {
		line := func() {
			elts := []interface{}{}
			for _, elt := range i {
				elts = append(elts, elt)
			}
			elts
		}()
		fullline := ""
		for _, i := range line {
			fullline += i
		}
		for _, i := range jsondic {
			if strings.Contains(fullline, string(base64.b64decode([]byte(jsondic[i])))) {
				fmt.Println(
					"Unknown",
					"'",
					string(base64.b64decode([]byte(jsondic[i]))),
					"'",
					"syntax found.\nPlease review your code...",
				)
				os.Exit(0)
			}
			func() {
				defer func() {
					if r := recover(); r != nil {
						if _, ok := r.(error); ok {
							return
						}
						panic(r)
					}
				}()
				fullline = re.sub(i, string(base64.b64decode([]byte(jsondic[i]))), fullline)
			}()
		}
		TheWholeCode += fullline
	}
	exec(TheWholeCode)
}

func init() {
	defer func() {
		if r := recover(); r != nil {
			if _, ok := r.(error); ok {
				fmt.Println("Something doesn't seem right, please review your code and try again.")
				return
			}
			panic(r)
		}
	}()
	lineArguments()
	commandsCodes()
}
