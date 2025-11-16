package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var scheduler_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Assigns one or more tags (key-value pairs) to the specified EventBridge Scheduler resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(scheduler_tagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(scheduler_tagResourceCmd).Standalone()

		scheduler_tagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the schedule group that you are adding tags to.")
		scheduler_tagResourceCmd.Flags().String("tags", "", "The list of tags to associate with the schedule group.")
		scheduler_tagResourceCmd.MarkFlagRequired("resource-arn")
		scheduler_tagResourceCmd.MarkFlagRequired("tags")
	})
	schedulerCmd.AddCommand(scheduler_tagResourceCmd)
}
