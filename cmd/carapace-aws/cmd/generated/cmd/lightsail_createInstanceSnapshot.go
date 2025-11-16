package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lightsail_createInstanceSnapshotCmd = &cobra.Command{
	Use:   "create-instance-snapshot",
	Short: "Creates a snapshot of a specific virtual private server, or *instance*.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lightsail_createInstanceSnapshotCmd).Standalone()

	lightsail_createInstanceSnapshotCmd.Flags().String("instance-name", "", "The Lightsail instance on which to base your snapshot.")
	lightsail_createInstanceSnapshotCmd.Flags().String("instance-snapshot-name", "", "The name for your new snapshot.")
	lightsail_createInstanceSnapshotCmd.Flags().String("tags", "", "The tag keys and optional values to add to the resource during create.")
	lightsail_createInstanceSnapshotCmd.MarkFlagRequired("instance-name")
	lightsail_createInstanceSnapshotCmd.MarkFlagRequired("instance-snapshot-name")
	lightsailCmd.AddCommand(lightsail_createInstanceSnapshotCmd)
}
