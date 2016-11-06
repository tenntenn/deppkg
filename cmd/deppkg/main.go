package main

import (
	"fmt"
	"go/build"
	"os"

	"github.com/tenntenn/deppkg"
)

func main() {
	if len(os.Args) <= 1 {
		fmt.Fprintln(os.Stderr, "%s rootpackage file1.go file2.go...", os.Args[0])
		os.Exit(1)
	}

	pkgs, err := deppkg.Main(&build.Default, os.Args[1], os.Args[2:])
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}

	for _, p := range pkgs {
		fmt.Println(p.ImportPath)
	}
}
