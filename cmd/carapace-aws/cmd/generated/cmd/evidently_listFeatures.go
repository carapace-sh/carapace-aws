package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var evidently_listFeaturesCmd = &cobra.Command{
	Use:   "list-features",
	Short: "Returns configuration details about all the features in the specified project.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(evidently_listFeaturesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(evidently_listFeaturesCmd).Standalone()

		evidently_listFeaturesCmd.Flags().String("max-results", "", "The maximum number of results to include in the response.")
		evidently_listFeaturesCmd.Flags().String("next-token", "", "The token to use when requesting the next set of results.")
		evidently_listFeaturesCmd.Flags().String("project", "", "The name or ARN of the project to return the feature list from.")
		evidently_listFeaturesCmd.MarkFlagRequired("project")
	})
	evidentlyCmd.AddCommand(evidently_listFeaturesCmd)
}
