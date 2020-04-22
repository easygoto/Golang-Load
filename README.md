# 目录

- [1 学习笔记](README.md#1-学习笔记)
    - [1.1 基本语法](README.md#11-基本语法)
        - [1.1.1 注意事项](README.md#111-注意事项)
        - [1.1.2 数据类型](README.md#112-数据类型)
    - [1.2 基础方法](README.md#12-基础方法)
    - [1.3 面向对象](README.md#13-面向对象)
    - [1.4 并发](README.md#14-并发)
    - [1.5 网络编程](README.md#15-网络编程)
- [2 案例](README.md#2-案例)
    - [2.1 并发思想归并排序](README.md#21-并发思想归并排序)
    - [2.2 抢红包算法](README.md#22-抢红包算法)
    - [2.3 哈希表数据结构](README.md#23-哈希表数据结构)

# 1 学习笔记

## 1.1 基本语法

### 1.1.1 注意事项

- 首字母大写的字段, 方法才会被外部调用
- func
    - 不定参数 myfunc(args ...int), 外部调用传输 slice 可以用 `myfunc(slice[:]...)`
    - 不支持重载函数
- select
    - 不需要 break 退出 case, 明确添加 fallthrough, 才会继续执行下一个 case
    - case 可以有多个值, 用 "," 隔开
- Printf
    - `%p` : 指针值
    - `%T` : 类型
    - `%t` : 布尔值
    - `%v / %+v / %#v` : (原始 / json / id 和 json) 形式内容
    - `%b / %o / %d / %x / %X` : (二 / 八 / 十进制 / 十六 / 大写的十六) 进制整型
    - `%f / %g / %e / %E` : (普通 / 精确 / 科学计数法 / 科学计数法) 浮点数, 复数
    - `%c / %s / %U / %q` : (字符 / 普通 / unicode / 带有引号) 字符串

### 1.1.2 数据类型

[点击查看测试代码](test/study/01_type_test.go)

### 1.1.3 module

- `go mod init modname` : 初始化模块
- `go mod graph` : 依赖
- `go mod why` : 依赖关系
- `go list -m all` : 列出依赖关系
- `go get` : 获取包
- `go build` : 此时也会添加依赖
- `go mod download` : 下载依赖
- `go mod tidy` : 整理依赖
- `go mod verify` : 验证依赖
- `go mod edit` : 编辑 mod 的属性, help
- `go mod vendor` : 依赖放到 vendor 中

## 1.2 基础方法

- 创建
    - make: 返回引用类型
    - new: 返回指针类型
- 基础方法
    - append: `切片`添加, 扩容
    - copy: `切片`拷贝
    - delete: 删除 map 中元素
    - close: 关闭 chan, 禁止写入
    - len: 实际长度, 支持 string, array, slice, map, chan
    - cap: 容器容量, 支持 array, slice, chan
- 异常处理
    - panic: 中断程序, 抛出信息
    - recover: 接收抛出的信息
    - defer: 关键字, 最后一定要做的事情
- 指针
    - 基本: 不支持指针++, --, 指向指针的指针和 C 语言一致, 没有函数指针
    - 指针数组: 一个数组里全是指针
    - 数组指针: 指向数组的指针
- json
    - Marshal: 序列化, 结构体中只会序列公用的字段, 结构体可以指定字段的映射
    - Unmarshal: 反序列化, 需要传递反序列化的原型

[点击查看测试代码](test/study/02_api_test.go)

## 1.3 面向对象

1. `静态` 文件中函数外部定义的变量, 类似于静态变量
1. `封装` type struct 可以作为类使用 `type Apple struct {}`
1. `多态` type interface 作为接口, 类不需要声明直接实现方法, 实现方法后就是该接口的子类
1. `构造` 构造方法需要对外提供若干函数, 来实现构造函数的功能 `func (apple Apple) NewApple() *Apple {return &Apple{}}`
1. `继承` 可以把要继承的结构体作为一个成员变量`组合`到结构体中, 成员变量默认访问自己的, 没有就访问 `父类` 的
1. 接口定义在任何地方都没有问题, 始终会被继承, 优先使用本包的接口
1. 对外提供的方法和结构体名称和文件名没有任何关系, 和文件夹名, 包名, 结构体名, 函数名有直接关系

[点击查看测试代码](test/study/04_oo_test.go)

- [源文件 a.go](src/demo/case/a.go)
- [源文件 b.go](src/demo/case/b.go)
- [源文件 c.go](src/demo/box/c.go)

## 1.4 并发

- go: 启动一个 goroutine, 协程没有运行的原因可能是主程序结束, 协程也跟着结束
- chan: 协程之间的通信, 必须在 goroutine 中执行
    - `func(c chan int) { }` 读写均可的 chan
    - `func(c <-chan int) { }` 只读的 chan
    - `func(c chan<- int) { }` 只写的 chan
    - 作为一个好的习惯, 写完 chan 需要 close
- select: 阻塞式从 chan 读取数据, 自上向下依次读取
- WaitGroup: 协程同步
    - Add(delta): 添加协程记录, delta 表示需要几次 `Done` 才能解除
    - Done(): 移除一个协程记录
    - Wait(): 等待所有协程完成, 若总 delta 不等于总 Done 的数量, 会发生死锁

[点击查看测试代码](test/study/03_goroutine_test.go)

## 1.5 网络编程

[点击查看测试代码](test/study/05_net_test.go)

# 2 案例

> 并发思想, 把大而复杂的事情拆分成小而简单的事情, 并行去做小而简单的事情, 以此达到计算机性能最大的发挥

## 2.1 并发思想归并排序

[点击查看源码](src/algo/mergesort)

## 2.2 抢红包算法

[点击查看源码](src/algo/luckymoney)

## 2.3 哈希表数据结构

[点击查看源码](src/ds/hashmap)
