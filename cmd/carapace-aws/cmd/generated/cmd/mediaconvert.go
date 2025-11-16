package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediaconvertCmd = &cobra.Command{
	Use:   "mediaconvert",
	Short: "AWS Elemental MediaConvert",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediaconvertCmd).Standalone()

	rootCmd.AddCommand(mediaconvertCmd)
}
