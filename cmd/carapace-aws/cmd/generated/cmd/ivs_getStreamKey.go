package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ivs_getStreamKeyCmd = &cobra.Command{
	Use:   "get-stream-key",
	Short: "Gets stream-key information for a specified ARN.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ivs_getStreamKeyCmd).Standalone()

	ivs_getStreamKeyCmd.Flags().String("arn", "", "ARN for the stream key to be retrieved.")
	ivs_getStreamKeyCmd.MarkFlagRequired("arn")
	ivsCmd.AddCommand(ivs_getStreamKeyCmd)
}
