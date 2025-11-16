package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ebs_completeSnapshotCmd = &cobra.Command{
	Use:   "complete-snapshot",
	Short: "Seals and completes the snapshot after all of the required blocks of data have been written to it.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ebs_completeSnapshotCmd).Standalone()

	ebs_completeSnapshotCmd.Flags().String("changed-blocks-count", "", "The number of blocks that were written to the snapshot.")
	ebs_completeSnapshotCmd.Flags().String("checksum", "", "An aggregated Base-64 SHA256 checksum based on the checksums of each written block.")
	ebs_completeSnapshotCmd.Flags().String("checksum-aggregation-method", "", "The aggregation method used to generate the checksum.")
	ebs_completeSnapshotCmd.Flags().String("checksum-algorithm", "", "The algorithm used to generate the checksum.")
	ebs_completeSnapshotCmd.Flags().String("snapshot-id", "", "The ID of the snapshot.")
	ebs_completeSnapshotCmd.MarkFlagRequired("changed-blocks-count")
	ebs_completeSnapshotCmd.MarkFlagRequired("snapshot-id")
	ebsCmd.AddCommand(ebs_completeSnapshotCmd)
}
