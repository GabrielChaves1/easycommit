package cmd

import (
	"github.com/GabrielChaves1/easycommit/internal/ai"
	"github.com/GabrielChaves1/easycommit/internal/config"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var setAgentCmd = &cobra.Command{
	Use:   "set-agent",
	Short: "Set the AI agent to use for generating commit messages",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		agentType := ai.AgentType(args[0])

		errColor := color.New(color.FgRed, color.Bold)
		successColor := color.New(color.FgGreen, color.Bold)

		// Validate the agent type
		if !agentType.IsValid() {
			errColor.Println("Invalid agent type. Available types are:")
			for _, at := range ai.ListAgentTypes() {
				cmd.Printf(" - %s\n", at)
			}
			return
		}

		// Initialize the agent based on the type
		switch agentType {
		default:
			apiKey, _ := cmd.Flags().GetString("api-key")
			if apiKey == "" {
				errColor.Println("API key is required")
				errColor.Println("Use --api-key to set the API key.")
				return
			}

			// Initialize the Agent
			_, err := ai.NewAgent(
				agentType.String(),
				ai.WithAPIKey(apiKey),
			)

			if err != nil {
				errColor.Println("Error initializing AI client: ", err)
				return
			}

			// Get the current configuration or return a empty one
			cfg, err := config.Load()
			if err != nil {
				errColor.Println("Error loading config: ", err)
				return
			}

			cfg.AgentType = string(agentType)
			cfg.APIKey = apiKey

			// Save the configuration of the agent
			if err := cfg.Save(); err != nil {
				errColor.Println("Error saving config: ", err)
				return
			}

			successColor.Println("Agent set successfully.")
		}
	},
}

var setLanguageCmd = &cobra.Command{
	Use:   "language [lang]",
	Short: "Set the language for commit messages (e.g., en, pt, pt-BR, es)",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		lang := args[0]

		errColor := color.New(color.FgRed, color.Bold)
		successColor := color.New(color.FgGreen, color.Bold)

		// Load the current configuration
		cfg, err := config.Load()
		if err != nil {
			errColor.Println("Error loading config:", err)
			return
		}

		// Save the new language setting to the yaml file
		cfg.Language = lang
		if err := cfg.Save(); err != nil {
			errColor.Println("Error saving config:", err)
			return
		}

		successColor.Printf("Language set to %s\n", lang)
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
