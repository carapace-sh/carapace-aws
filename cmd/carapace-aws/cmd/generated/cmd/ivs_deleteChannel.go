package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ivs_deleteChannelCmd = &cobra.Command{
	Use:   "delete-channel",
	Short: "Deletes the specified channel and its associated stream keys.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ivs_deleteChannelCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ivs_deleteChannelCmd).Standalone()

		ivs_deleteChannelCmd.Flags().String("arn", "", "ARN of the channel to be deleted.")
		ivs_deleteChannelCmd.MarkFlagRequired("arn")
	})
	ivsCmd.AddCommand(ivs_deleteChannelCmd)
}
