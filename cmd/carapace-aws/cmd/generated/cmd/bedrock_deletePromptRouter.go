package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrock_deletePromptRouterCmd = &cobra.Command{
	Use:   "delete-prompt-router",
	Short: "Deletes a specified prompt router.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrock_deletePromptRouterCmd).Standalone()

	bedrock_deletePromptRouterCmd.Flags().String("prompt-router-arn", "", "The Amazon Resource Name (ARN) of the prompt router to delete.")
	bedrock_deletePromptRouterCmd.MarkFlagRequired("prompt-router-arn")
	bedrockCmd.AddCommand(bedrock_deletePromptRouterCmd)
}
