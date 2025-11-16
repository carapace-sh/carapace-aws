package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediapackageVod_describeAssetCmd = &cobra.Command{
	Use:   "describe-asset",
	Short: "Returns a description of a MediaPackage VOD Asset resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediapackageVod_describeAssetCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mediapackageVod_describeAssetCmd).Standalone()

		mediapackageVod_describeAssetCmd.Flags().String("id", "", "The ID of an MediaPackage VOD Asset resource.")
		mediapackageVod_describeAssetCmd.MarkFlagRequired("id")
	})
	mediapackageVodCmd.AddCommand(mediapackageVod_describeAssetCmd)
}
