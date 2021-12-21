package chapter12_7

import (
	"fmt"
	"net/http"
	"reflect"
	"strconv"
	"strings"
)

// 下面的Unpack函数主要完成三件事情。
// 第一，它调用req.ParseForm()来解析HTTP请求。然后，req.Form将包含所有的请求参数，不管HTTP客户端使用的是GET还是POST请求方法。
// 下一步，Unpack函数将构建每个结构体成员有效参数名字到成员变量的映射。如果结构体成员有成员标签的话，有效参数名字可能和实际的成员名字不相同。
// reflect.Type的Field方法将返回一个reflect.StructField，里面含有每个成员的名字、类型和可选的成员标签等信息。
// 其中成员标签信息对应reflect.StructTag类型的字符串，并且提供了Get方法用于解析和根据特定key提取的子串，例如这里的http:"..."形式的子串。
// 最后，Unpack遍历HTTP请求的name/value参数键值对，并且根据更新相应的结构体成员。
// 回想一下，同一个名字的参数可能出现多次。如果发生这种情况，并且对应的结构体成员是一个slice，那么就将所有的参数添加到slice中。
// 其它情况，对应的成员值将被覆盖，只有最后一次出现的参数值才是起作用的。

// Unpack populates the fields of the struct pointed to by ptr
// from the HTTP request parameters in req.
func Unpack(req *http.Request, ptr interface{}) error {
	if err := req.ParseForm(); err != nil {
		return err
	}

	// Build map of fields keyed by effective name.
	fields := make(map[string]reflect.Value)
	v := reflect.ValueOf(ptr).Elem() // the struct variable
	for i := 0; i < v.NumField(); i++ {
		fieldInfo := v.Type().Field(i) // a reflect.StructField
		tag := fieldInfo.Tag           // a reflect.StructTag
		name := tag.Get("http")
		if name == "" {
			name = strings.ToLower(fieldInfo.Name)
		}
		fields[name] = v.Field(i)
	}

	// Update struct field for each parameter in the request.
	for name, values := range req.Form {
		f := fields[name]
		if !f.IsValid() {
			continue // ignore unrecognized HTTP parameters
		}
		for _, value := range values {
			if f.Kind() == reflect.Slice {
				elem := reflect.New(f.Type().Elem()).Elem()
				if err := populate(elem, value); err != nil {
					return fmt.Errorf("%s: %v", name, err)
				}
				f.Set(reflect.Append(f, elem))
			} else {
				if err := populate(f, value); err != nil {
					return fmt.Errorf("%s: %v", name, err)
				}
			}
		}
	}
	return nil
}

// 填充
// populate函数小心用请求的字符串类型参数值来填充单一的成员v（或者是slice类型成员中的单一的元素）。
func populate(v reflect.Value, value string) error {
	switch v.Kind() {
	case reflect.String:
		v.SetString(value)

	case reflect.Int:
		i, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return err
		}
		v.SetInt(i)

	case reflect.Bool:
		b, err := strconv.ParseBool(value)
		if err != nil {
			return err
		}
		v.SetBool(b)

	default:
		return fmt.Errorf("unsupported kind %s", v.Type())
	}
	return nil
}
