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

func main() {
	//minusR := flag.Bool("r", false, "r")
	minusQ := flag.Int("q", 32, "q")

	list := make([]string, 0)

	flag.Parse()

	for index, val := range flag.Args() {
		fmt.Println(index, ":", val)
	}
	//fmt.Println("-r :", *minusR)
	//fmt.Println("-q :", string(*minusQ))

	Ext := flag.Args()[0]

	err := filepath.Walk(".", func(file string, info os.FileInfo, err error) error {
		_, err = os.Stat(".")
		if err != nil {
			return err
		}
		if filepath.Ext(file) == "."+Ext {
			list = append(list, file)
			//fmt.Println("e", Ext, filepath.Ext(ext))
		}

		return nil
	})

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	for i := range list {
		dir := filepath.Dir(list[i])
		file := filepath.Base(list[i])
		fmt.Println(i, filepath.Dir(list[i]), " -> ", filepath.Base(list[i]))
		cmd := exec.Command("ffmpeg", "-n", "-i", path.Join(dir, file), "-map", "0", "-c:a", "copy", "-c:s", "copy", "-c:v", "libx265", "-crf", strconv.Itoa(*minusQ), strings.Replace(path.Join(dir, file), ".mp4", ".x265_"+strconv.Itoa(*minusQ)+".mkv", -1))
		stdout, err := cmd.Output()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(string(stdout))
	}

}
