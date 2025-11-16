package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockAgent_updatePromptCmd = &cobra.Command{
	Use:   "update-prompt",
	Short: "Modifies a prompt in your prompt library.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockAgent_updatePromptCmd).Standalone()

	bedrockAgent_updatePromptCmd.Flags().String("customer-encryption-key-arn", "", "The Amazon Resource Name (ARN) of the KMS key to encrypt the prompt.")
	bedrockAgent_updatePromptCmd.Flags().String("default-variant", "", "The name of the default variant for the prompt.")
	bedrockAgent_updatePromptCmd.Flags().String("description", "", "A description for the prompt.")
	bedrockAgent_updatePromptCmd.Flags().String("name", "", "A name for the prompt.")
	bedrockAgent_updatePromptCmd.Flags().String("prompt-identifier", "", "The unique identifier of the prompt.")
	bedrockAgent_updatePromptCmd.Flags().String("variants", "", "A list of objects, each containing details about a variant of the prompt.")
	bedrockAgent_updatePromptCmd.MarkFlagRequired("name")
	bedrockAgent_updatePromptCmd.MarkFlagRequired("prompt-identifier")
	bedrockAgentCmd.AddCommand(bedrockAgent_updatePromptCmd)
}
