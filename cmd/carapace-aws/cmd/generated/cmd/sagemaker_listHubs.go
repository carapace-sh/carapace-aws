package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_listHubsCmd = &cobra.Command{
	Use:   "list-hubs",
	Short: "List all existing hubs.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_listHubsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sagemaker_listHubsCmd).Standalone()

		sagemaker_listHubsCmd.Flags().String("creation-time-after", "", "Only list hubs that were created after the time specified.")
		sagemaker_listHubsCmd.Flags().String("creation-time-before", "", "Only list hubs that were created before the time specified.")
		sagemaker_listHubsCmd.Flags().String("last-modified-time-after", "", "Only list hubs that were last modified after the time specified.")
		sagemaker_listHubsCmd.Flags().String("last-modified-time-before", "", "Only list hubs that were last modified before the time specified.")
		sagemaker_listHubsCmd.Flags().String("max-results", "", "The maximum number of hubs to list.")
		sagemaker_listHubsCmd.Flags().String("name-contains", "", "Only list hubs with names that contain the specified string.")
		sagemaker_listHubsCmd.Flags().String("next-token", "", "If the response to a previous `ListHubs` request was truncated, the response includes a `NextToken`.")
		sagemaker_listHubsCmd.Flags().String("sort-by", "", "Sort hubs by either name or creation time.")
		sagemaker_listHubsCmd.Flags().String("sort-order", "", "Sort hubs by ascending or descending order.")
	})
	sagemakerCmd.AddCommand(sagemaker_listHubsCmd)
}
