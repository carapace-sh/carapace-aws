package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rekognition_deleteStreamProcessorCmd = &cobra.Command{
	Use:   "delete-stream-processor",
	Short: "Deletes the stream processor identified by `Name`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rekognition_deleteStreamProcessorCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(rekognition_deleteStreamProcessorCmd).Standalone()

		rekognition_deleteStreamProcessorCmd.Flags().String("name", "", "The name of the stream processor you want to delete.")
		rekognition_deleteStreamProcessorCmd.MarkFlagRequired("name")
	})
	rekognitionCmd.AddCommand(rekognition_deleteStreamProcessorCmd)
}
