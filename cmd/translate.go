package cmd

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"gpt/prompt"
	"gpt/util"
	"os"

	"github.com/sashabaranov/go-openai"
	"github.com/spf13/cobra"
)

var translateSetting openai.ChatCompletionRequest

var (
	language string
)

func init() {
	// 在 init 函数中读取 appSetting.json 文件
	loadTranslateAppSettings()
	CmdTranslate.Flags().StringVarP(&language, "lang", "l", "en", "输入语言")

	// SubCmd1.Flags().Float32Var(&temperature, "temperature", 0.7, "Set the temperature")
	// SubCmd1.Flags().Float32Var(&topP, "top_p", 1.0, "Set the top_p")

}

// 添加子命令
var CmdTranslate = &cobra.Command{
	Use:   "translate",
	Short: "translate code",
	Run:   cmdTranslateRun,
}

func cmdTranslateRun(cmd *cobra.Command, args []string) {
	if os.Getenv("OPENAI_API_KEY") == "" {
		fmt.Println("OPENAI_API_KEY is not set")
		return
	}

	// 设置 OpenAI API
	client := openai.NewClient(os.Getenv("OPENAI_API_KEY"))
	req := refactorSetting
	var lang string
	fmt.Println("請輸入要翻譯的語言：")
	scanner := bufio.NewScanner(os.Stdin)
	// 使用 Scan 方法读取输入，直到遇到换行符
	for scanner.Scan() {
		// 获取输入的文本
		inputText := scanner.Text()

		// 如果输入为空，则重新提示用户输入
		if inputText == "" {
			fmt.Print("请输入要翻譯的語言: ")
			continue
		} else {
			lang = inputText
			break
		}
	}

	var inputMsg string
	fmt.Println("請輸入您要翻譯的文字：")
	// 使用 Scan 方法读取输入，直到遇到换行符
	// for scanner.Scan() {
	// 	// 获取输入的文本
	// 	inputText := scanner.Text()

	// 	// 遇到換行符，则請輸入start 開始執行
	// 	if inputText == "" {
	// 		continue
	// 	} else {
	// 		inputMsg += inputText + "\n"
	// 	}
	// }
	for {
		t, _, _ := bufio.NewReader(os.Stdin).ReadLine()
		if string(t) != "" {
			inputMsg += string(t) + "\n"
		}
	}
	fmt.Println("inputMsg", inputMsg)
	panic(1)
	out, err := util.GetTemplateByString(
		prompt.TranslationTemplate,
		util.Data{
			"use_language":  lang,
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
	fmt.Println("gpt working...")
	resp, err := client.CreateChatCompletion(context.Background(), req)
	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
	}
	fmt.Printf("%s\n\n", resp.Choices[0].Message.Content)

}

func loadTranslateAppSettings() {
	// 讀取 translateSetting.json 文件
	file, err := os.Open("./configs/translateSetting.json")
	if err != nil {
		fmt.Println("Error opening translateSetting.json file:", err)
		return
	}
	defer file.Close()

	// 解析 JSON
	decoder := json.NewDecoder(file)

	err = decoder.Decode(&translateSetting)
	if err != nil {
		fmt.Println("Error decoding appSetting.json:", err)
		return
	}

	// 将设置应用于全局变量
	fmt.Println("translateSetting", translateSetting)
}
