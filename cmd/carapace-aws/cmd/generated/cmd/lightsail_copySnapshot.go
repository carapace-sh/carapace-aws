package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lightsail_copySnapshotCmd = &cobra.Command{
	Use:   "copy-snapshot",
	Short: "Copies a manual snapshot of an instance or disk as another manual snapshot, or copies an automatic snapshot of an instance or disk as a manual snapshot.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lightsail_copySnapshotCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lightsail_copySnapshotCmd).Standalone()

		lightsail_copySnapshotCmd.Flags().String("restore-date", "", "The date of the source automatic snapshot to copy.")
		lightsail_copySnapshotCmd.Flags().String("source-region", "", "The Amazon Web Services Region where the source manual or automatic snapshot is located.")
		lightsail_copySnapshotCmd.Flags().String("source-resource-name", "", "The name of the source instance or disk from which the source automatic snapshot was created.")
		lightsail_copySnapshotCmd.Flags().String("source-snapshot-name", "", "The name of the source manual snapshot to copy.")
		lightsail_copySnapshotCmd.Flags().String("target-snapshot-name", "", "The name of the new manual snapshot to be created as a copy.")
		lightsail_copySnapshotCmd.Flags().String("use-latest-restorable-auto-snapshot", "", "A Boolean value to indicate whether to use the latest available automatic snapshot of the specified source instance or disk.")
		lightsail_copySnapshotCmd.MarkFlagRequired("source-region")
		lightsail_copySnapshotCmd.MarkFlagRequired("target-snapshot-name")
	})
	lightsailCmd.AddCommand(lightsail_copySnapshotCmd)
}
