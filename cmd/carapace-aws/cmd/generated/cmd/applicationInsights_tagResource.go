package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var applicationInsights_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Add one or more tags (keys and values) to a specified application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(applicationInsights_tagResourceCmd).Standalone()

	applicationInsights_tagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the application that you want to add one or more tags to.")
	applicationInsights_tagResourceCmd.Flags().String("tags", "", "A list of tags that to add to the application.")
	applicationInsights_tagResourceCmd.MarkFlagRequired("resource-arn")
	applicationInsights_tagResourceCmd.MarkFlagRequired("tags")
	applicationInsightsCmd.AddCommand(applicationInsights_tagResourceCmd)
}
