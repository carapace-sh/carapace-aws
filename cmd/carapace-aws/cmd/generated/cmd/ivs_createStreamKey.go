package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ivs_createStreamKeyCmd = &cobra.Command{
	Use:   "create-stream-key",
	Short: "Creates a stream key, used to initiate a stream, for the specified channel ARN.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ivs_createStreamKeyCmd).Standalone()

	ivs_createStreamKeyCmd.Flags().String("channel-arn", "", "ARN of the channel for which to create the stream key.")
	ivs_createStreamKeyCmd.Flags().String("tags", "", "Array of 1-50 maps, each of the form `string:string (key:value)`.")
	ivs_createStreamKeyCmd.MarkFlagRequired("channel-arn")
	ivsCmd.AddCommand(ivs_createStreamKeyCmd)
}
