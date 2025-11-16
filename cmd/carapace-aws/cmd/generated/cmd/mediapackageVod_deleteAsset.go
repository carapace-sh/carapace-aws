package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediapackageVod_deleteAssetCmd = &cobra.Command{
	Use:   "delete-asset",
	Short: "Deletes an existing MediaPackage VOD Asset resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediapackageVod_deleteAssetCmd).Standalone()

	mediapackageVod_deleteAssetCmd.Flags().String("id", "", "The ID of the MediaPackage VOD Asset resource to delete.")
	mediapackageVod_deleteAssetCmd.MarkFlagRequired("id")
	mediapackageVodCmd.AddCommand(mediapackageVod_deleteAssetCmd)
}
