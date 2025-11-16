package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediaconvert_createQueueCmd = &cobra.Command{
	Use:   "create-queue",
	Short: "Create a new transcoding queue.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediaconvert_createQueueCmd).Standalone()

	mediaconvert_createQueueCmd.Flags().String("concurrent-jobs", "", "Specify the maximum number of jobs your queue can process concurrently.")
	mediaconvert_createQueueCmd.Flags().String("description", "", "Optional.")
	mediaconvert_createQueueCmd.Flags().String("name", "", "The name of the queue that you are creating.")
	mediaconvert_createQueueCmd.Flags().String("pricing-plan", "", "Specifies whether the pricing plan for the queue is on-demand or reserved.")
	mediaconvert_createQueueCmd.Flags().String("reservation-plan-settings", "", "Details about the pricing plan for your reserved queue.")
	mediaconvert_createQueueCmd.Flags().String("status", "", "Initial state of the queue.")
	mediaconvert_createQueueCmd.Flags().String("tags", "", "The tags that you want to add to the resource.")
	mediaconvert_createQueueCmd.MarkFlagRequired("name")
	mediaconvertCmd.AddCommand(mediaconvert_createQueueCmd)
}
