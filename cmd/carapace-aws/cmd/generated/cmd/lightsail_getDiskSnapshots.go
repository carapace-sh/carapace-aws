package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lightsail_getDiskSnapshotsCmd = &cobra.Command{
	Use:   "get-disk-snapshots",
	Short: "Returns information about all block storage disk snapshots in your AWS account and region.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lightsail_getDiskSnapshotsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lightsail_getDiskSnapshotsCmd).Standalone()

		lightsail_getDiskSnapshotsCmd.Flags().String("page-token", "", "The token to advance to the next page of results from your request.")
	})
	lightsailCmd.AddCommand(lightsail_getDiskSnapshotsCmd)
}
