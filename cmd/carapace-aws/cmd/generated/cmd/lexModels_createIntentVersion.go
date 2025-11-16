package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lexModels_createIntentVersionCmd = &cobra.Command{
	Use:   "create-intent-version",
	Short: "Creates a new version of an intent based on the `$LATEST` version of the intent.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lexModels_createIntentVersionCmd).Standalone()

	lexModels_createIntentVersionCmd.Flags().String("checksum", "", "Checksum of the `$LATEST` version of the intent that should be used to create the new version.")
	lexModels_createIntentVersionCmd.Flags().String("name", "", "The name of the intent that you want to create a new version of.")
	lexModels_createIntentVersionCmd.MarkFlagRequired("name")
	lexModelsCmd.AddCommand(lexModels_createIntentVersionCmd)
}
