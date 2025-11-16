package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ebs_getSnapshotBlockCmd = &cobra.Command{
	Use:   "get-snapshot-block",
	Short: "Returns the data in a block in an Amazon Elastic Block Store snapshot.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ebs_getSnapshotBlockCmd).Standalone()

	ebs_getSnapshotBlockCmd.Flags().String("block-index", "", "The block index of the block in which to read the data.")
	ebs_getSnapshotBlockCmd.Flags().String("block-token", "", "The block token of the block from which to get data.")
	ebs_getSnapshotBlockCmd.Flags().String("snapshot-id", "", "The ID of the snapshot containing the block from which to get data.")
	ebs_getSnapshotBlockCmd.MarkFlagRequired("block-index")
	ebs_getSnapshotBlockCmd.MarkFlagRequired("block-token")
	ebs_getSnapshotBlockCmd.MarkFlagRequired("snapshot-id")
	ebsCmd.AddCommand(ebs_getSnapshotBlockCmd)
}
