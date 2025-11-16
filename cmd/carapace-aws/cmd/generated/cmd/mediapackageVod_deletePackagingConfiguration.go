package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediapackageVod_deletePackagingConfigurationCmd = &cobra.Command{
	Use:   "delete-packaging-configuration",
	Short: "Deletes a MediaPackage VOD PackagingConfiguration resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediapackageVod_deletePackagingConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mediapackageVod_deletePackagingConfigurationCmd).Standalone()

		mediapackageVod_deletePackagingConfigurationCmd.Flags().String("id", "", "The ID of the MediaPackage VOD PackagingConfiguration resource to delete.")
		mediapackageVod_deletePackagingConfigurationCmd.MarkFlagRequired("id")
	})
	mediapackageVodCmd.AddCommand(mediapackageVod_deletePackagingConfigurationCmd)
}
