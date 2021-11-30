/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"os"
	"os/exec"
	"os/user"
	"path/filepath"

	"github.com/fatih/color"

	"github.com/arghyagod-coder/quickpkg/internal"
	"github.com/spf13/cobra"
)

// testCmd represents the test command
var testCmd = &cobra.Command{
	Use:   "test",
	Short: "Test your Package",
	Long: `Test the PKGBUILD and other files created by "pkgbd create" commmand. This command should be ran in the directory where you ran the create command. Not`,
	Run: func(cmd *cobra.Command, args []string) {
		usr, _ := user.Current()
		dir := usr.HomeDir
		dirpath := filepath.Join(dir, ".config", "quickpkg", "tmp")
		pkgpath := filepath.Join(dir, ".config", "quickpkg", "tmp", "PKGBUILD")
		files, _ := internal.WalkMatch("./", "*.install")
		inspath := filepath.Join(dir, ".config", "quickpkg", "tmp", files[0])
		os.MkdirAll(dirpath, os.ModePerm)
		internal.CopyFile("PKGBUILD", pkgpath)
		internal.CopyFile(files[0], inspath)
		// pwd, _ := os.Getwd()
		cm := exec.Command("makepkg", "-si", "--noconfirm")
		// var out bytes.Buffer
		cm.Stdout = os.Stdout
		cm.Stderr = os.Stderr
		cm.Stdin = os.Stdin
		cm.Dir = dirpath
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
