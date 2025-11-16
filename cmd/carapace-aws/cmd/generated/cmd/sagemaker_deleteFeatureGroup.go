package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_deleteFeatureGroupCmd = &cobra.Command{
	Use:   "delete-feature-group",
	Short: "Delete the `FeatureGroup` and any data that was written to the `OnlineStore` of the `FeatureGroup`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_deleteFeatureGroupCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sagemaker_deleteFeatureGroupCmd).Standalone()

		sagemaker_deleteFeatureGroupCmd.Flags().String("feature-group-name", "", "The name of the `FeatureGroup` you want to delete.")
		sagemaker_deleteFeatureGroupCmd.MarkFlagRequired("feature-group-name")
	})
	sagemakerCmd.AddCommand(sagemaker_deleteFeatureGroupCmd)
}
