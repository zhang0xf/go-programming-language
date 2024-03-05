# go-programming-language
the go programming language exercise

## reference
* [https://github.com/gopl-zh/gopl-zh.github.com](https://github.com/gopl-zh/gopl-zh.github.com)
* [https://golang-china.github.io/gopl-zh/](https://golang-china.github.io/gopl-zh/)

## Setup

### install go

[go downloads](https://golang.google.cn/dl/)

### confirm installed

`go --version`

### change proxy[optional]

```
go env -w GO111MODULE=on
go env -w GOPROXY=https://goproxy.io,direct
```

### init mod as "exercise"

`go mod init exercise`

### install libraries
`go get -u golang.org/x/net`

`go get gopl.io/ch8/thumbnail`

`sudo apt-get install libbz2-dev`

### build project

`go build`

### clean 

`go clean`
