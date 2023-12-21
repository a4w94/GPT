# cmdexample

`gpt` is a command-line tool that provides various functionalities for code-related tasks. It is built using Cobra, a popular Go library for creating powerful modern CLI applications.

## Features

- **Explain:** Provides explanations for code snippets.
- **Refactor:** Assists in refactoring code.
- **Translate:** Translates code snippets.

## Installation

Before using `gpt`, make sure to set up your environment variables by creating a `.env` file in the project root and populating it with the required values.

```sh
# .env file
API_KEY=your_api_key

Usage
Explain
sh
./gpt explain

This command sends a request to GPT to explain the provided code snippet. Include optional flags to customize the behavior.

Example
```sh
# Explain a code snippet
./gpt explain

Refactor
```sh
./gpt refactor 

Use this command to send a refactoring request to GPT for the provided code.

Translate
```sh
./gpt translate 

Translate the given code snippet using this command.





