package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediapackageVod_listPackagingConfigurationsCmd = &cobra.Command{
	Use:   "list-packaging-configurations",
	Short: "Returns a collection of MediaPackage VOD PackagingConfiguration resources.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediapackageVod_listPackagingConfigurationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mediapackageVod_listPackagingConfigurationsCmd).Standalone()

		mediapackageVod_listPackagingConfigurationsCmd.Flags().String("max-results", "", "Upper bound on number of records to return.")
		mediapackageVod_listPackagingConfigurationsCmd.Flags().String("next-token", "", "A token used to resume pagination from the end of a previous request.")
		mediapackageVod_listPackagingConfigurationsCmd.Flags().String("packaging-group-id", "", "Returns MediaPackage VOD PackagingConfigurations associated with the specified PackagingGroup.")
	})
	mediapackageVodCmd.AddCommand(mediapackageVod_listPackagingConfigurationsCmd)
}
