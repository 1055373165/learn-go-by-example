/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "go-gopher-grpc",
	Short: "gRPC app in Go",
	Long:  `gRPC application written in Go.`,
}

// Execute 会将所有子命令添加到根命令中，并适当设置标志。此操作由 main.main() 调用。只需对 rootCmd 执行一次。
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)
	// 您可以在此定义标记和配置设置。Cobra 支持持久标志，如果在此定义了这些标志，它们将成为应用程序的全局标志。
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.go-gopher-grpc.yaml")
	// Cobra 还支持本地标志，只有在直接调用此操作时才会运行。
	rootCmd.Flags().BoolP("toggle", "t", false, "help message for toggle")
	// rootCmd.AddCommand(configCmd)
	// rootCmd.AddCommand(serverCmd)
	// rootCmd.AddCommand(clientCmd)
}

// 如果设置了 initConfig，则读取配置文件和 ENV 变量。
func initConfig() {
	if cfgFile != "" {
		// use config file from the flag
		viper.SetConfigFile(cfgFile)
	} else {
		// find home dir
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)
		// 在主目录中搜索名为“ . go-gopher-grpc”的配置文件(不带扩展名)。
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".go-gopher-grpc")
	}
	viper.AutomaticEnv() // 读入匹配的环境变量
	// if a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
