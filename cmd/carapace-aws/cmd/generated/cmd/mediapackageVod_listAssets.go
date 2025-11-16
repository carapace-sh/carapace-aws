package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediapackageVod_listAssetsCmd = &cobra.Command{
	Use:   "list-assets",
	Short: "Returns a collection of MediaPackage VOD Asset resources.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediapackageVod_listAssetsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mediapackageVod_listAssetsCmd).Standalone()

		mediapackageVod_listAssetsCmd.Flags().String("max-results", "", "Upper bound on number of records to return.")
		mediapackageVod_listAssetsCmd.Flags().String("next-token", "", "A token used to resume pagination from the end of a previous request.")
		mediapackageVod_listAssetsCmd.Flags().String("packaging-group-id", "", "Returns Assets associated with the specified PackagingGroup.")
	})
	mediapackageVodCmd.AddCommand(mediapackageVod_listAssetsCmd)
}
