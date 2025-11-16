package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ebs_putSnapshotBlockCmd = &cobra.Command{
	Use:   "put-snapshot-block",
	Short: "Writes a block of data to a snapshot.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ebs_putSnapshotBlockCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ebs_putSnapshotBlockCmd).Standalone()

		ebs_putSnapshotBlockCmd.Flags().String("block-data", "", "The data to write to the block.")
		ebs_putSnapshotBlockCmd.Flags().String("block-index", "", "The block index of the block in which to write the data.")
		ebs_putSnapshotBlockCmd.Flags().String("checksum", "", "A Base64-encoded SHA256 checksum of the data.")
		ebs_putSnapshotBlockCmd.Flags().String("checksum-algorithm", "", "The algorithm used to generate the checksum.")
		ebs_putSnapshotBlockCmd.Flags().String("data-length", "", "The size of the data to write to the block, in bytes.")
		ebs_putSnapshotBlockCmd.Flags().String("progress", "", "The progress of the write process, as a percentage.")
		ebs_putSnapshotBlockCmd.Flags().String("snapshot-id", "", "The ID of the snapshot.")
		ebs_putSnapshotBlockCmd.MarkFlagRequired("block-data")
		ebs_putSnapshotBlockCmd.MarkFlagRequired("block-index")
		ebs_putSnapshotBlockCmd.MarkFlagRequired("checksum")
		ebs_putSnapshotBlockCmd.MarkFlagRequired("checksum-algorithm")
		ebs_putSnapshotBlockCmd.MarkFlagRequired("data-length")
		ebs_putSnapshotBlockCmd.MarkFlagRequired("snapshot-id")
	})
	ebsCmd.AddCommand(ebs_putSnapshotBlockCmd)
}
