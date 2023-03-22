package main

import (
	"flag"
	"fmt"
	"io"
	"net/url"
	"os"
	"strings"
)

func main() {
	encode := flag.Bool("encode", false, "encode input")
	decode := flag.Bool("decode", false, "decode input")
	input := flag.String("input", "", "input string")
	flag.Parse()

	var inp string
	if *input != "" {
		inp = *input
	} else {
		stdInput, err := io.ReadAll(os.Stdin)
		if err != nil {
			fmt.Fprintf(os.Stderr, "err: failed to read from stdin: %+v\n", err)
			os.Exit(1)
		}
		inp = strings.Trim(string(stdInput), "\n")
	}

	if *encode {
		fmt.Fprintln(os.Stdout, url.QueryEscape(inp))
	}

	if *decode {
		unescaped, err := url.QueryUnescape(inp)
		if err != nil {
			fmt.Fprintf(os.Stderr, "err: failed to unescape input: %+v\n", err)
		}
		fmt.Fprintln(os.Stdout, unescaped)
	}
}
