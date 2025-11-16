package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elastictranscoder_listPresetsCmd = &cobra.Command{
	Use:   "list-presets",
	Short: "The ListPresets operation gets a list of the default presets included with Elastic Transcoder and the presets that you've added in an AWS region.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elastictranscoder_listPresetsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(elastictranscoder_listPresetsCmd).Standalone()

		elastictranscoder_listPresetsCmd.Flags().String("ascending", "", "To list presets in chronological order by the date and time that they were created, enter `true`.")
		elastictranscoder_listPresetsCmd.Flags().String("page-token", "", "When Elastic Transcoder returns more than one page of results, use `pageToken` in subsequent `GET` requests to get each successive page of results.")
	})
	elastictranscoderCmd.AddCommand(elastictranscoder_listPresetsCmd)
}
