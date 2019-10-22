package main

import (
	"fmt"
	"github.com/go-cmd/cmd"
	"github.com/sirkon/goproxy/gomod"
	"io/ioutil"
	"os"
	"strings"
)

const (
	ENV_GO_PROXY = "GONOPROXY"
)

func main() {
	input, err := ioutil.ReadFile("go.mod")
	if err != nil {

	}
	got, err := gomod.Parse("go.mod", input)
	fmt.Println(got)
	if err != nil {
		fmt.Println("go mod parse err ")
		return
	}
	// Print each line of STDOUT from Cmd
	for k, v := range got.Require {
		np := os.Getenv(ENV_GO_PROXY)
		np = "git.cloud2go.cn"
		flag := strings.Contains(k, np)
		fmt.Println()
		fmt.Println(flag)
		if true {
			pkg := fmt.Sprintf("%s@%s", k, v)
			fmt.Printf("go get -insecure %v", pkg)
			envCmd := cmd.NewCmd("go", "get", "-u", "-v", "-insecure", pkg)
			//envCmd = cmd.NewCmd("go", "get", "-u", "-v", "-insecure", "git.cloud2go.cn/rd/orchor@v1.4.12")
			// Run and wait for Cmd to return Status
			gms := <-envCmd.Start()
			for _, gm := range gms.Stderr {
				fmt.Println(gm)
			}
		}
	}
/*	gb := cmd.NewCmd("go", "build")

	// Run and wait for Cmd to return Status
	gbs := <-gb.Start()
	for _, gs := range gbs.Stdout {
		fmt.Println(gs)
	}*/
}
