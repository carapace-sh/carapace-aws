package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockAgent_deletePromptCmd = &cobra.Command{
	Use:   "delete-prompt",
	Short: "Deletes a prompt or a version of it, depending on whether you include the `promptVersion` field or not.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockAgent_deletePromptCmd).Standalone()

	bedrockAgent_deletePromptCmd.Flags().String("prompt-identifier", "", "The unique identifier of the prompt.")
	bedrockAgent_deletePromptCmd.Flags().String("prompt-version", "", "The version of the prompt to delete.")
	bedrockAgent_deletePromptCmd.MarkFlagRequired("prompt-identifier")
	bedrockAgentCmd.AddCommand(bedrockAgent_deletePromptCmd)
}
