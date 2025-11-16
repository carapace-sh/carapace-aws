package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediapackageVod_describePackagingGroupCmd = &cobra.Command{
	Use:   "describe-packaging-group",
	Short: "Returns a description of a MediaPackage VOD PackagingGroup resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediapackageVod_describePackagingGroupCmd).Standalone()

	mediapackageVod_describePackagingGroupCmd.Flags().String("id", "", "The ID of a MediaPackage VOD PackagingGroup resource.")
	mediapackageVod_describePackagingGroupCmd.MarkFlagRequired("id")
	mediapackageVodCmd.AddCommand(mediapackageVod_describePackagingGroupCmd)
}
