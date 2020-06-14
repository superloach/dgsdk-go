package main

import (
	"archive/zip"
	"bytes"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

const (
	url    = "https://dl-game-sdk.discordapp.net/latest/discord_game_sdk.zip"
	outDir = "src"
)

func prefix(f string) bool {
	for _, s := range []string{
		filepath.Join(outDir, "c"),
		filepath.Join(outDir, "lib"),
	} {
		s += string([]rune{filepath.Separator})

		if strings.HasPrefix(f, s) {
			return true
		}
	}

	return false
}

func main() {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	buf1, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		panic(err)
	}

	buf2 := bytes.NewReader(buf1)

	zr, err := zip.NewReader(buf2, buf2.Size())
	if err != nil {
		panic(err)
	}

	for _, zf := range zr.File {
		out := filepath.Join(outDir, zf.Name)
		if !prefix(out) {
			continue
		}

		dir, _ := filepath.Split(out)

		err = os.MkdirAll(dir, 0755)
		if err != nil {
			panic(err)
		}

		zfd, err := zf.Open()
		if err != nil {
			panic(err)
		}
		defer zfd.Close()

		zfc, err := ioutil.ReadAll(zfd)
		if err != nil {
			panic(err)
		}

		f, err := os.Create(out)
		if err != nil {
			panic(err)
		}
		defer f.Close()

		_, err = f.Write(zfc)
		if err != nil {
			panic(err)
		}
		f.Close()
	}
}
