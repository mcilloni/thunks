package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/mcilloni/thunks/embed/embedder"
)

var (
	sym     = flag.String("sym", "", "package.varName of the exported array")
	outFile = flag.String("out", "", "output file")
)

func errf(f string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, f, args...)
	os.Exit(1)
}

func errln(args ...interface{}) {
	fmt.Fprintln(os.Stderr, args...)
	os.Exit(1)
}

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
		fmt.Fprintln(os.Stderr, "  in string\n    \tfile to be embedded")
	}

	flag.Parse()

	args := flag.Args()
	if len(args) != 1 {
		flag.Usage()
		os.Exit(1)
	}

	if *sym == "" {
		flag.Usage()
		os.Exit(1)
	}

	sms := strings.Split(*sym, ".")
	if len(sms) != 2 {
		errf("error: malformed sym parameter, something in the form `pkg.VarName` was expected, got %s\n", *sym)
	}

	var (
		pkg         = sms[0]
		varName     = sms[1]
		inFileName  = args[0]
		outFileName = ""
	)

	if *outFile != "" {
		outFileName = *outFile
	} else {
		outFileName = strings.Replace(filepath.Base(inFileName), ".", "_", -1) + ".go"
	}

	if err := embedder.Generate(pkg, varName, inFileName, outFileName); err != nil {
		errf("error while generating file %s: %v\n", outFileName, err)
	}
}
