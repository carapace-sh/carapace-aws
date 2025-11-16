package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elastictranscoder_readPresetCmd = &cobra.Command{
	Use:   "read-preset",
	Short: "The ReadPreset operation gets detailed information about a preset.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elastictranscoder_readPresetCmd).Standalone()

	elastictranscoder_readPresetCmd.Flags().String("id", "", "The identifier of the preset for which you want to get detailed information.")
	elastictranscoder_readPresetCmd.MarkFlagRequired("id")
	elastictranscoderCmd.AddCommand(elastictranscoder_readPresetCmd)
}
