package main

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"sync"

	"github.com/dlclark/regexp2"
)

const BuildPath = "./build"
const DotJS = ".js"

const (
	Comments          string = `\/{2}.*`
	MultilineComments        = `\/\*{2}(.|\n)*\*\/`
	RemovableSpaces          = `(?<![a-zA-Z])\s|\s(?![a-zA-Z])`
)

func FileExists(path string) bool {
	_, err := os.Stat(path)
	return !(err != nil)
}

func Minify(data []byte) []byte {
	// Remove all comments
	a := regexp.MustCompile(fmt.Sprintf("%s|%s", Comments, MultilineComments)).ReplaceAll(data, []byte(""))
	// Remove removable spaces
	b, _ := regexp2.MustCompile(RemovableSpaces, regexp2.None).Replace(string(a), "", -1, -1)

	return []byte(b)
}

func main() {
	if !FileExists(BuildPath) {
		panic(fmt.Sprintf("%s not found. Did you forget to compile?", BuildPath))
	}

	wg := sync.WaitGroup{}
	err := filepath.Walk(BuildPath, func(path string, info fs.FileInfo, err error) error {
		wg.Add(1)

		go func() {
			defer wg.Done()

			// Ignore directories and non .js files
			if info.IsDir() || filepath.Ext(info.Name()) != DotJS {
				return
			}

			data, _ := ioutil.ReadFile(path)

			// Remove the file if it's just an empty export file
			if regexp.MustCompile(`^export\s?{\s?};?\n?$`).Match(data) {
				os.Remove(path)
				return
			}

			minified := Minify(data)

			file, _ := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
			defer file.Close()
			file.Write(minified)
		}()

		return err
	})

	if err != nil {
		panic("Error ")
	}
	wg.Wait()
}
