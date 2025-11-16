package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var resourceGroups_startTagSyncTaskCmd = &cobra.Command{
	Use:   "start-tag-sync-task",
	Short: "Creates a new tag-sync task to onboard and sync resources tagged with a specific tag key-value pair to an application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(resourceGroups_startTagSyncTaskCmd).Standalone()

	resourceGroups_startTagSyncTaskCmd.Flags().String("group", "", "The Amazon resource name (ARN) or name of the application group for which you want to create a tag-sync task.")
	resourceGroups_startTagSyncTaskCmd.Flags().String("resource-query", "", "The query you can use to create the tag-sync task.")
	resourceGroups_startTagSyncTaskCmd.Flags().String("role-arn", "", "The Amazon resource name (ARN) of the role assumed by the service to tag and untag resources on your behalf.")
	resourceGroups_startTagSyncTaskCmd.Flags().String("tag-key", "", "The tag key.")
	resourceGroups_startTagSyncTaskCmd.Flags().String("tag-value", "", "The tag value.")
	resourceGroups_startTagSyncTaskCmd.MarkFlagRequired("group")
	resourceGroups_startTagSyncTaskCmd.MarkFlagRequired("role-arn")
	resourceGroupsCmd.AddCommand(resourceGroups_startTagSyncTaskCmd)
}
