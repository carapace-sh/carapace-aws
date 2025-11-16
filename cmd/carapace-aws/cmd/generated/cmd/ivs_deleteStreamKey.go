package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ivs_deleteStreamKeyCmd = &cobra.Command{
	Use:   "delete-stream-key",
	Short: "Deletes the stream key for the specified ARN, so it can no longer be used to stream.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ivs_deleteStreamKeyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ivs_deleteStreamKeyCmd).Standalone()

		ivs_deleteStreamKeyCmd.Flags().String("arn", "", "ARN of the stream key to be deleted.")
		ivs_deleteStreamKeyCmd.MarkFlagRequired("arn")
	})
	ivsCmd.AddCommand(ivs_deleteStreamKeyCmd)
}
