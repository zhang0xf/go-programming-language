package chapter13_4

import (
	"io"
	"log"
	"os"
)

// 注释见case1_cgo.go

// 下面的bzipper程序，使用我们自己包实现的bzip2压缩命令。它的行为和许多Unix系统的bzip2命令类似。

func Bzipper() {
	w := NewWriter(os.Stdout)
	if _, err := io.Copy(w, os.Stdin); err != nil {
		log.Fatalf("bzipper: %v\n", err)
	}
	if err := w.Close(); err != nil {
		log.Fatalf("bzipper: close: %v\n", err)
	}
}

// 在上面的场景中，我们使用bzipper压缩了/usr/share/dict/words系统自带的词典，从938,848字节压缩到335,405字节。大约是原始数据大小的三分之一。
// 然后使用系统自带的bunzip2命令进行解压。压缩前后文件的SHA256哈希码是相同了，这也说明了我们的压缩工具是正确的。
// (如果你的系统没有sha256sum命令，那么请先按照练习4.2实现一个类似的工具）

// usage :

// $ go build gopl.io/ch13/bzipper
// $ wc -c < /usr/share/dict/words
// 938848
// $ sha256sum < /usr/share/dict/words
// 126a4ef38493313edc50b86f90dfdaf7c59ec6c948451eac228f2f3a8ab1a6ed -
// $ ./bzipper < /usr/share/dict/words | wc -c
// 335405
// $ ./bzipper < /usr/share/dict/words | bunzip2 | sha256sum
// 126a4ef38493313edc50b86f90dfdaf7c59ec6c948451eac228f2f3a8ab1a6ed -
