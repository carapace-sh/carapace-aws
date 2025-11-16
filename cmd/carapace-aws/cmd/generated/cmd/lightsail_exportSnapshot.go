package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lightsail_exportSnapshotCmd = &cobra.Command{
	Use:   "export-snapshot",
	Short: "Exports an Amazon Lightsail instance or block storage disk snapshot to Amazon Elastic Compute Cloud (Amazon EC2).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lightsail_exportSnapshotCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lightsail_exportSnapshotCmd).Standalone()

		lightsail_exportSnapshotCmd.Flags().String("source-snapshot-name", "", "The name of the instance or disk snapshot to be exported to Amazon EC2.")
		lightsail_exportSnapshotCmd.MarkFlagRequired("source-snapshot-name")
	})
	lightsailCmd.AddCommand(lightsail_exportSnapshotCmd)
}
