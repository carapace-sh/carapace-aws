package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ivs_batchGetChannelCmd = &cobra.Command{
	Use:   "batch-get-channel",
	Short: "Performs [GetChannel]() on multiple ARNs simultaneously.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ivs_batchGetChannelCmd).Standalone()

	ivs_batchGetChannelCmd.Flags().String("arns", "", "Array of ARNs, one per channel.")
	ivs_batchGetChannelCmd.MarkFlagRequired("arns")
	ivsCmd.AddCommand(ivs_batchGetChannelCmd)
}
