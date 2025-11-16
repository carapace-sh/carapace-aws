package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var scheduler_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "Lists the tags associated with the Scheduler resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(scheduler_listTagsForResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(scheduler_listTagsForResourceCmd).Standalone()

		scheduler_listTagsForResourceCmd.Flags().String("resource-arn", "", "The ARN of the EventBridge Scheduler resource for which you want to view tags.")
		scheduler_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	})
	schedulerCmd.AddCommand(scheduler_listTagsForResourceCmd)
}
