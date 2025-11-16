package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var fsx_createSnapshotCmd = &cobra.Command{
	Use:   "create-snapshot",
	Short: "Creates a snapshot of an existing Amazon FSx for OpenZFS volume.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(fsx_createSnapshotCmd).Standalone()

	fsx_createSnapshotCmd.Flags().String("client-request-token", "", "")
	fsx_createSnapshotCmd.Flags().String("name", "", "The name of the snapshot.")
	fsx_createSnapshotCmd.Flags().String("tags", "", "")
	fsx_createSnapshotCmd.Flags().String("volume-id", "", "The ID of the volume that you are taking a snapshot of.")
	fsx_createSnapshotCmd.MarkFlagRequired("name")
	fsx_createSnapshotCmd.MarkFlagRequired("volume-id")
	fsxCmd.AddCommand(fsx_createSnapshotCmd)
}
