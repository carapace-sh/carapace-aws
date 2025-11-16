package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediaconvert_deletePresetCmd = &cobra.Command{
	Use:   "delete-preset",
	Short: "Permanently delete a preset you have created.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediaconvert_deletePresetCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mediaconvert_deletePresetCmd).Standalone()

		mediaconvert_deletePresetCmd.Flags().String("name", "", "The name of the preset to be deleted.")
		mediaconvert_deletePresetCmd.MarkFlagRequired("name")
	})
	mediaconvertCmd.AddCommand(mediaconvert_deletePresetCmd)
}
