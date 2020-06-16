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
		fmt.Println("(", i+1, "/", len(list), ")", dir, "/", file)
		inputName := path.Join(dir, file)
		if strings.Contains(inputName, ".x265_") {
			continue
		}
		outputName := makeOutputName(inputName, *minusQ, Ext)

		if _, err := os.Stat(outputName); !os.IsNotExist(err) {
			fmt.Println(outputName, " exists")
			continue
		}
	convert:
		cmd := exec.Command("ffmpeg", "-vaapi_device", "/dev/dri/renderD128", "-i", inputName, "-vf", "format=nv12,hwupload", "-map", "0", "-c:a", "copy", "-c:s", "copy", "-c:v", "hevc_vaapi", "-qp", strconv.Itoa(*minusQ), outputName)
		//ffmpeg -vaapi_device /dev/dri/renderD128  -vf 'format=nv12,hwupload' -c:v hevc_vaapi -qp 30 -c:a copy samu.mkv

		if *minusP {
			cmd = exec.Command("echo", "ffmpeg", "-vaapi_device", "/dev/dri/renderD128", "-i", inputName, "-vf", "format=nv12,hwupload", "-map", "0", "-c:a", "copy", "-c:s", "copy", "-c:v", "hevc_vaapi", "-qp", strconv.Itoa(*minusQ), outputName)
		}

		stdout, err := cmd.Output()
		if err != nil {
			fmt.Println(err)
			os.Remove(outputName)
			//os.Exit(3)
			goto convert
		}
		fmt.Println(string(stdout))

	}

}
