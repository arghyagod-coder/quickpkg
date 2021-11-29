package internal

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	// "github.com/fatih/color"
	"path/filepath"
)

func CopyFile(src, dst string) {
	bytesRead, err := ioutil.ReadFile(src)

    if err != nil {
        log.Fatal(err)
    }

    err = ioutil.WriteFile(dst, bytesRead, 0644)

    if err != nil {
        log.Fatal(err)
    }
}

func WalkMatch(root, pattern string) ([]string, error) {
    var matches []string
    err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
        if err != nil {
            return err
        }
        if info.IsDir() {
            return nil
        }
        if matched, err := filepath.Match(pattern, filepath.Base(path)); err != nil {
            return err
        } else if matched {
            matches = append(matches, path)
        }
        return nil
    })
    if err != nil {
        return nil, err
    }
    return matches, nil
}

func RunToShell(in string) {
	// command := strings.Split(in, " ")
    cmd := exec.Command(in)
    // var out bytes.Buffer
    cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

    cmd.Run()
}

type VSlice []string

func (s VSlice) String() string {
    var str string
    for _, i := range s {
        str += fmt.Sprintf("%d\n", i)
    }
    return str
}