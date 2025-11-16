package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediaconvert_getQueueCmd = &cobra.Command{
	Use:   "get-queue",
	Short: "Retrieve the JSON for a specific queue.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediaconvert_getQueueCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mediaconvert_getQueueCmd).Standalone()

		mediaconvert_getQueueCmd.Flags().String("name", "", "The name of the queue that you want information about.")
		mediaconvert_getQueueCmd.MarkFlagRequired("name")
	})
	mediaconvertCmd.AddCommand(mediaconvert_getQueueCmd)
}
