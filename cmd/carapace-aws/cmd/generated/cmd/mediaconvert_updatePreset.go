package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediaconvert_updatePresetCmd = &cobra.Command{
	Use:   "update-preset",
	Short: "Modify one of your existing presets.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediaconvert_updatePresetCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mediaconvert_updatePresetCmd).Standalone()

		mediaconvert_updatePresetCmd.Flags().String("category", "", "The new category for the preset, if you are changing it.")
		mediaconvert_updatePresetCmd.Flags().String("description", "", "The new description for the preset, if you are changing it.")
		mediaconvert_updatePresetCmd.Flags().String("name", "", "The name of the preset you are modifying.")
		mediaconvert_updatePresetCmd.Flags().String("settings", "", "Settings for preset")
		mediaconvert_updatePresetCmd.MarkFlagRequired("name")
	})
	mediaconvertCmd.AddCommand(mediaconvert_updatePresetCmd)
}
