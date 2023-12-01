package main

import (
	"fmt"

	"gpt/cmd"

	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

type Config struct {
	Model            string
	Prompt           string
	MaxTokens        int
	Temperature      float64
	TopP             float64
	FrequencyPenalty float64
	PresencePenalty  float64
}

var config Config

func init() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
		return
	}
}

func main() {
	rootCmd := &cobra.Command{Use: "cmdexample"}

	// 将子命令添加到根命令
	rootCmd.AddCommand(cmd.SubCmd1)
	rootCmd.AddCommand(cmd.CmdExplain)

	// 运行 Cobra 命令
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}