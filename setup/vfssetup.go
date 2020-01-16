package main

import (
	"flag"

	"github.com/siongui/goef"
	"github.com/siongui/gopalilib/dicutil"
)

const websiteDir = "../website"

const bookCSV = "data/dictionary/dict-books.csv"
const wordCSV1 = "data/dictionary/dict_words_1.csv"
const wordCSV2 = "data/dictionary/dict_words_2.csv"

const outBookJSON = websiteDir + "/bookIdAndInfos.json"
const wordJsonDir = websiteDir + "/json/"

const vfsPkgName = "gopaliwordvfs"
const vfsDir = "../src/github.com/siongui/" + vfsPkgName

func main() {
	action := flag.String("action", "", "What kind of action?")
	flag.Parse()

	if *action == "parsebooks" {
		dicutil.ParseDictionayBookCSV(bookCSV, outBookJSON)
	}

	if *action == "parsewords" {
		dicutil.ParseDictionayWordCSV(wordCSV1, wordCSV2, wordJsonDir)
	}

	if *action == "buildvfs" {
		err := goef.GenerateGoPackagePlainTextWithMaxFileSize("gopaliwordvfs", wordJsonDir, vfsDir, 31000000)
		if err != nil {
			panic(err)
		}
	}
}
