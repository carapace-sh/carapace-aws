package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elastictranscoder_createJobCmd = &cobra.Command{
	Use:   "create-job",
	Short: "When you create a job, Elastic Transcoder returns JSON data that includes the values that you specified plus information about the job that is created.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elastictranscoder_createJobCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(elastictranscoder_createJobCmd).Standalone()

		elastictranscoder_createJobCmd.Flags().String("input", "", "A section of the request body that provides information about the file that is being transcoded.")
		elastictranscoder_createJobCmd.Flags().String("inputs", "", "A section of the request body that provides information about the files that are being transcoded.")
		elastictranscoder_createJobCmd.Flags().String("output", "", "A section of the request body that provides information about the transcoded (target) file.")
		elastictranscoder_createJobCmd.Flags().String("output-key-prefix", "", "The value, if any, that you want Elastic Transcoder to prepend to the names of all files that this job creates, including output files, thumbnails, and playlists.")
		elastictranscoder_createJobCmd.Flags().String("outputs", "", "A section of the request body that provides information about the transcoded (target) files.")
		elastictranscoder_createJobCmd.Flags().String("pipeline-id", "", "The `Id` of the pipeline that you want Elastic Transcoder to use for transcoding.")
		elastictranscoder_createJobCmd.Flags().String("playlists", "", "If you specify a preset in `PresetId` for which the value of `Container` is fmp4 (Fragmented MP4) or ts (MPEG-TS), Playlists contains information about the master playlists that you want Elastic Transcoder to create.")
		elastictranscoder_createJobCmd.Flags().String("user-metadata", "", "User-defined metadata that you want to associate with an Elastic Transcoder job.")
		elastictranscoder_createJobCmd.MarkFlagRequired("pipeline-id")
	})
	elastictranscoderCmd.AddCommand(elastictranscoder_createJobCmd)
}
