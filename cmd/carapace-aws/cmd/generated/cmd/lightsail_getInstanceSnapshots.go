package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lightsail_getInstanceSnapshotsCmd = &cobra.Command{
	Use:   "get-instance-snapshots",
	Short: "Returns all instance snapshots for the user's account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lightsail_getInstanceSnapshotsCmd).Standalone()

	lightsail_getInstanceSnapshotsCmd.Flags().String("page-token", "", "The token to advance to the next page of results from your request.")
	lightsailCmd.AddCommand(lightsail_getInstanceSnapshotsCmd)
}
