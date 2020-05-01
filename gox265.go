package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strconv"
	"strings"
)

func makeOutputName(inputName string, quality int, ext string) string {
	return strings.Replace(inputName, "."+ext, ".x265_"+strconv.Itoa(quality)+".mkv", -1)
}

func main() {

	minusP := flag.Bool("p", false, "p")
	minusQ := flag.Int("q", 32, "q")
	flag.Parse()
	if len(flag.Args()) == 0 {
		fmt.Println("Usage: cmd <file extention>")
		os.Exit(1)
	}

	list := make([]string, 0)
	Ext := flag.Args()[0]

	err := filepath.Walk(".", func(file string, info os.FileInfo, err error) error {
		_, err = os.Stat(".")
		if err != nil {
			return err
		}
		if filepath.Ext(file) == "."+Ext {
			list = append(list, file)
		}

		return nil
	})

	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	for i := range list {
		dir := filepath.Dir(list[i])
		file := filepath.Base(list[i])
		inputName := path.Join(dir, file)
		outputName := makeOutputName(inputName, *minusQ, Ext)

		fmt.Println("(", i+1, "/", len(list), ")", filepath.Dir(list[i]), "/", filepath.Base(list[i]))
		cmd := exec.Command("ffmpeg", "-n", "-i", inputName, "-map", "0", "-c:a", "copy", "-c:s", "copy", "-c:v", "libx265", "-crf", strconv.Itoa(*minusQ), outputName)

		if *minusP {
			cmd = exec.Command("echo", "ffmpeg", "-n", "-i", inputName, "-map", "0", "-c:a", "copy", "-c:s", "copy", "-c:v", "libx265", "-crf", strconv.Itoa(*minusQ), outputName)
		}

		stdout, err := cmd.Output()
		if err != nil {
			fmt.Println(err)
			os.Exit(3)
		}
		fmt.Println(string(stdout))

	}

}
