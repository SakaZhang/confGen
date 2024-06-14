package cmd

import (
	"confGenerator/services"
	"confGenerator/tools"
	"github.com/spf13/cobra"
	"log"
	"os"
)

var cfgFile string

var RootCmd = &cobra.Command{
	Use:   "confGen",
	Short: "confGen is a tool to generate configuration files",
	Run: func(cmd *cobra.Command, args []string) {
		if cfgFile == "" {
			log.Fatal("Please specify a configuration file")
			return
		}

		if _, err := os.Stat(cfgFile); os.IsNotExist(err) {
			log.Fatalf("Configuration file does not exist: %s", cfgFile)
			return
		}

		if services.GitDir == "" {
			log.Fatal("Please specify Git directory")
			return
		}

		if info, err := os.Stat(services.GitDir); os.IsNotExist(err) || !info.IsDir() {
			log.Fatalf("The gitpath does not exist or is not a directory: %s", services.GitDir)
			return
		}
		// 根据config.yaml中的定义结构化对应服务的csv
		if err := tools.GenConfFromCfgFile(cfgFile); err != nil {
			log.Fatalf("Something went wrong: %v", err)
			return
		}
	},
}

func init() {

	RootCmd.PersistentFlags().StringVarP(&cfgFile, "file", "f", "", "config file (required)")
	RootCmd.MarkPersistentFlagRequired("file")

	// 添加新的目录参数
	RootCmd.PersistentFlags().StringVarP(&services.GitDir, "directory", "d", "", "gitpath directory (required)")
	RootCmd.MarkPersistentFlagRequired("directory")
}
