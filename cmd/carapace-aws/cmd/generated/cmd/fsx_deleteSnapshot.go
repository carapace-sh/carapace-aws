package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var fsx_deleteSnapshotCmd = &cobra.Command{
	Use:   "delete-snapshot",
	Short: "Deletes an Amazon FSx for OpenZFS snapshot.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(fsx_deleteSnapshotCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(fsx_deleteSnapshotCmd).Standalone()

		fsx_deleteSnapshotCmd.Flags().String("client-request-token", "", "")
		fsx_deleteSnapshotCmd.Flags().String("snapshot-id", "", "The ID of the snapshot that you want to delete.")
		fsx_deleteSnapshotCmd.MarkFlagRequired("snapshot-id")
	})
	fsxCmd.AddCommand(fsx_deleteSnapshotCmd)
}
