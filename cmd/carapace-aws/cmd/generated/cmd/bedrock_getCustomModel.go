package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrock_getCustomModelCmd = &cobra.Command{
	Use:   "get-custom-model",
	Short: "Get the properties associated with a Amazon Bedrock custom model that you have created.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrock_getCustomModelCmd).Standalone()

	bedrock_getCustomModelCmd.Flags().String("model-identifier", "", "Name or Amazon Resource Name (ARN) of the custom model.")
	bedrock_getCustomModelCmd.MarkFlagRequired("model-identifier")
	bedrockCmd.AddCommand(bedrock_getCustomModelCmd)
}
