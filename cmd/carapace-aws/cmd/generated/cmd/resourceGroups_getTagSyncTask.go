package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var resourceGroups_getTagSyncTaskCmd = &cobra.Command{
	Use:   "get-tag-sync-task",
	Short: "Returns information about a specified tag-sync task.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(resourceGroups_getTagSyncTaskCmd).Standalone()

	resourceGroups_getTagSyncTaskCmd.Flags().String("task-arn", "", "The Amazon resource name (ARN) of the tag-sync task.")
	resourceGroups_getTagSyncTaskCmd.MarkFlagRequired("task-arn")
	resourceGroupsCmd.AddCommand(resourceGroups_getTagSyncTaskCmd)
}
