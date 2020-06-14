package main

import (
	"fmt"
	"os"
	"regexp"

	"github.com/superloach/dgsdk-go/internal"
)

const out = "consts.go"

var defineExpr = regexp.MustCompile("#define DISCORD_([^ ]+) (.+)")

func main() {
	outf, err := os.Create(out)
	if err != nil {
		panic(err)
	}
	defer outf.Close()

	internal.GoHeader(outf)

	_, err = fmt.Fprintf(outf, "const (\n")
	if err != nil {
		panic(err)
	}

	for _, sm := range defineExpr.FindAllSubmatch(internal.HeaderBytes(), -1) {
		name := string(sm[1])

		_, err := fmt.Fprintf(
			outf, "\tDiscord%s = C.DISCORD_%s\n",
			internal.CamelCase(name), name,
		)
		if err != nil {
			panic(err)
		}
	}

	_, err = fmt.Fprintf(outf, ")\n")
	if err != nil {
		panic(err)
	}

	outf.Close()
}
