package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"gpt/plugins"
	"gpt/util"
	"os"
	"strconv"
	"strings"

	"github.com/fatih/color"
	"github.com/sashabaranov/go-openai"
	"github.com/spf13/cobra"
)

var (
	pluginSetting openai.ChatCompletionRequest
	pluginName    string
	templateRoute string
	promptMap     map[string]interface{}
)

const (
	prefix = "plugins"
)

func init() {
	CmdPlugin.Flags().StringVarP(&pluginName, "", "n", "", "Set the plugin name")
}

// 添加子命令
var CmdPlugin = &cobra.Command{
	Use:   "plugin",
	Short: "explain code",
	Run:   cmdPluginRun,
}

func cmdPluginRun(cmd *cobra.Command, args []string) {
	// 讀取設定檔
	loadAppSettings()

	if os.Getenv("OPENAI_API_KEY") == "" {
		fmt.Println("OPENAI_API_KEY is not set")
		return
	}
	out, err := util.GetTemplateByString(
		templateRoute,
		promptMap,
	)
	if err != nil {
		fmt.Println(err)
		return
	}

	client := openai.NewClient(os.Getenv("OPENAI_API_KEY"))
	req := pluginSetting

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

func loadAppSettings() {
	// 檔案目錄
	fileRoute := fmt.Sprintf("./%s/%s", prefix, pluginName)
	// 設定檔route
	configUrl := fmt.Sprintf("%s/config.json", fileRoute)
	// prompt 模板 route
	templateRoute = fmt.Sprintf("%s.tmpl", pluginName)
	// 讀取 explainSetting.json 文件
	file, err := os.Open(configUrl)
	if err != nil {
		fmt.Println(fmt.Sprintf("Error opening %s file:", configUrl), err)
		return
	}
	defer file.Close()

	// 解析 JSON
	decoder := json.NewDecoder(file)

	err = decoder.Decode(&pluginSetting)
	if err != nil {
		fmt.Println("Error decoding appSetting.json:", err)
		return
	}

	// prompt json route
	promptRoute := fmt.Sprintf("%s/%s.json", fileRoute, pluginName)
	// 讀取 JSON 文件
	promptFile, err := os.ReadFile(promptRoute)
	if err != nil {
		fmt.Println("Error reading config file:", err)
		return
	}
	err = json.Unmarshal(promptFile, &promptMap)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}

	plugins.InitTemplate(pluginName)

}
