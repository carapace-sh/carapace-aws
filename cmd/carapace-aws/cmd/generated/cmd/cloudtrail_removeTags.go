package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudtrail_removeTagsCmd = &cobra.Command{
	Use:   "remove-tags",
	Short: "Removes the specified tags from a trail, event data store, dashboard, or channel.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudtrail_removeTagsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloudtrail_removeTagsCmd).Standalone()

		cloudtrail_removeTagsCmd.Flags().String("resource-id", "", "Specifies the ARN of the trail, event data store, dashboard, or channel from which tags should be removed.")
		cloudtrail_removeTagsCmd.Flags().String("tags-list", "", "Specifies a list of tags to be removed.")
		cloudtrail_removeTagsCmd.MarkFlagRequired("resource-id")
		cloudtrail_removeTagsCmd.MarkFlagRequired("tags-list")
	})
	cloudtrailCmd.AddCommand(cloudtrail_removeTagsCmd)
}
