package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediapackageVod_deletePackagingGroupCmd = &cobra.Command{
	Use:   "delete-packaging-group",
	Short: "Deletes a MediaPackage VOD PackagingGroup resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediapackageVod_deletePackagingGroupCmd).Standalone()

	mediapackageVod_deletePackagingGroupCmd.Flags().String("id", "", "The ID of the MediaPackage VOD PackagingGroup resource to delete.")
	mediapackageVod_deletePackagingGroupCmd.MarkFlagRequired("id")
	mediapackageVodCmd.AddCommand(mediapackageVod_deletePackagingGroupCmd)
}
