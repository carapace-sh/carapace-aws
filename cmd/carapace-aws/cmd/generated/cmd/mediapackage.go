package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediapackageCmd = &cobra.Command{
	Use:   "mediapackage",
	Short: "AWS Elemental MediaPackage",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediapackageCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mediapackageCmd).Standalone()

	})
	rootCmd.AddCommand(mediapackageCmd)
}
