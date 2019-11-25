Pget - parallel file download pkg
=======

## Description

Download using a parallel requests

A pkg from Pget

## Installation

### go get
Install

    $ 

Update

    $ 

## Options

```
  Options:
  -h,  --help                   print usage and exit
  -v,  --version                display the version of pget and exit
  -p,  --procs <num>            split ratio to download file
  -o,  --output <filename>      output file to <filename>
  -d,  --target-dir <path>      path to the directory to save the downloaded file, filename will be taken from url
  -t,  --timeout <seconds>      timeout of checking request in seconds
  -u,  --user-agent <agent>     identify as <agent>
  -r,  --referer <referer>      identify as <referer>
  --check-update                check if there is update available
  --trace                       display detail error messages
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
