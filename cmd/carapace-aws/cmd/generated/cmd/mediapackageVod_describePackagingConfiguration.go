package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediapackageVod_describePackagingConfigurationCmd = &cobra.Command{
	Use:   "describe-packaging-configuration",
	Short: "Returns a description of a MediaPackage VOD PackagingConfiguration resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediapackageVod_describePackagingConfigurationCmd).Standalone()

	mediapackageVod_describePackagingConfigurationCmd.Flags().String("id", "", "The ID of a MediaPackage VOD PackagingConfiguration resource.")
	mediapackageVod_describePackagingConfigurationCmd.MarkFlagRequired("id")
	mediapackageVodCmd.AddCommand(mediapackageVod_describePackagingConfigurationCmd)
}
