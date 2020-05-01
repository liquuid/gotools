package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"sort"
	"testing"
)

var (
	tmpPath string
)

func TestCreateDirs(t *testing.T) {
	setUp()

	want := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}

	path := "/tmp/gotools-testing/"
	got := []string{}
	CreateAlphaDirs(path)
	filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println(err)
		} else {
			got = append(got, info.Name())
		}
		return nil
	})
	// remove first element because is parent dir
	got = got[1:]

	sort.Sort(sort.StringSlice(got))
	sort.Sort(sort.StringSlice(want))

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v , want %v", got, want)
	}
	tearDown(t)
}

func setUp() {
	tmpPath = "/tmp/gotools-testing"
	os.Mkdir(tmpPath, 0700)
}

func TestOrganizeFilesToRightPlace(t *testing.T) {
	want := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
	setUp()
	CreateAlphaDirs(tmpPath)
	tmpPath := "/tmp/gotools-testing/"
	for _, name := range want {
		createEmptyfile(tmpPath, name)
	}

	//got := []string{}
	//cmd := exec.Command("touch aaaaa")
	//err := cmd.Run()
	for _, name := range want {
		//fmt.Printf("%s - %s \n", tmpPath+name, tmpPath+name+"/"+name)
		cmd := exec.Command("mv", tmpPath+name+"*", tmpPath+name+"/")
		fmt.Println("mv", tmpPath+name+"*", tmpPath+name+"/")
		err := cmd.Run()
		fmt.Println(err)

		//err := os.Rename(tmpPath+name, tmpPath+name+"/"+name)

		//	if err != nil {
		//		panic(err)
		//	}
	}

	//Organize(path)

}

func createEmptyfile(tmpPath, name string) {
	os.Chdir(tmpPath)
	newFile, err := os.Create(name + name)
	if err != nil {
		fmt.Println("Error:", err)
	}
	newFile.Close()
}

func tearDown(t *testing.T) {
	err := os.RemoveAll(tmpPath)
	if err != nil {
		t.Fatal(err)
	}
}
