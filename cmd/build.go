/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	// "encoding/json"
	"fmt"
	"strings"
	// "bytes"
	// "time"
	"github.com/jochasinga/requests"
	// "io/ioutil"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"

	"github.com/arghyagod-coder/pkgbuilder/internal"
	"github.com/fatih/color"

	// "github.com/mikkeloscar/aur"
	"github.com/otiai10/copy"
	"github.com/spf13/cobra"
	// "github.com/arghyagod-coder/pkgbuilder/internal"
)

// buildCmd represents the build command
var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Builds tar.zst from AUR packages",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		res, _ := requests.Get(fmt.Sprintf("https://aur.archlinux.org/%v.git", args[0]))
		fmt.Println(res.Status)
		if (res.StatusCode==404){
			color.Red("Could Not Find Package: %v", args[0])
			os.Exit(5)
		}else{
		usr, _ := user.Current()
		dir := usr.HomeDir
		dirpath := filepath.Join(dir,".config","pkgbuilder", "tmp")
		pkgpath := filepath.Join(dir,".config","pkgbuilder", "tmp", args[0])
		os.MkdirAll(dirpath, os.ModePerm)
		copy.Copy(args[0], pkgpath)
		cm := exec.Command("git", "clone", fmt.Sprintf("https://aur.archlinux.org/%v.git", args[0]))
		pwd,_:= os.Getwd()
		cm.Stdout = os.Stdout
		cm.Dir=dirpath
    	cm.Run()
		cm = exec.Command("makepkg", "-sf")
    // var out bytes.Buffer
		cm.Stdout = os.Stdout
		cm.Stderr = os.Stderr
		cm.Dir=pkgpath
    	cm.Run()
		os.Remove(dirpath)
		color.Green("Package Built")
		files, _ := internal.WalkMatch(pkgpath, "*.tar.zst")
		splitpath := strings.Split(files[0], "/")
		filename := splitpath[len(splitpath)-1]
		internal.CopyFile(files[0], filepath.Join(pwd, filename))
		}
	},
}


func init() {
	rootCmd.AddCommand(buildCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// buildCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// buildCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
