package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lightsail_getInstanceSnapshotCmd = &cobra.Command{
	Use:   "get-instance-snapshot",
	Short: "Returns information about a specific instance snapshot.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lightsail_getInstanceSnapshotCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lightsail_getInstanceSnapshotCmd).Standalone()

		lightsail_getInstanceSnapshotCmd.Flags().String("instance-snapshot-name", "", "The name of the snapshot for which you are requesting information.")
		lightsail_getInstanceSnapshotCmd.MarkFlagRequired("instance-snapshot-name")
	})
	lightsailCmd.AddCommand(lightsail_getInstanceSnapshotCmd)
}
