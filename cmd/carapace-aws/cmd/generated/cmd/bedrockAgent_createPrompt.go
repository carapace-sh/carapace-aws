package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockAgent_createPromptCmd = &cobra.Command{
	Use:   "create-prompt",
	Short: "Creates a prompt in your prompt library that you can add to a flow.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockAgent_createPromptCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(bedrockAgent_createPromptCmd).Standalone()

		bedrockAgent_createPromptCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier to ensure that the API request completes no more than one time.")
		bedrockAgent_createPromptCmd.Flags().String("customer-encryption-key-arn", "", "The Amazon Resource Name (ARN) of the KMS key to encrypt the prompt.")
		bedrockAgent_createPromptCmd.Flags().String("default-variant", "", "The name of the default variant for the prompt.")
		bedrockAgent_createPromptCmd.Flags().String("description", "", "A description for the prompt.")
		bedrockAgent_createPromptCmd.Flags().String("name", "", "A name for the prompt.")
		bedrockAgent_createPromptCmd.Flags().String("tags", "", "Any tags that you want to attach to the prompt.")
		bedrockAgent_createPromptCmd.Flags().String("variants", "", "A list of objects, each containing details about a variant of the prompt.")
		bedrockAgent_createPromptCmd.MarkFlagRequired("name")
	})
	bedrockAgentCmd.AddCommand(bedrockAgent_createPromptCmd)
}
