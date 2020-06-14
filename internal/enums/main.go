package main

import (
	"fmt"
	"os"
	"regexp"

	"github.com/superloach/dgsdk-go/internal"
)

var nameExpr = regexp.MustCompile("enum EDiscord([^ ]+) {")

func valueExpr(name string) *regexp.Regexp {
	return regexp.MustCompile(fmt.Sprintf(
		"Discord%s_([^ ]+)(?: = ([^,]+))?,",
		name,
	))
}

func main() {
	hdr := internal.HeaderBytes()

	for _, sm := range nameExpr.FindAllSubmatch(hdr, -1) {
		name := string(sm[1])

		enum(hdr, name)
	}
}

func enum(hdr []byte, typ string) {
	out := internal.SnakeCase(typ) + ".go"

	outf, err := os.Create(out)
	if err != nil {
		panic(err)
	}
	defer outf.Close()

	internal.GoHeader(outf)

	_, err = fmt.Fprintf(
		outf, "type %s C.enum_EDiscord%s\n\nconst (\n",
		typ, typ,
	)
	if err != nil {
		panic(err)
	}

	expr := valueExpr(typ)
	for i, sm := range expr.FindAllSubmatch(hdr, -1) {
		name := string(sm[1])

		extra := ""
		if i == 0 {
			extra = " " + typ
		}

		_, err := fmt.Fprintf(
			outf, "\t%s%s%s = C.Discord%s_%s\n",
			typ, name, extra, typ, name,
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
