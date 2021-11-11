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

### build project

`go build`

### clean 

`go clean`
