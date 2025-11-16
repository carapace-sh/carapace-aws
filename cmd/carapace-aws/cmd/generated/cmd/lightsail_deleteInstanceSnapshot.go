package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lightsail_deleteInstanceSnapshotCmd = &cobra.Command{
	Use:   "delete-instance-snapshot",
	Short: "Deletes a specific snapshot of a virtual private server (or *instance*).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lightsail_deleteInstanceSnapshotCmd).Standalone()

	lightsail_deleteInstanceSnapshotCmd.Flags().String("instance-snapshot-name", "", "The name of the snapshot to delete.")
	lightsail_deleteInstanceSnapshotCmd.MarkFlagRequired("instance-snapshot-name")
	lightsailCmd.AddCommand(lightsail_deleteInstanceSnapshotCmd)
}
