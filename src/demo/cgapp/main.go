package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"demo/cgapp/cg"
	"demo/cgapp/ipc"
)

var centerClient *cg.CenterClient

func startCenterService() error {
	server := ipc.NewIpcServer(&cg.CenterServer{})
	client := ipc.NewIpcClient(server)
	centerClient = &cg.CenterClient{Client: client}
	return nil
}

func Help(args []string) int {
	_, _ = fmt.Println(` 
 Commands:
 login <username> <level> <exp>
 logout <username>
 send <message>
 list
 quit(q)
 help(h)
 `)
	return 0
}

func Quit(args []string) int {
	return 1
}

func Login(args []string) int {
	if len(args) != 4 {
		_, _ = fmt.Println("USAGE: login <username> <level> <exp>")
		return 0
	}
	level, err := strconv.Atoi(args[2])
	if err != nil {
		_, _ = fmt.Println("Invalid Parameter: <level> should be an integer.")
		return 0
	}
	exp, err := strconv.Atoi(args[3])
	if err != nil {
		_, _ = fmt.Println("Invalid Parameter: <exp> should be an integer.")
		return 0
	}

	player := cg.NewPlayer()
	player.Name = args[1]
	player.Level = level
	player.Exp = exp
	err = centerClient.AddPlayer(player)
	if err != nil {
		_, _ = fmt.Println("Failed adding player", err)
	}
	return 0
}

func Logout(args []string) int {
	if len(args) != 2 {
		_, _ = fmt.Println("USAGE: logout <username>")
		return 0
	}
	_ = centerClient.RemovePlayer(args[1])
	return 0
}

func ListPlayer(args []string) int {
	players, err := centerClient.ListPlayer("")
	if err != nil {
		_, _ = fmt.Println("Failed. ", err)
	} else {
		for i, v := range players {
			_, _ = fmt.Printf("No.%5d : %+v\n", i+1, v)
		}
	}
	return 0
}

func Send(args []string) int {
	message := strings.Join(args[1:], " ")
	err := centerClient.Broadcast(message)
	if err != nil {
		_, _ = fmt.Println("Failed.", err)
	}
	return 0
}

// 命令和函数名的映射
func GetCommandHandlers() map[string]func(args []string) int {
	return map[string]func([]string) int{
		"help":   Help,
		"h":      Help,
		"quit":   Quit,
		"q":      Quit,
		"login":  Login,
		"logout": Logout,
		"list":   ListPlayer,
		"send":   Send,
	}
}

func main() {
	_, _ = fmt.Println("Casual Game Server Solution")
	Help(nil)
	_ = startCenterService()

	r := bufio.NewReader(os.Stdin)
	handlers := GetCommandHandlers()
	for {
		_, _ = fmt.Print("Command > ")
		b, _, _ := r.ReadLine()
		line := string(b)
		tokens := strings.Split(line, " ")

		// 用户输入命令后, 执行失败直接退出
		if handler, ok := handlers[tokens[0]]; ok {
			ret := handler(tokens)
			if ret != 0 {
				break
			}
		} else {
			_, _ = fmt.Println("Unknown command:", tokens[0])
		}
	}
}
