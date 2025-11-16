package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rekognition_updateStreamProcessorCmd = &cobra.Command{
	Use:   "update-stream-processor",
	Short: "Allows you to update a stream processor.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rekognition_updateStreamProcessorCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(rekognition_updateStreamProcessorCmd).Standalone()

		rekognition_updateStreamProcessorCmd.Flags().String("data-sharing-preference-for-update", "", "Shows whether you are sharing data with Rekognition to improve model performance.")
		rekognition_updateStreamProcessorCmd.Flags().String("name", "", "Name of the stream processor that you want to update.")
		rekognition_updateStreamProcessorCmd.Flags().String("parameters-to-delete", "", "A list of parameters you want to delete from the stream processor.")
		rekognition_updateStreamProcessorCmd.Flags().String("regions-of-interest-for-update", "", "Specifies locations in the frames where Amazon Rekognition checks for objects or people.")
		rekognition_updateStreamProcessorCmd.Flags().String("settings-for-update", "", "The stream processor settings that you want to update.")
		rekognition_updateStreamProcessorCmd.MarkFlagRequired("name")
	})
	rekognitionCmd.AddCommand(rekognition_updateStreamProcessorCmd)
}
