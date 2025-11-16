package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediaconvert_createPresetCmd = &cobra.Command{
	Use:   "create-preset",
	Short: "Create a new preset.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediaconvert_createPresetCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mediaconvert_createPresetCmd).Standalone()

		mediaconvert_createPresetCmd.Flags().String("category", "", "Optional.")
		mediaconvert_createPresetCmd.Flags().String("description", "", "Optional.")
		mediaconvert_createPresetCmd.Flags().String("name", "", "The name of the preset you are creating.")
		mediaconvert_createPresetCmd.Flags().String("settings", "", "Settings for preset")
		mediaconvert_createPresetCmd.Flags().String("tags", "", "The tags that you want to add to the resource.")
		mediaconvert_createPresetCmd.MarkFlagRequired("name")
		mediaconvert_createPresetCmd.MarkFlagRequired("settings")
	})
	mediaconvertCmd.AddCommand(mediaconvert_createPresetCmd)
}
