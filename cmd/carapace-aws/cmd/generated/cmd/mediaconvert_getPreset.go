package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediaconvert_getPresetCmd = &cobra.Command{
	Use:   "get-preset",
	Short: "Retrieve the JSON for a specific preset.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediaconvert_getPresetCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mediaconvert_getPresetCmd).Standalone()

		mediaconvert_getPresetCmd.Flags().String("name", "", "The name of the preset.")
		mediaconvert_getPresetCmd.MarkFlagRequired("name")
	})
	mediaconvertCmd.AddCommand(mediaconvert_getPresetCmd)
}
