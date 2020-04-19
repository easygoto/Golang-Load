# 基本语法

## 数据类型

```go
package main

import (
    "fmt"
)

type Model interface {
    getId() int
}

type Admin struct {
    id int
}

func (admin Admin) getId() int {
    return admin.id
}

type User struct {
    id      int
    name    string
    isAdmin bool
}

func (user User) getId() int {
    return user.id
}

func main() {
    var ui uint = 0xffffffffffffffff
    var i, f, c, s, cx = 99999999, 3.099128981278123123, '和', "Welcome", 12 + 3i
    var pui, pf, pcx, ps = &ui, &f, &cx, &s
    fmt.Printf("%-24T: %t\n", false, false)
    fmt.Printf("%-24T: %d, %o, %x, %X\n", i, i, i, i, i)
    fmt.Printf("%-24T: %d, %X\n", ui, ui, ui)
    fmt.Printf("%-24T: %.6f\n", f, f)
    fmt.Printf("%-24T: %f\n", cx, cx)
    fmt.Printf("%-24T: %c(%d)\n", c, c, c)
    fmt.Printf("%-24T: %s\n", s, s)
    fmt.Printf("%-24T: %p = %d\n", pui, pui, *pui)
    fmt.Printf("%-24T: %p = %f\n", pf, pf, *pf)
    fmt.Printf("%-24T: %p = %f\n", pcx, pcx, *pcx)
    fmt.Printf("%-24T: %p = %s\n", ps, ps, *ps)
    fmt.Println()

    arr := [...]User{{11, "root", true}, {12, "test", false}, {21, "user", false}}
    mp := map[string]string{"a": "apple", "b": "banana", "d": "dog"}
    st := struct{ id int }{1}
    user := User{id: 1, name: "admin", isAdmin: true}
    mod := Admin{56}
    ch := make(chan User)
    fn := func(msg string) string { return "msg: " + msg }
    fmt.Printf("%-24T: %v\n", arr, arr)
    fmt.Printf("%-24T: %v\n", mp, mp)
    fmt.Printf("%-24T: %v, len=%d\n", s[1:5], s[1:5], len(s[1:5]))
    fmt.Printf("%-24T: %v, len=%d, cap=%d\n", arr[1:], arr[1:], len(arr[1:]), cap(arr[1:]))
    fmt.Printf("%-24T: %#v\n", st, st)
    fmt.Printf("%-24T: %+v\n", user, user)
    fmt.Printf("%-24T: %+v, id: %d\n", mod, mod, mod.getId())
    fmt.Printf("%-24T: %#v\n", ch, ch)
    fmt.Printf("%-24T: %#v, %s\n", fn, &fn, fn("hello"))
}
```

### Printf

- %v: 自然形式内容
- %+v: 自然形式 json
- %#v: 自然形式 id 和 json
- %T: 类型
- %t: 布尔值
- %b: 二进制
- %o: 八进制
- %d: 十进制整型
- %x: 十六进制
- %X: 大写的十六进制
- %x: base64 编码
- %f: 浮点数, 复数
- %g: 精确浮点数, 复数
- %e: 科学计数法
- %E: 科学计数法
- %c: unicode 字符
- %s: 字符串
- %U: unicode 字符串
- %q: 带有引号
- %p: 指针值

## 并发

> 并发很简单, 只需要在函数之前加入关键字 `go` 即可

```go
package main

import (
    "fmt"
    "net/http"
)

func main() {
    go HttpServer()
    ch := make(chan string)
    for i := 0; i < 5000; i++ {
        go PrintMessage(i, ch)
    }

    for {
        msg := <-ch
        fmt.Println(msg)
    }
}

func PrintMessage(id int, ch chan string) {
    ch <- fmt.Sprintf("the message from %d", id)
}

func HttpServer() {
    http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
        _, _ = fmt.Fprintf(writer, "<h1>Welcome! %s!</h1>", request.FormValue("name"))
    })
    _ = http.ListenAndServe(":8888", nil)
}
```

