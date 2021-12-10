/*
Copyright © 2021 Arghya Sarkar <arghyasarkar.nolan@gmail.com>

*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/arghyagod-coder/quickpkg/internal"
	"github.com/fatih/color"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create your PKGBUILD",
	Long: `Just enter a handful parameters and get your PKGBUILDs ready.`,
	Run: func(cmd *cobra.Command, args []string) {
		createPKGBUILD()
	},
}

var syntax string

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
	licenseurl string
	licensesum string
	deps []string
	mdeps []string
	cpkgs []string
	iss string
	srcs []string
	s256s []string

}


type VSlice []string

func (s VSlice) String() string {
    var str string
    for _, i := range s {
        str += fmt.Sprintf("%d\n", i)
    }
    return str
}

func createPKGBUILD() {
	p := project{}

	pwd, _ := os.Getwd()
	Pwd := strings.Split(pwd, "/")
	assumption:= Pwd[len(Pwd)-1]

	// fmt.Printf("Package Name [%v]: ", assumption)
    // var pname string
	pname := internal.Input(fmt.Sprintf("Package Name [%v] ",assumption))

	if (pname == ""){
		color.Cyan("Set package name to default value %v", assumption)
		p.name = assumption
	}else{
		p.name = pname
	}

	pkgver := internal.Input("Version [0.1.0]")

	if (pkgver == ""){
		color.Cyan("Set project name to default value 0.1.0")
		p.ver = "0.1.0"
	}else{
		p.ver = pkgver
	}

	pkgrel := internal.IntInput("Release Number")
	p.rel = pkgrel;

	desc := internal.Input("Short Description")
	p.desc = desc;

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
	};

	url:=internal.Input("Package URL: " )
	p.url = url;

	license:=internal.Input("License: " )
	p.license = license;

	licenseurl:=internal.Input("License URL (Raw File): " )
	p.licenseurl = licenseurl;

	licensesum:=internal.Input("License sha256sum: " )
	p.licensesum = licensesum;

	deps:=internal.Input("Package Dependencies [seperate by commas. no spaces]: " )
	rdeps:=strings.Split(deps, ",")
	p.deps = rdeps;

	mdeps:=internal.Input("Build Dependencies [seperate by commas. no spaces]: " )
	rmdeps:=strings.Split(mdeps, ",")
	p.mdeps = rmdeps;

	pkgs:=internal.Input("Conflicting Packages [seperate by commas. no spaces]: " )
	cpkgs:=strings.Split(pkgs, ",")
	p.cpkgs = cpkgs;

	iss:=internal.Input("Create a Post Install Script? [yes/no]: " )

	if (iss == "no"){
		
	}else if (iss=="yes"){
		p.iss = fmt.Sprintf("%v.install", p.name)
		f, err := os.Create(p.iss)

		if err != nil {
			log.Fatal(err)
		}

		defer f.Close()
	}else{
		fmt.Println("Invalid Value")
	}

	srcs:=internal.Input("Source Files [seperate by commas. no spaces]: " )
	lsrcs:=strings.Split(srcs, ",")
	p.srcs = lsrcs;

	s256s:=internal.Input("Sha256sums of the Source Files [seperate by commas. no spaces]: " )
	ls256s:=strings.Split(s256s, ",")
	p.s256s = ls256s;

	action, dest, target, buildi:= GetJSON()
	if (action=="copy"){
		syntax=fmt.Sprintf(`install -dm755 ${pkgdir}${_destname}
cp -r  ${srcdir}/%v/* ${pkgdir}${_destname}
		`, target)
	}else 
	if (action=="install"){
		syntax=fmt.Sprintf(`install -dm755 ${pkgdir}${_destname}
install -Dm755  ${srcdir}/%v/* ${pkgdir}${_destname}
`, target)
	}else
	if (len(buildi)==0){
		// buildl := strings.ReplaceAll((VSlice(buildi)), )
		syntax=fmt.Sprintf(`%v
install -dm755 ${pkgdir}${_destname}
install -Dm755  ${srcdir}/%v/* ${pkgdir}${_destname}
`, VSlice(buildi) ,target)
	}else{
		fmt.Println("Error in Build File")
	}

	user, mail := GetUserConfig()



	content := fmt.Sprintf(`# Maintainer: %v %v
pkgname=%v
_pkgname=%v
_destname="%v"
_licensedir="/usr/share/licenses/${_pkgname}/"
pkgver=%v
pkgrel=%v
epoch=
pkgdesc="%v"
arch=('%v')
url="%v"
license=(%v)
groups=()
depends=(%s)
makedepends=(%s)
checkdepends=()
optdepends=()
provides=(%v)
conflicts=(%v)
backup=()
options=()
install=%v
source=(%v
		"%v")
noextract=("${source[@]##*/}")
sha256sums=(%v
			"%v")
validpgpkeys=()
package() {
    %v
    install -dm755 ${pkgdir}${_licensedir}${_pkgname}
	install -m644  ${srcdir}/LICENSE ${pkgdir}${_licensedir}${_pkgname}
}`, user,mail, p.name, p.name,dest, p.ver, p.rel, p.desc, p.arch, p.url, p.license, internal.ArrStrings(p.deps), internal.ArrStrings(p.mdeps), p.name, internal.ArrStrings(p.cpkgs), p.iss, internal.ArrStrings(p.srcs), p.licenseurl, internal.ArrStrings(p.s256s), p.licensesum, syntax)
	f, err := os.Create("PKGBUILD")

    if err != nil {
        log.Fatal(err)
    }

    defer f.Close()

    _, err2 := f.WriteString(content)

    if err2 != nil {
        log.Fatal(err2)
    }

    color.Green("PKGBUILD created!")
}

type Build struct {
	Target             		string	 	`json:"target"`              // wallpaper search terms for unsplash
	Destination             string   	`json:"destination"`                // wallpaper resolution, defaults to 1600x900
	Action         			string  	`json:"action"`          // whether change wallpaper after a duration
	Instructions			[]string   	`json:"build_instructions"` // if wallpaper has to be changed, then after how many minutes
	}

func GetJSON()(string, string, string, []string){
	jsonFile, err := os.Open("buildfile.json")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Found buildfile.json..")
	time.Sleep(5)
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var build Build

	json.Unmarshal(byteValue, &build)
	return build.Action, build.Destination, build.Target, build.Instructions
}

type ConfigData struct {
	UserName string 	`json:"UserName"`
	Email string		`json:"Email"`
}	

func GetUserConfig()(string, string){
	homedir, _ := os.UserHomeDir()
	jsonFile, err := os.Open(filepath.Join(homedir, ".config", "quickpkg", "config.json"))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Found config.json..")
	time.Sleep(5)
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var config ConfigData

	json.Unmarshal(byteValue, &config)
	uname := strings.ReplaceAll(config.UserName, "\n", "")
	mail := strings.ReplaceAll(config.Email, "\n", "")
	return uname, mail
}