package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockRuntime_countTokensCmd = &cobra.Command{
	Use:   "count-tokens",
	Short: "Returns the token count for a given inference request.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockRuntime_countTokensCmd).Standalone()

	bedrockRuntime_countTokensCmd.Flags().String("input", "", "The input for which to count tokens.")
	bedrockRuntime_countTokensCmd.Flags().String("model-id", "", "The unique identifier or ARN of the foundation model to use for token counting.")
	bedrockRuntime_countTokensCmd.MarkFlagRequired("input")
	bedrockRuntime_countTokensCmd.MarkFlagRequired("model-id")
	bedrockRuntimeCmd.AddCommand(bedrockRuntime_countTokensCmd)
}
