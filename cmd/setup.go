/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
	"os/user"
	"path/filepath"
	"strings"
)

// setupCmd represents the setup command
var setupCmd = &cobra.Command{
	Use:   "setup",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		usr, _ := user.Current()
		dir := usr.HomeDir

		b, err := ioutil.ReadFile(filepath.Join(dir, ".ssh", "config"))
		if err != nil {
			panic(err)
		}
		s := string(b)
		sshconfig := `Host aur.archlinux.org
  IdentityFile ~/.ssh/aur
  User aur`
		fmt.Println(strings.Contains(s, sshconfig))
		if (strings.Contains(s, sshconfig)) == true {

		} else {
			os.MkdirAll(filepath.Join(dir, ".ssh"), os.ModePerm)
			f, err := os.OpenFile(filepath.Join(dir, ".ssh", "config"), os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
			if err != nil {
				panic(err)
			}

			defer f.Close()

			if _, err = f.WriteString(sshconfig); err != nil {
				panic(err)
			}

		}
	},
}

func init() {
	rootCmd.AddCommand(setupCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// setupCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// setupCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
