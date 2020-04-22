package study

import (
	"fmt"
	"testing"
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

func TestConst(t *testing.T) {
	const (
		c0 = iota
		c1
		c2 = 10
		c3
	)
	const (
		Sun = iota
		Mon
		Tue
		Wed
		Thu
		Fri
		Sat
	)
	const c10 = iota
	const c11 = iota
	const c12 = iota
	_, _ = fmt.Printf("c0=%d, c1=%d, c2=%d, c3=%d\n", c0, c1, c2, c3)
	_, _ = fmt.Printf("c10=%d, c11=%d, c12=%d\n", c10, c11, c12)
	_, _ = fmt.Printf("Sun=%d, Mon=%d, Tue=%d, Wed=%d, Thu=%d, Fri=%d, Sat=%d\n", Sun, Mon, Tue, Wed, Thu, Fri, Sat)
}

func TestType(t *testing.T) {
	var ui uint = 0xffffffffffffffff
	var b byte = '\x41'
	var i, f, c, s, cx = 99999999, 3.099128981278123123, '\u548c', "Welcome", 12 + 3i
	var pui, pf, pcx, ps = &ui, &f, &cx, &s
	_, _ = fmt.Printf("%-24T: %t\n", false, false)
	_, _ = fmt.Printf("%-24T: %d, %o, %x, %X\n", i, i, i, i, i)
	_, _ = fmt.Printf("%-24T: %d, %X\n", ui, ui, ui)
	_, _ = fmt.Printf("%-24T: %.6f\n", f, f)
	_, _ = fmt.Printf("%-24T: %f\n", cx, cx)
	_, _ = fmt.Printf("%-24T: %c(%d)\n", b, b, b)
	_, _ = fmt.Printf("%-24T: %c(%x)\n", c, c, c)
	_, _ = fmt.Printf("%-24T: %s\n", s, s)
	_, _ = fmt.Printf("%-24T: %p = %d\n", pui, pui, *pui)
	_, _ = fmt.Printf("%-24T: %p = %f\n", pf, pf, *pf)
	_, _ = fmt.Printf("%-24T: %p = %f\n", pcx, pcx, *pcx)
	_, _ = fmt.Printf("%-24T: %p = %s\n", ps, ps, *ps)
	_, _ = fmt.Println()

	arr := [...]User{{11, "root", true}, {12, "test", false}, {21, "user", false}}
	mp := map[string]string{"a": "apple", "b": "banana", "d": "dog"}
	st := struct{ id int }{1}
	user := User{id: 1, name: "admin", isAdmin: true}
	mod := Admin{56}
	ch := make(chan User)
	fn := func(msg string) string { return "msg: " + msg }
	_, _ = fmt.Printf("%-24T: %v\n", arr, arr)
	_, _ = fmt.Printf("%-24T: %v\n", mp, mp)
	_, _ = fmt.Printf("%-24T: %v, len=%d\n", s[1:5], s[1:5], len(s[1:5]))
	_, _ = fmt.Printf("%-24T: %v, len=%d, cap=%d\n", arr[1:], arr[1:], len(arr[1:]), cap(arr[1:]))
	_, _ = fmt.Printf("%-24T: %#v\n", st, st)
	_, _ = fmt.Printf("%-24T: %+v\n", user, user)
	_, _ = fmt.Printf("%-24T: %+v, id: %d\n", mod, mod, mod.getId())
	_, _ = fmt.Printf("%-24T: %#v\n", ch, ch)
	_, _ = fmt.Printf("%-24T: %#v, %s\n", fn, &fn, fn("hello"))
}
