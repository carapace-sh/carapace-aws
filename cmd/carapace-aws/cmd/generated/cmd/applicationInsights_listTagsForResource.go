package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var applicationInsights_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "Retrieve a list of the tags (keys and values) that are associated with a specified application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(applicationInsights_listTagsForResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(applicationInsights_listTagsForResourceCmd).Standalone()

		applicationInsights_listTagsForResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the application that you want to retrieve tag information for.")
		applicationInsights_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	})
	applicationInsightsCmd.AddCommand(applicationInsights_listTagsForResourceCmd)
}
