package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lexModels_getIntentVersionsCmd = &cobra.Command{
	Use:   "get-intent-versions",
	Short: "Gets information about all of the versions of an intent.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lexModels_getIntentVersionsCmd).Standalone()

	lexModels_getIntentVersionsCmd.Flags().String("max-results", "", "The maximum number of intent versions to return in the response.")
	lexModels_getIntentVersionsCmd.Flags().String("name", "", "The name of the intent for which versions should be returned.")
	lexModels_getIntentVersionsCmd.Flags().String("next-token", "", "A pagination token for fetching the next page of intent versions.")
	lexModels_getIntentVersionsCmd.MarkFlagRequired("name")
	lexModelsCmd.AddCommand(lexModels_getIntentVersionsCmd)
}
