package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediapackageVod_updatePackagingGroupCmd = &cobra.Command{
	Use:   "update-packaging-group",
	Short: "Updates a specific packaging group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediapackageVod_updatePackagingGroupCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mediapackageVod_updatePackagingGroupCmd).Standalone()

		mediapackageVod_updatePackagingGroupCmd.Flags().String("authorization", "", "")
		mediapackageVod_updatePackagingGroupCmd.Flags().String("id", "", "The ID of a MediaPackage VOD PackagingGroup resource.")
		mediapackageVod_updatePackagingGroupCmd.MarkFlagRequired("id")
	})
	mediapackageVodCmd.AddCommand(mediapackageVod_updatePackagingGroupCmd)
}
