package main

import (
	"io/ioutil"
	"os"
	"fmt"
)

func main() {
	main, err := ioutil.ReadFile("templates/main.html")
	if err != nil {
		panic(err)
	}

	toc, err := ioutil.ReadFile("templates/toc.html")
	if err != nil {
		panic(err)
	}
	
	f, err := os.Create("templates_embed.go")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	
	fmt.Fprintln(f, "package main")
	fmt.Fprintln(f, "const (")
	fmt.Fprintf(f, "main_template = `%s`\n\n", string(main))
	fmt.Fprintf(f, "toc_template = `%s`\n\n", string(toc))
	fmt.Fprintln(f, ")")
}
