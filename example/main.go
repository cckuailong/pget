package main

import (
	"fmt"
	"os"
	"github.com/cckuailong/pget"
)

func main() {
	url := "http://data.ris.ripe.net/rrc14/2019.10/bview.20191031.1600.gz"
	targetDir := "result/"
	output := "test1.gz"
	cli := pget.New()
	if err := cli.Run(url,targetDir, output, 10, 10, "","", true); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	os.Exit(0)
}
