package main

import (
	"archive/zip"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

const (
	album_path = "../acenet/photos/albums/places/usa/part3"
)

var (
	zipwr    *zip.Writer
	zipftail string
)

func zipDir(dir, zipfname string) {
	cmd := exec.Command("zip", "-u", "-i", "*.jpg", "-j", "-r", zipfname, dir)
	err := cmd.Run()

	fmt.Printf("%s %v", cmd.Path, cmd.Args)

	if err != nil && err.Error() == "exit status 12" {
		println(" -- no changes")
		return
	}

	if err != nil {
		panic(err)
	}

	println(" -- ok")
}

func convertDir(path string, info os.FileInfo, err error) error {
	if path == album_path {
		return nil
	}

	if info.IsDir() {
		return nil
	}

	ext := filepath.Ext(path)
	if ext != ".jpg" {
		return nil
	}

	basename := filepath.Base(path)
	dirname := filepath.Base(filepath.Dir(path))
	if len(basename) < 8 {
		panic("bad name: " + path)
	}

	timestr := basename[:8]
	//	return filepath.SkipDir

	time, err := time.Parse("20060102", timestr)
	if err != nil {
		panic(err)
	}

	zipfname := "ftpdir/" + time.Format("20060102_") + dirname + ".zip"
	fmt.Println(path, zipfname)

	zipDir(filepath.Dir(path), zipfname)

	return filepath.SkipDir

	/*		f, err := os.Open(zipfname)
				if err != nil {
					panic(err)
				}
				zipwr := zip.NewWriter(f)
			}
	*/
	return nil
}

func convertDirsToZips() {
	filepath.Walk(album_path, convertDir)
}

func main() {
	convertDirsToZips()
}
