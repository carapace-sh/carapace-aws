package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ds_createSnapshotCmd = &cobra.Command{
	Use:   "create-snapshot",
	Short: "Creates a snapshot of a Simple AD or Microsoft AD directory in the Amazon Web Services cloud.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ds_createSnapshotCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ds_createSnapshotCmd).Standalone()

		ds_createSnapshotCmd.Flags().String("directory-id", "", "The identifier of the directory of which to take a snapshot.")
		ds_createSnapshotCmd.Flags().String("name", "", "The descriptive name to apply to the snapshot.")
		ds_createSnapshotCmd.MarkFlagRequired("directory-id")
	})
	dsCmd.AddCommand(ds_createSnapshotCmd)
}
