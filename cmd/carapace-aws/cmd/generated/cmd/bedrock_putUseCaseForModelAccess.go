package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrock_putUseCaseForModelAccessCmd = &cobra.Command{
	Use:   "put-use-case-for-model-access",
	Short: "Put usecase for model access.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrock_putUseCaseForModelAccessCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(bedrock_putUseCaseForModelAccessCmd).Standalone()

		bedrock_putUseCaseForModelAccessCmd.Flags().String("form-data", "", "Put customer profile Request.")
		bedrock_putUseCaseForModelAccessCmd.MarkFlagRequired("form-data")
	})
	bedrockCmd.AddCommand(bedrock_putUseCaseForModelAccessCmd)
}
