package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/GabrielChaves1/easycommit/internal/ai"
	"github.com/GabrielChaves1/easycommit/internal/config"
	"github.com/GabrielChaves1/easycommit/internal/git"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var generateCmd = &cobra.Command{
	Use:     "easycommit",
	Short:   "Go CLI tool to generate commit messages using AI based on git staged changes",
	Version: "0.1.0",
	Run: func(cmd *cobra.Command, args []string) {
		errColor := color.New(color.FgRed, color.Bold)
		successColor := color.New(color.FgGreen, color.Bold)
		messageColor := color.New(color.FgCyan)
		warnColor := color.New(color.FgYellow, color.Italic)

		// Get differences between staged changes and the last commit
		diff, err := git.GetStagedDiff()
		if err != nil {
			if err == git.ErrNoStagedChanges {
				errColor.Println("❌ No staged changes found.")
				return
			}
			errColor.Printf("Error getting staged changes: %s\n", err)
			return
		}

		cfg, err := config.Load()
		if err != nil {
			errColor.Println("Error loading config: ", err)
			return
		}

		var aiClient ai.Agent

		// Initialize the AI client based on the agent type
		switch cfg.AgentType {
		case "openai":
			aiClient, err = ai.NewAgent(
				cfg.AgentType,
				ai.WithAPIKey(cfg.APIKey),
			)
		}

		if err != nil {
			cmd.PrintErr("Error initializing AI client: ", err)
			return
		}

		loading := color.New(color.FgYellow).Add(color.Bold)
		loading.Println("🎲 Generating commit message...")

		// Generate commit message using AI
		message, err := aiClient.GenerateCommitMessage(cmd.Context(), diff, cfg.Language)

		if err != nil {
			cmd.PrintErr("Error generating commit message: ", err)
			return
		}

		successColor.Println("✅ Generated commit message:")
		messageColor.Printf("%s\n", message)
		warnColor.Println("💻 To commit, run: git commit -m \"" + message + "\"")

		promptColor := color.New(color.FgWhite, color.Bold)
		promptColor.Println("Do you want to use this message for your commit? (y/n)")

		// Wait for user input
		var response string
		fmt.Scanln(&response)
		if strings.ToLower(response) == "y" {
			// Commit the changes with the generated message
			commit, err := git.CommitStagedChanges(message)
			if err != nil {
				errColor.Println("Error committing changes: ", err)
				return
			}

			successColor.Println("✅ Changes committed successfully!")
			messageColor.Println(commit)
		} else {
			errColor.Println("❌ Commit aborted.")
		}
	},
}

func Execute() {
	if err := generateCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
