package main

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"sync"
)

const BuildPath = "./build"
const DotJS = ".js"

func FileExists(path string) bool {
	_, err := os.Stat(path)
	return !(err != nil)
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

			// Remove comments from just JS files since the TypeScript compiler removes comments
			// from .d.ts files as well when provided the "removeComments" option
			minified := regexp.MustCompile(`\/{2}.*\n|\/\*{2}((.|\n)*)\*\/|\s{2,}|\n`).ReplaceAll(data, []byte(""))
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
