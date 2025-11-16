package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockAgent_createPromptVersionCmd = &cobra.Command{
	Use:   "create-prompt-version",
	Short: "Creates a static snapshot of your prompt that can be deployed to production.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockAgent_createPromptVersionCmd).Standalone()

	bedrockAgent_createPromptVersionCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier to ensure that the API request completes no more than one time.")
	bedrockAgent_createPromptVersionCmd.Flags().String("description", "", "A description for the version of the prompt.")
	bedrockAgent_createPromptVersionCmd.Flags().String("prompt-identifier", "", "The unique identifier of the prompt that you want to create a version of.")
	bedrockAgent_createPromptVersionCmd.Flags().String("tags", "", "Any tags that you want to attach to the version of the prompt.")
	bedrockAgent_createPromptVersionCmd.MarkFlagRequired("prompt-identifier")
	bedrockAgentCmd.AddCommand(bedrockAgent_createPromptVersionCmd)
}
