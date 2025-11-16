package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_listFeatureGroupsCmd = &cobra.Command{
	Use:   "list-feature-groups",
	Short: "List `FeatureGroup`s based on given filter and order.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_listFeatureGroupsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sagemaker_listFeatureGroupsCmd).Standalone()

		sagemaker_listFeatureGroupsCmd.Flags().String("creation-time-after", "", "Use this parameter to search for `FeatureGroups`s created after a specific date and time.")
		sagemaker_listFeatureGroupsCmd.Flags().String("creation-time-before", "", "Use this parameter to search for `FeatureGroups`s created before a specific date and time.")
		sagemaker_listFeatureGroupsCmd.Flags().String("feature-group-status-equals", "", "A `FeatureGroup` status.")
		sagemaker_listFeatureGroupsCmd.Flags().String("max-results", "", "The maximum number of results returned by `ListFeatureGroups`.")
		sagemaker_listFeatureGroupsCmd.Flags().String("name-contains", "", "A string that partially matches one or more `FeatureGroup`s names.")
		sagemaker_listFeatureGroupsCmd.Flags().String("next-token", "", "A token to resume pagination of `ListFeatureGroups` results.")
		sagemaker_listFeatureGroupsCmd.Flags().String("offline-store-status-equals", "", "An `OfflineStore` status.")
		sagemaker_listFeatureGroupsCmd.Flags().String("sort-by", "", "The value on which the feature group list is sorted.")
		sagemaker_listFeatureGroupsCmd.Flags().String("sort-order", "", "The order in which feature groups are listed.")
	})
	sagemakerCmd.AddCommand(sagemaker_listFeatureGroupsCmd)
}
