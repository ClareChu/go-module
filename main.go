package main

import (
	"fmt"
	"github.com/go-cmd/cmd"
	"os"
	"strings"
)

const (
	ENV_GO_PROXY = "GONOPROXY"
)

func main() {
	// Create Cmd, buffered output
	envCmd := cmd.NewCmd("go", "list", "-m", "all")

	// Run and wait for Cmd to return Status
	status := <-envCmd.Start()

	// Print each line of STDOUT from Cmd
	for _, line := range status.Stdout {
		np := os.Getenv(ENV_GO_PROXY)
		np = "git.cloud2go.cn"
		flag := strings.Contains(line, np)
		fmt.Println(flag)
		if true {
			ps := strings.Split(line, " ")
			fmt.Println(ps)
			pkg := ""
			if len(ps) == 2 {
				pkg = fmt.Sprintf("%s@%s", ps[0], ps[1])
			}
			fmt.Printf("go get -insecure %v", pkg)
			envCmd := cmd.NewCmd("go", "get", "-insecure", pkg)

			// Run and wait for Cmd to return Status
			gms := <-envCmd.Start()
			for _, gm := range gms.Stdout {
				fmt.Println(gm)
			}
		}
	}
	/*gb := cmd.NewCmd("go", "build")

	// Run and wait for Cmd to return Status
	gbs := <-gb.Start()
	for _, gs := range gbs.Stdout {
		fmt.Println(gs)
	}*/
}
