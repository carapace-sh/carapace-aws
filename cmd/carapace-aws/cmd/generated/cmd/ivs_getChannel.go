package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ivs_getChannelCmd = &cobra.Command{
	Use:   "get-channel",
	Short: "Gets the channel configuration for the specified channel ARN.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ivs_getChannelCmd).Standalone()

	ivs_getChannelCmd.Flags().String("arn", "", "ARN of the channel for which the configuration is to be retrieved.")
	ivs_getChannelCmd.MarkFlagRequired("arn")
	ivsCmd.AddCommand(ivs_getChannelCmd)
}
