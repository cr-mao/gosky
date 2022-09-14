package main

import (
	"fmt"
	"gosky/infra/conf"
	"os"
	"time"

	"github.com/spf13/cobra"

	"gosky/cmd"
	"gosky/cmd/job"
	"gosky/cmd/make"
	"gosky/cmd/server"
	"gosky/config"
)

var rootCmd = &cobra.Command{
	Use:   "gosky",
	Short: "this is gosky framework",
}

func init() {
	// 加载 config 目录下的配置信息
	config.Initialize()

	//全局设置时区
	var cstZone, _ = time.LoadLocation(conf.GetString("app.timezone"))
	time.Local = cstZone
}

func main() {

	//注册serve
	rootCmd.AddCommand(server.ServeCmd)
	//注册脚本
	rootCmd.AddCommand(job.JobCmd)
	//注册 make 生成代码模板
	rootCmd.AddCommand(make.CmdMake)
	//注册 sql2struct
	rootCmd.AddCommand(cmd.SqlCmd)

	// 注册全局参数，--env  是什么环境
	cmd.RegisterGlobalFlags(rootCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
