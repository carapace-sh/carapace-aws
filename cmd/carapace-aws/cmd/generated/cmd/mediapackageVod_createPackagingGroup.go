package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediapackageVod_createPackagingGroupCmd = &cobra.Command{
	Use:   "create-packaging-group",
	Short: "Creates a new MediaPackage VOD PackagingGroup resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediapackageVod_createPackagingGroupCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mediapackageVod_createPackagingGroupCmd).Standalone()

		mediapackageVod_createPackagingGroupCmd.Flags().String("authorization", "", "")
		mediapackageVod_createPackagingGroupCmd.Flags().String("egress-access-logs", "", "")
		mediapackageVod_createPackagingGroupCmd.Flags().String("id", "", "The ID of the PackagingGroup.")
		mediapackageVod_createPackagingGroupCmd.Flags().String("tags", "", "")
		mediapackageVod_createPackagingGroupCmd.MarkFlagRequired("id")
	})
	mediapackageVodCmd.AddCommand(mediapackageVod_createPackagingGroupCmd)
}
