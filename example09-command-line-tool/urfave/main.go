package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

// 批量声明类型
type (
	Config struct {
		username string
		password string
	}
)

var config Config

func main() {
	// 1. 创建 cli 应用实例（核心入口）
	app := cli.NewApp()
	// 2. 设置应用元信息（会显示在帮助文档中）
	app.Name = "Example"  // 应用名称
	app.Usage = "Example" // 应用用途描述
	// 3. 绑定默认执行逻辑：启动应用时执行 run 函数
	app.Action = run
	// 4. 定义命令行参数（Flags）：支持 --username/-u、--password/-p
	app.Flags = []cli.Flag{
		// 字符串类型参数：--username 或 -u
		cli.StringFlag{
			Name:  "username,u",   // 参数名：长名 + 短名（逗号分隔）
			Usage: "user account", // 参数说明（-h/--help 时显示）
		},
		// 字符串类型参数：--password 或 -p
		cli.StringFlag{
			Name:  "password,p",
			Usage: "user password",
		},
	}
	// 5. 启动应用：解析 os.Args 中的参数，执行 Action 绑定的 run 函数
	app.Run(os.Args)
}

// c *cli.Context：urfave/cli 的上下文对象，存储解析后的所有参数
func run(c *cli.Context) error {
	// 1. 从上下文解析参数，赋值给全局 config
	config = Config{
		username: c.String("username"),
		password: c.String("password"),
	}
	// 2. 执行核心业务逻辑（打印参数）
	return exec()
}

func exec() error {
	//打印全局 config 中存储的用户名和密码
	fmt.Println("username:", config.username)
	fmt.Println("password:", config.password)
	return nil
}
