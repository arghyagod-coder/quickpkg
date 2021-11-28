/*
Copyright © 2021 Arghya Sarkar <arghyasarkar.nolan@gmail.com>

*/
package cmd

import (
	"fmt"
	"os"
	"log"
	"strings"
    "github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/manifoldco/promptui"
)

// import . "github.com/ahmetb/go-linq/v3"

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create your PKGBUILD",
	Long: `Just enter a handful parameters and get your PKGBUILDs ready.`,
	Run: func(cmd *cobra.Command, args []string) {
		createPKGBUILD()
	},
}

func init() {
	rootCmd.AddCommand(createCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

type project struct {
    name string
    ver string
	rel int
	desc string
	arch string
	url string
	license string
	deps []string
	mdeps []string
	cpkgs []string
}

func createPKGBUILD() {
	p := project{}

	pwd, _ := os.Getwd()
	Pwd := strings.Split(pwd, "/")
	assumption:= Pwd[len(Pwd)-1]

	fmt.Printf("Package Name [%v]: ", assumption)
    var pname string
    fmt.Scanln(&pname)

	if (pname == ""){
		color.Cyan("Set package name to default value %v", assumption)
		p.name = assumption
	}else{
		p.name = pname
	}

	fmt.Printf("Version [0.1.0]: " )
    var pkgver string
    fmt.Scanln(&pkgver)

	if (pkgver == ""){
		color.Cyan("Set project name to default value 0.1.0")
		p.ver = "0.1.0"
	}else{
		p.ver = pkgver
	}

	fmt.Printf("Release Number: " )
    var pkgrel int
    fmt.Scanln(&pkgrel)
	p.rel = pkgrel;

	fmt.Printf("Short Description: " )
    var desc string
    fmt.Scanln(&desc)
	p.desc = desc

	prompt := promptui.Select{
		Label: "Select Architecture",
		Items: []string{"x86_64 (64-bit)", "i686 (32-bit)", "arm (ARM v5)", "armv6h (ARM v7)", "armv7h (ARM v7 Hardfloat)",
			"aarch64 (ARM v8 64-bit)", "any"},
	}

	_, result, err := prompt.Run()

	archs:= []string{"x86_64 (64-bit)", "i686 (32-bit)", "arm (ARM v5)", "armv6h (ARM v7)", "armv7h (ARM v7 Hardfloat)",
			"aarch64 (ARM v8 64-bit)", "any"}
	
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
	}

	if (result==archs[0]){
		p.arch = "x86_64"
	}else 
	if (result==archs[1]) {
		p.arch = "i686"
	}else
	if (result==archs[2]) {
		p.arch = "arm"
	}else
	if (result==archs[3]) {
		p.arch = "armv6h"
	}else
	if (result==archs[4]) {
		p.arch = "armv7h"
	}else
	if (result==archs[5]) {
		p.arch = "aarch64"
	}else
	if (result==archs[6]) {
		p.arch = "any"
	}

	fmt.Printf("Package URL: " )
    var url string
    fmt.Scanln(&url)
	p.url = url;

	fmt.Printf("License: " )
    var license string
    fmt.Scanln(&license)
	p.license = license;

	fmt.Printf("Package Dependencies: [seperate by commas. no spaces]" )
    var deps string
    fmt.Scanln(&deps)
	rdeps:=strings.Split(deps, ",")
	p.deps = rdeps;

	fmt.Printf("Build Dependencies: [seperate by commas. no spaces]" )
    var mdeps string
    fmt.Scanln(&mdeps)
	rmdeps:=strings.Split(mdeps, ",")
	p.mdeps = rmdeps;

	fmt.Printf("Conflicting Packages: [seperate by commas. no spaces]" )
    var pkgs string
    fmt.Scanln(&pkgs)
	cpkgs:=strings.Split(pkgs, ",")
	p.cpkgs = cpkgs;

	
	f, err := os.Create("PKGBUILD")

    if err != nil {
        log.Fatal(err)
    }

    defer f.Close()

    _, err2 := f.WriteString("fk")

    if err2 != nil {
        log.Fatal(err2)
    }

    fmt.Println("done")
}