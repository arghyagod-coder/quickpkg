/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"os/user"
	"github.com/fatih/color"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/arghyagod-coder/pkgbuilder/internal"
	"github.com/spf13/cobra"
)

// testCmd represents the test command
var testCmd = &cobra.Command{
	Use:   "test",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		usr, _ := user.Current()
		dir := usr.HomeDir
		dirpath := filepath.Join(dir,".config","pkgbuilder", "tmp")
		pkgpath := filepath.Join(dir,".config","pkgbuilder", "tmp", "PKGBUILD")
		files, _ := internal.WalkMatch("./", "*.install")
		inspath := filepath.Join(dir,".config","pkgbuilder", "tmp", files[0])
		os.MkdirAll(dirpath, os.ModePerm)
		internal.CopyFile("PKGBUILD", pkgpath)
		internal.CopyFile(files[0], inspath)
		// pwd, _ := os.Getwd()
		cm := exec.Command("makepkg", "-si")
    // var out bytes.Buffer
		cm.Stdout = os.Stdout
		cm.Stderr = os.Stderr
		cm.Dir=dirpath
    	cm.Run()
		cm = exec.Command("tree")
    // var out bytes.Buffer
		cm.Stdout = os.Stdout
		cm.Stderr = os.Stderr
		cm.Dir=dirpath
    	cm.Run()
		os.Remove(dirpath)
		color.Green("Testing Completed")
	},
}

func init() {
	rootCmd.AddCommand(testCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// testCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// testCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
