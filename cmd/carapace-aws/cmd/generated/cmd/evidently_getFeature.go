package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var evidently_getFeatureCmd = &cobra.Command{
	Use:   "get-feature",
	Short: "Returns the details about one feature.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(evidently_getFeatureCmd).Standalone()

	evidently_getFeatureCmd.Flags().String("feature", "", "The name of the feature that you want to retrieve information for.")
	evidently_getFeatureCmd.Flags().String("project", "", "The name or ARN of the project that contains the feature.")
	evidently_getFeatureCmd.MarkFlagRequired("feature")
	evidently_getFeatureCmd.MarkFlagRequired("project")
	evidentlyCmd.AddCommand(evidently_getFeatureCmd)
}
