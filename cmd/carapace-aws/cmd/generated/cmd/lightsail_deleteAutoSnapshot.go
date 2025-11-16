package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lightsail_deleteAutoSnapshotCmd = &cobra.Command{
	Use:   "delete-auto-snapshot",
	Short: "Deletes an automatic snapshot of an instance or disk.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lightsail_deleteAutoSnapshotCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lightsail_deleteAutoSnapshotCmd).Standalone()

		lightsail_deleteAutoSnapshotCmd.Flags().String("date", "", "The date of the automatic snapshot to delete in `YYYY-MM-DD` format.")
		lightsail_deleteAutoSnapshotCmd.Flags().String("resource-name", "", "The name of the source instance or disk from which to delete the automatic snapshot.")
		lightsail_deleteAutoSnapshotCmd.MarkFlagRequired("date")
		lightsail_deleteAutoSnapshotCmd.MarkFlagRequired("resource-name")
	})
	lightsailCmd.AddCommand(lightsail_deleteAutoSnapshotCmd)
}
