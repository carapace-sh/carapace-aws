package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrock_getInferenceProfileCmd = &cobra.Command{
	Use:   "get-inference-profile",
	Short: "Gets information about an inference profile.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrock_getInferenceProfileCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(bedrock_getInferenceProfileCmd).Standalone()

		bedrock_getInferenceProfileCmd.Flags().String("inference-profile-identifier", "", "The ID or Amazon Resource Name (ARN) of the inference profile.")
		bedrock_getInferenceProfileCmd.MarkFlagRequired("inference-profile-identifier")
	})
	bedrockCmd.AddCommand(bedrock_getInferenceProfileCmd)
}
