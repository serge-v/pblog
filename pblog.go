package main

import (
	"os"
	"os/exec"
	"fmt"
	"path/filepath"
	"io/ioutil"
	"strings"
)

var (
	cacheDir = "../cache"
	chunk = ""
	pagesCount = 1
	mainPage = ""
	lastFile = ""
	title = ""
)

func saveChunk() {
	prev := lastFile
	lastFile = fmt.Sprintf("../build/page-%03d.html", pagesCount)
	out := strings.Replace(string(mainPage), "{contents}", chunk, 1)
	out = strings.Replace(out, "{prev}", prev, 2)
	out = strings.Replace(out, "{title}", title, 1)
	err := ioutil.WriteFile(lastFile, []byte(out), 0666)
	if err != nil {
		panic(err)
	}
	
	pagesCount++
	println("saved:", lastFile)
	chunk = ""
}

func walk(path string, info os.FileInfo, err error) error {

	base := filepath.Base(path)
	dir := filepath.Dir(path)
	
	cachedName := filepath.Join(cacheDir, dir, base)
	
	_, err = os.Stat(cachedName)
	if os.IsNotExist(err) {
		if info.IsDir() {
			if err = os.Mkdir(cachedName, 0777); err != nil {
				panic(err)
			}
		} else {
			cmd := exec.Command("cp", path, cachedName)
			if err = cmd.Start(); err != nil {
				panic(err)
			}
			if err = cmd.Wait(); err != nil {
				panic(err)
			}
		}
	}

	if info.IsDir() && len(chunk) > 0 {
		saveChunk()
		return nil
	}
	
	title = dir
	chunk += "<br><br><br><br><br><img src=\"" + cachedName + "\"/>\n"
	return nil
}

func init_flags() {
	os.Mkdir(cacheDir, 0777)
	os.Chdir("ftpdir")

	text, err := ioutil.ReadFile("../templates/main.html")
	if err != nil {
		panic(err)
	}
	mainPage = string(text)
}

func main() {
	init_flags()
	filepath.Walk(".", walk)
	if len(chunk) > 0 {
		saveChunk()
	}
	
	os.Remove( "../build/index.html")
	if err := os.Link(lastFile, "../build/index.html"); err != nil {
		panic(err)
	}
}
