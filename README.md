# 🚀 EasyCommit

![image](https://github.com/user-attachments/assets/1d7edad1-09a1-4914-b180-a2d4048ec711)

**EasyCommit** is a Golang command-line tool that uses AI to generate concise, descriptive, and conventional Git commit messages for you, you just need to run a single command `easycommit` after staging your changes.

## 📜 Features

- **AI-Powered**: Uses OpenAI to generate commit messages.
- **Multi-Agent Support**: Choose between multiple AI agents to generate commit messages.
- **Easy to Use**: Just run `easycommit` after staging your changes.
- **Supports Multiple Languages**: Works with any programming language.

## 📦 Installation

Requires Go 1.20+ and Git.

```bash
go install github.com/GabrielChaves1/easycommit@latest
```

Or download the binary from [Releases](https://github.com/GabrielChaves1/easycommit/releases).

## 🛠️ Configuration Commands

### Set the AI Agent

Configure which AI agent will be used to generate commit messages.  
Now you can choose between OpenAI and Gemini.

```bash
# Set the AI agent to OpenAI
easycommit config set-agent openai --api-key YOUR_OPENAI_API_KEY

# Set the AI agent to Gemini
easycommit config set-agent gemini --api-key YOUR_GEMINI_API_KEY
```

### Set the Commit Message Language

Set the language in which commit messages will be generated.

```bash
easycommit config language en
```

**en**: The language for the commit message.

- Replace en with your preferred language (e.g., pt for Portuguese, es for Spanish, etc).

You can change this at any time to generate commit messages in a different language.

## 🖥️ Example Workflow

Generate a commit message:

```bash
easycommit config set-agent openai --api-key YOUR_OPENAI_API_KEY
easycommit config language en
git add .
easycommit
```

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Contributing

Contributions are welcome! Please read the [CONTRIBUTING.md](docs/CONTRIBUTING.md) for details on how to contribute to this project.
