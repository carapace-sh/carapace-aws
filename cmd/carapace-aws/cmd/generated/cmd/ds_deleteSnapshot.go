package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ds_deleteSnapshotCmd = &cobra.Command{
	Use:   "delete-snapshot",
	Short: "Deletes a directory snapshot.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ds_deleteSnapshotCmd).Standalone()

	ds_deleteSnapshotCmd.Flags().String("snapshot-id", "", "The identifier of the directory snapshot to be deleted.")
	ds_deleteSnapshotCmd.MarkFlagRequired("snapshot-id")
	dsCmd.AddCommand(ds_deleteSnapshotCmd)
}
