package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ivs_listStreamKeysCmd = &cobra.Command{
	Use:   "list-stream-keys",
	Short: "Gets summary information about stream keys for the specified channel.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ivs_listStreamKeysCmd).Standalone()

	ivs_listStreamKeysCmd.Flags().String("channel-arn", "", "Channel ARN used to filter the list.")
	ivs_listStreamKeysCmd.Flags().String("max-results", "", "Maximum number of streamKeys to return.")
	ivs_listStreamKeysCmd.Flags().String("next-token", "", "The first stream key to retrieve.")
	ivs_listStreamKeysCmd.MarkFlagRequired("channel-arn")
	ivsCmd.AddCommand(ivs_listStreamKeysCmd)
}
