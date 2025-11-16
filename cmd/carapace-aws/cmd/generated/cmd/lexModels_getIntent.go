package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lexModels_getIntentCmd = &cobra.Command{
	Use:   "get-intent",
	Short: "Returns information about an intent.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lexModels_getIntentCmd).Standalone()

	lexModels_getIntentCmd.Flags().String("name", "", "The name of the intent.")
	lexModels_getIntentCmd.Flags().String("version", "", "The version of the intent.")
	lexModels_getIntentCmd.MarkFlagRequired("name")
	lexModels_getIntentCmd.MarkFlagRequired("version")
	lexModelsCmd.AddCommand(lexModels_getIntentCmd)
}
