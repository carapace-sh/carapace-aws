package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrock_getUseCaseForModelAccessCmd = &cobra.Command{
	Use:   "get-use-case-for-model-access",
	Short: "Get usecase for model access.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrock_getUseCaseForModelAccessCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(bedrock_getUseCaseForModelAccessCmd).Standalone()

	})
	bedrockCmd.AddCommand(bedrock_getUseCaseForModelAccessCmd)
}
