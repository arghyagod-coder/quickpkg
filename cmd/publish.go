/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"github.com/fatih/color"
	"log"
	"github.com/arghyagod-coder/quickpkg/internal"
	"github.com/spf13/cobra"
)

// publishCmd represents the publish command
var publishCmd = &cobra.Command{
	Use:   "publish",
	Short: "Publish your package to the AUR",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		pwd,_:= os.Getwd()
		fmt.Printf("Enter your Package Name:" )
		var pkgname string
		fmt.Scanln(&pkgname)
		cm := exec.Command("git", "clone", fmt.Sprintf("ssh://aur@aur.archlinux.org/%v.git", pkgname))
		cm.Stdout = os.Stdout
		cm.Stderr = os.Stderr
		cm.Stdin = os.Stdin
		cm.Dir = pwd
		cm.Run()
		pkgdir := filepath.Join(pwd, pkgname)
		internal.CopyFile("PKGBUILD", filepath.Join(pkgdir, "PKGBUILD"))
		files, _ := internal.WalkMatch(pwd, "*.install")
		internal.CopyFile(files[0], files[0])
		srcinfo, _ := exec.Command("makepkg", "--printsrcinfo").Output()
		f, err := os.Create(filepath.Join(pkgdir, ".SRCINFO"))

			if err != nil {
				log.Fatal(err)
			}

			defer f.Close()

			_, err2 := f.WriteString(string(srcinfo))

			if err2 != nil {
				log.Fatal(err2)
			}

		fmt.Printf("Enter your Commit Message:" )
		var cmm string
		fmt.Scanln(&cmm)
		cm = exec.Command("git", "add", ".")
		cm.Dir = pkgdir
		cm.Run()
		cm = exec.Command("git", "commit", "-m", cmm)
		cm.Dir = pkgdir
		cm.Run()
		// cm = exec.Command("git", "push", "-u", "origin", "master")
		// cm.Dir = pkgdir
		// cm.Run()
		color.Green(`Package Published to AUR!
		
Check Out https://aur.archlinux.org/packages/%v`, pkgname)
	},
}

func init() {
	rootCmd.AddCommand(publishCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// publishCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// publishCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
