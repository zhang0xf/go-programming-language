package chapter4_6

import (
	"html/template"
	"log"
	"os"
)

// usage : ./exercise > data/autoescape.html

// 通过对信任的HTML字符串使用template.HTML类型来抑制这种自动转义的行为。
// 下面的程序演示了两个使用不同类型的相同字符串产生的不同结果：A是一个普通字符串，B是一个信任的template.HTML字符串类型。

// result : 出现在浏览器中的模板输出。我们看到A的黑体标记被转义失效了，但是B没有。

func AutoEscape() {
	const templ = `<p>A: {{.A}}</p><p>B: {{.B}}</p>`
	t := template.Must(template.New("escape").Parse(templ))
	var data struct {
		A string        // untrusted plain text
		B template.HTML // trusted HTML
	}
	data.A = "<b>Hello!</b>"
	data.B = "<b>Hello!</b>"
	if err := t.Execute(os.Stdout, data); err != nil {
		log.Fatal(err)
	}
}
