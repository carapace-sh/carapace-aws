package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediaconvert_listPresetsCmd = &cobra.Command{
	Use:   "list-presets",
	Short: "Retrieve a JSON array of up to twenty of your presets.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediaconvert_listPresetsCmd).Standalone()

	mediaconvert_listPresetsCmd.Flags().String("category", "", "Optionally, specify a preset category to limit responses to only presets from that category.")
	mediaconvert_listPresetsCmd.Flags().String("list-by", "", "Optional.")
	mediaconvert_listPresetsCmd.Flags().String("max-results", "", "Optional.")
	mediaconvert_listPresetsCmd.Flags().String("next-token", "", "Use this string, provided with the response to a previous request, to request the next batch of presets.")
	mediaconvert_listPresetsCmd.Flags().String("order", "", "Optional.")
	mediaconvertCmd.AddCommand(mediaconvert_listPresetsCmd)
}
