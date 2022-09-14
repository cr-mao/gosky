package cmd

import "github.com/spf13/cobra"

var Env string

// 注册全局参数，--env
func RegisterGlobalFlags(rootCmd *cobra.Command) {
	rootCmd.PersistentFlags().StringVarP(&Env, "env", "e", "local", "load `.config.yaml` file, example: --env=debug will use `debug.config.yaml` file")
}
