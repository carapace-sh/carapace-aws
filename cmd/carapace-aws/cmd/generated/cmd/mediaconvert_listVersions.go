package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediaconvert_listVersionsCmd = &cobra.Command{
	Use:   "list-versions",
	Short: "Retrieve a JSON array of all available Job engine versions and the date they expire.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediaconvert_listVersionsCmd).Standalone()

	mediaconvert_listVersionsCmd.Flags().String("max-results", "", "Optional.")
	mediaconvert_listVersionsCmd.Flags().String("next-token", "", "Optional.")
	mediaconvertCmd.AddCommand(mediaconvert_listVersionsCmd)
}
