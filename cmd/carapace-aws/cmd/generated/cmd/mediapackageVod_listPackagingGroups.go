package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediapackageVod_listPackagingGroupsCmd = &cobra.Command{
	Use:   "list-packaging-groups",
	Short: "Returns a collection of MediaPackage VOD PackagingGroup resources.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediapackageVod_listPackagingGroupsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mediapackageVod_listPackagingGroupsCmd).Standalone()

		mediapackageVod_listPackagingGroupsCmd.Flags().String("max-results", "", "Upper bound on number of records to return.")
		mediapackageVod_listPackagingGroupsCmd.Flags().String("next-token", "", "A token used to resume pagination from the end of a previous request.")
	})
	mediapackageVodCmd.AddCommand(mediapackageVod_listPackagingGroupsCmd)
}
