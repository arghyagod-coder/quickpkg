/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"

	"github.com/spf13/cobra"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		InitConfig()
		jokeTerm, _ := cmd.Flags().GetBool("auto")
		if (jokeTerm==true){
			name, err := exec.Command("git", "config", "--get", "user.name").Output()
			if err != nil {
				log.Fatal(err)
			}
			mail, err := exec.Command("git", "config", "--get", "user.email").Output()
			if err != nil {
				log.Fatal(err)
			}
			usr, _ := user.Current()
			dir := usr.HomeDir
			path := filepath.Join(dir,".config","pkgbuilder","config.json")
			config := Configuration{UserName: string(name), Email: string(mail)}
			jsonstruct, _ := json.MarshalIndent(config, "", " ")
			_ = ioutil.WriteFile(path, jsonstruct, 0644)
			
		}
	},
}

type Configuration struct {
	UserName, Email string
}

func init() {
	rootCmd.AddCommand(configCmd)
	configCmd.PersistentFlags().BoolP("auto", "a", false, "Use your git credentials for account")
}

func CheckFileExists(filePath string) bool {
	_, e := os.Stat(filePath)
	return !os.IsNotExist(e)
}

func InitConfig(){
	usr, _ := user.Current()
	dir := usr.HomeDir
	path := filepath.Join(dir,".config","pkgbuilder","config.json")
	dirpath := filepath.Join(dir,".config","pkgbuilder")
	if (CheckFileExists(path)==true){
		fmt.Println("Configuration File present at ~/.config/pkgbuilder/")
	}else{
	os.MkdirAll(dirpath, os.ModePerm)
	newFile, err := os.Create(path)
    if err != nil {
        log.Fatal(err)
    }
	newFile.Close()
    
}
}