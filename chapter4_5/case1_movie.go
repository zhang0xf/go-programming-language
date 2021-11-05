package chapter4_4

import (
	"encoding/json"
	"fmt"
	"log"
)

// 只有导出的结构体成员才会被编码，这也就是我们为什么选择用大写字母开头的成员名称
// 一个结构体成员Tag是和在编译阶段关联到该成员的元信息字符串
// 因为值中含有双引号字符，因此成员Tag一般用原生字符串面值的形式书写
// json开头键名对应的值用于控制encoding/json包的编码和解码的行为, 成员Tag中json对应值的第一部分用于指定JSON对象的名字
// Color成员的Tag还带了一个额外的omitempty选项，表示当Go语言结构体成员为空或零值时不生成该JSON对象（这里false为零值）。果然，Casablanca是一个黑白电影，并没有输出Color成员。
type Movie struct {
	Title  string
	Year   int  `json:"released"`
	Color  bool `json:"color,omitempty"`
	Actors []string
}

var movies = []Movie{
	{Title: "Casablanca", Year: 1942, Color: false,
		Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}},
	{Title: "Cool Hand Luke", Year: 1967, Color: true,
		Actors: []string{"Paul Newman"}},
	{Title: "Bullitt", Year: 1968, Color: true,
		Actors: []string{"Steve McQueen", "Jacqueline Bisset"}},
	// ...
}

func JsonMarshaling() []byte {
	data, err := json.Marshal(movies)
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}
	fmt.Printf("%s\n", data)
	return data
}

// 函数有两个额外的字符串参数用于表示每一行输出的前缀和每一个层级的缩进
func JsonMarshalIndent() []byte {
	data, err := json.MarshalIndent(movies, "", "    ")
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}
	fmt.Printf("%s\n", data)
	return data
}

// 通过定义合适的Go语言数据结构，我们可以选择性地解码JSON中感兴趣的成员
// 当Unmarshal函数调用返回，slice将被只含有Title信息的值填充，其它JSON成员将被忽略
func JsonUnMarshaling(data []byte) {
	var titles []struct{ Title string }
	if err := json.Unmarshal(data, &titles); err != nil {
		log.Fatalf("JSON unmarshaling failed: %s", err)
	}
	fmt.Println(titles) // "[{Casablanca} {Cool Hand Luke} {Bullitt}]"
}
