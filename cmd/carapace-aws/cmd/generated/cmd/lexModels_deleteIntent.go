package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lexModels_deleteIntentCmd = &cobra.Command{
	Use:   "delete-intent",
	Short: "Deletes all versions of the intent, including the `$LATEST` version.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lexModels_deleteIntentCmd).Standalone()

	lexModels_deleteIntentCmd.Flags().String("name", "", "The name of the intent.")
	lexModels_deleteIntentCmd.MarkFlagRequired("name")
	lexModelsCmd.AddCommand(lexModels_deleteIntentCmd)
}
