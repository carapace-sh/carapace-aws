package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var scheduler_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Removes one or more tags from the specified EventBridge Scheduler schedule group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(scheduler_untagResourceCmd).Standalone()

	scheduler_untagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the schedule group from which you are removing tags.")
	scheduler_untagResourceCmd.Flags().String("tag-keys", "", "The list of tag keys to remove from the resource.")
	scheduler_untagResourceCmd.MarkFlagRequired("resource-arn")
	scheduler_untagResourceCmd.MarkFlagRequired("tag-keys")
	schedulerCmd.AddCommand(scheduler_untagResourceCmd)
}
