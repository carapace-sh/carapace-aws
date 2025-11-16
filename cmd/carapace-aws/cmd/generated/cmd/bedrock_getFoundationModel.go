package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrock_getFoundationModelCmd = &cobra.Command{
	Use:   "get-foundation-model",
	Short: "Get details about a Amazon Bedrock foundation model.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrock_getFoundationModelCmd).Standalone()

	bedrock_getFoundationModelCmd.Flags().String("model-identifier", "", "The model identifier.")
	bedrock_getFoundationModelCmd.MarkFlagRequired("model-identifier")
	bedrockCmd.AddCommand(bedrock_getFoundationModelCmd)
}
