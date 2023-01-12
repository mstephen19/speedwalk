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

const PackageJsonPath = "./package.json"
const BuildPath = "./build"
const DotJS = ".js"

type Regex = string

const (
	Comments          Regex = `\/{2}.*`
	MultilineComments Regex = `\/\*{2}(.|\n)*\*\/`
	RemovableSpaces   Regex = `(?<![a-zA-Z])\s|\s(?![a-zA-Z])`
	JsonSpaces        Regex = `(?<=[:{}\n\[\]],?|(",?))\s{1,}`
)

func FileExists(path string) bool {
	_, err := os.Stat(path)
	return !(err != nil)
}

func MinifyJS(data []byte) []byte {
	// Remove all comments
	a := regexp.MustCompile(fmt.Sprintf("%s|%s", Comments, MultilineComments)).ReplaceAll(data, []byte(""))
	// Remove removable spaces
	b, _ := regexp2.MustCompile(RemovableSpaces, regexp2.None).Replace(string(a), "", -1, -1)
	return []byte(b)
}

func MinifyJson(data []byte) []byte {
	a, _ := regexp2.MustCompile(JsonSpaces, regexp2.None).Replace(string(data), "", -1, -1)
	return []byte(a)
}

func main() {
	if !FileExists(BuildPath) {
		panic(fmt.Sprintf("%s not found. Did you forget to compile?", BuildPath))
	}
	if !FileExists(PackageJsonPath) {
		panic("package.json somehow not found.")
	}

	wg := sync.WaitGroup{}

	wg.Add(1)
	// Minify package.json
	go func() {
		defer wg.Done()

		data, _ := ioutil.ReadFile(PackageJsonPath)
		file, _ := os.OpenFile(PackageJsonPath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
		defer file.Close()
		file.Write(MinifyJson(data))
	}()

	// Minify all JS files there.
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

			minified := MinifyJS(data)

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
