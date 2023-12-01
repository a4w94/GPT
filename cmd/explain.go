package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/sashabaranov/go-openai"
	"github.com/spf13/cobra"
)

var explainSetting openai.ChatCompletionRequest

func init() {
	// 在 init 函数中读取 appSetting.json 文件
	loadAppSettings()

	// 将 temperature 和 topP 绑定为 flag
	// SubCmd1.Flags().Float32Var(&temperature, "temperature", 0.7, "Set the temperature")
	// SubCmd1.Flags().Float32Var(&topP, "top_p", 1.0, "Set the top_p")

}

// 添加子命令
var CmdExplain = &cobra.Command{
	Use:   "explain",
	Short: "explain code",
	Run:   cmdExplainRun,
}

func cmdExplainRun(cmd *cobra.Command, args []string) {
	fmt.Println("cmdExplainRun")
	fmt.Println(os.Getenv("OPENAI_API_KEY"))
	if os.Getenv("OPENAI_API_KEY") == "" {
		fmt.Println("OPENAI_API_KEY is not set")
		return
	}
	client := openai.NewClient(os.Getenv("OPENAI_API_KEY"))

	req := explainSetting
	fmt.Println("請輸入要解釋的code")
	fmt.Println("---------------------")
	fmt.Print("> ")
	// 一次性读取所有输入
	var codeInput string
	fmt.Scanln(&codeInput)

	fmt.Println(codeInput)
	// 处理 codeInput，你可以将其作为一个字符串传递给 generatedPrompt 函数
	content := generatedPrompt(string(codeInput))
	fmt.Println("content", content)
	req.Messages = append(req.Messages, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleUser,
		Content: content,
	})
	resp, err := client.CreateChatCompletion(context.Background(), req)
	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return
	}
	fmt.Printf("%s\n\n", resp.Choices[0].Message.Content)
	req.Messages = append(req.Messages, resp.Choices[0].Message)
	fmt.Print("> ")

}

func loadAppSettings() {
	// 讀取 explainSetting.json 文件
	file, err := os.Open("./configs/explainSetting.json")
	if err != nil {
		fmt.Println("Error opening explainSetting.json file:", err)
		return
	}
	defer file.Close()

	// 解析 JSON
	decoder := json.NewDecoder(file)

	err = decoder.Decode(&explainSetting)
	if err != nil {
		fmt.Println("Error decoding appSetting.json:", err)
		return
	}

	// 将设置应用于全局变量
	fmt.Println("explainSetting", explainSetting)
}

func generatedPrompt(input string) string {
	prompt := "Please explain the following code in a short sentence:\n"
	prompt += "```\n"
	prompt += input
	prompt += "```\n"
	prompt += "Explanation:\n"
	return prompt
}
