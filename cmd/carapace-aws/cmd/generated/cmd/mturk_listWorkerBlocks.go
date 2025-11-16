package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mturk_listWorkerBlocksCmd = &cobra.Command{
	Use:   "list-worker-blocks",
	Short: "The `ListWorkersBlocks` operation retrieves a list of Workers who are blocked from working on your HITs.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mturk_listWorkerBlocksCmd).Standalone()

	mturk_listWorkerBlocksCmd.Flags().String("max-results", "", "")
	mturk_listWorkerBlocksCmd.Flags().String("next-token", "", "Pagination token")
	mturkCmd.AddCommand(mturk_listWorkerBlocksCmd)
}
