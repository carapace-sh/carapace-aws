package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ebs_listSnapshotBlocksCmd = &cobra.Command{
	Use:   "list-snapshot-blocks",
	Short: "Returns information about the blocks in an Amazon Elastic Block Store snapshot.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ebs_listSnapshotBlocksCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ebs_listSnapshotBlocksCmd).Standalone()

		ebs_listSnapshotBlocksCmd.Flags().String("max-results", "", "The maximum number of blocks to be returned by the request.")
		ebs_listSnapshotBlocksCmd.Flags().String("next-token", "", "The token to request the next page of results.")
		ebs_listSnapshotBlocksCmd.Flags().String("snapshot-id", "", "The ID of the snapshot from which to get block indexes and block tokens.")
		ebs_listSnapshotBlocksCmd.Flags().String("starting-block-index", "", "The block index from which the list should start.")
		ebs_listSnapshotBlocksCmd.MarkFlagRequired("snapshot-id")
	})
	ebsCmd.AddCommand(ebs_listSnapshotBlocksCmd)
}
