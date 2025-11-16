package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ebs_listChangedBlocksCmd = &cobra.Command{
	Use:   "list-changed-blocks",
	Short: "Returns information about the blocks that are different between two Amazon Elastic Block Store snapshots of the same volume/snapshot lineage.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ebs_listChangedBlocksCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ebs_listChangedBlocksCmd).Standalone()

		ebs_listChangedBlocksCmd.Flags().String("first-snapshot-id", "", "The ID of the first snapshot to use for the comparison.")
		ebs_listChangedBlocksCmd.Flags().String("max-results", "", "The maximum number of blocks to be returned by the request.")
		ebs_listChangedBlocksCmd.Flags().String("next-token", "", "The token to request the next page of results.")
		ebs_listChangedBlocksCmd.Flags().String("second-snapshot-id", "", "The ID of the second snapshot to use for the comparison.")
		ebs_listChangedBlocksCmd.Flags().String("starting-block-index", "", "The block index from which the comparison should start.")
		ebs_listChangedBlocksCmd.MarkFlagRequired("second-snapshot-id")
	})
	ebsCmd.AddCommand(ebs_listChangedBlocksCmd)
}
