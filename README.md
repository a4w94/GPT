# gpt

`gpt` is a command-line tool that provides various functionalities for code-related tasks. It is built using Cobra, a popular Go library for creating powerful modern CLI applications.

## Features

- **Summary:** Provides the article summary.
- **Refactor:** Assists in refactoring code.
- **Translate:** Translates the input paragraph.

## Installation

Before using `gpt`, make sure to set up your environment variables by creating a `.env` file in the project root and populating it with the required values.

```sh
# .env file
API_KEY=your_api_key
```

## Usage

### Translate

```sh
 go run main.go plugin -n translate
```
Use this command to send a translate request to GPT for the provided paragraph.

### Summary

```sh
 go run main.go plugin -n summary
```
Use this command to send a summary request to GPT for the provided article.

### Refactor
```sh
./gpt refactor
```
or
```sh
go run main.go refactor
```

Use this command to send a refactoring request to GPT for the provided code.


### 自行新增

```markdown
- 步驟 1: 複製目錄
- 步驟 2: 更改名稱為您要的功能名稱並確保目錄格式

    plugins/
    └── {function_name}/
        ├── config.json
        ├── {function_name}.json
        └── {function_name}.tmpl

- 步驟 3: 調整config.json參數
- 步驟 4: 依照您的需求修改{function_name}.tmpl並確保{function_name}.json包含{function_name}.tmpl所有key值
- 步驟 5: 執行 `go run main.go plugin -n {function_name}`
```

# Demo

## 1.
```php
```
## 2. Demo CodeGPT
### 安裝
```bash
brew install codegpt
```
or
```bash
go install github.com/appleboy/CodeGPT/cmd/codegpt@latest
```
### 設定api_key
```bash
codegpt config set openai.api_key sk-xxxxxxx
```
### 設定model (不一定要)
```bash
codegpt config set openai.model gpt-3.5-turbo-16k
```
### 執行
```bash
git add .
```

請gpt幫忙寫commit
```bash
Codegpt commit –preview //
```

請gpt幫忙code review 並翻譯
```bash
codegpt commit --lang zh-tw –preview
```

提交剛剛的commit
```bash
codegpt commit –amend
```

使用自訂格式commit
```bash
codegpt commit --preview --template_file ./commit_message.tpl --template_vars_file ./commit_vars.env
```







