package cmd

import (
	"github.com/GabrielChaves1/easycommit/internal/ai"
	"github.com/GabrielChaves1/easycommit/internal/config"
	"github.com/spf13/cobra"
)

var setAgentCmd = &cobra.Command{
	Use:   "set-agent",
	Short: "Set the AI agent to use for generating commit messages",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		agentType := ai.AgentType(args[0])

		// Validate the agent type
		if !agentType.IsValid() {
			cmd.Println("Invalid agent type. Available types are:")
			for _, at := range ai.ListAgentTypes() {
				cmd.Printf(" - %s\n", at)
			}
			return
		}

		// Initialize the agent based on the type
		switch agentType {
		case ai.AgentTypeOpenAI:
			apiKey, _ := cmd.Flags().GetString("api-key")
			if apiKey == "" {
				cmd.Println("API key is required for OpenAI agent")
				cmd.Println("Use --api-key to set the API key.")
				return
			}

			// Initialize the OpenAI agent
			_, err := ai.NewAgent(
				"openai",
				ai.WithAPIKey(apiKey),
			)

			if err != nil {
				cmd.Println("Error initializing OpenAI client: ", err)
				return
			}

			// Get the current configuration or return a empty one
			cfg, err := config.Load()
			if err != nil {
				cmd.Println("Error loading config: ", err)
				return
			}

			cfg.AgentType = string(agentType)
			cfg.APIKey = apiKey

			// Save the configuration of the agent
			if err := cfg.Save(); err != nil {
				cmd.Println("Error saving config: ", err)
				return
			}

			cmd.Println("Agent set successfully.")
		}
	},
}

var setLanguageCmd = &cobra.Command{
	Use:   "language [lang]",
	Short: "Set the language for commit messages (e.g., en, pt, pt-BR, es)",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		lang := args[0]

		cfg, err := config.Load()
		if err != nil {
			cmd.Println("Error loading config:", err)
			return
		}

		cfg.Language = lang
		if err := cfg.Save(); err != nil {
			cmd.Println("Error saving config:", err)
			return
		}
		cmd.Printf("Language set to %s\n", lang)
	},
}

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Manage easycommit configuration",
}

func init() {
	generateCmd.AddCommand(configCmd)
	configCmd.AddCommand(setAgentCmd)
	configCmd.AddCommand(setLanguageCmd)
	setAgentCmd.Flags().StringP("api-key", "k", "", "API key for OpenAI agent")
}
