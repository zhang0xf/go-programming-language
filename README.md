# go-programming-language
the go programming language exercise

## Setup

### install go

[go downloads](https://golang.google.cn/dl/)

### confirm installed

`go --version`

### change proxy for vscode

```
go env -w GO111MODULE=on
go env -w GOPROXY=https://goproxy.io,direct
```

### init mod as "exercise"

`go mod init exercise`

### install libraries
`go get -u golang.org/x/net`

`go get gopl.io/ch8/thumbnail`

`sudo apt-get install libbz2-dev` : 用于go调用c库（windows环境复杂，很多库安装不方便。另见：[WSL + vscode](https://github.com/zhang0xf/md/blob/main/WSL/WSL.md)）

### build project

`go build`

### clean 

`go clean`
