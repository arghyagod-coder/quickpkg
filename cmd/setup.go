/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	// "fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"github.com/fatih/color"
	"os/exec"
	"os"
	"os/user"
	"log"
	"path/filepath"
	"strings"
)

// setupCmd represents the setup command
var setupCmd = &cobra.Command{
	Use:   "setup",
	Short: "Setup your SSH keys for AUR",
	Long: `Setup your SSH keys through AUR. 
	
While registering to AUR (Arch User Repository), you will need to provide the public ssh key to gain write access to AUR. So this command does that automatically`,
	Run: func(cmd *cobra.Command, args []string) {
		usr, _ := user.Current()
		dir := usr.HomeDir
		myfile,_:=os.Create(filepath.Join(dir, ".ssh", "config"))
			myfile.Close()
		b, err := ioutil.ReadFile(filepath.Join(dir, ".ssh", "config"))
		if err != nil {
			panic(err)
		}
		s := string(b)
		var sshconfig string = `Host aur.archlinux.org
  IdentityFile ~/.ssh/aur
  User aur`
		if (strings.Contains(s, sshconfig)) == true {

		} else {
			os.MkdirAll(filepath.Join(dir, ".ssh"), os.ModePerm)
			f, err := os.OpenFile(filepath.Join(dir, ".ssh", "config"),
				os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				log.Println(err)
			}
			defer f.Close()
			if _, err := f.WriteString(sshconfig); err != nil {
				log.Println(err)
			}

		}
		cm := exec.Command("ssh-keygen", "-f", "~/.ssh/aur")
		cm.Stdout = os.Stdout
		cm.Stderr = os.Stderr
		cm.Stdin = os.Stdin
		cm.Dir = dir
		cm.Run()
		color.Green("Private Key Generated\n\n")
		f, err:= ioutil.ReadFile(filepath.Join(dir, ".ssh", "aur"))
		color.White(string(f))
		color.Green("\n\nPublic Key Generated\n\n")
		f, err= ioutil.ReadFile(filepath.Join(dir, ".ssh", "aur.pub"))
		color.White(string(f))
		color.Blue(`
		
Now Proceed by entering the Public SSH Key to the AUR Account Creation at https://aur.archlinux.org/register/

Your AUR Account is setup`)
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
