package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_listHubContentVersionsCmd = &cobra.Command{
	Use:   "list-hub-content-versions",
	Short: "List hub content versions.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_listHubContentVersionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sagemaker_listHubContentVersionsCmd).Standalone()

		sagemaker_listHubContentVersionsCmd.Flags().String("creation-time-after", "", "Only list hub content versions that were created after the time specified.")
		sagemaker_listHubContentVersionsCmd.Flags().String("creation-time-before", "", "Only list hub content versions that were created before the time specified.")
		sagemaker_listHubContentVersionsCmd.Flags().String("hub-content-name", "", "The name of the hub content.")
		sagemaker_listHubContentVersionsCmd.Flags().String("hub-content-type", "", "The type of hub content to list versions of.")
		sagemaker_listHubContentVersionsCmd.Flags().String("hub-name", "", "The name of the hub to list the content versions of.")
		sagemaker_listHubContentVersionsCmd.Flags().String("max-results", "", "The maximum number of hub content versions to list.")
		sagemaker_listHubContentVersionsCmd.Flags().String("max-schema-version", "", "The upper bound of the hub content schema version.")
		sagemaker_listHubContentVersionsCmd.Flags().String("min-version", "", "The lower bound of the hub content versions to list.")
		sagemaker_listHubContentVersionsCmd.Flags().String("next-token", "", "If the response to a previous `ListHubContentVersions` request was truncated, the response includes a `NextToken`.")
		sagemaker_listHubContentVersionsCmd.Flags().String("sort-by", "", "Sort hub content versions by either name or creation time.")
		sagemaker_listHubContentVersionsCmd.Flags().String("sort-order", "", "Sort hub content versions by ascending or descending order.")
		sagemaker_listHubContentVersionsCmd.MarkFlagRequired("hub-content-name")
		sagemaker_listHubContentVersionsCmd.MarkFlagRequired("hub-content-type")
		sagemaker_listHubContentVersionsCmd.MarkFlagRequired("hub-name")
	})
	sagemakerCmd.AddCommand(sagemaker_listHubContentVersionsCmd)
}
