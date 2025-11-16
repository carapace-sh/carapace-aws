package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediapackageVodCmd = &cobra.Command{
	Use:   "mediapackage-vod",
	Short: "AWS Elemental MediaPackage VOD",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediapackageVodCmd).Standalone()

	rootCmd.AddCommand(mediapackageVodCmd)
}
