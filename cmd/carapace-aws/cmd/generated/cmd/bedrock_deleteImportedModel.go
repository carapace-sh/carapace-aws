package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrock_deleteImportedModelCmd = &cobra.Command{
	Use:   "delete-imported-model",
	Short: "Deletes a custom model that you imported earlier.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrock_deleteImportedModelCmd).Standalone()

	bedrock_deleteImportedModelCmd.Flags().String("model-identifier", "", "Name of the imported model to delete.")
	bedrock_deleteImportedModelCmd.MarkFlagRequired("model-identifier")
	bedrockCmd.AddCommand(bedrock_deleteImportedModelCmd)
}
