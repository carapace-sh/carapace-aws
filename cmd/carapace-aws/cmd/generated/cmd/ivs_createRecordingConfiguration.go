package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ivs_createRecordingConfigurationCmd = &cobra.Command{
	Use:   "create-recording-configuration",
	Short: "Creates a new recording configuration, used to enable recording to Amazon S3.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ivs_createRecordingConfigurationCmd).Standalone()

	ivs_createRecordingConfigurationCmd.Flags().String("destination-configuration", "", "A complex type that contains a destination configuration for where recorded video will be stored.")
	ivs_createRecordingConfigurationCmd.Flags().String("name", "", "Recording-configuration name.")
	ivs_createRecordingConfigurationCmd.Flags().String("recording-reconnect-window-seconds", "", "If a broadcast disconnects and then reconnects within the specified interval, the multiple streams will be considered a single broadcast and merged together.")
	ivs_createRecordingConfigurationCmd.Flags().String("rendition-configuration", "", "Object that describes which renditions should be recorded for a stream.")
	ivs_createRecordingConfigurationCmd.Flags().String("tags", "", "Array of 1-50 maps, each of the form `string:string (key:value)`.")
	ivs_createRecordingConfigurationCmd.Flags().String("thumbnail-configuration", "", "A complex type that allows you to enable/disable the recording of thumbnails for a live session and modify the interval at which thumbnails are generated for the live session.")
	ivs_createRecordingConfigurationCmd.MarkFlagRequired("destination-configuration")
	ivsCmd.AddCommand(ivs_createRecordingConfigurationCmd)
}
