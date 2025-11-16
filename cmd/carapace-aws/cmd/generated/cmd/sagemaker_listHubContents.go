package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_listHubContentsCmd = &cobra.Command{
	Use:   "list-hub-contents",
	Short: "List the contents of a hub.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_listHubContentsCmd).Standalone()

	sagemaker_listHubContentsCmd.Flags().String("creation-time-after", "", "Only list hub content that was created after the time specified.")
	sagemaker_listHubContentsCmd.Flags().String("creation-time-before", "", "Only list hub content that was created before the time specified.")
	sagemaker_listHubContentsCmd.Flags().String("hub-content-type", "", "The type of hub content to list.")
	sagemaker_listHubContentsCmd.Flags().String("hub-name", "", "The name of the hub to list the contents of.")
	sagemaker_listHubContentsCmd.Flags().String("max-results", "", "The maximum amount of hub content to list.")
	sagemaker_listHubContentsCmd.Flags().String("max-schema-version", "", "The upper bound of the hub content schema verion.")
	sagemaker_listHubContentsCmd.Flags().String("name-contains", "", "Only list hub content if the name contains the specified string.")
	sagemaker_listHubContentsCmd.Flags().String("next-token", "", "If the response to a previous `ListHubContents` request was truncated, the response includes a `NextToken`.")
	sagemaker_listHubContentsCmd.Flags().String("sort-by", "", "Sort hub content versions by either name or creation time.")
	sagemaker_listHubContentsCmd.Flags().String("sort-order", "", "Sort hubs by ascending or descending order.")
	sagemaker_listHubContentsCmd.MarkFlagRequired("hub-content-type")
	sagemaker_listHubContentsCmd.MarkFlagRequired("hub-name")
	sagemakerCmd.AddCommand(sagemaker_listHubContentsCmd)
}
