package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var resourceGroups_cancelTagSyncTaskCmd = &cobra.Command{
	Use:   "cancel-tag-sync-task",
	Short: "Cancels the specified tag-sync task.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(resourceGroups_cancelTagSyncTaskCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(resourceGroups_cancelTagSyncTaskCmd).Standalone()

		resourceGroups_cancelTagSyncTaskCmd.Flags().String("task-arn", "", "The Amazon resource name (ARN) of the tag-sync task.")
		resourceGroups_cancelTagSyncTaskCmd.MarkFlagRequired("task-arn")
	})
	resourceGroupsCmd.AddCommand(resourceGroups_cancelTagSyncTaskCmd)
}
