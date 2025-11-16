package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrock_getFoundationModelAvailabilityCmd = &cobra.Command{
	Use:   "get-foundation-model-availability",
	Short: "Get information about the Foundation model availability.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrock_getFoundationModelAvailabilityCmd).Standalone()

	bedrock_getFoundationModelAvailabilityCmd.Flags().String("model-id", "", "The model Id of the foundation model.")
	bedrock_getFoundationModelAvailabilityCmd.MarkFlagRequired("model-id")
	bedrockCmd.AddCommand(bedrock_getFoundationModelAvailabilityCmd)
}
