package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var evidently_deleteFeatureCmd = &cobra.Command{
	Use:   "delete-feature",
	Short: "Deletes an Evidently feature.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(evidently_deleteFeatureCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(evidently_deleteFeatureCmd).Standalone()

		evidently_deleteFeatureCmd.Flags().String("feature", "", "The name of the feature to delete.")
		evidently_deleteFeatureCmd.Flags().String("project", "", "The name or ARN of the project that contains the feature to delete.")
		evidently_deleteFeatureCmd.MarkFlagRequired("feature")
		evidently_deleteFeatureCmd.MarkFlagRequired("project")
	})
	evidentlyCmd.AddCommand(evidently_deleteFeatureCmd)
}
