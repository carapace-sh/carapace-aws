package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrock_getPromptRouterCmd = &cobra.Command{
	Use:   "get-prompt-router",
	Short: "Retrieves details about a prompt router.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrock_getPromptRouterCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(bedrock_getPromptRouterCmd).Standalone()

		bedrock_getPromptRouterCmd.Flags().String("prompt-router-arn", "", "The prompt router's ARN")
		bedrock_getPromptRouterCmd.MarkFlagRequired("prompt-router-arn")
	})
	bedrockCmd.AddCommand(bedrock_getPromptRouterCmd)
}
