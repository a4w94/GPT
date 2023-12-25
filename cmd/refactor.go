package cmd

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"gpt/prompt"
	"gpt/util"
	"os"
	"strconv"
	"strings"

	"github.com/fatih/color"
	"github.com/sashabaranov/go-openai"
	"github.com/spf13/cobra"
)

var refactorSetting openai.ChatCompletionRequest

func init() {
	// 在 init 函数中读取 appSetting.json 文件
	loadRefactorAppSettings()

	// 将 temperature 和 topP 绑定为 flag
	// SubCmd1.Flags().Float32Var(&temperature, "temperature", 0.7, "Set the temperature")
	// SubCmd1.Flags().Float32Var(&topP, "top_p", 1.0, "Set the top_p")

}

// 添加子命令
var CmdRefactor = &cobra.Command{
	Use:   "refactor",
	Short: "refactor code",
	Run:   cmdRefactorRun,
}

func cmdRefactorRun(cmd *cobra.Command, args []string) {
	if os.Getenv("OPENAI_API_KEY") == "" {
		fmt.Println("OPENAI_API_KEY is not set")
		return
	}

	// 设置 OpenAI API
	client := openai.NewClient(os.Getenv("OPENAI_API_KEY"))
	req := refactorSetting

	var inputMsg string
	fmt.Println("請輸入您要重構為golang的php代碼：")
	// 使用 bufio.NewScanner 创建一个 Scanner 对象
	scanner := bufio.NewScanner(os.Stdin)
	// 使用 Scan 方法读取输入，直到遇到换行符
	for scanner.Scan() {
		// 获取输入的文本
		inputText := scanner.Text()

		// 遇到換行符，则請輸入start 開始執行
		if inputText == "" {
			fmt.Print("请输入 start 開始執行: ")
			continue
		} else if inputText == "start" {
			break
		} else {
			inputMsg += inputText + "\n"
		}
	}

	var moduleName string
	var moduleNameZhTw string
	fmt.Println("請輸入您要重構為golang的php代碼的模組名稱：")
	fmt.Scanln(&moduleName)
	fmt.Println("請輸入您要重構為golang的php代碼的模組中文名稱：")
	fmt.Scanln(&moduleNameZhTw)
	inputMsg = "module name：" + moduleName + "\n" + "module name in zh-tw" + moduleNameZhTw + "\n" + inputMsg + "\n"

	out, err := util.GetTemplateByString(
		prompt.RefactorTemplate,
		util.Data{
			"input_message": inputMsg,
		},
	)
	if err != nil {
		fmt.Println(err)
		return
	}

	msg := openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleUser,
		Content: out,
	}

	req.Messages = append(req.Messages, msg)
	color.Cyan("gpt is trying to work... ")
	resp, err := client.CreateChatCompletion(context.Background(), req)
	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
	}
	color.Magenta("PromptTokens: " + strconv.Itoa(resp.Usage.PromptTokens) +
		", CompletionTokens: " + strconv.Itoa(resp.Usage.CompletionTokens) +
		", TotalTokens: " + strconv.Itoa(resp.Usage.TotalTokens),
	)
	message := fmt.Sprintf("%s\n\n", resp.Choices[0].Message.Content)

	// Output commit summary data from AI
	color.Yellow("===================Response=======================")
	color.Yellow("\n" + strings.TrimSpace(message) + "\n\n")
	color.Yellow("==================================================")
}

func loadRefactorAppSettings() {
	// 讀取 refactorSetting.json 文件
	file, err := os.Open("./configs/refactorSetting.json")
	if err != nil {
		fmt.Println("Error opening refactorSetting.json file:", err)
		return
	}
	defer file.Close()

	// 解析 JSON
	decoder := json.NewDecoder(file)

	// 将设置应用于全局变量
	err = decoder.Decode(&refactorSetting)
	if err != nil {
		fmt.Println("Error decoding appSetting.json:", err)
		return
	}

}
