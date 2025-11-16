package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elastictranscoder_deletePresetCmd = &cobra.Command{
	Use:   "delete-preset",
	Short: "The DeletePreset operation removes a preset that you've added in an AWS region.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elastictranscoder_deletePresetCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(elastictranscoder_deletePresetCmd).Standalone()

		elastictranscoder_deletePresetCmd.Flags().String("id", "", "The identifier of the preset for which you want to get detailed information.")
		elastictranscoder_deletePresetCmd.MarkFlagRequired("id")
	})
	elastictranscoderCmd.AddCommand(elastictranscoder_deletePresetCmd)
}
