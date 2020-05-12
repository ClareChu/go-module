package main

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/go-cmd/cmd"
	"io/ioutil"
	"os"
	"strings"
)

const ENVPRIVATE = "GOPRIVATE"

func main() {
	sumContent, err := ioutil.ReadFile("go.sum")
	if err != nil {
		fmt.Println(fmt.Sprintf("read go.sum failed,got :%s", err.Error()))
	}
	scanner := bufio.NewScanner(bytes.NewReader(sumContent))
	private := os.Getenv(ENVPRIVATE)
	fmt.Println(fmt.Sprintf("GOPRIVATE:%s", private))
	for scanner.Scan() {
		// read line by line
		line := scanner.Text()
		parts := strings.Fields(line)
		pkg := parts[0]
		version := parts[1]
		// 对version做处理
		if strings.Contains(version, "/go.mod") {
			version = strings.Replace(version, "/go.mod","", -1)
		}
		// only go get the private pkg
		if strings.Contains(pkg, private) {
			target := fmt.Sprintf("%s@%s", pkg, version)
			fmt.Println(fmt.Sprintf("go get -insecure %v", target))
			envCmd := cmd.NewCmd("go", "get", "-v", "-u", "-insecure", target)
			gms := <- envCmd.Start()
			for _, gm := range gms.Stderr {
				fmt.Println(gm)
			}
		}
	}
}
