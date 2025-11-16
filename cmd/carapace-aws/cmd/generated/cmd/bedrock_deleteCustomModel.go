package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrock_deleteCustomModelCmd = &cobra.Command{
	Use:   "delete-custom-model",
	Short: "Deletes a custom model that you created earlier.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrock_deleteCustomModelCmd).Standalone()

	bedrock_deleteCustomModelCmd.Flags().String("model-identifier", "", "Name of the model to delete.")
	bedrock_deleteCustomModelCmd.MarkFlagRequired("model-identifier")
	bedrockCmd.AddCommand(bedrock_deleteCustomModelCmd)
}
