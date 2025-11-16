package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediaconvert_updateQueueCmd = &cobra.Command{
	Use:   "update-queue",
	Short: "Modify one of your existing queues.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediaconvert_updateQueueCmd).Standalone()

	mediaconvert_updateQueueCmd.Flags().String("concurrent-jobs", "", "Specify the maximum number of jobs your queue can process concurrently.")
	mediaconvert_updateQueueCmd.Flags().String("description", "", "The new description for the queue, if you are changing it.")
	mediaconvert_updateQueueCmd.Flags().String("name", "", "The name of the queue that you are modifying.")
	mediaconvert_updateQueueCmd.Flags().String("reservation-plan-settings", "", "The new details of your pricing plan for your reserved queue.")
	mediaconvert_updateQueueCmd.Flags().String("status", "", "Pause or activate a queue by changing its status between ACTIVE and PAUSED.")
	mediaconvert_updateQueueCmd.MarkFlagRequired("name")
	mediaconvertCmd.AddCommand(mediaconvert_updateQueueCmd)
}
