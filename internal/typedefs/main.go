package main

import (
	"fmt"
	"os"
	"regexp"

	"github.com/superloach/dgsdk-go/internal"
)

const out = "typedefs.go"

var expr = regexp.MustCompile("typedef [^ ]+ Discord([^;[]+)")

func main() {
	outf, err := os.Create(out)
	if err != nil {
		panic(err)
	}
	defer outf.Close()

	internal.GoHeader(outf)

	_, err = fmt.Fprintf(outf, "type (\n")
	if err != nil {
		panic(err)
	}

	for _, sm := range expr.FindAllSubmatch(internal.HeaderBytes(), -1) {
		typ := string(sm[1])

		_, err = fmt.Fprintf(
			outf, "\t%s C.Discord%s\n",
			typ, typ,
		)
		if err != nil {
			panic(err)
		}
	}

	_, err = fmt.Fprintf(outf, ")\n")
	if err != nil {
		panic(err)
	}
}
