package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediaconvert_deleteQueueCmd = &cobra.Command{
	Use:   "delete-queue",
	Short: "Permanently delete a queue you have created.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediaconvert_deleteQueueCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mediaconvert_deleteQueueCmd).Standalone()

		mediaconvert_deleteQueueCmd.Flags().String("name", "", "The name of the queue that you want to delete.")
		mediaconvert_deleteQueueCmd.MarkFlagRequired("name")
	})
	mediaconvertCmd.AddCommand(mediaconvert_deleteQueueCmd)
}
