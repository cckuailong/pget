Pget - parallel file download pkg
=======

## Description

Download using a parallel requests

A pkg from Pget

## Installation

### go get
Install

    $ go get github.com/cckuailong/pget

Update

    $ go get -u github.com/cckuailong/pget

## Func

Run(url,targetDir, output, procs, timeout, userAgent, referer, quiet)
```
  Params:   
  Name:     Type:       Note:
  url       [string]    url to download
  targetDir [string]    path to the directory to save the downloaded file, filename will be taken from url
  output    [string]    output file name
  procs     [int]       split ratio to download file
  timeout   [int]       timeout of checking request in seconds
  userAgent [string]    identify as <userAgent>
  referer   [string]    identify as <referer>
  quiet     [bool]      work on a silent model
```

## Example

```
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
```

## Pget vs Wget

URL: http://ubuntutym2.u-toyama.ac.jp/ubuntu/16.04/ubuntu-16.04-desktop-amd64.iso

Using
```
time wget http://ubuntutym2.u-toyama.ac.jp/ubuntu/16.04/ubuntu-16.04-desktop-amd64.iso
time pget -p 6 http://ubuntutym2.u-toyama.ac.jp/ubuntu/16.04/ubuntu-16.04-desktop-amd64.iso
```
Results

```
wget   3.92s user 23.52s system 3% cpu 13:35.24 total
pget -p 6   10.54s user 34.52s system 25% cpu 2:56.93 total
```
