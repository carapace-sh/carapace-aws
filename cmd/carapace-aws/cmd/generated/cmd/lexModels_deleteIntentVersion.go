package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lexModels_deleteIntentVersionCmd = &cobra.Command{
	Use:   "delete-intent-version",
	Short: "Deletes a specific version of an intent.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lexModels_deleteIntentVersionCmd).Standalone()

	lexModels_deleteIntentVersionCmd.Flags().String("name", "", "The name of the intent.")
	lexModels_deleteIntentVersionCmd.Flags().String("version", "", "The version of the intent to delete.")
	lexModels_deleteIntentVersionCmd.MarkFlagRequired("name")
	lexModels_deleteIntentVersionCmd.MarkFlagRequired("version")
	lexModelsCmd.AddCommand(lexModels_deleteIntentVersionCmd)
}
