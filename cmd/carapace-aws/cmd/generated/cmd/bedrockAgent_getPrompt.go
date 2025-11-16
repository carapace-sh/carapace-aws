package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockAgent_getPromptCmd = &cobra.Command{
	Use:   "get-prompt",
	Short: "Retrieves information about the working draft (`DRAFT` version) of a prompt or a version of it, depending on whether you include the `promptVersion` field or not.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockAgent_getPromptCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(bedrockAgent_getPromptCmd).Standalone()

		bedrockAgent_getPromptCmd.Flags().String("prompt-identifier", "", "The unique identifier of the prompt.")
		bedrockAgent_getPromptCmd.Flags().String("prompt-version", "", "The version of the prompt about which you want to retrieve information.")
		bedrockAgent_getPromptCmd.MarkFlagRequired("prompt-identifier")
	})
	bedrockAgentCmd.AddCommand(bedrockAgent_getPromptCmd)
}
