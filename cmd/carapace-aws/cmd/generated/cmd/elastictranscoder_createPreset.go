package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elastictranscoder_createPresetCmd = &cobra.Command{
	Use:   "create-preset",
	Short: "The CreatePreset operation creates a preset with settings that you specify.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elastictranscoder_createPresetCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(elastictranscoder_createPresetCmd).Standalone()

		elastictranscoder_createPresetCmd.Flags().String("audio", "", "A section of the request body that specifies the audio parameters.")
		elastictranscoder_createPresetCmd.Flags().String("container", "", "The container type for the output file.")
		elastictranscoder_createPresetCmd.Flags().String("description", "", "A description of the preset.")
		elastictranscoder_createPresetCmd.Flags().String("name", "", "The name of the preset.")
		elastictranscoder_createPresetCmd.Flags().String("thumbnails", "", "A section of the request body that specifies the thumbnail parameters, if any.")
		elastictranscoder_createPresetCmd.Flags().String("video", "", "A section of the request body that specifies the video parameters.")
		elastictranscoder_createPresetCmd.MarkFlagRequired("container")
		elastictranscoder_createPresetCmd.MarkFlagRequired("name")
	})
	elastictranscoderCmd.AddCommand(elastictranscoder_createPresetCmd)
}
