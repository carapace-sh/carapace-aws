package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var timestreamQuery_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "List all tags on a Timestream query resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(timestreamQuery_listTagsForResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(timestreamQuery_listTagsForResourceCmd).Standalone()

		timestreamQuery_listTagsForResourceCmd.Flags().String("max-results", "", "The maximum number of tags to return.")
		timestreamQuery_listTagsForResourceCmd.Flags().String("next-token", "", "A pagination token to resume pagination.")
		timestreamQuery_listTagsForResourceCmd.Flags().String("resource-arn", "", "The Timestream resource with tags to be listed.")
		timestreamQuery_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	})
	timestreamQueryCmd.AddCommand(timestreamQuery_listTagsForResourceCmd)
}
