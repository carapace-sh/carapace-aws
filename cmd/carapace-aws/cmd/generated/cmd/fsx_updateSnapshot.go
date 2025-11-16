package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var fsx_updateSnapshotCmd = &cobra.Command{
	Use:   "update-snapshot",
	Short: "Updates the name of an Amazon FSx for OpenZFS snapshot.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(fsx_updateSnapshotCmd).Standalone()

	fsx_updateSnapshotCmd.Flags().String("client-request-token", "", "")
	fsx_updateSnapshotCmd.Flags().String("name", "", "The name of the snapshot to update.")
	fsx_updateSnapshotCmd.Flags().String("snapshot-id", "", "The ID of the snapshot that you want to update, in the format `fsvolsnap-0123456789abcdef0`.")
	fsx_updateSnapshotCmd.MarkFlagRequired("name")
	fsx_updateSnapshotCmd.MarkFlagRequired("snapshot-id")
	fsxCmd.AddCommand(fsx_updateSnapshotCmd)
}
