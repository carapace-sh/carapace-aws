package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrock_deleteInferenceProfileCmd = &cobra.Command{
	Use:   "delete-inference-profile",
	Short: "Deletes an application inference profile.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrock_deleteInferenceProfileCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(bedrock_deleteInferenceProfileCmd).Standalone()

		bedrock_deleteInferenceProfileCmd.Flags().String("inference-profile-identifier", "", "The Amazon Resource Name (ARN) or ID of the application inference profile to delete.")
		bedrock_deleteInferenceProfileCmd.MarkFlagRequired("inference-profile-identifier")
	})
	bedrockCmd.AddCommand(bedrock_deleteInferenceProfileCmd)
}
