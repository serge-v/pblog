package main

import (
	"os"
	"os/exec"
	"fmt"
	"path/filepath"
	"io/ioutil"
	"strings"
	"regexp"
)

const maxToConvert = 200

type photo struct {
	src string
	cached string
}

type blogEntry struct {
	name string
	title string
	path string
	date string
	photos []*photo
}

var ftpDir, cacheDir, buildDir, mainPage string
var entries []*blogEntry
var lastEntry *blogEntry

func collectFtpDirs(path string, info os.FileInfo, err error) error {
	if path == "." {
		return nil
	}

	if info.IsDir() {
		lastEntry = &blogEntry{}
		lastEntry.name = info.Name()
		lastEntry.path = path
		entries = append(entries, lastEntry)
		return nil
	}
	
	ext := filepath.Ext(path)

	if ext == ".jpg" {
		base := filepath.Base(path)
		dir := filepath.Dir(path)
		dst := filepath.Join(cacheDir, dir, base)
		p := &photo{path, dst}
		lastEntry.photos = append(lastEntry.photos, p)
	}

	return nil
}

func collectCachedDirs(path string, info os.FileInfo, err error) error {
	if path == "." {
		return nil
	}

	if info.IsDir() {
		lastEntry = &blogEntry{}
		lastEntry.name = info.Name()
		lastEntry.title = getTitle(lastEntry.name)
		lastEntry.date = getDate(lastEntry.name)
		lastEntry.path = path
		entries = append(entries, lastEntry)
		return nil
	}
	
	ext := filepath.Ext(path)

	if ext == ".jpg" {
		base := filepath.Base(path)
		dir := filepath.Dir(path)
		src := filepath.Join(ftpDir, dir, base)
		if !exists(src) {
			src = ""
		}
		p := &photo{src, path}
		lastEntry.photos = append(lastEntry.photos, p)
	}

	return nil
}

func initVars() {
	var err error
	ftpDir, err = filepath.Abs("ftpdir")
	if err != nil {
		panic(err)
	}

	cacheDir, err = filepath.Abs("cache")
	if err != nil {
		panic(err)
	}

	buildDir, err = filepath.Abs(".")
	if err != nil {
		panic(err)
	}
	
	println("ftpDir:  ", ftpDir)
	println("cacheDir:", cacheDir)
	println("buildDir:", buildDir)

	if !exists(cacheDir) {
		if err := os.Mkdir(cacheDir, 0700); err != nil {
			panic(err)
		}
	}

	text, err := ioutil.ReadFile("templates/main.html")
	if err != nil {
		panic(err)
	}
	mainPage = string(text)
}

func dumpEntries(entries []*blogEntry) {
	for idx, entry := range entries {
		fmt.Println(idx, entry.name)
		for _, photo := range entry.photos {
			fmt.Println("    ", photo.src, photo.cached)
		}
	}
}

func getTitle(name string) string {
	re := regexp.MustCompile(`([A-Z\&])`)
	p := strings.SplitN(name, "_", 2)
	if len(p) < 2 {
		return p[0]
	}
	title := strings.Trim(re.ReplaceAllString(p[1], " $1"), " ")
	return title
}

func getDate(name string) string {
	p := strings.SplitN(name, "_", 2)
	return p[0]
}

func dumpTitles(entries []*blogEntry) {

	for idx, entry := range entries {
		fmt.Println(idx, entry.date, entry.title)
	}
}

func exists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

func convert(p *photo) {
	println("convert", "-auto-orient", p.src+"[480x480]", p.cached)
	cmd := exec.Command("convert", "-auto-orient", p.src+"[480x480]", p.cached)
	if err := cmd.Start(); err != nil {
		panic(err)
	}
	if err := cmd.Wait(); err != nil {
		panic(err)
	}
}

func convertPhotos(entries []*blogEntry) {
	count := 0
	for _, entry := range entries {

		base := filepath.Base(entry.path)
		dir := filepath.Dir(entry.path)
		d := filepath.Join(cacheDir, dir, base)
		if !exists(d) {
			if err := os.Mkdir(d, 0700); err != nil {
				panic(err)
			}
			println("created:", d)
		}

		for _, photo := range entry.photos {
			if exists(photo.cached) {
				continue
			}
			convert(photo)
			count++
			if count >= maxToConvert {
				return
			}
		}
	}
}

func renderPages(entries []*blogEntry) {
	for idx, entry := range entries {
		imgs := ""
		prev_href := ""
		next_href := ""
		page := fmt.Sprintf("%s/page-%03d.html", buildDir, idx)

		for _, photo := range entry.photos {
			rel_cached, err := filepath.Rel(buildDir, cacheDir + "/" + photo.cached)
			if err != nil {
				panic(err)
			}
			imgs += "<br><br><br><br><br>"
			if photo.src != "" {
				rel_src, err := filepath.Rel(buildDir, ftpDir + "/" + photo.cached)
				if err != nil {
					panic(err)
				}
				imgs += "<a href=\"../ftpdir/" + rel_src + "\">"
			}
			imgs += "<img src=\"" + rel_cached + "\">"
			if photo.src != "" {
				imgs += "</a>"
			}
		}

		if idx > 0 {
			prev_href = fmt.Sprintf("&lt;&lt; <a href=\"page-%03d.html\">%s %s</a>", idx-1, entries[idx-1].date, entries[idx-1].title)
		}

		if idx >= 0 {
			if idx == len(entries)-2 {
				next_href = fmt.Sprintf("&gt;&gt; <a href=\"index.html\">%s %s</a>", entries[idx+1].date, entries[idx+1].title)
			} else if idx < len(entries)-2 {
				next_href = fmt.Sprintf("&gt;&gt; <a href=\"page-%03d.html\">%s %s</a>", idx+1, entries[idx+1].date, entries[idx+1].title)
			}
		}

		out := strings.Replace(string(mainPage), "{contents}", imgs, 1)
//		println(imgs)
		out = strings.Replace(out, "{prev_href}", prev_href, 2)
		out = strings.Replace(out, "{next_href}", next_href, 2)
		out = strings.Replace(out, "{title}", entry.title, 1)
		err := ioutil.WriteFile(page, []byte(out), 0666)
		if err != nil {
			panic(err)
		}
		println("saved:", page)
	}
}

func linkIndexPage() {
	page := fmt.Sprintf("%s/page-%03d.html", buildDir, len(entries)-1)
	os.Remove(buildDir + "/index.html")
	if err := os.Link(page, buildDir + "/index.html"); err != nil {
		panic(err)
	}
	println("link:", page, " -> " + buildDir + "/index.html")
}

func main() {
	initVars()
	
	if err := os.Chdir(ftpDir); err != nil {
		panic(err)
	}

	filepath.Walk(".", collectFtpDirs)

	convertPhotos(entries)

	entries = nil
	if err := os.Chdir(cacheDir); err != nil {
		panic(err)
	}
	filepath.Walk(".", collectCachedDirs)

	renderPages(entries)
	linkIndexPage()
//	dumpEntries(entries)
	dumpTitles(entries)
}
