package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediapackageVod_createPackagingConfigurationCmd = &cobra.Command{
	Use:   "create-packaging-configuration",
	Short: "Creates a new MediaPackage VOD PackagingConfiguration resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediapackageVod_createPackagingConfigurationCmd).Standalone()

	mediapackageVod_createPackagingConfigurationCmd.Flags().String("cmaf-package", "", "")
	mediapackageVod_createPackagingConfigurationCmd.Flags().String("dash-package", "", "")
	mediapackageVod_createPackagingConfigurationCmd.Flags().String("hls-package", "", "")
	mediapackageVod_createPackagingConfigurationCmd.Flags().String("id", "", "The ID of the PackagingConfiguration.")
	mediapackageVod_createPackagingConfigurationCmd.Flags().String("mss-package", "", "")
	mediapackageVod_createPackagingConfigurationCmd.Flags().String("packaging-group-id", "", "The ID of a PackagingGroup.")
	mediapackageVod_createPackagingConfigurationCmd.Flags().String("tags", "", "")
	mediapackageVod_createPackagingConfigurationCmd.MarkFlagRequired("id")
	mediapackageVod_createPackagingConfigurationCmd.MarkFlagRequired("packaging-group-id")
	mediapackageVodCmd.AddCommand(mediapackageVod_createPackagingConfigurationCmd)
}
