package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudtrail_addTagsCmd = &cobra.Command{
	Use:   "add-tags",
	Short: "Adds one or more tags to a trail, event data store, dashboard, or channel, up to a limit of 50. Overwrites an existing tag's value when a new value is specified for an existing tag key.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudtrail_addTagsCmd).Standalone()

	cloudtrail_addTagsCmd.Flags().String("resource-id", "", "Specifies the ARN of the trail, event data store, dashboard, or channel to which one or more tags will be added.")
	cloudtrail_addTagsCmd.Flags().String("tags-list", "", "Contains a list of tags, up to a limit of 50")
	cloudtrail_addTagsCmd.MarkFlagRequired("resource-id")
	cloudtrail_addTagsCmd.MarkFlagRequired("tags-list")
	cloudtrailCmd.AddCommand(cloudtrail_addTagsCmd)
}
