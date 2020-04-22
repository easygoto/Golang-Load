package study

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Worker struct {
	Id    int     `json:"id"`
	Money float64 `json:"money"`
}

// json 序列化和反序列化
func TestJson(t *testing.T) {
	jsonStr, err := json.Marshal(Worker{1, 3.14})
	if err != nil {
		panic(err)
	}
	_, _ = fmt.Println(string(jsonStr))

	w1 := Worker{}
	w2 := new(Worker)
	err1 := json.Unmarshal(jsonStr, &w1)
	err2 := json.Unmarshal(jsonStr, w2)
	if err1 != nil || err2 != nil {
		panic(err)
	}
	_, _ = fmt.Println(w1, w2)
}

// 指针
func TestPoint(t *testing.T) {
	var ui uint = 123
	pUi := &ui
	ppUi := &pUi
	_, _ = fmt.Println(ui, &ui, *pUi, pUi, &pUi, **ppUi, *ppUi, ppUi)

	// 数组的指针
	var arr = [...]int{1, 2, 3}
	pArr := &arr
	_, _ = fmt.Println(pArr, (*pArr)[0], (*pArr)[1], (*pArr)[2])

	// 指针的数组
	a, b, c := 1, 2, 3
	pBrr := [...]*int{&a, &b, &c}
	_, _ = fmt.Println(pBrr, *pBrr[0], *pBrr[1], *pBrr[2])
}

// 基础内建方法
func TestBasicFn(t *testing.T) {
	list := make([]string, 2)
	list[0] = "root"
	list[1] = "admin"
	_, _ = fmt.Printf("len=%d,cap=%d,value=%v,addr=%p\n", len(list), cap(list), list, &list)

	list = append(list, "user", "test")
	_, _ = fmt.Printf("len=%d,cap=%d,value=%v,addr=%p\n", len(list), cap(list), list, &list)

	listCp := make([]string, 2)
	copy(listCp, list)
	_, _ = fmt.Printf("len=%d,cap=%d,value=%v,addr=%p\n", len(listCp), cap(listCp), listCp, &listCp)
}

// 异常处理
func TestExcept(t *testing.T) {
	defer func() {
		msg := recover()
		_, _ = fmt.Println("======== debug ========:", msg)
	}()
	panic("异常")
}

// 创建变量 —— new
func TestNew(t *testing.T) {
	worker := new(Worker)
	workerCh := new(chan Worker)
	workerMap := new(map[string]*Worker)

	worker.Id = 1
	worker.Money = 3.1415926535897932354626

	go func() {
		*workerCh <- Worker{3, 3.1415926}
		*workerCh <- Worker{4, 2.71828}
		close(*workerCh)
	}()

	*workerMap = map[string]*Worker{} // 需要初始化
	(*workerMap)["root"] = &Worker{0, 1.7e308}
	(*workerMap)["admin"] = &Worker{11, 3.4e38}
	delete(*workerMap, "admin") // 删除元素

	_, _ = fmt.Println(worker, workerCh, workerMap)
}

// 创建变量 —— make
func TestMake(t *testing.T) {
	workerSlice := make([]Worker, 2)
	workerCh := make(chan Worker)
	workerMap := make(map[string]Worker)

	workerSlice[0] = Worker{1, 0.618}
	//workerSlice[2] = Worker{2, 1.41421} // 数组越界

	go func() {
		workerCh <- Worker{3, 3.1415926}
		workerCh <- Worker{4, 2.71828}
		close(workerCh)
	}()

	workerMap["root"] = Worker{0, 1.7e308}
	workerMap["admin"] = Worker{11, 3.4e38}

	_, _ = fmt.Println(workerSlice, workerCh, workerMap)
}
