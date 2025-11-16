package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ivs_batchGetStreamKeyCmd = &cobra.Command{
	Use:   "batch-get-stream-key",
	Short: "Performs [GetStreamKey]() on multiple ARNs simultaneously.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ivs_batchGetStreamKeyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ivs_batchGetStreamKeyCmd).Standalone()

		ivs_batchGetStreamKeyCmd.Flags().String("arns", "", "Array of ARNs, one per stream key.")
		ivs_batchGetStreamKeyCmd.MarkFlagRequired("arns")
	})
	ivsCmd.AddCommand(ivs_batchGetStreamKeyCmd)
}
